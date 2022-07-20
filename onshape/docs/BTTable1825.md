# BTTable1825

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllRowValues** | Pointer to **[][]string** |  | [optional] 
**ColumnCount** | Pointer to **int32** |  | [optional] 
**FrozenColumns** | Pointer to **int32** |  | [optional] 
**IsFailed** | Pointer to **bool** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**RowCount** | Pointer to **int32** |  | [optional] 
**TableColumns** | Pointer to [**[]BTTableColumnInfo1222**](BTTableColumnInfo1222.md) |  | [optional] 
**TableId** | Pointer to **string** |  | [optional] 
**TableRows** | Pointer to [**[]BTTableRow1054**](BTTableRow1054.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewBTTable1825

`func NewBTTable1825() *BTTable1825`

NewBTTable1825 instantiates a new BTTable1825 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTable1825WithDefaults

`func NewBTTable1825WithDefaults() *BTTable1825`

NewBTTable1825WithDefaults instantiates a new BTTable1825 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllRowValues

`func (o *BTTable1825) GetAllRowValues() [][]string`

GetAllRowValues returns the AllRowValues field if non-nil, zero value otherwise.

### GetAllRowValuesOk

`func (o *BTTable1825) GetAllRowValuesOk() (*[][]string, bool)`

GetAllRowValuesOk returns a tuple with the AllRowValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRowValues

`func (o *BTTable1825) SetAllRowValues(v [][]string)`

SetAllRowValues sets AllRowValues field to given value.

### HasAllRowValues

`func (o *BTTable1825) HasAllRowValues() bool`

HasAllRowValues returns a boolean if a field has been set.

### GetColumnCount

`func (o *BTTable1825) GetColumnCount() int32`

GetColumnCount returns the ColumnCount field if non-nil, zero value otherwise.

### GetColumnCountOk

`func (o *BTTable1825) GetColumnCountOk() (*int32, bool)`

GetColumnCountOk returns a tuple with the ColumnCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnCount

`func (o *BTTable1825) SetColumnCount(v int32)`

SetColumnCount sets ColumnCount field to given value.

### HasColumnCount

`func (o *BTTable1825) HasColumnCount() bool`

HasColumnCount returns a boolean if a field has been set.

### GetFrozenColumns

`func (o *BTTable1825) GetFrozenColumns() int32`

GetFrozenColumns returns the FrozenColumns field if non-nil, zero value otherwise.

### GetFrozenColumnsOk

`func (o *BTTable1825) GetFrozenColumnsOk() (*int32, bool)`

GetFrozenColumnsOk returns a tuple with the FrozenColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenColumns

`func (o *BTTable1825) SetFrozenColumns(v int32)`

SetFrozenColumns sets FrozenColumns field to given value.

### HasFrozenColumns

`func (o *BTTable1825) HasFrozenColumns() bool`

HasFrozenColumns returns a boolean if a field has been set.

### GetIsFailed

`func (o *BTTable1825) GetIsFailed() bool`

GetIsFailed returns the IsFailed field if non-nil, zero value otherwise.

### GetIsFailedOk

`func (o *BTTable1825) GetIsFailedOk() (*bool, bool)`

GetIsFailedOk returns a tuple with the IsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailed

`func (o *BTTable1825) SetIsFailed(v bool)`

SetIsFailed sets IsFailed field to given value.

### HasIsFailed

`func (o *BTTable1825) HasIsFailed() bool`

HasIsFailed returns a boolean if a field has been set.

### GetNodeId

`func (o *BTTable1825) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTTable1825) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTTable1825) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTTable1825) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetReadOnly

`func (o *BTTable1825) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *BTTable1825) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *BTTable1825) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *BTTable1825) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRowCount

`func (o *BTTable1825) GetRowCount() int32`

GetRowCount returns the RowCount field if non-nil, zero value otherwise.

### GetRowCountOk

`func (o *BTTable1825) GetRowCountOk() (*int32, bool)`

GetRowCountOk returns a tuple with the RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowCount

`func (o *BTTable1825) SetRowCount(v int32)`

SetRowCount sets RowCount field to given value.

### HasRowCount

`func (o *BTTable1825) HasRowCount() bool`

HasRowCount returns a boolean if a field has been set.

### GetTableColumns

`func (o *BTTable1825) GetTableColumns() []BTTableColumnInfo1222`

GetTableColumns returns the TableColumns field if non-nil, zero value otherwise.

### GetTableColumnsOk

`func (o *BTTable1825) GetTableColumnsOk() (*[]BTTableColumnInfo1222, bool)`

GetTableColumnsOk returns a tuple with the TableColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableColumns

`func (o *BTTable1825) SetTableColumns(v []BTTableColumnInfo1222)`

SetTableColumns sets TableColumns field to given value.

### HasTableColumns

`func (o *BTTable1825) HasTableColumns() bool`

HasTableColumns returns a boolean if a field has been set.

### GetTableId

`func (o *BTTable1825) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *BTTable1825) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *BTTable1825) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *BTTable1825) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetTableRows

`func (o *BTTable1825) GetTableRows() []BTTableRow1054`

GetTableRows returns the TableRows field if non-nil, zero value otherwise.

### GetTableRowsOk

`func (o *BTTable1825) GetTableRowsOk() (*[]BTTableRow1054, bool)`

GetTableRowsOk returns a tuple with the TableRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableRows

`func (o *BTTable1825) SetTableRows(v []BTTableRow1054)`

SetTableRows sets TableRows field to given value.

### HasTableRows

`func (o *BTTable1825) HasTableRows() bool`

HasTableRows returns a boolean if a field has been set.

### GetTitle

`func (o *BTTable1825) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BTTable1825) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BTTable1825) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BTTable1825) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


