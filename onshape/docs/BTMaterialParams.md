# BTMaterialParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LibraryName** | Pointer to **string** |  | [optional] 
**LibraryReference** | Pointer to [**BTExternalElementReferenceInfo**](BTExternalElementReferenceInfo.md) |  | [optional] 
**Properties** | Pointer to [**[]BTMaterialPropertyParams**](BTMaterialPropertyParams.md) |  | [optional] 

## Methods

### NewBTMaterialParams

`func NewBTMaterialParams() *BTMaterialParams`

NewBTMaterialParams instantiates a new BTMaterialParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMaterialParamsWithDefaults

`func NewBTMaterialParamsWithDefaults() *BTMaterialParams`

NewBTMaterialParamsWithDefaults instantiates a new BTMaterialParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *BTMaterialParams) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BTMaterialParams) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BTMaterialParams) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BTMaterialParams) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *BTMaterialParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTMaterialParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTMaterialParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTMaterialParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibraryName

`func (o *BTMaterialParams) GetLibraryName() string`

GetLibraryName returns the LibraryName field if non-nil, zero value otherwise.

### GetLibraryNameOk

`func (o *BTMaterialParams) GetLibraryNameOk() (*string, bool)`

GetLibraryNameOk returns a tuple with the LibraryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryName

`func (o *BTMaterialParams) SetLibraryName(v string)`

SetLibraryName sets LibraryName field to given value.

### HasLibraryName

`func (o *BTMaterialParams) HasLibraryName() bool`

HasLibraryName returns a boolean if a field has been set.

### GetLibraryReference

`func (o *BTMaterialParams) GetLibraryReference() BTExternalElementReferenceInfo`

GetLibraryReference returns the LibraryReference field if non-nil, zero value otherwise.

### GetLibraryReferenceOk

`func (o *BTMaterialParams) GetLibraryReferenceOk() (*BTExternalElementReferenceInfo, bool)`

GetLibraryReferenceOk returns a tuple with the LibraryReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryReference

`func (o *BTMaterialParams) SetLibraryReference(v BTExternalElementReferenceInfo)`

SetLibraryReference sets LibraryReference field to given value.

### HasLibraryReference

`func (o *BTMaterialParams) HasLibraryReference() bool`

HasLibraryReference returns a boolean if a field has been set.

### GetProperties

`func (o *BTMaterialParams) GetProperties() []BTMaterialPropertyParams`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTMaterialParams) GetPropertiesOk() (*[]BTMaterialPropertyParams, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTMaterialParams) SetProperties(v []BTMaterialPropertyParams)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTMaterialParams) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


