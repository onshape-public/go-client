# BTTranslateFormatParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowFaultyParts** | Pointer to **bool** | If true, parts with faults are imported. If false, faulty parts are omitted. | [optional] 
**AngularTolerance** | Pointer to **float64** | Determines the maximum angular deviation, between the analytical surface and its triangulation. Lower values result in a finer geometry and higher values result in coarser geometry. | [optional] 
**BlobElementId** | Pointer to **string** |  | [optional] 
**BlobMicroversionId** | Pointer to **string** |  | [optional] 
**CloudObjectId** | Pointer to **string** |  | [optional] 
**CloudStorageAccountId** | Pointer to **string** |  | [optional] 
**ColorMethod** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**CreateComposite** | Pointer to **bool** |  | [optional] 
**CurrentSheetOnly** | Pointer to **bool** |  | [optional] 
**DestinationName** | Pointer to **string** |  | [optional] 
**DistanceTolerance** | Pointer to **float64** |  | [optional] 
**ElementId** | Pointer to **string** | The id of the element in which to perform the operation. | [optional] 
**ElementIds** | Pointer to **[]string** | An array of element ids. | [optional] 
**EmailLink** | Pointer to **bool** |  | [optional] 
**EmailMessage** | Pointer to **string** |  | [optional] 
**EmailSubject** | Pointer to **string** |  | [optional] 
**EmailTo** | Pointer to **[]string** |  | [optional] 
**ExtractAssemblyHierarchy** | Pointer to **bool** |  | [optional] 
**Flatten** | Pointer to **bool** |  | [optional] 
**FlattenAssemblies** | Pointer to **bool** |  | [optional] 
**ForeignId** | Pointer to **string** |  | [optional] 
**FormatName** | **string** | The name of the file format. | 
**FromUserId** | Pointer to **string** |  | [optional] 
**GetyAxisIsUp** | Pointer to **bool** |  | [optional] 
**Grouping** | Pointer to **bool** |  | [optional] 
**HideInspectionItems** | Pointer to **bool** |  | [optional] 
**IgnoreExportRulesForContents** | Pointer to **bool** |  | [optional] 
**ImageHeight** | Pointer to **int32** |  | [optional] 
**ImageWidth** | Pointer to **int32** |  | [optional] 
**ImportAppearances** | Pointer to **bool** |  | [optional] 
**ImportInBackground** | Pointer to **bool** |  | [optional] 
**ImportWithinDocument** | Pointer to **bool** |  | [optional] 
**IncludeExportIds** | Pointer to **bool** |  | [optional] 
**JoinAdjacentSurfaces** | Pointer to **bool** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**LinkDocumentId** | Pointer to **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [optional] 
**LinkDocumentWorkspaceId** | Pointer to **string** |  | [optional] 
**MaximumChordLength** | Pointer to **float64** | Determines the maximum distance, between the analytical surface and its triangulation. Lower values result in a finer geometry and higher values result in coarser geometry. | [optional] 
**NotifyUser** | Pointer to **bool** |  | [optional] 
**OnePartPerDoc** | Pointer to **bool** |  | [optional] 
**OriginalForeignId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**PartIds** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordRequired** | Pointer to **bool** |  | [optional] 
**PdfVersion** | Pointer to **string** |  | [optional] 
**ProcessedForeignId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ProxyDocumentId** | Pointer to **string** |  | [optional] 
**ProxyElementId** | Pointer to **string** |  | [optional] 
**ProxyWorkspaceVersion** | Pointer to **string** |  | [optional] 
**ProxyWorkspaceVersionId** | Pointer to **string** |  | [optional] 
**SelectablePdfText** | Pointer to **bool** |  | [optional] 
**SendCopyToMe** | Pointer to **bool** |  | [optional] 
**SheetIndices** | Pointer to **[]int32** |  | [optional] 
**ShowOverriddenDimensions** | Pointer to **bool** |  | [optional] 
**SkipBodyshop** | Pointer to **bool** |  | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 
**SpecifyUnits** | Pointer to **bool** |  | [optional] 
**SplinesAsPolylines** | Pointer to **bool** |  | [optional] 
**SplitAssembliesIntoMultipleDocuments** | Pointer to **bool** |  | [optional] 
**StepParasolidPreprocessingOption** | Pointer to [**GBTPreProcessParasolidOption**](GBTPreProcessParasolidOption.md) |  | [optional] 
**StepVersionString** | Pointer to **string** |  | [optional] 
**StoreInDocument** | Pointer to **bool** |  | [optional] 
**TextAsGeometry** | Pointer to **bool** |  | [optional] 
**TextOption** | Pointer to **string** |  | [optional] 
**TriggerAutoDownload** | Pointer to **bool** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**UploadId** | Pointer to **string** |  | [optional] 
**UseGltfCompression** | Pointer to **bool** |  | [optional] 
**UseIgesCompatibilityMode** | Pointer to **bool** |  | [optional] 
**ValidForDays** | Pointer to **int32** |  | [optional] 
**VersionString** | Pointer to **string** |  | [optional] 

## Methods

### NewBTTranslateFormatParams

`func NewBTTranslateFormatParams(formatName string, ) *BTTranslateFormatParams`

NewBTTranslateFormatParams instantiates a new BTTranslateFormatParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTranslateFormatParamsWithDefaults

`func NewBTTranslateFormatParamsWithDefaults() *BTTranslateFormatParams`

NewBTTranslateFormatParamsWithDefaults instantiates a new BTTranslateFormatParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowFaultyParts

`func (o *BTTranslateFormatParams) GetAllowFaultyParts() bool`

GetAllowFaultyParts returns the AllowFaultyParts field if non-nil, zero value otherwise.

### GetAllowFaultyPartsOk

`func (o *BTTranslateFormatParams) GetAllowFaultyPartsOk() (*bool, bool)`

GetAllowFaultyPartsOk returns a tuple with the AllowFaultyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFaultyParts

`func (o *BTTranslateFormatParams) SetAllowFaultyParts(v bool)`

SetAllowFaultyParts sets AllowFaultyParts field to given value.

### HasAllowFaultyParts

`func (o *BTTranslateFormatParams) HasAllowFaultyParts() bool`

HasAllowFaultyParts returns a boolean if a field has been set.

### GetAngularTolerance

`func (o *BTTranslateFormatParams) GetAngularTolerance() float64`

GetAngularTolerance returns the AngularTolerance field if non-nil, zero value otherwise.

