# BTCloudStorageAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudStorageAccountId** | Pointer to **string** |  | [optional] 
**CloudStorageProvider** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExportFolder** | Pointer to [**BTCloudStorageObjectInfo**](BTCloudStorageObjectInfo.md) |  | [optional] 
**ImportFolder** | Pointer to [**BTCloudStorageObjectInfo**](BTCloudStorageObjectInfo.md) |  | [optional] 
**LastAuthenticated** | Pointer to **JSONTime** |  | [optional] 

## Methods

### NewBTCloudStorageAccountInfo

`func NewBTCloudStorageAccountInfo() *BTCloudStorageAccountInfo`

NewBTCloudStorageAccountInfo instantiates a new BTCloudStorageAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCloudStorageAccountInfoWithDefaults

`func NewBTCloudStorageAccountInfoWithDefaults() *BTCloudStorageAccountInfo`

NewBTCloudStorageAccountInfoWithDefaults instantiates a new BTCloudStorageAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudStorageAccountId

`func (o *BTCloudStorageAccountInfo) GetCloudStorageAccountId() string`

GetCloudStorageAccountId returns the CloudStorageAccountId field if non-nil, zero value otherwise.

### GetCloudStorageAccountIdOk

`func (o *BTCloudStorageAccountInfo) GetCloudStorageAccountIdOk() (*string, bool)`

GetCloudStorageAccountIdOk returns a tuple with the CloudStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageAccountId

`func (o *BTCloudStorageAccountInfo) SetCloudStorageAccountId(v string)`

SetCloudStorageAccountId sets CloudStorageAccountId field to given value.

### HasCloudStorageAccountId

`func (o *BTCloudStorageAccountInfo) HasCloudStorageAccountId() bool`

HasCloudStorageAccountId returns a boolean if a field has been set.

### GetCloudStorageProvider

`func (o *BTCloudStorageAccountInfo) GetCloudStorageProvider() int32`

GetCloudStorageProvider returns the CloudStorageProvider field if non-nil, zero value otherwise.

### GetCloudStorageProviderOk

`func (o *BTCloudStorageAccountInfo) GetCloudStorageProviderOk() (*int32, bool)`

GetCloudStorageProviderOk returns a tuple with the CloudStorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageProvider

`func (o *BTCloudStorageAccountInfo) SetCloudStorageProvider(v int32)`

SetCloudStorageProvider sets CloudStorageProvider field to given value.

### HasCloudStorageProvider

`func (o *BTCloudStorageAccountInfo) HasCloudStorageProvider() bool`

HasCloudStorageProvider returns a boolean if a field has been set.

### GetEnabled

`func (o *BTCloudStorageAccountInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BTCloudStorageAccountInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BTCloudStorageAccountInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BTCloudStorageAccountInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExportFolder

`func (o *BTCloudStorageAccountInfo) GetExportFolder() BTCloudStorageObjectInfo`

GetExportFolder returns the ExportFolder field if non-nil, zero value otherwise.

### GetExportFolderOk

`func (o *BTCloudStorageAccountInfo) GetExportFolderOk() (*BTCloudStorageObjectInfo, bool)`

GetExportFolderOk returns a tuple with the ExportFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFolder

`func (o *BTCloudStorageAccountInfo) SetExportFolder(v BTCloudStorageObjectInfo)`

SetExportFolder sets ExportFolder field to given value.

### HasExportFolder

`func (o *BTCloudStorageAccountInfo) HasExportFolder() bool`

HasExportFolder returns a boolean if a field has been set.

### GetImportFolder

`func (o *BTCloudStorageAccountInfo) GetImportFolder() BTCloudStorageObjectInfo`

GetImportFolder returns the ImportFolder field if non-nil, zero value otherwise.

### GetImportFolderOk

`func (o *BTCloudStorageAccountInfo) GetImportFolderOk() (*BTCloudStorageObjectInfo, bool)`

GetImportFolderOk returns a tuple with the ImportFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFolder

`func (o *BTCloudStorageAccountInfo) SetImportFolder(v BTCloudStorageObjectInfo)`

SetImportFolder sets ImportFolder field to given value.

### HasImportFolder

`func (o *BTCloudStorageAccountInfo) HasImportFolder() bool`

HasImportFolder returns a boolean if a field has been set.

### GetLastAuthenticated

`func (o *BTCloudStorageAccountInfo) GetLastAuthenticated() JSONTime`

GetLastAuthenticated returns the LastAuthenticated field if non-nil, zero value otherwise.

### GetLastAuthenticatedOk

`func (o *BTCloudStorageAccountInfo) GetLastAuthenticatedOk() (*JSONTime, bool)`

GetLastAuthenticatedOk returns a tuple with the LastAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthenticated

`func (o *BTCloudStorageAccountInfo) SetLastAuthenticated(v JSONTime)`

SetLastAuthenticated sets LastAuthenticated field to given value.

### HasLastAuthenticated

`func (o *BTCloudStorageAccountInfo) HasLastAuthenticated() bool`

HasLastAuthenticated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


