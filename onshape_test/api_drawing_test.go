package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeDrawingTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.DrawingApi,
	}
}

func TestDrawingAPI(t *testing.T) {
    ctx := initializeDrawingTests(t)

    OpenAPITest{
        Call: onshape.ApiCreateDrawingAppElementRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetDrawingTranslatorFormatsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiCreateDrawingTranslationRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}