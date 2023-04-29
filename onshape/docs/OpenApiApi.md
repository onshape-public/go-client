# \OpenApiApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenApi**](OpenApiApi.md#GetOpenApi) | **Get** /openapi | 
[**GetTags**](OpenApiApi.md#GetTags) | **Get** /openapi/tags | 



## GetOpenApi

> GetOpenApi(ctx).ForceReload(forceReload).Version(version).VersionAlias(versionAlias).NoFilter(noFilter).IncludedTags(includedTags).ExcludedTags(excludedTags).IncludeDeprecated(includeDeprecated).OnlyDeprecated(onlyDeprecated).DocumentationStatuses(documentationStatuses).RestUserRole(restUserRole).OperationIds(operationIds).ExcludedOperationIds(excludedOperationIds).Execute()



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
    forceReload := true // bool | Force reload the OpenApi definition. Only works when asking for the latest version. (optional)
    version := "version_example" // string | Specify a version of Onshape from which the OpenAPI is generated. If '*' is specified in any of the version fields, that indicates any version if acceptable. (optional)
    versionAlias := openapiclient.VersionAlias("LAST_MINOR") // VersionAlias | Version aliases based on the currently released version. (optional)
    noFilter := true // bool | Do not filter the specification at all. (optional) (default to false)
    includedTags := []string{"Inner_example"} // []string | Return only operations with tags included in includedTags. (optional)
    excludedTags := []string{"Inner_example"} // []string | If an operation contains an excluded tag, it is not returned from this endpoint. (optional)
    includeDeprecated := true // bool | Include deprecated endpoints. (optional) (default to false)
    onlyDeprecated := true // bool | Only include deprecated endpoints. (optional) (default to false)
    documentationStatuses := []openapiclient.Status{openapiclient.Status("DEVELOPMENT")} // []Status | Only return endpoints that have the specified documentation status. Default is to return all the endpoints the user should have access to. (optional)
    restUserRole := openapiclient.BTRestUserRole("PUBLIC") // BTRestUserRole | The REST user role for which this spec is requested. (optional)
    operationIds := []string{"Inner_example"} // []string | Only return operations with specified ids. (optional)
    excludedOperationIds := []string{"Inner_example"} // []string | Do not return operations with specified ids. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenApiApi.GetOpenApi(context.Background()).ForceReload(forceReload).Version(version).VersionAlias(versionAlias).NoFilter(noFilter).IncludedTags(includedTags).ExcludedTags(excludedTags).IncludeDeprecated(includeDeprecated).OnlyDeprecated(onlyDeprecated).DocumentationStatuses(documentationStatuses).RestUserRole(restUserRole).OperationIds(operationIds).ExcludedOperationIds(excludedOperationIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenApiApi.GetOpenApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forceReload** | **bool** | Force reload the OpenApi definition. Only works when asking for the latest version. | 
 **version** | **string** | Specify a version of Onshape from which the OpenAPI is generated. If &#39;*&#39; is specified in any of the version fields, that indicates any version if acceptable. | 
 **versionAlias** | [**VersionAlias**](VersionAlias.md) | Version aliases based on the currently released version. | 
 **noFilter** | **bool** | Do not filter the specification at all. | [default to false]
 **includedTags** | **[]string** | Return only operations with tags included in includedTags. | 
 **excludedTags** | **[]string** | If an operation contains an excluded tag, it is not returned from this endpoint. | 
 **includeDeprecated** | **bool** | Include deprecated endpoints. | [default to false]
 **onlyDeprecated** | **bool** | Only include deprecated endpoints. | [default to false]
 **documentationStatuses** | [**[]Status**](Status.md) | Only return endpoints that have the specified documentation status. Default is to return all the endpoints the user should have access to. | 
 **restUserRole** | [**BTRestUserRole**](BTRestUserRole.md) | The REST user role for which this spec is requested. | 
 **operationIds** | **[]string** | Only return operations with specified ids. | 
 **excludedOperationIds** | **[]string** | Do not return operations with specified ids. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> GetTags(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenApiApi.GetTags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenApiApi.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

