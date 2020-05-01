# ContentDisposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**ModificationDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**ReadDate** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewContentDisposition

`func NewContentDisposition() *ContentDisposition`

NewContentDisposition instantiates a new ContentDisposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentDispositionWithDefaults

`func NewContentDispositionWithDefaults() *ContentDisposition`

NewContentDispositionWithDefaults instantiates a new ContentDisposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *ContentDisposition) GetCreationDate() JSONTime`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ContentDisposition) GetCreationDateOk() (*JSONTime, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ContentDisposition) SetCreationDate(v JSONTime)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ContentDisposition) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetFileName

`func (o *ContentDisposition) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ContentDisposition) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ContentDisposition) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ContentDisposition) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetModificationDate

`func (o *ContentDisposition) GetModificationDate() JSONTime`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *ContentDisposition) GetModificationDateOk() (*JSONTime, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *ContentDisposition) SetModificationDate(v JSONTime)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *ContentDisposition) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetParameters

`func (o *ContentDisposition) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ContentDisposition) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ContentDisposition) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ContentDisposition) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetReadDate

`func (o *ContentDisposition) GetReadDate() JSONTime`

GetReadDate returns the ReadDate field if non-nil, zero value otherwise.

### GetReadDateOk

`func (o *ContentDisposition) GetReadDateOk() (*JSONTime, bool)`

GetReadDateOk returns a tuple with the ReadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadDate

`func (o *ContentDisposition) SetReadDate(v JSONTime)`

SetReadDate sets ReadDate field to given value.

### HasReadDate

`func (o *ContentDisposition) HasReadDate() bool`

HasReadDate returns a boolean if a field has been set.

### GetSize

`func (o *ContentDisposition) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentDisposition) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentDisposition) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentDisposition) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ContentDisposition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentDisposition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentDisposition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentDisposition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


