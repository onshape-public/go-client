# BTMSketchGeomEntity5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ControlBoxIds** | Pointer to **[]string** |  | [optional] 
**EntityType** | Pointer to [**GBTSketchEntityType**](GBTSketchEntityType.md) |  | [optional] 
**FunctionName** | Pointer to **string** |  | [optional] 
**IsConstruction** | Pointer to **bool** |  | [optional] 
**IsFromEndpointSplineHandle** | Pointer to **bool** |  | [optional] 
**IsFromSplineControlPolygon** | Pointer to **bool** |  | [optional] 
**IsFromSplineHandle** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTMSketchGeomEntity5

`func NewBTMSketchGeomEntity5() *BTMSketchGeomEntity5`

NewBTMSketchGeomEntity5 instantiates a new BTMSketchGeomEntity5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMSketchGeomEntity5WithDefaults

`func NewBTMSketchGeomEntity5WithDefaults() *BTMSketchGeomEntity5`

NewBTMSketchGeomEntity5WithDefaults instantiates a new BTMSketchGeomEntity5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTMSketchGeomEntity5) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMSketchGeomEntity5) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMSketchGeomEntity5) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMSketchGeomEntity5) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetControlBoxIds

`func (o *BTMSketchGeomEntity5) GetControlBoxIds() []string`

GetControlBoxIds returns the ControlBoxIds field if non-nil, zero value otherwise.

### GetControlBoxIdsOk

`func (o *BTMSketchGeomEntity5) GetControlBoxIdsOk() (*[]string, bool)`

GetControlBoxIdsOk returns a tuple with the ControlBoxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlBoxIds

`func (o *BTMSketchGeomEntity5) SetControlBoxIds(v []string)`

SetControlBoxIds sets ControlBoxIds field to given value.

### HasControlBoxIds

`func (o *BTMSketchGeomEntity5) HasControlBoxIds() bool`

HasControlBoxIds returns a boolean if a field has been set.

### GetEntityType

`func (o *BTMSketchGeomEntity5) GetEntityType() GBTSketchEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BTMSketchGeomEntity5) GetEntityTypeOk() (*GBTSketchEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BTMSketchGeomEntity5) SetEntityType(v GBTSketchEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BTMSketchGeomEntity5) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFunctionName

`func (o *BTMSketchGeomEntity5) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *BTMSketchGeomEntity5) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *BTMSketchGeomEntity5) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *BTMSketchGeomEntity5) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetIsConstruction

`func (o *BTMSketchGeomEntity5) GetIsConstruction() bool`

GetIsConstruction returns the IsConstruction field if non-nil, zero value otherwise.

### GetIsConstructionOk

`func (o *BTMSketchGeomEntity5) GetIsConstructionOk() (*bool, bool)`

GetIsConstructionOk returns a tuple with the IsConstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstruction

`func (o *BTMSketchGeomEntity5) SetIsConstruction(v bool)`

SetIsConstruction sets IsConstruction field to given value.

### HasIsConstruction

`func (o *BTMSketchGeomEntity5) HasIsConstruction() bool`

HasIsConstruction returns a boolean if a field has been set.

### GetIsFromEndpointSplineHandle

`func (o *BTMSketchGeomEntity5) GetIsFromEndpointSplineHandle() bool`

GetIsFromEndpointSplineHandle returns the IsFromEndpointSplineHandle field if non-nil, zero value otherwise.

### GetIsFromEndpointSplineHandleOk

`func (o *BTMSketchGeomEntity5) GetIsFromEndpointSplineHandleOk() (*bool, bool)`

GetIsFromEndpointSplineHandleOk returns a tuple with the IsFromEndpointSplineHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromEndpointSplineHandle

`func (o *BTMSketchGeomEntity5) SetIsFromEndpointSplineHandle(v bool)`

SetIsFromEndpointSplineHandle sets IsFromEndpointSplineHandle field to given value.

### HasIsFromEndpointSplineHandle

`func (o *BTMSketchGeomEntity5) HasIsFromEndpointSplineHandle() bool`

HasIsFromEndpointSplineHandle returns a boolean if a field has been set.

### GetIsFromSplineControlPolygon

`func (o *BTMSketchGeomEntity5) GetIsFromSplineControlPolygon() bool`

GetIsFromSplineControlPolygon returns the IsFromSplineControlPolygon field if non-nil, zero value otherwise.

### GetIsFromSplineControlPolygonOk

`func (o *BTMSketchGeomEntity5) GetIsFromSplineControlPolygonOk() (*bool, bool)`

GetIsFromSplineControlPolygonOk returns a tuple with the IsFromSplineControlPolygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSplineControlPolygon

`func (o *BTMSketchGeomEntity5) SetIsFromSplineControlPolygon(v bool)`

SetIsFromSplineControlPolygon sets IsFromSplineControlPolygon field to given value.

### HasIsFromSplineControlPolygon

`func (o *BTMSketchGeomEntity5) HasIsFromSplineControlPolygon() bool`

HasIsFromSplineControlPolygon returns a boolean if a field has been set.

### GetIsFromSplineHandle

`func (o *BTMSketchGeomEntity5) GetIsFromSplineHandle() bool`

GetIsFromSplineHandle returns the IsFromSplineHandle field if non-nil, zero value otherwise.

### GetIsFromSplineHandleOk

`func (o *BTMSketchGeomEntity5) GetIsFromSplineHandleOk() (*bool, bool)`

GetIsFromSplineHandleOk returns a tuple with the IsFromSplineHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFromSplineHandle

`func (o *BTMSketchGeomEntity5) SetIsFromSplineHandle(v bool)`

SetIsFromSplineHandle sets IsFromSplineHandle field to given value.

### HasIsFromSplineHandle

`func (o *BTMSketchGeomEntity5) HasIsFromSplineHandle() bool`

HasIsFromSplineHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


