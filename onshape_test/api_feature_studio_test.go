package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestFeatureStudioAPI(t *testing.T) {
    InitializeTester[*onshape.FeatureStudioAPIService](t)

    OpenAPITest{
        Call: onshape.ApiCreateFeatureStudioRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetFeatureStudioContentsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateFeatureStudioContentsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetFeatureStudioSpecsRequest{},
        Expect: Todo(),
    }.Execute()
    
}