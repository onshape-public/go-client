package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeMetadataCategoryTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.MetadataCategoryApi,
	}
}

func TestMetadataCategoryAPI(t *testing.T) {
    ctx := initializeMetadataCategoryTests(t)

    OpenAPITest{
        Call: onshape.ApiGetCategoryPropertiesRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}