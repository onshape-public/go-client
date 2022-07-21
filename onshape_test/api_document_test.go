package onshape_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

const (
	ShareTestGroupId         = "62203d19a8823131729808eb-nonexistant"
	ShareTestGroupType int32 = 2
)

func initializeDocumentTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.DocumentApi,
	}
}

func TestDocumentAPI(t *testing.T) {
	ctx := TestingContext{
		"label":     Ptr("test-doc"),
		"docPublic": false,
		"bTDocumentParams": &onshape.BTDocumentParams{
			Description: Ptr("This is a test document"),
		},
	}.InheritDefaults(initializeDocumentTests(t))

	ctx = CreateDocumentPreTest(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetDocumentsRequest{},
		Expect: func(ctx TestingContext, t *testing.T, result *onshape.BTGlobalTreeNodeListResponse, _ *http.Response, err error) {
			assert.NoError(t, err)
			assert.NotZero(t, len(result.Items))

			assert.GreaterOrEqual(t, slices.IndexFunc(result.Items, func(q onshape.BTGlobalTreeNodeInfo) bool {
				btd := q.GetActualInstance().(*onshape.BTDocumentSummaryInfo)
				return btd.GetId() == ctx["did"] && Ptr(btd.GetCreatedBy()).GetId() == *ctx["creator"].(*string)
			}), 0)
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiUpdateDocumentAttributesRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetDocumentRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r *onshape.BTDocumentInfo, _ *http.Response, err error) {
			assert.NoError(t, err)
			assert.Equal(t, r.GetId(), ctx["did"])
			assert.Equal(t, r.GetName(), *ctx["label"].(*string))
			assert.Equal(t, r.GetPublic(), ctx["docPublic"])
			assert.Equal(t, r.GetDescription(), *ctx["bTDocumentParams"].(*onshape.BTDocumentParams).Description)
			assert.True(t, r.HasDefaultWorkspace())
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Name:   "Test ApiGetDocumentRequest nonexistant",
		Call:   onshape.ApiGetDocumentRequest{},
		Expect: APIError,
	}.Execute(ctx.SetDefault("did", "this-doesnt-exist"), t)

	OpenAPITest{
		Call: onshape.ApiGetDocumentAclRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r *onshape.BTAclInfo, _ *http.Response, err error) {
			assert.NoError(t, err)
			assert.Equal(t, r.Owner.GetId(), *ctx["owner"].(*string))
			assert.Equal(t, r.GetPublic(), ctx["docPublic"])
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetDocumentPermissionSetRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r []string, _ *http.Response, err error) {
			assert.NoError(t, err)
			val := []string{"READ", "COPY", "WRITE", "RESHARE", "EXPORT", "DELETE", "COMMENT", "LINK"}
			for _, c := range val {
				assert.True(t, slices.Contains(r, c))
			}
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiCreateVersionRequest{}.BTVersionOrWorkspaceParams(onshape.BTVersionOrWorkspaceParams{
			Name:        Ptr("Test version"),
			Description: Ptr("bindings test"),
			DocumentId:  Ptr(ctx["did"].(string)),
			FromHistory: Ptr(true),
		}),
		Expect: NoAPIError,
	}.Execute(ctx, t)

	var originalVersion string
	OpenAPITest{
		Call: onshape.ApiGetDocumentVersionsRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r []onshape.BTVersionInfo, _ *http.Response, err error) {
			assert.NoError(t, err)
			assert.Equal(t, 2, len(r))
			originalVersion = r[0].GetId()
		},
	}.Execute(ctx, t)

	var swid string
	OpenAPITest{
		Call: onshape.ApiCreateWorkspaceRequest{
			ApiService: ctx["client"].(*onshape.APIClient).DocumentApi,
		}.BTVersionOrWorkspaceParams(onshape.BTVersionOrWorkspaceParams{Name: Ptr("branch")}),
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTWorkspaceInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			swid = r.GetId()
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetDocumentWorkspacesRequest{},
		Expect: func(_ TestingContext, t *testing.T, r []onshape.BTWorkspaceInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.Len(t, r, 2)
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetCurrentMicroversionRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTMicroversionInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.Equal(t, ctx["mv"], r.GetMicroversion())
		},
	}.Execute(ctx.SetDefault("wv", "w").SetDefault("wvid", ctx["wid"]), t)

	OpenAPITest{
		Call:   onshape.ApiMergeIntoWorkspaceRequest{}.BTVersionOrWorkspaceInfo(onshape.BTVersionOrWorkspaceInfo{Id: &swid, Type: Ptr("workspace")}),
		Expect: NoAPIError,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiGetElementsInDocumentRequest{},
		Expect: NoAPIError,
	}.Execute(ctx.SetDefault("wvm", "w").SetDefault("wvmid", ctx["wid"]), t)

	OpenAPITest{
		Call: onshape.ApiGetCurrentMicroversionRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTMicroversionInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.NotEqual(t, ctx["mv"], r.GetMicroversion())
		},
	}.Execute(ctx.SetDefault("wv", "w").SetDefault("wvid", ctx["wid"]), t)

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{},
		Expect: func(_ TestingContext, t *testing.T, r []onshape.BTDocumentHistoryInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.Len(t, r, 2)
		},
	}.Execute(ctx.SetDefault("wm", "w").SetDefault("wmid", ctx["wid"]), t)

	OpenAPITest{
		Call:   onshape.ApiRestoreFromHistoryRequest{},
		Expect: NoAPIError,
	}.Execute(ctx.SetDefault("vm", "v").SetDefault("vmid", originalVersion), t)

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{},
		Expect: func(_ TestingContext, t *testing.T, r []onshape.BTDocumentHistoryInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.Len(t, r, 3)
		},
	}.Execute(ctx.SetDefault("wm", "w").SetDefault("wmid", ctx["wid"]), t)

	OpenAPITest{
		Call: onshape.ApiShareDocumentRequest{}.BTShareParams(onshape.BTShareParams{
			DocumentId:    Ptr(ctx["did"].(string)),
			WorkspaceId:   Ptr(ctx["wid"].(string)),
			PermissionSet: []string{"COPY", "EXPORT", "COMMENT", "WRITE"},
			Entries: []onshape.BTShareEntryParams{
				{
					TeamId:    Ptr(ShareTestGroupId),
					EntryType: Ptr(ShareTestGroupType),
				},
			},
		}),
		Expect: APIError,
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiUnShareDocumentRequest{},
		Expect: APIError,
	}.Execute(ctx.SetDefault("entryType", Ptr(ShareTestGroupType)).SetDefault("eid", ShareTestGroupId), t)

	DeleteDocumentPostTest(ctx, t)
}

