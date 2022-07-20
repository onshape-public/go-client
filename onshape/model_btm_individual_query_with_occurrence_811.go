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
	"fmt"
)

// BTMIndividualQueryWithOccurrence811 - struct for BTMIndividualQueryWithOccurrence811
type BTMIndividualQueryWithOccurrence811 struct {
	implBTMIndividualQueryWithOccurrence811 interface{}
}

// BTMInferenceQueryWithOccurrence1083AsBTMIndividualQueryWithOccurrence811 is a convenience function that returns BTMInferenceQueryWithOccurrence1083 wrapped in BTMIndividualQueryWithOccurrence811
func (o *BTMInferenceQueryWithOccurrence1083) AsBTMIndividualQueryWithOccurrence811() *BTMIndividualQueryWithOccurrence811 {
	return &BTMIndividualQueryWithOccurrence811{o}
}

// NewBTMIndividualQueryWithOccurrence811 instantiates a new BTMIndividualQueryWithOccurrence811 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualQueryWithOccurrence811() *BTMIndividualQueryWithOccurrence811 {
	this := BTMIndividualQueryWithOccurrence811{Newbase_BTMIndividualQueryWithOccurrence811()}
	return &this
}

// NewBTMIndividualQueryWithOccurrence811WithDefaults instantiates a new BTMIndividualQueryWithOccurrence811 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualQueryWithOccurrence811WithDefaults() *BTMIndividualQueryWithOccurrence811 {
	this := BTMIndividualQueryWithOccurrence811{Newbase_BTMIndividualQueryWithOccurrence811WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetBtType() string {
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
func (o *BTMIndividualQueryWithOccurrence811) GetBtTypeOk() (*string, bool) {
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
func (o *BTMIndividualQueryWithOccurrence811) HasBtType() bool {
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
func (o *BTMIndividualQueryWithOccurrence811) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetDeterministicIdList() BTMIndividualQueryBase139 {
	type getResult interface {
		GetDeterministicIdList() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdList()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdListOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasDeterministicIdList() bool {
	type getResult interface {
		HasDeterministicIdList() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIdList()
	} else {
		return false
	}
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualQueryWithOccurrence811) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetDeterministicIdList(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetDeterministicIdList(v)
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetDeterministicIds() []string {
	type getResult interface {
		GetDeterministicIds() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIds()
	} else {
		var de []string
		return de
	}
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetDeterministicIdsOk() ([]string, bool) {
	type getResult interface {
		GetDeterministicIdsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdsOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasDeterministicIds() bool {
	type getResult interface {
		HasDeterministicIds() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIds()
	} else {
		return false
	}
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualQueryWithOccurrence811) SetDeterministicIds(v []string) {
	type getResult interface {
		SetDeterministicIds(v []string)
	}

	o.GetActualInstance().(getResult).SetDeterministicIds(v)
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetFullPathAsString() string {
	type getResult interface {
		GetFullPathAsString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullPathAsString()
	} else {
		var de string
		return de
	}
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetFullPathAsStringOk() (*string, bool) {
	type getResult interface {
		GetFullPathAsStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetFullPathAsStringOk()
	} else {
		return nil, false
	}
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasFullPathAsString() bool {
	type getResult interface {
		HasFullPathAsString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasFullPathAsString()
	} else {
		return false
	}
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTMIndividualQueryWithOccurrence811) SetFullPathAsString(v string) {
	type getResult interface {
		SetFullPathAsString(v string)
	}

	o.GetActualInstance().(getResult).SetFullPathAsString(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMIndividualQueryWithOccurrence811) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetNodeId() string {
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
func (o *BTMIndividualQueryWithOccurrence811) GetNodeIdOk() (*string, bool) {
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
func (o *BTMIndividualQueryWithOccurrence811) HasNodeId() bool {
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
func (o *BTMIndividualQueryWithOccurrence811) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetOccurrence() BTOccurrence74 {
	type getResult interface {
		GetOccurrence() BTOccurrence74
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrence()
	} else {
		var de BTOccurrence74
		return de
	}
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetOccurrenceOk() (*BTOccurrence74, bool) {
	type getResult interface {
		GetOccurrenceOk() (*BTOccurrence74, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetOccurrenceOk()
	} else {
		return nil, false
	}
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasOccurrence() bool {
	type getResult interface {
		HasOccurrence() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasOccurrence()
	} else {
		return false
	}
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTMIndividualQueryWithOccurrence811) SetOccurrence(v BTOccurrence74) {
	type getResult interface {
		SetOccurrence(v BTOccurrence74)
	}

	o.GetActualInstance().(getResult).SetOccurrence(v)
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetPath() []string {
	type getResult interface {
		GetPath() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPath()
	} else {
		var de []string
		return de
	}
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetPathOk() ([]string, bool) {
	type getResult interface {
		GetPathOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPathOk()
	} else {
		return nil, false
	}
}

// HasPath returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasPath() bool {
	type getResult interface {
		HasPath() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPath()
	} else {
		return false
	}
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTMIndividualQueryWithOccurrence811) SetPath(v []string) {
	type getResult interface {
		SetPath(v []string)
	}

	o.GetActualInstance().(getResult).SetPath(v)
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetQuery() BTMIndividualQueryBase139 {
	type getResult interface {
		GetQuery() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQuery()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetQueryOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryOk()
	} else {
		return nil, false
	}
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasQuery() bool {
	type getResult interface {
		HasQuery() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQuery()
	} else {
		return false
	}
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualQueryWithOccurrence811) SetQuery(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetQuery(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetQuery(v)
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetQueryString() string {
	type getResult interface {
		GetQueryString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryString()
	} else {
		var de string
		return de
	}
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetQueryStringOk() (*string, bool) {
	type getResult interface {
		GetQueryStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryStringOk()
	} else {
		return nil, false
	}
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasQueryString() bool {
	type getResult interface {
		HasQueryString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQueryString()
	} else {
		return false
	}
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualQueryWithOccurrence811) SetQueryString(v string) {
	type getResult interface {
		SetQueryString(v string)
	}

	o.GetActualInstance().(getResult).SetQueryString(v)
}

// GetEntityQuery returns the EntityQuery field value if set, zero value otherwise.
func (o *BTMIndividualQueryWithOccurrence811) GetEntityQuery() string {
	type getResult interface {
		GetEntityQuery() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEntityQuery()
	} else {
		var de string
		return de
	}
}

// GetEntityQueryOk returns a tuple with the EntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQueryWithOccurrence811) GetEntityQueryOk() (*string, bool) {
	type getResult interface {
		GetEntityQueryOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetEntityQueryOk()
	} else {
		return nil, false
	}
}

// HasEntityQuery returns a boolean if a field has been set.
func (o *BTMIndividualQueryWithOccurrence811) HasEntityQuery() bool {
	type getResult interface {
		HasEntityQuery() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasEntityQuery()
	} else {
		return false
	}
}

// SetEntityQuery gets a reference to the given string and assigns it to the EntityQuery field.
func (o *BTMIndividualQueryWithOccurrence811) SetEntityQuery(v string) {
	type getResult interface {
		SetEntityQuery(v string)
	}

	o.GetActualInstance().(getResult).SetEntityQuery(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMIndividualQueryWithOccurrence811) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMInferenceQueryWithOccurrence-1083'
	if jsonDict["btType"] == "BTMInferenceQueryWithOccurrence-1083" {
		// try to unmarshal JSON data into BTMInferenceQueryWithOccurrence1083
		var qr *BTMInferenceQueryWithOccurrence1083
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQueryWithOccurrence811 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQueryWithOccurrence811 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrence811 as BTMInferenceQueryWithOccurrence1083: %s", err.Error())
		}
	}

	var qtx *base_BTMIndividualQueryWithOccurrence811
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMIndividualQueryWithOccurrence811 = qtx
		return nil // data stored in dst.base_BTMIndividualQueryWithOccurrence811, return on the first match
	} else {
		dst.implBTMIndividualQueryWithOccurrence811 = nil
		return fmt.Errorf("Failed to unmarshal BTMIndividualQueryWithOccurrence811 as base_BTMIndividualQueryWithOccurrence811: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMIndividualQueryWithOccurrence811) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMIndividualQueryWithOccurrence811) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMIndividualQueryWithOccurrence811
}

type NullableBTMIndividualQueryWithOccurrence811 struct {
	value *BTMIndividualQueryWithOccurrence811
	isSet bool
}

func (v NullableBTMIndividualQueryWithOccurrence811) Get() *BTMIndividualQueryWithOccurrence811 {
	return v.value
}

func (v *NullableBTMIndividualQueryWithOccurrence811) Set(val *BTMIndividualQueryWithOccurrence811) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualQueryWithOccurrence811) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualQueryWithOccurrence811) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualQueryWithOccurrence811(val *BTMIndividualQueryWithOccurrence811) *NullableBTMIndividualQueryWithOccurrence811 {
	return &NullableBTMIndividualQueryWithOccurrence811{value: val, isSet: true}
}

func (v NullableBTMIndividualQueryWithOccurrence811) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualQueryWithOccurrence811) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMIndividualQueryWithOccurrence811 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	FullPathAsString    *string                    `json:"fullPathAsString,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Occurrence          *BTOccurrence74            `json:"occurrence,omitempty"`
	Path                []string                   `json:"path,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
	EntityQuery         *string                    `json:"entityQuery,omitempty"`
}

// Newbase_BTMIndividualQueryWithOccurrence811 instantiates a new base_BTMIndividualQueryWithOccurrence811 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMIndividualQueryWithOccurrence811() *base_BTMIndividualQueryWithOccurrence811 {
	this := base_BTMIndividualQueryWithOccurrence811{}
	return &this
}

// Newbase_BTMIndividualQueryWithOccurrence811WithDefaults instantiates a new base_BTMIndividualQueryWithOccurrence811 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMIndividualQueryWithOccurrence811WithDefaults() *base_BTMIndividualQueryWithOccurrence811 {
	this := base_BTMIndividualQueryWithOccurrence811{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetPath(v []string) {
	o.Path = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetQueryString(v string) {
	o.QueryString = &v
}

// GetEntityQuery returns the EntityQuery field value if set, zero value otherwise.
func (o *base_BTMIndividualQueryWithOccurrence811) GetEntityQuery() string {
	if o == nil || o.EntityQuery == nil {
		var ret string
		return ret
	}
	return *o.EntityQuery
}

// GetEntityQueryOk returns a tuple with the EntityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) GetEntityQueryOk() (*string, bool) {
	if o == nil || o.EntityQuery == nil {
		return nil, false
	}
	return o.EntityQuery, true
}

// HasEntityQuery returns a boolean if a field has been set.
func (o *base_BTMIndividualQueryWithOccurrence811) HasEntityQuery() bool {
	if o != nil && o.EntityQuery != nil {
		return true
	}

	return false
}

// SetEntityQuery gets a reference to the given string and assigns it to the EntityQuery field.
func (o *base_BTMIndividualQueryWithOccurrence811) SetEntityQuery(v string) {
	o.EntityQuery = &v
}

func (o base_BTMIndividualQueryWithOccurrence811) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.EntityQuery != nil {
		toSerialize["entityQuery"] = o.EntityQuery
	}
	return json.Marshal(toSerialize)
}