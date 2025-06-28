# BTBGltfExportParams

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
**MeshParams** | Pointer to [**BTBExportMeshParams**](BTBExportMeshParams.md) |  | [optional] 
**NotifyUser** | Pointer to **bool** | Send notification to the user client. | [optional] [default to true]
**StoreInDocument** | Pointer to **bool** | Create a blob with exported file in the source document. | [optional] [default to true]
**TriggerAutoDownload** | Pointer to **bool** | Automatically download a translated file. | [optional] [default to false]

## Methods

### NewBTBGltfExportParams

`func NewBTBGltfExportParams() *BTBGltfExportParams`

NewBTBGltfExportParams instantiates a new BTBGltfExportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBGltfExportParamsWithDefaults

`func NewBTBGltfExportParamsWithDefaults() *BTBGltfExportParams`

NewBTBGltfExportParamsWithDefaults instantiates a new BTBGltfExportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedParams

`func (o *BTBGltfExportParams) GetAdvancedParams() BTBExportAdvancedParams`

GetAdvancedParams returns the AdvancedParams field if non-nil, zero value otherwise.

### GetAdvancedParamsOk

`func (o *BTBGltfExportParams) GetAdvancedParamsOk() (*BTBExportAdvancedParams, bool)`

GetAdvancedParamsOk returns a tuple with the AdvancedParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedParams

`func (o *BTBGltfExportParams) SetAdvancedParams(v BTBExportAdvancedParams)`

SetAdvancedParams sets AdvancedParams field to given value.

### HasAdvancedParams

`func (o *BTBGltfExportParams) HasAdvancedParams() bool`

HasAdvancedParams returns a boolean if a field has been set.

### GetCloudStorageOptions

`func (o *BTBGltfExportParams) GetCloudStorageOptions() BTBCloudStorageOptions`

GetCloudStorageOptions returns the CloudStorageOptions field if non-nil, zero value otherwise.

### GetCloudStorageOptionsOk

`func (o *BTBGltfExportParams) GetCloudStorageOptionsOk() (*BTBCloudStorageOptions, bool)`

GetCloudStorageOptionsOk returns a tuple with the CloudStorageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageOptions

`func (o *BTBGltfExportParams) SetCloudStorageOptions(v BTBCloudStorageOptions)`

SetCloudStorageOptions sets CloudStorageOptions field to given value.

### HasCloudStorageOptions

`func (o *BTBGltfExportParams) HasCloudStorageOptions() bool`

HasCloudStorageOptions returns a boolean if a field has been set.

### GetDestinationName

