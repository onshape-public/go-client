# BTExportModelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AngleTolerance** | Pointer to **float64** |  | [optional] 
**BatchFlatPatterns** | Pointer to **bool** |  | [optional] 
**ChordTolerance** | Pointer to **float64** |  | [optional] 
**CloudObjectId** | Pointer to **string** |  | [optional] 
**CloudStorageAccountId** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**DeepSearchForForeignData** | Pointer to **bool** |  | [optional] 
**DestinationName** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DocumentVersionId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementIds** | Pointer to **string** |  | [optional] 
**EmailLink** | Pointer to **bool** |  | [optional] 
**EmailMessage** | Pointer to **string** |  | [optional] 
**EmailSubject** | Pointer to **string** |  | [optional] 
**EmailTo** | Pointer to **string** |  | [optional] 
**ExtractToS3** | Pointer to **bool** |  | [optional] 
**FeatureIds** | Pointer to **string** |  | [optional] 
**Flatten** | Pointer to **bool** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**FromUserId** | Pointer to **string** |  | [optional] 
**Grouping** | Pointer to **string** |  | [optional] 
**IncludeBendCenterlines** | Pointer to **bool** |  | [optional] 
**IncludeBendLines** | Pointer to **bool** |  | [optional] 
**IncludeCustomProperties** | Pointer to **bool** |  | [optional] 
**IncludeCustomPropertiesData** | Pointer to **bool** |  | [optional] 
**IncludeExportIds** | Pointer to **bool** |  | [optional] 
**IncludeForeignData** | Pointer to **bool** |  | [optional] 
**IncludeItemsData** | Pointer to **bool** |  | [optional] 
**IncludeLinkedDocuments** | Pointer to **bool** |  | [optional] 
**IncludeReleaseManagementData** | Pointer to **bool** |  | [optional] 
**IncludeSketches** | Pointer to **bool** |  | [optional] 
**IncludeStd** | Pointer to **bool** |  | [optional] 
**IsPartingOut** | Pointer to **bool** |  | [optional] 
**LinkDocumentId** | Pointer to **string** |  | [optional] 
**LinkDocumentWorkspaceId** | Pointer to **string** |  | [optional] 
**MaxFacetWidth** | Pointer to **float64** |  | [optional] 
**Microversion** | Pointer to **string** |  | [optional] 
**MinFacetWidth** | Pointer to **float64** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**PartIds** | Pointer to **string** |  | [optional] 
**PartQuery** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordRequired** | Pointer to **bool** |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**Scale** | Pointer to **float64** |  | [optional] 
**SendCopyToMe** | Pointer to **bool** |  | [optional] 
**SheetMetalFlat** | Pointer to **bool** |  | [optional] 
**SplinesAsPolylines** | Pointer to **bool** |  | [optional] 
**StoreInDocument** | Pointer to **bool** |  | [optional] 
**TriggerAutoDownload** | Pointer to **bool** |  | [optional] 
**Units** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ValidForDays** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**View** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**ZipSingleFileOutput** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTExportModelParams

`func NewBTExportModelParams() *BTExportModelParams`

NewBTExportModelParams instantiates a new BTExportModelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTExportModelParamsWithDefaults

`func NewBTExportModelParamsWithDefaults() *BTExportModelParams`

NewBTExportModelParamsWithDefaults instantiates a new BTExportModelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAngleTolerance

`func (o *BTExportModelParams) GetAngleTolerance() float64`

GetAngleTolerance returns the AngleTolerance field if non-nil, zero value otherwise.

### GetAngleToleranceOk

`func (o *BTExportModelParams) GetAngleToleranceOk() (*float64, bool)`

GetAngleToleranceOk returns a tuple with the AngleTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngleTolerance

`func (o *BTExportModelParams) SetAngleTolerance(v float64)`

SetAngleTolerance sets AngleTolerance field to given value.

### HasAngleTolerance

`func (o *BTExportModelParams) HasAngleTolerance() bool`

HasAngleTolerance returns a boolean if a field has been set.

### GetBatchFlatPatterns

`func (o *BTExportModelParams) GetBatchFlatPatterns() bool`

