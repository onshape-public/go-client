package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestAppAssociativeDataAPI(t *testing.T) {
    InitializeTester[*onshape.AppAssociativeDataApiService](t)

    OpenAPITest{
        Call: onshape.ApiCopyAssociativeDataRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetAssociativeDataRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiPostAssociativeDataRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDeleteAssociativeDataRequest{},
        Expect: Todo(),
    }.Execute()
    
}