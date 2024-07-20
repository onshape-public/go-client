# BTMIndividualQuery138

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**DeterministicIdList** | Pointer to [**BTMIndividualQueryBase139**](BTMIndividualQueryBase139.md) |  | [optional] 
**DeterministicIds** | Pointer to **[]string** |  | [optional] 
**GenerateSectionEntityQuery** | Pointer to **bool** |  | [optional] 
**GeneratedSectionQueryId** | Pointer to **string** |  | [optional] 
**ImportMicroversion** | Pointer to **string** | Microversion that resulted from the import. | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**PersistentQuery** | Pointer to [**BTPStatement269**](BTPStatement269.md) |  | [optional] 
**Query** | Pointer to [**BTMIndividualQueryBase139**](BTMIndividualQueryBase139.md) |  | [optional] 
**QueryStatement** | Pointer to [**BTPStatement269**](BTPStatement269.md) |  | [optional] 
**QueryString** | Pointer to **string** |  | [optional] 
**VariableName** | Pointer to [**BTMIndividualQuery138**](BTMIndividualQuery138.md) |  | [optional] 

## Methods

### NewBTMIndividualQuery138

`func NewBTMIndividualQuery138() *BTMIndividualQuery138`

NewBTMIndividualQuery138 instantiates a new BTMIndividualQuery138 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMIndividualQuery138WithDefaults

`func NewBTMIndividualQuery138WithDefaults() *BTMIndividualQuery138`

NewBTMIndividualQuery138WithDefaults instantiates a new BTMIndividualQuery138 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTMIndividualQuery138) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMIndividualQuery138) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMIndividualQuery138) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMIndividualQuery138) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDeterministicIdList

`func (o *BTMIndividualQuery138) GetDeterministicIdList() BTMIndividualQueryBase139`

GetDeterministicIdList returns the DeterministicIdList field if non-nil, zero value otherwise.

### GetDeterministicIdListOk

`func (o *BTMIndividualQuery138) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool)`

GetDeterministicIdListOk returns a tuple with the DeterministicIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicIdList

`func (o *BTMIndividualQuery138) SetDeterministicIdList(v BTMIndividualQueryBase139)`

SetDeterministicIdList sets DeterministicIdList field to given value.

### HasDeterministicIdList

`func (o *BTMIndividualQuery138) HasDeterministicIdList() bool`

HasDeterministicIdList returns a boolean if a field has been set.

### GetDeterministicIds

`func (o *BTMIndividualQuery138) GetDeterministicIds() []string`

GetDeterministicIds returns the DeterministicIds field if non-nil, zero value otherwise.

### GetDeterministicIdsOk

`func (o *BTMIndividualQuery138) GetDeterministicIdsOk() (*[]string, bool)`

GetDeterministicIdsOk returns a tuple with the DeterministicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicIds

`func (o *BTMIndividualQuery138) SetDeterministicIds(v []string)`

SetDeterministicIds sets DeterministicIds field to given value.

### HasDeterministicIds

`func (o *BTMIndividualQuery138) HasDeterministicIds() bool`

HasDeterministicIds returns a boolean if a field has been set.

### GetGenerateSectionEntityQuery

`func (o *BTMIndividualQuery138) GetGenerateSectionEntityQuery() bool`

GetGenerateSectionEntityQuery returns the GenerateSectionEntityQuery field if non-nil, zero value otherwise.

### GetGenerateSectionEntityQueryOk

`func (o *BTMIndividualQuery138) GetGenerateSectionEntityQueryOk() (*bool, bool)`

GetGenerateSectionEntityQueryOk returns a tuple with the GenerateSectionEntityQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateSectionEntityQuery

`func (o *BTMIndividualQuery138) SetGenerateSectionEntityQuery(v bool)`

SetGenerateSectionEntityQuery sets GenerateSectionEntityQuery field to given value.

### HasGenerateSectionEntityQuery

`func (o *BTMIndividualQuery138) HasGenerateSectionEntityQuery() bool`

HasGenerateSectionEntityQuery returns a boolean if a field has been set.