GetBatchFlatPatterns returns the BatchFlatPatterns field if non-nil, zero value otherwise.

### GetBatchFlatPatternsOk

`func (o *BTExportModelParams) GetBatchFlatPatternsOk() (*bool, bool)`

GetBatchFlatPatternsOk returns a tuple with the BatchFlatPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFlatPatterns

`func (o *BTExportModelParams) SetBatchFlatPatterns(v bool)`

SetBatchFlatPatterns sets BatchFlatPatterns field to given value.

### HasBatchFlatPatterns

`func (o *BTExportModelParams) HasBatchFlatPatterns() bool`

HasBatchFlatPatterns returns a boolean if a field has been set.

### GetChordTolerance

`func (o *BTExportModelParams) GetChordTolerance() float64`

GetChordTolerance returns the ChordTolerance field if non-nil, zero value otherwise.

### GetChordToleranceOk

`func (o *BTExportModelParams) GetChordToleranceOk() (*float64, bool)`

GetChordToleranceOk returns a tuple with the ChordTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChordTolerance

`func (o *BTExportModelParams) SetChordTolerance(v float64)`

SetChordTolerance sets ChordTolerance field to given value.

### HasChordTolerance

`func (o *BTExportModelParams) HasChordTolerance() bool`

HasChordTolerance returns a boolean if a field has been set.

### GetCloudObjectId

`func (o *BTExportModelParams) GetCloudObjectId() string`

GetCloudObjectId returns the CloudObjectId field if non-nil, zero value otherwise.

### GetCloudObjectIdOk

`func (o *BTExportModelParams) GetCloudObjectIdOk() (*string, bool)`

GetCloudObjectIdOk returns a tuple with the CloudObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudObjectId

`func (o *BTExportModelParams) SetCloudObjectId(v string)`

SetCloudObjectId sets CloudObjectId field to given value.

### HasCloudObjectId

`func (o *BTExportModelParams) HasCloudObjectId() bool`

HasCloudObjectId returns a boolean if a field has been set.

### GetCloudStorageAccountId

`func (o *BTExportModelParams) GetCloudStorageAccountId() string`

GetCloudStorageAccountId returns the CloudStorageAccountId field if non-nil, zero value otherwise.

### GetCloudStorageAccountIdOk

`func (o *BTExportModelParams) GetCloudStorageAccountIdOk() (*string, bool)`

GetCloudStorageAccountIdOk returns a tuple with the CloudStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageAccountId

`func (o *BTExportModelParams) SetCloudStorageAccountId(v string)`

SetCloudStorageAccountId sets CloudStorageAccountId field to given value.

### HasCloudStorageAccountId

`func (o *BTExportModelParams) HasCloudStorageAccountId() bool`

HasCloudStorageAccountId returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTExportModelParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTExportModelParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTExportModelParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTExportModelParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDeepSearchForForeignData

`func (o *BTExportModelParams) GetDeepSearchForForeignData() bool`

GetDeepSearchForForeignData returns the DeepSearchForForeignData field if non-nil, zero value otherwise.

### GetDeepSearchForForeignDataOk

`func (o *BTExportModelParams) GetDeepSearchForForeignDataOk() (*bool, bool)`

GetDeepSearchForForeignDataOk returns a tuple with the DeepSearchForForeignData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepSearchForForeignData

`func (o *BTExportModelParams) SetDeepSearchForForeignData(v bool)`

SetDeepSearchForForeignData sets DeepSearchForForeignData field to given value.

### HasDeepSearchForForeignData

`func (o *BTExportModelParams) HasDeepSearchForForeignData() bool`

HasDeepSearchForForeignData returns a boolean if a field has been set.

### GetDestinationName

