# PathItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delete** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Get** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Getref** | Pointer to **string** |  | [optional] 
**Head** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Options** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Parameters** | Pointer to [**[]Parameter**](Parameter.md) |  | [optional] 
**Patch** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Post** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Put** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Servers** | Pointer to [**[]Server**](Server.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Trace** | Pointer to [**Operation**](Operation.md) |  | [optional] 

## Methods

### NewPathItem

`func NewPathItem() *PathItem`

NewPathItem instantiates a new PathItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathItemWithDefaults

`func NewPathItemWithDefaults() *PathItem`

NewPathItemWithDefaults instantiates a new PathItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelete

`func (o *PathItem) GetDelete() Operation`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *PathItem) GetDeleteOk() (*Operation, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *PathItem) SetDelete(v Operation)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *PathItem) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescription

`func (o *PathItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PathItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PathItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PathItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtensions

`func (o *PathItem) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *PathItem) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *PathItem) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *PathItem) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGet

`func (o *PathItem) GetGet() Operation`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *PathItem) GetGetOk() (*Operation, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *PathItem) SetGet(v Operation)`

SetGet sets Get field to given value.

### HasGet

`func (o *PathItem) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetGetref

`func (o *PathItem) GetGetref() string`

GetGetref returns the Getref field if non-nil, zero value otherwise.

### GetGetrefOk

`func (o *PathItem) GetGetrefOk() (*string, bool)`

GetGetrefOk returns a tuple with the Getref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetref

`func (o *PathItem) SetGetref(v string)`

SetGetref sets Getref field to given value.

### HasGetref

`func (o *PathItem) HasGetref() bool`

HasGetref returns a boolean if a field has been set.

### GetHead

`func (o *PathItem) GetHead() Operation`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PathItem) GetHeadOk() (*Operation, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PathItem) SetHead(v Operation)`

SetHead sets Head field to given value.

### HasHead

`func (o *PathItem) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetOptions

`func (o *PathItem) GetOptions() Operation`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PathItem) GetOptionsOk() (*Operation, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PathItem) SetOptions(v Operation)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PathItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetParameters

`func (o *PathItem) GetParameters() []Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PathItem) GetParametersOk() (*[]Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PathItem) SetParameters(v []Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PathItem) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPatch

`func (o *PathItem) GetPatch() Operation`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *PathItem) GetPatchOk() (*Operation, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *PathItem) SetPatch(v Operation)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *PathItem) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetPost

`func (o *PathItem) GetPost() Operation`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *PathItem) GetPostOk() (*Operation, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *PathItem) SetPost(v Operation)`

SetPost sets Post field to given value.

### HasPost

`func (o *PathItem) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPut

`func (o *PathItem) GetPut() Operation`

GetPut returns the Put field if non-nil, zero value otherwise.

### GetPutOk

`func (o *PathItem) GetPutOk() (*Operation, bool)`

GetPutOk returns a tuple with the Put field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPut

`func (o *PathItem) SetPut(v Operation)`

SetPut sets Put field to given value.

### HasPut

`func (o *PathItem) HasPut() bool`

HasPut returns a boolean if a field has been set.

### GetServers

`func (o *PathItem) GetServers() []Server`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PathItem) GetServersOk() (*[]Server, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PathItem) SetServers(v []Server)`

SetServers sets Servers field to given value.

### HasServers

`func (o *PathItem) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSummary

`func (o *PathItem) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PathItem) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PathItem) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PathItem) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTrace

`func (o *PathItem) GetTrace() Operation`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *PathItem) GetTraceOk() (*Operation, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *PathItem) SetTrace(v Operation)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *PathItem) HasTrace() bool`

HasTrace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