### GetAngularToleranceOk

`func (o *BTTranslateFormatParams) GetAngularToleranceOk() (*float64, bool)`

GetAngularToleranceOk returns a tuple with the AngularTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngularTolerance

`func (o *BTTranslateFormatParams) SetAngularTolerance(v float64)`

SetAngularTolerance sets AngularTolerance field to given value.

### HasAngularTolerance

`func (o *BTTranslateFormatParams) HasAngularTolerance() bool`

HasAngularTolerance returns a boolean if a field has been set.

### GetBlobElementId

`func (o *BTTranslateFormatParams) GetBlobElementId() string`

GetBlobElementId returns the BlobElementId field if non-nil, zero value otherwise.

### GetBlobElementIdOk

`func (o *BTTranslateFormatParams) GetBlobElementIdOk() (*string, bool)`

GetBlobElementIdOk returns a tuple with the BlobElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobElementId

`func (o *BTTranslateFormatParams) SetBlobElementId(v string)`

SetBlobElementId sets BlobElementId field to given value.

### HasBlobElementId

`func (o *BTTranslateFormatParams) HasBlobElementId() bool`

HasBlobElementId returns a boolean if a field has been set.

### GetBlobMicroversionId

`func (o *BTTranslateFormatParams) GetBlobMicroversionId() string`

GetBlobMicroversionId returns the BlobMicroversionId field if non-nil, zero value otherwise.

### GetBlobMicroversionIdOk

`func (o *BTTranslateFormatParams) GetBlobMicroversionIdOk() (*string, bool)`

GetBlobMicroversionIdOk returns a tuple with the BlobMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobMicroversionId

`func (o *BTTranslateFormatParams) SetBlobMicroversionId(v string)`

SetBlobMicroversionId sets BlobMicroversionId field to given value.

### HasBlobMicroversionId

`func (o *BTTranslateFormatParams) HasBlobMicroversionId() bool`

HasBlobMicroversionId returns a boolean if a field has been set.

### GetCloudObjectId

`func (o *BTTranslateFormatParams) GetCloudObjectId() string`

GetCloudObjectId returns the CloudObjectId field if non-nil, zero value otherwise.

### GetCloudObjectIdOk

`func (o *BTTranslateFormatParams) GetCloudObjectIdOk() (*string, bool)`

GetCloudObjectIdOk returns a tuple with the CloudObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudObjectId

`func (o *BTTranslateFormatParams) SetCloudObjectId(v string)`

SetCloudObjectId sets CloudObjectId field to given value.

### HasCloudObjectId

`func (o *BTTranslateFormatParams) HasCloudObjectId() bool`

HasCloudObjectId returns a boolean if a field has been set.

### GetCloudStorageAccountId

`func (o *BTTranslateFormatParams) GetCloudStorageAccountId() string`

GetCloudStorageAccountId returns the CloudStorageAccountId field if non-nil, zero value otherwise.

### GetCloudStorageAccountIdOk

`func (o *BTTranslateFormatParams) GetCloudStorageAccountIdOk() (*string, bool)`

GetCloudStorageAccountIdOk returns a tuple with the CloudStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageAccountId

`func (o *BTTranslateFormatParams) SetCloudStorageAccountId(v string)`

SetCloudStorageAccountId sets CloudStorageAccountId field to given value.

### HasCloudStorageAccountId

`func (o *BTTranslateFormatParams) HasCloudStorageAccountId() bool`

HasCloudStorageAccountId returns a boolean if a field has been set.

### GetColorMethod

`func (o *BTTranslateFormatParams) GetColorMethod() string`

GetColorMethod returns the ColorMethod field if non-nil, zero value otherwise.

### GetColorMethodOk

`func (o *BTTranslateFormatParams) GetColorMethodOk() (*string, bool)`

GetColorMethodOk returns a tuple with the ColorMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorMethod

`func (o *BTTranslateFormatParams) SetColorMethod(v string)`

SetColorMethod sets ColorMethod field to given value.

### HasColorMethod

`func (o *BTTranslateFormatParams) HasColorMethod() bool`

HasColorMethod returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTTranslateFormatParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTTranslateFormatParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTTranslateFormatParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTTranslateFormatParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConnectionId

`func (o *BTTranslateFormatParams) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BTTranslateFormatParams) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BTTranslateFormatParams) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *BTTranslateFormatParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreateComposite

`func (o *BTTranslateFormatParams) GetCreateComposite() bool`

GetCreateComposite returns the CreateComposite field if non-nil, zero value otherwise.

### GetCreateCompositeOk

`func (o *BTTranslateFormatParams) GetCreateCompositeOk() (*bool, bool)`

GetCreateCompositeOk returns a tuple with the CreateComposite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateComposite

`func (o *BTTranslateFormatParams) SetCreateComposite(v bool)`

SetCreateComposite sets CreateComposite field to given value.

### HasCreateComposite

`func (o *BTTranslateFormatParams) HasCreateComposite() bool`

HasCreateComposite returns a boolean if a field has been set.

### GetCurrentSheetOnly

`func (o *BTTranslateFormatParams) GetCurrentSheetOnly() bool`

GetCurrentSheetOnly returns the CurrentSheetOnly field if non-nil, zero value otherwise.

### GetCurrentSheetOnlyOk

`func (o *BTTranslateFormatParams) GetCurrentSheetOnlyOk() (*bool, bool)`

GetCurrentSheetOnlyOk returns a tuple with the CurrentSheetOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSheetOnly

`func (o *BTTranslateFormatParams) SetCurrentSheetOnly(v bool)`

SetCurrentSheetOnly sets CurrentSheetOnly field to given value.

### HasCurrentSheetOnly

`func (o *BTTranslateFormatParams) HasCurrentSheetOnly() bool`

HasCurrentSheetOnly returns a boolean if a field has been set.

### GetDestinationName

`func (o *BTTranslateFormatParams) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *BTTranslateFormatParams) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *BTTranslateFormatParams) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *BTTranslateFormatParams) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetDistanceTolerance

`func (o *BTTranslateFormatParams) GetDistanceTolerance() float64`

GetDistanceTolerance returns the DistanceTolerance field if non-nil, zero value otherwise.

### GetDistanceToleranceOk

`func (o *BTTranslateFormatParams) GetDistanceToleranceOk() (*float64, bool)`

GetDistanceToleranceOk returns a tuple with the DistanceTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceTolerance

`func (o *BTTranslateFormatParams) SetDistanceTolerance(v float64)`

SetDistanceTolerance sets DistanceTolerance field to given value.

### HasDistanceTolerance

`func (o *BTTranslateFormatParams) HasDistanceTolerance() bool`

HasDistanceTolerance returns a boolean if a field has been set.

### GetElementId

`func (o *BTTranslateFormatParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTTranslateFormatParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTTranslateFormatParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTTranslateFormatParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementIds

`func (o *BTTranslateFormatParams) GetElementIds() []string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *BTTranslateFormatParams) GetElementIdsOk() (*[]string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *BTTranslateFormatParams) SetElementIds(v []string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *BTTranslateFormatParams) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### GetEmailLink

`func (o *BTTranslateFormatParams) GetEmailLink() bool`

GetEmailLink returns the EmailLink field if non-nil, zero value otherwise.

### GetEmailLinkOk

`func (o *BTTranslateFormatParams) GetEmailLinkOk() (*bool, bool)`

GetEmailLinkOk returns a tuple with the EmailLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLink

`func (o *BTTranslateFormatParams) SetEmailLink(v bool)`

SetEmailLink sets EmailLink field to given value.

### HasEmailLink

`func (o *BTTranslateFormatParams) HasEmailLink() bool`

HasEmailLink returns a boolean if a field has been set.

### GetEmailMessage

`func (o *BTTranslateFormatParams) GetEmailMessage() string`

GetEmailMessage returns the EmailMessage field if non-nil, zero value otherwise.

### GetEmailMessageOk

`func (o *BTTranslateFormatParams) GetEmailMessageOk() (*string, bool)`

GetEmailMessageOk returns a tuple with the EmailMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMessage

`func (o *BTTranslateFormatParams) SetEmailMessage(v string)`

SetEmailMessage sets EmailMessage field to given value.

### HasEmailMessage

`func (o *BTTranslateFormatParams) HasEmailMessage() bool`

HasEmailMessage returns a boolean if a field has been set.

### GetEmailSubject

`func (o *BTTranslateFormatParams) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *BTTranslateFormatParams) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *BTTranslateFormatParams) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *BTTranslateFormatParams) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailTo

`func (o *BTTranslateFormatParams) GetEmailTo() []string`

GetEmailTo returns the EmailTo field if non-nil, zero value otherwise.

### GetEmailToOk

`func (o *BTTranslateFormatParams) GetEmailToOk() (*[]string, bool)`

GetEmailToOk returns a tuple with the EmailTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTo

`func (o *BTTranslateFormatParams) SetEmailTo(v []string)`

SetEmailTo sets EmailTo field to given value.

### HasEmailTo

`func (o *BTTranslateFormatParams) HasEmailTo() bool`

HasEmailTo returns a boolean if a field has been set.

### GetExtractAssemblyHierarchy

`func (o *BTTranslateFormatParams) GetExtractAssemblyHierarchy() bool`

GetExtractAssemblyHierarchy returns the ExtractAssemblyHierarchy field if non-nil, zero value otherwise.

### GetExtractAssemblyHierarchyOk

`func (o *BTTranslateFormatParams) GetExtractAssemblyHierarchyOk() (*bool, bool)`

GetExtractAssemblyHierarchyOk returns a tuple with the ExtractAssemblyHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractAssemblyHierarchy

`func (o *BTTranslateFormatParams) SetExtractAssemblyHierarchy(v bool)`

SetExtractAssemblyHierarchy sets ExtractAssemblyHierarchy field to given value.

### HasExtractAssemblyHierarchy

`func (o *BTTranslateFormatParams) HasExtractAssemblyHierarchy() bool`

HasExtractAssemblyHierarchy returns a boolean if a field has been set.

### GetFlatten

`func (o *BTTranslateFormatParams) GetFlatten() bool`

GetFlatten returns the Flatten field if non-nil, zero value otherwise.

### GetFlattenOk

`func (o *BTTranslateFormatParams) GetFlattenOk() (*bool, bool)`

GetFlattenOk returns a tuple with the Flatten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatten

`func (o *BTTranslateFormatParams) SetFlatten(v bool)`

SetFlatten sets Flatten field to given value.

### HasFlatten

`func (o *BTTranslateFormatParams) HasFlatten() bool`

HasFlatten returns a boolean if a field has been set.

### GetFlattenAssemblies

`func (o *BTTranslateFormatParams) GetFlattenAssemblies() bool`

GetFlattenAssemblies returns the FlattenAssemblies field if non-nil, zero value otherwise.

### GetFlattenAssembliesOk

`func (o *BTTranslateFormatParams) GetFlattenAssembliesOk() (*bool, bool)`

GetFlattenAssembliesOk returns a tuple with the FlattenAssemblies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenAssemblies

`func (o *BTTranslateFormatParams) SetFlattenAssemblies(v bool)`

SetFlattenAssemblies sets FlattenAssemblies field to given value.

### HasFlattenAssemblies

`func (o *BTTranslateFormatParams) HasFlattenAssemblies() bool`

HasFlattenAssemblies returns a boolean if a field has been set.

### GetForeignId

`func (o *BTTranslateFormatParams) GetForeignId() string`

GetForeignId returns the ForeignId field if non-nil, zero value otherwise.

### GetForeignIdOk

`func (o *BTTranslateFormatParams) GetForeignIdOk() (*string, bool)`

GetForeignIdOk returns a tuple with the ForeignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignId

`func (o *BTTranslateFormatParams) SetForeignId(v string)`

SetForeignId sets ForeignId field to given value.

### HasForeignId

`func (o *BTTranslateFormatParams) HasForeignId() bool`

HasForeignId returns a boolean if a field has been set.

### GetFormatName

`func (o *BTTranslateFormatParams) GetFormatName() string`

GetFormatName returns the FormatName field if non-nil, zero value otherwise.

### GetFormatNameOk

`func (o *BTTranslateFormatParams) GetFormatNameOk() (*string, bool)`

GetFormatNameOk returns a tuple with the FormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatName

`func (o *BTTranslateFormatParams) SetFormatName(v string)`

SetFormatName sets FormatName field to given value.


### GetFromUserId

`func (o *BTTranslateFormatParams) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *BTTranslateFormatParams) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *BTTranslateFormatParams) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *BTTranslateFormatParams) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetGetyAxisIsUp

`func (o *BTTranslateFormatParams) GetGetyAxisIsUp() bool`

