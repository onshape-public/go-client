package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestAPIApplicationAPI(t *testing.T) {
    InitializeTester[*onshape.APIApplicationApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetCompanyAppSettingsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateAppCompanySettingsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDeleteCompanyAppSettingsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetUserAppSettingsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateAppSettingsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDeleteAppSettingsRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetApplicableExtensionsForClientRequest{},
        Expect: Todo(),
    }.Execute()
    
}