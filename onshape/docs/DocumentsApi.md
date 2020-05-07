# \DocumentsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyWorkspace**](DocumentsApi.md#CopyWorkspace) | **Post** /api/documents/{did}/workspaces/{wid}/copy | Copy Workspace
[**CreateDocument**](DocumentsApi.md#CreateDocument) | **Post** /api/documents | Create document.
[**CreateVersion**](DocumentsApi.md#CreateVersion) | **Post** /api/documents/d/{did}/versions | Create Version.
[**CreateWorkspace**](DocumentsApi.md#CreateWorkspace) | **Post** /api/documents/d/{did}/workspaces | Create Workspace
[**DeleteDocument**](DocumentsApi.md#DeleteDocument) | **Delete** /api/documents/{did} | Delete Document
[**DeleteWorkspace**](DocumentsApi.md#DeleteWorkspace) | **Delete** /api/documents/d/{did}/workspaces/{wid} | Delete Workspace
[**DownloadExternalData**](DocumentsApi.md#DownloadExternalData) | **Get** /api/documents/d/{did}/externaldata/{fid} | Download External Data
[**Export2Json**](DocumentsApi.md#Export2Json) | **Post** /api/documents/d/{did}/{wv}/{wvid}/e/{eid}/export | 
[**GetCurrentMicroversion**](DocumentsApi.md#GetCurrentMicroversion) | **Get** /api/documents/d/{did}/{wv}/{wvid}/currentmicroversion | Get Current Document Microversion
[**GetDocument**](DocumentsApi.md#GetDocument) | **Get** /api/documents/{did} | Get Document
[**GetDocumentAcl**](DocumentsApi.md#GetDocumentAcl) | **Get** /api/documents/{did}/acl | Get Access Control List
[**GetDocumentPermissionSet**](DocumentsApi.md#GetDocumentPermissionSet) | **Get** /api/documents/{did}/permissionset | Get Document Permissions
[**GetDocumentVersions**](DocumentsApi.md#GetDocumentVersions) | **Get** /api/documents/d/{did}/versions | Get Versions
[**GetDocumentWorkspaces**](DocumentsApi.md#GetDocumentWorkspaces) | **Get** /api/documents/d/{did}/workspaces | Get Workspaces
[**GetDocuments**](DocumentsApi.md#GetDocuments) | **Get** /api/documents | Get Documents
[**GetElementsInDocument**](DocumentsApi.md#GetElementsInDocument) | **Get** /api/documents/d/{did}/{wvm}/{wvmid}/elements | Get a list of elements in the workspace, version, or microversion of the document.
[**GetInsertables**](DocumentsApi.md#GetInsertables) | **Get** /api/documents/d/{did}/{wvm}/{wvmid}/insertables | Insertable List for Document Version.
[**GetVersion**](DocumentsApi.md#GetVersion) | **Get** /api/documents/d/{did}/versions/{vid} | Get Version
[**MergeIntoWorkspace**](DocumentsApi.md#MergeIntoWorkspace) | **Post** /api/documents/{did}/workspaces/{wid}/merge | Merge into workspace
[**MoveElementsToDocument**](DocumentsApi.md#MoveElementsToDocument) | **Post** /api/documents/d/{did}/w/{wid}/moveelement | Move Elements
[**RestoreFromHistory**](DocumentsApi.md#RestoreFromHistory) | **Post** /api/documents/{did}/w/{wid}/restore/{vm}/{vmid} | Restore version or microversion to workspace.
[**Search**](DocumentsApi.md#Search) | **Post** /api/documents/search | 
[**ShareDocument**](DocumentsApi.md#ShareDocument) | **Post** /api/documents/{did}/share | Share Document
[**SyncApplicationElements**](DocumentsApi.md#SyncApplicationElements) | **Post** /api/documents/d/{did}/w/{wid}/syncApplicationElements | Sync Application Elements
[**UnShareDocument**](DocumentsApi.md#UnShareDocument) | **Delete** /api/documents/{did}/share/{eid} | Unshare Document
[**UpdateDocumentAttributes**](DocumentsApi.md#UpdateDocumentAttributes) | **Post** /api/documents/{did} | Update Document Attributes.
[**UpdateExternalReferencesToLatestDocuments**](DocumentsApi.md#UpdateExternalReferencesToLatestDocuments) | **Post** /api/documents/d/{did}/w/{wid}/e/{eid}/latestdocumentreferences | Update External References to Latest



## CopyWorkspace

> BTCopyDocumentInfo CopyWorkspace(ctx, did, wid).BTCopyDocumentParams(bTCopyDocumentParams).Execute()

Copy Workspace

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocument

> BTDocumentInfo CreateDocument(ctx).BTDocumentParams(bTDocumentParams).Execute()

Create document.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTDocumentParams** | [**BTDocumentParams**](BTDocumentParams.md) |  | 

### Return type

[**BTDocumentInfo**](BTDocumentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVersion

> BTVersionInfo CreateVersion(ctx, did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()

Create Version.

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> BTWorkspaceInfo CreateWorkspace(ctx, did).BTVersionOrWorkspaceParams(bTVersionOrWorkspaceParams).Execute()

Create Workspace

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, did).Forever(forever).Execute()

Delete Document

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

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspace

> DeleteWorkspace(ctx, did, wid).Execute()

Delete Workspace

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

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadExternalData

> *os.File DownloadExternalData(ctx, did, fid).IfNoneMatch(ifNoneMatch).Execute()

Download External Data

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

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+octet-stream;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Export2Json

> Export2Json(ctx, did, wv, wvid, eid).BTExportModelParams(bTExportModelParams).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExport2JsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTExportModelParams** | [**BTExportModelParams**](BTExportModelParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+octet-stream;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentMicroversion

> BTMicroversionInfo GetCurrentMicroversion(ctx, did, wv, wvid).Execute()

Get Current Document Microversion

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> BTDocumentInfo GetDocument(ctx, did).Execute()

Get Document

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentAcl

> BTAclInfo GetDocumentAcl(ctx, did).Execute()

Get Access Control List

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentPermissionSet

> []string GetDocumentPermissionSet(ctx, did).Execute()

Get Document Permissions

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentVersions

> []BTVersionInfo GetDocumentVersions(ctx, did).Offset(offset).Limit(limit).Execute()

Get Versions

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentWorkspaces

> []BTWorkspaceInfo GetDocumentWorkspaces(ctx, did).Execute()

Get Workspaces

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocuments

> BTGlobalTreeNodeListResponse GetDocuments(ctx).Q(q).Filter(filter).Owner(owner).OwnerType(ownerType).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Label(label).Project(project).ParentId(parentId).Execute()

Get Documents

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementsInDocument

> []BTDocumentElementInfo GetElementsInDocument(ctx, did, wvm, wvmid).ElementType(elementType).ElementId(elementId).WithThumbnails(withThumbnails).LinkDocumentId(linkDocumentId).Execute()

Get a list of elements in the workspace, version, or microversion of the document.

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInsertables

> BTInsertablesListResponse GetInsertables(ctx, did, wvm, wvmid).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeWires(includeWires).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatures(includeFeatures).IncludeFeatureStudios(includeFeatureStudios).IncludePartStudios(includePartStudios).IncludeBlobs(includeBlobs).IncludeMeshes(includeMeshes).IncludeFlattenedBodies(includeFlattenedBodies).AllowedBlobMimeTypes(allowedBlobMimeTypes).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).Execute()

Insertable List for Document Version.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInsertablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **betaCapabilityIds** | [**[]string**](string.md) |  | 
 **includeParts** | **bool** |  | [default to false]
 **includeSurfaces** | **bool** |  | [default to false]
 **includeWires** | **bool** |  | [default to false]
 **includeSketches** | **bool** |  | [default to false]
 **includeReferenceFeatures** | **bool** |  | [default to false]
 **includeAssemblies** | **bool** |  | [default to false]
 **includeFeatures** | **bool** |  | [default to false]
 **includeFeatureStudios** | **bool** |  | [default to false]
 **includePartStudios** | **bool** |  | [default to false]
 **includeBlobs** | **bool** |  | [default to false]
 **includeMeshes** | **bool** |  | [default to false]
 **includeFlattenedBodies** | **bool** |  | [default to false]
 **allowedBlobMimeTypes** | **string** |  | [default to &quot;&quot;]
 **maxFeatureScriptVersion** | **int32** |  | [default to 0]
 **includeApplications** | **bool** |  | [default to false]
 **allowedApplicationMimeTypes** | **string** |  | [default to &quot;&quot;]
 **includeCompositeParts** | **bool** |  | [default to false]
 **includeFSTables** | **bool** |  | [default to false]

### Return type

[**BTInsertablesListResponse**](BTInsertablesListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> BTVersionInfo GetVersion(ctx, did, vid).Parents(parents).LinkDocumentId(linkDocumentId).Execute()

Get Version

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeIntoWorkspace

> BTDocumentMergeInfo MergeIntoWorkspace(ctx, did, wid).BTVersionOrWorkspaceInfo(bTVersionOrWorkspaceInfo).Execute()

Merge into workspace

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


 **bTVersionOrWorkspaceInfo** | [**BTVersionOrWorkspaceInfo**](BTVersionOrWorkspaceInfo.md) |  | 

### Return type

[**BTDocumentMergeInfo**](BTDocumentMergeInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveElementsToDocument

> BTMoveElementInfo MoveElementsToDocument(ctx, did, wid).BTMoveElementParams(bTMoveElementParams).Execute()

Move Elements

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFromHistory

> RestoreFromHistory(ctx, did, wid, vm, vmid).Execute()

Restore version or microversion to workspace.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**vm** | **string** |  | 
**vmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFromHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> Search(ctx).BTDocumentSearchParams(bTDocumentSearchParams).Execute()



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTDocumentSearchParams** | [**BTDocumentSearchParams**](BTDocumentSearchParams.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareDocument

> BTAclInfo ShareDocument(ctx, did).BTShareParams(bTShareParams).Execute()

Share Document

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncApplicationElements

> SyncApplicationElements(ctx, did, wid).ApplicationElementIds(applicationElementIds).Description(description).Execute()

Sync Application Elements

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


 **applicationElementIds** | [**[]string**](string.md) |  | 
 **description** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnShareDocument

> UnShareDocument(ctx, did, eid).EntryType(entryType).Execute()

Unshare Document

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

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentAttributes

> UpdateDocumentAttributes(ctx, did).BTDocumentParams(bTDocumentParams).Execute()

Update Document Attributes.

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

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExternalReferencesToLatestDocuments

> BTLinkToLatestDocumentInfo UpdateExternalReferencesToLatestDocuments(ctx, did, wid, eid).BTLinkToLatestDocumentParams(bTLinkToLatestDocumentParams).Execute()

Update External References to Latest

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

