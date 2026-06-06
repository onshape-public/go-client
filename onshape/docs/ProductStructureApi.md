# \ProductStructureApi

All URIs are relative to *https://cad.onshape.com/api/v16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWhereUsed**](ProductStructureApi.md#GetWhereUsed) | **Get** /productstructure/whereused | Find where an item is used



## GetWhereUsed

> BTWhereUsedItemInfoList GetWhereUsed(ctx).DocumentId(documentId).ElementId(elementId).VersionId(versionId).Configuration(configuration).PartId(partId).PartNumber(partNumber).IncludeProperties(includeProperties).Filter(filter).IncludeVersionInfo(includeVersionInfo).UseLatestVersion(useLatestVersion).LimitToTypes(limitToTypes).Execute()

Find where an item is used



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
    documentId := "documentId_example" // string | The id of the document in which to perform the operation. (optional)
    elementId := "elementId_example" // string | The id of the element in which to perform the operation. (optional)
    versionId := "versionId_example" // string | The id of the version on which the operation should be performed. (optional)
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    partId := "partId_example" // string | Id of the part on which to perform the operation. (optional)
    partNumber := "partNumber_example" // string | Part number on which to perform the operation. (optional)
    includeProperties := true // bool | If true, include metadata properties for each result item. (optional) (default to false)
    filter := int32(56) // int32 | Result filter ordinal. `0: LATEST_REVISION, 1: ALL_REVISIONS, 2: ALL_VERSIONS (default), 3: ALL, 4: ALL_WORKSPACES (deprecated), 5: LATEST_VERSIONS, 6: MOST_RECENT_REVISION`. (optional)
    includeVersionInfo := true // bool | If true, include the list of document versions in which the queried item has been referenced by other documents. Each entry contains the version ID and name. (optional)
    useLatestVersion := true // bool | If true, automatically resolve to the most recent referenced document version for the queried item. This is only used when includeVersionInfo=true, documentId is provided, and versionId is not specified. (optional) (default to false)
    limitToTypes := "limitToTypes_example" // string | Comma-separated list of referring element type ordinals to limit results to. Valid values: 0 (PARTSTUDIO), 1 (ASSEMBLY), 2 (DRAWING). (optional) (default to "")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ProductStructureApi.GetWhereUsed(context.Background()).DocumentId(documentId).ElementId(elementId).VersionId(versionId).Configuration(configuration).PartId(partId).PartNumber(partNumber).IncludeProperties(includeProperties).Filter(filter).IncludeVersionInfo(includeVersionInfo).UseLatestVersion(useLatestVersion).LimitToTypes(limitToTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductStructureApi.GetWhereUsed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWhereUsed`: BTWhereUsedItemInfoList
    fmt.Fprintf(os.Stdout, "Response from `ProductStructureApi.GetWhereUsed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhereUsedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentId** | **string** | The id of the document in which to perform the operation. | 
 **elementId** | **string** | The id of the element in which to perform the operation. | 
 **versionId** | **string** | The id of the version on which the operation should be performed. | 
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **partId** | **string** | Id of the part on which to perform the operation. | 
 **partNumber** | **string** | Part number on which to perform the operation. | 
 **includeProperties** | **bool** | If true, include metadata properties for each result item. | [default to false]
 **filter** | **int32** | Result filter ordinal. &#x60;0: LATEST_REVISION, 1: ALL_REVISIONS, 2: ALL_VERSIONS (default), 3: ALL, 4: ALL_WORKSPACES (deprecated), 5: LATEST_VERSIONS, 6: MOST_RECENT_REVISION&#x60;. | 
 **includeVersionInfo** | **bool** | If true, include the list of document versions in which the queried item has been referenced by other documents. Each entry contains the version ID and name. | 
 **useLatestVersion** | **bool** | If true, automatically resolve to the most recent referenced document version for the queried item. This is only used when includeVersionInfo&#x3D;true, documentId is provided, and versionId is not specified. | [default to false]
 **limitToTypes** | **string** | Comma-separated list of referring element type ordinals to limit results to. Valid values: 0 (PARTSTUDIO), 1 (ASSEMBLY), 2 (DRAWING). | [default to &quot;&quot;]

### Return type

[**BTWhereUsedItemInfoList**](BTWhereUsedItemInfoList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

