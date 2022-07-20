package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeFolderTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.FolderApi,
	}
}

func TestFolderAPI(t *testing.T) {
    ctx := initializeFolderTests(t)

    OpenAPITest{
        Call: onshape.ApiGetFolderAclRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiShareRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUnShareRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}