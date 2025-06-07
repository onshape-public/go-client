# BTBSolidworksExportParams

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
**StoreInDocument** | Pointer to **bool** | Create a blob with exported file in the source document. | [optional] [default to true]
**TriggerAutoDownload** | Pointer to **bool** | Automatically download a translated file. | [optional] [default to false]

## Methods

### NewBTBSolidworksExportParams

`func NewBTBSolidworksExportParams() *BTBSolidworksExportParams`

NewBTBSolidworksExportParams instantiates a new BTBSolidworksExportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBSolidworksExportParamsWithDefaults

`func NewBTBSolidworksExportParamsWithDefaults() *BTBSolidworksExportParams`

NewBTBSolidworksExportParamsWithDefaults instantiates a new BTBSolidworksExportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedParams

`func (o *BTBSolidworksExportParams) GetAdvancedParams() BTBExportAdvancedParams`

GetAdvancedParams returns the AdvancedParams field if non-nil, zero value otherwise.

### GetAdvancedParamsOk

`func (o *BTBSolidworksExportParams) GetAdvancedParamsOk() (*BTBExportAdvancedParams, bool)`

GetAdvancedParamsOk returns a tuple with the AdvancedParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedParams

`func (o *BTBSolidworksExportParams) SetAdvancedParams(v BTBExportAdvancedParams)`

SetAdvancedParams sets AdvancedParams field to given value.

### HasAdvancedParams

`func (o *BTBSolidworksExportParams) HasAdvancedParams() bool`

HasAdvancedParams returns a boolean if a field has been set.

### GetCloudStorageOptions

`func (o *BTBSolidworksExportParams) GetCloudStorageOptions() BTBCloudStorageOptions`

GetCloudStorageOptions returns the CloudStorageOptions field if non-nil, zero value otherwise.

### GetCloudStorageOptionsOk

`func (o *BTBSolidworksExportParams) GetCloudStorageOptionsOk() (*BTBCloudStorageOptions, bool)`

GetCloudStorageOptionsOk returns a tuple with the CloudStorageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageOptions

`func (o *BTBSolidworksExportParams) SetCloudStorageOptions(v BTBCloudStorageOptions)`

SetCloudStorageOptions sets CloudStorageOptions field to given value.

### HasCloudStorageOptions

`func (o *BTBSolidworksExportParams) HasCloudStorageOptions() bool`

HasCloudStorageOptions returns a boolean if a field has been set.

### GetDestinationName

