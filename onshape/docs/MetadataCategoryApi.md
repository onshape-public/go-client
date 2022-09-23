# \MetadataCategoryApi

All URIs are relative to *https://staging.dev.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCategoryProperties**](MetadataCategoryApi.md#GetCategoryProperties) | **Get** /metadatacategory/categoryproperties | Retrieve category properties for metadata.



## GetCategoryProperties

> BTListResponseBTCategoryPropertyInfo GetCategoryProperties(ctx).OwnerId(ownerId).OwnerType(ownerType).DocumentId(documentId).CategoryIds(categoryIds).ObjectType(objectType).Strict(strict).IncludeObjectTypeDefaults(includeObjectTypeDefaults).IncludeComputedProperties(includeComputedProperties).IncludePartPropertiesTableOnlyProperties(includePartPropertiesTableOnlyProperties).OnlyActive(onlyActive).OnlyObjectTypeDefaults(onlyObjectTypeDefaults).Execute()

Retrieve category properties for metadata.

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
    ownerId := "ownerId_example" // string |  (optional)
    ownerType := int32(56) // int32 |  (optional) (default to 1)
    documentId := "documentId_example" // string |  (optional)
    categoryIds := []string{"Inner_example"} // []string |  (optional)
    objectType := int32(56) // int32 |  (optional)
    strict := true // bool |  (optional) (default to true)
    includeObjectTypeDefaults := true // bool |  (optional) (default to false)
    includeComputedProperties := true // bool |  (optional) (default to true)
    includePartPropertiesTableOnlyProperties := true // bool |  (optional) (default to true)
    onlyActive := true // bool |  (optional) (default to false)
    onlyObjectTypeDefaults := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataCategoryApi.GetCategoryProperties(context.Background()).OwnerId(ownerId).OwnerType(ownerType).DocumentId(documentId).CategoryIds(categoryIds).ObjectType(objectType).Strict(strict).IncludeObjectTypeDefaults(includeObjectTypeDefaults).IncludeComputedProperties(includeComputedProperties).IncludePartPropertiesTableOnlyProperties(includePartPropertiesTableOnlyProperties).OnlyActive(onlyActive).OnlyObjectTypeDefaults(onlyObjectTypeDefaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataCategoryApi.GetCategoryProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryProperties`: BTListResponseBTCategoryPropertyInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataCategoryApi.GetCategoryProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** |  | 
 **ownerType** | **int32** |  | [default to 1]
 **documentId** | **string** |  | 
 **categoryIds** | **[]string** |  | 
 **objectType** | **int32** |  | 
 **strict** | **bool** |  | [default to true]
 **includeObjectTypeDefaults** | **bool** |  | [default to false]
 **includeComputedProperties** | **bool** |  | [default to true]
 **includePartPropertiesTableOnlyProperties** | **bool** |  | [default to true]
 **onlyActive** | **bool** |  | [default to false]
 **onlyObjectTypeDefaults** | **bool** |  | [default to false]

### Return type

[**BTListResponseBTCategoryPropertyInfo**](BTListResponseBTCategoryPropertyInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

