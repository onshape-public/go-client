# Go Client for Onshape public API

## The simplest possible thing to try: Creating a document
```Go
package main

import (
	"context"
	"fmt"

	"github.com/onshape-public/go-client/onshape"
)

func main() {
	config := onshape.NewConfiguration()
	config.Debug = true

	client := onshape.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), onshape.ContextAPIKeys,
		onshape.APIKeys{"Your_Secret_Key", "Your_Access_Key"})

	docParams := onshape.NewBTDocumentParams()
	docParams.SetName("Name For Your Document")

	docInfo, rawResp, err := client.DocumentsApi.CreateDocument(ctx).BTDocumentParams(*docParams).Execute()

	if err != nil || (rawResp != nil && rawResp.StatusCode >= 300) {
		fmt.Print("err: ", err, " -- Response status: ", rawResp)
	} else {
		fmt.Println("Created a document w/the name: ", docInfo.GetName())
    }
}
```

## The Detailed API Documentation
Could be found [here](./onshape/README.md)

## Breaking Changes
### v1.201.59934-5fdafe848d7f
* Replaces BTDocumentSummaryInfo with BTGlobalTreeNodeSummaryInfo
### v1.193.50007-3a3fc4483ede
* UpdateDocumentAttributes now returns BTDocumentSummaryInfo instead of map[string]interface
### v1.171.24257-687de06de652
* ServerVariable was renamed to APIServerVariable after OnShape introduced a data structure with the same name
### v1.189.45939-469a4bd6b788
* map replaces some BT structure when appropriate for simplicity
* discriminators are added as embedded structs. Adjustment is needed for successful deserialization