`func (o *BTExportModelParams) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *BTExportModelParams) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *BTExportModelParams) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *BTExportModelParams) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTExportModelParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTExportModelParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTExportModelParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTExportModelParams) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentVersionId

`func (o *BTExportModelParams) GetDocumentVersionId() string`

GetDocumentVersionId returns the DocumentVersionId field if non-nil, zero value otherwise.

### GetDocumentVersionIdOk

`func (o *BTExportModelParams) GetDocumentVersionIdOk() (*string, bool)`

GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVersionId

`func (o *BTExportModelParams) SetDocumentVersionId(v string)`

SetDocumentVersionId sets DocumentVersionId field to given value.

### HasDocumentVersionId

`func (o *BTExportModelParams) HasDocumentVersionId() bool`

HasDocumentVersionId returns a boolean if a field has been set.

### GetElementId

`func (o *BTExportModelParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTExportModelParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTExportModelParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTExportModelParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementIds

`func (o *BTExportModelParams) GetElementIds() string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *BTExportModelParams) GetElementIdsOk() (*string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *BTExportModelParams) SetElementIds(v string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *BTExportModelParams) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### GetEmailLink

`func (o *BTExportModelParams) GetEmailLink() bool`

GetEmailLink returns the EmailLink field if non-nil, zero value otherwise.

### GetEmailLinkOk

`func (o *BTExportModelParams) GetEmailLinkOk() (*bool, bool)`

GetEmailLinkOk returns a tuple with the EmailLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLink

`func (o *BTExportModelParams) SetEmailLink(v bool)`

SetEmailLink sets EmailLink field to given value.

### HasEmailLink

`func (o *BTExportModelParams) HasEmailLink() bool`

HasEmailLink returns a boolean if a field has been set.

### GetEmailMessage

`func (o *BTExportModelParams) GetEmailMessage() string`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *BTExportModelParams) GetEmailMessageOk() (*string, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *BTExportModelParams) SetEmailMessage(v string)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *BTExportModelParams) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.

### GetEmailSubject

`func (o *BTExportModelParams) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *BTExportModelParams) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *BTExportModelParams) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *BTExportModelParams) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailTo

`func (o *BTExportModelParams) GetEmailTo() string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *BTExportModelParams) GetEmailToOk() (*string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *BTExportModelParams) SetEmailTo(v string)`

SetEmailTo sets EmailTo field to given value.

### HasEmailTo

`func (o *BTExportModelParams) HasEmailTo() bool`

HasEmailTo returns a boolean if a field has been set.

### GetExtractToS3

`func (o *BTExportModelParams) GetExtractToS3() bool`

GetExtractToS3 returns the ExtractToS3 field if non-nil, zero value otherwise.

### GetExtractToS3Ok

`func (o *BTExportModelParams) GetExtractToS3Ok() (*bool, bool)`

GetExtractToS3Ok returns a tuple with the ExtractToS3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractToS3

`func (o *BTExportModelParams) SetExtractToS3(v bool)`

SetExtractToS3 sets ExtractToS3 field to given value.

### HasExtractToS3

`func (o *BTExportModelParams) HasExtractToS3() bool`

HasExtractToS3 returns a boolean if a field has been set.

### GetFeatureIds

`func (o *BTExportModelParams) GetFeatureIds() string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *BTExportModelParams) GetFeatureIdsOk() (*string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *BTExportModelParams) SetFeatureIds(v string)`

SetFeatureIds sets FeatureIds field to given value.

### HasFeatureIds

`func (o *BTExportModelParams) HasFeatureIds() bool`

HasFeatureIds returns a boolean if a field has been set.

### GetFlatten

`func (o *BTExportModelParams) GetFlatten() bool`

GetFlatten returns the Flatten field if non-nil, zero value otherwise.

### GetFlattenOk

`func (o *BTExportModelParams) GetFlattenOk() (*bool, bool)`

GetFlattenOk returns a tuple with the Flatten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatten

`func (o *BTExportModelParams) SetFlatten(v bool)`

SetFlatten sets Flatten field to given value.

### HasFlatten

`func (o *BTExportModelParams) HasFlatten() bool`

HasFlatten returns a boolean if a field has been set.

### GetFormat

`func (o *BTExportModelParams) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BTExportModelParams) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BTExportModelParams) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *BTExportModelParams) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFromUserId

`func (o *BTExportModelParams) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *BTExportModelParams) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *BTExportModelParams) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *BTExportModelParams) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetGrouping

