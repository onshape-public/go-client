package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
)

func TestWebhookAPI(t *testing.T) {
    InitializeTester[*onshape.WebhookAPIService](t)

    OpenAPITest{
        Call: onshape.ApiGetWebhooksRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiCreateWebhookRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiGetWebhookRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUpdateWebhookRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiUnregisterWebhookRequest{},
        Expect: Todo(),
    }.Execute()
    
    OpenAPITest{
        Call: onshape.ApiPingWebhookRequest{},
        Expect: Todo(),
    }.Execute()
    
}