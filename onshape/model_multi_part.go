/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
	"encoding/json"
)

// MultiPart struct for MultiPart
type MultiPart struct {
	BodyParts *[]BodyPart `json:"bodyParts,omitempty"`
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	Entity *map[string]interface{} `json:"entity,omitempty"`
	Headers *map[string][]string `json:"headers,omitempty"`
	MediaType *BodyPartMediaType `json:"mediaType,omitempty"`
	MessageBodyWorkers *map[string]interface{} `json:"messageBodyWorkers,omitempty"`
	ParameterizedHeaders *map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
	Parent *MultiPart `json:"parent,omitempty"`
	Providers *map[string]interface{} `json:"providers,omitempty"`
}

// NewMultiPart instantiates a new MultiPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiPart() *MultiPart {
	this := MultiPart{}
	return &this
}

// NewMultiPartWithDefaults instantiates a new MultiPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiPartWithDefaults() *MultiPart {
	this := MultiPart{}
	return &this
}

// GetBodyParts returns the BodyParts field value if set, zero value otherwise.
func (o *MultiPart) GetBodyParts() []BodyPart {
	if o == nil || o.BodyParts == nil {
		var ret []BodyPart
		return ret
	}
	return *o.BodyParts
}

// GetBodyPartsOk returns a tuple with the BodyParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetBodyPartsOk() (*[]BodyPart, bool) {
	if o == nil || o.BodyParts == nil {
		return nil, false
	}
	return o.BodyParts, true
}

// HasBodyParts returns a boolean if a field has been set.
func (o *MultiPart) HasBodyParts() bool {
	if o != nil && o.BodyParts != nil {
		return true
	}

	return false
}

// SetBodyParts gets a reference to the given []BodyPart and assigns it to the BodyParts field.
func (o *MultiPart) SetBodyParts(v []BodyPart) {
	o.BodyParts = &v
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *MultiPart) GetContentDisposition() ContentDisposition {
	if o == nil || o.ContentDisposition == nil {
		var ret ContentDisposition
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetContentDispositionOk() (*ContentDisposition, bool) {
	if o == nil || o.ContentDisposition == nil {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *MultiPart) HasContentDisposition() bool {
	if o != nil && o.ContentDisposition != nil {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given ContentDisposition and assigns it to the ContentDisposition field.
func (o *MultiPart) SetContentDisposition(v ContentDisposition) {
	o.ContentDisposition = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *MultiPart) GetEntity() map[string]interface{} {
	if o == nil || o.Entity == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetEntityOk() (*map[string]interface{}, bool) {
	if o == nil || o.Entity == nil {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *MultiPart) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *MultiPart) SetEntity(v map[string]interface{}) {
	o.Entity = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *MultiPart) GetHeaders() map[string][]string {
	if o == nil || o.Headers == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetHeadersOk() (*map[string][]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *MultiPart) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string][]string and assigns it to the Headers field.
func (o *MultiPart) SetHeaders(v map[string][]string) {
	o.Headers = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *MultiPart) GetMediaType() BodyPartMediaType {
	if o == nil || o.MediaType == nil {
		var ret BodyPartMediaType
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetMediaTypeOk() (*BodyPartMediaType, bool) {
	if o == nil || o.MediaType == nil {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *MultiPart) HasMediaType() bool {
	if o != nil && o.MediaType != nil {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given BodyPartMediaType and assigns it to the MediaType field.
func (o *MultiPart) SetMediaType(v BodyPartMediaType) {
	o.MediaType = &v
}

// GetMessageBodyWorkers returns the MessageBodyWorkers field value if set, zero value otherwise.
func (o *MultiPart) GetMessageBodyWorkers() map[string]interface{} {
	if o == nil || o.MessageBodyWorkers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.MessageBodyWorkers
}

// GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetMessageBodyWorkersOk() (*map[string]interface{}, bool) {
	if o == nil || o.MessageBodyWorkers == nil {
		return nil, false
	}
	return o.MessageBodyWorkers, true
}

// HasMessageBodyWorkers returns a boolean if a field has been set.
func (o *MultiPart) HasMessageBodyWorkers() bool {
	if o != nil && o.MessageBodyWorkers != nil {
		return true
	}

	return false
}

// SetMessageBodyWorkers gets a reference to the given map[string]interface{} and assigns it to the MessageBodyWorkers field.
func (o *MultiPart) SetMessageBodyWorkers(v map[string]interface{}) {
	o.MessageBodyWorkers = &v
}

// GetParameterizedHeaders returns the ParameterizedHeaders field value if set, zero value otherwise.
func (o *MultiPart) GetParameterizedHeaders() map[string][]ParameterizedHeader {
	if o == nil || o.ParameterizedHeaders == nil {
		var ret map[string][]ParameterizedHeader
		return ret
	}
	return *o.ParameterizedHeaders
}

// GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool) {
	if o == nil || o.ParameterizedHeaders == nil {
		return nil, false
	}
	return o.ParameterizedHeaders, true
}

// HasParameterizedHeaders returns a boolean if a field has been set.
func (o *MultiPart) HasParameterizedHeaders() bool {
	if o != nil && o.ParameterizedHeaders != nil {
		return true
	}

	return false
}

// SetParameterizedHeaders gets a reference to the given map[string][]ParameterizedHeader and assigns it to the ParameterizedHeaders field.
func (o *MultiPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader) {
	o.ParameterizedHeaders = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *MultiPart) GetParent() MultiPart {
	if o == nil || o.Parent == nil {
		var ret MultiPart
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetParentOk() (*MultiPart, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *MultiPart) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given MultiPart and assigns it to the Parent field.
func (o *MultiPart) SetParent(v MultiPart) {
	o.Parent = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *MultiPart) GetProviders() map[string]interface{} {
	if o == nil || o.Providers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPart) GetProvidersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *MultiPart) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given map[string]interface{} and assigns it to the Providers field.
func (o *MultiPart) SetProviders(v map[string]interface{}) {
	o.Providers = &v
}

func (o MultiPart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BodyParts != nil {
		toSerialize["bodyParts"] = o.BodyParts
	}
	if o.ContentDisposition != nil {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if o.Entity != nil {
		toSerialize["entity"] = o.Entity
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.MediaType != nil {
		toSerialize["mediaType"] = o.MediaType
	}
	if o.MessageBodyWorkers != nil {
		toSerialize["messageBodyWorkers"] = o.MessageBodyWorkers
	}
	if o.ParameterizedHeaders != nil {
		toSerialize["parameterizedHeaders"] = o.ParameterizedHeaders
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	return json.Marshal(toSerialize)
}

type NullableMultiPart struct {
	value *MultiPart
	isSet bool
}

func (v NullableMultiPart) Get() *MultiPart {
	return v.value
}

func (v *NullableMultiPart) Set(val *MultiPart) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiPart) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiPart(val *MultiPart) *NullableMultiPart {
	return &NullableMultiPart{value: val, isSet: true}
}

func (v NullableMultiPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
