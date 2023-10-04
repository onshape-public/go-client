# BTMetadataPropertyUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **string** | The id of the property that should be edited. This can be retrieved from MetadataCategory:getCategoryProperties. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The new value for the property. | [optional] 

## Methods

### NewBTMetadataPropertyUpdateParams

`func NewBTMetadataPropertyUpdateParams() *BTMetadataPropertyUpdateParams`

NewBTMetadataPropertyUpdateParams instantiates a new BTMetadataPropertyUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMetadataPropertyUpdateParamsWithDefaults

`func NewBTMetadataPropertyUpdateParamsWithDefaults() *BTMetadataPropertyUpdateParams`

NewBTMetadataPropertyUpdateParamsWithDefaults instantiates a new BTMetadataPropertyUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *BTMetadataPropertyUpdateParams) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *BTMetadataPropertyUpdateParams) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *BTMetadataPropertyUpdateParams) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *BTMetadataPropertyUpdateParams) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetValue

`func (o *BTMetadataPropertyUpdateParams) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BTMetadataPropertyUpdateParams) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BTMetadataPropertyUpdateParams) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BTMetadataPropertyUpdateParams) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


