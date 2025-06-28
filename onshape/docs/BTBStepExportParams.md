# BTBStepExportParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedParams** | Pointer to [**BTBExportAdvancedParams**](BTBExportAdvancedParams.md) |  | [optional] 
**CloudStorageOptions** | Pointer to [**BTBCloudStorageOptions**](BTBCloudStorageOptions.md) |  | [optional] 
**DestinationName** | Pointer to **string** | The name of the exported file. | [optional] [default to "Untitled"]
**EmailExportOptions** | Pointer to [**BTBEmailExportOptions**](BTBEmailExportOptions.md) |  | [optional] 
**ExcludeHiddenEntities** | Pointer to **bool** | Whether or not to exclude hidden parts from export. | [optional] [default to false]
**Grouping** | Pointer to **bool** | Whether parts should be exported as a group or individually in a .zip file. | [optional] [default to true]
**IncludeExportIds** | Pointer to **bool** | Whether topology ids should be exported as parasolid attributes. | [optional] [default to false]
**IsYAxisUp** | Pointer to **bool** | Rotate model from Z-axis-up orientation to Y-axis-up. | [optional] [default to false]
**NotifyUser** | Pointer to **bool** | Send notification to the user client. | [optional] [default to true]
**StepParasolidPreprocessingOption** | Pointer to [**GBTPreProcessParasolidOption**](GBTPreProcessParasolidOption.md) |  | [optional] 
**StepUnit** | Pointer to [**GBTExportUnit**](GBTExportUnit.md) |  | [optional] [default to GBTExportUnitMeter]
**StepVersionString** | Pointer to **string** | Export STEP in version: &#x60;AP242&#x60; | &#x60;AP203&#x60; | &#x60;AP214&#x60; | [optional] [default to "AP242"]
**StoreInDocument** | Pointer to **bool** | Create a blob with exported file in the source document. | [optional] [default to true]
**TriggerAutoDownload** | Pointer to **bool** | Automatically download a translated file. | [optional] [default to false]

## Methods

### NewBTBStepExportParams

`func NewBTBStepExportParams() *BTBStepExportParams`

NewBTBStepExportParams instantiates a new BTBStepExportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBStepExportParamsWithDefaults

`func NewBTBStepExportParamsWithDefaults() *BTBStepExportParams`

NewBTBStepExportParamsWithDefaults instantiates a new BTBStepExportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedParams

`func (o *BTBStepExportParams) GetAdvancedParams() BTBExportAdvancedParams`

GetAdvancedParams returns the AdvancedParams field if non-nil, zero value otherwise.

### GetAdvancedParamsOk

`func (o *BTBStepExportParams) GetAdvancedParamsOk() (*BTBExportAdvancedParams, bool)`

GetAdvancedParamsOk returns a tuple with the AdvancedParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedParams

`func (o *BTBStepExportParams) SetAdvancedParams(v BTBExportAdvancedParams)`

SetAdvancedParams sets AdvancedParams field to given value.

### HasAdvancedParams

`func (o *BTBStepExportParams) HasAdvancedParams() bool`

HasAdvancedParams returns a boolean if a field has been set.

### GetCloudStorageOptions

`func (o *BTBStepExportParams) GetCloudStorageOptions() BTBCloudStorageOptions`

GetCloudStorageOptions returns the CloudStorageOptions field if non-nil, zero value otherwise.

### GetCloudStorageOptionsOk

`func (o *BTBStepExportParams) GetCloudStorageOptionsOk() (*BTBCloudStorageOptions, bool)`

GetCloudStorageOptionsOk returns a tuple with the CloudStorageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageOptions

`func (o *BTBStepExportParams) SetCloudStorageOptions(v BTBCloudStorageOptions)`

SetCloudStorageOptions sets CloudStorageOptions field to given value.

### HasCloudStorageOptions

`func (o *BTBStepExportParams) HasCloudStorageOptions() bool`

HasCloudStorageOptions returns a boolean if a field has been set.

### GetDestinationName

