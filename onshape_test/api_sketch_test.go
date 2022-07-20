package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeSketchTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.SketchApi,
	}
}

func TestSketchAPI(t *testing.T) {
    ctx := initializeSketchTests(t)

    OpenAPITest{
        Call: onshape.ApiGetSketchInfoRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetSketchBoundingBoxesRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetTessellatedEntitiesRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}