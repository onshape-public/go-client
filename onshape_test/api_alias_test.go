package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeAliasTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.AliasApi,
	}
}

func TestAliasAPI(t *testing.T) {
    ctx := initializeAliasTests(t)

    OpenAPITest{
        Call: onshape.ApiGetAliasesInCompanyRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiCreateAliasRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetAliasRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUpdateAliasRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiDeleteAliasRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetAliasMembersRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}