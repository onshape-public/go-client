package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeExportRuleTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.ExportRuleApi,
	}
}

func TestExportRuleAPI(t *testing.T) {
    ctx := initializeExportRuleTests(t)

    OpenAPITest{
        Call: onshape.ApiGetValidRuleOptionsRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}