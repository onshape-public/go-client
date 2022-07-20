# BTElementLocationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementId** | Pointer to **string** | The id of an element which provides context for the position value specified. | [optional] 
**Position** | Pointer to **int32** | An indicator for the relative placement of the new element. If elementId is specified, a negative number indicates insertion prior to the element and a non-negative number indicates insertion following the element. If no elementId is specified, a negative value indicates insertion at the end of the element list and a non-negative number indicates insertion at the start of the element list. | [optional] 

## Methods

### NewBTElementLocationParams

`func NewBTElementLocationParams() *BTElementLocationParams`

NewBTElementLocationParams instantiates a new BTElementLocationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTElementLocationParamsWithDefaults

`func NewBTElementLocationParamsWithDefaults() *BTElementLocationParams`

NewBTElementLocationParamsWithDefaults instantiates a new BTElementLocationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElementId

`func (o *BTElementLocationParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTElementLocationParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTElementLocationParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTElementLocationParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetPosition

`func (o *BTElementLocationParams) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BTElementLocationParams) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BTElementLocationParams) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BTElementLocationParams) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


