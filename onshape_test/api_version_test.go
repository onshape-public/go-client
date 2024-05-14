package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestVersionAPI(t *testing.T) {
    InitializeTester[*onshape.VersionAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetAllVersionsRequest{},
        Expect: Todo(),
    }.Execute()
    
}