GetGetyAxisIsUp returns the GetyAxisIsUp field if non-nil, zero value otherwise.

### GetGetyAxisIsUpOk

`func (o *BTTranslateFormatParams) GetGetyAxisIsUpOk() (*bool, bool)`

GetGetyAxisIsUpOk returns a tuple with the GetyAxisIsUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetyAxisIsUp

`func (o *BTTranslateFormatParams) SetGetyAxisIsUp(v bool)`

SetGetyAxisIsUp sets GetyAxisIsUp field to given value.

### HasGetyAxisIsUp

`func (o *BTTranslateFormatParams) HasGetyAxisIsUp() bool`

HasGetyAxisIsUp returns a boolean if a field has been set.

### GetGrouping

`func (o *BTTranslateFormatParams) GetGrouping() bool`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BTTranslateFormatParams) GetGroupingOk() (*bool, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BTTranslateFormatParams) SetGrouping(v bool)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *BTTranslateFormatParams) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetHideInspectionItems

`func (o *BTTranslateFormatParams) GetHideInspectionItems() bool`

GetHideInspectionItems returns the HideInspectionItems field if non-nil, zero value otherwise.

### GetHideInspectionItemsOk

`func (o *BTTranslateFormatParams) GetHideInspectionItemsOk() (*bool, bool)`

GetHideInspectionItemsOk returns a tuple with the HideInspectionItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideInspectionItems

`func (o *BTTranslateFormatParams) SetHideInspectionItems(v bool)`

SetHideInspectionItems sets HideInspectionItems field to given value.

### HasHideInspectionItems

`func (o *BTTranslateFormatParams) HasHideInspectionItems() bool`

HasHideInspectionItems returns a boolean if a field has been set.

### GetIgnoreExportRulesForContents

`func (o *BTTranslateFormatParams) GetIgnoreExportRulesForContents() bool`

GetIgnoreExportRulesForContents returns the IgnoreExportRulesForContents field if non-nil, zero value otherwise.

### GetIgnoreExportRulesForContentsOk

`func (o *BTTranslateFormatParams) GetIgnoreExportRulesForContentsOk() (*bool, bool)`

GetIgnoreExportRulesForContentsOk returns a tuple with the IgnoreExportRulesForContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreExportRulesForContents

`func (o *BTTranslateFormatParams) SetIgnoreExportRulesForContents(v bool)`

SetIgnoreExportRulesForContents sets IgnoreExportRulesForContents field to given value.

### HasIgnoreExportRulesForContents

`func (o *BTTranslateFormatParams) HasIgnoreExportRulesForContents() bool`

HasIgnoreExportRulesForContents returns a boolean if a field has been set.

### GetImageHeight

`func (o *BTTranslateFormatParams) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *BTTranslateFormatParams) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *BTTranslateFormatParams) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.

### HasImageHeight

`func (o *BTTranslateFormatParams) HasImageHeight() bool`

HasImageHeight returns a boolean if a field has been set.

### GetImageWidth

`func (o *BTTranslateFormatParams) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *BTTranslateFormatParams) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *BTTranslateFormatParams) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.

### HasImageWidth

`func (o *BTTranslateFormatParams) HasImageWidth() bool`

HasImageWidth returns a boolean if a field has been set.

### GetImportAppearances

`func (o *BTTranslateFormatParams) GetImportAppearances() bool`

GetImportAppearances returns the ImportAppearances field if non-nil, zero value otherwise.

### GetImportAppearancesOk

`func (o *BTTranslateFormatParams) GetImportAppearancesOk() (*bool, bool)`

GetImportAppearancesOk returns a tuple with the ImportAppearances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAppearances

`func (o *BTTranslateFormatParams) SetImportAppearances(v bool)`

SetImportAppearances sets ImportAppearances field to given value.

### HasImportAppearances

`func (o *BTTranslateFormatParams) HasImportAppearances() bool`

HasImportAppearances returns a boolean if a field has been set.

### GetImportInBackground

`func (o *BTTranslateFormatParams) GetImportInBackground() bool`

GetImportInBackground returns the ImportInBackground field if non-nil, zero value otherwise.

### GetImportInBackgroundOk

`func (o *BTTranslateFormatParams) GetImportInBackgroundOk() (*bool, bool)`

GetImportInBackgroundOk returns a tuple with the ImportInBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportInBackground

`func (o *BTTranslateFormatParams) SetImportInBackground(v bool)`

SetImportInBackground sets ImportInBackground field to given value.

### HasImportInBackground

`func (o *BTTranslateFormatParams) HasImportInBackground() bool`

HasImportInBackground returns a boolean if a field has been set.

### GetImportWithinDocument

`func (o *BTTranslateFormatParams) GetImportWithinDocument() bool`

GetImportWithinDocument returns the ImportWithinDocument field if non-nil, zero value otherwise.

### GetImportWithinDocumentOk

`func (o *BTTranslateFormatParams) GetImportWithinDocumentOk() (*bool, bool)`

GetImportWithinDocumentOk returns a tuple with the ImportWithinDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportWithinDocument

`func (o *BTTranslateFormatParams) SetImportWithinDocument(v bool)`

SetImportWithinDocument sets ImportWithinDocument field to given value.

### HasImportWithinDocument

`func (o *BTTranslateFormatParams) HasImportWithinDocument() bool`

HasImportWithinDocument returns a boolean if a field has been set.

### GetIncludeExportIds

`func (o *BTTranslateFormatParams) GetIncludeExportIds() bool`

GetIncludeExportIds returns the IncludeExportIds field if non-nil, zero value otherwise.

### GetIncludeExportIdsOk

`func (o *BTTranslateFormatParams) GetIncludeExportIdsOk() (*bool, bool)`

GetIncludeExportIdsOk returns a tuple with the IncludeExportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExportIds

`func (o *BTTranslateFormatParams) SetIncludeExportIds(v bool)`

SetIncludeExportIds sets IncludeExportIds field to given value.

### HasIncludeExportIds

`func (o *BTTranslateFormatParams) HasIncludeExportIds() bool`

HasIncludeExportIds returns a boolean if a field has been set.

### GetJoinAdjacentSurfaces

`func (o *BTTranslateFormatParams) GetJoinAdjacentSurfaces() bool`

GetJoinAdjacentSurfaces returns the JoinAdjacentSurfaces field if non-nil, zero value otherwise.

