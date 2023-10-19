# BTBExportModelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AngleTolerance** | Pointer to **float64** |  | [optional] 
**BatchAllFlatPatterns** | Pointer to **bool** |  | [optional] 
**BatchFlatPatterns** | Pointer to **bool** |  | [optional] 
**ChordTolerance** | Pointer to **float64** |  | [optional] 
**CloudObjectId** | Pointer to **string** |  | [optional] 
**CloudStorageAccountId** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**DestinationName** | Pointer to **string** |  | [optional] 
**DocumentId** | **string** |  | 
**DocumentVersionId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementIds** | Pointer to **string** |  | [optional] 
**EmailLink** | Pointer to **bool** |  | [optional] 
**EmailMessage** | Pointer to **string** |  | [optional] 
**EmailSubject** | Pointer to **string** |  | [optional] 
**EmailTo** | Pointer to **string** |  | [optional] 
**FeatureIds** | Pointer to **string** |  | [optional] 
**Flatten** | Pointer to **bool** |  | [optional] 
**Format** | **string** |  | 
**FromUserId** | Pointer to **string** |  | [optional] 
**Grouping** | Pointer to **string** |  | [optional] 
**IgnoreExportRulesForContents** | Pointer to **bool** |  | [optional] 
**IncludeBendCenterlines** | Pointer to **bool** |  | [optional] 
**IncludeBendLines** | Pointer to **bool** |  | [optional] 
**IncludeExportIds** | Pointer to **bool** |  | [optional] 
**IncludeSketches** | Pointer to **bool** |  | [optional] 
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

### NewBTBExportModelParams

`func NewBTBExportModelParams(documentId string, format string, ) *BTBExportModelParams`

NewBTBExportModelParams instantiates a new BTBExportModelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBExportModelParamsWithDefaults

`func NewBTBExportModelParamsWithDefaults() *BTBExportModelParams`

NewBTBExportModelParamsWithDefaults instantiates a new BTBExportModelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAngleTolerance

`func (o *BTBExportModelParams) GetAngleTolerance() float64`

GetAngleTolerance returns the AngleTolerance field if non-nil, zero value otherwise.

### GetAngleToleranceOk

`func (o *BTBExportModelParams) GetAngleToleranceOk() (*float64, bool)`

GetAngleToleranceOk returns a tuple with the AngleTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngleTolerance

`func (o *BTBExportModelParams) SetAngleTolerance(v float64)`

SetAngleTolerance sets AngleTolerance field to given value.

### HasAngleTolerance

`func (o *BTBExportModelParams) HasAngleTolerance() bool`

HasAngleTolerance returns a boolean if a field has been set.

### GetBatchAllFlatPatterns

`func (o *BTBExportModelParams) GetBatchAllFlatPatterns() bool`

GetBatchAllFlatPatterns returns the BatchAllFlatPatterns field if non-nil, zero value otherwise.

### GetBatchAllFlatPatternsOk

`func (o *BTBExportModelParams) GetBatchAllFlatPatternsOk() (*bool, bool)`

GetBatchAllFlatPatternsOk returns a tuple with the BatchAllFlatPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchAllFlatPatterns

`func (o *BTBExportModelParams) SetBatchAllFlatPatterns(v bool)`

SetBatchAllFlatPatterns sets BatchAllFlatPatterns field to given value.

### HasBatchAllFlatPatterns

`func (o *BTBExportModelParams) HasBatchAllFlatPatterns() bool`

HasBatchAllFlatPatterns returns a boolean if a field has been set.

### GetBatchFlatPatterns

`func (o *BTBExportModelParams) GetBatchFlatPatterns() bool`

GetBatchFlatPatterns returns the BatchFlatPatterns field if non-nil, zero value otherwise.

### GetBatchFlatPatternsOk

`func (o *BTBExportModelParams) GetBatchFlatPatternsOk() (*bool, bool)`

GetBatchFlatPatternsOk returns a tuple with the BatchFlatPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFlatPatterns

`func (o *BTBExportModelParams) SetBatchFlatPatterns(v bool)`

SetBatchFlatPatterns sets BatchFlatPatterns field to given value.

### HasBatchFlatPatterns

`func (o *BTBExportModelParams) HasBatchFlatPatterns() bool`

HasBatchFlatPatterns returns a boolean if a field has been set.

### GetChordTolerance

`func (o *BTBExportModelParams) GetChordTolerance() float64`

GetChordTolerance returns the ChordTolerance field if non-nil, zero value otherwise.