`func (o *BTExportModelParams) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BTExportModelParams) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BTExportModelParams) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *BTExportModelParams) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetIncludeBendCenterlines

`func (o *BTExportModelParams) GetIncludeBendCenterlines() bool`

GetIncludeBendCenterlines returns the IncludeBendCenterlines field if non-nil, zero value otherwise.

### GetIncludeBendCenterlinesOk

`func (o *BTExportModelParams) GetIncludeBendCenterlinesOk() (*bool, bool)`

GetIncludeBendCenterlinesOk returns a tuple with the IncludeBendCenterlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBendCenterlines

`func (o *BTExportModelParams) SetIncludeBendCenterlines(v bool)`

SetIncludeBendCenterlines sets IncludeBendCenterlines field to given value.

### HasIncludeBendCenterlines

`func (o *BTExportModelParams) HasIncludeBendCenterlines() bool`

HasIncludeBendCenterlines returns a boolean if a field has been set.

### GetIncludeBendLines

`func (o *BTExportModelParams) GetIncludeBendLines() bool`

GetIncludeBendLines returns the IncludeBendLines field if non-nil, zero value otherwise.

### GetIncludeBendLinesOk

`func (o *BTExportModelParams) GetIncludeBendLinesOk() (*bool, bool)`

GetIncludeBendLinesOk returns a tuple with the IncludeBendLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBendLines

`func (o *BTExportModelParams) SetIncludeBendLines(v bool)`

SetIncludeBendLines sets IncludeBendLines field to given value.

### HasIncludeBendLines

`func (o *BTExportModelParams) HasIncludeBendLines() bool`

HasIncludeBendLines returns a boolean if a field has been set.

### GetIncludeCustomProperties

`func (o *BTExportModelParams) GetIncludeCustomProperties() bool`

GetIncludeCustomProperties returns the IncludeCustomProperties field if non-nil, zero value otherwise.

### GetIncludeCustomPropertiesOk

`func (o *BTExportModelParams) GetIncludeCustomPropertiesOk() (*bool, bool)`

GetIncludeCustomPropertiesOk returns a tuple with the IncludeCustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCustomProperties

`func (o *BTExportModelParams) SetIncludeCustomProperties(v bool)`

SetIncludeCustomProperties sets IncludeCustomProperties field to given value.

### HasIncludeCustomProperties

`func (o *BTExportModelParams) HasIncludeCustomProperties() bool`

HasIncludeCustomProperties returns a boolean if a field has been set.

### GetIncludeCustomPropertiesData

`func (o *BTExportModelParams) GetIncludeCustomPropertiesData() bool`

GetIncludeCustomPropertiesData returns the IncludeCustomPropertiesData field if non-nil, zero value otherwise.

### GetIncludeCustomPropertiesDataOk

`func (o *BTExportModelParams) GetIncludeCustomPropertiesDataOk() (*bool, bool)`

GetIncludeCustomPropertiesDataOk returns a tuple with the IncludeCustomPropertiesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCustomPropertiesData

`func (o *BTExportModelParams) SetIncludeCustomPropertiesData(v bool)`

SetIncludeCustomPropertiesData sets IncludeCustomPropertiesData field to given value.

### HasIncludeCustomPropertiesData

`func (o *BTExportModelParams) HasIncludeCustomPropertiesData() bool`

HasIncludeCustomPropertiesData returns a boolean if a field has been set.

### GetIncludeExportIds

`func (o *BTExportModelParams) GetIncludeExportIds() bool`

GetIncludeExportIds returns the IncludeExportIds field if non-nil, zero value otherwise.

### GetIncludeExportIdsOk

`func (o *BTExportModelParams) GetIncludeExportIdsOk() (*bool, bool)`

GetIncludeExportIdsOk returns a tuple with the IncludeExportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExportIds

`func (o *BTExportModelParams) SetIncludeExportIds(v bool)`

SetIncludeExportIds sets IncludeExportIds field to given value.

### HasIncludeExportIds

`func (o *BTExportModelParams) HasIncludeExportIds() bool`

HasIncludeExportIds returns a boolean if a field has been set.

### GetIncludeForeignData

`func (o *BTExportModelParams) GetIncludeForeignData() bool`

GetIncludeForeignData returns the IncludeForeignData field if non-nil, zero value otherwise.

### GetIncludeForeignDataOk

`func (o *BTExportModelParams) GetIncludeForeignDataOk() (*bool, bool)`

GetIncludeForeignDataOk returns a tuple with the IncludeForeignData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeForeignData

`func (o *BTExportModelParams) SetIncludeForeignData(v bool)`

SetIncludeForeignData sets IncludeForeignData field to given value.

### HasIncludeForeignData

`func (o *BTExportModelParams) HasIncludeForeignData() bool`

HasIncludeForeignData returns a boolean if a field has been set.

### GetIncludeItemsData

`func (o *BTExportModelParams) GetIncludeItemsData() bool`

GetIncludeItemsData returns the IncludeItemsData field if non-nil, zero value otherwise.

### GetIncludeItemsDataOk

`func (o *BTExportModelParams) GetIncludeItemsDataOk() (*bool, bool)`

GetIncludeItemsDataOk returns a tuple with the IncludeItemsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeItemsData

`func (o *BTExportModelParams) SetIncludeItemsData(v bool)`

SetIncludeItemsData sets IncludeItemsData field to given value.

### HasIncludeItemsData

`func (o *BTExportModelParams) HasIncludeItemsData() bool`

HasIncludeItemsData returns a boolean if a field has been set.

### GetIncludeLinkedDocuments

`func (o *BTExportModelParams) GetIncludeLinkedDocuments() bool`

GetIncludeLinkedDocuments returns the IncludeLinkedDocuments field if non-nil, zero value otherwise.

### GetIncludeLinkedDocumentsOk

`func (o *BTExportModelParams) GetIncludeLinkedDocumentsOk() (*bool, bool)`

GetIncludeLinkedDocumentsOk returns a tuple with the IncludeLinkedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLinkedDocuments

`func (o *BTExportModelParams) SetIncludeLinkedDocuments(v bool)`

SetIncludeLinkedDocuments sets IncludeLinkedDocuments field to given value.

### HasIncludeLinkedDocuments

`func (o *BTExportModelParams) HasIncludeLinkedDocuments() bool`

HasIncludeLinkedDocuments returns a boolean if a field has been set.

### GetIncludeReleaseManagementData

`func (o *BTExportModelParams) GetIncludeReleaseManagementData() bool`

GetIncludeReleaseManagementData returns the IncludeReleaseManagementData field if non-nil, zero value otherwise.

### GetIncludeReleaseManagementDataOk

`func (o *BTExportModelParams) GetIncludeReleaseManagementDataOk() (*bool, bool)`

GetIncludeReleaseManagementDataOk returns a tuple with the IncludeReleaseManagementData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeReleaseManagementData

`func (o *BTExportModelParams) SetIncludeReleaseManagementData(v bool)`

SetIncludeReleaseManagementData sets IncludeReleaseManagementData field to given value.

### HasIncludeReleaseManagementData

`func (o *BTExportModelParams) HasIncludeReleaseManagementData() bool`

HasIncludeReleaseManagementData returns a boolean if a field has been set.

### GetIncludeSketches

`func (o *BTExportModelParams) GetIncludeSketches() bool`

GetIncludeSketches returns the IncludeSketches field if non-nil, zero value otherwise.

### GetIncludeSketchesOk

`func (o *BTExportModelParams) GetIncludeSketchesOk() (*bool, bool)`

GetIncludeSketchesOk returns a tuple with the IncludeSketches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSketches

`func (o *BTExportModelParams) SetIncludeSketches(v bool)`

SetIncludeSketches sets IncludeSketches field to given value.

### HasIncludeSketches

`func (o *BTExportModelParams) HasIncludeSketches() bool`

HasIncludeSketches returns a boolean if a field has been set.

### GetIncludeStd

`func (o *BTExportModelParams) GetIncludeStd() bool`

GetIncludeStd returns the IncludeStd field if non-nil, zero value otherwise.

### GetIncludeStdOk

`func (o *BTExportModelParams) GetIncludeStdOk() (*bool, bool)`

GetIncludeStdOk returns a tuple with the IncludeStd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStd

`func (o *BTExportModelParams) SetIncludeStd(v bool)`

SetIncludeStd sets IncludeStd field to given value.

### HasIncludeStd

`func (o *BTExportModelParams) HasIncludeStd() bool`

HasIncludeStd returns a boolean if a field has been set.

### GetIsPartingOut

`func (o *BTExportModelParams) GetIsPartingOut() bool`

GetIsPartingOut returns the IsPartingOut field if non-nil, zero value otherwise.

### GetIsPartingOutOk

`func (o *BTExportModelParams) GetIsPartingOutOk() (*bool, bool)`

GetIsPartingOutOk returns a tuple with the IsPartingOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartingOut

`func (o *BTExportModelParams) SetIsPartingOut(v bool)`

SetIsPartingOut sets IsPartingOut field to given value.

### HasIsPartingOut

`func (o *BTExportModelParams) HasIsPartingOut() bool`

HasIsPartingOut returns a boolean if a field has been set.

### GetLinkDocumentId

`func (o *BTExportModelParams) GetLinkDocumentId() string`

GetLinkDocumentId returns the LinkDocumentId field if non-nil, zero value otherwise.

### GetLinkDocumentIdOk

`func (o *BTExportModelParams) GetLinkDocumentIdOk() (*string, bool)`

GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentId

`func (o *BTExportModelParams) SetLinkDocumentId(v string)`

SetLinkDocumentId sets LinkDocumentId field to given value.

### HasLinkDocumentId

`func (o *BTExportModelParams) HasLinkDocumentId() bool`

HasLinkDocumentId returns a boolean if a field has been set.

### GetLinkDocumentWorkspaceId

`func (o *BTExportModelParams) GetLinkDocumentWorkspaceId() string`

GetLinkDocumentWorkspaceId returns the LinkDocumentWorkspaceId field if non-nil, zero value otherwise.

### GetLinkDocumentWorkspaceIdOk

`func (o *BTExportModelParams) GetLinkDocumentWorkspaceIdOk() (*string, bool)`

GetLinkDocumentWorkspaceIdOk returns a tuple with the LinkDocumentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentWorkspaceId

`func (o *BTExportModelParams) SetLinkDocumentWorkspaceId(v string)`

SetLinkDocumentWorkspaceId sets LinkDocumentWorkspaceId field to given value.

### HasLinkDocumentWorkspaceId

`func (o *BTExportModelParams) HasLinkDocumentWorkspaceId() bool`

HasLinkDocumentWorkspaceId returns a boolean if a field has been set.

### GetMaxFacetWidth

`func (o *BTExportModelParams) GetMaxFacetWidth() float64`

GetMaxFacetWidth returns the MaxFacetWidth field if non-nil, zero value otherwise.

### GetMaxFacetWidthOk

`func (o *BTExportModelParams) GetMaxFacetWidthOk() (*float64, bool)`

GetMaxFacetWidthOk returns a tuple with the MaxFacetWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFacetWidth

`func (o *BTExportModelParams) SetMaxFacetWidth(v float64)`

SetMaxFacetWidth sets MaxFacetWidth field to given value.

### HasMaxFacetWidth

`func (o *BTExportModelParams) HasMaxFacetWidth() bool`

HasMaxFacetWidth returns a boolean if a field has been set.

### GetMicroversion

`func (o *BTExportModelParams) GetMicroversion() string`

GetMicroversion returns the Microversion field if non-nil, zero value otherwise.

### GetMicroversionOk

`func (o *BTExportModelParams) GetMicroversionOk() (*string, bool)`

GetMicroversionOk returns a tuple with the Microversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversion

`func (o *BTExportModelParams) SetMicroversion(v string)`

SetMicroversion sets Microversion field to given value.

### HasMicroversion

`func (o *BTExportModelParams) HasMicroversion() bool`

HasMicroversion returns a boolean if a field has been set.

### GetMinFacetWidth

`func (o *BTExportModelParams) GetMinFacetWidth() float64`

GetMinFacetWidth returns the MinFacetWidth field if non-nil, zero value otherwise.

### GetMinFacetWidthOk

`func (o *BTExportModelParams) GetMinFacetWidthOk() (*float64, bool)`

GetMinFacetWidthOk returns a tuple with the MinFacetWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFacetWidth

`func (o *BTExportModelParams) SetMinFacetWidth(v float64)`

SetMinFacetWidth sets MinFacetWidth field to given value.

### HasMinFacetWidth

`func (o *BTExportModelParams) HasMinFacetWidth() bool`

HasMinFacetWidth returns a boolean if a field has been set.

### GetMode

`func (o *BTExportModelParams) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BTExportModelParams) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BTExportModelParams) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *BTExportModelParams) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPartIds

`func (o *BTExportModelParams) GetPartIds() string`

GetPartIds returns the PartIds field if non-nil, zero value otherwise.

### GetPartIdsOk

`func (o *BTExportModelParams) GetPartIdsOk() (*string, bool)`

GetPartIdsOk returns a tuple with the PartIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIds

`func (o *BTExportModelParams) SetPartIds(v string)`

SetPartIds sets PartIds field to given value.

### HasPartIds

`func (o *BTExportModelParams) HasPartIds() bool`

HasPartIds returns a boolean if a field has been set.

### GetPartQuery

`func (o *BTExportModelParams) GetPartQuery() string`

GetPartQuery returns the PartQuery field if non-nil, zero value otherwise.

### GetPartQueryOk

`func (o *BTExportModelParams) GetPartQueryOk() (*string, bool)`

GetPartQueryOk returns a tuple with the PartQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartQuery

`func (o *BTExportModelParams) SetPartQuery(v string)`

SetPartQuery sets PartQuery field to given value.

### HasPartQuery

`func (o *BTExportModelParams) HasPartQuery() bool`

HasPartQuery returns a boolean if a field has been set.

### GetPassword

`func (o *BTExportModelParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BTExportModelParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BTExportModelParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BTExportModelParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordRequired

`func (o *BTExportModelParams) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *BTExportModelParams) GetPasswordRequiredOk() (*bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequired

`func (o *BTExportModelParams) SetPasswordRequired(v bool)`

SetPasswordRequired sets PasswordRequired field to given value.

### HasPasswordRequired

`func (o *BTExportModelParams) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### GetResolution

`func (o *BTExportModelParams) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *BTExportModelParams) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *BTExportModelParams) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *BTExportModelParams) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetScale

`func (o *BTExportModelParams) GetScale() float64`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *BTExportModelParams) GetScaleOk() (*float64, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *BTExportModelParams) SetScale(v float64)`

SetScale sets Scale field to given value.

### HasScale

`func (o *BTExportModelParams) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetSendCopyToMe

`func (o *BTExportModelParams) GetSendCopyToMe() bool`

GetSendCopyToMe returns the SendCopyToMe field if non-nil, zero value otherwise.

### GetSendCopyToMeOk

`func (o *BTExportModelParams) GetSendCopyToMeOk() (*bool, bool)`

GetSendCopyToMeOk returns a tuple with the SendCopyToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCopyToMe

`func (o *BTExportModelParams) SetSendCopyToMe(v bool)`

SetSendCopyToMe sets SendCopyToMe field to given value.

### HasSendCopyToMe

`func (o *BTExportModelParams) HasSendCopyToMe() bool`

HasSendCopyToMe returns a boolean if a field has been set.

### GetSheetMetalFlat

`func (o *BTExportModelParams) GetSheetMetalFlat() bool`

GetSheetMetalFlat returns the SheetMetalFlat field if non-nil, zero value otherwise.

### GetSheetMetalFlatOk

`func (o *BTExportModelParams) GetSheetMetalFlatOk() (*bool, bool)`

