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

// BTCloudStorageAccountInfo struct for BTCloudStorageAccountInfo
type BTCloudStorageAccountInfo struct {
	CanMove         *bool                   `json:"canMove,omitempty"`
	ConnectionName  *string                 `json:"connectionName,omitempty"`
	ConnectionNames []string                `json:"connectionNames,omitempty"`
	CreatedAt       *JSONTime               `json:"createdAt,omitempty"`
	CreatedBy       *BTUserBasicSummaryInfo `json:"createdBy,omitempty"`
	Description     *string                 `json:"description,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id                           *string                 `json:"id,omitempty"`
	IsContainer                  *bool                   `json:"isContainer,omitempty"`
	IsEnterpriseOwned            *bool                   `json:"isEnterpriseOwned,omitempty"`
	IsExternalConnectionResource *bool                   `json:"isExternalConnectionResource,omitempty"`
	IsMutable                    *bool                   `json:"isMutable,omitempty"`
	JsonType                     string                  `json:"jsonType"`
	ModifiedAt                   *JSONTime               `json:"modifiedAt,omitempty"`
	ModifiedBy                   *BTUserBasicSummaryInfo `json:"modifiedBy,omitempty"`
	// Name of the resource.
	Name         *string      `json:"name,omitempty"`
	Owner        *BTOwnerInfo `json:"owner,omitempty"`
	ProjectId    *string      `json:"projectId,omitempty"`
	ResourceType *string      `json:"resourceType,omitempty"`
	TreeHref     *string      `json:"treeHref,omitempty"`
	UnparentHref *string      `json:"unparentHref,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef               *string                   `json:"viewRef,omitempty"`
	CloudStorageAccountId *string                   `json:"cloudStorageAccountId,omitempty"`
	CloudStorageProvider  *int32                    `json:"cloudStorageProvider,omitempty"`
	Enabled               *bool                     `json:"enabled,omitempty"`
	ExportFolder          *BTCloudStorageObjectInfo `json:"exportFolder,omitempty"`
	ImportFolder          *BTCloudStorageObjectInfo `json:"importFolder,omitempty"`
	LastAuthenticated     *JSONTime                 `json:"lastAuthenticated,omitempty"`
}

// NewBTCloudStorageAccountInfo instantiates a new BTCloudStorageAccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCloudStorageAccountInfo(jsonType string) *BTCloudStorageAccountInfo {
	this := BTCloudStorageAccountInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTCloudStorageAccountInfoWithDefaults instantiates a new BTCloudStorageAccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCloudStorageAccountInfoWithDefaults() *BTCloudStorageAccountInfo {
	this := BTCloudStorageAccountInfo{}
	return &this
}

// GetCanMove returns the CanMove field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetCanMove() bool {
	if o == nil || o.CanMove == nil {
		var ret bool
		return ret
	}
	return *o.CanMove
}

// GetCanMoveOk returns a tuple with the CanMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetCanMoveOk() (*bool, bool) {
	if o == nil || o.CanMove == nil {
		return nil, false
	}
	return o.CanMove, true
}

// HasCanMove returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasCanMove() bool {
	if o != nil && o.CanMove != nil {
		return true
	}

	return false
}

// SetCanMove gets a reference to the given bool and assigns it to the CanMove field.
func (o *BTCloudStorageAccountInfo) SetCanMove(v bool) {
	o.CanMove = &v
}

// GetConnectionName returns the ConnectionName field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetConnectionName() string {
	if o == nil || o.ConnectionName == nil {
		var ret string
		return ret
	}
	return *o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetConnectionNameOk() (*string, bool) {
	if o == nil || o.ConnectionName == nil {
		return nil, false
	}
	return o.ConnectionName, true
}

// HasConnectionName returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasConnectionName() bool {
	if o != nil && o.ConnectionName != nil {
		return true
	}

	return false
}

// SetConnectionName gets a reference to the given string and assigns it to the ConnectionName field.
func (o *BTCloudStorageAccountInfo) SetConnectionName(v string) {
	o.ConnectionName = &v
}

// GetConnectionNames returns the ConnectionNames field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetConnectionNames() []string {
	if o == nil || o.ConnectionNames == nil {
		var ret []string
		return ret
	}
	return o.ConnectionNames
}

