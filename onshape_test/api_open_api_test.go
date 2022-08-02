package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestOpenApiAPI(t *testing.T) {
    InitializeTester[*onshape.OpenApiApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetOpenApiRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetTagsRequest{},
        Expect: Todo(),
    }.Execute()
    
}