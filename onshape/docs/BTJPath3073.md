# BTJPath3073

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**Path** | Pointer to [**[]BTJPathElement2297**](BTJPathElement2297.md) |  | [optional] 
**StartNode** | **string** | Either empty (root) or the nodeId of a node in the tree. | 

## Methods

### NewBTJPath3073

`func NewBTJPath3073(startNode string, ) *BTJPath3073`

NewBTJPath3073 instantiates a new BTJPath3073 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTJPath3073WithDefaults

`func NewBTJPath3073WithDefaults() *BTJPath3073`

NewBTJPath3073WithDefaults instantiates a new BTJPath3073 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTJPath3073) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTJPath3073) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTJPath3073) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTJPath3073) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetPath

`func (o *BTJPath3073) GetPath() []BTJPathElement2297`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BTJPath3073) GetPathOk() (*[]BTJPathElement2297, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BTJPath3073) SetPath(v []BTJPathElement2297)`

SetPath sets Path field to given value.

### HasPath

`func (o *BTJPath3073) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStartNode

`func (o *BTJPath3073) GetStartNode() string`

GetStartNode returns the StartNode field if non-nil, zero value otherwise.

### GetStartNodeOk

`func (o *BTJPath3073) GetStartNodeOk() (*string, bool)`

GetStartNodeOk returns a tuple with the StartNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartNode

`func (o *BTJPath3073) SetStartNode(v string)`

SetStartNode sets StartNode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


