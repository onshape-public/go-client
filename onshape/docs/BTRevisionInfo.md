# BTRevisionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvers** | Pointer to [**[]BTRevisionApproverInfo**](BTRevisionApproverInfo.md) | The users who approved the release package that created this revision. | [optional] 
**AutoObsoletionReleaseId** | Pointer to **string** |  | [optional] 
**AutoObsoletionReleaseName** | Pointer to **string** |  | [optional] 
**CanChangeType** | Pointer to **bool** | Whether the revision can change object type. Used in reuse part number flow. | [optional] [default to false]
**CanExport** | Pointer to **bool** |  | [optional] 
**CompanyId** | Pointer to **string** | The company or enterprise ID that owns the resource. | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**Description** | Pointer to **string** | The Revision Description metadata property if revision is of a drawing. | [optional] 
**DocumentId** | Pointer to **string** | The document that contains the revision object. | [optional] 
**DocumentName** | Pointer to **string** | The name of the document that contains the revision object. | [optional] 
**DocumentState** | Pointer to **int32** | The state of document containing this revision. Used in reuse part number flow | [optional] 
**ElementId** | Pointer to **string** | The element that contains the revision object. | [optional] 
**ElementType** | Pointer to **int32** | The type of item Element Type. Must be one of: &#x60;-1&#x60;: Unknown, &#x60;0&#x60;: Part Studio, &#x60;1&#x60;: Assembly, &#x60;2&#x60;: Drawing. &#x60;4&#x60; : Blob, &#x60;8&#x60;: Variable Studio | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FlatPartInsertableId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**InsertableId** | Pointer to **string** |  | [optional] 
**IsObsolete** | Pointer to **bool** | Whether the revision has been obsoleted. | [optional] 
**IsRereleasable** | Pointer to **bool** | If true, the revision can be created again. | [optional] 
**IsTranslatable** | Pointer to **bool** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**NextRevisionId** | Pointer to **string** | The next revision if applicable. null for the latest revision. | [optional] 
**ObsoletionPackageId** | Pointer to **string** | The OBSOLETION release package that obsoleted this revision if applicable. | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** | The part number with which the object was revised. | [optional] 
**PreviousRevisionId** | Pointer to **string** | The previous revision if applicable. null for first revision. | [optional] 
**ReleaseCreatedDate** | Pointer to **JSONTime** |  | [optional] 
**ReleaseId** | Pointer to **string** | The release package that created this revision. | [optional] 
**ReleaseName** | Pointer to **string** | The name of the release package that created this item. | [optional] 
**ReleasedBy** | Pointer to [**BTUserSummaryInfo**](BTUserSummaryInfo.md) |  | [optional] 
**Revision** | Pointer to **string** | The id of the revision. | [optional] 
**RevisionRuleId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** | The version of the document that contains this revision. | [optional] 
**VersionName** | Pointer to **string** | The name of the version of the document that contains this revision. | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTRevisionInfo

`func NewBTRevisionInfo() *BTRevisionInfo`

NewBTRevisionInfo instantiates a new BTRevisionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTRevisionInfoWithDefaults

`func NewBTRevisionInfoWithDefaults() *BTRevisionInfo`

NewBTRevisionInfoWithDefaults instantiates a new BTRevisionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovers

