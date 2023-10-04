# BTShareParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**EncodedConfiguration** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]BTShareEntryParams**](BTShareEntryParams.md) |  | [optional] 
**FolderId** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Permission** | Pointer to **int64** |  | [optional] 
**PermissionSet** | Pointer to **[]string** |  | [optional] 
**Update** | Pointer to **bool** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTShareParams

`func NewBTShareParams() *BTShareParams`

NewBTShareParams instantiates a new BTShareParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTShareParamsWithDefaults

`func NewBTShareParamsWithDefaults() *BTShareParams`

NewBTShareParamsWithDefaults instantiates a new BTShareParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *BTShareParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTShareParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTShareParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTShareParams) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTShareParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTShareParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTShareParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTShareParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetEncodedConfiguration

`func (o *BTShareParams) GetEncodedConfiguration() string`

GetEncodedConfiguration returns the EncodedConfiguration field if non-nil, zero value otherwise.

### GetEncodedConfigurationOk

`func (o *BTShareParams) GetEncodedConfigurationOk() (*string, bool)`

GetEncodedConfigurationOk returns a tuple with the EncodedConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedConfiguration

`func (o *BTShareParams) SetEncodedConfiguration(v string)`

SetEncodedConfiguration sets EncodedConfiguration field to given value.

### HasEncodedConfiguration

`func (o *BTShareParams) HasEncodedConfiguration() bool`

HasEncodedConfiguration returns a boolean if a field has been set.

### GetEntries

`func (o *BTShareParams) GetEntries() []BTShareEntryParams`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BTShareParams) GetEntriesOk() (*[]BTShareEntryParams, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BTShareParams) SetEntries(v []BTShareEntryParams)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BTShareParams) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetFolderId

`func (o *BTShareParams) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *BTShareParams) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *BTShareParams) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *BTShareParams) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetMessage

`func (o *BTShareParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BTShareParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BTShareParams) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BTShareParams) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPermission

`func (o *BTShareParams) GetPermission() int64`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BTShareParams) GetPermissionOk() (*int64, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BTShareParams) SetPermission(v int64)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *BTShareParams) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetPermissionSet

`func (o *BTShareParams) GetPermissionSet() []string`

GetPermissionSet returns the PermissionSet field if non-nil, zero value otherwise.

### GetPermissionSetOk

`func (o *BTShareParams) GetPermissionSetOk() (*[]string, bool)`

GetPermissionSetOk returns a tuple with the PermissionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSet

`func (o *BTShareParams) SetPermissionSet(v []string)`

SetPermissionSet sets PermissionSet field to given value.

### HasPermissionSet

`func (o *BTShareParams) HasPermissionSet() bool`

HasPermissionSet returns a boolean if a field has been set.

### GetUpdate

`func (o *BTShareParams) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *BTShareParams) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *BTShareParams) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *BTShareParams) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTShareParams) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTShareParams) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTShareParams) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTShareParams) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


