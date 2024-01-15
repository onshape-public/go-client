package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/onshape-public/go-client/onshape_test/testhelper"
)

func TestElementAPI(t *testing.T) {
	InitializeTester[*onshape.ElementApiService](t)

	SetContext(TestingContext{
		"wvm": "w",
	}.InheritDefaults(Context()))

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

	OpenAPITest{
		Call: onshape.ApiCreateElementRequest{
			ApiService: Context()["client"].(*onshape.APIClient).AppElementApi,
		}.BTAppElementParams(*GetDefaultAppElementParams()),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiDeleteElementRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiCopyElementFromSourceDocumentRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiEncodeConfigurationMapRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiDeleteElementRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateReferencesRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetConfigurationRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateConfigurationRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiDecodeConfigurationRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest{},
		Expect: Todo(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDeleteDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()
}