### GetChordToleranceOk

`func (o *BTBExportModelParams) GetChordToleranceOk() (*float64, bool)`

GetChordToleranceOk returns a tuple with the ChordTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChordTolerance

`func (o *BTBExportModelParams) SetChordTolerance(v float64)`

SetChordTolerance sets ChordTolerance field to given value.

### HasChordTolerance

`func (o *BTBExportModelParams) HasChordTolerance() bool`

HasChordTolerance returns a boolean if a field has been set.

### GetCloudObjectId

`func (o *BTBExportModelParams) GetCloudObjectId() string`

GetCloudObjectId returns the CloudObjectId field if non-nil, zero value otherwise.

### GetCloudObjectIdOk

`func (o *BTBExportModelParams) GetCloudObjectIdOk() (*string, bool)`

GetCloudObjectIdOk returns a tuple with the CloudObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudObjectId

`func (o *BTBExportModelParams) SetCloudObjectId(v string)`

SetCloudObjectId sets CloudObjectId field to given value.

### HasCloudObjectId

`func (o *BTBExportModelParams) HasCloudObjectId() bool`

HasCloudObjectId returns a boolean if a field has been set.

### GetCloudStorageAccountId

`func (o *BTBExportModelParams) GetCloudStorageAccountId() string`

GetCloudStorageAccountId returns the CloudStorageAccountId field if non-nil, zero value otherwise.

### GetCloudStorageAccountIdOk

`func (o *BTBExportModelParams) GetCloudStorageAccountIdOk() (*string, bool)`

GetCloudStorageAccountIdOk returns a tuple with the CloudStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageAccountId

`func (o *BTBExportModelParams) SetCloudStorageAccountId(v string)`

SetCloudStorageAccountId sets CloudStorageAccountId field to given value.

### HasCloudStorageAccountId

`func (o *BTBExportModelParams) HasCloudStorageAccountId() bool`

HasCloudStorageAccountId returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTBExportModelParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTBExportModelParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTBExportModelParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTBExportModelParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDestinationName

`func (o *BTBExportModelParams) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *BTBExportModelParams) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *BTBExportModelParams) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *BTBExportModelParams) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTBExportModelParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTBExportModelParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTBExportModelParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetDocumentVersionId

`func (o *BTBExportModelParams) GetDocumentVersionId() string`

GetDocumentVersionId returns the DocumentVersionId field if non-nil, zero value otherwise.

### GetDocumentVersionIdOk

`func (o *BTBExportModelParams) GetDocumentVersionIdOk() (*string, bool)`

GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentVersionId

`func (o *BTBExportModelParams) SetDocumentVersionId(v string)`

SetDocumentVersionId sets DocumentVersionId field to given value.

### HasDocumentVersionId

`func (o *BTBExportModelParams) HasDocumentVersionId() bool`

HasDocumentVersionId returns a boolean if a field has been set.

### GetElementId

`func (o *BTBExportModelParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTBExportModelParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTBExportModelParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTBExportModelParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementIds

`func (o *BTBExportModelParams) GetElementIds() string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *BTBExportModelParams) GetElementIdsOk() (*string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *BTBExportModelParams) SetElementIds(v string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *BTBExportModelParams) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### GetEmailLink

`func (o *BTBExportModelParams) GetEmailLink() bool`

GetEmailLink returns the EmailLink field if non-nil, zero value otherwise.

### GetEmailLinkOk

`func (o *BTBExportModelParams) GetEmailLinkOk() (*bool, bool)`

GetEmailLinkOk returns a tuple with the EmailLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLink

`func (o *BTBExportModelParams) SetEmailLink(v bool)`

SetEmailLink sets EmailLink field to given value.

### HasEmailLink

`func (o *BTBExportModelParams) HasEmailLink() bool`

HasEmailLink returns a boolean if a field has been set.

### GetEmailMessage

`func (o *BTBExportModelParams) GetEmailMessage() string`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *BTBExportModelParams) GetEmailMessageOk() (*string, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *BTBExportModelParams) SetEmailMessage(v string)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *BTBExportModelParams) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.

### GetEmailSubject

`func (o *BTBExportModelParams) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *BTBExportModelParams) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *BTBExportModelParams) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *BTBExportModelParams) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailTo

`func (o *BTBExportModelParams) GetEmailTo() string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *BTBExportModelParams) GetEmailToOk() (*string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *BTBExportModelParams) SetEmailTo(v string)`

