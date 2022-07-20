package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeWorkflowableTestObjectTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.WorkflowableTestObjectApi,
	}
}

func TestWorkflowableTestObjectAPI(t *testing.T) {
    ctx := initializeWorkflowableTestObjectTests(t)

    OpenAPITest{
        Call: onshape.ApiCreateWorkflowableTestObjectRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetWorkflowableTestObjectRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUpdateWorkflowableTestObjectRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiTransitionWorkflowableTestObjectRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}