# \AppElementAPI

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortTransaction**](AppElementAPI.md#AbortTransaction) | **Delete** /appelements/d/{did}/w/{wid}/e/{eid}/transactions/{tid} | Abort a transaction.
[**BulkCreateElement**](AppElementAPI.md#BulkCreateElement) | **Post** /appelements/d/{did}/w/{wid}/bulkcreate | Create multiple empty application elements at once.
[**CommitTransactions**](AppElementAPI.md#CommitTransactions) | **Post** /appelements/d/{did}/w/{wid}/transactions | Merge multiple transactions into one microversion.
[**CompareAppElementJson**](AppElementAPI.md#CompareAppElementJson) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/compare | Compare app element JSON trees between workspaces/versions/microversions in a document.
[**CreateElement**](AppElementAPI.md#CreateElement) | **Post** /appelements/d/{did}/w/{wid} | Create a new application element.
[**CreateReference**](AppElementAPI.md#CreateReference) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references | Creates a reference to an app element.
[**DeleteAppElementContent**](AppElementAPI.md#DeleteAppElementContent) | **Delete** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/subelements/{sid} | Deletes the content from the specified app element.
[**DeleteAppElementContentBatch**](AppElementAPI.md#DeleteAppElementContentBatch) | **Delete** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/subelements | Delete multiple subelements array by document ID, workspace or version or microversion ID, tab ID, and subelement IDs.
[**DeleteBlobSubelement**](AppElementAPI.md#DeleteBlobSubelement) | **Delete** /appelements/d/{did}/w/{wid}/e/{eid}/blob/{bid} | Delete a blob subelement from an app element.
[**DeleteReference**](AppElementAPI.md#DeleteReference) | **Delete** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Delete an app element reference.
[**DownloadBlobSubelement**](AppElementAPI.md#DownloadBlobSubelement) | **Get** /appelements/d/{did}/{vm}/{vmid}/e/{eid}/blob/{bid} | Download a blob subelement from the specified app element.
[**DownloadBlobSubelementWorkspace**](AppElementAPI.md#DownloadBlobSubelementWorkspace) | **Get** /appelements/d/{did}/w/{wid}/e/{eid}/blob/{bid} | Download the blob element (i.e., a file) stored in an app element in a document&#39;s workspace.
[**GetAppElementHistory**](AppElementAPI.md#GetAppElementHistory) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/history | Get the history of the specified all element.
[**GetBlobSubelementIds**](AppElementAPI.md#GetBlobSubelementIds) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/blob | Get a list of all blob subelement IDs for the specified workspace, version, or microversion.
[**GetElementTransactions**](AppElementAPI.md#GetElementTransactions) | **Get** /appelements/d/{did}/w/{wid}/e/{eid}/transactions | Get a list of all transactions performed on an element.
[**GetJson**](AppElementAPI.md#GetJson) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/json | Get the full JSON tree for the specified workspace/version/microversion.
[**GetJsonPaths**](AppElementAPI.md#GetJsonPaths) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/jsonpaths | Get the JSON at specified paths for an element.
[**GetSubElementContent**](AppElementAPI.md#GetSubElementContent) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Get a list of all subelement IDs in a specified workspace/version/microversion.
[**GetSubelementIds**](AppElementAPI.md#GetSubelementIds) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/ids | Get a list of all subelement IDs in a specified workspace/version/microversion.
[**ResolveReference**](AppElementAPI.md#ResolveReference) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Resolves a single reference to an app element.
[**ResolveReferences**](AppElementAPI.md#ResolveReferences) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/resolvereferences | Resolves bulk app element references.
[**StartTransaction**](AppElementAPI.md#StartTransaction) | **Post** /appelements/d/{did}/w/{wid}/e/{eid}/transactions | Start a transaction
[**UpdateAppElement**](AppElementAPI.md#UpdateAppElement) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Update the content for the specified app element.
[**UpdateReference**](AppElementAPI.md#UpdateReference) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Update an app element reference.
[**UploadBlobSubelement**](AppElementAPI.md#UploadBlobSubelement) | **Post** /appelements/d/{did}/w/{wid}/e/{eid}/blob/{bid} | Create a new blob subelement from an uploaded file.



## AbortTransaction

> map[string]interface{} AbortTransaction(ctx, did, eid, wid, tid).ReturnError(returnError).Execute()

Abort a transaction.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wid := "wid_example" // string | 
    tid := "tid_example" // string | 
    returnError := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.AbortTransaction(context.Background(), did, eid, wid, tid).ReturnError(returnError).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.AbortTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortTransaction`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.AbortTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **returnError** | **bool** |  | [default to true]

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateElement

> BTAppElementBulkCreateInfo BulkCreateElement(ctx, did, wid).BTAppElementBulkCreateParams(bTAppElementBulkCreateParams).LinkDocumentId(linkDocumentId).Execute()

Create multiple empty application elements at once.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    bTAppElementBulkCreateParams := *openapiclient.NewBTAppElementBulkCreateParams("FormatId_example") // BTAppElementBulkCreateParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.BulkCreateElement(context.Background(), did, wid).BTAppElementBulkCreateParams(bTAppElementBulkCreateParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.BulkCreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateElement`: BTAppElementBulkCreateInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.BulkCreateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTAppElementBulkCreateParams** | [**BTAppElementBulkCreateParams**](BTAppElementBulkCreateParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTAppElementBulkCreateInfo**](BTAppElementBulkCreateInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommitTransactions

> BTAppElementModifyInfo CommitTransactions(ctx, did, wid).BTAppElementCommitTransactionParams(bTAppElementCommitTransactionParams).LinkDocumentId(linkDocumentId).Execute()

Merge multiple transactions into one microversion.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    bTAppElementCommitTransactionParams := *openapiclient.NewBTAppElementCommitTransactionParams() // BTAppElementCommitTransactionParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.CommitTransactions(context.Background(), did, wid).BTAppElementCommitTransactionParams(bTAppElementCommitTransactionParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.CommitTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommitTransactions`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.CommitTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTAppElementCommitTransactionParams** | [**BTAppElementCommitTransactionParams**](BTAppElementCommitTransactionParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompareAppElementJson

> BTDiffJsonResponse2725 CompareAppElementJson(ctx, did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).LinkDocumentId(linkDocumentId).JsonDifferenceFormat(jsonDifferenceFormat).Execute()

Compare app element JSON trees between workspaces/versions/microversions in a document.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 
    workspaceId := "workspaceId_example" // string |  (optional)
    versionId := "versionId_example" // string |  (optional)
    microversionId := "microversionId_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    jsonDifferenceFormat := "jsonDifferenceFormat_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.CompareAppElementJson(context.Background(), did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).LinkDocumentId(linkDocumentId).JsonDifferenceFormat(jsonDifferenceFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.CompareAppElementJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompareAppElementJson`: BTDiffJsonResponse2725
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.CompareAppElementJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompareAppElementJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **workspaceId** | **string** |  | 
 **versionId** | **string** |  | 
 **microversionId** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **jsonDifferenceFormat** | **string** |  | 

### Return type

[**BTDiffJsonResponse2725**](BTDiffJsonResponse2725.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateElement

> BTAppElementModifyInfo CreateElement(ctx, did, wid).BTAppElementParams(bTAppElementParams).LinkDocumentId(linkDocumentId).Execute()

Create a new application element.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    bTAppElementParams := *openapiclient.NewBTAppElementParams("FormatId_example") // BTAppElementParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.CreateElement(context.Background(), did, wid).BTAppElementParams(bTAppElementParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.CreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateElement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.CreateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTAppElementParams** | [**BTAppElementParams**](BTAppElementParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReference

> BTAppElementReferenceInfo CreateReference(ctx, did, eid, wvm, wvmid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Creates a reference to an app element.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    bTAppElementReferenceParams := *openapiclient.NewBTAppElementReferenceParams() // BTAppElementReferenceParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.CreateReference(context.Background(), did, eid, wvm, wvmid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.CreateReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.CreateReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementReferenceParams** | [**BTAppElementReferenceParams**](BTAppElementReferenceParams.md) |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppElementContent

> BTAppElementModifyInfo DeleteAppElementContent(ctx, did, eid, wvm, wvmid, sid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Deletes the content from the specified app element.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    sid := "sid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    description := "description_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.DeleteAppElementContent(context.Background(), did, eid, wvm, wvmid, sid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.DeleteAppElementContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppElementContent`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.DeleteAppElementContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppElementContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppElementContentBatch

> BTAppElementModifyInfo DeleteAppElementContentBatch(ctx, did, eid, wvm, wvmid).SubelementIds(subelementIds).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete multiple subelements array by document ID, workspace or version or microversion ID, tab ID, and subelement IDs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    subelementIds := []string{"Inner_example"} // []string |  (optional)
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    description := "description_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.DeleteAppElementContentBatch(context.Background(), did, eid, wvm, wvmid).SubelementIds(subelementIds).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.DeleteAppElementContentBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppElementContentBatch`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.DeleteAppElementContentBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppElementContentBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **subelementIds** | **[]string** |  | 
 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlobSubelement

> BTAppElementModifyInfo DeleteBlobSubelement(ctx, did, wid, eid, bid).TransactionId(transactionId).ChangeId(changeId).Execute()

Delete a blob subelement from an app element.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    eid := "eid_example" // string | 
    bid := "bid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.DeleteBlobSubelement(context.Background(), did, wid, eid, bid).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.DeleteBlobSubelement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlobSubelement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.DeleteBlobSubelement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlobSubelementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReference

> BTAppElementReferenceInfo DeleteReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete an app element reference.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    rid := "rid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    description := "description_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.DeleteReference(context.Background(), did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.DeleteReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.DeleteReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadBlobSubelement

> HttpFile DownloadBlobSubelement(ctx, did, vm, vmid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).LinkDocumentId(linkDocumentId).Execute()

Download a blob subelement from the specified app element.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    vm := "vm_example" // string | 
    vmid := "vmid_example" // string | 
    eid := "eid_example" // string | 
    bid := "bid_example" // string | 
    contentDisposition := "contentDisposition_example" // string |  (optional)
    ifNoneMatch := "ifNoneMatch_example" // string |  (optional)
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.DownloadBlobSubelement(context.Background(), did, vm, vmid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.DownloadBlobSubelement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadBlobSubelement`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.DownloadBlobSubelement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vm** | **string** |  | 
**vmid** | **string** |  | 
**eid** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBlobSubelementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **contentDisposition** | **string** |  | 
 **ifNoneMatch** | **string** |  | 
 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**HttpFile**](HttpFile.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadBlobSubelementWorkspace

> HttpFile DownloadBlobSubelementWorkspace(ctx, did, wid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).Execute()

Download the blob element (i.e., a file) stored in an app element in a document's workspace.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    eid := "eid_example" // string | 
    bid := "bid_example" // string | 
    contentDisposition := "contentDisposition_example" // string |  (optional)
    ifNoneMatch := "ifNoneMatch_example" // string |  (optional)
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.DownloadBlobSubelementWorkspace(context.Background(), did, wid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.DownloadBlobSubelementWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadBlobSubelementWorkspace`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.DownloadBlobSubelementWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBlobSubelementWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **contentDisposition** | **string** |  | 
 **ifNoneMatch** | **string** |  | 
 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 

### Return type

[**HttpFile**](HttpFile.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppElementHistory

> BTAppElementHistoryInfo GetAppElementHistory(ctx, did, eid, wvm, wvmid).Execute()

Get the history of the specified all element.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.GetAppElementHistory(context.Background(), did, eid, wvm, wvmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.GetAppElementHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppElementHistory`: BTAppElementHistoryInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.GetAppElementHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppElementHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTAppElementHistoryInfo**](BTAppElementHistoryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobSubelementIds

> BTAppElementIdsInfo GetBlobSubelementIds(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()

Get a list of all blob subelement IDs for the specified workspace, version, or microversion.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.GetBlobSubelementIds(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.GetBlobSubelementIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlobSubelementIds`: BTAppElementIdsInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.GetBlobSubelementIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobSubelementIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 

### Return type

[**BTAppElementIdsInfo**](BTAppElementIdsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementTransactions

> BTAppElementTransactionsInfo GetElementTransactions(ctx, did, eid, wid).Execute()

Get a list of all transactions performed on an element.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wid := "wid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.GetElementTransactions(context.Background(), did, eid, wid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.GetElementTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementTransactions`: BTAppElementTransactionsInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.GetElementTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BTAppElementTransactionsInfo**](BTAppElementTransactionsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJson

> BTGetJsonResponse2137 GetJson(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()

Get the full JSON tree for the specified workspace/version/microversion.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wvm := "wvm_example" // string | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    transactionId := "transactionId_example" // string | The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API. (optional)
    changeId := "changeId_example" // string | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.GetJson(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.GetJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJson`: BTGetJsonResponse2137
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.GetJson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** | The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API. | 
 **changeId** | **string** | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. | 

### Return type

[**BTGetJsonResponse2137**](BTGetJsonResponse2137.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsonPaths

> BTGetJsonPathsResponse1544 GetJsonPaths(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).BTGetJsonPaths1697(bTGetJsonPaths1697).Execute()

Get the JSON at specified paths for an element.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wvm := "wvm_example" // string | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)
    bTGetJsonPaths1697 := *openapiclient.NewBTGetJsonPaths1697() // BTGetJsonPaths1697 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.GetJsonPaths(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).BTGetJsonPaths1697(bTGetJsonPaths1697).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.GetJsonPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonPaths`: BTGetJsonPathsResponse1544
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.GetJsonPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsonPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 
 **bTGetJsonPaths1697** | [**BTGetJsonPaths1697**](BTGetJsonPaths1697.md) |  | 

### Return type

[**BTGetJsonPathsResponse1544**](BTGetJsonPathsResponse1544.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubElementContent

> BTAppElementContentInfo GetSubElementContent(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).BaseChangeId(baseChangeId).SubelementId(subelementId).Execute()

Get a list of all subelement IDs in a specified workspace/version/microversion.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wvm := "wvm_example" // string | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)
    baseChangeId := "baseChangeId_example" // string |  (optional)
    subelementId := "subelementId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.GetSubElementContent(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).BaseChangeId(baseChangeId).SubelementId(subelementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.GetSubElementContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubElementContent`: BTAppElementContentInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.GetSubElementContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubElementContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 
 **baseChangeId** | **string** |  | 
 **subelementId** | **string** |  | 

### Return type

[**BTAppElementContentInfo**](BTAppElementContentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubelementIds

> BTAppElementIdsInfo GetSubelementIds(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()

Get a list of all subelement IDs in a specified workspace/version/microversion.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    changeId := "changeId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.GetSubelementIds(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.GetSubelementIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubelementIds`: BTAppElementIdsInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.GetSubelementIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubelementIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 

### Return type

[**BTAppElementIdsInfo**](BTAppElementIdsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReference

> BTAppElementReferenceResolveInfo ResolveReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).Execute()

Resolves a single reference to an app element.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    rid := "rid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    includeInternal := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.ResolveReference(context.Background(), did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.ResolveReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveReference`: BTAppElementReferenceResolveInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.ResolveReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **includeInternal** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTAppElementReferenceResolveInfo**](BTAppElementReferenceResolveInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReferences

> BTAppElementReferencesResolveInfo ResolveReferences(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).ReferenceIds(referenceIds).Execute()

Resolves bulk app element references.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | The id of the document in which to perform the operation.
    wvm := "wvm_example" // string | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    includeInternal := true // bool |  (optional) (default to false)
    referenceIds := "referenceIds_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.ResolveReferences(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).ReferenceIds(referenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.ResolveReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveReferences`: BTAppElementReferencesResolveInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.ResolveReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **includeInternal** | **bool** |  | [default to false]
 **referenceIds** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppElementReferencesResolveInfo**](BTAppElementReferencesResolveInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTransaction

> BTAppElementModifyInfo StartTransaction(ctx, did, eid, wid).BTAppElementStartTransactionParams(bTAppElementStartTransactionParams).Execute()

Start a transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wid := "wid_example" // string | 
    bTAppElementStartTransactionParams := *openapiclient.NewBTAppElementStartTransactionParams() // BTAppElementStartTransactionParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.StartTransaction(context.Background(), did, eid, wid).BTAppElementStartTransactionParams(bTAppElementStartTransactionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.StartTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartTransaction`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.StartTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAppElementStartTransactionParams** | [**BTAppElementStartTransactionParams**](BTAppElementStartTransactionParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppElement

> BTAppElementModifyInfo UpdateAppElement(ctx, did, eid, wvm, wvmid).BTAppElementUpdateParams(bTAppElementUpdateParams).Execute()

Update the content for the specified app element.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    bTAppElementUpdateParams := *openapiclient.NewBTAppElementUpdateParams() // BTAppElementUpdateParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.UpdateAppElement(context.Background(), did, eid, wvm, wvmid).BTAppElementUpdateParams(bTAppElementUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.UpdateAppElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppElement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.UpdateAppElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementUpdateParams** | [**BTAppElementUpdateParams**](BTAppElementUpdateParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReference

> BTAppElementReferenceInfo UpdateReference(ctx, did, eid, wvm, wvmid, rid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Update an app element reference.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    rid := "rid_example" // string | 
    bTAppElementReferenceParams := *openapiclient.NewBTAppElementReferenceParams() // BTAppElementReferenceParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.UpdateReference(context.Background(), did, eid, wvm, wvmid, rid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.UpdateReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.UpdateReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bTAppElementReferenceParams** | [**BTAppElementReferenceParams**](BTAppElementReferenceParams.md) |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBlobSubelement

> BTAppElementModifyInfo UploadBlobSubelement(ctx, did, wid, eid, bid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).File(file).FileContentLength(fileContentLength).Execute()

Create a new blob subelement from an uploaded file.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    eid := "eid_example" // string | 
    bid := "bid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional)
    parentChangeId := "parentChangeId_example" // string |  (optional)
    description := "description_example" // string |  (optional)
    file := os.NewFile(1234, "some_file") // HttpFile | File to upload. (optional)
    fileContentLength := int64(789) // int64 |  (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementAPI.UploadBlobSubelement(context.Background(), did, wid, eid, bid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).File(file).FileContentLength(fileContentLength).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementAPI.UploadBlobSubelement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBlobSubelement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementAPI.UploadBlobSubelement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**bid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadBlobSubelementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 
 **file** | **HttpFile** | File to upload. | 
 **fileContentLength** | **int64** |  | [default to -1]

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