### GetJoinAdjacentSurfacesOk

`func (o *BTTranslateFormatParams) GetJoinAdjacentSurfacesOk() (*bool, bool)`

GetJoinAdjacentSurfacesOk returns a tuple with the JoinAdjacentSurfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAdjacentSurfaces

`func (o *BTTranslateFormatParams) SetJoinAdjacentSurfaces(v bool)`

SetJoinAdjacentSurfaces sets JoinAdjacentSurfaces field to given value.

### HasJoinAdjacentSurfaces

`func (o *BTTranslateFormatParams) HasJoinAdjacentSurfaces() bool`

HasJoinAdjacentSurfaces returns a boolean if a field has been set.

### GetLevel

`func (o *BTTranslateFormatParams) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *BTTranslateFormatParams) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *BTTranslateFormatParams) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *BTTranslateFormatParams) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLinkDocumentId

`func (o *BTTranslateFormatParams) GetLinkDocumentId() string`

GetLinkDocumentId returns the LinkDocumentId field if non-nil, zero value otherwise.

### GetLinkDocumentIdOk

`func (o *BTTranslateFormatParams) GetLinkDocumentIdOk() (*string, bool)`

GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentId

`func (o *BTTranslateFormatParams) SetLinkDocumentId(v string)`

SetLinkDocumentId sets LinkDocumentId field to given value.

### HasLinkDocumentId

`func (o *BTTranslateFormatParams) HasLinkDocumentId() bool`

HasLinkDocumentId returns a boolean if a field has been set.

### GetLinkDocumentWorkspaceId

`func (o *BTTranslateFormatParams) GetLinkDocumentWorkspaceId() string`

GetLinkDocumentWorkspaceId returns the LinkDocumentWorkspaceId field if non-nil, zero value otherwise.

### GetLinkDocumentWorkspaceIdOk

`func (o *BTTranslateFormatParams) GetLinkDocumentWorkspaceIdOk() (*string, bool)`

GetLinkDocumentWorkspaceIdOk returns a tuple with the LinkDocumentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentWorkspaceId

`func (o *BTTranslateFormatParams) SetLinkDocumentWorkspaceId(v string)`

SetLinkDocumentWorkspaceId sets LinkDocumentWorkspaceId field to given value.

### HasLinkDocumentWorkspaceId

`func (o *BTTranslateFormatParams) HasLinkDocumentWorkspaceId() bool`

HasLinkDocumentWorkspaceId returns a boolean if a field has been set.

### GetMaximumChordLength

`func (o *BTTranslateFormatParams) GetMaximumChordLength() float64`

GetMaximumChordLength returns the MaximumChordLength field if non-nil, zero value otherwise.

### GetMaximumChordLengthOk

`func (o *BTTranslateFormatParams) GetMaximumChordLengthOk() (*float64, bool)`

GetMaximumChordLengthOk returns a tuple with the MaximumChordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumChordLength

`func (o *BTTranslateFormatParams) SetMaximumChordLength(v float64)`

SetMaximumChordLength sets MaximumChordLength field to given value.

### HasMaximumChordLength

`func (o *BTTranslateFormatParams) HasMaximumChordLength() bool`

HasMaximumChordLength returns a boolean if a field has been set.

### GetNotifyUser

`func (o *BTTranslateFormatParams) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *BTTranslateFormatParams) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *BTTranslateFormatParams) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *BTTranslateFormatParams) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetOnePartPerDoc

`func (o *BTTranslateFormatParams) GetOnePartPerDoc() bool`

GetOnePartPerDoc returns the OnePartPerDoc field if non-nil, zero value otherwise.

### GetOnePartPerDocOk

`func (o *BTTranslateFormatParams) GetOnePartPerDocOk() (*bool, bool)`

GetOnePartPerDocOk returns a tuple with the OnePartPerDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnePartPerDoc

`func (o *BTTranslateFormatParams) SetOnePartPerDoc(v bool)`

SetOnePartPerDoc sets OnePartPerDoc field to given value.

### HasOnePartPerDoc

`func (o *BTTranslateFormatParams) HasOnePartPerDoc() bool`

HasOnePartPerDoc returns a boolean if a field has been set.

### GetOriginalForeignId

`func (o *BTTranslateFormatParams) GetOriginalForeignId() string`

GetOriginalForeignId returns the OriginalForeignId field if non-nil, zero value otherwise.

### GetOriginalForeignIdOk

`func (o *BTTranslateFormatParams) GetOriginalForeignIdOk() (*string, bool)`

GetOriginalForeignIdOk returns a tuple with the OriginalForeignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalForeignId

`func (o *BTTranslateFormatParams) SetOriginalForeignId(v string)`

SetOriginalForeignId sets OriginalForeignId field to given value.

### HasOriginalForeignId

`func (o *BTTranslateFormatParams) HasOriginalForeignId() bool`

HasOriginalForeignId returns a boolean if a field has been set.

### GetParentId

