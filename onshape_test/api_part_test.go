package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestPartAPI(t *testing.T) {
    InitializeTester[*onshape.PartApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetPartsWMVRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetPartsWMVERequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetBodyDetailsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetBoundingBoxesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiExportPartGltfRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetMassPropertiesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetPartMetadataRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdatePartMetadataRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiExportPSRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetPartShadedViewsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetBendTableRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiExportStlRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetEdgesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetFaces1Request{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetStandardContentPartMetadataRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateStandardContentPartMetadataRequest{},
        Expect: Todo(),
    }.Execute()
    
}