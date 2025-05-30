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

// BTThumbnailInfo struct for BTThumbnailInfo
type BTThumbnailInfo struct {
	Href           *string                 `json:"href,omitempty"`
	Id             *string                 `json:"id,omitempty"`
	SecondarySizes [][]BTThumbnailSizeInfo `json:"secondarySizes,omitempty"`
	Sizes          []BTThumbnailSizeInfo   `json:"sizes,omitempty"`
}

// NewBTThumbnailInfo instantiates a new BTThumbnailInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTThumbnailInfo() *BTThumbnailInfo {
	this := BTThumbnailInfo{}
	return &this
}

// NewBTThumbnailInfoWithDefaults instantiates a new BTThumbnailInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTThumbnailInfoWithDefaults() *BTThumbnailInfo {
	this := BTThumbnailInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTThumbnailInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTThumbnailInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTThumbnailInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTThumbnailInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTThumbnailInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTThumbnailInfo) SetId(v string) {
	o.Id = &v
}

// GetSecondarySizes returns the SecondarySizes field value if set, zero value otherwise.
func (o *BTThumbnailInfo) GetSecondarySizes() [][]BTThumbnailSizeInfo {
	if o == nil || o.SecondarySizes == nil {
		var ret [][]BTThumbnailSizeInfo
		return ret
	}
	return o.SecondarySizes
}

// GetSecondarySizesOk returns a tuple with the SecondarySizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailInfo) GetSecondarySizesOk() ([][]BTThumbnailSizeInfo, bool) {
	if o == nil || o.SecondarySizes == nil {
		return nil, false
	}
	return o.SecondarySizes, true
}

// HasSecondarySizes returns a boolean if a field has been set.
func (o *BTThumbnailInfo) HasSecondarySizes() bool {
	if o != nil && o.SecondarySizes != nil {
		return true
	}

	return false
}

// SetSecondarySizes gets a reference to the given [][]BTThumbnailSizeInfo and assigns it to the SecondarySizes field.
func (o *BTThumbnailInfo) SetSecondarySizes(v [][]BTThumbnailSizeInfo) {
	o.SecondarySizes = v
}

// GetSizes returns the Sizes field value if set, zero value otherwise.
func (o *BTThumbnailInfo) GetSizes() []BTThumbnailSizeInfo {
	if o == nil || o.Sizes == nil {
		var ret []BTThumbnailSizeInfo
		return ret
	}
	return o.Sizes
}

// GetSizesOk returns a tuple with the Sizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTThumbnailInfo) GetSizesOk() ([]BTThumbnailSizeInfo, bool) {
	if o == nil || o.Sizes == nil {
		return nil, false
	}
	return o.Sizes, true
}

// HasSizes returns a boolean if a field has been set.
func (o *BTThumbnailInfo) HasSizes() bool {
	if o != nil && o.Sizes != nil {
		return true
	}

	return false
}

// SetSizes gets a reference to the given []BTThumbnailSizeInfo and assigns it to the Sizes field.
func (o *BTThumbnailInfo) SetSizes(v []BTThumbnailSizeInfo) {
	o.Sizes = v
}

func (o BTThumbnailInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SecondarySizes != nil {
		toSerialize["secondarySizes"] = o.SecondarySizes
	}
	if o.Sizes != nil {
		toSerialize["sizes"] = o.Sizes
	}
	return json.Marshal(toSerialize)
}

type NullableBTThumbnailInfo struct {
	value *BTThumbnailInfo
	isSet bool
}

func (v NullableBTThumbnailInfo) Get() *BTThumbnailInfo {
	return v.value
}

func (v *NullableBTThumbnailInfo) Set(val *BTThumbnailInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTThumbnailInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTThumbnailInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTThumbnailInfo(val *BTThumbnailInfo) *NullableBTThumbnailInfo {
	return &NullableBTThumbnailInfo{value: val, isSet: true}
}

func (v NullableBTThumbnailInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTThumbnailInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
