package onshape_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/require"
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

func TestMetadataAPI(t *testing.T) {
	InitializeTester[*onshape.MetadataApiService](t)

	SetContext(TestingContext{
		"wv": "w",
	}.InheritDefaults(Context()))

	OpenAPITest{
		Call: onshape.ApiCreateDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		}.BTDocumentParams(onshape.BTDocumentParams{
			Name:        "test-doc",
			Description: "This is a test document",
			IsPublic:    Ptr(false),
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetWVMetadataRequest{},
		Expect: func(r interface{}, v *http.Response, err error) {
			Context()["href"] = v.Request.URL.String()
		},
	}.Execute()

	it := WVMetadataItem{Href: Context()["href"].(string), Properties: []WVMetadataProperty{
		{
			PropertyID: "57f3fb8efa3416c06701d60d", // id of Name property
			Value:      "test",
		},
	}}
	its := WVMetadataItems{Items: []WVMetadataItem{it}}
	body, err := json.Marshal(its)
	require.NoError(Tester(), err)

	Context()["body"] = Ptr(string(body))

	OpenAPITest{
		Call:   onshape.ApiUpdateWVMetadataRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetWVMetadataRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTMetadataObjectInfo) {
			// Check all included properties are updated
			co := 0
			for _, prop := range r.GetProperties() {
				for _, x := range it.Properties {
					if prop.GetPropertyId() == x.PropertyID {
						require.Equal(Tester(), prop.Value, x.Value)
						co++
					}
				}
			}
			require.Len(Tester(), it.Properties, co)
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetWMVEsMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetWMVEMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateWVEMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetWMVEPsMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetWMVEPMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateWVEPMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetVEOPStandardContentMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateVEOPStandardContentPartMetadataRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDeleteDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()
}