// GetConnectionNamesOk returns a tuple with the ConnectionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetConnectionNamesOk() ([]string, bool) {
	if o == nil || o.ConnectionNames == nil {
		return nil, false
	}
	return o.ConnectionNames, true
}

// HasConnectionNames returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasConnectionNames() bool {
	if o != nil && o.ConnectionNames != nil {
		return true
	}

	return false
}

// SetConnectionNames gets a reference to the given []string and assigns it to the ConnectionNames field.
func (o *BTCloudStorageAccountInfo) SetConnectionNames(v []string) {
	o.ConnectionNames = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTCloudStorageAccountInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTCloudStorageAccountInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTCloudStorageAccountInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTCloudStorageAccountInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTCloudStorageAccountInfo) SetId(v string) {
	o.Id = &v
}

// GetIsContainer returns the IsContainer field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetIsContainer() bool {
	if o == nil || o.IsContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetIsContainerOk() (*bool, bool) {
	if o == nil || o.IsContainer == nil {
		return nil, false
	}
	return o.IsContainer, true
}

// HasIsContainer returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasIsContainer() bool {
	if o != nil && o.IsContainer != nil {
		return true
	}

	return false
}

// SetIsContainer gets a reference to the given bool and assigns it to the IsContainer field.
func (o *BTCloudStorageAccountInfo) SetIsContainer(v bool) {
	o.IsContainer = &v
}

// GetIsEnterpriseOwned returns the IsEnterpriseOwned field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetIsEnterpriseOwned() bool {
	if o == nil || o.IsEnterpriseOwned == nil {
		var ret bool
		return ret
	}
	return *o.IsEnterpriseOwned
}

// GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetIsEnterpriseOwnedOk() (*bool, bool) {
	if o == nil || o.IsEnterpriseOwned == nil {
		return nil, false
	}
	return o.IsEnterpriseOwned, true
}

// HasIsEnterpriseOwned returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasIsEnterpriseOwned() bool {
	if o != nil && o.IsEnterpriseOwned != nil {
		return true
	}

	return false
}

// SetIsEnterpriseOwned gets a reference to the given bool and assigns it to the IsEnterpriseOwned field.
func (o *BTCloudStorageAccountInfo) SetIsEnterpriseOwned(v bool) {
	o.IsEnterpriseOwned = &v
}

// GetIsExternalConnectionResource returns the IsExternalConnectionResource field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetIsExternalConnectionResource() bool {
	if o == nil || o.IsExternalConnectionResource == nil {
		var ret bool
		return ret
	}
	return *o.IsExternalConnectionResource
}

// GetIsExternalConnectionResourceOk returns a tuple with the IsExternalConnectionResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetIsExternalConnectionResourceOk() (*bool, bool) {
	if o == nil || o.IsExternalConnectionResource == nil {
		return nil, false
	}
	return o.IsExternalConnectionResource, true
}

// HasIsExternalConnectionResource returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasIsExternalConnectionResource() bool {
	if o != nil && o.IsExternalConnectionResource != nil {
		return true
	}

	return false
}

// SetIsExternalConnectionResource gets a reference to the given bool and assigns it to the IsExternalConnectionResource field.
func (o *BTCloudStorageAccountInfo) SetIsExternalConnectionResource(v bool) {
	o.IsExternalConnectionResource = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *BTCloudStorageAccountInfo) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetJsonType returns the JsonType field value
func (o *BTCloudStorageAccountInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTCloudStorageAccountInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTCloudStorageAccountInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTCloudStorageAccountInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTCloudStorageAccountInfo) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetOwner() BTOwnerInfo {
	if o == nil || o.Owner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BTOwnerInfo and assigns it to the Owner field.
func (o *BTCloudStorageAccountInfo) SetOwner(v BTOwnerInfo) {
	o.Owner = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTCloudStorageAccountInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BTCloudStorageAccountInfo) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetTreeHref returns the TreeHref field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetTreeHref() string {
	if o == nil || o.TreeHref == nil {
		var ret string
		return ret
	}
	return *o.TreeHref
}

// GetTreeHrefOk returns a tuple with the TreeHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetTreeHrefOk() (*string, bool) {
	if o == nil || o.TreeHref == nil {
		return nil, false
	}
	return o.TreeHref, true
}

