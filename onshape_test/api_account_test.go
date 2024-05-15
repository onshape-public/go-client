package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestAccountAPI(t *testing.T) {
    InitializeTester[*onshape.AccountAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetPlanPurchasesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetPurchasesRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiConsumePurchaseRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCancelPurchaseNewRequest{},
        Expect: Todo(),
    }.Execute()
    
}