package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeVersionTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.VersionApi,
	}
}

func TestVersionAPI(t *testing.T) {
    ctx := initializeVersionTests(t)

    OpenAPITest{
        Call: onshape.ApiGetAllVersionsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}