`func (o *BTBStepExportParams) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *BTBStepExportParams) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *BTBStepExportParams) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *BTBStepExportParams) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetEmailExportOptions

`func (o *BTBStepExportParams) GetEmailExportOptions() BTBEmailExportOptions`

GetEmailExportOptions returns the EmailExportOptions field if non-nil, zero value otherwise.

### GetEmailExportOptionsOk

`func (o *BTBStepExportParams) GetEmailExportOptionsOk() (*BTBEmailExportOptions, bool)`

GetEmailExportOptionsOk returns a tuple with the EmailExportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailExportOptions

`func (o *BTBStepExportParams) SetEmailExportOptions(v BTBEmailExportOptions)`

SetEmailExportOptions sets EmailExportOptions field to given value.

### HasEmailExportOptions

`func (o *BTBStepExportParams) HasEmailExportOptions() bool`

HasEmailExportOptions returns a boolean if a field has been set.

### GetExcludeHiddenEntities

`func (o *BTBStepExportParams) GetExcludeHiddenEntities() bool`

GetExcludeHiddenEntities returns the ExcludeHiddenEntities field if non-nil, zero value otherwise.

### GetExcludeHiddenEntitiesOk

`func (o *BTBStepExportParams) GetExcludeHiddenEntitiesOk() (*bool, bool)`

GetExcludeHiddenEntitiesOk returns a tuple with the ExcludeHiddenEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHiddenEntities

`func (o *BTBStepExportParams) SetExcludeHiddenEntities(v bool)`

SetExcludeHiddenEntities sets ExcludeHiddenEntities field to given value.

### HasExcludeHiddenEntities

`func (o *BTBStepExportParams) HasExcludeHiddenEntities() bool`

HasExcludeHiddenEntities returns a boolean if a field has been set.

### GetGrouping

`func (o *BTBStepExportParams) GetGrouping() bool`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BTBStepExportParams) GetGroupingOk() (*bool, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BTBStepExportParams) SetGrouping(v bool)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *BTBStepExportParams) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetIncludeExportIds

`func (o *BTBStepExportParams) GetIncludeExportIds() bool`

GetIncludeExportIds returns the IncludeExportIds field if non-nil, zero value otherwise.

### GetIncludeExportIdsOk

`func (o *BTBStepExportParams) GetIncludeExportIdsOk() (*bool, bool)`

GetIncludeExportIdsOk returns a tuple with the IncludeExportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExportIds

`func (o *BTBStepExportParams) SetIncludeExportIds(v bool)`

SetIncludeExportIds sets IncludeExportIds field to given value.

### HasIncludeExportIds

`func (o *BTBStepExportParams) HasIncludeExportIds() bool`

HasIncludeExportIds returns a boolean if a field has been set.

### GetIsYAxisUp

`func (o *BTBStepExportParams) GetIsYAxisUp() bool`

GetIsYAxisUp returns the IsYAxisUp field if non-nil, zero value otherwise.

### GetIsYAxisUpOk

`func (o *BTBStepExportParams) GetIsYAxisUpOk() (*bool, bool)`

GetIsYAxisUpOk returns a tuple with the IsYAxisUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYAxisUp

`func (o *BTBStepExportParams) SetIsYAxisUp(v bool)`

SetIsYAxisUp sets IsYAxisUp field to given value.

### HasIsYAxisUp

`func (o *BTBStepExportParams) HasIsYAxisUp() bool`

HasIsYAxisUp returns a boolean if a field has been set.

### GetNotifyUser

`func (o *BTBStepExportParams) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *BTBStepExportParams) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *BTBStepExportParams) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *BTBStepExportParams) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetStepParasolidPreprocessingOption

`func (o *BTBStepExportParams) GetStepParasolidPreprocessingOption() GBTPreProcessParasolidOption`

GetStepParasolidPreprocessingOption returns the StepParasolidPreprocessingOption field if non-nil, zero value otherwise.

### GetStepParasolidPreprocessingOptionOk

`func (o *BTBStepExportParams) GetStepParasolidPreprocessingOptionOk() (*GBTPreProcessParasolidOption, bool)`

GetStepParasolidPreprocessingOptionOk returns a tuple with the StepParasolidPreprocessingOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepParasolidPreprocessingOption

`func (o *BTBStepExportParams) SetStepParasolidPreprocessingOption(v GBTPreProcessParasolidOption)`

SetStepParasolidPreprocessingOption sets StepParasolidPreprocessingOption field to given value.

### HasStepParasolidPreprocessingOption

`func (o *BTBStepExportParams) HasStepParasolidPreprocessingOption() bool`

HasStepParasolidPreprocessingOption returns a boolean if a field has been set.

### GetStepUnit

`func (o *BTBStepExportParams) GetStepUnit() GBTExportUnit`

GetStepUnit returns the StepUnit field if non-nil, zero value otherwise.

### GetStepUnitOk

`func (o *BTBStepExportParams) GetStepUnitOk() (*GBTExportUnit, bool)`

GetStepUnitOk returns a tuple with the StepUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUnit

`func (o *BTBStepExportParams) SetStepUnit(v GBTExportUnit)`

SetStepUnit sets StepUnit field to given value.

### HasStepUnit

`func (o *BTBStepExportParams) HasStepUnit() bool`

HasStepUnit returns a boolean if a field has been set.

### GetStepVersionString

`func (o *BTBStepExportParams) GetStepVersionString() string`

GetStepVersionString returns the StepVersionString field if non-nil, zero value otherwise.

### GetStepVersionStringOk

`func (o *BTBStepExportParams) GetStepVersionStringOk() (*string, bool)`

GetStepVersionStringOk returns a tuple with the StepVersionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepVersionString

`func (o *BTBStepExportParams) SetStepVersionString(v string)`

SetStepVersionString sets StepVersionString field to given value.

### HasStepVersionString

`func (o *BTBStepExportParams) HasStepVersionString() bool`

HasStepVersionString returns a boolean if a field has been set.

### GetStoreInDocument

`func (o *BTBStepExportParams) GetStoreInDocument() bool`

GetStoreInDocument returns the StoreInDocument field if non-nil, zero value otherwise.

### GetStoreInDocumentOk

`func (o *BTBStepExportParams) GetStoreInDocumentOk() (*bool, bool)`

GetStoreInDocumentOk returns a tuple with the StoreInDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInDocument

`func (o *BTBStepExportParams) SetStoreInDocument(v bool)`

SetStoreInDocument sets StoreInDocument field to given value.

### HasStoreInDocument

`func (o *BTBStepExportParams) HasStoreInDocument() bool`

HasStoreInDocument returns a boolean if a field has been set.

### GetTriggerAutoDownload

`func (o *BTBStepExportParams) GetTriggerAutoDownload() bool`

GetTriggerAutoDownload returns the TriggerAutoDownload field if non-nil, zero value otherwise.

### GetTriggerAutoDownloadOk

`func (o *BTBStepExportParams) GetTriggerAutoDownloadOk() (*bool, bool)`

GetTriggerAutoDownloadOk returns a tuple with the TriggerAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAutoDownload

`func (o *BTBStepExportParams) SetTriggerAutoDownload(v bool)`

SetTriggerAutoDownload sets TriggerAutoDownload field to given value.

### HasTriggerAutoDownload

`func (o *BTBStepExportParams) HasTriggerAutoDownload() bool`

HasTriggerAutoDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


