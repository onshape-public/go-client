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

// BTExportTessellatedEdgesBody890 struct for BTExportTessellatedEdgesBody890
type BTExportTessellatedEdgesBody890 struct {
	BTExportTessellatedBody3398
	BtType       *string                            `json:"btType,omitempty"`
	Constituents []string                           `json:"constituents,omitempty"`
	Id           *string                            `json:"id,omitempty"`
	Name         *string                            `json:"name,omitempty"`
	Edges        []BTExportTessellatedEdgesEdge1364 `json:"edges,omitempty"`
}

// NewBTExportTessellatedEdgesBody890 instantiates a new BTExportTessellatedEdgesBody890 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedEdgesBody890() *BTExportTessellatedEdgesBody890 {
	this := BTExportTessellatedEdgesBody890{}
	return &this
}

// NewBTExportTessellatedEdgesBody890WithDefaults instantiates a new BTExportTessellatedEdgesBody890 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedEdgesBody890WithDefaults() *BTExportTessellatedEdgesBody890 {
	this := BTExportTessellatedEdgesBody890{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesBody890) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesBody890) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesBody890) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportTessellatedEdgesBody890) SetBtType(v string) {
	o.BtType = &v
}

// GetConstituents returns the Constituents field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesBody890) GetConstituents() []string {
	if o == nil || o.Constituents == nil {
		var ret []string
		return ret
	}
	return o.Constituents
}

// GetConstituentsOk returns a tuple with the Constituents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesBody890) GetConstituentsOk() ([]string, bool) {
	if o == nil || o.Constituents == nil {
		return nil, false
	}
	return o.Constituents, true
}

// HasConstituents returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesBody890) HasConstituents() bool {
	if o != nil && o.Constituents != nil {
		return true
	}

	return false
}

// SetConstituents gets a reference to the given []string and assigns it to the Constituents field.
func (o *BTExportTessellatedEdgesBody890) SetConstituents(v []string) {
	o.Constituents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesBody890) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesBody890) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesBody890) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportTessellatedEdgesBody890) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesBody890) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesBody890) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesBody890) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTExportTessellatedEdgesBody890) SetName(v string) {
	o.Name = &v
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *BTExportTessellatedEdgesBody890) GetEdges() []BTExportTessellatedEdgesEdge1364 {
	if o == nil || o.Edges == nil {
		var ret []BTExportTessellatedEdgesEdge1364
		return ret
	}
	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedEdgesBody890) GetEdgesOk() ([]BTExportTessellatedEdgesEdge1364, bool) {
	if o == nil || o.Edges == nil {
		return nil, false
	}
	return o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *BTExportTessellatedEdgesBody890) HasEdges() bool {
	if o != nil && o.Edges != nil {
		return true
	}

	return false
}

// SetEdges gets a reference to the given []BTExportTessellatedEdgesEdge1364 and assigns it to the Edges field.
func (o *BTExportTessellatedEdgesBody890) SetEdges(v []BTExportTessellatedEdgesEdge1364) {
	o.Edges = v
}

func (o BTExportTessellatedEdgesBody890) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTExportTessellatedBody3398, errBTExportTessellatedBody3398 := json.Marshal(o.BTExportTessellatedBody3398)
	if errBTExportTessellatedBody3398 != nil {
		return []byte{}, errBTExportTessellatedBody3398
	}
	errBTExportTessellatedBody3398 = json.Unmarshal([]byte(serializedBTExportTessellatedBody3398), &toSerialize)
	if errBTExportTessellatedBody3398 != nil {
		return []byte{}, errBTExportTessellatedBody3398
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Constituents != nil {
		toSerialize["constituents"] = o.Constituents
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Edges != nil {
		toSerialize["edges"] = o.Edges
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportTessellatedEdgesBody890 struct {
	value *BTExportTessellatedEdgesBody890
	isSet bool
}

func (v NullableBTExportTessellatedEdgesBody890) Get() *BTExportTessellatedEdgesBody890 {
	return v.value
}

func (v *NullableBTExportTessellatedEdgesBody890) Set(val *BTExportTessellatedEdgesBody890) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedEdgesBody890) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedEdgesBody890) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedEdgesBody890(val *BTExportTessellatedEdgesBody890) *NullableBTExportTessellatedEdgesBody890 {
	return &NullableBTExportTessellatedEdgesBody890{value: val, isSet: true}
}

func (v NullableBTExportTessellatedEdgesBody890) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedEdgesBody890) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
