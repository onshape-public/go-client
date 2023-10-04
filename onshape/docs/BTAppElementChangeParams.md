# BTAppElementChangeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseContent** | Pointer to **string** | This base64-encoded data is treated as a full representation of the sub-element&#39;s data and all deltas previously added will no longer be returned. | [optional] 
**Delta** | Pointer to **string** | This base64-encoded data is a delta which the application can apply to the last added baseContent data. | [optional] 
**SubelementId** | **string** | The id of the subelement to edit. The value is determined by the application. | 

## Methods

### NewBTAppElementChangeParams

`func NewBTAppElementChangeParams(subelementId string, ) *BTAppElementChangeParams`

NewBTAppElementChangeParams instantiates a new BTAppElementChangeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementChangeParamsWithDefaults

`func NewBTAppElementChangeParamsWithDefaults() *BTAppElementChangeParams`

NewBTAppElementChangeParamsWithDefaults instantiates a new BTAppElementChangeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseContent

`func (o *BTAppElementChangeParams) GetBaseContent() string`

GetBaseContent returns the BaseContent field if non-nil, zero value otherwise.

### GetBaseContentOk

`func (o *BTAppElementChangeParams) GetBaseContentOk() (*string, bool)`

GetBaseContentOk returns a tuple with the BaseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContent

`func (o *BTAppElementChangeParams) SetBaseContent(v string)`

SetBaseContent sets BaseContent field to given value.

### HasBaseContent

`func (o *BTAppElementChangeParams) HasBaseContent() bool`

HasBaseContent returns a boolean if a field has been set.

### GetDelta

`func (o *BTAppElementChangeParams) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *BTAppElementChangeParams) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *BTAppElementChangeParams) SetDelta(v string)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *BTAppElementChangeParams) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetSubelementId

`func (o *BTAppElementChangeParams) GetSubelementId() string`

GetSubelementId returns the SubelementId field if non-nil, zero value otherwise.

### GetSubelementIdOk

`func (o *BTAppElementChangeParams) GetSubelementIdOk() (*string, bool)`

GetSubelementIdOk returns a tuple with the SubelementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelementId

`func (o *BTAppElementChangeParams) SetSubelementId(v string)`

SetSubelementId sets SubelementId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


