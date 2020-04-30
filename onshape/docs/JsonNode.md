# JsonNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Array** | Pointer to **bool** |  | [optional] 
**BigDecimal** | Pointer to **bool** |  | [optional] 
**BigInteger** | Pointer to **bool** |  | [optional] 
**Binary** | Pointer to **bool** |  | [optional] 
**Boolean** | Pointer to **bool** |  | [optional] 
**ContainerNode** | Pointer to **bool** |  | [optional] 
**Double** | Pointer to **bool** |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 
**Float** | Pointer to **bool** |  | [optional] 
**FloatingPointNumber** | Pointer to **bool** |  | [optional] 
**Int** | Pointer to **bool** |  | [optional] 
**IntegralNumber** | Pointer to **bool** |  | [optional] 
**Long** | Pointer to **bool** |  | [optional] 
**MissingNode** | Pointer to **bool** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**Null** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **bool** |  | [optional] 
**Pojo** | Pointer to **bool** |  | [optional] 
**Short** | Pointer to **bool** |  | [optional] 
**Textual** | Pointer to **bool** |  | [optional] 
**ValueNode** | Pointer to **bool** |  | [optional] 

## Methods

### NewJsonNode

`func NewJsonNode() *JsonNode`

NewJsonNode instantiates a new JsonNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonNodeWithDefaults

`func NewJsonNodeWithDefaults() *JsonNode`

NewJsonNodeWithDefaults instantiates a new JsonNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArray

`func (o *JsonNode) GetArray() bool`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *JsonNode) GetArrayOk() (*bool, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *JsonNode) SetArray(v bool)`

SetArray sets Array field to given value.

### HasArray

`func (o *JsonNode) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetBigDecimal

`func (o *JsonNode) GetBigDecimal() bool`

GetBigDecimal returns the BigDecimal field if non-nil, zero value otherwise.

### GetBigDecimalOk

`func (o *JsonNode) GetBigDecimalOk() (*bool, bool)`

GetBigDecimalOk returns a tuple with the BigDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigDecimal

`func (o *JsonNode) SetBigDecimal(v bool)`

SetBigDecimal sets BigDecimal field to given value.

### HasBigDecimal

`func (o *JsonNode) HasBigDecimal() bool`

HasBigDecimal returns a boolean if a field has been set.

### GetBigInteger

`func (o *JsonNode) GetBigInteger() bool`

GetBigInteger returns the BigInteger field if non-nil, zero value otherwise.

### GetBigIntegerOk

`func (o *JsonNode) GetBigIntegerOk() (*bool, bool)`

GetBigIntegerOk returns a tuple with the BigInteger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigInteger

`func (o *JsonNode) SetBigInteger(v bool)`

SetBigInteger sets BigInteger field to given value.

### HasBigInteger

`func (o *JsonNode) HasBigInteger() bool`

HasBigInteger returns a boolean if a field has been set.

### GetBinary

`func (o *JsonNode) GetBinary() bool`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *JsonNode) GetBinaryOk() (*bool, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *JsonNode) SetBinary(v bool)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *JsonNode) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetBoolean

`func (o *JsonNode) GetBoolean() bool`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *JsonNode) GetBooleanOk() (*bool, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolean

`func (o *JsonNode) SetBoolean(v bool)`

SetBoolean sets Boolean field to given value.

### HasBoolean

`func (o *JsonNode) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### GetContainerNode

`func (o *JsonNode) GetContainerNode() bool`

GetContainerNode returns the ContainerNode field if non-nil, zero value otherwise.

### GetContainerNodeOk

`func (o *JsonNode) GetContainerNodeOk() (*bool, bool)`

GetContainerNodeOk returns a tuple with the ContainerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNode

`func (o *JsonNode) SetContainerNode(v bool)`

SetContainerNode sets ContainerNode field to given value.

### HasContainerNode

`func (o *JsonNode) HasContainerNode() bool`

HasContainerNode returns a boolean if a field has been set.

### GetDouble

`func (o *JsonNode) GetDouble() bool`

GetDouble returns the Double field if non-nil, zero value otherwise.

### GetDoubleOk

`func (o *JsonNode) GetDoubleOk() (*bool, bool)`

GetDoubleOk returns a tuple with the Double field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDouble

`func (o *JsonNode) SetDouble(v bool)`

SetDouble sets Double field to given value.

### HasDouble

`func (o *JsonNode) HasDouble() bool`

HasDouble returns a boolean if a field has been set.

### GetEmpty

`func (o *JsonNode) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *JsonNode) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *JsonNode) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *JsonNode) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetFloat

`func (o *JsonNode) GetFloat() bool`

GetFloat returns the Float field if non-nil, zero value otherwise.

### GetFloatOk

`func (o *JsonNode) GetFloatOk() (*bool, bool)`

GetFloatOk returns a tuple with the Float field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloat

`func (o *JsonNode) SetFloat(v bool)`

SetFloat sets Float field to given value.

### HasFloat

`func (o *JsonNode) HasFloat() bool`

HasFloat returns a boolean if a field has been set.

### GetFloatingPointNumber

`func (o *JsonNode) GetFloatingPointNumber() bool`

GetFloatingPointNumber returns the FloatingPointNumber field if non-nil, zero value otherwise.

### GetFloatingPointNumberOk

`func (o *JsonNode) GetFloatingPointNumberOk() (*bool, bool)`

