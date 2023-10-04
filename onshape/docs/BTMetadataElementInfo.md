# BTMetadataElementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementId** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **int32** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Parts** | Pointer to [**BTMetadataObjectListInfoBTMetadataPartInfo**](BTMetadataObjectListInfoBTMetadataPartInfo.md) |  | [optional] 

## Methods

### NewBTMetadataElementInfo

`func NewBTMetadataElementInfo() *BTMetadataElementInfo`

NewBTMetadataElementInfo instantiates a new BTMetadataElementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataElementInfoWithDefaults

`func NewBTMetadataElementInfoWithDefaults() *BTMetadataElementInfo`

NewBTMetadataElementInfoWithDefaults instantiates a new BTMetadataElementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElementId

`func (o *BTMetadataElementInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTMetadataElementInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTMetadataElementInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTMetadataElementInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTMetadataElementInfo) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTMetadataElementInfo) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTMetadataElementInfo) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTMetadataElementInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetMimeType

`func (o *BTMetadataElementInfo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTMetadataElementInfo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTMetadataElementInfo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTMetadataElementInfo) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetParts

`func (o *BTMetadataElementInfo) GetParts() BTMetadataObjectListInfoBTMetadataPartInfo`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *BTMetadataElementInfo) GetPartsOk() (*BTMetadataObjectListInfoBTMetadataPartInfo, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *BTMetadataElementInfo) SetParts(v BTMetadataObjectListInfoBTMetadataPartInfo)`

SetParts sets Parts field to given value.

### HasParts

`func (o *BTMetadataElementInfo) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


