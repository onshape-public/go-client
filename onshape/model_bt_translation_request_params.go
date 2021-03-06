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
	"os"
)

// BTTranslationRequestParams struct for BTTranslationRequestParams
type BTTranslationRequestParams struct {
	AllowFaultyParts *bool `json:"allowFaultyParts,omitempty"`
	CreateComposite *bool `json:"createComposite,omitempty"`
	CreateDrawingIfPossible *bool `json:"createDrawingIfPossible,omitempty"`
	EncodedFilename *string `json:"encodedFilename,omitempty"`
	ExtractAssemblyHierarchy *bool `json:"extractAssemblyHierarchy,omitempty"`
	File **os.File `json:"file,omitempty"`
	FileBodyWithDetails *FormDataBodyPart `json:"fileBodyWithDetails,omitempty"`
	FileContentLength *int64 `json:"fileContentLength,omitempty"`
	FileDetail *FormDataContentDisposition `json:"fileDetail,omitempty"`
	FlattenAssemblies *bool `json:"flattenAssemblies,omitempty"`
	FormatName *string `json:"formatName,omitempty"`
	IsyAxisIsUp *bool `json:"isyAxisIsUp,omitempty"`
	JoinAdjacentSurfaces *bool `json:"joinAdjacentSurfaces,omitempty"`
	LocationElementId *string `json:"locationElementId,omitempty"`
	LocationGroupId *string `json:"locationGroupId,omitempty"`
	LocationPosition *int32 `json:"locationPosition,omitempty"`
	NotifyUser *bool `json:"notifyUser,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	Public *bool `json:"public,omitempty"`
	SplitAssembliesIntoMultipleDocuments *bool `json:"splitAssembliesIntoMultipleDocuments,omitempty"`
	StoreInDocument *bool `json:"storeInDocument,omitempty"`
	Translate *bool `json:"translate,omitempty"`
	Unit *string `json:"unit,omitempty"`
	UploadId *string `json:"uploadId,omitempty"`
	VersionString *string `json:"versionString,omitempty"`
}

// NewBTTranslationRequestParams instantiates a new BTTranslationRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTranslationRequestParams() *BTTranslationRequestParams {
	this := BTTranslationRequestParams{}
	return &this
}

// NewBTTranslationRequestParamsWithDefaults instantiates a new BTTranslationRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTranslationRequestParamsWithDefaults() *BTTranslationRequestParams {
	this := BTTranslationRequestParams{}
	return &this
}

// GetAllowFaultyParts returns the AllowFaultyParts field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetAllowFaultyParts() bool {
	if o == nil || o.AllowFaultyParts == nil {
		var ret bool
		return ret
	}
	return *o.AllowFaultyParts
}

// GetAllowFaultyPartsOk returns a tuple with the AllowFaultyParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetAllowFaultyPartsOk() (*bool, bool) {
	if o == nil || o.AllowFaultyParts == nil {
		return nil, false
	}
	return o.AllowFaultyParts, true
}

// HasAllowFaultyParts returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasAllowFaultyParts() bool {
	if o != nil && o.AllowFaultyParts != nil {
		return true
	}

	return false
}

// SetAllowFaultyParts gets a reference to the given bool and assigns it to the AllowFaultyParts field.
func (o *BTTranslationRequestParams) SetAllowFaultyParts(v bool) {
	o.AllowFaultyParts = &v
}

// GetCreateComposite returns the CreateComposite field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetCreateComposite() bool {
	if o == nil || o.CreateComposite == nil {
		var ret bool
		return ret
	}
	return *o.CreateComposite
}

// GetCreateCompositeOk returns a tuple with the CreateComposite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetCreateCompositeOk() (*bool, bool) {
	if o == nil || o.CreateComposite == nil {
		return nil, false
	}
	return o.CreateComposite, true
}

// HasCreateComposite returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasCreateComposite() bool {
	if o != nil && o.CreateComposite != nil {
		return true
	}

	return false
}

// SetCreateComposite gets a reference to the given bool and assigns it to the CreateComposite field.
func (o *BTTranslationRequestParams) SetCreateComposite(v bool) {
	o.CreateComposite = &v
}

// GetCreateDrawingIfPossible returns the CreateDrawingIfPossible field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetCreateDrawingIfPossible() bool {
	if o == nil || o.CreateDrawingIfPossible == nil {
		var ret bool
		return ret
	}
	return *o.CreateDrawingIfPossible
}

// GetCreateDrawingIfPossibleOk returns a tuple with the CreateDrawingIfPossible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetCreateDrawingIfPossibleOk() (*bool, bool) {
	if o == nil || o.CreateDrawingIfPossible == nil {
		return nil, false
	}
	return o.CreateDrawingIfPossible, true
}

// HasCreateDrawingIfPossible returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasCreateDrawingIfPossible() bool {
	if o != nil && o.CreateDrawingIfPossible != nil {
		return true
	}

	return false
}

// SetCreateDrawingIfPossible gets a reference to the given bool and assigns it to the CreateDrawingIfPossible field.
func (o *BTTranslationRequestParams) SetCreateDrawingIfPossible(v bool) {
	o.CreateDrawingIfPossible = &v
}

// GetEncodedFilename returns the EncodedFilename field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetEncodedFilename() string {
	if o == nil || o.EncodedFilename == nil {
		var ret string
		return ret
	}
	return *o.EncodedFilename
}

// GetEncodedFilenameOk returns a tuple with the EncodedFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetEncodedFilenameOk() (*string, bool) {
	if o == nil || o.EncodedFilename == nil {
		return nil, false
	}
	return o.EncodedFilename, true
}

// HasEncodedFilename returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasEncodedFilename() bool {
	if o != nil && o.EncodedFilename != nil {
		return true
	}

	return false
}

// SetEncodedFilename gets a reference to the given string and assigns it to the EncodedFilename field.
func (o *BTTranslationRequestParams) SetEncodedFilename(v string) {
	o.EncodedFilename = &v
}

// GetExtractAssemblyHierarchy returns the ExtractAssemblyHierarchy field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetExtractAssemblyHierarchy() bool {
	if o == nil || o.ExtractAssemblyHierarchy == nil {
		var ret bool
		return ret
	}
	return *o.ExtractAssemblyHierarchy
}

// GetExtractAssemblyHierarchyOk returns a tuple with the ExtractAssemblyHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetExtractAssemblyHierarchyOk() (*bool, bool) {
	if o == nil || o.ExtractAssemblyHierarchy == nil {
		return nil, false
	}
	return o.ExtractAssemblyHierarchy, true
}

// HasExtractAssemblyHierarchy returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasExtractAssemblyHierarchy() bool {
	if o != nil && o.ExtractAssemblyHierarchy != nil {
		return true
	}

	return false
}

// SetExtractAssemblyHierarchy gets a reference to the given bool and assigns it to the ExtractAssemblyHierarchy field.
func (o *BTTranslationRequestParams) SetExtractAssemblyHierarchy(v bool) {
	o.ExtractAssemblyHierarchy = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetFile() *os.File {
	if o == nil || o.File == nil {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetFileOk() (**os.File, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *BTTranslationRequestParams) SetFile(v *os.File) {
	o.File = &v
}

// GetFileBodyWithDetails returns the FileBodyWithDetails field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetFileBodyWithDetails() FormDataBodyPart {
	if o == nil || o.FileBodyWithDetails == nil {
		var ret FormDataBodyPart
		return ret
	}
	return *o.FileBodyWithDetails
}

// GetFileBodyWithDetailsOk returns a tuple with the FileBodyWithDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetFileBodyWithDetailsOk() (*FormDataBodyPart, bool) {
	if o == nil || o.FileBodyWithDetails == nil {
		return nil, false
	}
	return o.FileBodyWithDetails, true
}

// HasFileBodyWithDetails returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasFileBodyWithDetails() bool {
	if o != nil && o.FileBodyWithDetails != nil {
		return true
	}

	return false
}

// SetFileBodyWithDetails gets a reference to the given FormDataBodyPart and assigns it to the FileBodyWithDetails field.
func (o *BTTranslationRequestParams) SetFileBodyWithDetails(v FormDataBodyPart) {
	o.FileBodyWithDetails = &v
}

// GetFileContentLength returns the FileContentLength field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetFileContentLength() int64 {
	if o == nil || o.FileContentLength == nil {
		var ret int64
		return ret
	}
	return *o.FileContentLength
}

// GetFileContentLengthOk returns a tuple with the FileContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetFileContentLengthOk() (*int64, bool) {
	if o == nil || o.FileContentLength == nil {
		return nil, false
	}
	return o.FileContentLength, true
}

// HasFileContentLength returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasFileContentLength() bool {
	if o != nil && o.FileContentLength != nil {
		return true
	}

	return false
}

// SetFileContentLength gets a reference to the given int64 and assigns it to the FileContentLength field.
func (o *BTTranslationRequestParams) SetFileContentLength(v int64) {
	o.FileContentLength = &v
}

// GetFileDetail returns the FileDetail field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetFileDetail() FormDataContentDisposition {
	if o == nil || o.FileDetail == nil {
		var ret FormDataContentDisposition
		return ret
	}
	return *o.FileDetail
}

// GetFileDetailOk returns a tuple with the FileDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetFileDetailOk() (*FormDataContentDisposition, bool) {
	if o == nil || o.FileDetail == nil {
		return nil, false
	}
	return o.FileDetail, true
}

// HasFileDetail returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasFileDetail() bool {
	if o != nil && o.FileDetail != nil {
		return true
	}

	return false
}

// SetFileDetail gets a reference to the given FormDataContentDisposition and assigns it to the FileDetail field.
func (o *BTTranslationRequestParams) SetFileDetail(v FormDataContentDisposition) {
	o.FileDetail = &v
}

// GetFlattenAssemblies returns the FlattenAssemblies field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetFlattenAssemblies() bool {
	if o == nil || o.FlattenAssemblies == nil {
		var ret bool
		return ret
	}
	return *o.FlattenAssemblies
}

// GetFlattenAssembliesOk returns a tuple with the FlattenAssemblies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetFlattenAssembliesOk() (*bool, bool) {
	if o == nil || o.FlattenAssemblies == nil {
		return nil, false
	}
	return o.FlattenAssemblies, true
}

// HasFlattenAssemblies returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasFlattenAssemblies() bool {
	if o != nil && o.FlattenAssemblies != nil {
		return true
	}

	return false
}

// SetFlattenAssemblies gets a reference to the given bool and assigns it to the FlattenAssemblies field.
func (o *BTTranslationRequestParams) SetFlattenAssemblies(v bool) {
	o.FlattenAssemblies = &v
}

// GetFormatName returns the FormatName field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetFormatName() string {
	if o == nil || o.FormatName == nil {
		var ret string
		return ret
	}
	return *o.FormatName
}

// GetFormatNameOk returns a tuple with the FormatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetFormatNameOk() (*string, bool) {
	if o == nil || o.FormatName == nil {
		return nil, false
	}
	return o.FormatName, true
}

// HasFormatName returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasFormatName() bool {
	if o != nil && o.FormatName != nil {
		return true
	}

	return false
}

// SetFormatName gets a reference to the given string and assigns it to the FormatName field.
func (o *BTTranslationRequestParams) SetFormatName(v string) {
	o.FormatName = &v
}

// GetIsyAxisIsUp returns the IsyAxisIsUp field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetIsyAxisIsUp() bool {
	if o == nil || o.IsyAxisIsUp == nil {
		var ret bool
		return ret
	}
	return *o.IsyAxisIsUp
}

// GetIsyAxisIsUpOk returns a tuple with the IsyAxisIsUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetIsyAxisIsUpOk() (*bool, bool) {
	if o == nil || o.IsyAxisIsUp == nil {
		return nil, false
	}
	return o.IsyAxisIsUp, true
}

// HasIsyAxisIsUp returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasIsyAxisIsUp() bool {
	if o != nil && o.IsyAxisIsUp != nil {
		return true
	}

	return false
}

// SetIsyAxisIsUp gets a reference to the given bool and assigns it to the IsyAxisIsUp field.
func (o *BTTranslationRequestParams) SetIsyAxisIsUp(v bool) {
	o.IsyAxisIsUp = &v
}

// GetJoinAdjacentSurfaces returns the JoinAdjacentSurfaces field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetJoinAdjacentSurfaces() bool {
	if o == nil || o.JoinAdjacentSurfaces == nil {
		var ret bool
		return ret
	}
	return *o.JoinAdjacentSurfaces
}

// GetJoinAdjacentSurfacesOk returns a tuple with the JoinAdjacentSurfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetJoinAdjacentSurfacesOk() (*bool, bool) {
	if o == nil || o.JoinAdjacentSurfaces == nil {
		return nil, false
	}
	return o.JoinAdjacentSurfaces, true
}

// HasJoinAdjacentSurfaces returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasJoinAdjacentSurfaces() bool {
	if o != nil && o.JoinAdjacentSurfaces != nil {
		return true
	}

	return false
}

// SetJoinAdjacentSurfaces gets a reference to the given bool and assigns it to the JoinAdjacentSurfaces field.
func (o *BTTranslationRequestParams) SetJoinAdjacentSurfaces(v bool) {
	o.JoinAdjacentSurfaces = &v
}

// GetLocationElementId returns the LocationElementId field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetLocationElementId() string {
	if o == nil || o.LocationElementId == nil {
		var ret string
		return ret
	}
	return *o.LocationElementId
}

// GetLocationElementIdOk returns a tuple with the LocationElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetLocationElementIdOk() (*string, bool) {
	if o == nil || o.LocationElementId == nil {
		return nil, false
	}
	return o.LocationElementId, true
}

// HasLocationElementId returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasLocationElementId() bool {
	if o != nil && o.LocationElementId != nil {
		return true
	}

	return false
}

// SetLocationElementId gets a reference to the given string and assigns it to the LocationElementId field.
func (o *BTTranslationRequestParams) SetLocationElementId(v string) {
	o.LocationElementId = &v
}

// GetLocationGroupId returns the LocationGroupId field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetLocationGroupId() string {
	if o == nil || o.LocationGroupId == nil {
		var ret string
		return ret
	}
	return *o.LocationGroupId
}

// GetLocationGroupIdOk returns a tuple with the LocationGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetLocationGroupIdOk() (*string, bool) {
	if o == nil || o.LocationGroupId == nil {
		return nil, false
	}
	return o.LocationGroupId, true
}

// HasLocationGroupId returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasLocationGroupId() bool {
	if o != nil && o.LocationGroupId != nil {
		return true
	}

	return false
}

// SetLocationGroupId gets a reference to the given string and assigns it to the LocationGroupId field.
func (o *BTTranslationRequestParams) SetLocationGroupId(v string) {
	o.LocationGroupId = &v
}

// GetLocationPosition returns the LocationPosition field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetLocationPosition() int32 {
	if o == nil || o.LocationPosition == nil {
		var ret int32
		return ret
	}
	return *o.LocationPosition
}

// GetLocationPositionOk returns a tuple with the LocationPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetLocationPositionOk() (*int32, bool) {
	if o == nil || o.LocationPosition == nil {
		return nil, false
	}
	return o.LocationPosition, true
}

// HasLocationPosition returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasLocationPosition() bool {
	if o != nil && o.LocationPosition != nil {
		return true
	}

	return false
}

// SetLocationPosition gets a reference to the given int32 and assigns it to the LocationPosition field.
func (o *BTTranslationRequestParams) SetLocationPosition(v int32) {
	o.LocationPosition = &v
}

// GetNotifyUser returns the NotifyUser field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetNotifyUser() bool {
	if o == nil || o.NotifyUser == nil {
		var ret bool
		return ret
	}
	return *o.NotifyUser
}

// GetNotifyUserOk returns a tuple with the NotifyUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetNotifyUserOk() (*bool, bool) {
	if o == nil || o.NotifyUser == nil {
		return nil, false
	}
	return o.NotifyUser, true
}

// HasNotifyUser returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasNotifyUser() bool {
	if o != nil && o.NotifyUser != nil {
		return true
	}

	return false
}

// SetNotifyUser gets a reference to the given bool and assigns it to the NotifyUser field.
func (o *BTTranslationRequestParams) SetNotifyUser(v bool) {
	o.NotifyUser = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTTranslationRequestParams) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetOwnerType() string {
	if o == nil || o.OwnerType == nil {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetOwnerTypeOk() (*string, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *BTTranslationRequestParams) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTTranslationRequestParams) SetParentId(v string) {
	o.ParentId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTTranslationRequestParams) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *BTTranslationRequestParams) SetPublic(v bool) {
	o.Public = &v
}

// GetSplitAssembliesIntoMultipleDocuments returns the SplitAssembliesIntoMultipleDocuments field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetSplitAssembliesIntoMultipleDocuments() bool {
	if o == nil || o.SplitAssembliesIntoMultipleDocuments == nil {
		var ret bool
		return ret
	}
	return *o.SplitAssembliesIntoMultipleDocuments
}

// GetSplitAssembliesIntoMultipleDocumentsOk returns a tuple with the SplitAssembliesIntoMultipleDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetSplitAssembliesIntoMultipleDocumentsOk() (*bool, bool) {
	if o == nil || o.SplitAssembliesIntoMultipleDocuments == nil {
		return nil, false
	}
	return o.SplitAssembliesIntoMultipleDocuments, true
}

// HasSplitAssembliesIntoMultipleDocuments returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasSplitAssembliesIntoMultipleDocuments() bool {
	if o != nil && o.SplitAssembliesIntoMultipleDocuments != nil {
		return true
	}

	return false
}

// SetSplitAssembliesIntoMultipleDocuments gets a reference to the given bool and assigns it to the SplitAssembliesIntoMultipleDocuments field.
func (o *BTTranslationRequestParams) SetSplitAssembliesIntoMultipleDocuments(v bool) {
	o.SplitAssembliesIntoMultipleDocuments = &v
}

// GetStoreInDocument returns the StoreInDocument field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetStoreInDocument() bool {
	if o == nil || o.StoreInDocument == nil {
		var ret bool
		return ret
	}
	return *o.StoreInDocument
}

// GetStoreInDocumentOk returns a tuple with the StoreInDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetStoreInDocumentOk() (*bool, bool) {
	if o == nil || o.StoreInDocument == nil {
		return nil, false
	}
	return o.StoreInDocument, true
}

// HasStoreInDocument returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasStoreInDocument() bool {
	if o != nil && o.StoreInDocument != nil {
		return true
	}

	return false
}

// SetStoreInDocument gets a reference to the given bool and assigns it to the StoreInDocument field.
func (o *BTTranslationRequestParams) SetStoreInDocument(v bool) {
	o.StoreInDocument = &v
}

// GetTranslate returns the Translate field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetTranslate() bool {
	if o == nil || o.Translate == nil {
		var ret bool
		return ret
	}
	return *o.Translate
}

// GetTranslateOk returns a tuple with the Translate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetTranslateOk() (*bool, bool) {
	if o == nil || o.Translate == nil {
		return nil, false
	}
	return o.Translate, true
}

// HasTranslate returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasTranslate() bool {
	if o != nil && o.Translate != nil {
		return true
	}

	return false
}

// SetTranslate gets a reference to the given bool and assigns it to the Translate field.
func (o *BTTranslationRequestParams) SetTranslate(v bool) {
	o.Translate = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *BTTranslationRequestParams) SetUnit(v string) {
	o.Unit = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetUploadId() string {
	if o == nil || o.UploadId == nil {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetUploadIdOk() (*string, bool) {
	if o == nil || o.UploadId == nil {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasUploadId() bool {
	if o != nil && o.UploadId != nil {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *BTTranslationRequestParams) SetUploadId(v string) {
	o.UploadId = &v
}

// GetVersionString returns the VersionString field value if set, zero value otherwise.
func (o *BTTranslationRequestParams) GetVersionString() string {
	if o == nil || o.VersionString == nil {
		var ret string
		return ret
	}
	return *o.VersionString
}

// GetVersionStringOk returns a tuple with the VersionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTranslationRequestParams) GetVersionStringOk() (*string, bool) {
	if o == nil || o.VersionString == nil {
		return nil, false
	}
	return o.VersionString, true
}

// HasVersionString returns a boolean if a field has been set.
func (o *BTTranslationRequestParams) HasVersionString() bool {
	if o != nil && o.VersionString != nil {
		return true
	}

	return false
}

// SetVersionString gets a reference to the given string and assigns it to the VersionString field.
func (o *BTTranslationRequestParams) SetVersionString(v string) {
	o.VersionString = &v
}

func (o BTTranslationRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowFaultyParts != nil {
		toSerialize["allowFaultyParts"] = o.AllowFaultyParts
	}
	if o.CreateComposite != nil {
		toSerialize["createComposite"] = o.CreateComposite
	}
	if o.CreateDrawingIfPossible != nil {
		toSerialize["createDrawingIfPossible"] = o.CreateDrawingIfPossible
	}
	if o.EncodedFilename != nil {
		toSerialize["encodedFilename"] = o.EncodedFilename
	}
	if o.ExtractAssemblyHierarchy != nil {
		toSerialize["extractAssemblyHierarchy"] = o.ExtractAssemblyHierarchy
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.FileBodyWithDetails != nil {
		toSerialize["fileBodyWithDetails"] = o.FileBodyWithDetails
	}
	if o.FileContentLength != nil {
		toSerialize["fileContentLength"] = o.FileContentLength
	}
	if o.FileDetail != nil {
		toSerialize["fileDetail"] = o.FileDetail
	}
	if o.FlattenAssemblies != nil {
		toSerialize["flattenAssemblies"] = o.FlattenAssemblies
	}
	if o.FormatName != nil {
		toSerialize["formatName"] = o.FormatName
	}
	if o.IsyAxisIsUp != nil {
		toSerialize["isyAxisIsUp"] = o.IsyAxisIsUp
	}
	if o.JoinAdjacentSurfaces != nil {
		toSerialize["joinAdjacentSurfaces"] = o.JoinAdjacentSurfaces
	}
	if o.LocationElementId != nil {
		toSerialize["locationElementId"] = o.LocationElementId
	}
	if o.LocationGroupId != nil {
		toSerialize["locationGroupId"] = o.LocationGroupId
	}
	if o.LocationPosition != nil {
		toSerialize["locationPosition"] = o.LocationPosition
	}
	if o.NotifyUser != nil {
		toSerialize["notifyUser"] = o.NotifyUser
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.SplitAssembliesIntoMultipleDocuments != nil {
		toSerialize["splitAssembliesIntoMultipleDocuments"] = o.SplitAssembliesIntoMultipleDocuments
	}
	if o.StoreInDocument != nil {
		toSerialize["storeInDocument"] = o.StoreInDocument
	}
	if o.Translate != nil {
		toSerialize["translate"] = o.Translate
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.UploadId != nil {
		toSerialize["uploadId"] = o.UploadId
	}
	if o.VersionString != nil {
		toSerialize["versionString"] = o.VersionString
	}
	return json.Marshal(toSerialize)
}

type NullableBTTranslationRequestParams struct {
	value *BTTranslationRequestParams
	isSet bool
}

func (v NullableBTTranslationRequestParams) Get() *BTTranslationRequestParams {
	return v.value
}

func (v *NullableBTTranslationRequestParams) Set(val *BTTranslationRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTranslationRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTranslationRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTranslationRequestParams(val *BTTranslationRequestParams) *NullableBTTranslationRequestParams {
	return &NullableBTTranslationRequestParams{value: val, isSet: true}
}

func (v NullableBTTranslationRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTranslationRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
