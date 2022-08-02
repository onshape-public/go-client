package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestBlobElementAPI(t *testing.T) {
    InitializeTester[*onshape.BlobElementApiService](t)

    OpenAPITest{
        Call: onshape.ApiUploadFileCreateElementRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDownloadFileWorkspaceRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUploadFileUpdateElementRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateUnitsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateBlobTranslationRequest{},
        Expect: Todo(),
    }.Execute()
    
}