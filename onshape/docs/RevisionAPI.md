# \RevisionAPI

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRevisionHistory**](RevisionAPI.md#DeleteRevisionHistory) | **Delete** /revisions/companies/{cid}/partnumber/{pnum}/elementType/{et} | Delete all revisions for a part number.
[**EnumerateRevisions**](RevisionAPI.md#EnumerateRevisions) | **Get** /revisions/companies/{cid} | Enumerate all of a company&#39;s revisions.
[**GetAllInDocument**](RevisionAPI.md#GetAllInDocument) | **Get** /revisions/d/{did} | Get all revisions for the specified document.
[**GetAllInDocumentVersion**](RevisionAPI.md#GetAllInDocumentVersion) | **Get** /revisions/d/{did}/v/{vid} | Get all revisions for a version.
[**GetLatestInDocumentOrCompany**](RevisionAPI.md#GetLatestInDocumentOrCompany) | **Get** /revisions/{cd}/{cdid}/p/{pnum}/latest | Get the latest revision for a part number in a document or company.
[**GetRevisionByPartNumber**](RevisionAPI.md#GetRevisionByPartNumber) | **Get** /revisions/c/{cid}/partnumber/{pnum} | Get a list of revisions by part number.
[**GetRevisionHistoryInCompanyByElementId**](RevisionAPI.md#GetRevisionHistoryInCompanyByElementId) | **Get** /revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid} | Get a list of all revisions for an element in a company-owned document.
[**GetRevisionHistoryInCompanyByPartId**](RevisionAPI.md#GetRevisionHistoryInCompanyByPartId) | **Get** /revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid}/p/{pid} | Get a list of all revisions for a part in a company-owned document by part ID.
[**GetRevisionHistoryInCompanyByPartNumber**](RevisionAPI.md#GetRevisionHistoryInCompanyByPartNumber) | **Get** /revisions/companies/{cid}/partnumber/{pnum} | Get a list of all revisions for a part or element in a company-owned document by part number.



## DeleteRevisionHistory

> map[string]interface{} DeleteRevisionHistory(ctx, cid, pnum, et).IgnoreLinkedDocuments(ignoreLinkedDocuments).Execute()

Delete all revisions for a part number.



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
    cid := "cid_example" // string | 
    pnum := "pnum_example" // string | 
    et := "et_example" // string | 
    ignoreLinkedDocuments := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.DeleteRevisionHistory(context.Background(), cid, pnum, et).IgnoreLinkedDocuments(ignoreLinkedDocuments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.DeleteRevisionHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRevisionHistory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.DeleteRevisionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**pnum** | **string** |  | 
**et** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRevisionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ignoreLinkedDocuments** | **bool** |  | [default to false]

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


## EnumerateRevisions

> BTListResponseBTRevisionInfo EnumerateRevisions(ctx, cid).ElementType(elementType).Limit(limit).LatestOnly(latestOnly).After(after).Execute()

Enumerate all of a company's revisions.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    cid := "cid_example" // string | The company or enterprise ID that owns the resource.
    elementType := int32(56) // int32 | 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob (optional)
    limit := int32(56) // int32 | The number of items to return in a single API call (optional) (default to 20)
    latestOnly := true // bool | Whether to limit search to only latest revisions. (optional) (default to false)
    after := time.Now() // JSONTime | The earliest creation date of the revision to find. (optional) (default to "2000-01-01T00:00Z")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.EnumerateRevisions(context.Background(), cid).ElementType(elementType).Limit(limit).LatestOnly(latestOnly).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.EnumerateRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnumerateRevisions`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.EnumerateRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | The company or enterprise ID that owns the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnumerateRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementType** | **int32** | 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob | 
 **limit** | **int32** | The number of items to return in a single API call | [default to 20]
 **latestOnly** | **bool** | Whether to limit search to only latest revisions. | [default to false]
 **after** | **JSONTime** | The earliest creation date of the revision to find. | [default to &quot;2000-01-01T00:00Z&quot;]

### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInDocument

> BTListResponseBTRevisionInfo GetAllInDocument(ctx, did).Execute()

Get all revisions for the specified document.



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
    resp, r, err := apiClient.RevisionAPI.GetAllInDocument(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.GetAllInDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInDocument`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.GetAllInDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInDocumentVersion

> BTListResponseBTRevisionInfo GetAllInDocumentVersion(ctx, did, vid).Execute()

Get all revisions for a version.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.GetAllInDocumentVersion(context.Background(), did, vid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.GetAllInDocumentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInDocumentVersion`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.GetAllInDocumentVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInDocumentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestInDocumentOrCompany

> BTRevisionInfo GetLatestInDocumentOrCompany(ctx, cd, cdid, pnum).Et(et).Execute()

Get the latest revision for a part number in a document or company.



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
    cd := "cd_example" // string | 
    cdid := "cdid_example" // string | 
    pnum := "pnum_example" // string | 
    et := "et_example" // string | 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.GetLatestInDocumentOrCompany(context.Background(), cd, cdid, pnum).Et(et).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.GetLatestInDocumentOrCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestInDocumentOrCompany`: BTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.GetLatestInDocumentOrCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cd** | **string** |  | 
**cdid** | **string** |  | 
**pnum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInDocumentOrCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **et** | **string** | 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob | 

### Return type

[**BTRevisionInfo**](BTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionByPartNumber

> BTRevisionInfo GetRevisionByPartNumber(ctx, cid, pnum).Revision(revision).ElementType(elementType).Execute()

Get a list of revisions by part number.

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
    cid := "cid_example" // string | Company id
    pnum := "pnum_example" // string | Part Number
    revision := "revision_example" // string | Revision (optional)
    elementType := int32(56) // int32 | 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.GetRevisionByPartNumber(context.Background(), cid, pnum).Revision(revision).ElementType(elementType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.GetRevisionByPartNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionByPartNumber`: BTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.GetRevisionByPartNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | Company id | 
**pnum** | **string** | Part Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionByPartNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **string** | Revision | 
 **elementType** | **int32** | 0: Part Studio, 1: Assembly, 2: Drawing. 4: Blob | 

### Return type

[**BTRevisionInfo**](BTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompanyByElementId

> BTRevisionListResponse GetRevisionHistoryInCompanyByElementId(ctx, cid, did, wv, wvid, eid).ElementType(elementType).LinkDocumentId(linkDocumentId).Configuration(configuration).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()

Get a list of all revisions for an element in a company-owned document.

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
    cid := "cid_example" // string | 
    did := "did_example" // string | The id of the document in which to perform the operation.
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    elementType := "elementType_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`) for the feature definition. (optional) (default to "")
    fillApprovers := true // bool |  (optional) (default to false)
    fillExportPermission := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.GetRevisionHistoryInCompanyByElementId(context.Background(), cid, did, wv, wvid, eid).ElementType(elementType).LinkDocumentId(linkDocumentId).Configuration(configuration).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.GetRevisionHistoryInCompanyByElementId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionHistoryInCompanyByElementId`: BTRevisionListResponse
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.GetRevisionHistoryInCompanyByElementId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyByElementIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **elementType** | **string** |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;) for the feature definition. | [default to &quot;&quot;]
 **fillApprovers** | **bool** |  | [default to false]
 **fillExportPermission** | **bool** |  | [default to false]

### Return type

[**BTRevisionListResponse**](BTRevisionListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompanyByPartId

> BTRevisionListResponse GetRevisionHistoryInCompanyByPartId(ctx, cid, did, wv, wvid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()

Get a list of all revisions for a part in a company-owned document by part ID.

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
    cid := "cid_example" // string | 
    did := "did_example" // string | 
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | 
    pid := "pid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    fillApprovers := true // bool |  (optional) (default to false)
    fillExportPermission := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.GetRevisionHistoryInCompanyByPartId(context.Background(), cid, did, wv, wvid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.GetRevisionHistoryInCompanyByPartId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionHistoryInCompanyByPartId`: BTRevisionListResponse
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.GetRevisionHistoryInCompanyByPartId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyByPartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **fillApprovers** | **bool** |  | [default to false]
 **fillExportPermission** | **bool** |  | [default to false]

### Return type

[**BTRevisionListResponse**](BTRevisionListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompanyByPartNumber

> BTRevisionListResponse GetRevisionHistoryInCompanyByPartNumber(ctx, cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()

Get a list of all revisions for a part or element in a company-owned document by part number.



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
    cid := "cid_example" // string | 
    pnum := "pnum_example" // string | 
    elementType := "elementType_example" // string | 
    fillApprovers := true // bool |  (optional) (default to false)
    fillExportPermission := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionAPI.GetRevisionHistoryInCompanyByPartNumber(context.Background(), cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionAPI.GetRevisionHistoryInCompanyByPartNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionHistoryInCompanyByPartNumber`: BTRevisionListResponse
    fmt.Fprintf(os.Stdout, "Response from `RevisionAPI.GetRevisionHistoryInCompanyByPartNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**pnum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyByPartNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementType** | **string** |  | 
 **fillApprovers** | **bool** |  | [default to false]
 **fillExportPermission** | **bool** |  | [default to false]

### Return type

[**BTRevisionListResponse**](BTRevisionListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

