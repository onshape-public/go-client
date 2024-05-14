package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestSketchAPI(t *testing.T) {
    InitializeTester[*onshape.SketchAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetSketchInfoRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetSketchBoundingBoxesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetTessellatedEntitiesRequest{},
        Expect: Todo(),
    }.Execute()
    
}