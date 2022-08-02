package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestThumbnailAPI(t *testing.T) {
    InitializeTester[*onshape.ThumbnailApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentAndVersionRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetDocumentThumbnailRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetElementThumbnailWithApiConfigurationRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetDocumentThumbnailWithSizeRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetElementThumbnailRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiSetApplicationElementThumbnailRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDeleteApplicationThumbnailsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetElementThumbnailWithSizeRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentOldRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentAndVersionOldRequest{},
        Expect: Todo(),
    }.Execute()
    
}