`func (o *BTRevisionInfo) GetApprovers() []BTRevisionApproverInfo`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *BTRevisionInfo) GetApproversOk() (*[]BTRevisionApproverInfo, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *BTRevisionInfo) SetApprovers(v []BTRevisionApproverInfo)`

SetApprovers sets Approvers field to given value.

### HasApprovers

`func (o *BTRevisionInfo) HasApprovers() bool`

HasApprovers returns a boolean if a field has been set.

### GetAutoObsoletionReleaseId

`func (o *BTRevisionInfo) GetAutoObsoletionReleaseId() string`

GetAutoObsoletionReleaseId returns the AutoObsoletionReleaseId field if non-nil, zero value otherwise.

### GetAutoObsoletionReleaseIdOk

`func (o *BTRevisionInfo) GetAutoObsoletionReleaseIdOk() (*string, bool)`

GetAutoObsoletionReleaseIdOk returns a tuple with the AutoObsoletionReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoObsoletionReleaseId

`func (o *BTRevisionInfo) SetAutoObsoletionReleaseId(v string)`

SetAutoObsoletionReleaseId sets AutoObsoletionReleaseId field to given value.

### HasAutoObsoletionReleaseId

`func (o *BTRevisionInfo) HasAutoObsoletionReleaseId() bool`

HasAutoObsoletionReleaseId returns a boolean if a field has been set.

### GetAutoObsoletionReleaseName

`func (o *BTRevisionInfo) GetAutoObsoletionReleaseName() string`

GetAutoObsoletionReleaseName returns the AutoObsoletionReleaseName field if non-nil, zero value otherwise.

### GetAutoObsoletionReleaseNameOk

`func (o *BTRevisionInfo) GetAutoObsoletionReleaseNameOk() (*string, bool)`

GetAutoObsoletionReleaseNameOk returns a tuple with the AutoObsoletionReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoObsoletionReleaseName

`func (o *BTRevisionInfo) SetAutoObsoletionReleaseName(v string)`

SetAutoObsoletionReleaseName sets AutoObsoletionReleaseName field to given value.

### HasAutoObsoletionReleaseName

`func (o *BTRevisionInfo) HasAutoObsoletionReleaseName() bool`

HasAutoObsoletionReleaseName returns a boolean if a field has been set.

### GetCanChangeType

`func (o *BTRevisionInfo) GetCanChangeType() bool`

GetCanChangeType returns the CanChangeType field if non-nil, zero value otherwise.

### GetCanChangeTypeOk

`func (o *BTRevisionInfo) GetCanChangeTypeOk() (*bool, bool)`

GetCanChangeTypeOk returns a tuple with the CanChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanChangeType

`func (o *BTRevisionInfo) SetCanChangeType(v bool)`

SetCanChangeType sets CanChangeType field to given value.

### HasCanChangeType

`func (o *BTRevisionInfo) HasCanChangeType() bool`

HasCanChangeType returns a boolean if a field has been set.

### GetCanExport

`func (o *BTRevisionInfo) GetCanExport() bool`

GetCanExport returns the CanExport field if non-nil, zero value otherwise.

### GetCanExportOk

`func (o *BTRevisionInfo) GetCanExportOk() (*bool, bool)`

GetCanExportOk returns a tuple with the CanExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanExport

`func (o *BTRevisionInfo) SetCanExport(v bool)`

SetCanExport sets CanExport field to given value.

### HasCanExport

`func (o *BTRevisionInfo) HasCanExport() bool`

HasCanExport returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTRevisionInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTRevisionInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTRevisionInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTRevisionInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTRevisionInfo) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTRevisionInfo) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTRevisionInfo) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTRevisionInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTRevisionInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTRevisionInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTRevisionInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTRevisionInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *BTRevisionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTRevisionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTRevisionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTRevisionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTRevisionInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTRevisionInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTRevisionInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTRevisionInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentName

`func (o *BTRevisionInfo) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *BTRevisionInfo) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *BTRevisionInfo) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *BTRevisionInfo) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetDocumentState

`func (o *BTRevisionInfo) GetDocumentState() int32`

GetDocumentState returns the DocumentState field if non-nil, zero value otherwise.

### GetDocumentStateOk

`func (o *BTRevisionInfo) GetDocumentStateOk() (*int32, bool)`

GetDocumentStateOk returns a tuple with the DocumentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentState

`func (o *BTRevisionInfo) SetDocumentState(v int32)`

SetDocumentState sets DocumentState field to given value.

### HasDocumentState

`func (o *BTRevisionInfo) HasDocumentState() bool`

HasDocumentState returns a boolean if a field has been set.

### GetElementId

`func (o *BTRevisionInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTRevisionInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTRevisionInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTRevisionInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTRevisionInfo) GetElementType() int32`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTRevisionInfo) GetElementTypeOk() (*int32, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTRevisionInfo) SetElementType(v int32)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTRevisionInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BTRevisionInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BTRevisionInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BTRevisionInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BTRevisionInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFlatPartInsertableId

`func (o *BTRevisionInfo) GetFlatPartInsertableId() string`

GetFlatPartInsertableId returns the FlatPartInsertableId field if non-nil, zero value otherwise.

### GetFlatPartInsertableIdOk

`func (o *BTRevisionInfo) GetFlatPartInsertableIdOk() (*string, bool)`

GetFlatPartInsertableIdOk returns a tuple with the FlatPartInsertableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatPartInsertableId

`func (o *BTRevisionInfo) SetFlatPartInsertableId(v string)`