`func (o *BTBGltfExportParams) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *BTBGltfExportParams) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *BTBGltfExportParams) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *BTBGltfExportParams) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.

### GetEmailExportOptions

`func (o *BTBGltfExportParams) GetEmailExportOptions() BTBEmailExportOptions`

GetEmailExportOptions returns the EmailExportOptions field if non-nil, zero value otherwise.

### GetEmailExportOptionsOk

`func (o *BTBGltfExportParams) GetEmailExportOptionsOk() (*BTBEmailExportOptions, bool)`

GetEmailExportOptionsOk returns a tuple with the EmailExportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailExportOptions

`func (o *BTBGltfExportParams) SetEmailExportOptions(v BTBEmailExportOptions)`

SetEmailExportOptions sets EmailExportOptions field to given value.

### HasEmailExportOptions

`func (o *BTBGltfExportParams) HasEmailExportOptions() bool`

HasEmailExportOptions returns a boolean if a field has been set.

### GetExcludeHiddenEntities

`func (o *BTBGltfExportParams) GetExcludeHiddenEntities() bool`

GetExcludeHiddenEntities returns the ExcludeHiddenEntities field if non-nil, zero value otherwise.

### GetExcludeHiddenEntitiesOk

`func (o *BTBGltfExportParams) GetExcludeHiddenEntitiesOk() (*bool, bool)`

GetExcludeHiddenEntitiesOk returns a tuple with the ExcludeHiddenEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHiddenEntities

`func (o *BTBGltfExportParams) SetExcludeHiddenEntities(v bool)`

SetExcludeHiddenEntities sets ExcludeHiddenEntities field to given value.

### HasExcludeHiddenEntities

`func (o *BTBGltfExportParams) HasExcludeHiddenEntities() bool`

HasExcludeHiddenEntities returns a boolean if a field has been set.

### GetGrouping

`func (o *BTBGltfExportParams) GetGrouping() bool`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *BTBGltfExportParams) GetGroupingOk() (*bool, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *BTBGltfExportParams) SetGrouping(v bool)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *BTBGltfExportParams) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetIncludeExportIds

`func (o *BTBGltfExportParams) GetIncludeExportIds() bool`

GetIncludeExportIds returns the IncludeExportIds field if non-nil, zero value otherwise.

### GetIncludeExportIdsOk

`func (o *BTBGltfExportParams) GetIncludeExportIdsOk() (*bool, bool)`

GetIncludeExportIdsOk returns a tuple with the IncludeExportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExportIds

`func (o *BTBGltfExportParams) SetIncludeExportIds(v bool)`

SetIncludeExportIds sets IncludeExportIds field to given value.

### HasIncludeExportIds

`func (o *BTBGltfExportParams) HasIncludeExportIds() bool`

HasIncludeExportIds returns a boolean if a field has been set.

### GetIsYAxisUp

`func (o *BTBGltfExportParams) GetIsYAxisUp() bool`

GetIsYAxisUp returns the IsYAxisUp field if non-nil, zero value otherwise.

### GetIsYAxisUpOk

`func (o *BTBGltfExportParams) GetIsYAxisUpOk() (*bool, bool)`

GetIsYAxisUpOk returns a tuple with the IsYAxisUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYAxisUp

`func (o *BTBGltfExportParams) SetIsYAxisUp(v bool)`

SetIsYAxisUp sets IsYAxisUp field to given value.

### HasIsYAxisUp

`func (o *BTBGltfExportParams) HasIsYAxisUp() bool`

HasIsYAxisUp returns a boolean if a field has been set.

### GetMeshParams

`func (o *BTBGltfExportParams) GetMeshParams() BTBExportMeshParams`

GetMeshParams returns the MeshParams field if non-nil, zero value otherwise.

### GetMeshParamsOk

`func (o *BTBGltfExportParams) GetMeshParamsOk() (*BTBExportMeshParams, bool)`

GetMeshParamsOk returns a tuple with the MeshParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshParams

`func (o *BTBGltfExportParams) SetMeshParams(v BTBExportMeshParams)`

SetMeshParams sets MeshParams field to given value.

### HasMeshParams

`func (o *BTBGltfExportParams) HasMeshParams() bool`

HasMeshParams returns a boolean if a field has been set.

### GetNotifyUser

`func (o *BTBGltfExportParams) GetNotifyUser() bool`

GetNotifyUser returns the NotifyUser field if non-nil, zero value otherwise.

### GetNotifyUserOk

`func (o *BTBGltfExportParams) GetNotifyUserOk() (*bool, bool)`

GetNotifyUserOk returns a tuple with the NotifyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUser

`func (o *BTBGltfExportParams) SetNotifyUser(v bool)`

SetNotifyUser sets NotifyUser field to given value.

### HasNotifyUser

`func (o *BTBGltfExportParams) HasNotifyUser() bool`

HasNotifyUser returns a boolean if a field has been set.

### GetStoreInDocument

`func (o *BTBGltfExportParams) GetStoreInDocument() bool`

GetStoreInDocument returns the StoreInDocument field if non-nil, zero value otherwise.

### GetStoreInDocumentOk

`func (o *BTBGltfExportParams) GetStoreInDocumentOk() (*bool, bool)`

GetStoreInDocumentOk returns a tuple with the StoreInDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInDocument

`func (o *BTBGltfExportParams) SetStoreInDocument(v bool)`

SetStoreInDocument sets StoreInDocument field to given value.

### HasStoreInDocument

`func (o *BTBGltfExportParams) HasStoreInDocument() bool`

HasStoreInDocument returns a boolean if a field has been set.

### GetTriggerAutoDownload

`func (o *BTBGltfExportParams) GetTriggerAutoDownload() bool`

GetTriggerAutoDownload returns the TriggerAutoDownload field if non-nil, zero value otherwise.

### GetTriggerAutoDownloadOk

`func (o *BTBGltfExportParams) GetTriggerAutoDownloadOk() (*bool, bool)`

GetTriggerAutoDownloadOk returns a tuple with the TriggerAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerAutoDownload

`func (o *BTBGltfExportParams) SetTriggerAutoDownload(v bool)`

SetTriggerAutoDownload sets TriggerAutoDownload field to given value.

### HasTriggerAutoDownload

`func (o *BTBGltfExportParams) HasTriggerAutoDownload() bool`

HasTriggerAutoDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


