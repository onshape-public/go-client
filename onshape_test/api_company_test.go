package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeCompanyTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.CompanyApi,
	}
}

func TestCompanyAPI(t *testing.T) {
    ctx := initializeCompanyTests(t)

    OpenAPITest{
        Call: onshape.ApiFindCompanyRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetCompanyRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetDocumentsByNameRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}