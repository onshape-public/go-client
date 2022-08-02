package onshape_test

import (
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

const (
	ShareTestGroupId         = "62203d19a8823131729808eb-nonexistant"
	ShareTestGroupType int32 = 2
)

func TestDocumentAPI(t *testing.T) {
	InitializeTester[*onshape.DocumentApiService](t)

	SetContext(TestingContext{
		"label":     Ptr("test-doc"),
		"docPublic": false,
		"bTDocumentParams": &onshape.BTDocumentParams{
			Name:        Ptr("test-doc"),
			Description: Ptr("This is a test document"),
			IsPublic:    Ptr(false),
		},
		"wm":  "w",
		"wv":  "w",
		"wvm": "w",
		"vm":  "v",
	}.InheritDefaults(Context()))

	OpenAPITest{
		Call:   onshape.ApiCreateDocumentRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentsRequest{},
		Expect: NoAPIErrorAnd(func(result *onshape.BTGlobalTreeNodeListResponse) {
			assert.NotZero(Tester(), len(result.Items))
			assert.GreaterOrEqual(Tester(), slices.IndexFunc(result.Items, func(q onshape.BTGlobalTreeNodeInfo) bool {
				btd := q.GetActualInstance().(*onshape.BTDocumentSummaryInfo)
				return btd.GetId() == Context()["did"] && Ptr(btd.GetCreatedBy()).GetId() == *Context()["creator"].(*string)
			}), 0)
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateDocumentAttributesRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTDocumentInfo) {
			assert.Equal(Tester(), r.GetId(), Context()["did"])
			assert.Equal(Tester(), r.GetName(), *Context()["label"].(*string))
			assert.Equal(Tester(), r.GetPublic(), Context()["docPublic"])
			assert.Equal(Tester(), r.GetDescription(), *Context()["bTDocumentParams"].(*onshape.BTDocumentParams).Description)
			assert.True(Tester(), r.HasDefaultWorkspace())
		}),
	}.Execute()

	OpenAPITest{
		Name:    "Test ApiGetDocumentRequest nonexistant",
		Call:    onshape.ApiGetDocumentRequest{},
		Context: TestingContext{"did": "this-doesnt-exit"},
		Expect:  APIError(),
		Pass:    Nothing(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentAclRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAclInfo) {
			assert.Equal(Tester(), r.Owner.GetId(), *Context()["owner"].(*string))
			assert.Equal(Tester(), r.GetPublic(), Context()["docPublic"])
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentPermissionSetRequest{},
		Expect: NoAPIErrorAnd(func(r []string) {
			val := []string{"READ", "COPY", "WRITE", "RESHARE", "EXPORT", "DELETE", "COMMENT", "LINK"}
			for _, c := range val {
				assert.True(Tester(), slices.Contains(r, c))
			}
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateVersionRequest{}.BTVersionOrWorkspaceParams(onshape.BTVersionOrWorkspaceParams{
			Name:        Ptr("Test version"),
			Description: Ptr("bindings test"),
			DocumentId:  Ptr(Context()["did"].(string)),
			FromHistory: Ptr(true),
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentVersionsRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTVersionInfo) {
			assert.Equal(Tester(), 2, len(r))
			Context()["vmid"] = r[0].GetId()
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateWorkspaceRequest{}.BTVersionOrWorkspaceParams(onshape.BTVersionOrWorkspaceParams{Name: Ptr("branch")}),
		Expect: NoAPIErrorAnd(func(r *onshape.BTWorkspaceInfo) {
			Context()["swid"] = r.Id
		}),
		Pass: Nothing(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentWorkspacesRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTWorkspaceInfo) {
			assert.Len(Tester(), r, 2)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetCurrentMicroversionRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTMicroversionInfo) {
			assert.Equal(Tester(), Context()["mv"], r.GetMicroversion())
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiMergeIntoWorkspaceRequest{}.
			BTVersionOrWorkspaceMergeInfo(onshape.BTVersionOrWorkspaceMergeInfo{
				Id:                            Context()["swid"].(*string),
				DefaultMergeStrategy:          Ptr("KEEP"),
				MergeStrategyElementOverrides: Ptr(map[string]string{}),
				Type:                          Ptr("workspace")}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiGetElementsInDocumentRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetCurrentMicroversionRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTMicroversionInfo) {
			assert.NotEqual(Tester(), Context()["mv"], r.GetMicroversion())
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTDocumentHistoryInfo) {
			assert.Len(Tester(), r, 2)
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiRestoreFromHistoryRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{},
		Expect: NoAPIErrorAnd(func(r []onshape.BTDocumentHistoryInfo) {
			assert.Len(Tester(), r, 3)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiShareDocumentRequest{}.BTShareParams(onshape.BTShareParams{
			DocumentId:    Ptr(Context()["did"].(string)),
			WorkspaceId:   Ptr(Context()["wid"].(string)),
			PermissionSet: []string{"COPY", "EXPORT", "COMMENT", "WRITE"},
			Entries: []onshape.BTShareEntryParams{
				{
					TeamId:    Ptr(ShareTestGroupId),
					EntryType: Ptr(ShareTestGroupType),
				},
			},
		}),
		Expect: APIError(),
	}.Execute()

	OpenAPITest{
		Call:    onshape.ApiUnShareDocumentRequest{},
		Context: TestingContext{"entryType": Ptr(ShareTestGroupType), "eid": ShareTestGroupId},
		Expect:  APIError(),
	}.Execute()

	OpenAPITest{
		Call:    onshape.ApiMergePreviewRequest{},
		Context: TestingContext{"sourceType": Ptr("w"), "sourceId": Context()["swid"].(*string)},
		Expect: NoAPIErrorAnd(func(r *onshape.BTMergePreviewInfo) {
			assert.NotNil(Tester(), r)
			assert.Equalf(Tester(), r.GetTargetMicroversionId(), Context()["mv"], "Target Microversion Id should be the same as the workspace id")
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiDeleteDocumentRequest{},
		Expect: NoAPIError(),
	}.Execute()
}

/*** ADDITIONAL TESTS

OpenAPITest{
    Call: onshape.ApiDownloadExternalDataRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetVersionRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiUpdateExternalReferencesToLatestDocumentsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiMoveElementsToDocumentRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiRevertUnchangedToRevisionsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiSyncApplicationElementsRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiDeleteWorkspaceRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetUnitInfoRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiExport2JsonRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiGetInsertablesRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiSearchRequest{},
    Expect: Todo(),
}.Execute()

OpenAPITest{
    Call: onshape.ApiCopyWorkspaceRequest{},
    Expect: Todo(),
}.Execute()


***/
