package onshape_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/onshape-public/go-client/onshape"
)

func initializeWebhookTests(t *testing.T) TestingContext {
    client, err := onshape.NewAPIClientFromEnv()
    assert.NoError(t, err)

    ctx := context.Background()

    return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.WebhookApi,
	}
}

func TestWebhookAPI(t *testing.T) {
    ctx := initializeWebhookTests(t)

    OpenAPITest{
        Call: onshape.ApiGetWebhooksRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiCreateWebhookRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiGetWebhookRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUpdateWebhookRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiUnregisterWebhookRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
    OpenAPITest{
        Call: onshape.ApiPingWebhookRequest{},
        Expect: Todo,
    }.Execute(ctx, t)
    
}