SetEmailTo sets EmailTo field to given value.

### HasEmailTo

`func (o *BTBExportModelParams) HasEmailTo() bool`

HasEmailTo returns a boolean if a field has been set.

### GetFeatureIds

`func (o *BTBExportModelParams) GetFeatureIds() string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *BTBExportModelParams) GetFeatureIdsOk() (*string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *BTBExportModelParams) SetFeatureIds(v string)`

SetFeatureIds sets FeatureIds field to given value.

### HasFeatureIds

`func (o *BTBExportModelParams) HasFeatureIds() bool`

HasFeatureIds returns a boolean if a field has been set.

### GetFlatten

`func (o *BTBExportModelParams) GetFlatten() bool`

GetFlatten returns the Flatten field if non-nil, zero value otherwise.

### GetFlattenOk

`func (o *BTBExportModelParams) GetFlattenOk() (*bool, bool)`

GetFlattenOk returns a tuple with the Flatten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatten

`func (o *BTBExportModelParams) SetFlatten(v bool)`

SetFlatten sets Flatten field to given value.

### HasFlatten

`func (o *BTBExportModelParams) HasFlatten() bool`

HasFlatten returns a boolean if a field has been set.

### GetFormat

`func (o *BTBExportModelParams) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BTBExportModelParams) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BTBExportModelParams) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetFromUserId

`func (o *BTBExportModelParams) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *BTBExportModelParams) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *BTBExportModelParams) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *BTBExportModelParams) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetGrouping

