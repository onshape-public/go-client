package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestVariablesAPI(t *testing.T) {
    InitializeTester[*onshape.VariablesApiService](t)

    OpenAPITest{
        Call: onshape.ApiSetVariablesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiSetVariableStudioReferencesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiSetVariableStudioScopeRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateVariableStudioRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetVariablesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetVariableStudioReferencesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetVariableStudioScopeRequest{},
        Expect: Todo(),
    }.Execute()
    
}