`func (o *BTTranslateFormatParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTTranslateFormatParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTTranslateFormatParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTTranslateFormatParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartIds

`func (o *BTTranslateFormatParams) GetPartIds() string`

GetPartIds returns the PartIds field if non-nil, zero value otherwise.

### GetPartIdsOk

`func (o *BTTranslateFormatParams) GetPartIdsOk() (*string, bool)`

GetPartIdsOk returns a tuple with the PartIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIds

`func (o *BTTranslateFormatParams) SetPartIds(v string)`

SetPartIds sets PartIds field to given value.

### HasPartIds

`func (o *BTTranslateFormatParams) HasPartIds() bool`

HasPartIds returns a boolean if a field has been set.

### GetPassword

`func (o *BTTranslateFormatParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BTTranslateFormatParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BTTranslateFormatParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BTTranslateFormatParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordRequired

`func (o *BTTranslateFormatParams) GetPasswordRequired() bool`

GetPasswordRequired returns the PasswordRequired field if non-nil, zero value otherwise.

### GetPasswordRequiredOk

`func (o *BTTranslateFormatParams) GetPasswordRequiredOk() (*bool, bool)`

GetPasswordRequiredOk returns a tuple with the PasswordRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRequired

`func (o *BTTranslateFormatParams) SetPasswordRequired(v bool)`

SetPasswordRequired sets PasswordRequired field to given value.

### HasPasswordRequired

`func (o *BTTranslateFormatParams) HasPasswordRequired() bool`

HasPasswordRequired returns a boolean if a field has been set.

### GetPdfVersion

`func (o *BTTranslateFormatParams) GetPdfVersion() string`

GetPdfVersion returns the PdfVersion field if non-nil, zero value otherwise.

### GetPdfVersionOk

`func (o *BTTranslateFormatParams) GetPdfVersionOk() (*string, bool)`

GetPdfVersionOk returns a tuple with the PdfVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfVersion

`func (o *BTTranslateFormatParams) SetPdfVersion(v string)`

SetPdfVersion sets PdfVersion field to given value.

### HasPdfVersion

`func (o *BTTranslateFormatParams) HasPdfVersion() bool`

HasPdfVersion returns a boolean if a field has been set.

### GetProcessedForeignId

`func (o *BTTranslateFormatParams) GetProcessedForeignId() string`

GetProcessedForeignId returns the ProcessedForeignId field if non-nil, zero value otherwise.

### GetProcessedForeignIdOk

`func (o *BTTranslateFormatParams) GetProcessedForeignIdOk() (*string, bool)`

GetProcessedForeignIdOk returns a tuple with the ProcessedForeignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedForeignId

`func (o *BTTranslateFormatParams) SetProcessedForeignId(v string)`

SetProcessedForeignId sets ProcessedForeignId field to given value.

### HasProcessedForeignId

`func (o *BTTranslateFormatParams) HasProcessedForeignId() bool`

HasProcessedForeignId returns a boolean if a field has been set.

### GetProjectId

`func (o *BTTranslateFormatParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTTranslateFormatParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTTranslateFormatParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTTranslateFormatParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProxyDocumentId

`func (o *BTTranslateFormatParams) GetProxyDocumentId() string`

GetProxyDocumentId returns the ProxyDocumentId field if non-nil, zero value otherwise.

### GetProxyDocumentIdOk

`func (o *BTTranslateFormatParams) GetProxyDocumentIdOk() (*string, bool)`

GetProxyDocumentIdOk returns a tuple with the ProxyDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDocumentId

`func (o *BTTranslateFormatParams) SetProxyDocumentId(v string)`

SetProxyDocumentId sets ProxyDocumentId field to given value.

### HasProxyDocumentId

`func (o *BTTranslateFormatParams) HasProxyDocumentId() bool`

HasProxyDocumentId returns a boolean if a field has been set.

### GetProxyElementId

`func (o *BTTranslateFormatParams) GetProxyElementId() string`

GetProxyElementId returns the ProxyElementId field if non-nil, zero value otherwise.

### GetProxyElementIdOk

`func (o *BTTranslateFormatParams) GetProxyElementIdOk() (*string, bool)`

GetProxyElementIdOk returns a tuple with the ProxyElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyElementId

`func (o *BTTranslateFormatParams) SetProxyElementId(v string)`

SetProxyElementId sets ProxyElementId field to given value.

### HasProxyElementId

`func (o *BTTranslateFormatParams) HasProxyElementId() bool`

HasProxyElementId returns a boolean if a field has been set.

### GetProxyWorkspaceVersion

`func (o *BTTranslateFormatParams) GetProxyWorkspaceVersion() string`

GetProxyWorkspaceVersion returns the ProxyWorkspaceVersion field if non-nil, zero value otherwise.

### GetProxyWorkspaceVersionOk

`func (o *BTTranslateFormatParams) GetProxyWorkspaceVersionOk() (*string, bool)`

GetProxyWorkspaceVersionOk returns a tuple with the ProxyWorkspaceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkspaceVersion

`func (o *BTTranslateFormatParams) SetProxyWorkspaceVersion(v string)`

SetProxyWorkspaceVersion sets ProxyWorkspaceVersion field to given value.

### HasProxyWorkspaceVersion

`func (o *BTTranslateFormatParams) HasProxyWorkspaceVersion() bool`

HasProxyWorkspaceVersion returns a boolean if a field has been set.

### GetProxyWorkspaceVersionId

`func (o *BTTranslateFormatParams) GetProxyWorkspaceVersionId() string`

GetProxyWorkspaceVersionId returns the ProxyWorkspaceVersionId field if non-nil, zero value otherwise.

### GetProxyWorkspaceVersionIdOk

`func (o *BTTranslateFormatParams) GetProxyWorkspaceVersionIdOk() (*string, bool)`

GetProxyWorkspaceVersionIdOk returns a tuple with the ProxyWorkspaceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkspaceVersionId

`func (o *BTTranslateFormatParams) SetProxyWorkspaceVersionId(v string)`

SetProxyWorkspaceVersionId sets ProxyWorkspaceVersionId field to given value.

### HasProxyWorkspaceVersionId

`func (o *BTTranslateFormatParams) HasProxyWorkspaceVersionId() bool`

HasProxyWorkspaceVersionId returns a boolean if a field has been set.

### GetSelectablePdfText

`func (o *BTTranslateFormatParams) GetSelectablePdfText() bool`

GetSelectablePdfText returns the SelectablePdfText field if non-nil, zero value otherwise.

### GetSelectablePdfTextOk

`func (o *BTTranslateFormatParams) GetSelectablePdfTextOk() (*bool, bool)`

GetSelectablePdfTextOk returns a tuple with the SelectablePdfText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectablePdfText

`func (o *BTTranslateFormatParams) SetSelectablePdfText(v bool)`

SetSelectablePdfText sets SelectablePdfText field to given value.

### HasSelectablePdfText

`func (o *BTTranslateFormatParams) HasSelectablePdfText() bool`

HasSelectablePdfText returns a boolean if a field has been set.

### GetSendCopyToMe

`func (o *BTTranslateFormatParams) GetSendCopyToMe() bool`

GetSendCopyToMe returns the SendCopyToMe field if non-nil, zero value otherwise.

### GetSendCopyToMeOk

`func (o *BTTranslateFormatParams) GetSendCopyToMeOk() (*bool, bool)`

GetSendCopyToMeOk returns a tuple with the SendCopyToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCopyToMe

`func (o *BTTranslateFormatParams) SetSendCopyToMe(v bool)`

SetSendCopyToMe sets SendCopyToMe field to given value.

### HasSendCopyToMe

`func (o *BTTranslateFormatParams) HasSendCopyToMe() bool`

HasSendCopyToMe returns a boolean if a field has been set.

### GetSheetIndices

`func (o *BTTranslateFormatParams) GetSheetIndices() []int32`

GetSheetIndices returns the SheetIndices field if non-nil, zero value otherwise.

### GetSheetIndicesOk

`func (o *BTTranslateFormatParams) GetSheetIndicesOk() (*[]int32, bool)`

GetSheetIndicesOk returns a tuple with the SheetIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetIndices

`func (o *BTTranslateFormatParams) SetSheetIndices(v []int32)`

SetSheetIndices sets SheetIndices field to given value.

### HasSheetIndices

`func (o *BTTranslateFormatParams) HasSheetIndices() bool`

HasSheetIndices returns a boolean if a field has been set.

### GetShowOverriddenDimensions

`func (o *BTTranslateFormatParams) GetShowOverriddenDimensions() bool`

GetShowOverriddenDimensions returns the ShowOverriddenDimensions field if non-nil, zero value otherwise.

### GetShowOverriddenDimensionsOk

`func (o *BTTranslateFormatParams) GetShowOverriddenDimensionsOk() (*bool, bool)`

GetShowOverriddenDimensionsOk returns a tuple with the ShowOverriddenDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOverriddenDimensions

`func (o *BTTranslateFormatParams) SetShowOverriddenDimensions(v bool)`

SetShowOverriddenDimensions sets ShowOverriddenDimensions field to given value.

### HasShowOverriddenDimensions

`func (o *BTTranslateFormatParams) HasShowOverriddenDimensions() bool`

HasShowOverriddenDimensions returns a boolean if a field has been set.

### GetSkipBodyshop

`func (o *BTTranslateFormatParams) GetSkipBodyshop() bool`

GetSkipBodyshop returns the SkipBodyshop field if non-nil, zero value otherwise.

### GetSkipBodyshopOk

`func (o *BTTranslateFormatParams) GetSkipBodyshopOk() (*bool, bool)`

GetSkipBodyshopOk returns a tuple with the SkipBodyshop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipBodyshop

`func (o *BTTranslateFormatParams) SetSkipBodyshop(v bool)`

SetSkipBodyshop sets SkipBodyshop field to given value.

### HasSkipBodyshop

`func (o *BTTranslateFormatParams) HasSkipBodyshop() bool`

HasSkipBodyshop returns a boolean if a field has been set.

### GetSourceName

`func (o *BTTranslateFormatParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BTTranslateFormatParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BTTranslateFormatParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *BTTranslateFormatParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSpecifyUnits

`func (o *BTTranslateFormatParams) GetSpecifyUnits() bool`

GetSpecifyUnits returns the SpecifyUnits field if non-nil, zero value otherwise.

### GetSpecifyUnitsOk

`func (o *BTTranslateFormatParams) GetSpecifyUnitsOk() (*bool, bool)`

GetSpecifyUnitsOk returns a tuple with the SpecifyUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifyUnits

`func (o *BTTranslateFormatParams) SetSpecifyUnits(v bool)`

SetSpecifyUnits sets SpecifyUnits field to given value.

### HasSpecifyUnits

`func (o *BTTranslateFormatParams) HasSpecifyUnits() bool`

HasSpecifyUnits returns a boolean if a field has been set.

### GetSplinesAsPolylines

`func (o *BTTranslateFormatParams) GetSplinesAsPolylines() bool`

GetSplinesAsPolylines returns the SplinesAsPolylines field if non-nil, zero value otherwise.

### GetSplinesAsPolylinesOk

`func (o *BTTranslateFormatParams) GetSplinesAsPolylinesOk() (*bool, bool)`

GetSplinesAsPolylinesOk returns a tuple with the SplinesAsPolylines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplinesAsPolylines

`func (o *BTTranslateFormatParams) SetSplinesAsPolylines(v bool)`

SetSplinesAsPolylines sets SplinesAsPolylines field to given value.

### HasSplinesAsPolylines

`func (o *BTTranslateFormatParams) HasSplinesAsPolylines() bool`

HasSplinesAsPolylines returns a boolean if a field has been set.

### GetSplitAssembliesIntoMultipleDocuments

`func (o *BTTranslateFormatParams) GetSplitAssembliesIntoMultipleDocuments() bool`

GetSplitAssembliesIntoMultipleDocuments returns the SplitAssembliesIntoMultipleDocuments field if non-nil, zero value otherwise.

### GetSplitAssembliesIntoMultipleDocumentsOk

`func (o *BTTranslateFormatParams) GetSplitAssembliesIntoMultipleDocumentsOk() (*bool, bool)`

GetSplitAssembliesIntoMultipleDocumentsOk returns a tuple with the SplitAssembliesIntoMultipleDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAssembliesIntoMultipleDocuments

`func (o *BTTranslateFormatParams) SetSplitAssembliesIntoMultipleDocuments(v bool)`

SetSplitAssembliesIntoMultipleDocuments sets SplitAssembliesIntoMultipleDocuments field to given value.

### HasSplitAssembliesIntoMultipleDocuments

`func (o *BTTranslateFormatParams) HasSplitAssembliesIntoMultipleDocuments() bool`

HasSplitAssembliesIntoMultipleDocuments returns a boolean if a field has been set.

### GetStepParasolidPreprocessingOption

`func (o *BTTranslateFormatParams) GetStepParasolidPreprocessingOption() GBTPreProcessParasolidOption`

GetStepParasolidPreprocessingOption returns the StepParasolidPreprocessingOption field if non-nil, zero value otherwise.

### GetStepParasolidPreprocessingOptionOk

`func (o *BTTranslateFormatParams) GetStepParasolidPreprocessingOptionOk() (*GBTPreProcessParasolidOption, bool)`

GetStepParasolidPreprocessingOptionOk returns a tuple with the StepParasolidPreprocessingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepParasolidPreprocessingOption

`func (o *BTTranslateFormatParams) SetStepParasolidPreprocessingOption(v GBTPreProcessParasolidOption)`

SetStepParasolidPreprocessingOption sets StepParasolidPreprocessingOption field to given value.

### HasStepParasolidPreprocessingOption

`func (o *BTTranslateFormatParams) HasStepParasolidPreprocessingOption() bool`

HasStepParasolidPreprocessingOption returns a boolean if a field has been set.

### GetStepVersionString

`func (o *BTTranslateFormatParams) GetStepVersionString() string`

GetStepVersionString returns the StepVersionString field if non-nil, zero value otherwise.

### GetStepVersionStringOk

`func (o *BTTranslateFormatParams) GetStepVersionStringOk() (*string, bool)`

GetStepVersionStringOk returns a tuple with the StepVersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepVersionString

`func (o *BTTranslateFormatParams) SetStepVersionString(v string)`

SetStepVersionString sets StepVersionString field to given value.

### HasStepVersionString

`func (o *BTTranslateFormatParams) HasStepVersionString() bool`

HasStepVersionString returns a boolean if a field has been set.

### GetStoreInDocument

`func (o *BTTranslateFormatParams) GetStoreInDocument() bool`

GetStoreInDocument returns the StoreInDocument field if non-nil, zero value otherwise.

### GetStoreInDocumentOk

`func (o *BTTranslateFormatParams) GetStoreInDocumentOk() (*bool, bool)`

GetStoreInDocumentOk returns a tuple with the StoreInDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInDocument

`func (o *BTTranslateFormatParams) SetStoreInDocument(v bool)`

SetStoreInDocument sets StoreInDocument field to given value.

### HasStoreInDocument

`func (o *BTTranslateFormatParams) HasStoreInDocument() bool`

HasStoreInDocument returns a boolean if a field has been set.

### GetTextAsGeometry

`func (o *BTTranslateFormatParams) GetTextAsGeometry() bool`

GetTextAsGeometry returns the TextAsGeometry field if non-nil, zero value otherwise.

### GetTextAsGeometryOk

`func (o *BTTranslateFormatParams) GetTextAsGeometryOk() (*bool, bool)`

GetTextAsGeometryOk returns a tuple with the TextAsGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAsGeometry

`func (o *BTTranslateFormatParams) SetTextAsGeometry(v bool)`

SetTextAsGeometry sets TextAsGeometry field to given value.

### HasTextAsGeometry

`func (o *BTTranslateFormatParams) HasTextAsGeometry() bool`

HasTextAsGeometry returns a boolean if a field has been set.

### GetTextOption

`func (o *BTTranslateFormatParams) GetTextOption() string`

GetTextOption returns the TextOption field if non-nil, zero value otherwise.

### GetTextOptionOk

`func (o *BTTranslateFormatParams) GetTextOptionOk() (*string, bool)`

GetTextOptionOk returns a tuple with the TextOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextOption

`func (o *BTTranslateFormatParams) SetTextOption(v string)`

SetTextOption sets TextOption field to given value.

### HasTextOption

`func (o *BTTranslateFormatParams) HasTextOption() bool`

HasTextOption returns a boolean if a field has been set.

### GetTriggerAutoDownload

`func (o *BTTranslateFormatParams) GetTriggerAutoDownload() bool`

GetTriggerAutoDownload returns the TriggerAutoDownload field if non-nil, zero value otherwise.

### GetTriggerAutoDownloadOk

`func (o *BTTranslateFormatParams) GetTriggerAutoDownloadOk() (*bool, bool)`

GetTriggerAutoDownloadOk returns a tuple with the TriggerAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAutoDownload

`func (o *BTTranslateFormatParams) SetTriggerAutoDownload(v bool)`

SetTriggerAutoDownload sets TriggerAutoDownload field to given value.

### HasTriggerAutoDownload

`func (o *BTTranslateFormatParams) HasTriggerAutoDownload() bool`

HasTriggerAutoDownload returns a boolean if a field has been set.

### GetUnit

`func (o *BTTranslateFormatParams) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BTTranslateFormatParams) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BTTranslateFormatParams) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *BTTranslateFormatParams) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUploadId

`func (o *BTTranslateFormatParams) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *BTTranslateFormatParams) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *BTTranslateFormatParams) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *BTTranslateFormatParams) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUseGltfCompression

`func (o *BTTranslateFormatParams) GetUseGltfCompression() bool`

GetUseGltfCompression returns the UseGltfCompression field if non-nil, zero value otherwise.

### GetUseGltfCompressionOk

`func (o *BTTranslateFormatParams) GetUseGltfCompressionOk() (*bool, bool)`

GetUseGltfCompressionOk returns a tuple with the UseGltfCompression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGltfCompression

`func (o *BTTranslateFormatParams) SetUseGltfCompression(v bool)`

SetUseGltfCompression sets UseGltfCompression field to given value.

### HasUseGltfCompression

`func (o *BTTranslateFormatParams) HasUseGltfCompression() bool`

HasUseGltfCompression returns a boolean if a field has been set.

### GetUseIgesCompatibilityMode

`func (o *BTTranslateFormatParams) GetUseIgesCompatibilityMode() bool`

GetUseIgesCompatibilityMode returns the UseIgesCompatibilityMode field if non-nil, zero value otherwise.

### GetUseIgesCompatibilityModeOk

`func (o *BTTranslateFormatParams) GetUseIgesCompatibilityModeOk() (*bool, bool)`

GetUseIgesCompatibilityModeOk returns a tuple with the UseIgesCompatibilityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIgesCompatibilityMode

`func (o *BTTranslateFormatParams) SetUseIgesCompatibilityMode(v bool)`