`func (o *BTBExportModelParams) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BTBExportModelParams) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BTBExportModelParams) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *BTBExportModelParams) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetIgnoreExportRulesForContents

`func (o *BTBExportModelParams) GetIgnoreExportRulesForContents() bool`

GetIgnoreExportRulesForContents returns the IgnoreExportRulesForContents field if non-nil, zero value otherwise.

### GetIgnoreExportRulesForContentsOk

`func (o *BTBExportModelParams) GetIgnoreExportRulesForContentsOk() (*bool, bool)`

GetIgnoreExportRulesForContentsOk returns a tuple with the IgnoreExportRulesForContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreExportRulesForContents

`func (o *BTBExportModelParams) SetIgnoreExportRulesForContents(v bool)`

SetIgnoreExportRulesForContents sets IgnoreExportRulesForContents field to given value.

### HasIgnoreExportRulesForContents

`func (o *BTBExportModelParams) HasIgnoreExportRulesForContents() bool`

HasIgnoreExportRulesForContents returns a boolean if a field has been set.

### GetIncludeBendCenterlines

`func (o *BTBExportModelParams) GetIncludeBendCenterlines() bool`

GetIncludeBendCenterlines returns the IncludeBendCenterlines field if non-nil, zero value otherwise.

### GetIncludeBendCenterlinesOk

`func (o *BTBExportModelParams) GetIncludeBendCenterlinesOk() (*bool, bool)`

GetIncludeBendCenterlinesOk returns a tuple with the IncludeBendCenterlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBendCenterlines

`func (o *BTBExportModelParams) SetIncludeBendCenterlines(v bool)`

SetIncludeBendCenterlines sets IncludeBendCenterlines field to given value.

### HasIncludeBendCenterlines

`func (o *BTBExportModelParams) HasIncludeBendCenterlines() bool`

HasIncludeBendCenterlines returns a boolean if a field has been set.

### GetIncludeBendLines

`func (o *BTBExportModelParams) GetIncludeBendLines() bool`

GetIncludeBendLines returns the IncludeBendLines field if non-nil, zero value otherwise.

### GetIncludeBendLinesOk

`func (o *BTBExportModelParams) GetIncludeBendLinesOk() (*bool, bool)`

GetIncludeBendLinesOk returns a tuple with the IncludeBendLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBendLines

`func (o *BTBExportModelParams) SetIncludeBendLines(v bool)`

SetIncludeBendLines sets IncludeBendLines field to given value.

### HasIncludeBendLines

`func (o *BTBExportModelParams) HasIncludeBendLines() bool`

HasIncludeBendLines returns a boolean if a field has been set.

### GetIncludeExportIds

`func (o *BTBExportModelParams) GetIncludeExportIds() bool`

GetIncludeExportIds returns the IncludeExportIds field if non-nil, zero value otherwise.

### GetIncludeExportIdsOk

`func (o *BTBExportModelParams) GetIncludeExportIdsOk() (*bool, bool)`

GetIncludeExportIdsOk returns a tuple with the IncludeExportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExportIds

`func (o *BTBExportModelParams) SetIncludeExportIds(v bool)`

SetIncludeExportIds sets IncludeExportIds field to given value.

### HasIncludeExportIds

`func (o *BTBExportModelParams) HasIncludeExportIds() bool`

HasIncludeExportIds returns a boolean if a field has been set.

### GetIncludeSketches

`func (o *BTBExportModelParams) GetIncludeSketches() bool`

GetIncludeSketches returns the IncludeSketches field if non-nil, zero value otherwise.

### GetIncludeSketchesOk

`func (o *BTBExportModelParams) GetIncludeSketchesOk() (*bool, bool)`

GetIncludeSketchesOk returns a tuple with the IncludeSketches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSketches

`func (o *BTBExportModelParams) SetIncludeSketches(v bool)`

SetIncludeSketches sets IncludeSketches field to given value.

### HasIncludeSketches

`func (o *BTBExportModelParams) HasIncludeSketches() bool`

HasIncludeSketches returns a boolean if a field has been set.

### GetIsPartingOut

`func (o *BTBExportModelParams) GetIsPartingOut() bool`

GetIsPartingOut returns the IsPartingOut field if non-nil, zero value otherwise.

### GetIsPartingOutOk

`func (o *BTBExportModelParams) GetIsPartingOutOk() (*bool, bool)`

GetIsPartingOutOk returns a tuple with the IsPartingOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartingOut

`func (o *BTBExportModelParams) SetIsPartingOut(v bool)`

SetIsPartingOut sets IsPartingOut field to given value.

### HasIsPartingOut

`func (o *BTBExportModelParams) HasIsPartingOut() bool`

HasIsPartingOut returns a boolean if a field has been set.

### GetLinkDocumentId

`func (o *BTBExportModelParams) GetLinkDocumentId() string`

GetLinkDocumentId returns the LinkDocumentId field if non-nil, zero value otherwise.

### GetLinkDocumentIdOk

`func (o *BTBExportModelParams) GetLinkDocumentIdOk() (*string, bool)`

GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentId

`func (o *BTBExportModelParams) SetLinkDocumentId(v string)`

SetLinkDocumentId sets LinkDocumentId field to given value.

### HasLinkDocumentId

`func (o *BTBExportModelParams) HasLinkDocumentId() bool`

HasLinkDocumentId returns a boolean if a field has been set.

### GetLinkDocumentWorkspaceId

`func (o *BTBExportModelParams) GetLinkDocumentWorkspaceId() string`

GetLinkDocumentWorkspaceId returns the LinkDocumentWorkspaceId field if non-nil, zero value otherwise.

### GetLinkDocumentWorkspaceIdOk

`func (o *BTBExportModelParams) GetLinkDocumentWorkspaceIdOk() (*string, bool)`

GetLinkDocumentWorkspaceIdOk returns a tuple with the LinkDocumentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentWorkspaceId

`func (o *BTBExportModelParams) SetLinkDocumentWorkspaceId(v string)`

SetLinkDocumentWorkspaceId sets LinkDocumentWorkspaceId field to given value.

### HasLinkDocumentWorkspaceId

`func (o *BTBExportModelParams) HasLinkDocumentWorkspaceId() bool`

HasLinkDocumentWorkspaceId returns a boolean if a field has been set.

### GetMaxFacetWidth

`func (o *BTBExportModelParams) GetMaxFacetWidth() float64`

GetMaxFacetWidth returns the MaxFacetWidth field if non-nil, zero value otherwise.

### GetMaxFacetWidthOk

`func (o *BTBExportModelParams) GetMaxFacetWidthOk() (*float64, bool)`

GetMaxFacetWidthOk returns a tuple with the MaxFacetWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFacetWidth

`func (o *BTBExportModelParams) SetMaxFacetWidth(v float64)`

SetMaxFacetWidth sets MaxFacetWidth field to given value.

### HasMaxFacetWidth

`func (o *BTBExportModelParams) HasMaxFacetWidth() bool`

HasMaxFacetWidth returns a boolean if a field has been set.

### GetMicroversion

`func (o *BTBExportModelParams) GetMicroversion() string`

GetMicroversion returns the Microversion field if non-nil, zero value otherwise.

### GetMicroversionOk

`func (o *BTBExportModelParams) GetMicroversionOk() (*string, bool)`

GetMicroversionOk returns a tuple with the Microversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversion

`func (o *BTBExportModelParams) SetMicroversion(v string)`

SetMicroversion sets Microversion field to given value.

### HasMicroversion

`func (o *BTBExportModelParams) HasMicroversion() bool`

HasMicroversion returns a boolean if a field has been set.

### GetMinFacetWidth

`func (o *BTBExportModelParams) GetMinFacetWidth() float64`

GetMinFacetWidth returns the MinFacetWidth field if non-nil, zero value otherwise.

### GetMinFacetWidthOk

`func (o *BTBExportModelParams) GetMinFacetWidthOk() (*float64, bool)`

GetMinFacetWidthOk returns a tuple with the MinFacetWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinFacetWidth

`func (o *BTBExportModelParams) SetMinFacetWidth(v float64)`

SetMinFacetWidth sets MinFacetWidth field to given value.

### HasMinFacetWidth

`func (o *BTBExportModelParams) HasMinFacetWidth() bool`

HasMinFacetWidth returns a boolean if a field has been set.

### GetMode

`func (o *BTBExportModelParams) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BTBExportModelParams) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BTBExportModelParams) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *BTBExportModelParams) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPartIds

