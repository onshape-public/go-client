package onshape_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/onshape-public/go-client/onshape"
	"github.com/onshape-public/go-client/onshape_test/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateAndGetAppElement(t *testing.T) {
	//var docInfo onshape.BTDocumentInfo
	documentName := "App-Elements-testDoc-" + uuid.New().String()
	did, wid, teardownDoc := testhelper.SetupDocument(ctx, client, documentName)

	defer teardownDoc()

	type args struct {
		elementName  string
		subElements  []onshape.BTAppElementChangeParams
		propMap      map[string]interface{}
		blobFilePath string
	}
	tests := []struct {
		name string
		args args
		want *onshape.BTAppElementContentInfo
	}{
		{"CreateAppElementWithNoSubElement",
			args{"Test-element-1",
				nil,
				map[string]interface{}{
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
				},
				"./test_data/hf.txt",
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appElementParams := onshape.NewBTAppElementParams("String")
			appElementName := tt.args.elementName
			appElementParams.Name = &appElementName
			appElementParams.Subelements = tt.args.subElements
			propMap := tt.args.propMap
			propMapJsonWanted, err := json.Marshal(propMap)
			if err != nil {
				t.Error("Should be able to marshal the Input Map")
			}
			appElementParams.SetJsonTree(propMap)

			appElementModifyInfo, rawResp, err :=
				client.AppElementApi.CreateElement(ctx, did, wid).BTAppElementParams(*appElementParams).Execute()

			if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
				t.Error("err: ", err, " -- Response status: ", rawResp)
			} else {
				t.Log("Created an element w/the name: ", appElementModifyInfo.GetElementId())
			}
			eid := appElementModifyInfo.GetElementId()

			appElementContentInfo, rawResp, err := client.AppElementApi.GetSubElementContent(ctx, did, "w", wid, eid).Execute()
			if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
				t.Error("err: ", err, " -- Response status: ", rawResp)
			}
			if tt.want == nil {
				if len(appElementContentInfo.GetData()) != 0 {
					t.Errorf("expected 0 subelement, got %d", len(appElementContentInfo.GetData()))
				}
			}
			//Check the JSONTreeContent
			jsonResponse, rawResp, err := client.AppElementApi.GetJson(ctx, did, "w", wid, eid).Execute()
			if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
				t.Error("err: ", err, " -- Response status: ", rawResp)
			}
			require.True(t, jsonResponse.HasTree())
			propMapJsonResult, err := json.Marshal(jsonResponse.Tree)
			if err != nil {
				t.Error("Should be able to marshal the Received Map")
			}
			eq, err := testhelper.JSONBytesEqual(propMapJsonWanted, propMapJsonResult)
			if err != nil {
				t.Error("Can't compare JSONs")
			}
			require.True(t, eq, string(propMapJsonWanted), string(propMapJsonResult))

			//Update JSONTree: Insert
			//Too much work for now, I know :)
			//I'll make it better
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
			client.AppElementApi.UpdateAppElement(ctx, did, eid, "w", wid).BTAppElementUpdateParams(*bTAppElementUpdateParams).Execute()

			//Now get the JSONTree data again and see if the "chapterProperties" are there ...
			jsonResponse, rawResp, err = client.AppElementApi.GetJson(ctx, did, "w", wid, eid).Execute()
			if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
				t.Error("err: ", err, " -- Response status: ", rawResp)
			}
			require.True(t, jsonResponse.HasTree())
			propMapJsonResult, err = json.Marshal(jsonResponse.Tree)
			if err != nil {
				t.Error("Should be able to marshal the Received Map")
			}
			require.True(t, strings.Contains(string(propMapJsonResult), "chapterProperties"))

			//Try to send AppElement Upload data
			f, err := os.Open(tt.args.blobFilePath)
			require.NoError(t, err, "Error opening file")
			fi, err := f.Stat()
			require.NoError(t, err, "Can't stat file")
			hf := onshape.NewHttpFileFromOsFile(f)

			appElementModifyInfo, rawResp, err = client.AppElementApi.UploadBlobSubelement(ctx, did, wid, eid, "blob0001").Description("Test Blob Element").File(hf).Execute()
			require.NoError(t, err, "Error Uploading Blob ...")

			//Read Blob data ...
			_, resp, err := client.AppElementApi.DownloadBlobSubelementWorkspace(ctx, did, wid, eid, "blob0001").Execute()
			if err != nil {
				//openapi-generator 5.1.1 is trying to be smart and create a decode call for xml and json but we have neither
				//so check the error and move on ...
				_, ok := err.(onshape.GenericOpenAPIError)
				if !ok {
					t.Error("Unexpected Error: Unable to download Blob Subelement", err)
				}
			}
			//require.NoError(t, err, "Unable to download Blob Subelement")
			defer resp.Body.Close()

			fileBytes, err := ioutil.ReadAll(resp.Body)
			require.True(t, fi.Size() == int64(len(fileBytes)))
			require.NoError(t, err, "can't read response body")
			newFile, err := os.Create(tt.args.blobFilePath + ".new")
			require.NoError(t, err, "can't create file ")
			defer newFile.Close()
			_, err = newFile.Write(fileBytes)
			require.NoError(t, err, "can't write data into the new file")
		})
	}
}

func TestCreateBulkAppElement(t *testing.T) {
	documentName := "App-Elements-testDocBulk-" + uuid.New().String()
	did, wid, teardownDoc := testhelper.SetupDocument(ctx, client, documentName)

	defer teardownDoc()

	type args struct {
		elementNames []string
	}
	tests := []struct {
		name string
		args args
		want *interface{}
	}{
		{"CreateAppElementWithNoSubElement",
			args{[]string{"Q.PRT", "R.PRT", "S.PRT", "T.PRT", "U.PRT", "V.PRT", "W.PRT", "X.PRT", "Y.PRT", "Z.PRT"}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bulkAppElementParams := onshape.NewBTAppElementBulkCreateParams("CollabBulkCreate")
			bulkAppElementParams.SetNames(tt.args.elementNames)
			bulkAppElementParams.SetDescription("bulk test creation")

			appElementBulkCreateInfo, rawResp, err :=
				client.AppElementApi.BulkCreateElement(ctx, did, wid).BTAppElementBulkCreateParams(*bulkAppElementParams).Execute()
			require.NoError(t, err, "Error Creating Bulk App Element")
			require.NotNil(t, rawResp, "Response should not be nil")
			require.True(t, rawResp.StatusCode < 300, "Status code should be less than 300")
			require.NotNil(t, appElementBulkCreateInfo, "AppElementBulkCreateInfo should not be nil")
			require.True(t, len(appElementBulkCreateInfo.GetElementIds()) == len(tt.args.elementNames))
		})
	}
}
