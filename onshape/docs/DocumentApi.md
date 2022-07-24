# \DocumentApi

All URIs are relative to *https://staging.dev.onshape.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyWorkspace**](DocumentApi.md#CopyWorkspace) | **Post** /documents/{did}/workspaces/{wid}/copy | Copy workspace by document ID and workspace ID.
[**CreateDocument**](DocumentApi.md#CreateDocument) | **Post** /documents | Create and upload a document.
[**CreateVersion**](DocumentApi.md#CreateVersion) | **Post** /documents/d/{did}/versions | Create version by document ID.
[**CreateWorkspace**](DocumentApi.md#CreateWorkspace) | **Post** /documents/d/{did}/workspaces | Create workspace by document ID.
[**DeleteDocument**](DocumentApi.md#DeleteDocument) | **Delete** /documents/{did} | Delete document by document ID.
[**DeleteWorkspace**](DocumentApi.md#DeleteWorkspace) | **Delete** /documents/d/{did}/workspaces/{wid} | Delete workspace by document ID and workspace ID.
[**DownloadExternalData**](DocumentApi.md#DownloadExternalData) | **Get** /documents/d/{did}/externaldata/{fid} | Retrieve external data by document ID and foreign ID.
[**Export2Json**](DocumentApi.md#Export2Json) | **Post** /documents/d/{did}/{wv}/{wvid}/e/{eid}/export | Export document by document ID, workspace or version ID, and tab ID.
[**GetCurrentMicroversion**](DocumentApi.md#GetCurrentMicroversion) | **Get** /documents/d/{did}/{wv}/{wvid}/currentmicroversion | Retrieve current microversion by document ID and workspace or version ID.
[**GetDocument**](DocumentApi.md#GetDocument) | **Get** /documents/{did} | Retrieve document by document ID.
[**GetDocumentAcl**](DocumentApi.md#GetDocumentAcl) | **Get** /documents/{did}/acl | Retrieve access control list by document ID.
[**GetDocumentHistory**](DocumentApi.md#GetDocumentHistory) | **Get** /documents/d/{did}/{wm}/{wmid}/documenthistory | Retrieve document history by document ID and workspace or microversion ID.
[**GetDocumentPermissionSet**](DocumentApi.md#GetDocumentPermissionSet) | **Get** /documents/{did}/permissionset | Retrieve Document permissions by document ID.
[**GetDocumentVersions**](DocumentApi.md#GetDocumentVersions) | **Get** /documents/d/{did}/versions | Retrieve versions by document ID.
[**GetDocumentWorkspaces**](DocumentApi.md#GetDocumentWorkspaces) | **Get** /documents/d/{did}/workspaces | Retrieve workspaces by document ID.
[**GetDocuments**](DocumentApi.md#GetDocuments) | **Get** /documents | Retrieve a document.
[**GetElementsInDocument**](DocumentApi.md#GetElementsInDocument) | **Get** /documents/d/{did}/{wvm}/{wvmid}/elements | Retrieve tabs by document ID and workspace or version or microversion ID.
[**GetInsertables**](DocumentApi.md#GetInsertables) | **Get** /documents/d/{did}/{wv}/{wvid}/insertables | Retrieve insertables by document ID and workspace or version ID.
[**GetUnitInfo**](DocumentApi.md#GetUnitInfo) | **Get** /documents/d/{did}/{wvm}/{wvmid}/unitinfo | Get the selected units and precision by document ID and workspace or version or microversion ID.
[**GetVersion**](DocumentApi.md#GetVersion) | **Get** /documents/d/{did}/versions/{vid} | Retrieve version by document ID and version ID.
[**MergeIntoWorkspace**](DocumentApi.md#MergeIntoWorkspace) | **Post** /documents/{did}/workspaces/{wid}/merge | Merge into workspace by document ID and workspace ID.
[**MergePreview**](DocumentApi.md#MergePreview) | **Get** /documents/{did}/w/{wid}/mergePreview | Merge preview of changes that will occur based on document ID, workspace ID and source workspace/version ID
[**MoveElementsToDocument**](DocumentApi.md#MoveElementsToDocument) | **Post** /documents/d/{did}/w/{wid}/moveelement | Move tab by document ID and workspace ID.
[**RestoreFromHistory**](DocumentApi.md#RestoreFromHistory) | **Post** /documents/{did}/w/{wid}/restore/{vm}/{vmid} | Restore version or microversion to workspace by document ID, workspace ID, and version or microversion ID.
[**RevertUnchangedToRevisions**](DocumentApi.md#RevertUnchangedToRevisions) | **Post** /documents/d/{did}/w/{wid}/revertunchangedtorevisions | 
[**Search**](DocumentApi.md#Search) | **Post** /documents/search | Search document.
[**ShareDocument**](DocumentApi.md#ShareDocument) | **Post** /documents/{did}/share | Share document by document ID.
[**SyncApplicationElements**](DocumentApi.md#SyncApplicationElements) | **Post** /documents/d/{did}/w/{wid}/syncAppElements | 
[**UnShareDocument**](DocumentApi.md#UnShareDocument) | **Delete** /documents/{did}/share/{eid} | Unshare document by document ID and tab ID.
[**UpdateDocumentAttributes**](DocumentApi.md#UpdateDocumentAttributes) | **Post** /documents/{did} | Update document attributes by document ID.
[**UpdateExternalReferencesToLatestDocuments**](DocumentApi.md#UpdateExternalReferencesToLatestDocuments) | **Post** /documents/d/{did}/w/{wid}/e/{eid}/latestdocumentreferences | Update external references to latest by document ID, workspace ID, and tab ID.



## CopyWorkspace

> BTCopyDocumentInfo CopyWorkspace(ctx, did, wid).BTCopyDocumentParams(bTCopyDocumentParams).Execute()

Copy workspace by document ID and workspace ID.

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
    bTCopyDocumentParams := *openapiclient.NewBTCopyDocumentParams() // BTCopyDocumentParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.CopyWorkspace(context.Background(), did, wid).BTCopyDocumentParams(bTCopyDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.CopyWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyWorkspace`: BTCopyDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.CopyWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTCopyDocumentParams** | [**BTCopyDocumentParams**](BTCopyDocumentParams.md) |  | 

### Return type

[**BTCopyDocumentInfo**](BTCopyDocumentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocument

> BTDocumentInfo CreateDocument(ctx).BTDocumentParams(bTDocumentParams).Execute()

Create and upload a document.

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
    bTDocumentParams := *openapiclient.NewBTDocumentParams() // BTDocumentParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.CreateDocument(context.Background()).BTDocumentParams(bTDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.CreateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocument`: BTDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.CreateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTDocumentParams** | [**BTDocumentParams**](BTDocumentParams.md) |  | 

### Return type

[**BTDocumentInfo**](BTDocumentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVersion

> BTVersionInfo CreateVersion(ctx, did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()

Create version by document ID.

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
    bTVersionOrWorkspaceParams := *openapiclient.NewBTVersionOrWorkspaceParams() // BTVersionOrWorkspaceParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.CreateVersion(context.Background(), did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.CreateVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVersion`: BTVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.CreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTVersionOrWorkspaceParams** | [**BTVersionOrWorkspaceParams**](BTVersionOrWorkspaceParams.md) |  | 

### Return type

[**BTVersionInfo**](BTVersionInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> BTWorkspaceInfo CreateWorkspace(ctx, did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()

Create workspace by document ID.

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
    bTVersionOrWorkspaceParams := *openapiclient.NewBTVersionOrWorkspaceParams() // BTVersionOrWorkspaceParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.CreateWorkspace(context.Background(), did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.CreateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkspace`: BTWorkspaceInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTVersionOrWorkspaceParams** | [**BTVersionOrWorkspaceParams**](BTVersionOrWorkspaceParams.md) |  | 

### Return type

[**BTWorkspaceInfo**](BTWorkspaceInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> map[string]interface{} DeleteDocument(ctx, did).Forever(forever).Execute()

Delete document by document ID.

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
    forever := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.DeleteDocument(context.Background(), did).Forever(forever).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.DeleteDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDocument`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.DeleteDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forever** | **bool** |  | [default to false]

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


## DeleteWorkspace

> map[string]interface{} DeleteWorkspace(ctx, did, wid).Execute()

Delete workspace by document ID and workspace ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.DeleteWorkspace(context.Background(), did, wid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.DeleteWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkspace`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.DeleteWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DownloadExternalData

> HttpFile DownloadExternalData(ctx, did, fid).IfNoneMatch(ifNoneMatch).Execute()

Retrieve external data by document ID and foreign ID.

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
    fid := "fid_example" // string | 
    ifNoneMatch := "ifNoneMatch_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.DownloadExternalData(context.Background(), did, fid).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.DownloadExternalData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadExternalData`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.DownloadExternalData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadExternalDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** |  | 

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


## Export2Json

> map[string]interface{} Export2Json(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).BTBExportModelParams(bTBExportModelParams).Execute()

Export document by document ID, workspace or version ID, and tab ID.

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
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    bTBExportModelParams := *openapiclient.NewBTBExportModelParams("DocumentId_example", "Format_example") // BTBExportModelParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.Export2Json(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).BTBExportModelParams(bTBExportModelParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.Export2Json``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export2Json`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.Export2Json`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExport2JsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **bTBExportModelParams** | [**BTBExportModelParams**](BTBExportModelParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentMicroversion

> BTMicroversionInfo GetCurrentMicroversion(ctx, did, wv, wvid).Execute()

Retrieve current microversion by document ID and workspace or version ID.

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
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetCurrentMicroversion(context.Background(), did, wv, wvid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetCurrentMicroversion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentMicroversion`: BTMicroversionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetCurrentMicroversion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentMicroversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BTMicroversionInfo**](BTMicroversionInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> BTDocumentInfo GetDocument(ctx, did).Execute()

Retrieve document by document ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetDocument(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocument`: BTDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTDocumentInfo**](BTDocumentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentAcl

> BTAclInfo GetDocumentAcl(ctx, did).Execute()

Retrieve access control list by document ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetDocumentAcl(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetDocumentAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentAcl`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetDocumentAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTAclInfo**](BTAclInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentHistory

> []BTDocumentHistoryInfo GetDocumentHistory(ctx, did, wm, wmid).Execute()

Retrieve document history by document ID and workspace or microversion ID.

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
    wm := "wm_example" // string | 
    wmid := "wmid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetDocumentHistory(context.Background(), did, wm, wmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetDocumentHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentHistory`: []BTDocumentHistoryInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetDocumentHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wm** | **string** |  | 
**wmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]BTDocumentHistoryInfo**](BTDocumentHistoryInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentPermissionSet

> []string GetDocumentPermissionSet(ctx, did).Execute()

Retrieve Document permissions by document ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetDocumentPermissionSet(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetDocumentPermissionSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentPermissionSet`: []string
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetDocumentPermissionSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentPermissionSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentVersions

> []BTVersionInfo GetDocumentVersions(ctx, did).Offset(offset).Limit(limit).Execute()

Retrieve versions by document ID.

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
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetDocumentVersions(context.Background(), did).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetDocumentVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentVersions`: []BTVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetDocumentVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 0]

### Return type

[**[]BTVersionInfo**](BTVersionInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentWorkspaces

> []BTWorkspaceInfo GetDocumentWorkspaces(ctx, did).Execute()

Retrieve workspaces by document ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetDocumentWorkspaces(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetDocumentWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentWorkspaces`: []BTWorkspaceInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetDocumentWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BTWorkspaceInfo**](BTWorkspaceInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocuments

> BTGlobalTreeNodeListResponse GetDocuments(ctx).Q(q).Filter(filter).Owner(owner).OwnerType(ownerType).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Label(label).Project(project).ParentId(parentId).Execute()

Retrieve a document.

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
    q := "q_example" // string |  (optional) (default to "")
    filter := int32(56) // int32 |  (optional)
    owner := "owner_example" // string |  (optional) (default to "")
    ownerType := int32(56) // int32 |  (optional) (default to 1)
    sortColumn := "sortColumn_example" // string |  (optional) (default to "createdAt")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "desc")
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)
    label := "label_example" // string |  (optional)
    project := "project_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetDocuments(context.Background()).Q(q).Filter(filter).Owner(owner).OwnerType(ownerType).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Label(label).Project(project).ParentId(parentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocuments`: BTGlobalTreeNodeListResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | [default to &quot;&quot;]
 **filter** | **int32** |  | 
 **owner** | **string** |  | [default to &quot;&quot;]
 **ownerType** | **int32** |  | [default to 1]
 **sortColumn** | **string** |  | [default to &quot;createdAt&quot;]
 **sortOrder** | **string** |  | [default to &quot;desc&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **label** | **string** |  | 
 **project** | **string** |  | 
 **parentId** | **string** |  | 

### Return type

[**BTGlobalTreeNodeListResponse**](BTGlobalTreeNodeListResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementsInDocument

> []BTDocumentElementInfo GetElementsInDocument(ctx, did, wvm, wvmid).ElementType(elementType).ElementId(elementId).WithThumbnails(withThumbnails).LinkDocumentId(linkDocumentId).Execute()

Retrieve tabs by document ID and workspace or version or microversion ID.

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
    elementType := "elementType_example" // string |  (optional) (default to "")
    elementId := "elementId_example" // string |  (optional) (default to "")
    withThumbnails := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetElementsInDocument(context.Background(), did, wvm, wvmid).ElementType(elementType).ElementId(elementId).WithThumbnails(withThumbnails).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetElementsInDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementsInDocument`: []BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetElementsInDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementsInDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **elementType** | **string** |  | [default to &quot;&quot;]
 **elementId** | **string** |  | [default to &quot;&quot;]
 **withThumbnails** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**[]BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsertables

> BTInsertablesListResponse GetInsertables(ctx, did, wv, wvid).ElementId(elementId).Configuration(configuration).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatureStudios(includeFeatureStudios).IncludeBlobs(includeBlobs).AllowedBlobMimeTypes(allowedBlobMimeTypes).ExcludeNewerFSVersions(excludeNewerFSVersions).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludePartStudios(includePartStudios).IncludeFeatures(includeFeatures).IncludeMeshes(includeMeshes).IncludeWires(includeWires).IncludeFlattenedBodies(includeFlattenedBodies).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).IncludeFSComputedPartPropertyFunctions(includeFSComputedPartPropertyFunctions).IncludeVariables(includeVariables).IncludeVariableStudios(includeVariableStudios).AllowedBlobExtensions(allowedBlobExtensions).Execute()

Retrieve insertables by document ID and workspace or version ID.

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
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    elementId := "elementId_example" // string |  (optional)
    configuration := "configuration_example" // string |  (optional)
    betaCapabilityIds := []string{"Inner_example"} // []string |  (optional)
    includeParts := true // bool |  (optional) (default to false)
    includeSurfaces := true // bool |  (optional) (default to false)
    includeSketches := true // bool |  (optional) (default to false)
    includeReferenceFeatures := true // bool |  (optional) (default to false)
    includeAssemblies := true // bool |  (optional) (default to false)
    includeFeatureStudios := true // bool |  (optional) (default to false)
    includeBlobs := true // bool |  (optional) (default to false)
    allowedBlobMimeTypes := "allowedBlobMimeTypes_example" // string |  (optional) (default to "")
    excludeNewerFSVersions := true // bool |  (optional) (default to false)
    maxFeatureScriptVersion := int32(56) // int32 |  (optional)
    includePartStudios := true // bool |  (optional) (default to false)
    includeFeatures := true // bool |  (optional) (default to false)
    includeMeshes := true // bool |  (optional) (default to false)
    includeWires := true // bool |  (optional) (default to false)
    includeFlattenedBodies := true // bool |  (optional) (default to false)
    includeApplications := true // bool |  (optional) (default to false)
    allowedApplicationMimeTypes := "allowedApplicationMimeTypes_example" // string |  (optional) (default to "")
    includeCompositeParts := true // bool |  (optional) (default to false)
    includeFSTables := true // bool |  (optional) (default to false)
    includeFSComputedPartPropertyFunctions := true // bool |  (optional) (default to false)
    includeVariables := true // bool |  (optional) (default to false)
    includeVariableStudios := true // bool |  (optional) (default to false)
    allowedBlobExtensions := "allowedBlobExtensions_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetInsertables(context.Background(), did, wv, wvid).ElementId(elementId).Configuration(configuration).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatureStudios(includeFeatureStudios).IncludeBlobs(includeBlobs).AllowedBlobMimeTypes(allowedBlobMimeTypes).ExcludeNewerFSVersions(excludeNewerFSVersions).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludePartStudios(includePartStudios).IncludeFeatures(includeFeatures).IncludeMeshes(includeMeshes).IncludeWires(includeWires).IncludeFlattenedBodies(includeFlattenedBodies).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).IncludeFSComputedPartPropertyFunctions(includeFSComputedPartPropertyFunctions).IncludeVariables(includeVariables).IncludeVariableStudios(includeVariableStudios).AllowedBlobExtensions(allowedBlobExtensions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetInsertables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInsertables`: BTInsertablesListResponse
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetInsertables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsertablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **elementId** | **string** |  | 
 **configuration** | **string** |  | 
 **betaCapabilityIds** | **[]string** |  | 
 **includeParts** | **bool** |  | [default to false]
 **includeSurfaces** | **bool** |  | [default to false]
 **includeSketches** | **bool** |  | [default to false]
 **includeReferenceFeatures** | **bool** |  | [default to false]
 **includeAssemblies** | **bool** |  | [default to false]
 **includeFeatureStudios** | **bool** |  | [default to false]
 **includeBlobs** | **bool** |  | [default to false]
 **allowedBlobMimeTypes** | **string** |  | [default to &quot;&quot;]
 **excludeNewerFSVersions** | **bool** |  | [default to false]
 **maxFeatureScriptVersion** | **int32** |  | 
 **includePartStudios** | **bool** |  | [default to false]
 **includeFeatures** | **bool** |  | [default to false]
 **includeMeshes** | **bool** |  | [default to false]
 **includeWires** | **bool** |  | [default to false]
 **includeFlattenedBodies** | **bool** |  | [default to false]
 **includeApplications** | **bool** |  | [default to false]
 **allowedApplicationMimeTypes** | **string** |  | [default to &quot;&quot;]
 **includeCompositeParts** | **bool** |  | [default to false]
 **includeFSTables** | **bool** |  | [default to false]
 **includeFSComputedPartPropertyFunctions** | **bool** |  | [default to false]
 **includeVariables** | **bool** |  | [default to false]
 **includeVariableStudios** | **bool** |  | [default to false]
 **allowedBlobExtensions** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTInsertablesListResponse**](BTInsertablesListResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnitInfo

> BTUnitInfo GetUnitInfo(ctx, did, wvm, wvmid).LinkDocumentId(linkDocumentId).Execute()

Get the selected units and precision by document ID and workspace or version or microversion ID.

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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetUnitInfo(context.Background(), did, wvm, wvmid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetUnitInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnitInfo`: BTUnitInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetUnitInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnitInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTUnitInfo**](BTUnitInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> BTVersionInfo GetVersion(ctx, did, vid).Parents(parents).LinkDocumentId(linkDocumentId).Execute()

Retrieve version by document ID and version ID.

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
    vid := "vid_example" // string | 
    parents := true // bool |  (optional) (default to false)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.GetVersion(context.Background(), did, vid).Parents(parents).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: BTVersionInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.GetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parents** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTVersionInfo**](BTVersionInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeIntoWorkspace

> BTDocumentMergeInfo MergeIntoWorkspace(ctx, did, wid).BTVersionOrWorkspaceMergeInfo(bTVersionOrWorkspaceMergeInfo).Execute()

Merge into workspace by document ID and workspace ID.

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
    bTVersionOrWorkspaceMergeInfo := *openapiclient.NewBTVersionOrWorkspaceMergeInfo() // BTVersionOrWorkspaceMergeInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.MergeIntoWorkspace(context.Background(), did, wid).BTVersionOrWorkspaceMergeInfo(bTVersionOrWorkspaceMergeInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.MergeIntoWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MergeIntoWorkspace`: BTDocumentMergeInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.MergeIntoWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeIntoWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTVersionOrWorkspaceMergeInfo** | [**BTVersionOrWorkspaceMergeInfo**](BTVersionOrWorkspaceMergeInfo.md) |  | 

### Return type

[**BTDocumentMergeInfo**](BTDocumentMergeInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergePreview

> BTMergePreviewInfo MergePreview(ctx, did, wid).SourceType(sourceType).SourceId(sourceId).LinkDocumentId(linkDocumentId).Execute()

Merge preview of changes that will occur based on document ID, workspace ID and source workspace/version ID

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
    sourceType := "sourceType_example" // string | 
    sourceId := "sourceId_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.MergePreview(context.Background(), did, wid).SourceType(sourceType).SourceId(sourceId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.MergePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MergePreview`: BTMergePreviewInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.MergePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceType** | **string** |  | 
 **sourceId** | **string** |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTMergePreviewInfo**](BTMergePreviewInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveElementsToDocument

> BTMoveElementInfo MoveElementsToDocument(ctx, did, wid).BTMoveElementParams(bTMoveElementParams).Execute()

Move tab by document ID and workspace ID.

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
    bTMoveElementParams := *openapiclient.NewBTMoveElementParams() // BTMoveElementParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.MoveElementsToDocument(context.Background(), did, wid).BTMoveElementParams(bTMoveElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.MoveElementsToDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveElementsToDocument`: BTMoveElementInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.MoveElementsToDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveElementsToDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTMoveElementParams** | [**BTMoveElementParams**](BTMoveElementParams.md) |  | 

### Return type

[**BTMoveElementInfo**](BTMoveElementInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFromHistory

> BTRestoreFromHistoryInfo RestoreFromHistory(ctx, did, wid, vm, vmid).LinkDocumentId(linkDocumentId).Execute()

Restore version or microversion to workspace by document ID, workspace ID, and version or microversion ID.

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
    vm := "vm_example" // string | 
    vmid := "vmid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.RestoreFromHistory(context.Background(), did, wid, vm, vmid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.RestoreFromHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreFromHistory`: BTRestoreFromHistoryInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.RestoreFromHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**vm** | **string** |  | 
**vmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFromHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTRestoreFromHistoryInfo**](BTRestoreFromHistoryInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertUnchangedToRevisions

> []BTUnchangedElementInfo RevertUnchangedToRevisions(ctx, did, wid).BTRevertUnchangedParams(bTRevertUnchangedParams).Execute()



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
    bTRevertUnchangedParams := *openapiclient.NewBTRevertUnchangedParams() // BTRevertUnchangedParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.RevertUnchangedToRevisions(context.Background(), did, wid).BTRevertUnchangedParams(bTRevertUnchangedParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.RevertUnchangedToRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertUnchangedToRevisions`: []BTUnchangedElementInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.RevertUnchangedToRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertUnchangedToRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTRevertUnchangedParams** | [**BTRevertUnchangedParams**](BTRevertUnchangedParams.md) |  | 

### Return type

[**[]BTUnchangedElementInfo**](BTUnchangedElementInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> map[string]interface{} Search(ctx).BTDocumentSearchParams(bTDocumentSearchParams).Execute()

Search document.

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
    bTDocumentSearchParams := *openapiclient.NewBTDocumentSearchParams() // BTDocumentSearchParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.Search(context.Background()).BTDocumentSearchParams(bTDocumentSearchParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTDocumentSearchParams** | [**BTDocumentSearchParams**](BTDocumentSearchParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareDocument

> BTAclInfo ShareDocument(ctx, did).BTShareParams(bTShareParams).Execute()

Share document by document ID.

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
    bTShareParams := *openapiclient.NewBTShareParams() // BTShareParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.ShareDocument(context.Background(), did).BTShareParams(bTShareParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.ShareDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShareDocument`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.ShareDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTShareParams** | [**BTShareParams**](BTShareParams.md) |  | 

### Return type

[**BTAclInfo**](BTAclInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncApplicationElements

> map[string]interface{} SyncApplicationElements(ctx, did, wid).BTSyncAppElementParams(bTSyncAppElementParams).Execute()



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
    bTSyncAppElementParams := *openapiclient.NewBTSyncAppElementParams() // BTSyncAppElementParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.SyncApplicationElements(context.Background(), did, wid).BTSyncAppElementParams(bTSyncAppElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.SyncApplicationElements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyncApplicationElements`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.SyncApplicationElements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncApplicationElementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTSyncAppElementParams** | [**BTSyncAppElementParams**](BTSyncAppElementParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnShareDocument

> map[string]interface{} UnShareDocument(ctx, did, eid).EntryType(entryType).Execute()

Unshare document by document ID and tab ID.

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
    entryType := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.UnShareDocument(context.Background(), did, eid).EntryType(entryType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.UnShareDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnShareDocument`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.UnShareDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnShareDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entryType** | **int32** |  | [default to 0]

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


## UpdateDocumentAttributes

> map[string]interface{} UpdateDocumentAttributes(ctx, did).BTDocumentParams(bTDocumentParams).Execute()

Update document attributes by document ID.

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
    bTDocumentParams := *openapiclient.NewBTDocumentParams() // BTDocumentParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.UpdateDocumentAttributes(context.Background(), did).BTDocumentParams(bTDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.UpdateDocumentAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocumentAttributes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.UpdateDocumentAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTDocumentParams** | [**BTDocumentParams**](BTDocumentParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalReferencesToLatestDocuments

> BTLinkToLatestDocumentInfo UpdateExternalReferencesToLatestDocuments(ctx, did, wid, eid).BTLinkToLatestDocumentParams(bTLinkToLatestDocumentParams).Execute()

Update external references to latest by document ID, workspace ID, and tab ID.

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
    bTLinkToLatestDocumentParams := *openapiclient.NewBTLinkToLatestDocumentParams() // BTLinkToLatestDocumentParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.UpdateExternalReferencesToLatestDocuments(context.Background(), did, wid, eid).BTLinkToLatestDocumentParams(bTLinkToLatestDocumentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.UpdateExternalReferencesToLatestDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExternalReferencesToLatestDocuments`: BTLinkToLatestDocumentInfo
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.UpdateExternalReferencesToLatestDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalReferencesToLatestDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTLinkToLatestDocumentParams** | [**BTLinkToLatestDocumentParams**](BTLinkToLatestDocumentParams.md) |  | 

### Return type

[**BTLinkToLatestDocumentInfo**](BTLinkToLatestDocumentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

