# BTPartMaterialInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LibraryName** | Pointer to **string** |  | [optional] 
**LibraryReference** | Pointer to [**BTExternalElementReferenceInfo**](BTExternalElementReferenceInfo.md) |  | [optional] 
**Properties** | Pointer to [**[]BTPartMaterialPropertyInfo**](BTPartMaterialPropertyInfo.md) |  | [optional] 

## Methods

### NewBTPartMaterialInfo

`func NewBTPartMaterialInfo() *BTPartMaterialInfo`

NewBTPartMaterialInfo instantiates a new BTPartMaterialInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPartMaterialInfoWithDefaults

`func NewBTPartMaterialInfoWithDefaults() *BTPartMaterialInfo`

NewBTPartMaterialInfoWithDefaults instantiates a new BTPartMaterialInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *BTPartMaterialInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BTPartMaterialInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BTPartMaterialInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BTPartMaterialInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *BTPartMaterialInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTPartMaterialInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTPartMaterialInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTPartMaterialInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibraryName

`func (o *BTPartMaterialInfo) GetLibraryName() string`

GetLibraryName returns the LibraryName field if non-nil, zero value otherwise.

### GetLibraryNameOk

`func (o *BTPartMaterialInfo) GetLibraryNameOk() (*string, bool)`

GetLibraryNameOk returns a tuple with the LibraryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryName

`func (o *BTPartMaterialInfo) SetLibraryName(v string)`

SetLibraryName sets LibraryName field to given value.

### HasLibraryName

`func (o *BTPartMaterialInfo) HasLibraryName() bool`

HasLibraryName returns a boolean if a field has been set.

### GetLibraryReference

`func (o *BTPartMaterialInfo) GetLibraryReference() BTExternalElementReferenceInfo`

GetLibraryReference returns the LibraryReference field if non-nil, zero value otherwise.

### GetLibraryReferenceOk

`func (o *BTPartMaterialInfo) GetLibraryReferenceOk() (*BTExternalElementReferenceInfo, bool)`

GetLibraryReferenceOk returns a tuple with the LibraryReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryReference

`func (o *BTPartMaterialInfo) SetLibraryReference(v BTExternalElementReferenceInfo)`

SetLibraryReference sets LibraryReference field to given value.

### HasLibraryReference

`func (o *BTPartMaterialInfo) HasLibraryReference() bool`

HasLibraryReference returns a boolean if a field has been set.

### GetProperties

`func (o *BTPartMaterialInfo) GetProperties() []BTPartMaterialPropertyInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTPartMaterialInfo) GetPropertiesOk() (*[]BTPartMaterialPropertyInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTPartMaterialInfo) SetProperties(v []BTPartMaterialPropertyInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTPartMaterialInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


