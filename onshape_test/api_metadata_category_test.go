package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestMetadataCategoryAPI(t *testing.T) {
    InitializeTester[*onshape.MetadataCategoryAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetCategoryPropertiesRequest{},
        Expect: Todo(),
    }.Execute()
    
}