// HasTreeHref returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasTreeHref() bool {
	if o != nil && o.TreeHref != nil {
		return true
	}

	return false
}

// SetTreeHref gets a reference to the given string and assigns it to the TreeHref field.
func (o *BTCloudStorageAccountInfo) SetTreeHref(v string) {
	o.TreeHref = &v
}

// GetUnparentHref returns the UnparentHref field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetUnparentHref() string {
	if o == nil || o.UnparentHref == nil {
		var ret string
		return ret
	}
	return *o.UnparentHref
}

// GetUnparentHrefOk returns a tuple with the UnparentHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetUnparentHrefOk() (*string, bool) {
	if o == nil || o.UnparentHref == nil {
		return nil, false
	}
	return o.UnparentHref, true
}

// HasUnparentHref returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasUnparentHref() bool {
	if o != nil && o.UnparentHref != nil {
		return true
	}

	return false
}

// SetUnparentHref gets a reference to the given string and assigns it to the UnparentHref field.
func (o *BTCloudStorageAccountInfo) SetUnparentHref(v string) {
	o.UnparentHref = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTCloudStorageAccountInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetCloudStorageAccountId returns the CloudStorageAccountId field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetCloudStorageAccountId() string {
	if o == nil || o.CloudStorageAccountId == nil {
		var ret string
		return ret
	}
	return *o.CloudStorageAccountId
}

// GetCloudStorageAccountIdOk returns a tuple with the CloudStorageAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetCloudStorageAccountIdOk() (*string, bool) {
	if o == nil || o.CloudStorageAccountId == nil {
		return nil, false
	}
	return o.CloudStorageAccountId, true
}

// HasCloudStorageAccountId returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasCloudStorageAccountId() bool {
	if o != nil && o.CloudStorageAccountId != nil {
		return true
	}

	return false
}

// SetCloudStorageAccountId gets a reference to the given string and assigns it to the CloudStorageAccountId field.
func (o *BTCloudStorageAccountInfo) SetCloudStorageAccountId(v string) {
	o.CloudStorageAccountId = &v
}

// GetCloudStorageProvider returns the CloudStorageProvider field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetCloudStorageProvider() int32 {
	if o == nil || o.CloudStorageProvider == nil {
		var ret int32
		return ret
	}
	return *o.CloudStorageProvider
}

// GetCloudStorageProviderOk returns a tuple with the CloudStorageProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetCloudStorageProviderOk() (*int32, bool) {
	if o == nil || o.CloudStorageProvider == nil {
		return nil, false
	}
	return o.CloudStorageProvider, true
}

// HasCloudStorageProvider returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasCloudStorageProvider() bool {
	if o != nil && o.CloudStorageProvider != nil {
		return true
	}

	return false
}

// SetCloudStorageProvider gets a reference to the given int32 and assigns it to the CloudStorageProvider field.
func (o *BTCloudStorageAccountInfo) SetCloudStorageProvider(v int32) {
	o.CloudStorageProvider = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BTCloudStorageAccountInfo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExportFolder returns the ExportFolder field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetExportFolder() BTCloudStorageObjectInfo {
	if o == nil || o.ExportFolder == nil {
		var ret BTCloudStorageObjectInfo
		return ret
	}
	return *o.ExportFolder
}

// GetExportFolderOk returns a tuple with the ExportFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetExportFolderOk() (*BTCloudStorageObjectInfo, bool) {
	if o == nil || o.ExportFolder == nil {
		return nil, false
	}
	return o.ExportFolder, true
}

// HasExportFolder returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasExportFolder() bool {
	if o != nil && o.ExportFolder != nil {
		return true
	}

	return false
}

// SetExportFolder gets a reference to the given BTCloudStorageObjectInfo and assigns it to the ExportFolder field.
func (o *BTCloudStorageAccountInfo) SetExportFolder(v BTCloudStorageObjectInfo) {
	o.ExportFolder = &v
}

