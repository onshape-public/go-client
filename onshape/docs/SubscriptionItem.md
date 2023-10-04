# SubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubscriptionItem

`func NewSubscriptionItem() *SubscriptionItem`

NewSubscriptionItem instantiates a new SubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemWithDefaults

`func NewSubscriptionItemWithDefaults() *SubscriptionItem`

NewSubscriptionItemWithDefaults instantiates a new SubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SubscriptionItem) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubscriptionItem) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubscriptionItem) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SubscriptionItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *SubscriptionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *SubscriptionItem) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionItem) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionItem) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubscriptionItem) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPlan

`func (o *SubscriptionItem) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionItem) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionItem) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriptionItem) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SubscriptionItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