`func (o *BTBExportModelParams) GetPartIds() string`

GetPartIds returns the PartIds field if non-nil, zero value otherwise.

### GetPartIdsOk

`func (o *BTBExportModelParams) GetPartIdsOk() (*string, bool)`

GetPartIdsOk returns a tuple with the PartIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIds

`func (o *BTBExportModelParams) SetPartIds(v string)`

SetPartIds sets PartIds field to given value.

### HasPartIds

`func (o *BTBExportModelParams) HasPartIds() bool`

HasPartIds returns a boolean if a field has been set.

### GetPartQuery

`func (o *BTBExportModelParams) GetPartQuery() string`

GetPartQuery returns the PartQuery field if non-nil, zero value otherwise.

### GetPartQueryOk

`func (o *BTBExportModelParams) GetPartQueryOk() (*string, bool)`

GetPartQueryOk returns a tuple with the PartQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartQuery

`func (o *BTBExportModelParams) SetPartQuery(v string)`

SetPartQuery sets PartQuery field to given value.

### HasPartQuery

`func (o *BTBExportModelParams) HasPartQuery() bool`

HasPartQuery returns a boolean if a field has been set.

### GetPassword

`func (o *BTBExportModelParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BTBExportModelParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BTBExportModelParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BTBExportModelParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordRequired

`func (o *BTBExportModelParams) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *BTBExportModelParams) GetPasswordRequiredOk() (*bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequired

`func (o *BTBExportModelParams) SetPasswordRequired(v bool)`

SetPasswordRequired sets PasswordRequired field to given value.

### HasPasswordRequired

`func (o *BTBExportModelParams) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### GetResolution

`func (o *BTBExportModelParams) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *BTBExportModelParams) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *BTBExportModelParams) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *BTBExportModelParams) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetScale

`func (o *BTBExportModelParams) GetScale() float64`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *BTBExportModelParams) GetScaleOk() (*float64, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *BTBExportModelParams) SetScale(v float64)`

SetScale sets Scale field to given value.

### HasScale

`func (o *BTBExportModelParams) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetSendCopyToMe

`func (o *BTBExportModelParams) GetSendCopyToMe() bool`

GetSendCopyToMe returns the SendCopyToMe field if non-nil, zero value otherwise.

### GetSendCopyToMeOk

`func (o *BTBExportModelParams) GetSendCopyToMeOk() (*bool, bool)`

GetSendCopyToMeOk returns a tuple with the SendCopyToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCopyToMe

`func (o *BTBExportModelParams) SetSendCopyToMe(v bool)`

SetSendCopyToMe sets SendCopyToMe field to given value.

### HasSendCopyToMe

`func (o *BTBExportModelParams) HasSendCopyToMe() bool`

HasSendCopyToMe returns a boolean if a field has been set.

### GetSheetMetalFlat

`func (o *BTBExportModelParams) GetSheetMetalFlat() bool`

GetSheetMetalFlat returns the SheetMetalFlat field if non-nil, zero value otherwise.

### GetSheetMetalFlatOk

`func (o *BTBExportModelParams) GetSheetMetalFlatOk() (*bool, bool)`

GetSheetMetalFlatOk returns a tuple with the SheetMetalFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetMetalFlat

`func (o *BTBExportModelParams) SetSheetMetalFlat(v bool)`

SetSheetMetalFlat sets SheetMetalFlat field to given value.

### HasSheetMetalFlat

