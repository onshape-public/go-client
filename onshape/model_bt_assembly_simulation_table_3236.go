/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.150.5616-564f6a8676f1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblySimulationTable3236 struct for BTAssemblySimulationTable3236
type BTAssemblySimulationTable3236 struct {
	AllRowValues  [][]string              `json:"allRowValues,omitempty"`
	ColumnCount   *int32                  `json:"columnCount,omitempty"`
	FrozenColumns *int32                  `json:"frozenColumns,omitempty"`
	IsFailed      *bool                   `json:"isFailed,omitempty"`
	NodeId        *string                 `json:"nodeId,omitempty"`
	ReadOnly      *bool                   `json:"readOnly,omitempty"`
	RowCount      *int32                  `json:"rowCount,omitempty"`
	TableColumns  []BTTableColumnInfo1222 `json:"tableColumns,omitempty"`
	TableId       *string                 `json:"tableId,omitempty"`
	TableRows     []BTTableRow1054        `json:"tableRows,omitempty"`
	Title         *string                 `json:"title,omitempty"`
	BtType        *string                 `json:"btType,omitempty"`
}

// NewBTAssemblySimulationTable3236 instantiates a new BTAssemblySimulationTable3236 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblySimulationTable3236() *BTAssemblySimulationTable3236 {
	this := BTAssemblySimulationTable3236{}
	return &this
}

// NewBTAssemblySimulationTable3236WithDefaults instantiates a new BTAssemblySimulationTable3236 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblySimulationTable3236WithDefaults() *BTAssemblySimulationTable3236 {
	this := BTAssemblySimulationTable3236{}
	return &this
}

// GetAllRowValues returns the AllRowValues field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetAllRowValues() [][]string {
	if o == nil || o.AllRowValues == nil {
		var ret [][]string
		return ret
	}
	return o.AllRowValues
}

// GetAllRowValuesOk returns a tuple with the AllRowValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetAllRowValuesOk() ([][]string, bool) {
	if o == nil || o.AllRowValues == nil {
		return nil, false
	}
	return o.AllRowValues, true
}

// HasAllRowValues returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasAllRowValues() bool {
	if o != nil && o.AllRowValues != nil {
		return true
	}

	return false
}

// SetAllRowValues gets a reference to the given [][]string and assigns it to the AllRowValues field.
func (o *BTAssemblySimulationTable3236) SetAllRowValues(v [][]string) {
	o.AllRowValues = v
}

// GetColumnCount returns the ColumnCount field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetColumnCount() int32 {
	if o == nil || o.ColumnCount == nil {
		var ret int32
		return ret
	}
	return *o.ColumnCount
}

// GetColumnCountOk returns a tuple with the ColumnCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetColumnCountOk() (*int32, bool) {
	if o == nil || o.ColumnCount == nil {
		return nil, false
	}
	return o.ColumnCount, true
}

// HasColumnCount returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasColumnCount() bool {
	if o != nil && o.ColumnCount != nil {
		return true
	}

	return false
}

// SetColumnCount gets a reference to the given int32 and assigns it to the ColumnCount field.
func (o *BTAssemblySimulationTable3236) SetColumnCount(v int32) {
	o.ColumnCount = &v
}

// GetFrozenColumns returns the FrozenColumns field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetFrozenColumns() int32 {
	if o == nil || o.FrozenColumns == nil {
		var ret int32
		return ret
	}
	return *o.FrozenColumns
}

// GetFrozenColumnsOk returns a tuple with the FrozenColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetFrozenColumnsOk() (*int32, bool) {
	if o == nil || o.FrozenColumns == nil {
		return nil, false
	}
	return o.FrozenColumns, true
}

// HasFrozenColumns returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasFrozenColumns() bool {
	if o != nil && o.FrozenColumns != nil {
		return true
	}

	return false
}

// SetFrozenColumns gets a reference to the given int32 and assigns it to the FrozenColumns field.
func (o *BTAssemblySimulationTable3236) SetFrozenColumns(v int32) {
	o.FrozenColumns = &v
}

// GetIsFailed returns the IsFailed field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetIsFailed() bool {
	if o == nil || o.IsFailed == nil {
		var ret bool
		return ret
	}
	return *o.IsFailed
}

// GetIsFailedOk returns a tuple with the IsFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetIsFailedOk() (*bool, bool) {
	if o == nil || o.IsFailed == nil {
		return nil, false
	}
	return o.IsFailed, true
}

// HasIsFailed returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasIsFailed() bool {
	if o != nil && o.IsFailed != nil {
		return true
	}

	return false
}

