# BTDisplayStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the view feature. | [optional] 
**IsOnshapeDefault** | Pointer to **bool** | &#x60;True&#x60; if this display state is in all assemblies by default; &#x60;false&#x60; if the display state is user-created. | [optional] 
**Name** | Pointer to **string** | The name of the view feature. | [optional] 

## Methods

### NewBTDisplayStateInfo

`func NewBTDisplayStateInfo() *BTDisplayStateInfo`

NewBTDisplayStateInfo instantiates a new BTDisplayStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDisplayStateInfoWithDefaults

`func NewBTDisplayStateInfoWithDefaults() *BTDisplayStateInfo`

NewBTDisplayStateInfoWithDefaults instantiates a new BTDisplayStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BTDisplayStateInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDisplayStateInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDisplayStateInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDisplayStateInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOnshapeDefault

`func (o *BTDisplayStateInfo) GetIsOnshapeDefault() bool`

GetIsOnshapeDefault returns the IsOnshapeDefault field if non-nil, zero value otherwise.

### GetIsOnshapeDefaultOk

`func (o *BTDisplayStateInfo) GetIsOnshapeDefaultOk() (*bool, bool)`

GetIsOnshapeDefaultOk returns a tuple with the IsOnshapeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnshapeDefault

`func (o *BTDisplayStateInfo) SetIsOnshapeDefault(v bool)`

SetIsOnshapeDefault sets IsOnshapeDefault field to given value.

### HasIsOnshapeDefault

`func (o *BTDisplayStateInfo) HasIsOnshapeDefault() bool`

HasIsOnshapeDefault returns a boolean if a field has been set.

### GetName

`func (o *BTDisplayStateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDisplayStateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDisplayStateInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDisplayStateInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


