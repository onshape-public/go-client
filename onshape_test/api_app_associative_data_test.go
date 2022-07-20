package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeAppAssociativeDataTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.AppAssociativeDataApi,
	}
}

func TestAppAssociativeDataAPI(t *testing.T) {
    ctx := initializeAppAssociativeDataTests(t)

    OpenAPITest{
        Call: onshape.ApiCopyAssociativeDataRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetAssociativeDataRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiPostAssociativeDataRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiDeleteAssociativeDataRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}