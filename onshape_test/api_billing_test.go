package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestBillingAPI(t *testing.T) {
    InitializeTester[*onshape.BillingAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetClientPlansRequest{},
        Expect: Todo(),
    }.Execute()
    
}