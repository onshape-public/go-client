package onshape_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

func initializeElementTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.ElementApi,
	}
}

func TestElementAPI(t *testing.T) {
	ctx := initializeElementTests(t)
	ctx = CreateDocumentPreTest(ctx, t)
	ctx = ctx.SetDefault("wvm", "w").SetDefault("bTAppElementParams", GetDefaultAppElementParams())

	OpenAPITest{
		Call: onshape.ApiCreateElementRequest{
			ApiService: ctx["client"].(*onshape.APIClient).AppElementApi,
		},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTAppElementModifyInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			ctx = ctx.SetDefault("eid", r.GetElementId())
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiDeleteElementRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)

	DeleteDocumentPostTest(ctx, t)
}

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiCopyElementFromSourceDocumentRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiEncodeConfigurationMapRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiUpdateReferencesRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetConfigurationRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiUpdateConfigurationRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiDecodeConfigurationRequest{},
    Expect: Todo,
},
OpenAPITest{
    Call: onshape.ApiGetElementTranslatorFormatsByVersionOrWorkspaceRequest{},
    Expect: Todo,
},

***/
