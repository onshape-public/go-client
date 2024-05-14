package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestDrawingAPI(t *testing.T) {
    InitializeTester[*onshape.DrawingAPIService](t)

    OpenAPITest{
        Call: onshape.ApiCreateDrawingAppElementRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetDrawingTranslatorFormatsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateDrawingTranslationRequest{},
        Expect: Todo(),
    }.Execute()
    
}