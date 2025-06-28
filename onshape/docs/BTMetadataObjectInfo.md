# BTMetadataObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**JsonType** | **string** |  | 
**Properties** | Pointer to [**[]BTMetadataPropertyInfo**](BTMetadataPropertyInfo.md) | Properties associated with this metadata object | [optional] 
**Thumbnail** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 

## Methods

### NewBTMetadataObjectInfo

`func NewBTMetadataObjectInfo(jsonType string, ) *BTMetadataObjectInfo`

NewBTMetadataObjectInfo instantiates a new BTMetadataObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataObjectInfoWithDefaults

`func NewBTMetadataObjectInfoWithDefaults() *BTMetadataObjectInfo`

NewBTMetadataObjectInfoWithDefaults instantiates a new BTMetadataObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTMetadataObjectInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTMetadataObjectInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTMetadataObjectInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTMetadataObjectInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetJsonType

`func (o *BTMetadataObjectInfo) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *BTMetadataObjectInfo) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *BTMetadataObjectInfo) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.


### GetProperties

`func (o *BTMetadataObjectInfo) GetProperties() []BTMetadataPropertyInfo`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTMetadataObjectInfo) GetPropertiesOk() (*[]BTMetadataPropertyInfo, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTMetadataObjectInfo) SetProperties(v []BTMetadataPropertyInfo)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTMetadataObjectInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetThumbnail

`func (o *BTMetadataObjectInfo) GetThumbnail() BTThumbnailInfo`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *BTMetadataObjectInfo) GetThumbnailOk() (*BTThumbnailInfo, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *BTMetadataObjectInfo) SetThumbnail(v BTThumbnailInfo)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *BTMetadataObjectInfo) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


