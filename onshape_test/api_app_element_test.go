package onshape_test

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

func initializeAppElementTests(t *testing.T) TestingContext {
	client, err := onshape.NewAPIClientFromEnv()
	assert.NoError(t, err)

	ctx := context.Background()

	return TestingContext{
		"client":     client,
		"ctx":        ctx,
		"ApiService": client.AppElementApi,
	}
}

func TestAppElementAPI(t *testing.T) {
	ctx := TestingContext{
		"wvm":                      "w",
		"bTAppElementParams":       GetDefaultAppElementParams(),
		"bTAppElementUpdateParams": GetDefaultAppUpdateParams(),
		"bTAppElementBulkCreateParams": &onshape.BTAppElementBulkCreateParams{
			Names:    []string{"Q.PRT", "R.PRT", "S.PRT", "T.PRT", "U.PRT", "V.PRT", "W.PRT", "X.PRT", "Y.PRT", "Z.PRT"},
			FormatId: "CollabBulkCreate",
		},
		"bid":      "blob1",
		"fileName": "./test_data/hf.txt",
	}.InheritDefaults(initializeAppElementTests(t))

	ctx = CreateDocumentPreTest(ctx, t)
	var eid string

	fileName := ctx["fileName"].(string)
	osFile, err := os.Open(fileName)
	assert.NoError(t, err)
	file := onshape.NewHttpFileFromOsFile(osFile)

	OpenAPITest{
		Call: onshape.ApiCreateElementRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r *onshape.BTAppElementModifyInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			eid = *r.ElementId
		},
	}.Execute(ctx, t)

	ctx = TestingContext{"file": &file, "eid": eid, "wvmid": ctx["wid"]}.InheritDefaults(ctx)

	OpenAPITest{
		Call: onshape.ApiBulkCreateElementRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r *onshape.BTAppElementBulkCreateInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			bcp := *ctx["bTAppElementBulkCreateParams"].(*onshape.BTAppElementBulkCreateParams)
			assert.EqualValues(t, len(r.ElementIds), len(bcp.Names))
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetSubElementContentRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r *onshape.BTAppElementContentInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.Zero(t, len(r.GetData()))
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiUploadBlobSubelementRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiDownloadBlobSubelementWorkspaceRequest{},
		Expect: func(ctx TestingContext, t *testing.T, res *onshape.HttpFile, v *http.Response, err error) {
			assert.NoError(t, err)
			defer v.Body.Close()
			fileBytes, err := ioutil.ReadAll(res.Data)
			assert.NoError(t, err)
			buf, err := os.ReadFile(ctx["fileName"].(string))
			assert.NoError(t, err)
			assert.EqualValues(t, buf, fileBytes)
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetJsonRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r *onshape.BTGetJsonResponse2137, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.EqualValues(t, r.GetTree().AdditionalProperties, *ctx["bTAppElementParams"].(*onshape.BTAppElementParams).JsonTree)
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiUpdateAppElementRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetJsonRequest{},
		Expect: func(ctx TestingContext, t *testing.T, r *onshape.BTGetJsonResponse2137, v *http.Response, err error) {
			assert.NoError(t, err)

			nm := make(map[string]interface{})
			for k, v := range *ctx["bTAppElementParams"].(*onshape.BTAppElementParams).JsonTree {
				nm[k] = v
			}

			insert := ctx["bTAppElementUpdateParams"].(*onshape.BTAppElementUpdateParams).JsonTreeEdit.GetActualInstance().(*onshape.BTJEditInsert2523)
			nm[insert.Path.GetPath()[0].GetActualInstance().(*onshape.BTJPathKey3221).GetKey()] = insert.GetValue()

			assert.EqualValues(t, r.GetTree().AdditionalProperties, nm)
		},
	}.Execute(ctx, t)

	DeleteDocumentPostTest(ctx, t)
}

