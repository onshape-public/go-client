# FormDataContentDisposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**ModificationDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**ReadDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewFormDataContentDisposition

`func NewFormDataContentDisposition() *FormDataContentDisposition`

NewFormDataContentDisposition instantiates a new FormDataContentDisposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDataContentDispositionWithDefaults

`func NewFormDataContentDispositionWithDefaults() *FormDataContentDisposition`

NewFormDataContentDispositionWithDefaults instantiates a new FormDataContentDisposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *FormDataContentDisposition) GetCreationDate() JSONTime`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *FormDataContentDisposition) GetCreationDateOk() (*JSONTime, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *FormDataContentDisposition) SetCreationDate(v JSONTime)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *FormDataContentDisposition) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetFileName

`func (o *FormDataContentDisposition) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FormDataContentDisposition) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FormDataContentDisposition) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FormDataContentDisposition) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetModificationDate

`func (o *FormDataContentDisposition) GetModificationDate() JSONTime`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *FormDataContentDisposition) GetModificationDateOk() (*JSONTime, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *FormDataContentDisposition) SetModificationDate(v JSONTime)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *FormDataContentDisposition) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetName

`func (o *FormDataContentDisposition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormDataContentDisposition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormDataContentDisposition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormDataContentDisposition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *FormDataContentDisposition) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FormDataContentDisposition) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FormDataContentDisposition) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *FormDataContentDisposition) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReadDate

`func (o *FormDataContentDisposition) GetReadDate() JSONTime`

GetReadDate returns the ReadDate field if non-nil, zero value otherwise.

### GetReadDateOk

`func (o *FormDataContentDisposition) GetReadDateOk() (*JSONTime, bool)`

GetReadDateOk returns a tuple with the ReadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadDate

`func (o *FormDataContentDisposition) SetReadDate(v JSONTime)`

SetReadDate sets ReadDate field to given value.

### HasReadDate

`func (o *FormDataContentDisposition) HasReadDate() bool`

HasReadDate returns a boolean if a field has been set.

### GetSize

`func (o *FormDataContentDisposition) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FormDataContentDisposition) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FormDataContentDisposition) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FormDataContentDisposition) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *FormDataContentDisposition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormDataContentDisposition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormDataContentDisposition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormDataContentDisposition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


