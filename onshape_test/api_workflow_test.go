package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeWorkflowTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.WorkflowApi,
	}
}

func TestWorkflowAPI(t *testing.T) {
    ctx := initializeWorkflowTests(t)

    OpenAPITest{
        Call: onshape.ApiGetActiveWorkflowsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetAllowedApproversRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetAuditLogRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}