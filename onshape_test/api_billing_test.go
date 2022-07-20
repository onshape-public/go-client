package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeBillingTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.BillingApi,
	}
}

func TestBillingAPI(t *testing.T) {
    ctx := initializeBillingTests(t)

    OpenAPITest{
        Call: onshape.ApiGetClientPlansRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}