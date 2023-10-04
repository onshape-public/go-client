# BTUnchangedElementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**UnchangedReferences** | Pointer to [**[]BTUnchangedReferenceInfo**](BTUnchangedReferenceInfo.md) |  | [optional] 

## Methods

### NewBTUnchangedElementInfo

`func NewBTUnchangedElementInfo() *BTUnchangedElementInfo`

NewBTUnchangedElementInfo instantiates a new BTUnchangedElementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUnchangedElementInfoWithDefaults

`func NewBTUnchangedElementInfoWithDefaults() *BTUnchangedElementInfo`

NewBTUnchangedElementInfoWithDefaults instantiates a new BTUnchangedElementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *BTUnchangedElementInfo) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BTUnchangedElementInfo) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BTUnchangedElementInfo) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *BTUnchangedElementInfo) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetElementId

`func (o *BTUnchangedElementInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTUnchangedElementInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTUnchangedElementInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTUnchangedElementInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetUnchangedReferences

`func (o *BTUnchangedElementInfo) GetUnchangedReferences() []BTUnchangedReferenceInfo`

GetUnchangedReferences returns the UnchangedReferences field if non-nil, zero value otherwise.

### GetUnchangedReferencesOk

`func (o *BTUnchangedElementInfo) GetUnchangedReferencesOk() (*[]BTUnchangedReferenceInfo, bool)`

GetUnchangedReferencesOk returns a tuple with the UnchangedReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnchangedReferences

`func (o *BTUnchangedElementInfo) SetUnchangedReferences(v []BTUnchangedReferenceInfo)`

SetUnchangedReferences sets UnchangedReferences field to given value.

### HasUnchangedReferences

`func (o *BTUnchangedElementInfo) HasUnchangedReferences() bool`

HasUnchangedReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


