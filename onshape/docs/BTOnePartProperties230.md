# BTOnePartProperties230

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appearance** | Pointer to [**BTGraphicsAppearance1152**](BTGraphicsAppearance1152.md) |  | [optional] 
**AppearanceForNewCell** | Pointer to [**BTGraphicsAppearance1152**](BTGraphicsAppearance1152.md) |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ChangedPropertiesSet** | Pointer to **[]string** |  | [optional] 
**CustomProperties** | Pointer to [**BTPartCustomProperties1338**](BTPartCustomProperties1338.md) |  | [optional] 
**Material** | Pointer to [**BTPartMaterial1445**](BTPartMaterial1445.md) |  | [optional] 
**MaterialForNewCell** | Pointer to [**BTPartMaterial1445**](BTPartMaterial1445.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameForNewCell** | Pointer to **string** |  | [optional] 
**NameIfNotNull** | Pointer to [**BTOnePartProperties230**](BTOnePartProperties230.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ParsedQuery** | Pointer to [**BTPFunctionDeclaration246**](BTPFunctionDeclaration246.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**QueryListParameter** | Pointer to [**BTMParameterQueryList148**](BTMParameterQueryList148.md) |  | [optional] 
**SheetMetalBendOrder** | Pointer to **[]string** |  | [optional] 
**SheetMetalBendOrderIfNotNull** | Pointer to [**BTOnePartProperties230**](BTOnePartProperties230.md) |  | [optional] 
**Visibility** | Pointer to [**GBTPartVisibility**](GBTPartVisibility.md) |  | [optional] 

## Methods

### NewBTOnePartProperties230

`func NewBTOnePartProperties230() *BTOnePartProperties230`

NewBTOnePartProperties230 instantiates a new BTOnePartProperties230 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTOnePartProperties230WithDefaults

`func NewBTOnePartProperties230WithDefaults() *BTOnePartProperties230`

NewBTOnePartProperties230WithDefaults instantiates a new BTOnePartProperties230 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearance

`func (o *BTOnePartProperties230) GetAppearance() BTGraphicsAppearance1152`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *BTOnePartProperties230) GetAppearanceOk() (*BTGraphicsAppearance1152, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *BTOnePartProperties230) SetAppearance(v BTGraphicsAppearance1152)`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *BTOnePartProperties230) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### GetAppearanceForNewCell

`func (o *BTOnePartProperties230) GetAppearanceForNewCell() BTGraphicsAppearance1152`

GetAppearanceForNewCell returns the AppearanceForNewCell field if non-nil, zero value otherwise.

### GetAppearanceForNewCellOk

`func (o *BTOnePartProperties230) GetAppearanceForNewCellOk() (*BTGraphicsAppearance1152, bool)`

GetAppearanceForNewCellOk returns a tuple with the AppearanceForNewCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearanceForNewCell

`func (o *BTOnePartProperties230) SetAppearanceForNewCell(v BTGraphicsAppearance1152)`

SetAppearanceForNewCell sets AppearanceForNewCell field to given value.

### HasAppearanceForNewCell

`func (o *BTOnePartProperties230) HasAppearanceForNewCell() bool`

HasAppearanceForNewCell returns a boolean if a field has been set.

### GetBtType

`func (o *BTOnePartProperties230) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTOnePartProperties230) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTOnePartProperties230) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTOnePartProperties230) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetChangedPropertiesSet

`func (o *BTOnePartProperties230) GetChangedPropertiesSet() []string`

GetChangedPropertiesSet returns the ChangedPropertiesSet field if non-nil, zero value otherwise.

### GetChangedPropertiesSetOk

`func (o *BTOnePartProperties230) GetChangedPropertiesSetOk() (*[]string, bool)`

GetChangedPropertiesSetOk returns a tuple with the ChangedPropertiesSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedPropertiesSet

`func (o *BTOnePartProperties230) SetChangedPropertiesSet(v []string)`

SetChangedPropertiesSet sets ChangedPropertiesSet field to given value.

### HasChangedPropertiesSet

`func (o *BTOnePartProperties230) HasChangedPropertiesSet() bool`

HasChangedPropertiesSet returns a boolean if a field has been set.

### GetCustomProperties

`func (o *BTOnePartProperties230) GetCustomProperties() BTPartCustomProperties1338`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BTOnePartProperties230) GetCustomPropertiesOk() (*BTPartCustomProperties1338, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BTOnePartProperties230) SetCustomProperties(v BTPartCustomProperties1338)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BTOnePartProperties230) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetMaterial

`func (o *BTOnePartProperties230) GetMaterial() BTPartMaterial1445`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *BTOnePartProperties230) GetMaterialOk() (*BTPartMaterial1445, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *BTOnePartProperties230) SetMaterial(v BTPartMaterial1445)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *BTOnePartProperties230) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetMaterialForNewCell

`func (o *BTOnePartProperties230) GetMaterialForNewCell() BTPartMaterial1445`

GetMaterialForNewCell returns the MaterialForNewCell field if non-nil, zero value otherwise.

### GetMaterialForNewCellOk

`func (o *BTOnePartProperties230) GetMaterialForNewCellOk() (*BTPartMaterial1445, bool)`

GetMaterialForNewCellOk returns a tuple with the MaterialForNewCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialForNewCell

`func (o *BTOnePartProperties230) SetMaterialForNewCell(v BTPartMaterial1445)`

SetMaterialForNewCell sets MaterialForNewCell field to given value.

### HasMaterialForNewCell

`func (o *BTOnePartProperties230) HasMaterialForNewCell() bool`

HasMaterialForNewCell returns a boolean if a field has been set.

### GetName

