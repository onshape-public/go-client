/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
	"fmt"
)

// BTTableRow1054 - struct for BTTableRow1054
type BTTableRow1054 struct {
	implBTTableRow1054 interface{}
}

// BTBillOfMaterialsTableRow1425AsBTTableRow1054 is a convenience function that returns BTBillOfMaterialsTableRow1425 wrapped in BTTableRow1054
func (o *BTBillOfMaterialsTableRow1425) AsBTTableRow1054() *BTTableRow1054 {
	return &BTTableRow1054{o}
}

// NewBTTableRow1054 instantiates a new BTTableRow1054 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableRow1054() *BTTableRow1054 {
	this := BTTableRow1054{Newbase_BTTableRow1054()}
	return &this
}

// NewBTTableRow1054WithDefaults instantiates a new BTTableRow1054 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableRow1054WithDefaults() *BTTableRow1054 {
	this := BTTableRow1054{Newbase_BTTableRow1054WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableRow1054) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableRow1054) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableRow1054) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableRow1054) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetColumnIdToCell returns the ColumnIdToCell field value if set, zero value otherwise.
func (o *BTTableRow1054) GetColumnIdToCell() map[string]BTTableCell1114 {
	type getResult interface {
		GetColumnIdToCell() map[string]BTTableCell1114
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetColumnIdToCell()
	} else {
		var de map[string]BTTableCell1114
		return de
	}
}

// GetColumnIdToCellOk returns a tuple with the ColumnIdToCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableRow1054) GetColumnIdToCellOk() (*map[string]BTTableCell1114, bool) {
	type getResult interface {
		GetColumnIdToCellOk() (*map[string]BTTableCell1114, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetColumnIdToCellOk()
	} else {
		return nil, false
	}
}

// HasColumnIdToCell returns a boolean if a field has been set.
func (o *BTTableRow1054) HasColumnIdToCell() bool {
	type getResult interface {
		HasColumnIdToCell() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasColumnIdToCell()
	} else {
		return false
	}
}

// SetColumnIdToCell gets a reference to the given map[string]BTTableCell1114 and assigns it to the ColumnIdToCell field.
func (o *BTTableRow1054) SetColumnIdToCell(v map[string]BTTableCell1114) {
	type getResult interface {
		SetColumnIdToCell(v map[string]BTTableCell1114)
	}

	o.GetActualInstance().(getResult).SetColumnIdToCell(v)
}

// GetExpandableState returns the ExpandableState field value if set, zero value otherwise.
func (o *BTTableRow1054) GetExpandableState() ExpandableState {
	type getResult interface {
		GetExpandableState() ExpandableState
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExpandableState()
	} else {
		var de ExpandableState
		return de
	}
}

// GetExpandableStateOk returns a tuple with the ExpandableState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableRow1054) GetExpandableStateOk() (*ExpandableState, bool) {
	type getResult interface {
		GetExpandableStateOk() (*ExpandableState, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetExpandableStateOk()
	} else {
		return nil, false
	}
}

// HasExpandableState returns a boolean if a field has been set.
func (o *BTTableRow1054) HasExpandableState() bool {
	type getResult interface {
		HasExpandableState() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasExpandableState()
	} else {
		return false
	}
}

// SetExpandableState gets a reference to the given ExpandableState and assigns it to the ExpandableState field.
func (o *BTTableRow1054) SetExpandableState(v ExpandableState) {
	type getResult interface {
		SetExpandableState(v ExpandableState)
	}

	o.GetActualInstance().(getResult).SetExpandableState(v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTTableRow1054) GetId() string {
	type getResult interface {
		GetId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetId()
	} else {
		var de string
		return de
	}
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableRow1054) GetIdOk() (*string, bool) {
	type getResult interface {
		GetIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetIdOk()
	} else {
		return nil, false
	}
}

// HasId returns a boolean if a field has been set.
func (o *BTTableRow1054) HasId() bool {
	type getResult interface {
		HasId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasId()
	} else {
		return false
	}
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTTableRow1054) SetId(v string) {
	type getResult interface {
		SetId(v string)
	}

	o.GetActualInstance().(getResult).SetId(v)
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *BTTableRow1054) GetMetaData() BTTreeNode20 {
	type getResult interface {
		GetMetaData() BTTreeNode20
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMetaData()
	} else {
		var de BTTreeNode20
		return de
	}
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableRow1054) GetMetaDataOk() (*BTTreeNode20, bool) {
	type getResult interface {
		GetMetaDataOk() (*BTTreeNode20, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetMetaDataOk()
	} else {
		return nil, false
	}
}

// HasMetaData returns a boolean if a field has been set.
func (o *BTTableRow1054) HasMetaData() bool {
	type getResult interface {
		HasMetaData() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasMetaData()
	} else {
		return false
	}
}

// SetMetaData gets a reference to the given BTTreeNode20 and assigns it to the MetaData field.
func (o *BTTableRow1054) SetMetaData(v BTTreeNode20) {
	type getResult interface {
		SetMetaData(v BTTreeNode20)
	}

	o.GetActualInstance().(getResult).SetMetaData(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTTableRow1054) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableRow1054) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTTableRow1054) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTTableRow1054) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetRowMetadata returns the RowMetadata field value if set, zero value otherwise.
