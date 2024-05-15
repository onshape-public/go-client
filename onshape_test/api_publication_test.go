package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestPublicationAPI(t *testing.T) {
    InitializeTester[*onshape.PublicationAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetPublicationItemsRequest{},
        Expect: Todo(),
    }.Execute()
    
}