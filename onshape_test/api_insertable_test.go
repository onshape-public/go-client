package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeInsertableTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.InsertableApi,
	}
}

func TestInsertableAPI(t *testing.T) {
    ctx := initializeInsertableTests(t)

    OpenAPITest{
        Call: onshape.ApiGetLatestInDocumentRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}