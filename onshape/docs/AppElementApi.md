# \AppElementApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortTransaction**](AppElementApi.md#AbortTransaction) | **Delete** /appelements/d/{did}/w/{wid}/e/{eid}/transactions/{tid} | Abort transaction by document ID, workspace ID, tab ID, and transaction ID.
[**BulkCreateElement**](AppElementApi.md#BulkCreateElement) | **Post** /appelements/d/{did}/w/{wid}/bulkcreate | Create multiple empty application tabs by document ID and workspace ID.
[**CommitTransactions**](AppElementApi.md#CommitTransactions) | **Post** /appelements/d/{did}/w/{wid}/transactions | Commit transactions by document ID and workspace ID.
[**CompareAppElementJson**](AppElementApi.md#CompareAppElementJson) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/compare | Compare JSON by document ID, workspace or version or microversion ID, and tab ID.
[**CreateElement**](AppElementApi.md#CreateElement) | **Post** /appelements/d/{did}/w/{wid} | Create application tab by document ID and workspace ID.
[**CreateReference**](AppElementApi.md#CreateReference) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references | Create referencesby document ID, workspace or version or microversion ID, and tab ID.
[**DeleteAppElementContent**](AppElementApi.md#DeleteAppElementContent) | **Delete** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/subelements/{sid} | Delete subelement array by document ID, workspace or version or microversion ID, tab ID, and subelement ID. A Subelement is used to store and organize data.
[**DeleteBlobSubelement**](AppElementApi.md#DeleteBlobSubelement) | **Delete** /appelements/d/{did}/w/{wid}/e/{eid}/blob/{bid} | Delete blob subelement file by document ID, workspace ID, tab ID, and blob ID. A Subelement is used to store and organize data.
[**DeleteReference**](AppElementApi.md#DeleteReference) | **Delete** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Delete references by document ID, workspace or version or microversion ID, tab ID, and resolve ID.
[**DownloadBlobSubelement**](AppElementApi.md#DownloadBlobSubelement) | **Get** /appelements/d/{did}/{vm}/{vmid}/e/{eid}/blob/{bid} | Download blob subelement file by document ID, version or microversion ID, tab ID, and blob ID. A Subelement is used to store and organize data.
[**DownloadBlobSubelementWorkspace**](AppElementApi.md#DownloadBlobSubelementWorkspace) | **Get** /appelements/d/{did}/w/{wid}/e/{eid}/blob/{bid} | Download blob subelement as a file by document ID, workspace ID, tab ID, and blob ID. A Subelement is used to store and organize data.
[**GetAppElementHistory**](AppElementApi.md#GetAppElementHistory) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/history | Retrieve history by document ID, workspace or version or microversion ID, and tab ID.
[**GetBlobSubelementIds**](AppElementApi.md#GetBlobSubelementIds) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/blob | Retrieve an array of blob subelement IDs by document ID and workspace or microversion ID. A Subelement is used to store and organize data.
[**GetElementTransactions**](AppElementApi.md#GetElementTransactions) | **Get** /appelements/d/{did}/w/{wid}/e/{eid}/transactions | Retrieve an array of tab transactions by document ID, workspace ID, and tab ID.
[**GetJson**](AppElementApi.md#GetJson) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/json | Retrieve JSON by document ID, workspace or version or microversion ID, and tab ID.
[**GetJsonPaths**](AppElementApi.md#GetJsonPaths) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/jsonpaths | Retrieve JSON paths by document ID, workspace or version or microversion ID, and tab ID.
[**GetSubElementContent**](AppElementApi.md#GetSubElementContent) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Retrieve subelement content by document ID, tab ID, and workspace or version or microversion ID. A Subelement is used to store and organize data.
[**GetSubelementIds**](AppElementApi.md#GetSubelementIds) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/ids | Retrieve subelement IDs by document ID, workspace or version or microversion ID, and tab ID. A Subelement is used to store and organize data.
[**ResolveReference**](AppElementApi.md#ResolveReference) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Resolve references by document ID, workspace or version or microversion ID, tab ID and resolve ID.
[**ResolveReferences**](AppElementApi.md#ResolveReferences) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/resolvereferences | Resolve references by document ID, workspace or version or microversion ID, and tab ID.
[**StartTransaction**](AppElementApi.md#StartTransaction) | **Post** /appelements/d/{did}/w/{wid}/e/{eid}/transactions | Start application tab transaction by document ID, workspace ID, and tab ID.
[**UpdateAppElement**](AppElementApi.md#UpdateAppElement) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Update application tab by document ID, workspace or version or microversion ID, and tab ID.
[**UpdateReference**](AppElementApi.md#UpdateReference) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Update references by document ID, workspace or version or microversion ID, tab ID, and resolve ID.
[**UploadBlobSubelement**](AppElementApi.md#UploadBlobSubelement) | **Post** /appelements/d/{did}/w/{wid}/e/{eid}/blob/{bid} | Upload blob subelement file by document ID, workspace ID, tab ID, and blob ID. A Subelement is used to store and organize data.



## AbortTransaction

> map[string]interface{} AbortTransaction(ctx, did, eid, wid, tid).ReturnError(returnError).Execute()

Abort transaction by document ID, workspace ID, tab ID, and transaction ID.

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
    resp, r, err := apiClient.AppElementApi.AbortTransaction(context.Background(), did, eid, wid, tid).ReturnError(returnError).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.AbortTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortTransaction`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.AbortTransaction`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateElement

> BTAppElementBulkCreateInfo BulkCreateElement(ctx, did, wid).BTAppElementBulkCreateParams(bTAppElementBulkCreateParams).LinkDocumentId(linkDocumentId).Execute()

Create multiple empty application tabs by document ID and workspace ID.



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
    resp, r, err := apiClient.AppElementApi.BulkCreateElement(context.Background(), did, wid).BTAppElementBulkCreateParams(bTAppElementBulkCreateParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.BulkCreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateElement`: BTAppElementBulkCreateInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.BulkCreateElement`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommitTransactions

> BTAppElementModifyInfo CommitTransactions(ctx, did, wid).BTAppElementCommitTransactionParams(bTAppElementCommitTransactionParams).LinkDocumentId(linkDocumentId).Execute()

Commit transactions by document ID and workspace ID.

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
    resp, r, err := apiClient.AppElementApi.CommitTransactions(context.Background(), did, wid).BTAppElementCommitTransactionParams(bTAppElementCommitTransactionParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.CommitTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommitTransactions`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.CommitTransactions`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompareAppElementJson

> BTDiffJsonResponse2725 CompareAppElementJson(ctx, did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).LinkDocumentId(linkDocumentId).JsonDifferenceFormat(jsonDifferenceFormat).Execute()

Compare JSON by document ID, workspace or version or microversion ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.CompareAppElementJson(context.Background(), did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).LinkDocumentId(linkDocumentId).JsonDifferenceFormat(jsonDifferenceFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.CompareAppElementJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompareAppElementJson`: BTDiffJsonResponse2725
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.CompareAppElementJson`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateElement

> BTAppElementModifyInfo CreateElement(ctx, did, wid).BTAppElementParams(bTAppElementParams).LinkDocumentId(linkDocumentId).Execute()

Create application tab by document ID and workspace ID.

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
    resp, r, err := apiClient.AppElementApi.CreateElement(context.Background(), did, wid).BTAppElementParams(bTAppElementParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.CreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateElement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.CreateElement`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReference

> BTAppElementReferenceInfo CreateReference(ctx, did, eid, wvm, wvmid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Create referencesby document ID, workspace or version or microversion ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.CreateReference(context.Background(), did, eid, wvm, wvmid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.CreateReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.CreateReference`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppElementContent

> BTAppElementModifyInfo DeleteAppElementContent(ctx, did, eid, wvm, wvmid, sid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete subelement array by document ID, workspace or version or microversion ID, tab ID, and subelement ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.DeleteAppElementContent(context.Background(), did, eid, wvm, wvmid, sid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.DeleteAppElementContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppElementContent`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.DeleteAppElementContent`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlobSubelement

> BTAppElementModifyInfo DeleteBlobSubelement(ctx, did, wid, eid, bid).TransactionId(transactionId).ChangeId(changeId).Execute()

Delete blob subelement file by document ID, workspace ID, tab ID, and blob ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.DeleteBlobSubelement(context.Background(), did, wid, eid, bid).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.DeleteBlobSubelement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlobSubelement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.DeleteBlobSubelement`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReference

> BTAppElementReferenceInfo DeleteReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete references by document ID, workspace or version or microversion ID, tab ID, and resolve ID.

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
    resp, r, err := apiClient.AppElementApi.DeleteReference(context.Background(), did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.DeleteReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.DeleteReference`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadBlobSubelement

> HttpFile DownloadBlobSubelement(ctx, did, vm, vmid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).LinkDocumentId(linkDocumentId).Execute()

Download blob subelement file by document ID, version or microversion ID, tab ID, and blob ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.DownloadBlobSubelement(context.Background(), did, vm, vmid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.DownloadBlobSubelement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadBlobSubelement`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.DownloadBlobSubelement`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadBlobSubelementWorkspace

> HttpFile DownloadBlobSubelementWorkspace(ctx, did, wid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).Execute()

Download blob subelement as a file by document ID, workspace ID, tab ID, and blob ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.DownloadBlobSubelementWorkspace(context.Background(), did, wid, eid, bid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.DownloadBlobSubelementWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadBlobSubelementWorkspace`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.DownloadBlobSubelementWorkspace`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppElementHistory

> BTAppElementHistoryInfo GetAppElementHistory(ctx, did, eid, wvm, wvmid).Execute()

Retrieve history by document ID, workspace or version or microversion ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.GetAppElementHistory(context.Background(), did, eid, wvm, wvmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.GetAppElementHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppElementHistory`: BTAppElementHistoryInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.GetAppElementHistory`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobSubelementIds

> BTAppElementIdsInfo GetBlobSubelementIds(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()

Retrieve an array of blob subelement IDs by document ID and workspace or microversion ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.GetBlobSubelementIds(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.GetBlobSubelementIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlobSubelementIds`: BTAppElementIdsInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.GetBlobSubelementIds`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementTransactions

> BTAppElementTransactionsInfo GetElementTransactions(ctx, did, eid, wid).Execute()

Retrieve an array of tab transactions by document ID, workspace ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.GetElementTransactions(context.Background(), did, eid, wid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.GetElementTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementTransactions`: BTAppElementTransactionsInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.GetElementTransactions`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJson

> BTGetJsonResponse2137 GetJson(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()

Retrieve JSON by document ID, workspace or version or microversion ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.GetJson(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.GetJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJson`: BTGetJsonResponse2137
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.GetJson`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsonPaths

> BTGetJsonPathsResponse1544 GetJsonPaths(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).BTGetJsonPaths1697(bTGetJsonPaths1697).Execute()

Retrieve JSON paths by document ID, workspace or version or microversion ID, and tab ID.

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
    bTGetJsonPaths1697 := *openapiclient.NewBTGetJsonPaths1697() // BTGetJsonPaths1697 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppElementApi.GetJsonPaths(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).BTGetJsonPaths1697(bTGetJsonPaths1697).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.GetJsonPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsonPaths`: BTGetJsonPathsResponse1544
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.GetJsonPaths`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetJsonPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 
 **bTGetJsonPaths1697** | [**BTGetJsonPaths1697**](BTGetJsonPaths1697.md) |  | 

### Return type

[**BTGetJsonPathsResponse1544**](BTGetJsonPathsResponse1544.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubElementContent

> BTAppElementContentInfo GetSubElementContent(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).BaseChangeId(baseChangeId).SubelementId(subelementId).Execute()

Retrieve subelement content by document ID, tab ID, and workspace or version or microversion ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.GetSubElementContent(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).BaseChangeId(baseChangeId).SubelementId(subelementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.GetSubElementContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubElementContent`: BTAppElementContentInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.GetSubElementContent`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubelementIds

> BTAppElementIdsInfo GetSubelementIds(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()

Retrieve subelement IDs by document ID, workspace or version or microversion ID, and tab ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.GetSubelementIds(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.GetSubelementIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubelementIds`: BTAppElementIdsInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.GetSubelementIds`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReference

> BTAppElementReferenceResolveInfo ResolveReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).Execute()

Resolve references by document ID, workspace or version or microversion ID, tab ID and resolve ID.

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
    resp, r, err := apiClient.AppElementApi.ResolveReference(context.Background(), did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.ResolveReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveReference`: BTAppElementReferenceResolveInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.ResolveReference`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReferences

> BTAppElementReferencesResolveInfo ResolveReferences(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).ReferenceIds(referenceIds).Execute()

Resolve references by document ID, workspace or version or microversion ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.ResolveReferences(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).ReferenceIds(referenceIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.ResolveReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveReferences`: BTAppElementReferencesResolveInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.ResolveReferences`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTransaction

> BTAppElementModifyInfo StartTransaction(ctx, did, eid, wid).BTAppElementStartTransactionParams(bTAppElementStartTransactionParams).Execute()

Start application tab transaction by document ID, workspace ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.StartTransaction(context.Background(), did, eid, wid).BTAppElementStartTransactionParams(bTAppElementStartTransactionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.StartTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartTransaction`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.StartTransaction`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppElement

> BTAppElementModifyInfo UpdateAppElement(ctx, did, eid, wvm, wvmid).BTAppElementUpdateParams(bTAppElementUpdateParams).Execute()

Update application tab by document ID, workspace or version or microversion ID, and tab ID.

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
    resp, r, err := apiClient.AppElementApi.UpdateAppElement(context.Background(), did, eid, wvm, wvmid).BTAppElementUpdateParams(bTAppElementUpdateParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.UpdateAppElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppElement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.UpdateAppElement`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReference

> BTAppElementReferenceInfo UpdateReference(ctx, did, eid, wvm, wvmid, rid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Update references by document ID, workspace or version or microversion ID, tab ID, and resolve ID.

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
    resp, r, err := apiClient.AppElementApi.UpdateReference(context.Background(), did, eid, wvm, wvmid, rid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.UpdateReference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReference`: BTAppElementReferenceInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.UpdateReference`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadBlobSubelement

> BTAppElementModifyInfo UploadBlobSubelement(ctx, did, wid, eid, bid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).File(file).FileContentLength(fileContentLength).Execute()

Upload blob subelement file by document ID, workspace ID, tab ID, and blob ID. A Subelement is used to store and organize data.

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
    resp, r, err := apiClient.AppElementApi.UploadBlobSubelement(context.Background(), did, wid, eid, bid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).File(file).FileContentLength(fileContentLength).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppElementApi.UploadBlobSubelement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadBlobSubelement`: BTAppElementModifyInfo
    fmt.Fprintf(os.Stdout, "Response from `AppElementApi.UploadBlobSubelement`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