SetFlatPartInsertableId sets FlatPartInsertableId field to given value.

### HasFlatPartInsertableId

`func (o *BTRevisionInfo) HasFlatPartInsertableId() bool`

HasFlatPartInsertableId returns a boolean if a field has been set.

### GetHref

`func (o *BTRevisionInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTRevisionInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTRevisionInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTRevisionInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTRevisionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTRevisionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTRevisionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTRevisionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertableId

`func (o *BTRevisionInfo) GetInsertableId() string`

GetInsertableId returns the InsertableId field if non-nil, zero value otherwise.

### GetInsertableIdOk

`func (o *BTRevisionInfo) GetInsertableIdOk() (*string, bool)`

GetInsertableIdOk returns a tuple with the InsertableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertableId

`func (o *BTRevisionInfo) SetInsertableId(v string)`

SetInsertableId sets InsertableId field to given value.

### HasInsertableId

`func (o *BTRevisionInfo) HasInsertableId() bool`

HasInsertableId returns a boolean if a field has been set.

### GetIsObsolete

`func (o *BTRevisionInfo) GetIsObsolete() bool`

GetIsObsolete returns the IsObsolete field if non-nil, zero value otherwise.

### GetIsObsoleteOk

`func (o *BTRevisionInfo) GetIsObsoleteOk() (*bool, bool)`

GetIsObsoleteOk returns a tuple with the IsObsolete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObsolete

`func (o *BTRevisionInfo) SetIsObsolete(v bool)`

SetIsObsolete sets IsObsolete field to given value.

### HasIsObsolete

`func (o *BTRevisionInfo) HasIsObsolete() bool`

HasIsObsolete returns a boolean if a field has been set.

### GetIsRereleasable

`func (o *BTRevisionInfo) GetIsRereleasable() bool`

GetIsRereleasable returns the IsRereleasable field if non-nil, zero value otherwise.

### GetIsRereleasableOk

`func (o *BTRevisionInfo) GetIsRereleasableOk() (*bool, bool)`

GetIsRereleasableOk returns a tuple with the IsRereleasable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRereleasable

`func (o *BTRevisionInfo) SetIsRereleasable(v bool)`

SetIsRereleasable sets IsRereleasable field to given value.

### HasIsRereleasable

`func (o *BTRevisionInfo) HasIsRereleasable() bool`

HasIsRereleasable returns a boolean if a field has been set.

### GetIsTranslatable

`func (o *BTRevisionInfo) GetIsTranslatable() bool`

GetIsTranslatable returns the IsTranslatable field if non-nil, zero value otherwise.

### GetIsTranslatableOk

`func (o *BTRevisionInfo) GetIsTranslatableOk() (*bool, bool)`

GetIsTranslatableOk returns a tuple with the IsTranslatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTranslatable

`func (o *BTRevisionInfo) SetIsTranslatable(v bool)`

SetIsTranslatable sets IsTranslatable field to given value.

### HasIsTranslatable

`func (o *BTRevisionInfo) HasIsTranslatable() bool`

HasIsTranslatable returns a boolean if a field has been set.

### GetMimeType

`func (o *BTRevisionInfo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTRevisionInfo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTRevisionInfo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTRevisionInfo) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetName

`func (o *BTRevisionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTRevisionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTRevisionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTRevisionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRevisionId

`func (o *BTRevisionInfo) GetNextRevisionId() string`

GetNextRevisionId returns the NextRevisionId field if non-nil, zero value otherwise.

### GetNextRevisionIdOk

`func (o *BTRevisionInfo) GetNextRevisionIdOk() (*string, bool)`

GetNextRevisionIdOk returns a tuple with the NextRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRevisionId

`func (o *BTRevisionInfo) SetNextRevisionId(v string)`

SetNextRevisionId sets NextRevisionId field to given value.

### HasNextRevisionId

`func (o *BTRevisionInfo) HasNextRevisionId() bool`

HasNextRevisionId returns a boolean if a field has been set.

### GetObsoletionPackageId

`func (o *BTRevisionInfo) GetObsoletionPackageId() string`

GetObsoletionPackageId returns the ObsoletionPackageId field if non-nil, zero value otherwise.

### GetObsoletionPackageIdOk

`func (o *BTRevisionInfo) GetObsoletionPackageIdOk() (*string, bool)`

GetObsoletionPackageIdOk returns a tuple with the ObsoletionPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionPackageId

`func (o *BTRevisionInfo) SetObsoletionPackageId(v string)`

SetObsoletionPackageId sets ObsoletionPackageId field to given value.

### HasObsoletionPackageId

`func (o *BTRevisionInfo) HasObsoletionPackageId() bool`

HasObsoletionPackageId returns a boolean if a field has been set.

### GetPartId

`func (o *BTRevisionInfo) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTRevisionInfo) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTRevisionInfo) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTRevisionInfo) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTRevisionInfo) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTRevisionInfo) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTRevisionInfo) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTRevisionInfo) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTRevisionInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTRevisionInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTRevisionInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTRevisionInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPreviousRevisionId

`func (o *BTRevisionInfo) GetPreviousRevisionId() string`

GetPreviousRevisionId returns the PreviousRevisionId field if non-nil, zero value otherwise.

### GetPreviousRevisionIdOk

`func (o *BTRevisionInfo) GetPreviousRevisionIdOk() (*string, bool)`

GetPreviousRevisionIdOk returns a tuple with the PreviousRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRevisionId

`func (o *BTRevisionInfo) SetPreviousRevisionId(v string)`

SetPreviousRevisionId sets PreviousRevisionId field to given value.

### HasPreviousRevisionId

`func (o *BTRevisionInfo) HasPreviousRevisionId() bool`

HasPreviousRevisionId returns a boolean if a field has been set.

### GetReleaseCreatedDate

`func (o *BTRevisionInfo) GetReleaseCreatedDate() JSONTime`

GetReleaseCreatedDate returns the ReleaseCreatedDate field if non-nil, zero value otherwise.

### GetReleaseCreatedDateOk

`func (o *BTRevisionInfo) GetReleaseCreatedDateOk() (*JSONTime, bool)`

GetReleaseCreatedDateOk returns a tuple with the ReleaseCreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCreatedDate

`func (o *BTRevisionInfo) SetReleaseCreatedDate(v JSONTime)`

SetReleaseCreatedDate sets ReleaseCreatedDate field to given value.

### HasReleaseCreatedDate

`func (o *BTRevisionInfo) HasReleaseCreatedDate() bool`

HasReleaseCreatedDate returns a boolean if a field has been set.

### GetReleaseId

`func (o *BTRevisionInfo) GetReleaseId() string`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *BTRevisionInfo) GetReleaseIdOk() (*string, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *BTRevisionInfo) SetReleaseId(v string)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *BTRevisionInfo) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### GetReleaseName

`func (o *BTRevisionInfo) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *BTRevisionInfo) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *BTRevisionInfo) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *BTRevisionInfo) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetReleasedBy

`func (o *BTRevisionInfo) GetReleasedBy() BTUserSummaryInfo`

GetReleasedBy returns the ReleasedBy field if non-nil, zero value otherwise.

### GetReleasedByOk

`func (o *BTRevisionInfo) GetReleasedByOk() (*BTUserSummaryInfo, bool)`

GetReleasedByOk returns a tuple with the ReleasedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedBy

`func (o *BTRevisionInfo) SetReleasedBy(v BTUserSummaryInfo)`

SetReleasedBy sets ReleasedBy field to given value.

### HasReleasedBy

`func (o *BTRevisionInfo) HasReleasedBy() bool`

HasReleasedBy returns a boolean if a field has been set.

### GetRevision

`func (o *BTRevisionInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTRevisionInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTRevisionInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTRevisionInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRevisionRuleId

`func (o *BTRevisionInfo) GetRevisionRuleId() string`

GetRevisionRuleId returns the RevisionRuleId field if non-nil, zero value otherwise.

### GetRevisionRuleIdOk

`func (o *BTRevisionInfo) GetRevisionRuleIdOk() (*string, bool)`

GetRevisionRuleIdOk returns a tuple with the RevisionRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionRuleId

`func (o *BTRevisionInfo) SetRevisionRuleId(v string)`

SetRevisionRuleId sets RevisionRuleId field to given value.

### HasRevisionRuleId

`func (o *BTRevisionInfo) HasRevisionRuleId() bool`

HasRevisionRuleId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTRevisionInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTRevisionInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTRevisionInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTRevisionInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetVersionName

`func (o *BTRevisionInfo) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *BTRevisionInfo) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *BTRevisionInfo) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *BTRevisionInfo) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.

### GetViewRef

`func (o *BTRevisionInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTRevisionInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTRevisionInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTRevisionInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


