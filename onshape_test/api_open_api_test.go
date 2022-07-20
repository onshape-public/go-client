package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeOpenApiTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.OpenApiApi,
	}
}

func TestOpenApiAPI(t *testing.T) {
    ctx := initializeOpenApiTests(t)

    OpenAPITest{
        Call: onshape.ApiGetOpenApiRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetTagsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}