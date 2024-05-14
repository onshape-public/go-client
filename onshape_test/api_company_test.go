package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestCompanyAPI(t *testing.T) {
    InitializeTester[*onshape.CompanyAPIService](t)

    OpenAPITest{
        Call: onshape.ApiFindCompanyRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetCompanyRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetDocumentsByNameRequest{},
        Expect: Todo(),
    }.Execute()
    
}