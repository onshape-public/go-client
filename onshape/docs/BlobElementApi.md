# \BlobElementApi

All URIs are relative to *https://staging.dev.onshape.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlobTranslation**](BlobElementApi.md#CreateBlobTranslation) | **Post** /blobelements/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create translation (export) of blob element (document tab) by document id, workspace or version ID, and tab ID.
[**DownloadFileWorkspace**](BlobElementApi.md#DownloadFileWorkspace) | **Get** /blobelements/d/{did}/w/{wid}/e/{eid} | 
[**UpdateUnits**](BlobElementApi.md#UpdateUnits) | **Post** /blobelements/d/{did}/w/{wid}/e/{eid}/units | Update mesh units of a previously imported STL or OBJ file by document ID, workspace ID, and tab ID.
[**UploadFileCreateElement**](BlobElementApi.md#UploadFileCreateElement) | **Post** /blobelements/d/{did}/w/{wid} | Upload the file to a new tab by document ID and workspace ID.
[**UploadFileUpdateElement**](BlobElementApi.md#UploadFileUpdateElement) | **Post** /blobelements/d/{did}/w/{wid}/e/{eid} | Update a blob element by uploading a file by document ID, workspace ID, and tab ID.



## CreateBlobTranslation

> BTTranslationRequestInfo CreateBlobTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create translation (export) of blob element (document tab) by document id, workspace or version ID, and tab ID.

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
    eid := "eid_example" // string | 
    bTTranslateFormatParams := *openapiclient.NewBTTranslateFormatParams() // BTTranslateFormatParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobElementApi.CreateBlobTranslation(context.Background(), did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.CreateBlobTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlobTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.CreateBlobTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTTranslateFormatParams** | [**BTTranslateFormatParams**](BTTranslateFormatParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFileWorkspace

> HttpFile DownloadFileWorkspace(ctx, did, wid, eid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).LinkDocumentId(linkDocumentId).Execute()



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
    contentDisposition := "contentDisposition_example" // string |  (optional)
    ifNoneMatch := "ifNoneMatch_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobElementApi.DownloadFileWorkspace(context.Background(), did, wid, eid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.DownloadFileWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFileWorkspace`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.DownloadFileWorkspace`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiDownloadFileWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentDisposition** | **string** |  | 
 **ifNoneMatch** | **string** |  | 
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


## UpdateUnits

> BTDocumentElementProcessingInfo UpdateUnits(ctx, did, eid, wid).BTUpdateMeshUnitsParams(bTUpdateMeshUnitsParams).Execute()

Update mesh units of a previously imported STL or OBJ file by document ID, workspace ID, and tab ID.

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
    bTUpdateMeshUnitsParams := *openapiclient.NewBTUpdateMeshUnitsParams() // BTUpdateMeshUnitsParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlobElementApi.UpdateUnits(context.Background(), did, eid, wid).BTUpdateMeshUnitsParams(bTUpdateMeshUnitsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.UpdateUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUnits`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.UpdateUnits`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateMeshUnitsParams** | [**BTUpdateMeshUnitsParams**](BTUpdateMeshUnitsParams.md) |  | 

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileCreateElement

> BTDocumentElementProcessingInfo UploadFileCreateElement(ctx, did, wid).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).Execute()

Upload the file to a new tab by document ID and workspace ID.

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
    resp, r, err := apiClient.BlobElementApi.UploadFileCreateElement(context.Background(), did, wid).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.UploadFileCreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileCreateElement`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.UploadFileCreateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileCreateElementRequest struct via the builder pattern


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

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileUpdateElement

> BTDocumentElementProcessingInfo UploadFileUpdateElement(ctx, did, eid, wid).ParentChangeId(parentChangeId).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).Execute()

Update a blob element by uploading a file by document ID, workspace ID, and tab ID.

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
    parentChangeId := "parentChangeId_example" // string |  (optional)
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
    resp, r, err := apiClient.BlobElementApi.UploadFileUpdateElement(context.Background(), did, eid, wid).ParentChangeId(parentChangeId).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.UploadFileUpdateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileUpdateElement`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.UploadFileUpdateElement`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUploadFileUpdateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentChangeId** | **string** |  | 
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

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

