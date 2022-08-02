package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestPublicationAPI(t *testing.T) {
    InitializeTester[*onshape.PublicationApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetPublicationItemsRequest{},
        Expect: Todo(),
    }.Execute()
    
}