func TestTransactionAppElementAPI(t *testing.T) {
	ctx := TestingContext{
		"bTVersionOrWorkspaceParams": &onshape.BTVersionOrWorkspaceParams{
			Name: Ptr("branch"),
		},
		"bTAppElementParams": &onshape.BTAppElementParams{
			FormatId: "String",
			Name:     Ptr("test-element"),
		},
		"wvm": "w",
		"wm":  "w",
	}.InheritDefaults(initializeAppElementTests(t))

	ctx = CreateDocumentPreTest(ctx, t)

	OpenAPITest{
		Call: onshape.ApiCreateWorkspaceRequest{
			ApiService: ctx["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTWorkspaceInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			ctx = ctx.SetDefault("wid", r.GetId())
			ctx = ctx.SetDefault("wvmid", r.GetId())
			ctx = ctx.SetDefault("wmid", r.GetId())
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiCreateElementRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTAppElementModifyInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			ctx = ctx.SetDefault("eid", r.GetElementId())
			ctx = ctx.SetDefault("bTAppElementStartTransactionParams", &onshape.BTAppElementStartTransactionParams{
				Description:    Ptr("Transaction to populate Assembly w/Data"),
				ParentChangeId: Ptr(r.GetChangeId()),
			})
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiStartTransactionRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTAppElementModifyInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			ctx = ctx.SetDefault("txnId", r.GetTransactionId())
		},
	}.Execute(ctx, t)

	for _, eup := range []*onshape.BTAppElementUpdateParams{
		CreateAppElementChangeUpdate("Chapter 1"), CreateAppElementChangeUpdate("Chapter 2"), CreateAppElementChangeUpdate("Chapter 3"),
	} {
		eup.SetTransactionId(ctx["txnId"].(string))
		ctx = ctx.SetDefault("bTAppElementUpdateParams", eup)

		OpenAPITest{
			Call: onshape.ApiUpdateAppElementRequest{},
			Expect: func(_ TestingContext, t *testing.T, r *onshape.BTAppElementModifyInfo, v *http.Response, err error) {
				assert.NoError(t, err)
				assert.Equal(t, ctx["txnId"], r.GetTransactionId())
			},
		}.Execute(ctx, t)
	}

	btAppElementCommitTransactionParams := onshape.NewBTAppElementCommitTransactionParams()
	btAppElementCommitTransactionParams.SetDescription("Done Updating the Assembly")
	btAppElementCommitTransactionParams.SetTransactionIds([]string{ctx["txnId"].(string)})

	OpenAPITest{
		Call:   onshape.ApiCommitTransactionsRequest{}.BTAppElementCommitTransactionParams(*btAppElementCommitTransactionParams),
		Expect: NoAPIError,
	}.Execute(ctx, t)

	var latestCommit string
	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{
			ApiService: ctx["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: func(_ TestingContext, t *testing.T, r []onshape.BTDocumentHistoryInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.LessOrEqual(t, 3, len(r))
			latestCommit = r[0].GetMicroversionId()
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call: onshape.ApiGetSubelementIdsRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTAppElementIdsInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			assert.Equal(t, 3, len(r.GetSubelementIds()))
		},
	}.Execute(TestingContext{
		"wvm":   "m",
		"wvmid": latestCommit,
	}.InheritDefaults(ctx), t)

	OpenAPITest{
		Call: onshape.ApiStartTransactionRequest{},
		Expect: func(_ TestingContext, t *testing.T, r *onshape.BTAppElementModifyInfo, v *http.Response, err error) {
			assert.NoError(t, err)
			ctx = ctx.SetDefault("tid", r.GetTransactionId())
		},
	}.Execute(ctx, t)

	OpenAPITest{
		Call:   onshape.ApiAbortTransactionRequest{},
		Expect: NoAPIError,
	}.Execute(ctx, t)
}

func GetDefaultAppElementParams() *onshape.BTAppElementParams {
	return &onshape.BTAppElementParams{
		Name:     Ptr("Test-element-1"),
		FormatId: "String",
		JsonTree: Ptr(map[string]interface{}{
			"name":    "p.prt",
			"_nodeId": "master",
			"masterProperties": map[string]interface{}{
				"attr1":   "val1",
				"attr2":   "val2",
				"_nodeId": "masterProperties",
			},
			"versionProperties": map[string]interface{}{
				"attr1":   "v11",
				"attr2":   "v22",
				"attr3":   "v33",
				"_nodeId": "versionProperties",
			},
		}),
	}
}

func GetDefaultAppUpdateParams() *onshape.BTAppElementUpdateParams {
	bTAppElementUpdateParams := onshape.NewBTAppElementUpdateParams()
	//Top Level Edit Element
	btjEditInsert := onshape.NewBTJEditInsert2523()
	btjEditInsert.SetBtType("BTJEditInsert-2523")
	//Path Element of the Edit
	path := onshape.NewBTJPath3073("master")
	pathType := string("BTJPath-3073")
	path.BtType = &pathType
	// path.SetStartNode("master") -- not needed any more, passing it in the constructor
	//Path.path element
	pathKey := onshape.NewBTJPathKey3221()
	pathKey.SetBtType("BTJPathKey-3221")
	pathKey.SetKey("chapterProperties")
	path.SetPath([]onshape.BTJPathElement2297{*pathKey.AsBTJPathElement2297()})
	btjEditInsert.SetPath(*path)
	value := map[string]interface{}{
		"_nodeId": "chapterProperties",
		"chp1":    "v1",
		"chp2":    "v2",
	}
	btjEditInsert.SetValue(value)
	bTAppElementUpdateParams.SetJsonTreeEdit(*btjEditInsert.AsBTJEdit3734())

	return bTAppElementUpdateParams
}

func CreateChapterTestSubelement(name string) onshape.BTAppElementChangeParams {
	return onshape.BTAppElementChangeParams{SubelementId: name, BaseContent: Ptr(base64.StdEncoding.EncodeToString([]byte(name)))}
}

func CreateAppElementChangeUpdate(name string) *onshape.BTAppElementUpdateParams {
	return &onshape.BTAppElementUpdateParams{
		Description: Ptr("Update " + name),
		Changes:     []onshape.BTAppElementChangeParams{CreateChapterTestSubelement(name)},
	}
}

/*** ADDITIONAL TESTS

OpenAPITest{
	Call:   onshape.ApiDownloadBlobSubelementRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiDeleteBlobSubelementRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetElementTransactionsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiStartTransactionRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiCommitTransactionsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetBlobSubelementIdsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiCompareAppElementJsonRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetAppElementHistoryRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetSubelementIdsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiGetJsonPathsRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiDeleteAppElementContentRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiCreateReferenceRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiResolveReferenceRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiUpdateReferenceRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiDeleteReferenceRequest{},
	Expect: Todo,
},
OpenAPITest{
	Call:   onshape.ApiResolveReferencesRequest{},
	Expect: Todo,
},

***/