`func (o *BTBExportModelParams) HasSheetMetalFlat() bool`

HasSheetMetalFlat returns a boolean if a field has been set.

### GetSplinesAsPolylines

`func (o *BTBExportModelParams) GetSplinesAsPolylines() bool`

GetSplinesAsPolylines returns the SplinesAsPolylines field if non-nil, zero value otherwise.

### GetSplinesAsPolylinesOk

`func (o *BTBExportModelParams) GetSplinesAsPolylinesOk() (*bool, bool)`

GetSplinesAsPolylinesOk returns a tuple with the SplinesAsPolylines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplinesAsPolylines

`func (o *BTBExportModelParams) SetSplinesAsPolylines(v bool)`

SetSplinesAsPolylines sets SplinesAsPolylines field to given value.

### HasSplinesAsPolylines

`func (o *BTBExportModelParams) HasSplinesAsPolylines() bool`

HasSplinesAsPolylines returns a boolean if a field has been set.

### GetStoreInDocument

`func (o *BTBExportModelParams) GetStoreInDocument() bool`

GetStoreInDocument returns the StoreInDocument field if non-nil, zero value otherwise.

### GetStoreInDocumentOk

`func (o *BTBExportModelParams) GetStoreInDocumentOk() (*bool, bool)`

GetStoreInDocumentOk returns a tuple with the StoreInDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInDocument

`func (o *BTBExportModelParams) SetStoreInDocument(v bool)`

SetStoreInDocument sets StoreInDocument field to given value.

### HasStoreInDocument

`func (o *BTBExportModelParams) HasStoreInDocument() bool`

HasStoreInDocument returns a boolean if a field has been set.

### GetTriggerAutoDownload

`func (o *BTBExportModelParams) GetTriggerAutoDownload() bool`

GetTriggerAutoDownload returns the TriggerAutoDownload field if non-nil, zero value otherwise.

### GetTriggerAutoDownloadOk

`func (o *BTBExportModelParams) GetTriggerAutoDownloadOk() (*bool, bool)`

GetTriggerAutoDownloadOk returns a tuple with the TriggerAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAutoDownload

`func (o *BTBExportModelParams) SetTriggerAutoDownload(v bool)`

SetTriggerAutoDownload sets TriggerAutoDownload field to given value.

### HasTriggerAutoDownload

`func (o *BTBExportModelParams) HasTriggerAutoDownload() bool`

HasTriggerAutoDownload returns a boolean if a field has been set.

### GetUnits

`func (o *BTBExportModelParams) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BTBExportModelParams) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BTBExportModelParams) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BTBExportModelParams) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetUserId

`func (o *BTBExportModelParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTBExportModelParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTBExportModelParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTBExportModelParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetValidForDays

`func (o *BTBExportModelParams) GetValidForDays() int32`

GetValidForDays returns the ValidForDays field if non-nil, zero value otherwise.

### GetValidForDaysOk

`func (o *BTBExportModelParams) GetValidForDaysOk() (*int32, bool)`

GetValidForDaysOk returns a tuple with the ValidForDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidForDays

`func (o *BTBExportModelParams) SetValidForDays(v int32)`

SetValidForDays sets ValidForDays field to given value.

### HasValidForDays

`func (o *BTBExportModelParams) HasValidForDays() bool`

HasValidForDays returns a boolean if a field has been set.

### GetVersion

`func (o *BTBExportModelParams) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTBExportModelParams) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTBExportModelParams) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTBExportModelParams) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetView

`func (o *BTBExportModelParams) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *BTBExportModelParams) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *BTBExportModelParams) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *BTBExportModelParams) HasView() bool`

HasView returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTBExportModelParams) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTBExportModelParams) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTBExportModelParams) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTBExportModelParams) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetZipSingleFileOutput

`func (o *BTBExportModelParams) GetZipSingleFileOutput() bool`

GetZipSingleFileOutput returns the ZipSingleFileOutput field if non-nil, zero value otherwise.

### GetZipSingleFileOutputOk

`func (o *BTBExportModelParams) GetZipSingleFileOutputOk() (*bool, bool)`

GetZipSingleFileOutputOk returns a tuple with the ZipSingleFileOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipSingleFileOutput

`func (o *BTBExportModelParams) SetZipSingleFileOutput(v bool)`

SetZipSingleFileOutput sets ZipSingleFileOutput field to given value.

### HasZipSingleFileOutput

`func (o *BTBExportModelParams) HasZipSingleFileOutput() bool`

HasZipSingleFileOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


