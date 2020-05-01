# BTRootDiffInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to [**map[string]BTDiffInfo**](BTDiffInfo.md) |  | [optional] 
**CollectionChanges** | Pointer to [**map[string][]BTDiffInfo**](array.md) |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**GeometryChangeMessages** | Pointer to **[]string** |  | [optional] 
**SourceConfiguration** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceMicroversionId** | Pointer to **string** |  | [optional] 
**SourceValue** | Pointer to **string** |  | [optional] 
**SourceVersionId** | Pointer to **string** |  | [optional] 
**SourceWorkspaceId** | Pointer to **string** |  | [optional] 
**TargetConfiguration** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**TargetMicroversionId** | Pointer to **string** |  | [optional] 
**TargetValue** | Pointer to **string** |  | [optional] 
**TargetVersionId** | Pointer to **string** |  | [optional] 
**TargetWorkspaceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBTRootDiffInfo

`func NewBTRootDiffInfo() *BTRootDiffInfo`

NewBTRootDiffInfo instantiates a new BTRootDiffInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTRootDiffInfoWithDefaults

`func NewBTRootDiffInfoWithDefaults() *BTRootDiffInfo`

NewBTRootDiffInfoWithDefaults instantiates a new BTRootDiffInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *BTRootDiffInfo) GetChanges() map[string]BTDiffInfo`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *BTRootDiffInfo) GetChangesOk() (*map[string]BTDiffInfo, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *BTRootDiffInfo) SetChanges(v map[string]BTDiffInfo)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *BTRootDiffInfo) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetCollectionChanges

`func (o *BTRootDiffInfo) GetCollectionChanges() map[string][]BTDiffInfo`

GetCollectionChanges returns the CollectionChanges field if non-nil, zero value otherwise.

### GetCollectionChangesOk

`func (o *BTRootDiffInfo) GetCollectionChangesOk() (*map[string][]BTDiffInfo, bool)`

GetCollectionChangesOk returns a tuple with the CollectionChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionChanges

`func (o *BTRootDiffInfo) SetCollectionChanges(v map[string][]BTDiffInfo)`

SetCollectionChanges sets CollectionChanges field to given value.

### HasCollectionChanges

`func (o *BTRootDiffInfo) HasCollectionChanges() bool`

HasCollectionChanges returns a boolean if a field has been set.

### GetEntityType

`func (o *BTRootDiffInfo) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BTRootDiffInfo) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BTRootDiffInfo) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BTRootDiffInfo) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetGeometryChangeMessages

`func (o *BTRootDiffInfo) GetGeometryChangeMessages() []string`

GetGeometryChangeMessages returns the GeometryChangeMessages field if non-nil, zero value otherwise.

### GetGeometryChangeMessagesOk

`func (o *BTRootDiffInfo) GetGeometryChangeMessagesOk() (*[]string, bool)`

GetGeometryChangeMessagesOk returns a tuple with the GeometryChangeMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometryChangeMessages

`func (o *BTRootDiffInfo) SetGeometryChangeMessages(v []string)`

SetGeometryChangeMessages sets GeometryChangeMessages field to given value.

### HasGeometryChangeMessages

`func (o *BTRootDiffInfo) HasGeometryChangeMessages() bool`

HasGeometryChangeMessages returns a boolean if a field has been set.

### GetSourceConfiguration

`func (o *BTRootDiffInfo) GetSourceConfiguration() string`

GetSourceConfiguration returns the SourceConfiguration field if non-nil, zero value otherwise.

### GetSourceConfigurationOk

`func (o *BTRootDiffInfo) GetSourceConfigurationOk() (*string, bool)`

GetSourceConfigurationOk returns a tuple with the SourceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfiguration

`func (o *BTRootDiffInfo) SetSourceConfiguration(v string)`

SetSourceConfiguration sets SourceConfiguration field to given value.

### HasSourceConfiguration

`func (o *BTRootDiffInfo) HasSourceConfiguration() bool`

HasSourceConfiguration returns a boolean if a field has been set.

### GetSourceId

