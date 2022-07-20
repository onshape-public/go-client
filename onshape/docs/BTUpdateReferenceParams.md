# BTUpdateReferenceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** |  | [optional] 
**EditDescription** | Pointer to **string** |  | [optional] 
**ReferenceUpdates** | Pointer to [**[]UpdateParams**](UpdateParams.md) |  | [optional] 

## Methods

### NewBTUpdateReferenceParams

`func NewBTUpdateReferenceParams() *BTUpdateReferenceParams`

NewBTUpdateReferenceParams instantiates a new BTUpdateReferenceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUpdateReferenceParamsWithDefaults

`func NewBTUpdateReferenceParamsWithDefaults() *BTUpdateReferenceParams`

NewBTUpdateReferenceParamsWithDefaults instantiates a new BTUpdateReferenceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *BTUpdateReferenceParams) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BTUpdateReferenceParams) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BTUpdateReferenceParams) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *BTUpdateReferenceParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetEditDescription

`func (o *BTUpdateReferenceParams) GetEditDescription() string`

GetEditDescription returns the EditDescription field if non-nil, zero value otherwise.

### GetEditDescriptionOk

`func (o *BTUpdateReferenceParams) GetEditDescriptionOk() (*string, bool)`

GetEditDescriptionOk returns a tuple with the EditDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditDescription

`func (o *BTUpdateReferenceParams) SetEditDescription(v string)`

SetEditDescription sets EditDescription field to given value.

### HasEditDescription

`func (o *BTUpdateReferenceParams) HasEditDescription() bool`

HasEditDescription returns a boolean if a field has been set.

### GetReferenceUpdates

`func (o *BTUpdateReferenceParams) GetReferenceUpdates() []UpdateParams`

GetReferenceUpdates returns the ReferenceUpdates field if non-nil, zero value otherwise.

### GetReferenceUpdatesOk

`func (o *BTUpdateReferenceParams) GetReferenceUpdatesOk() (*[]UpdateParams, bool)`

GetReferenceUpdatesOk returns a tuple with the ReferenceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceUpdates

`func (o *BTUpdateReferenceParams) SetReferenceUpdates(v []UpdateParams)`

SetReferenceUpdates sets ReferenceUpdates field to given value.

### HasReferenceUpdates

`func (o *BTUpdateReferenceParams) HasReferenceUpdates() bool`

HasReferenceUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