func (o *BTTableRow1054) GetRowMetadata() BTTableBaseRowMetadata3181 {
	type getResult interface {
		GetRowMetadata() BTTableBaseRowMetadata3181
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRowMetadata()
	} else {
		var de BTTableBaseRowMetadata3181
		return de
	}
}

// GetRowMetadataOk returns a tuple with the RowMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableRow1054) GetRowMetadataOk() (*BTTableBaseRowMetadata3181, bool) {
	type getResult interface {
		GetRowMetadataOk() (*BTTableBaseRowMetadata3181, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetRowMetadataOk()
	} else {
		return nil, false
	}
}

// HasRowMetadata returns a boolean if a field has been set.
func (o *BTTableRow1054) HasRowMetadata() bool {
	type getResult interface {
		HasRowMetadata() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasRowMetadata()
	} else {
		return false
	}
}

// SetRowMetadata gets a reference to the given BTTableBaseRowMetadata3181 and assigns it to the RowMetadata field.
func (o *BTTableRow1054) SetRowMetadata(v BTTableBaseRowMetadata3181) {
	type getResult interface {
		SetRowMetadata(v BTTableBaseRowMetadata3181)
	}

	o.GetActualInstance().(getResult).SetRowMetadata(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTTableRow1054) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'BTBillOfMaterialsTableRow-1425'
	if jsonDict["btType"] == "BTBillOfMaterialsTableRow-1425" {
		// try to unmarshal JSON data into BTBillOfMaterialsTableRow1425
		var qr *BTBillOfMaterialsTableRow1425
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTTableRow1054 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTTableRow1054 = nil
			return fmt.Errorf("failed to unmarshal BTTableRow1054 as BTBillOfMaterialsTableRow1425: %s", err.Error())
		}
	}

	var qtx *base_BTTableRow1054
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTTableRow1054 = qtx
		return nil // data stored in dst.base_BTTableRow1054, return on the first match
	} else {
		dst.implBTTableRow1054 = nil
		return fmt.Errorf("failed to unmarshal BTTableRow1054 as base_BTTableRow1054: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTTableRow1054) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTTableRow1054) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTTableRow1054
}

type NullableBTTableRow1054 struct {
	value *BTTableRow1054
	isSet bool
}

func (v NullableBTTableRow1054) Get() *BTTableRow1054 {
	return v.value
}

func (v *NullableBTTableRow1054) Set(val *BTTableRow1054) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableRow1054) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableRow1054) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableRow1054(val *BTTableRow1054) *NullableBTTableRow1054 {
	return &NullableBTTableRow1054{value: val, isSet: true}
}

func (v NullableBTTableRow1054) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableRow1054) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTTableRow1054 struct {
	// Type of JSON object.
	BtType          *string                     `json:"btType,omitempty"`
	ColumnIdToCell  *map[string]BTTableCell1114 `json:"columnIdToCell,omitempty"`
	ExpandableState *ExpandableState            `json:"expandableState,omitempty"`
	Id              *string                     `json:"id,omitempty"`
	MetaData        *BTTreeNode20               `json:"metaData,omitempty"`
	NodeId          *string                     `json:"nodeId,omitempty"`
	RowMetadata     *BTTableBaseRowMetadata3181 `json:"rowMetadata,omitempty"`
}

// Newbase_BTTableRow1054 instantiates a new base_BTTableRow1054 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTTableRow1054() *base_BTTableRow1054 {
	this := base_BTTableRow1054{}
	return &this
}

// Newbase_BTTableRow1054WithDefaults instantiates a new base_BTTableRow1054 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTTableRow1054WithDefaults() *base_BTTableRow1054 {
	this := base_BTTableRow1054{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTTableRow1054) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableRow1054) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTTableRow1054) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTTableRow1054) SetBtType(v string) {
	o.BtType = &v
}

