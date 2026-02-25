# \RevisionApi

All URIs are relative to *https://cad.onshape.com/api/v14*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRevisionHistory**](RevisionApi.md#DeleteRevisionHistory) | **Delete** /revisions/companies/{cid}/partnumber/{pnum}/elementType/{et} | Delete all revisions for a part number.
[**EnumerateRevisions**](RevisionApi.md#EnumerateRevisions) | **Get** /revisions/companies/{cid} | Get all revisions for a company.
[**GetAllInDocument**](RevisionApi.md#GetAllInDocument) | **Get** /revisions/d/{did} | Get all revisions for a document.
[**GetAllInDocumentVersion**](RevisionApi.md#GetAllInDocumentVersion) | **Get** /revisions/d/{did}/v/{vid} | Get all revisions for a version.
[**GetLatestInDocumentOrCompany**](RevisionApi.md#GetLatestInDocumentOrCompany) | **Get** /revisions/{cd}/{cdid}/p/{pnum}/latest | Get the latest revision information for a part.
[**GetRevisionByPartNumber**](RevisionApi.md#GetRevisionByPartNumber) | **Get** /revisions/c/{cid}/partnumber/{pnum} | Get details for the specified revision.
[**GetRevisionHistoryInCompanyByElementId**](RevisionApi.md#GetRevisionHistoryInCompanyByElementId) | **Get** /revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid} | Get all revisions for an element (tab).
[**GetRevisionHistoryInCompanyByPartId**](RevisionApi.md#GetRevisionHistoryInCompanyByPartId) | **Get** /revisions/companies/{cid}/d/{did}/{wv}/{wvid}/e/{eid}/p/{pid} | Get all revisions for a part ID.
[**GetRevisionHistoryInCompanyByPartNumber**](RevisionApi.md#GetRevisionHistoryInCompanyByPartNumber) | **Get** /revisions/companies/{cid}/partnumber/{pnum} | Get all revisions for a part number.



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
    cid := "cid_example" // string | The company or enterprise ID that owns the resource.
    pnum := "pnum_example" // string | Part number.
    et := "et_example" // string | Element Type. Must be one of: `0`: Part Studio, `1`: Assembly, `2`: Drawing. `4` : Blob
    ignoreLinkedDocuments := true // bool | By default, revisions will be deleted for the part number in the specified, and all linked documents. Set to `true` to only delete revisions in the specified document. (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**cid** | **string** | The company or enterprise ID that owns the resource. | 
**pnum** | **string** | Part number. | 
**et** | **string** | Element Type. Must be one of: &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRevisionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ignoreLinkedDocuments** | **bool** | By default, revisions will be deleted for the part number in the specified, and all linked documents. Set to &#x60;true&#x60; to only delete revisions in the specified document. | [default to false]

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

Get all revisions for a company.



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
    elementType := int32(56) // int32 | Element Type. Must be one of: `-1`: Unknown, `0`: Part Studio, `1`: Assembly, `2`: Drawing. `4` : Blob, `8`: Variable Studio (optional)
    limit := int32(56) // int32 | The number of list entries to return in a single API call. (optional) (default to 20)
    latestOnly := true // bool | Whether to limit search to only latest revisions. (optional) (default to false)
    after := time.Now() // JSONTime | The earliest creation date of the revision to find. (optional) (default to "2000-01-01T00:00Z")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.RevisionApi.EnumerateRevisions(context.Background(), cid).ElementType(elementType).Limit(limit).LatestOnly(latestOnly).After(after).Execute()
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
**cid** | **string** | The company or enterprise ID that owns the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnumerateRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementType** | **int32** | Element Type. Must be one of: &#x60;-1&#x60;: Unknown, &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob, &#x60;8&#x60;: Variable Studio | 
 **limit** | **int32** | The number of list entries to return in a single API call. | [default to 20]
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

Get all revisions for a document.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.RevisionApi.GetAllInDocument(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevisionApi.GetAllInDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllInDocument`: BTListResponseBTRevisionInfo
    fmt.Fprintf(os.Stdout, "Response from `RevisionApi.GetAllInDocument`: %v\n", resp)
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
    did := "did_example" // string | Document ID.
    vid := "vid_example" // string | Version ID.

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**did** | **string** | Document ID. | 
**vid** | **string** | Version ID. | 

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

Get the latest revision information for a part.



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
    cd := "cd_example" // string | Use `c` to specify a company ID or `d` to specify a document ID.
    cdid := "cdid_example" // string | Company ID or document ID
    pnum := "pnum_example" // string | Part number.
    et := "et_example" // string | Element Type. Must be one of: `-1`: Unknown, `0`: Part Studio, `1`: Assembly, `2`: Drawing. `4` : Blob, `8`: Variable Studio

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**cd** | **string** | Use &#x60;c&#x60; to specify a company ID or &#x60;d&#x60; to specify a document ID. | 
**cdid** | **string** | Company ID or document ID | 
**pnum** | **string** | Part number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInDocumentOrCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **et** | **string** | Element Type. Must be one of: &#x60;-1&#x60;: Unknown, &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob, &#x60;8&#x60;: Variable Studio | 

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

Get details for the specified revision.



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
    cid := "cid_example" // string | The company or enterprise ID that owns the resource.
    pnum := "pnum_example" // string | Part number.
    revision := "revision_example" // string | ID of the revision to get (optional)
    elementType := int32(56) // int32 | Element Type. Must be one of: `-1`: Unknown, `0`: Part Studio, `1`: Assembly, `2`: Drawing. `4` : Blob, `8`: Variable Studio (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**cid** | **string** | The company or enterprise ID that owns the resource. | 
**pnum** | **string** | Part number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionByPartNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revision** | **string** | ID of the revision to get | 
 **elementType** | **int32** | Element Type. Must be one of: &#x60;-1&#x60;: Unknown, &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob, &#x60;8&#x60;: Variable Studio | 

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

> BTRevisionListResponse GetRevisionHistoryInCompanyByElementId(ctx, cid, did, wv, wvid, eid).ElementType(elementType).LinkDocumentId(linkDocumentId).Configuration(configuration).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).SupportChangeType(supportChangeType).Execute()

Get all revisions for an element (tab).



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
    cid := "cid_example" // string | The company or enterprise ID that owns the resource.
    did := "did_example" // string | The id of the document in which to perform the operation.
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    elementType := "elementType_example" // string | Element Type. Must be one of: `-1`: Unknown, `0`: Part Studio, `1`: Assembly, `2`: Drawing. `4` : Blob, `8`: Variable Studio
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional) (default to "")
    fillApprovers := true // bool | Set to `true` to return a list of approvers. Default is `false` and will return `null`. (optional) (default to false)
    fillExportPermission := true // bool | Set to `true` to return a list of export permissions. Default is `false` and will return `null`. (optional) (default to false)
    supportChangeType := true // bool | Whether the revision can change object type. Used in reuse part number flow. (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.RevisionApi.GetRevisionHistoryInCompanyByElementId(context.Background(), cid, did, wv, wvid, eid).ElementType(elementType).LinkDocumentId(linkDocumentId).Configuration(configuration).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).SupportChangeType(supportChangeType).Execute()
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
**cid** | **string** | The company or enterprise ID that owns the resource. | 
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyByElementIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **elementType** | **string** | Element Type. Must be one of: &#x60;-1&#x60;: Unknown, &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob, &#x60;8&#x60;: Variable Studio | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **fillApprovers** | **bool** | Set to &#x60;true&#x60; to return a list of approvers. Default is &#x60;false&#x60; and will return &#x60;null&#x60;. | [default to false]
 **fillExportPermission** | **bool** | Set to &#x60;true&#x60; to return a list of export permissions. Default is &#x60;false&#x60; and will return &#x60;null&#x60;. | [default to false]
 **supportChangeType** | **bool** | Whether the revision can change object type. Used in reuse part number flow. | [default to false]

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

> BTRevisionListResponse GetRevisionHistoryInCompanyByPartId(ctx, cid, did, wv, wvid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).SupportChangeType(supportChangeType).Execute()

Get all revisions for a part ID.



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
    cid := "cid_example" // string | The company or enterprise ID that owns the resource.
    did := "did_example" // string | Document ID.
    wv := "wv_example" // string | One of w or v corresponding to whether a workspace or version was specified.
    wvid := "wvid_example" // string | Workspace (w) or Version (v) ID.
    eid := "eid_example" // string | Element ID.
    pid := "pid_example" // string | Part ID.
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)
    fillApprovers := true // bool | Set to `true` to return a list of approvers. Default is `false` and will return `null`. (optional) (default to false)
    fillExportPermission := true // bool | Set to `true` to return a list of export permissions. Default is `false` and will return `null`. (optional) (default to false)
    supportChangeType := true // bool | Whether the revision can change object type. Used in reuse part number flow. (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.RevisionApi.GetRevisionHistoryInCompanyByPartId(context.Background(), cid, did, wv, wvid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).SupportChangeType(supportChangeType).Execute()
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
**cid** | **string** | The company or enterprise ID that owns the resource. | 
**did** | **string** | Document ID. | 
**wv** | **string** | One of w or v corresponding to whether a workspace or version was specified. | 
**wvid** | **string** | Workspace (w) or Version (v) ID. | 
**eid** | **string** | Element ID. | 
**pid** | **string** | Part ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyByPartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 
 **fillApprovers** | **bool** | Set to &#x60;true&#x60; to return a list of approvers. Default is &#x60;false&#x60; and will return &#x60;null&#x60;. | [default to false]
 **fillExportPermission** | **bool** | Set to &#x60;true&#x60; to return a list of export permissions. Default is &#x60;false&#x60; and will return &#x60;null&#x60;. | [default to false]
 **supportChangeType** | **bool** | Whether the revision can change object type. Used in reuse part number flow. | [default to false]

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

> BTRevisionListResponse GetRevisionHistoryInCompanyByPartNumber(ctx, cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).SupportChangeType(supportChangeType).Execute()

Get all revisions for a part number.



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
    cid := "cid_example" // string | The company or enterprise ID that owns the resource.
    pnum := "pnum_example" // string | Part number.
    elementType := "elementType_example" // string | Element Type. Must be one of: `-1`: Unknown, `0`: Part Studio, `1`: Assembly, `2`: Drawing. `4` : Blob, `8`: Variable Studio
    fillApprovers := true // bool | Set to `true` to return a list of approvers. Default is `false` and will return `null`. (optional) (default to false)
    fillExportPermission := true // bool | Set to `true` to return a list of export permissions. Default is `false` and will return `null`. (optional) (default to false)
    supportChangeType := true // bool | Whether the revision can change object type. Used in reuse part number flow. (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.RevisionApi.GetRevisionHistoryInCompanyByPartNumber(context.Background(), cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).SupportChangeType(supportChangeType).Execute()
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
**cid** | **string** | The company or enterprise ID that owns the resource. | 
**pnum** | **string** | Part number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyByPartNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementType** | **string** | Element Type. Must be one of: &#x60;-1&#x60;: Unknown, &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob, &#x60;8&#x60;: Variable Studio | 
 **fillApprovers** | **bool** | Set to &#x60;true&#x60; to return a list of approvers. Default is &#x60;false&#x60; and will return &#x60;null&#x60;. | [default to false]
 **fillExportPermission** | **bool** | Set to &#x60;true&#x60; to return a list of export permissions. Default is &#x60;false&#x60; and will return &#x60;null&#x60;. | [default to false]
 **supportChangeType** | **bool** | Whether the revision can change object type. Used in reuse part number flow. | [default to false]

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

