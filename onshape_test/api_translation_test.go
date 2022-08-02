package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestTranslationAPI(t *testing.T) {
    InitializeTester[*onshape.TranslationApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetDocumentTranslationsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateTranslationRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetAllTranslatorFormatsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetTranslationRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDeleteTranslationRequest{},
        Expect: Todo(),
    }.Execute()
    
}