GetFloatingPointNumberOk returns a tuple with the FloatingPointNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingPointNumber

`func (o *JsonNode) SetFloatingPointNumber(v bool)`

SetFloatingPointNumber sets FloatingPointNumber field to given value.

### HasFloatingPointNumber

`func (o *JsonNode) HasFloatingPointNumber() bool`

HasFloatingPointNumber returns a boolean if a field has been set.

### GetInt

`func (o *JsonNode) GetInt() bool`

GetInt returns the Int field if non-nil, zero value otherwise.

### GetIntOk

`func (o *JsonNode) GetIntOk() (*bool, bool)`

GetIntOk returns a tuple with the Int field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInt

`func (o *JsonNode) SetInt(v bool)`

SetInt sets Int field to given value.

### HasInt

`func (o *JsonNode) HasInt() bool`

HasInt returns a boolean if a field has been set.

### GetIntegralNumber

`func (o *JsonNode) GetIntegralNumber() bool`

GetIntegralNumber returns the IntegralNumber field if non-nil, zero value otherwise.

### GetIntegralNumberOk

`func (o *JsonNode) GetIntegralNumberOk() (*bool, bool)`

GetIntegralNumberOk returns a tuple with the IntegralNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegralNumber

`func (o *JsonNode) SetIntegralNumber(v bool)`

SetIntegralNumber sets IntegralNumber field to given value.

### HasIntegralNumber

`func (o *JsonNode) HasIntegralNumber() bool`

HasIntegralNumber returns a boolean if a field has been set.

### GetLong

`func (o *JsonNode) GetLong() bool`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *JsonNode) GetLongOk() (*bool, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *JsonNode) SetLong(v bool)`

SetLong sets Long field to given value.

### HasLong

`func (o *JsonNode) HasLong() bool`

HasLong returns a boolean if a field has been set.

### GetMissingNode

`func (o *JsonNode) GetMissingNode() bool`

GetMissingNode returns the MissingNode field if non-nil, zero value otherwise.

### GetMissingNodeOk

`func (o *JsonNode) GetMissingNodeOk() (*bool, bool)`

GetMissingNodeOk returns a tuple with the MissingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingNode

`func (o *JsonNode) SetMissingNode(v bool)`

SetMissingNode sets MissingNode field to given value.

### HasMissingNode

`func (o *JsonNode) HasMissingNode() bool`

HasMissingNode returns a boolean if a field has been set.

### GetNodeType

`func (o *JsonNode) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *JsonNode) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *JsonNode) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *JsonNode) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetNull

`func (o *JsonNode) GetNull() bool`

GetNull returns the Null field if non-nil, zero value otherwise.

### GetNullOk

`func (o *JsonNode) GetNullOk() (*bool, bool)`

GetNullOk returns a tuple with the Null field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNull

`func (o *JsonNode) SetNull(v bool)`

SetNull sets Null field to given value.

### HasNull

`func (o *JsonNode) HasNull() bool`

HasNull returns a boolean if a field has been set.

### GetNumber

`func (o *JsonNode) GetNumber() bool`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *JsonNode) GetNumberOk() (*bool, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *JsonNode) SetNumber(v bool)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *JsonNode) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetObject

`func (o *JsonNode) GetObject() bool`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *JsonNode) GetObjectOk() (*bool, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *JsonNode) SetObject(v bool)`

SetObject sets Object field to given value.

### HasObject

`func (o *JsonNode) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPojo

`func (o *JsonNode) GetPojo() bool`

GetPojo returns the Pojo field if non-nil, zero value otherwise.

### GetPojoOk

`func (o *JsonNode) GetPojoOk() (*bool, bool)`

GetPojoOk returns a tuple with the Pojo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPojo

`func (o *JsonNode) SetPojo(v bool)`

SetPojo sets Pojo field to given value.

### HasPojo

`func (o *JsonNode) HasPojo() bool`

HasPojo returns a boolean if a field has been set.

### GetShort

`func (o *JsonNode) GetShort() bool`

GetShort returns the Short field if non-nil, zero value otherwise.

### GetShortOk

`func (o *JsonNode) GetShortOk() (*bool, bool)`

GetShortOk returns a tuple with the Short field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShort

`func (o *JsonNode) SetShort(v bool)`

SetShort sets Short field to given value.

### HasShort

`func (o *JsonNode) HasShort() bool`

HasShort returns a boolean if a field has been set.

### GetTextual

`func (o *JsonNode) GetTextual() bool`

GetTextual returns the Textual field if non-nil, zero value otherwise.

### GetTextualOk

`func (o *JsonNode) GetTextualOk() (*bool, bool)`

GetTextualOk returns a tuple with the Textual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextual

`func (o *JsonNode) SetTextual(v bool)`

SetTextual sets Textual field to given value.

### HasTextual

`func (o *JsonNode) HasTextual() bool`

HasTextual returns a boolean if a field has been set.

### GetValueNode

`func (o *JsonNode) GetValueNode() bool`

GetValueNode returns the ValueNode field if non-nil, zero value otherwise.

### GetValueNodeOk

`func (o *JsonNode) GetValueNodeOk() (*bool, bool)`

GetValueNodeOk returns a tuple with the ValueNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNode

`func (o *JsonNode) SetValueNode(v bool)`

SetValueNode sets ValueNode field to given value.

### HasValueNode

`func (o *JsonNode) HasValueNode() bool`

HasValueNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