`func (o *BTOnePartProperties230) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTOnePartProperties230) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTOnePartProperties230) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTOnePartProperties230) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameForNewCell

`func (o *BTOnePartProperties230) GetNameForNewCell() string`

GetNameForNewCell returns the NameForNewCell field if non-nil, zero value otherwise.

### GetNameForNewCellOk

`func (o *BTOnePartProperties230) GetNameForNewCellOk() (*string, bool)`

GetNameForNewCellOk returns a tuple with the NameForNewCell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameForNewCell

`func (o *BTOnePartProperties230) SetNameForNewCell(v string)`

SetNameForNewCell sets NameForNewCell field to given value.

### HasNameForNewCell

`func (o *BTOnePartProperties230) HasNameForNewCell() bool`

HasNameForNewCell returns a boolean if a field has been set.

### GetNameIfNotNull

`func (o *BTOnePartProperties230) GetNameIfNotNull() BTOnePartProperties230`

GetNameIfNotNull returns the NameIfNotNull field if non-nil, zero value otherwise.

### GetNameIfNotNullOk

`func (o *BTOnePartProperties230) GetNameIfNotNullOk() (*BTOnePartProperties230, bool)`

GetNameIfNotNullOk returns a tuple with the NameIfNotNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIfNotNull

`func (o *BTOnePartProperties230) SetNameIfNotNull(v BTOnePartProperties230)`

SetNameIfNotNull sets NameIfNotNull field to given value.

### HasNameIfNotNull

`func (o *BTOnePartProperties230) HasNameIfNotNull() bool`

HasNameIfNotNull returns a boolean if a field has been set.

### GetNodeId

`func (o *BTOnePartProperties230) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTOnePartProperties230) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTOnePartProperties230) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTOnePartProperties230) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetParsedQuery

`func (o *BTOnePartProperties230) GetParsedQuery() BTPFunctionDeclaration246`

GetParsedQuery returns the ParsedQuery field if non-nil, zero value otherwise.

### GetParsedQueryOk

`func (o *BTOnePartProperties230) GetParsedQueryOk() (*BTPFunctionDeclaration246, bool)`

GetParsedQueryOk returns a tuple with the ParsedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedQuery

`func (o *BTOnePartProperties230) SetParsedQuery(v BTPFunctionDeclaration246)`

SetParsedQuery sets ParsedQuery field to given value.

### HasParsedQuery

`func (o *BTOnePartProperties230) HasParsedQuery() bool`

HasParsedQuery returns a boolean if a field has been set.

### GetQuery

`func (o *BTOnePartProperties230) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BTOnePartProperties230) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BTOnePartProperties230) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *BTOnePartProperties230) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryListParameter

`func (o *BTOnePartProperties230) GetQueryListParameter() BTMParameterQueryList148`

GetQueryListParameter returns the QueryListParameter field if non-nil, zero value otherwise.

### GetQueryListParameterOk

`func (o *BTOnePartProperties230) GetQueryListParameterOk() (*BTMParameterQueryList148, bool)`

GetQueryListParameterOk returns a tuple with the QueryListParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryListParameter

`func (o *BTOnePartProperties230) SetQueryListParameter(v BTMParameterQueryList148)`

SetQueryListParameter sets QueryListParameter field to given value.

### HasQueryListParameter

`func (o *BTOnePartProperties230) HasQueryListParameter() bool`

HasQueryListParameter returns a boolean if a field has been set.

### GetSheetMetalBendOrder

`func (o *BTOnePartProperties230) GetSheetMetalBendOrder() []string`

GetSheetMetalBendOrder returns the SheetMetalBendOrder field if non-nil, zero value otherwise.

### GetSheetMetalBendOrderOk

`func (o *BTOnePartProperties230) GetSheetMetalBendOrderOk() (*[]string, bool)`

GetSheetMetalBendOrderOk returns a tuple with the SheetMetalBendOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetMetalBendOrder

`func (o *BTOnePartProperties230) SetSheetMetalBendOrder(v []string)`

SetSheetMetalBendOrder sets SheetMetalBendOrder field to given value.

### HasSheetMetalBendOrder

`func (o *BTOnePartProperties230) HasSheetMetalBendOrder() bool`

HasSheetMetalBendOrder returns a boolean if a field has been set.

### GetSheetMetalBendOrderIfNotNull

`func (o *BTOnePartProperties230) GetSheetMetalBendOrderIfNotNull() BTOnePartProperties230`

GetSheetMetalBendOrderIfNotNull returns the SheetMetalBendOrderIfNotNull field if non-nil, zero value otherwise.

### GetSheetMetalBendOrderIfNotNullOk

`func (o *BTOnePartProperties230) GetSheetMetalBendOrderIfNotNullOk() (*BTOnePartProperties230, bool)`

GetSheetMetalBendOrderIfNotNullOk returns a tuple with the SheetMetalBendOrderIfNotNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetMetalBendOrderIfNotNull

`func (o *BTOnePartProperties230) SetSheetMetalBendOrderIfNotNull(v BTOnePartProperties230)`

SetSheetMetalBendOrderIfNotNull sets SheetMetalBendOrderIfNotNull field to given value.

### HasSheetMetalBendOrderIfNotNull

`func (o *BTOnePartProperties230) HasSheetMetalBendOrderIfNotNull() bool`

HasSheetMetalBendOrderIfNotNull returns a boolean if a field has been set.

### GetVisibility

`func (o *BTOnePartProperties230) GetVisibility() GBTPartVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BTOnePartProperties230) GetVisibilityOk() (*GBTPartVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BTOnePartProperties230) SetVisibility(v GBTPartVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BTOnePartProperties230) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


