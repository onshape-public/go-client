package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeThumbnailTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.ThumbnailApi,
	}
}

func TestThumbnailAPI(t *testing.T) {
    ctx := initializeThumbnailTests(t)

    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentAndVersionRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetDocumentThumbnailRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetElementThumbnailWithApiConfigurationRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetDocumentThumbnailWithSizeRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetElementThumbnailRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiSetApplicationElementThumbnailRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiDeleteApplicationThumbnailsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetElementThumbnailWithSizeRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentOldRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetThumbnailForDocumentAndVersionOldRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}