GetSheetMetalFlatOk returns a tuple with the SheetMetalFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetMetalFlat

`func (o *BTExportModelParams) SetSheetMetalFlat(v bool)`

SetSheetMetalFlat sets SheetMetalFlat field to given value.

### HasSheetMetalFlat

`func (o *BTExportModelParams) HasSheetMetalFlat() bool`

HasSheetMetalFlat returns a boolean if a field has been set.

### GetSplinesAsPolylines

`func (o *BTExportModelParams) GetSplinesAsPolylines() bool`

GetSplinesAsPolylines returns the SplinesAsPolylines field if non-nil, zero value otherwise.

### GetSplinesAsPolylinesOk

`func (o *BTExportModelParams) GetSplinesAsPolylinesOk() (*bool, bool)`

GetSplinesAsPolylinesOk returns a tuple with the SplinesAsPolylines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplinesAsPolylines

`func (o *BTExportModelParams) SetSplinesAsPolylines(v bool)`

SetSplinesAsPolylines sets SplinesAsPolylines field to given value.

### HasSplinesAsPolylines

`func (o *BTExportModelParams) HasSplinesAsPolylines() bool`

HasSplinesAsPolylines returns a boolean if a field has been set.

### GetStoreInDocument

`func (o *BTExportModelParams) GetStoreInDocument() bool`

