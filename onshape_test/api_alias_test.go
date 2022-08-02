package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestAliasAPI(t *testing.T) {
    InitializeTester[*onshape.AliasApiService](t)

    OpenAPITest{
        Call: onshape.ApiGetAliasesInCompanyRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateAliasRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetAliasRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateAliasRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiDeleteAliasRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetAliasMembersRequest{},
        Expect: Todo(),
    }.Execute()
    
}