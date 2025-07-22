package onshape_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/onshape-public/go-client/onshape"
	"github.com/onshape-public/go-client/onshape_test/testhelper"
	"github.com/stretchr/testify/require"
)

var client *onshape.APIClient
var ctx context.Context

// TestMain is a common method for all tests in package onshape_test
func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func setup() {
	var err error

	client, err = onshape.NewAPIClientFromEnv()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// TODO: check if default workspace is also created
func TestCreateAndGetDocument(t *testing.T) {
	uid := uuid.NewString()
	document1Name := "Documents-testDoc-1-" + uid
	type args struct {
		docName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test001",
			args{document1Name},
			document1Name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docParams := onshape.NewBTDocumentParams()
			docParams.Name = &testhelper.DocumentName
			docParams.Description = &testhelper.DocumentDescription
			docParams.SetName(tt.args.docName)
			docParams.SetIsPublic(false)

			t.Log("Creating document")
			docInfo, rawResp, err := client.DocumentApi.CreateDocument(ctx).BTDocumentParams(*docParams).Execute()
			if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
				t.Error("err: ", err, " -- Response status: ", rawResp)
			} else {
				if docInfo.GetName() != tt.want {
					t.Errorf("CreateDocument() got = %s, want %s", docInfo.GetName(), tt.want)
				}
			}
			//check default workspace is also created
			if !docInfo.HasDefaultWorkspace() {
				t.Error("Create Document should have created a default workspace ")
				return
			}

			t.Log("Getting a document")
			getDocInfo, rawResp, err := client.DocumentApi.GetDocument(ctx, docInfo.GetId()).Execute()
			if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
				t.Error("err: ", err, " -- Response status: ", rawResp)
			} else {
				if getDocInfo.GetName() != tt.want {
					t.Errorf("GetDocument() got = %s, want %s", getDocInfo.GetName(), tt.want)
				}
			}

			t.Log("Deleting a document")
			_, rawResp, err = client.DocumentApi.DeleteDocument(ctx, getDocInfo.GetId()).Execute()

			if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
				t.Error("err: ", err, " -- Response status: ", rawResp)
			} else {
				t.Log("Deleted successfully document w/the name: ", docInfo.GetName())
			}
		})
	}

}

// TODO: check if default workspace is also created
func TestGetDocument(t *testing.T) {
	uid := uuid.NewString()
	document1Name := "Documents-testDoc-1-" + uid
	type args struct {
		docName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test001",
			args{document1Name},
			document1Name,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			resp, _, err := client.DocumentApi.GetDocuments(ctx).Filter(6).Owner("5ecd79b2a3e6eb1251369b10").SortColumn("modifiedAt").SortOrder("desc").Execute()
			require.NoError(t, err)
			require.NotNil(t, resp)
			//require.True(t, len(*resp.Items) == 20)

			for _, item := range resp.GetItems() {
				cb := item.GetCreatedBy()
				fmt.Println(item.GetName(), cb.GetName(), item.GetCreatedAt(), item.GetOwner().Id)

				doc := item.GetActualInstance().(*onshape.BTGlobalTreeNodeSummaryInfo)
				if doc == nil {
					t.Error("Error: expected item to be BTGlobalTreeNodeSummaryInfo, but it was actually", item.GetActualInstance())
				} else {
					require.True(t, doc.HasDefaultWorkspace())
				}
			}
		})
	}

}
