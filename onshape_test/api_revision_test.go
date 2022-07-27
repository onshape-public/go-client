package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeRevisionTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.RevisionApi,
	}
}

func TestRevisionAPI(t *testing.T) {
    ctx := initializeRevisionTests(t)

    OpenAPITest{
        Call: onshape.ApiEnumerateRevisionsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetRevisionHistoryInCompanyByElementIdRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetRevisionHistoryInCompanyByPartIdRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetRevisionHistoryInCompanyByPartNumberRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiDeleteRevisionHistoryRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetLatestInDocumentOrCompanyRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}
/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiGetRevisionByPartNumberRequest{},
    Expect: Todo,
}.Execute(ctx, t)


***/