GetStoreInDocument returns the StoreInDocument field if non-nil, zero value otherwise.

### GetStoreInDocumentOk

`func (o *BTExportModelParams) GetStoreInDocumentOk() (*bool, bool)`

GetStoreInDocumentOk returns a tuple with the StoreInDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInDocument

`func (o *BTExportModelParams) SetStoreInDocument(v bool)`

SetStoreInDocument sets StoreInDocument field to given value.

### HasStoreInDocument

`func (o *BTExportModelParams) HasStoreInDocument() bool`

HasStoreInDocument returns a boolean if a field has been set.

### GetTriggerAutoDownload

`func (o *BTExportModelParams) GetTriggerAutoDownload() bool`

GetTriggerAutoDownload returns the TriggerAutoDownload field if non-nil, zero value otherwise.

### GetTriggerAutoDownloadOk

`func (o *BTExportModelParams) GetTriggerAutoDownloadOk() (*bool, bool)`

GetTriggerAutoDownloadOk returns a tuple with the TriggerAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAutoDownload

`func (o *BTExportModelParams) SetTriggerAutoDownload(v bool)`

SetTriggerAutoDownload sets TriggerAutoDownload field to given value.

### HasTriggerAutoDownload

`func (o *BTExportModelParams) HasTriggerAutoDownload() bool`

HasTriggerAutoDownload returns a boolean if a field has been set.

### GetUnits

`func (o *BTExportModelParams) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BTExportModelParams) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BTExportModelParams) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BTExportModelParams) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetUserId

`func (o *BTExportModelParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTExportModelParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTExportModelParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTExportModelParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetValidForDays

`func (o *BTExportModelParams) GetValidForDays() int32`

GetValidForDays returns the ValidForDays field if non-nil, zero value otherwise.

### GetValidForDaysOk

`func (o *BTExportModelParams) GetValidForDaysOk() (*int32, bool)`

GetValidForDaysOk returns a tuple with the ValidForDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidForDays

`func (o *BTExportModelParams) SetValidForDays(v int32)`

SetValidForDays sets ValidForDays field to given value.

### HasValidForDays

`func (o *BTExportModelParams) HasValidForDays() bool`

HasValidForDays returns a boolean if a field has been set.

### GetVersion

`func (o *BTExportModelParams) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTExportModelParams) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTExportModelParams) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTExportModelParams) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetView

`func (o *BTExportModelParams) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *BTExportModelParams) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *BTExportModelParams) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *BTExportModelParams) HasView() bool`

HasView returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTExportModelParams) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTExportModelParams) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTExportModelParams) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTExportModelParams) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetZipSingleFileOutput

`func (o *BTExportModelParams) GetZipSingleFileOutput() bool`

GetZipSingleFileOutput returns the ZipSingleFileOutput field if non-nil, zero value otherwise.

### GetZipSingleFileOutputOk

`func (o *BTExportModelParams) GetZipSingleFileOutputOk() (*bool, bool)`

GetZipSingleFileOutputOk returns a tuple with the ZipSingleFileOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipSingleFileOutput

`func (o *BTExportModelParams) SetZipSingleFileOutput(v bool)`

SetZipSingleFileOutput sets ZipSingleFileOutput field to given value.

### HasZipSingleFileOutput

`func (o *BTExportModelParams) HasZipSingleFileOutput() bool`

HasZipSingleFileOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


