# BTAssemblyItemMetadataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]BTAssemblyItemMetadataInfo**](BTAssemblyItemMetadataInfo.md) |  | [optional] 
**PropertyIdToError** | Pointer to **map[string]string** |  | [optional] 
**PropertyIdToEvalInfo** | Pointer to **map[string]string** |  | [optional] 
**PropertyIdToOverrideStatus** | Pointer to **map[string]string** |  | [optional] 
**PropertyIdToSourceType** | Pointer to **map[string]string** |  | [optional] 
**PropertyIdToValue** | Pointer to **map[string]string** |  | [optional] 
**RequestInfo** | Pointer to [**BTAssemblyItemMetadataRequestInfo**](BTAssemblyItemMetadataRequestInfo.md) |  | [optional] 

## Methods

### NewBTAssemblyItemMetadataInfo

`func NewBTAssemblyItemMetadataInfo() *BTAssemblyItemMetadataInfo`

NewBTAssemblyItemMetadataInfo instantiates a new BTAssemblyItemMetadataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyItemMetadataInfoWithDefaults

`func NewBTAssemblyItemMetadataInfoWithDefaults() *BTAssemblyItemMetadataInfo`

NewBTAssemblyItemMetadataInfoWithDefaults instantiates a new BTAssemblyItemMetadataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *BTAssemblyItemMetadataInfo) GetChildren() []BTAssemblyItemMetadataInfo`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *BTAssemblyItemMetadataInfo) GetChildrenOk() (*[]BTAssemblyItemMetadataInfo, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *BTAssemblyItemMetadataInfo) SetChildren(v []BTAssemblyItemMetadataInfo)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *BTAssemblyItemMetadataInfo) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetPropertyIdToError

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToError() map[string]string`

GetPropertyIdToError returns the PropertyIdToError field if non-nil, zero value otherwise.

### GetPropertyIdToErrorOk

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToErrorOk() (*map[string]string, bool)`

GetPropertyIdToErrorOk returns a tuple with the PropertyIdToError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIdToError

`func (o *BTAssemblyItemMetadataInfo) SetPropertyIdToError(v map[string]string)`

SetPropertyIdToError sets PropertyIdToError field to given value.

### HasPropertyIdToError

`func (o *BTAssemblyItemMetadataInfo) HasPropertyIdToError() bool`

HasPropertyIdToError returns a boolean if a field has been set.

### GetPropertyIdToEvalInfo

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToEvalInfo() map[string]string`

GetPropertyIdToEvalInfo returns the PropertyIdToEvalInfo field if non-nil, zero value otherwise.

### GetPropertyIdToEvalInfoOk

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToEvalInfoOk() (*map[string]string, bool)`

GetPropertyIdToEvalInfoOk returns a tuple with the PropertyIdToEvalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIdToEvalInfo

`func (o *BTAssemblyItemMetadataInfo) SetPropertyIdToEvalInfo(v map[string]string)`

SetPropertyIdToEvalInfo sets PropertyIdToEvalInfo field to given value.

### HasPropertyIdToEvalInfo

`func (o *BTAssemblyItemMetadataInfo) HasPropertyIdToEvalInfo() bool`

HasPropertyIdToEvalInfo returns a boolean if a field has been set.

### GetPropertyIdToOverrideStatus

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToOverrideStatus() map[string]string`

GetPropertyIdToOverrideStatus returns the PropertyIdToOverrideStatus field if non-nil, zero value otherwise.

### GetPropertyIdToOverrideStatusOk

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToOverrideStatusOk() (*map[string]string, bool)`

GetPropertyIdToOverrideStatusOk returns a tuple with the PropertyIdToOverrideStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIdToOverrideStatus

`func (o *BTAssemblyItemMetadataInfo) SetPropertyIdToOverrideStatus(v map[string]string)`

SetPropertyIdToOverrideStatus sets PropertyIdToOverrideStatus field to given value.

### HasPropertyIdToOverrideStatus

`func (o *BTAssemblyItemMetadataInfo) HasPropertyIdToOverrideStatus() bool`

HasPropertyIdToOverrideStatus returns a boolean if a field has been set.

### GetPropertyIdToSourceType

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToSourceType() map[string]string`

GetPropertyIdToSourceType returns the PropertyIdToSourceType field if non-nil, zero value otherwise.

### GetPropertyIdToSourceTypeOk

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToSourceTypeOk() (*map[string]string, bool)`

GetPropertyIdToSourceTypeOk returns a tuple with the PropertyIdToSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIdToSourceType

`func (o *BTAssemblyItemMetadataInfo) SetPropertyIdToSourceType(v map[string]string)`

SetPropertyIdToSourceType sets PropertyIdToSourceType field to given value.

### HasPropertyIdToSourceType

`func (o *BTAssemblyItemMetadataInfo) HasPropertyIdToSourceType() bool`

HasPropertyIdToSourceType returns a boolean if a field has been set.

### GetPropertyIdToValue

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToValue() map[string]string`

GetPropertyIdToValue returns the PropertyIdToValue field if non-nil, zero value otherwise.

### GetPropertyIdToValueOk

`func (o *BTAssemblyItemMetadataInfo) GetPropertyIdToValueOk() (*map[string]string, bool)`

GetPropertyIdToValueOk returns a tuple with the PropertyIdToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyIdToValue

`func (o *BTAssemblyItemMetadataInfo) SetPropertyIdToValue(v map[string]string)`

SetPropertyIdToValue sets PropertyIdToValue field to given value.

### HasPropertyIdToValue

`func (o *BTAssemblyItemMetadataInfo) HasPropertyIdToValue() bool`

HasPropertyIdToValue returns a boolean if a field has been set.

### GetRequestInfo

`func (o *BTAssemblyItemMetadataInfo) GetRequestInfo() BTAssemblyItemMetadataRequestInfo`

GetRequestInfo returns the RequestInfo field if non-nil, zero value otherwise.

### GetRequestInfoOk

`func (o *BTAssemblyItemMetadataInfo) GetRequestInfoOk() (*BTAssemblyItemMetadataRequestInfo, bool)`

GetRequestInfoOk returns a tuple with the RequestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestInfo

`func (o *BTAssemblyItemMetadataInfo) SetRequestInfo(v BTAssemblyItemMetadataRequestInfo)`

SetRequestInfo sets RequestInfo field to given value.

### HasRequestInfo

`func (o *BTAssemblyItemMetadataInfo) HasRequestInfo() bool`

HasRequestInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