`func (o *BTBSolidworksExportParams) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *BTBSolidworksExportParams) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *BTBSolidworksExportParams) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *BTBSolidworksExportParams) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetEmailExportOptions

`func (o *BTBSolidworksExportParams) GetEmailExportOptions() BTBEmailExportOptions`

GetEmailExportOptions returns the EmailExportOptions field if non-nil, zero value otherwise.

### GetEmailExportOptionsOk

`func (o *BTBSolidworksExportParams) GetEmailExportOptionsOk() (*BTBEmailExportOptions, bool)`

GetEmailExportOptionsOk returns a tuple with the EmailExportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailExportOptions

`func (o *BTBSolidworksExportParams) SetEmailExportOptions(v BTBEmailExportOptions)`

SetEmailExportOptions sets EmailExportOptions field to given value.

### HasEmailExportOptions

`func (o *BTBSolidworksExportParams) HasEmailExportOptions() bool`

HasEmailExportOptions returns a boolean if a field has been set.

### GetExcludeHiddenEntities

`func (o *BTBSolidworksExportParams) GetExcludeHiddenEntities() bool`

GetExcludeHiddenEntities returns the ExcludeHiddenEntities field if non-nil, zero value otherwise.

### GetExcludeHiddenEntitiesOk

`func (o *BTBSolidworksExportParams) GetExcludeHiddenEntitiesOk() (*bool, bool)`

GetExcludeHiddenEntitiesOk returns a tuple with the ExcludeHiddenEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHiddenEntities

`func (o *BTBSolidworksExportParams) SetExcludeHiddenEntities(v bool)`

SetExcludeHiddenEntities sets ExcludeHiddenEntities field to given value.

### HasExcludeHiddenEntities

`func (o *BTBSolidworksExportParams) HasExcludeHiddenEntities() bool`

HasExcludeHiddenEntities returns a boolean if a field has been set.

### GetGrouping

`func (o *BTBSolidworksExportParams) GetGrouping() bool`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BTBSolidworksExportParams) GetGroupingOk() (*bool, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BTBSolidworksExportParams) SetGrouping(v bool)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *BTBSolidworksExportParams) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetIncludeExportIds

`func (o *BTBSolidworksExportParams) GetIncludeExportIds() bool`

GetIncludeExportIds returns the IncludeExportIds field if non-nil, zero value otherwise.

### GetIncludeExportIdsOk

`func (o *BTBSolidworksExportParams) GetIncludeExportIdsOk() (*bool, bool)`

GetIncludeExportIdsOk returns a tuple with the IncludeExportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExportIds

`func (o *BTBSolidworksExportParams) SetIncludeExportIds(v bool)`

SetIncludeExportIds sets IncludeExportIds field to given value.

### HasIncludeExportIds

`func (o *BTBSolidworksExportParams) HasIncludeExportIds() bool`

HasIncludeExportIds returns a boolean if a field has been set.

### GetIsYAxisUp

`func (o *BTBSolidworksExportParams) GetIsYAxisUp() bool`

GetIsYAxisUp returns the IsYAxisUp field if non-nil, zero value otherwise.

### GetIsYAxisUpOk

`func (o *BTBSolidworksExportParams) GetIsYAxisUpOk() (*bool, bool)`

GetIsYAxisUpOk returns a tuple with the IsYAxisUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYAxisUp

`func (o *BTBSolidworksExportParams) SetIsYAxisUp(v bool)`

SetIsYAxisUp sets IsYAxisUp field to given value.

### HasIsYAxisUp

`func (o *BTBSolidworksExportParams) HasIsYAxisUp() bool`

HasIsYAxisUp returns a boolean if a field has been set.

### GetNotifyUser

`func (o *BTBSolidworksExportParams) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *BTBSolidworksExportParams) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *BTBSolidworksExportParams) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *BTBSolidworksExportParams) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetStoreInDocument

`func (o *BTBSolidworksExportParams) GetStoreInDocument() bool`

GetStoreInDocument returns the StoreInDocument field if non-nil, zero value otherwise.

### GetStoreInDocumentOk

`func (o *BTBSolidworksExportParams) GetStoreInDocumentOk() (*bool, bool)`

GetStoreInDocumentOk returns a tuple with the StoreInDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInDocument

`func (o *BTBSolidworksExportParams) SetStoreInDocument(v bool)`

SetStoreInDocument sets StoreInDocument field to given value.

### HasStoreInDocument

`func (o *BTBSolidworksExportParams) HasStoreInDocument() bool`

HasStoreInDocument returns a boolean if a field has been set.

### GetTriggerAutoDownload

`func (o *BTBSolidworksExportParams) GetTriggerAutoDownload() bool`

GetTriggerAutoDownload returns the TriggerAutoDownload field if non-nil, zero value otherwise.

### GetTriggerAutoDownloadOk

`func (o *BTBSolidworksExportParams) GetTriggerAutoDownloadOk() (*bool, bool)`

GetTriggerAutoDownloadOk returns a tuple with the TriggerAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAutoDownload

`func (o *BTBSolidworksExportParams) SetTriggerAutoDownload(v bool)`

SetTriggerAutoDownload sets TriggerAutoDownload field to given value.

### HasTriggerAutoDownload

`func (o *BTBSolidworksExportParams) HasTriggerAutoDownload() bool`

HasTriggerAutoDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


