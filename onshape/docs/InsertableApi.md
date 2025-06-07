# \InsertableApi

All URIs are relative to *https://cad.onshape.com/api/v11*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestInDocument**](InsertableApi.md#GetLatestInDocument) | **Get** /insertables/d/{did}/latest | Get a list of things in this document that can be inserted elsewhere.



## GetLatestInDocument

> BTListResponseBTInsertableInfo GetLatestInDocument(ctx, did).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatureStudios(includeFeatureStudios).IncludeBlobs(includeBlobs).AllowedBlobMimeTypes(allowedBlobMimeTypes).ExcludeNewerFSVersions(excludeNewerFSVersions).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludePartStudios(includePartStudios).IncludeFeatures(includeFeatures).IncludeMeshes(includeMeshes).IncludeWires(includeWires).IncludeFlattenedBodies(includeFlattenedBodies).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).IncludeFSComputedPartPropertyFunctions(includeFSComputedPartPropertyFunctions).IncludeVariables(includeVariables).IncludeVariableStudios(includeVariableStudios).AllowedBlobExtensions(allowedBlobExtensions).IsObsoletion(isObsoletion).Execute()

Get a list of things in this document that can be inserted elsewhere.



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
    isObsoletion := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.InsertableApi.GetLatestInDocument(context.Background(), did).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatureStudios(includeFeatureStudios).IncludeBlobs(includeBlobs).AllowedBlobMimeTypes(allowedBlobMimeTypes).ExcludeNewerFSVersions(excludeNewerFSVersions).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludePartStudios(includePartStudios).IncludeFeatures(includeFeatures).IncludeMeshes(includeMeshes).IncludeWires(includeWires).IncludeFlattenedBodies(includeFlattenedBodies).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).IncludeFSComputedPartPropertyFunctions(includeFSComputedPartPropertyFunctions).IncludeVariables(includeVariables).IncludeVariableStudios(includeVariableStudios).AllowedBlobExtensions(allowedBlobExtensions).IsObsoletion(isObsoletion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsertableApi.GetLatestInDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestInDocument`: BTListResponseBTInsertableInfo
    fmt.Fprintf(os.Stdout, "Response from `InsertableApi.GetLatestInDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
 **isObsoletion** | **bool** |  | [default to false]

### Return type

[**BTListResponseBTInsertableInfo**](BTListResponseBTInsertableInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

