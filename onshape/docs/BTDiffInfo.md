# BTDiffInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**map[string]BTDiffInfo**](BTDiffInfo.md) |  | [optional] 
**EntityType** | Pointer to [**BTDiffInfoCollectionType**](BTDiffInfoCollectionType.md) |  | [optional] 
**GeometryChangeMessages** | Pointer to **[]string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceValue** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**TargetValue** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GBTNodeChange**](GBTNodeChange.md) |  | [optional] 

## Methods

### NewBTDiffInfo

`func NewBTDiffInfo() *BTDiffInfo`

NewBTDiffInfo instantiates a new BTDiffInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDiffInfoWithDefaults

`func NewBTDiffInfoWithDefaults() *BTDiffInfo`

NewBTDiffInfoWithDefaults instantiates a new BTDiffInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *BTDiffInfo) GetChanges() map[string]BTDiffInfo`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *BTDiffInfo) GetChangesOk() (*map[string]BTDiffInfo, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *BTDiffInfo) SetChanges(v map[string]BTDiffInfo)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *BTDiffInfo) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetEntityType

`func (o *BTDiffInfo) GetEntityType() BTDiffInfoCollectionType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BTDiffInfo) GetEntityTypeOk() (*BTDiffInfoCollectionType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BTDiffInfo) SetEntityType(v BTDiffInfoCollectionType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BTDiffInfo) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetGeometryChangeMessages

`func (o *BTDiffInfo) GetGeometryChangeMessages() []string`

GetGeometryChangeMessages returns the GeometryChangeMessages field if non-nil, zero value otherwise.

### GetGeometryChangeMessagesOk

`func (o *BTDiffInfo) GetGeometryChangeMessagesOk() (*[]string, bool)`

GetGeometryChangeMessagesOk returns a tuple with the GeometryChangeMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometryChangeMessages

`func (o *BTDiffInfo) SetGeometryChangeMessages(v []string)`

SetGeometryChangeMessages sets GeometryChangeMessages field to given value.

### HasGeometryChangeMessages

`func (o *BTDiffInfo) HasGeometryChangeMessages() bool`

HasGeometryChangeMessages returns a boolean if a field has been set.

### GetSourceId

`func (o *BTDiffInfo) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BTDiffInfo) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BTDiffInfo) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BTDiffInfo) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceValue

`func (o *BTDiffInfo) GetSourceValue() string`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *BTDiffInfo) GetSourceValueOk() (*string, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *BTDiffInfo) SetSourceValue(v string)`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *BTDiffInfo) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.

### GetTargetId

`func (o *BTDiffInfo) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *BTDiffInfo) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *BTDiffInfo) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *BTDiffInfo) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetValue

`func (o *BTDiffInfo) GetTargetValue() string`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *BTDiffInfo) GetTargetValueOk() (*string, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *BTDiffInfo) SetTargetValue(v string)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *BTDiffInfo) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetType

`func (o *BTDiffInfo) GetType() GBTNodeChange`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTDiffInfo) GetTypeOk() (*GBTNodeChange, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTDiffInfo) SetType(v GBTNodeChange)`

SetType sets Type field to given value.

### HasType

`func (o *BTDiffInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