`func (o *BTRootDiffInfo) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BTRootDiffInfo) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BTRootDiffInfo) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BTRootDiffInfo) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceMicroversionId

`func (o *BTRootDiffInfo) GetSourceMicroversionId() string`

GetSourceMicroversionId returns the SourceMicroversionId field if non-nil, zero value otherwise.

### GetSourceMicroversionIdOk

`func (o *BTRootDiffInfo) GetSourceMicroversionIdOk() (*string, bool)`

GetSourceMicroversionIdOk returns a tuple with the SourceMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMicroversionId

`func (o *BTRootDiffInfo) SetSourceMicroversionId(v string)`

SetSourceMicroversionId sets SourceMicroversionId field to given value.

### HasSourceMicroversionId

`func (o *BTRootDiffInfo) HasSourceMicroversionId() bool`

HasSourceMicroversionId returns a boolean if a field has been set.

### GetSourceValue

`func (o *BTRootDiffInfo) GetSourceValue() string`

GetSourceValue returns the SourceValue field if non-nil, zero value otherwise.

### GetSourceValueOk

`func (o *BTRootDiffInfo) GetSourceValueOk() (*string, bool)`

GetSourceValueOk returns a tuple with the SourceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValue

`func (o *BTRootDiffInfo) SetSourceValue(v string)`

SetSourceValue sets SourceValue field to given value.

### HasSourceValue

`func (o *BTRootDiffInfo) HasSourceValue() bool`

HasSourceValue returns a boolean if a field has been set.

### GetSourceVersionId

`func (o *BTRootDiffInfo) GetSourceVersionId() string`

GetSourceVersionId returns the SourceVersionId field if non-nil, zero value otherwise.

### GetSourceVersionIdOk

`func (o *BTRootDiffInfo) GetSourceVersionIdOk() (*string, bool)`

GetSourceVersionIdOk returns a tuple with the SourceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersionId

`func (o *BTRootDiffInfo) SetSourceVersionId(v string)`

SetSourceVersionId sets SourceVersionId field to given value.

### HasSourceVersionId

`func (o *BTRootDiffInfo) HasSourceVersionId() bool`

HasSourceVersionId returns a boolean if a field has been set.

### GetSourceWorkspaceId

`func (o *BTRootDiffInfo) GetSourceWorkspaceId() string`

GetSourceWorkspaceId returns the SourceWorkspaceId field if non-nil, zero value otherwise.

### GetSourceWorkspaceIdOk

`func (o *BTRootDiffInfo) GetSourceWorkspaceIdOk() (*string, bool)`

GetSourceWorkspaceIdOk returns a tuple with the SourceWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceWorkspaceId

`func (o *BTRootDiffInfo) SetSourceWorkspaceId(v string)`

SetSourceWorkspaceId sets SourceWorkspaceId field to given value.

### HasSourceWorkspaceId

`func (o *BTRootDiffInfo) HasSourceWorkspaceId() bool`

HasSourceWorkspaceId returns a boolean if a field has been set.

### GetTargetConfiguration

`func (o *BTRootDiffInfo) GetTargetConfiguration() string`

GetTargetConfiguration returns the TargetConfiguration field if non-nil, zero value otherwise.

### GetTargetConfigurationOk

`func (o *BTRootDiffInfo) GetTargetConfigurationOk() (*string, bool)`

GetTargetConfigurationOk returns a tuple with the TargetConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConfiguration

`func (o *BTRootDiffInfo) SetTargetConfiguration(v string)`

SetTargetConfiguration sets TargetConfiguration field to given value.

### HasTargetConfiguration

`func (o *BTRootDiffInfo) HasTargetConfiguration() bool`

HasTargetConfiguration returns a boolean if a field has been set.

### GetTargetId

`func (o *BTRootDiffInfo) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *BTRootDiffInfo) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *BTRootDiffInfo) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *BTRootDiffInfo) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetMicroversionId

`func (o *BTRootDiffInfo) GetTargetMicroversionId() string`

GetTargetMicroversionId returns the TargetMicroversionId field if non-nil, zero value otherwise.

### GetTargetMicroversionIdOk

`func (o *BTRootDiffInfo) GetTargetMicroversionIdOk() (*string, bool)`

GetTargetMicroversionIdOk returns a tuple with the TargetMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMicroversionId

`func (o *BTRootDiffInfo) SetTargetMicroversionId(v string)`

SetTargetMicroversionId sets TargetMicroversionId field to given value.

### HasTargetMicroversionId

`func (o *BTRootDiffInfo) HasTargetMicroversionId() bool`

HasTargetMicroversionId returns a boolean if a field has been set.

### GetTargetValue

`func (o *BTRootDiffInfo) GetTargetValue() string`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *BTRootDiffInfo) GetTargetValueOk() (*string, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *BTRootDiffInfo) SetTargetValue(v string)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *BTRootDiffInfo) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetTargetVersionId

`func (o *BTRootDiffInfo) GetTargetVersionId() string`

GetTargetVersionId returns the TargetVersionId field if non-nil, zero value otherwise.

### GetTargetVersionIdOk

`func (o *BTRootDiffInfo) GetTargetVersionIdOk() (*string, bool)`

GetTargetVersionIdOk returns a tuple with the TargetVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersionId

`func (o *BTRootDiffInfo) SetTargetVersionId(v string)`

SetTargetVersionId sets TargetVersionId field to given value.

### HasTargetVersionId

`func (o *BTRootDiffInfo) HasTargetVersionId() bool`

HasTargetVersionId returns a boolean if a field has been set.

### GetTargetWorkspaceId

`func (o *BTRootDiffInfo) GetTargetWorkspaceId() string`

GetTargetWorkspaceId returns the TargetWorkspaceId field if non-nil, zero value otherwise.

### GetTargetWorkspaceIdOk

`func (o *BTRootDiffInfo) GetTargetWorkspaceIdOk() (*string, bool)`

GetTargetWorkspaceIdOk returns a tuple with the TargetWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWorkspaceId

`func (o *BTRootDiffInfo) SetTargetWorkspaceId(v string)`

SetTargetWorkspaceId sets TargetWorkspaceId field to given value.

### HasTargetWorkspaceId

`func (o *BTRootDiffInfo) HasTargetWorkspaceId() bool`

HasTargetWorkspaceId returns a boolean if a field has been set.

### GetType

`func (o *BTRootDiffInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTRootDiffInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTRootDiffInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTRootDiffInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


