package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeFeatureStudioTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.FeatureStudioApi,
	}
}

func TestFeatureStudioAPI(t *testing.T) {
    ctx := initializeFeatureStudioTests(t)

    OpenAPITest{
        Call: onshape.ApiCreateFeatureStudioRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetFeatureStudioContentsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUpdateFeatureStudioContentsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetFeatureStudioSpecsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}