// GetImportFolder returns the ImportFolder field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetImportFolder() BTCloudStorageObjectInfo {
	if o == nil || o.ImportFolder == nil {
		var ret BTCloudStorageObjectInfo
		return ret
	}
	return *o.ImportFolder
}

// GetImportFolderOk returns a tuple with the ImportFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetImportFolderOk() (*BTCloudStorageObjectInfo, bool) {
	if o == nil || o.ImportFolder == nil {
		return nil, false
	}
	return o.ImportFolder, true
}

// HasImportFolder returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasImportFolder() bool {
	if o != nil && o.ImportFolder != nil {
		return true
	}

	return false
}

// SetImportFolder gets a reference to the given BTCloudStorageObjectInfo and assigns it to the ImportFolder field.
func (o *BTCloudStorageAccountInfo) SetImportFolder(v BTCloudStorageObjectInfo) {
	o.ImportFolder = &v
}

// GetLastAuthenticated returns the LastAuthenticated field value if set, zero value otherwise.
func (o *BTCloudStorageAccountInfo) GetLastAuthenticated() JSONTime {
	if o == nil || o.LastAuthenticated == nil {
		var ret JSONTime
		return ret
	}
	return *o.LastAuthenticated
}

// GetLastAuthenticatedOk returns a tuple with the LastAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCloudStorageAccountInfo) GetLastAuthenticatedOk() (*JSONTime, bool) {
	if o == nil || o.LastAuthenticated == nil {
		return nil, false
	}
	return o.LastAuthenticated, true
}

// HasLastAuthenticated returns a boolean if a field has been set.
func (o *BTCloudStorageAccountInfo) HasLastAuthenticated() bool {
	if o != nil && o.LastAuthenticated != nil {
		return true
	}

	return false
}

// SetLastAuthenticated gets a reference to the given JSONTime and assigns it to the LastAuthenticated field.
func (o *BTCloudStorageAccountInfo) SetLastAuthenticated(v JSONTime) {
	o.LastAuthenticated = &v
}

func (o BTCloudStorageAccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanMove != nil {
		toSerialize["canMove"] = o.CanMove
	}
	if o.ConnectionName != nil {
		toSerialize["connectionName"] = o.ConnectionName
	}
	if o.ConnectionNames != nil {
		toSerialize["connectionNames"] = o.ConnectionNames
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsContainer != nil {
		toSerialize["isContainer"] = o.IsContainer
	}
	if o.IsEnterpriseOwned != nil {
		toSerialize["isEnterpriseOwned"] = o.IsEnterpriseOwned
	}
	if o.IsExternalConnectionResource != nil {
		toSerialize["isExternalConnectionResource"] = o.IsExternalConnectionResource
	}
	if o.IsMutable != nil {
		toSerialize["isMutable"] = o.IsMutable
	}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.TreeHref != nil {
		toSerialize["treeHref"] = o.TreeHref
	}
	if o.UnparentHref != nil {
		toSerialize["unparentHref"] = o.UnparentHref
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.CloudStorageAccountId != nil {
		toSerialize["cloudStorageAccountId"] = o.CloudStorageAccountId
	}
	if o.CloudStorageProvider != nil {
		toSerialize["cloudStorageProvider"] = o.CloudStorageProvider
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExportFolder != nil {
		toSerialize["exportFolder"] = o.ExportFolder
	}
	if o.ImportFolder != nil {
		toSerialize["importFolder"] = o.ImportFolder
	}
	if o.LastAuthenticated != nil {
		toSerialize["lastAuthenticated"] = o.LastAuthenticated
	}
	return json.Marshal(toSerialize)
}

type NullableBTCloudStorageAccountInfo struct {
	value *BTCloudStorageAccountInfo
	isSet bool
}

func (v NullableBTCloudStorageAccountInfo) Get() *BTCloudStorageAccountInfo {
	return v.value
}

func (v *NullableBTCloudStorageAccountInfo) Set(val *BTCloudStorageAccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCloudStorageAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCloudStorageAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCloudStorageAccountInfo(val *BTCloudStorageAccountInfo) *NullableBTCloudStorageAccountInfo {
	return &NullableBTCloudStorageAccountInfo{value: val, isSet: true}
}

func (v NullableBTCloudStorageAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCloudStorageAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