SetUseIgesCompatibilityMode sets UseIgesCompatibilityMode field to given value.

### HasUseIgesCompatibilityMode

`func (o *BTTranslateFormatParams) HasUseIgesCompatibilityMode() bool`

HasUseIgesCompatibilityMode returns a boolean if a field has been set.

### GetValidForDays

`func (o *BTTranslateFormatParams) GetValidForDays() int32`

GetValidForDays returns the ValidForDays field if non-nil, zero value otherwise.

### GetValidForDaysOk

`func (o *BTTranslateFormatParams) GetValidForDaysOk() (*int32, bool)`

GetValidForDaysOk returns a tuple with the ValidForDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidForDays

`func (o *BTTranslateFormatParams) SetValidForDays(v int32)`

SetValidForDays sets ValidForDays field to given value.

### HasValidForDays

`func (o *BTTranslateFormatParams) HasValidForDays() bool`

HasValidForDays returns a boolean if a field has been set.

### GetVersionString

`func (o *BTTranslateFormatParams) GetVersionString() string`

GetVersionString returns the VersionString field if non-nil, zero value otherwise.

### GetVersionStringOk

`func (o *BTTranslateFormatParams) GetVersionStringOk() (*string, bool)`

GetVersionStringOk returns a tuple with the VersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionString

`func (o *BTTranslateFormatParams) SetVersionString(v string)`

SetVersionString sets VersionString field to given value.

### HasVersionString

`func (o *BTTranslateFormatParams) HasVersionString() bool`

HasVersionString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


