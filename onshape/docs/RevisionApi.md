# \RevisionApi

All URIs are relative to *https://staging.dev.onshape.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRevisionHistory**](RevisionApi.md#DeleteRevisionHistory) | **Delete** /revisions/companies/{cid}/partnumber/{pnum}/elementType/{et} | 
[**EnumerateRevisions**](RevisionApi.md#EnumerateRevisions) | **Get** /revisions/companies/{cid} | Enumerate all revisions released in a company by company ID.
[**GetAllInDocumentVersion**](RevisionApi.md#GetAllInDocumentVersion) | **Get** /revisions/d/{did}/v/{vid} | Retrieve a list of all revisions that exist in a document version.
[**GetLatestInDocumentOrCompany**](RevisionApi.md#GetLatestInDocumentOrCompany) | **Get** /revisions/{cd}/{cdid}/p/{pnum}/latest | Retrieve latest revisions for a part number in a document or company by document ID, workspace or version or microversion ID, and tab ID.
[**GetRevisionByPartNumber**](RevisionApi.md#GetRevisionByPartNumber) | **Get** /revisions/c/{cid}/partnumber/{pnum} | Get Navigation URL
[**GetRevisionHistoryInCompanyByElementId**](RevisionApi.md#GetRevisionHistoryInCompanyByElementId) | **Get** /revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid} | Retrieve a list of all revisions for a part in a company by company ID, document ID, workspace or version ID, and tab ID.
[**GetRevisionHistoryInCompanyByPartId**](RevisionApi.md#GetRevisionHistoryInCompanyByPartId) | **Get** /revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid}/p/{pid} | Retrieve a list of all revisions for a part in a company by company ID, document ID, workspace or version ID, tab ID, and part ID.
[**GetRevisionHistoryInCompanyByPartNumber**](RevisionApi.md#GetRevisionHistoryInCompanyByPartNumber) | **Get** /revisions/companies/{cid}/partnumber/{pnum} | Retrieve a list of all revisions for a part number in a company by company ID.



## DeleteRevisionHistory

> map[string]interface{} DeleteRevisionHistory(ctx, cid, pnum, et).IgnoreLinkedDocuments(ignoreLinkedDocuments).Execute()



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
    resp, r, err := apiClient.RevisionApi.DeleteRevisionHistory(context.Background(), cid, pnum, et).IgnoreLinkedDocuments(ignoreLinkedDocuments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.DeleteRevisionHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRevisionHistory`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.DeleteRevisionHistory`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnumerateRevisions

> BTListResponseBTRevisionInfo EnumerateRevisions(ctx, cid).ElementType(elementType).Limit(limit).Offset(offset).LatestOnly(latestOnly).After(after).Execute()

Enumerate all revisions released in a company by company ID.

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
    cid := "cid_example" // string | 
    elementType := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)
    offset := int32(56) // int32 |  (optional) (default to 0)
    latestOnly := true // bool |  (optional) (default to false)
    after := time.Now() // JSONTime |  (optional) (default to "2000-01-01T00:00Z")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionApi.EnumerateRevisions(context.Background(), cid).ElementType(elementType).Limit(limit).Offset(offset).LatestOnly(latestOnly).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.EnumerateRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnumerateRevisions`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.EnumerateRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnumerateRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementType** | **int32** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **latestOnly** | **bool** |  | [default to false]
 **after** | **JSONTime** |  | [default to &quot;2000-01-01T00:00Z&quot;]

### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInDocumentVersion

> BTListResponseBTRevisionInfo GetAllInDocumentVersion(ctx, did, vid).Execute()

Retrieve a list of all revisions that exist in a document version.



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
    resp, r, err := apiClient.RevisionApi.GetAllInDocumentVersion(context.Background(), did, vid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.GetAllInDocumentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInDocumentVersion`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.GetAllInDocumentVersion`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestInDocumentOrCompany

> BTRevisionInfo GetLatestInDocumentOrCompany(ctx, cd, cdid, pnum).Et(et).Execute()

Retrieve latest revisions for a part number in a document or company by document ID, workspace or version or microversion ID, and tab ID.

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
    resp, r, err := apiClient.RevisionApi.GetLatestInDocumentOrCompany(context.Background(), cd, cdid, pnum).Et(et).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.GetLatestInDocumentOrCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestInDocumentOrCompany`: BTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.GetLatestInDocumentOrCompany`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionByPartNumber

> BTRevisionInfo GetRevisionByPartNumber(ctx, cid, pnum).Revision(revision).ElementType(elementType).Execute()

Get Navigation URL



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
    resp, r, err := apiClient.RevisionApi.GetRevisionByPartNumber(context.Background(), cid, pnum).Revision(revision).ElementType(elementType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.GetRevisionByPartNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionByPartNumber`: BTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.GetRevisionByPartNumber`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompanyByElementId

> BTRevisionListResponse GetRevisionHistoryInCompanyByElementId(ctx, cid, did, wv, wvid, eid).ElementType(elementType).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()

Retrieve a list of all revisions for a part in a company by company ID, document ID, workspace or version ID, and tab ID.

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
    elementType := "elementType_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    fillApprovers := true // bool |  (optional) (default to false)
    fillExportPermission := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevisionApi.GetRevisionHistoryInCompanyByElementId(context.Background(), cid, did, wv, wvid, eid).ElementType(elementType).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.GetRevisionHistoryInCompanyByElementId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionHistoryInCompanyByElementId`: BTRevisionListResponse
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.GetRevisionHistoryInCompanyByElementId`: %v\n", resp)
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

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyByElementIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **elementType** | **string** |  | 
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **fillApprovers** | **bool** |  | [default to false]
 **fillExportPermission** | **bool** |  | [default to false]

### Return type

[**BTRevisionListResponse**](BTRevisionListResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompanyByPartId

> BTRevisionListResponse GetRevisionHistoryInCompanyByPartId(ctx, cid, did, wv, wvid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()

Retrieve a list of all revisions for a part in a company by company ID, document ID, workspace or version ID, tab ID, and part ID.

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
    resp, r, err := apiClient.RevisionApi.GetRevisionHistoryInCompanyByPartId(context.Background(), cid, did, wv, wvid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.GetRevisionHistoryInCompanyByPartId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionHistoryInCompanyByPartId`: BTRevisionListResponse
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.GetRevisionHistoryInCompanyByPartId`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompanyByPartNumber

> BTRevisionListResponse GetRevisionHistoryInCompanyByPartNumber(ctx, cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()

Retrieve a list of all revisions for a part number in a company by company ID.

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
    resp, r, err := apiClient.RevisionApi.GetRevisionHistoryInCompanyByPartNumber(context.Background(), cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.GetRevisionHistoryInCompanyByPartNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRevisionHistoryInCompanyByPartNumber`: BTRevisionListResponse
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.GetRevisionHistoryInCompanyByPartNumber`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

