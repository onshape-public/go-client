package onshape_test

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	"github.com/onshape-public/go-client/onshape"
	"github.com/stretchr/testify/assert"
)

func TestAppElementAPI(t *testing.T) {
	InitializeTester[*onshape.AppElementApiService](t)

	fileName := "./test_data/hf.txt"
	osFile, err := os.Open(fileName)
	assert.NoError(Tester(), err)
	file := onshape.NewHttpFileFromOsFile(osFile)

	SetContext(TestingContext{
		"wvm":                      "w",
		"bTAppElementParams":       GetDefaultAppElementParams(),
		"bTAppElementUpdateParams": GetDefaultAppUpdateParams(),
		"bTAppElementBulkCreateParams": &onshape.BTAppElementBulkCreateParams{
			Names:    []string{"Q.PRT", "R.PRT", "S.PRT", "T.PRT", "U.PRT", "V.PRT", "W.PRT", "X.PRT", "Y.PRT", "Z.PRT"},
			FormatId: "CollabBulkCreate",
		},
		"bid":      "blob1",
		"fileName": fileName,
		"file":     &file,
	}.InheritDefaults(Context()))

	OpenAPITest{
		Call: onshape.ApiCreateDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		}.BTDocumentParams(onshape.BTDocumentParams{
			Name:        Ptr("test-doc"),
			Description: Ptr("This is a test document"),
			IsPublic:    Ptr(false),
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiCreateElementRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiBulkCreateElementRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementBulkCreateInfo) {
			bcp := *Context()["bTAppElementBulkCreateParams"].(*onshape.BTAppElementBulkCreateParams)
			assert.EqualValues(Tester(), len(r.ElementIds), len(bcp.Names))
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetSubElementContentRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementContentInfo) {
			assert.Zero(Tester(), len(r.GetData()))
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUploadBlobSubelementRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDownloadBlobSubelementWorkspaceRequest{},
		Expect: NoAPIErrorAnd(func(res *onshape.HttpFile) {
			fileBytes, err := ioutil.ReadAll(res.Data)
			assert.NoError(Tester(), err)
			buf, err := os.ReadFile(Context()["fileName"].(string))
			assert.NoError(Tester(), err)
			assert.EqualValues(Tester(), buf, fileBytes)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetJsonRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTGetJsonResponse2137) {
			assert.EqualValues(Tester(), r.GetTree().AdditionalProperties, *Context()["bTAppElementParams"].(*onshape.BTAppElementParams).JsonTree)
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiUpdateAppElementRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetJsonRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTGetJsonResponse2137) {
			nm := make(map[string]interface{})
			for k, v := range *Context()["bTAppElementParams"].(*onshape.BTAppElementParams).JsonTree {
				nm[k] = v
			}

			insert := Context()["bTAppElementUpdateParams"].(*onshape.BTAppElementUpdateParams).JsonTreeEdit.GetActualInstance().(*onshape.BTJEditInsert2523)
			nm[insert.Path.GetPath()[0].GetActualInstance().(*onshape.BTJPathKey3221).GetKey()] = insert.GetValue()

			assert.EqualValues(t, r.GetTree().AdditionalProperties, nm)
		}),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDeleteDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()
}

func TestTransactionAppElementAPI(t *testing.T) {
	InitializeTester[*onshape.AppElementApiService](t)

	SetContext(TestingContext{
		"bTVersionOrWorkspaceParams": &onshape.BTVersionOrWorkspaceParams{
			Name: Ptr("branch"),
		},
		"bTAppElementParams": &onshape.BTAppElementParams{
			FormatId: "String",
			Name:     Ptr("test-element"),
		},
		"wvm": "w",
		"wm":  "w",
	}.InheritDefaults(Context()))

	OpenAPITest{
		Call: onshape.ApiCreateDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		}.BTDocumentParams(onshape.BTDocumentParams{
			Name:        Ptr("test-doc"),
			Description: Ptr("This is a test document"),
			IsPublic:    Ptr(false),
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateWorkspaceRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiCreateElementRequest{},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementModifyInfo) {
			Context()["bTAppElementStartTransactionParams"] = &onshape.BTAppElementStartTransactionParams{
				Description:    Ptr("Transaction to populate Assembly w/Data"),
				ParentChangeId: Ptr(r.GetChangeId()),
			}
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiStartTransactionRequest{},
		Expect: NoAPIError(),
	}.Execute()

	for _, eup := range []*onshape.BTAppElementUpdateParams{
		CreateAppElementChangeUpdate("Chapter 1"), CreateAppElementChangeUpdate("Chapter 2"), CreateAppElementChangeUpdate("Chapter 3"),
	} {
		eup.SetTransactionId(Context()["tid"].(string))

		OpenAPITest{
			Call: onshape.ApiUpdateAppElementRequest{}.BTAppElementUpdateParams(*eup),
			Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementModifyInfo) {
				assert.Equal(Tester(), Context()["tid"].(string), r.GetTransactionId())
			}),
		}.Execute()
	}

	OpenAPITest{
		Call: onshape.ApiCommitTransactionsRequest{}.BTAppElementCommitTransactionParams(onshape.BTAppElementCommitTransactionParams{
			Description:    Ptr("Done Updating the Assembly"),
			TransactionIds: []string{Context()["tid"].(string)},
		}),
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiGetDocumentHistoryRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIErrorAnd(func(r []onshape.BTDocumentHistoryInfo) {
			assert.LessOrEqual(Tester(), 3, len(r))
			Context()["wvmid"] = r[0].GetMicroversionId()
		}),
	}.Execute()

	OpenAPITest{
		Call:    onshape.ApiGetSubelementIdsRequest{},
		Context: TestingContext{"wvm": "m"},
		Expect: NoAPIErrorAnd(func(r *onshape.BTAppElementIdsInfo) {
			assert.Equal(Tester(), 3, len(r.GetSubelementIds()))
		}),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiStartTransactionRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call:   onshape.ApiAbortTransactionRequest{},
		Expect: NoAPIError(),
	}.Execute()

	OpenAPITest{
		Call: onshape.ApiDeleteDocumentRequest{
			ApiService: Context()["client"].(*onshape.APIClient).DocumentApi,
		},
		Expect: NoAPIError(),
	}.Execute()
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
