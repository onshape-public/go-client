/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://cad.onshape.com/appstore/dev-portal): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTApiTable2300 struct for BTApiTable2300
type BTApiTable2300 struct {
	// Type of JSON object.
	BtType    *string                `json:"btType,omitempty"`
	Columns   []BTApiTableColumn3141 `json:"columns,omitempty"`
	EntityIds []string               `json:"entityIds,omitempty"`
	Id        *string                `json:"id,omitempty"`
	Rows      []BTApiTableRow2915    `json:"rows,omitempty"`
	Title     *string                `json:"title,omitempty"`
}

// NewBTApiTable2300 instantiates a new BTApiTable2300 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApiTable2300() *BTApiTable2300 {
	this := BTApiTable2300{}
	return &this
}

// NewBTApiTable2300WithDefaults instantiates a new BTApiTable2300 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApiTable2300WithDefaults() *BTApiTable2300 {
	this := BTApiTable2300{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTApiTable2300) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTApiTable2300) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTApiTable2300) SetBtType(v string) {
	o.BtType = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *BTApiTable2300) GetColumns() []BTApiTableColumn3141 {
	if o == nil || o.Columns == nil {
		var ret []BTApiTableColumn3141
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetColumnsOk() ([]BTApiTableColumn3141, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *BTApiTable2300) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []BTApiTableColumn3141 and assigns it to the Columns field.
func (o *BTApiTable2300) SetColumns(v []BTApiTableColumn3141) {
	o.Columns = v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *BTApiTable2300) GetEntityIds() []string {
	if o == nil || o.EntityIds == nil {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetEntityIdsOk() ([]string, bool) {
	if o == nil || o.EntityIds == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *BTApiTable2300) HasEntityIds() bool {
	if o != nil && o.EntityIds != nil {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *BTApiTable2300) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTApiTable2300) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTApiTable2300) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTApiTable2300) SetId(v string) {
	o.Id = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *BTApiTable2300) GetRows() []BTApiTableRow2915 {
	if o == nil || o.Rows == nil {
		var ret []BTApiTableRow2915
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetRowsOk() ([]BTApiTableRow2915, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *BTApiTable2300) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []BTApiTableRow2915 and assigns it to the Rows field.
func (o *BTApiTable2300) SetRows(v []BTApiTableRow2915) {
	o.Rows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BTApiTable2300) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BTApiTable2300) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BTApiTable2300) SetTitle(v string) {
	o.Title = &v
}

func (o BTApiTable2300) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.EntityIds != nil {
		toSerialize["entityIds"] = o.EntityIds
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableBTApiTable2300 struct {
	value *BTApiTable2300
	isSet bool
}

func (v NullableBTApiTable2300) Get() *BTApiTable2300 {
	return v.value
}

func (v *NullableBTApiTable2300) Set(val *BTApiTable2300) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApiTable2300) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApiTable2300) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApiTable2300(val *BTApiTable2300) *NullableBTApiTable2300 {
	return &NullableBTApiTable2300{value: val, isSet: true}
}

func (v NullableBTApiTable2300) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApiTable2300) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
