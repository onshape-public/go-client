package onshape_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/onshape-public/go-client/onshape"
	"github.com/onshape-public/go-client/onshape_test/testhelper"
	"github.com/stretchr/testify/require"
)

func TestEventAPI(t *testing.T) {
	InitializeTester[*onshape.EventApiService](t)

	OpenAPITest{
		Call: onshape.ApiCreateDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		}.BTDocumentParams(onshape.BTDocumentParams{
			Name:        "test-doc",
			Description: &testhelper.DocumentDescription,
			IsPublic:    Ptr(false),
		}),
		Expect: NoAPIError(),
	}.Execute()

	//First wait for 3 sec to wait for doc to settle...
	time.Sleep(3 * time.Second)

	doc, err := requestRecentlyOpenedDocument()
	if doc != nil {
		require.NoError(Tester(), err)
		require.NotEqual(t, Context()["did"], doc.GetId())
	}

	OpenAPITest{
		Call: onshape.ApiFireEventRequest{}.BTEventParams(*Ptr(onshape.BTDocumentOpenEventParams{
			JsonType:   Ptr("DocumentOpenEventInfo"),
			DocumentId: Ptr(Context()["did"].(string)),
		}).AsBTEventParams()),
		Expect: NoAPIError(),
	}.Execute()

	//wait a bit again to let the event to percolate
	time.Sleep(3 * time.Second)

	doc, err = requestRecentlyOpenedDocument()
	require.NoError(t, err)
	require.Equal(t, Context()["did"], doc.GetId())

	OpenAPITest{
		Call: onshape.ApiDeleteDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()

	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	return localVarBody, nil, nil
}

func getBaseURL() string {
	return onshape.GetDefaultConfig().Servers[0].URL
}
