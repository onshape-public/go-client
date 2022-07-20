package onshape_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

type WVMetadataProperty struct {
	Value      interface{} `json:"value"`
	PropertyID string      `json:"propertyId"`
}
type WVMetadataItem struct {
	Href       string               `json:"href"`
	Properties []WVMetadataProperty `json:"properties"`
}
type WVMetadataItems struct {
	Items []WVMetadataItem `json:"items"`
}

type WVMetadataPropertyResponse struct {
	ErrorMessage string `json:"errorMessage"`
	PropertyID   string `json:"propertyId"`
	Status       string `json:"status"`
}

type WVMetadataResponse struct {
	Properties []WVMetadataPropertyResponse `json:"properties,omitempty"`
	Href       string                       `json:"href,omitempty"`
	Status     string                       `json:"status,omitempty"`
}

func initializeMetadataTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.MetadataApi,
	}
}

func TestMetadataAPI(t *testing.T) {
	ctx := initializeMetadataTests(t)
	ctx = CreateDocumentPreTest(ctx, t)
	ctx = ctx.SetDefault("wv", "w").SetDefault("wvid", ctx["wid"])

	var href string
	OpenAPITest{
		Call: onshape.ApiGetWVMetadataRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTMetadataObjectInfo, v *http.Response, err error) {
			href = v.Request.URL.String()
		},
	}.Execute(ctx, t)

	it := WVMetadataItem{Href: href, Properties: []WVMetadataProperty{
		{
			PropertyID: "57f3fb8efa3416c06701d60d", // id of Name property
			Value:      "test",
		},
	}}
	its := WVMetadataItems{Items: []WVMetadataItem{it}}
	body, err := json.Marshal(its)
	assert.NoError(t, err)

	ctx = ctx.SetDefault("body", Ptr(string(body)))

	OpenAPITest{
		Call:   onshape.ApiUpdateWVMetadataRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetWVMetadataRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTMetadataObjectInfo, v *http.Response, err error) {
			assert.NoError(t, err)

			// Check all included properties are updated
			co := 0
			for _, prop := range r.GetProperties() {
				for _, x := range it.Properties {
					if prop.GetPropertyId() == x.PropertyID {
						assert.Equal(t, prop.Value, x.Value)
						co++
					}
				}
			}
			assert.Len(t, it.Properties, co)
		},
	}.Execute(ctx, t)
}

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiGetWMVEsMetadataRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetWMVEMetadataRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiUpdateWVEMetadataRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetWMVEPsMetadataRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetWMVEPMetadataRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiUpdateWVEPMetadataRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetVEOPStandardContentMetadataRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiUpdateVEOPStandardContentPartMetadataRequest{},
    Expect: Todo,
},

***/
