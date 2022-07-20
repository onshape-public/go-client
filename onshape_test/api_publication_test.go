package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializePublicationTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.PublicationApi,
	}
}

func TestPublicationAPI(t *testing.T) {
    ctx := initializePublicationTests(t)

    OpenAPITest{
        Call: onshape.ApiGetPublicationItemsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}