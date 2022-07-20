package onshape_test

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

func initializeEventTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.EventApi,
	}
}

func TestEventAPI(t *testing.T) {
	ctx := CreateDocumentPreTest(initializeEventTests(t).SetDefault("label", Ptr("test"+uuid.NewString())), t)
	ctx = ctx.SetDefault("bTEventParams", Ptr(onshape.BTDocumentOpenEventParams{
		JsonType:   Ptr("DocumentOpenEventInfo"),
		DocumentId: Ptr(ctx["did"].(string)),
	}).AsBTEventParams())

	//First wait for 3 sec to wait for doc to settle...
	time.Sleep(3 * time.Second)

	doc, err := requestRecentlyOpenedDocument()
	if doc != nil {
		assert.NoError(t, err)
		assert.NotEqual(t, ctx["did"], doc.GetId())
	}

	OpenAPITest{
		Call:   onshape.ApiFireEventRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)

	doc, err = requestRecentlyOpenedDocument()
	assert.NoError(t, err)
	assert.Equal(t, ctx["did"], doc.GetId())

	DeleteDocumentPostTest(ctx, t)
}

func requestRecentlyOpenedDocument() (*onshape.BTDocumentSummaryInfo, error) {
	doc, _, err := requestMagicTreeNodes()

	if err != nil {
		return nil, err
	}

	for _, x := range doc {
		if x.GetName() == "Recently opened" {
			res, _, err := requestMagicTreeNode(x.GetTreeHref())

			if err != nil {
				return nil, err
			}

			if len(res) == 0 {
				return nil, nil
			}

			return &res[0], err
		}
	}

	return nil, errors.New("No recentlyOpened document found")
}

func requestMagicTreeNodes() ([]onshape.BTGlobalTreeMagicNodeInfo, *http.Response, error) {
	localVarBody, localVarHTTPResponse, err := makeHTTPRequest(getBaseURL() + "/globaltreenodes/")

	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	type bTMagicResponse struct {
		Items []onshape.BTGlobalTreeMagicNodeInfo `json:"items"`
	}

	var nodes bTMagicResponse

	err = json.Unmarshal(localVarBody, &nodes)

	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	return nodes.Items, localVarHTTPResponse, err
}

func requestMagicTreeNode(url string) ([]onshape.BTDocumentSummaryInfo, *http.Response, error) {
	localVarBody, localVarHTTPResponse, err := makeHTTPRequest(url)

	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	type bTMagicResponse struct {
		Items []onshape.BTDocumentSummaryInfo `json:"items"`
	}

	var nodes bTMagicResponse

	err = json.Unmarshal(localVarBody, &nodes)

	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	return nodes.Items, localVarHTTPResponse, err
}

func makeHTTPRequest(url string) ([]byte, *http.Response, error) {
	ht := &http.Client{Transport: onshape.APIKeysRoundTripper{Proxied: http.DefaultTransport, HdrProcessors: []onshape.HdrProcFunc{onshape.GetAddAPIKeysReqHdrsFunc()}}}

	localVarRequest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	localVarRequest.Header.Add("accept", "application/json;charset=UTF-8; qs=0.09")

	localVarHTTPResponse, err := ht.Do(localVarRequest)
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()

	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	return localVarBody, nil, nil
}

func getBaseURL() string {
	return onshape.GetDefaultConfig().Servers[0].URL
}
