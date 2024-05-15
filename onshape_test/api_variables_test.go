package onshape_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/require"
)

func TestVariablesAPI(t *testing.T) {
    InitializeTester[*onshape.VariablesAPIService](t)

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

func TestGetVariablesExecute(t *testing.T) {
    InitializeTester[*onshape.VariablesAPIService](t)
    obj, r, err := Context()["client"].(*onshape.APIClient).VariablesAPI.GetVariablesExecute(onshape.ApiGetVariablesRequest{})
    require.Error(t, err)
    // Check if the returned values are of the expected types
    type objOnshape []onshape.BTVariableTableInfo 
    if reflect.TypeOf(obj).Kind() != reflect.TypeOf(objOnshape{}).Kind()  {
        t.Errorf("Expected an []BTVariableTableInfo, got %T", obj)
    }

    type httpRes http.Response
    if reflect.TypeOf(r).Kind() != reflect.TypeOf(&httpRes{}).Kind()  {
        t.Errorf("Expected a *http.Response, got %T", r)
    }
    if _, ok := err.(error); !ok {
        t.Errorf("Expected a error, got %T", err)
    }
}