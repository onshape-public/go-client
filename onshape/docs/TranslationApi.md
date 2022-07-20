# \TranslationApi

All URIs are relative to *https://staging.dev.onshape.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranslation**](TranslationApi.md#CreateTranslation) | **Post** /translations/d/{did}/w/{wid} | Upload foreign data (for example, an X_T file) into Onshape, and then translate the data to generate a part, Part Studio, Assembly, or subassembly.
[**DeleteTranslation**](TranslationApi.md#DeleteTranslation) | **Delete** /translations/{tid} | Delete translation status entry.
[**GetAllTranslatorFormats**](TranslationApi.md#GetAllTranslatorFormats) | **Get** /translations/translationformats | Retrieve a list of translation formats that can work for this translation. Some are valid only as an input format and cannot be used as an output format.
[**GetDocumentTranslations**](TranslationApi.md#GetDocumentTranslations) | **Get** /translations/d/{did} | Request an array of translations that were made against this document.
[**GetTranslation**](TranslationApi.md#GetTranslation) | **Get** /translations/{tid} | Request information on an in-progress or completed translation.



## CreateTranslation

> BTTranslationRequestInfo CreateTranslation(ctx, did, wid).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).Execute()

Upload foreign data (for example, an X_T file) into Onshape, and then translate the data to generate a part, Part Studio, Assembly, or subassembly.

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
    file := map[string]interface{}{ ... } // map[string]interface{} | The file to upload. (optional)
    allowFaultyParts := true // bool |  (optional)
    createComposite := true // bool |  (optional)
    createDrawingIfPossible := true // bool |  (optional)
    encodedFilename := "encodedFilename_example" // string |  (optional)
    extractAssemblyHierarchy := true // bool |  (optional)
    flattenAssemblies := true // bool |  (optional)
    formatName := "formatName_example" // string |  (optional)
    joinAdjacentSurfaces := true // bool |  (optional)
    locationElementId := "locationElementId_example" // string |  (optional)
    locationGroupId := "locationGroupId_example" // string |  (optional)
    locationPosition := int32(56) // int32 |  (optional) (default to -1)
    notifyUser := true // bool |  (optional) (default to true)
    ownerId := "ownerId_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)
    projectId := "projectId_example" // string |  (optional)
    public := true // bool |  (optional)
    onePartPerDoc := true // bool |  (optional) (default to false)
    splitAssembliesIntoMultipleDocuments := true // bool |  (optional) (default to false)
    storeInDocument := true // bool |  (optional) (default to true)
    translate := true // bool |  (optional) (default to true)
    unit := "unit_example" // string |  (optional) (default to "")
    uploadId := "uploadId_example" // string |  (optional)
    versionString := "versionString_example" // string |  (optional)
    yAxisIsUp := true // bool |  (optional)
    importWithinDocument := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.CreateTranslation(context.Background(), did, wid).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.CreateTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.CreateTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | [**map[string]interface{}**](map[string]interface{}.md) | The file to upload. | 
 **allowFaultyParts** | **bool** |  | 
 **createComposite** | **bool** |  | 
 **createDrawingIfPossible** | **bool** |  | 
 **encodedFilename** | **string** |  | 
 **extractAssemblyHierarchy** | **bool** |  | 
 **flattenAssemblies** | **bool** |  | 
 **formatName** | **string** |  | 
 **joinAdjacentSurfaces** | **bool** |  | 
 **locationElementId** | **string** |  | 
 **locationGroupId** | **string** |  | 
 **locationPosition** | **int32** |  | [default to -1]
 **notifyUser** | **bool** |  | [default to true]
 **ownerId** | **string** |  | 
 **parentId** | **string** |  | 
 **projectId** | **string** |  | 
 **public** | **bool** |  | 
 **onePartPerDoc** | **bool** |  | [default to false]
 **splitAssembliesIntoMultipleDocuments** | **bool** |  | [default to false]
 **storeInDocument** | **bool** |  | [default to true]
 **translate** | **bool** |  | [default to true]
 **unit** | **string** |  | [default to &quot;&quot;]
 **uploadId** | **string** |  | 
 **versionString** | **string** |  | 
 **yAxisIsUp** | **bool** |  | 
 **importWithinDocument** | **bool** |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranslation

> map[string]interface{} DeleteTranslation(ctx, tid).Execute()

Delete translation status entry.

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
    tid := "tid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.DeleteTranslation(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.DeleteTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTranslation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.DeleteTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTranslationRequest struct via the builder pattern


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


## GetAllTranslatorFormats

> []BTModelFormatFullInfo GetAllTranslatorFormats(ctx).Execute()

Retrieve a list of translation formats that can work for this translation. Some are valid only as an input format and cannot be used as an output format.

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
    resp, r, err := apiClient.TranslationApi.GetAllTranslatorFormats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.GetAllTranslatorFormats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTranslatorFormats`: []BTModelFormatFullInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.GetAllTranslatorFormats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTranslatorFormatsRequest struct via the builder pattern


### Return type

[**[]BTModelFormatFullInfo**](BTModelFormatFullInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentTranslations

> BTListResponseBTTranslationRequestInfo GetDocumentTranslations(ctx, did).Offset(offset).Limit(limit).Execute()

Request an array of translations that were made against this document.

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
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.GetDocumentTranslations(context.Background(), did).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.GetDocumentTranslations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentTranslations`: BTListResponseBTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.GetDocumentTranslations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTTranslationRequestInfo**](BTListResponseBTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslation

> BTTranslationRequestInfo GetTranslation(ctx, tid).Execute()

Request information on an in-progress or completed translation.

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
    tid := "tid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslationApi.GetTranslation(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.GetTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.GetTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