// SetIsFailed gets a reference to the given bool and assigns it to the IsFailed field.
func (o *BTAssemblySimulationTable3236) SetIsFailed(v bool) {
	o.IsFailed = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTAssemblySimulationTable3236) SetNodeId(v string) {
	o.NodeId = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BTAssemblySimulationTable3236) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRowCount returns the RowCount field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetRowCount() int32 {
	if o == nil || o.RowCount == nil {
		var ret int32
		return ret
	}
	return *o.RowCount
}

// GetRowCountOk returns a tuple with the RowCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetRowCountOk() (*int32, bool) {
	if o == nil || o.RowCount == nil {
		return nil, false
	}
	return o.RowCount, true
}

// HasRowCount returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasRowCount() bool {
	if o != nil && o.RowCount != nil {
		return true
	}

	return false
}

// SetRowCount gets a reference to the given int32 and assigns it to the RowCount field.
func (o *BTAssemblySimulationTable3236) SetRowCount(v int32) {
	o.RowCount = &v
}

// GetTableColumns returns the TableColumns field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetTableColumns() []BTTableColumnInfo1222 {
	if o == nil || o.TableColumns == nil {
		var ret []BTTableColumnInfo1222
		return ret
	}
	return o.TableColumns
}

// GetTableColumnsOk returns a tuple with the TableColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetTableColumnsOk() ([]BTTableColumnInfo1222, bool) {
	if o == nil || o.TableColumns == nil {
		return nil, false
	}
	return o.TableColumns, true
}

// HasTableColumns returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasTableColumns() bool {
	if o != nil && o.TableColumns != nil {
		return true
	}

	return false
}

// SetTableColumns gets a reference to the given []BTTableColumnInfo1222 and assigns it to the TableColumns field.
func (o *BTAssemblySimulationTable3236) SetTableColumns(v []BTTableColumnInfo1222) {
	o.TableColumns = v
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetTableId() string {
	if o == nil || o.TableId == nil {
		var ret string
		return ret
	}
	return *o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetTableIdOk() (*string, bool) {
	if o == nil || o.TableId == nil {
		return nil, false
	}
	return o.TableId, true
}

// HasTableId returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasTableId() bool {
	if o != nil && o.TableId != nil {
		return true
	}

	return false
}

// SetTableId gets a reference to the given string and assigns it to the TableId field.
func (o *BTAssemblySimulationTable3236) SetTableId(v string) {
	o.TableId = &v
}

// GetTableRows returns the TableRows field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetTableRows() []BTTableRow1054 {
	if o == nil || o.TableRows == nil {
		var ret []BTTableRow1054
		return ret
	}
	return o.TableRows
}

// GetTableRowsOk returns a tuple with the TableRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetTableRowsOk() ([]BTTableRow1054, bool) {
	if o == nil || o.TableRows == nil {
		return nil, false
	}
	return o.TableRows, true
}

// HasTableRows returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasTableRows() bool {
	if o != nil && o.TableRows != nil {
		return true
	}

	return false
}

// SetTableRows gets a reference to the given []BTTableRow1054 and assigns it to the TableRows field.
func (o *BTAssemblySimulationTable3236) SetTableRows(v []BTTableRow1054) {
	o.TableRows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BTAssemblySimulationTable3236) SetTitle(v string) {
	o.Title = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAssemblySimulationTable3236) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblySimulationTable3236) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAssemblySimulationTable3236) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAssemblySimulationTable3236) SetBtType(v string) {
	o.BtType = &v
}

func (o BTAssemblySimulationTable3236) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllRowValues != nil {
		toSerialize["allRowValues"] = o.AllRowValues
	}
	if o.ColumnCount != nil {
		toSerialize["columnCount"] = o.ColumnCount
	}
	if o.FrozenColumns != nil {
		toSerialize["frozenColumns"] = o.FrozenColumns
	}
	if o.IsFailed != nil {
		toSerialize["isFailed"] = o.IsFailed
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.RowCount != nil {
		toSerialize["rowCount"] = o.RowCount
	}
	if o.TableColumns != nil {
		toSerialize["tableColumns"] = o.TableColumns
	}
	if o.TableId != nil {
		toSerialize["tableId"] = o.TableId
	}
	if o.TableRows != nil {
		toSerialize["tableRows"] = o.TableRows
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblySimulationTable3236 struct {
	value *BTAssemblySimulationTable3236
	isSet bool
}

func (v NullableBTAssemblySimulationTable3236) Get() *BTAssemblySimulationTable3236 {
	return v.value
}

func (v *NullableBTAssemblySimulationTable3236) Set(val *BTAssemblySimulationTable3236) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblySimulationTable3236) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblySimulationTable3236) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblySimulationTable3236(val *BTAssemblySimulationTable3236) *NullableBTAssemblySimulationTable3236 {
	return &NullableBTAssemblySimulationTable3236{value: val, isSet: true}
}

func (v NullableBTAssemblySimulationTable3236) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblySimulationTable3236) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}