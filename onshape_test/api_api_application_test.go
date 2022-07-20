package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeAPIApplicationTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.APIApplicationApi,
	}
}

func TestAPIApplicationAPI(t *testing.T) {
    ctx := initializeAPIApplicationTests(t)

    OpenAPITest{
        Call: onshape.ApiGetCompanyAppSettingsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUpdateAppCompanySettingsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiDeleteCompanyAppSettingsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetUserAppSettingsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUpdateAppSettingsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiDeleteAppSettingsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetApplicableExtensionsForClientRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}