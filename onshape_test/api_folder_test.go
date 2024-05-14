package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestFolderAPI(t *testing.T) {
    InitializeTester[*onshape.FolderAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetFolderAclRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiShareRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUnShareRequest{},
        Expect: Todo(),
    }.Execute()
    
}