// GetColumnIdToCell returns the ColumnIdToCell field value if set, zero value otherwise.
func (o *base_BTTableRow1054) GetColumnIdToCell() map[string]BTTableCell1114 {
	if o == nil || o.ColumnIdToCell == nil {
		var ret map[string]BTTableCell1114
		return ret
	}
	return *o.ColumnIdToCell
}

// GetColumnIdToCellOk returns a tuple with the ColumnIdToCell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableRow1054) GetColumnIdToCellOk() (*map[string]BTTableCell1114, bool) {
	if o == nil || o.ColumnIdToCell == nil {
		return nil, false
	}
	return o.ColumnIdToCell, true
}

// HasColumnIdToCell returns a boolean if a field has been set.
func (o *base_BTTableRow1054) HasColumnIdToCell() bool {
	if o != nil && o.ColumnIdToCell != nil {
		return true
	}

	return false
}

// SetColumnIdToCell gets a reference to the given map[string]BTTableCell1114 and assigns it to the ColumnIdToCell field.
func (o *base_BTTableRow1054) SetColumnIdToCell(v map[string]BTTableCell1114) {
	o.ColumnIdToCell = &v
}

// GetExpandableState returns the ExpandableState field value if set, zero value otherwise.
func (o *base_BTTableRow1054) GetExpandableState() ExpandableState {
	if o == nil || o.ExpandableState == nil {
		var ret ExpandableState
		return ret
	}
	return *o.ExpandableState
}

// GetExpandableStateOk returns a tuple with the ExpandableState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableRow1054) GetExpandableStateOk() (*ExpandableState, bool) {
	if o == nil || o.ExpandableState == nil {
		return nil, false
	}
	return o.ExpandableState, true
}

// HasExpandableState returns a boolean if a field has been set.
func (o *base_BTTableRow1054) HasExpandableState() bool {
	if o != nil && o.ExpandableState != nil {
		return true
	}

	return false
}

// SetExpandableState gets a reference to the given ExpandableState and assigns it to the ExpandableState field.
func (o *base_BTTableRow1054) SetExpandableState(v ExpandableState) {
	o.ExpandableState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *base_BTTableRow1054) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableRow1054) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *base_BTTableRow1054) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *base_BTTableRow1054) SetId(v string) {
	o.Id = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *base_BTTableRow1054) GetMetaData() BTTreeNode20 {
	if o == nil || o.MetaData == nil {
		var ret BTTreeNode20
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableRow1054) GetMetaDataOk() (*BTTreeNode20, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *base_BTTableRow1054) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given BTTreeNode20 and assigns it to the MetaData field.
func (o *base_BTTableRow1054) SetMetaData(v BTTreeNode20) {
	o.MetaData = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTTableRow1054) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableRow1054) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTTableRow1054) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTTableRow1054) SetNodeId(v string) {
	o.NodeId = &v
}

// GetRowMetadata returns the RowMetadata field value if set, zero value otherwise.
func (o *base_BTTableRow1054) GetRowMetadata() BTTableBaseRowMetadata3181 {
	if o == nil || o.RowMetadata == nil {
		var ret BTTableBaseRowMetadata3181
		return ret
	}
	return *o.RowMetadata
}

// GetRowMetadataOk returns a tuple with the RowMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTTableRow1054) GetRowMetadataOk() (*BTTableBaseRowMetadata3181, bool) {
	if o == nil || o.RowMetadata == nil {
		return nil, false
	}
	return o.RowMetadata, true
}

// HasRowMetadata returns a boolean if a field has been set.
func (o *base_BTTableRow1054) HasRowMetadata() bool {
	if o != nil && o.RowMetadata != nil {
		return true
	}

	return false
}

// SetRowMetadata gets a reference to the given BTTableBaseRowMetadata3181 and assigns it to the RowMetadata field.
func (o *base_BTTableRow1054) SetRowMetadata(v BTTableBaseRowMetadata3181) {
	o.RowMetadata = &v
}

func (o base_BTTableRow1054) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ColumnIdToCell != nil {
		toSerialize["columnIdToCell"] = o.ColumnIdToCell
	}
	if o.ExpandableState != nil {
		toSerialize["expandableState"] = o.ExpandableState
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetaData != nil {
		toSerialize["metaData"] = o.MetaData
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.RowMetadata != nil {
		toSerialize["rowMetadata"] = o.RowMetadata
	}
	return json.Marshal(toSerialize)
}
