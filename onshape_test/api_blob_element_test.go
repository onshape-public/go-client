package onshape_test

import (
	"context"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

func initializeBlobElementTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.BlobElementApi,
	}
}

func TestBlobElementAPI(t *testing.T) {
	ctx := initializeBlobElementTests(t)

	OpenAPITest{
		Call:   onshape.ApiUploadFileCreateElementRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiDownloadFileWorkspaceRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiUploadFileUpdateElementRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiUpdateUnitsRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiCreateBlobTranslationRequest{},
		Expect: Todo,
	}.Execute(ctx, t)

}
