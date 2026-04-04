# BTAssemblyParametricInstanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]BTAssemblyParametricInstanceChildInfo**](BTAssemblyParametricInstanceChildInfo.md) | Child instances. | [optional] 
**Id** | Pointer to **string** | Id of the Part Studio instance. | [optional] 
**Name** | Pointer to **string** | Name of the parametric instance. | [optional] 
**Suppressed** | Pointer to **bool** | If the parametric is suppressed. | [optional] 
**Type** | Pointer to **string** | Type of parametric instance. | [optional] 

## Methods

### NewBTAssemblyParametricInstanceInfo

`func NewBTAssemblyParametricInstanceInfo() *BTAssemblyParametricInstanceInfo`

NewBTAssemblyParametricInstanceInfo instantiates a new BTAssemblyParametricInstanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyParametricInstanceInfoWithDefaults

`func NewBTAssemblyParametricInstanceInfoWithDefaults() *BTAssemblyParametricInstanceInfo`

NewBTAssemblyParametricInstanceInfoWithDefaults instantiates a new BTAssemblyParametricInstanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *BTAssemblyParametricInstanceInfo) GetChildren() []BTAssemblyParametricInstanceChildInfo`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *BTAssemblyParametricInstanceInfo) GetChildrenOk() (*[]BTAssemblyParametricInstanceChildInfo, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *BTAssemblyParametricInstanceInfo) SetChildren(v []BTAssemblyParametricInstanceChildInfo)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *BTAssemblyParametricInstanceInfo) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetId

`func (o *BTAssemblyParametricInstanceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTAssemblyParametricInstanceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTAssemblyParametricInstanceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTAssemblyParametricInstanceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTAssemblyParametricInstanceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAssemblyParametricInstanceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAssemblyParametricInstanceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAssemblyParametricInstanceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTAssemblyParametricInstanceInfo) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTAssemblyParametricInstanceInfo) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTAssemblyParametricInstanceInfo) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTAssemblyParametricInstanceInfo) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetType

`func (o *BTAssemblyParametricInstanceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTAssemblyParametricInstanceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTAssemblyParametricInstanceInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTAssemblyParametricInstanceInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