### GetGeneratedSectionQueryId

`func (o *BTMIndividualQuery138) GetGeneratedSectionQueryId() string`

GetGeneratedSectionQueryId returns the GeneratedSectionQueryId field if non-nil, zero value otherwise.

### GetGeneratedSectionQueryIdOk

`func (o *BTMIndividualQuery138) GetGeneratedSectionQueryIdOk() (*string, bool)`

GetGeneratedSectionQueryIdOk returns a tuple with the GeneratedSectionQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedSectionQueryId

`func (o *BTMIndividualQuery138) SetGeneratedSectionQueryId(v string)`

SetGeneratedSectionQueryId sets GeneratedSectionQueryId field to given value.

### HasGeneratedSectionQueryId

`func (o *BTMIndividualQuery138) HasGeneratedSectionQueryId() bool`

HasGeneratedSectionQueryId returns a boolean if a field has been set.

### GetImportMicroversion

`func (o *BTMIndividualQuery138) GetImportMicroversion() string`

GetImportMicroversion returns the ImportMicroversion field if non-nil, zero value otherwise.

### GetImportMicroversionOk

`func (o *BTMIndividualQuery138) GetImportMicroversionOk() (*string, bool)`

GetImportMicroversionOk returns a tuple with the ImportMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMicroversion

`func (o *BTMIndividualQuery138) SetImportMicroversion(v string)`

SetImportMicroversion sets ImportMicroversion field to given value.

### HasImportMicroversion

`func (o *BTMIndividualQuery138) HasImportMicroversion() bool`

HasImportMicroversion returns a boolean if a field has been set.

### GetNodeId

`func (o *BTMIndividualQuery138) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTMIndividualQuery138) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTMIndividualQuery138) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTMIndividualQuery138) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetPersistentQuery

`func (o *BTMIndividualQuery138) GetPersistentQuery() BTPStatement269`

GetPersistentQuery returns the PersistentQuery field if non-nil, zero value otherwise.

### GetPersistentQueryOk

`func (o *BTMIndividualQuery138) GetPersistentQueryOk() (*BTPStatement269, bool)`

GetPersistentQueryOk returns a tuple with the PersistentQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentQuery

`func (o *BTMIndividualQuery138) SetPersistentQuery(v BTPStatement269)`

SetPersistentQuery sets PersistentQuery field to given value.

### HasPersistentQuery

`func (o *BTMIndividualQuery138) HasPersistentQuery() bool`

HasPersistentQuery returns a boolean if a field has been set.

### GetQuery

`func (o *BTMIndividualQuery138) GetQuery() BTMIndividualQueryBase139`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *BTMIndividualQuery138) GetQueryOk() (*BTMIndividualQueryBase139, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *BTMIndividualQuery138) SetQuery(v BTMIndividualQueryBase139)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *BTMIndividualQuery138) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryStatement

`func (o *BTMIndividualQuery138) GetQueryStatement() BTPStatement269`

GetQueryStatement returns the QueryStatement field if non-nil, zero value otherwise.

### GetQueryStatementOk

`func (o *BTMIndividualQuery138) GetQueryStatementOk() (*BTPStatement269, bool)`

GetQueryStatementOk returns a tuple with the QueryStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatement

`func (o *BTMIndividualQuery138) SetQueryStatement(v BTPStatement269)`

SetQueryStatement sets QueryStatement field to given value.

### HasQueryStatement

`func (o *BTMIndividualQuery138) HasQueryStatement() bool`

HasQueryStatement returns a boolean if a field has been set.

### GetQueryString

`func (o *BTMIndividualQuery138) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *BTMIndividualQuery138) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *BTMIndividualQuery138) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *BTMIndividualQuery138) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetVariableName

`func (o *BTMIndividualQuery138) GetVariableName() BTMIndividualQuery138`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *BTMIndividualQuery138) GetVariableNameOk() (*BTMIndividualQuery138, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *BTMIndividualQuery138) SetVariableName(v BTMIndividualQuery138)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *BTMIndividualQuery138) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