func CreateDocumentPreTest(ctx TestingContext, t *testing.T) TestingContext {
	ctx = ctx.InheritDefaults(TestingContext{"label": Ptr("test-doc"), "docPublic": false})

	client := ctx["client"].(*onshape.APIClient)
	docParams := onshape.NewBTDocumentParams()
	docParams.Name = ctx["label"].(*string)
	docParams.SetIsPublic(ctx["docPublic"].(bool))
	doc, v, err := client.DocumentApi.CreateDocument(ctx["ctx"].(context.Context)).BTDocumentParams(*docParams).Execute()
	assert.NoError(t, err)
	if err != nil {
		spew.Dump(v)
	}
	return TestingContext{
		"did":     doc.GetId(),
		"wid":     *doc.GetDefaultWorkspace().Id,
		"mv":      *doc.GetDefaultWorkspace().Microversion,
		"owner":   doc.GetOwner().Id,
		"creator": Ptr(Ptr(doc.GetCreatedBy()).GetId()),
	}.InheritDefaults(ctx)
}

func DeleteDocumentPostTest(ctx TestingContext, t *testing.T) {
	client := ctx["client"].(*onshape.APIClient)
	_, _, err := client.DocumentApi.DeleteDocument(ctx["ctx"].(context.Context), ctx["did"].(string)).Execute()
	assert.NoError(t, err)
}

/*** ADDITIONAL TESTS

OpenAPITest{
	Call:   onshape.ApiRestoreFromHistoryRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiCopyWorkspaceRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetDocumentHistoryRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetCurrentMicroversionRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiDownloadExternalDataRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetVersionRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiUpdateExternalReferencesToLatestDocumentsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiMoveElementsToDocumentRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiRevertUnchangedToRevisionsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiSyncApplicationElementsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiExport2JsonRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetInsertablesRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiSearchRequest{},
	Expect: Todo,
},

***/
