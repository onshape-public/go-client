package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializePartNumberTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.PartNumberApi,
	}
}

func TestPartNumberAPI(t *testing.T) {
    ctx := initializePartNumberTests(t)

    OpenAPITest{
        Call: onshape.ApiNextNumbersRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}