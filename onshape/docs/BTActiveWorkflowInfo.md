# BTActiveWorkflowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCreateReleases** | Pointer to **bool** |  | [optional] 
**CanCurrentUserCreateReleases** | Pointer to **bool** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**DrawingCanDuplicatePartNumber** | Pointer to **bool** |  | [optional] 
**EnabledActiveMultipleWorkflows** | Pointer to **bool** |  | [optional] 
**ObsoletionWorkflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ObsoletionWorkflowId** | Pointer to **string** |  | [optional] 
**PickableWorkflows** | Pointer to [**[]BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ReleaseWorkflow** | Pointer to [**BTPublishedWorkflowInfo**](BTPublishedWorkflowInfo.md) |  | [optional] 
**ReleaseWorkflowId** | Pointer to **string** |  | [optional] 
**UsingAutoPartNumbering** | Pointer to **bool** |  | [optional] 
**UsingAutoPartNumberingScheme** | Pointer to **bool** |  | [optional] 
**UsingManagedWorkflow** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTActiveWorkflowInfo

`func NewBTActiveWorkflowInfo() *BTActiveWorkflowInfo`

NewBTActiveWorkflowInfo instantiates a new BTActiveWorkflowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTActiveWorkflowInfoWithDefaults

`func NewBTActiveWorkflowInfoWithDefaults() *BTActiveWorkflowInfo`

NewBTActiveWorkflowInfoWithDefaults instantiates a new BTActiveWorkflowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCreateReleases

`func (o *BTActiveWorkflowInfo) GetCanCreateReleases() bool`

GetCanCreateReleases returns the CanCreateReleases field if non-nil, zero value otherwise.

### GetCanCreateReleasesOk

`func (o *BTActiveWorkflowInfo) GetCanCreateReleasesOk() (*bool, bool)`

GetCanCreateReleasesOk returns a tuple with the CanCreateReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCreateReleases

`func (o *BTActiveWorkflowInfo) SetCanCreateReleases(v bool)`

SetCanCreateReleases sets CanCreateReleases field to given value.

### HasCanCreateReleases

`func (o *BTActiveWorkflowInfo) HasCanCreateReleases() bool`

HasCanCreateReleases returns a boolean if a field has been set.

### GetCanCurrentUserCreateReleases

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleases() bool`

GetCanCurrentUserCreateReleases returns the CanCurrentUserCreateReleases field if non-nil, zero value otherwise.

### GetCanCurrentUserCreateReleasesOk

`func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleasesOk() (*bool, bool)`

GetCanCurrentUserCreateReleasesOk returns a tuple with the CanCurrentUserCreateReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCurrentUserCreateReleases

`func (o *BTActiveWorkflowInfo) SetCanCurrentUserCreateReleases(v bool)`

SetCanCurrentUserCreateReleases sets CanCurrentUserCreateReleases field to given value.

### HasCanCurrentUserCreateReleases

`func (o *BTActiveWorkflowInfo) HasCanCurrentUserCreateReleases() bool`

HasCanCurrentUserCreateReleases returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTActiveWorkflowInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTActiveWorkflowInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTActiveWorkflowInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTActiveWorkflowInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTActiveWorkflowInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTActiveWorkflowInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTActiveWorkflowInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTActiveWorkflowInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDrawingCanDuplicatePartNumber

`func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumber() bool`

GetDrawingCanDuplicatePartNumber returns the DrawingCanDuplicatePartNumber field if non-nil, zero value otherwise.

### GetDrawingCanDuplicatePartNumberOk

`func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumberOk() (*bool, bool)`

GetDrawingCanDuplicatePartNumberOk returns a tuple with the DrawingCanDuplicatePartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawingCanDuplicatePartNumber

`func (o *BTActiveWorkflowInfo) SetDrawingCanDuplicatePartNumber(v bool)`

SetDrawingCanDuplicatePartNumber sets DrawingCanDuplicatePartNumber field to given value.

### HasDrawingCanDuplicatePartNumber

`func (o *BTActiveWorkflowInfo) HasDrawingCanDuplicatePartNumber() bool`

HasDrawingCanDuplicatePartNumber returns a boolean if a field has been set.

### GetEnabledActiveMultipleWorkflows

`func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflows() bool`

GetEnabledActiveMultipleWorkflows returns the EnabledActiveMultipleWorkflows field if non-nil, zero value otherwise.

### GetEnabledActiveMultipleWorkflowsOk

`func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflowsOk() (*bool, bool)`

GetEnabledActiveMultipleWorkflowsOk returns a tuple with the EnabledActiveMultipleWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledActiveMultipleWorkflows

`func (o *BTActiveWorkflowInfo) SetEnabledActiveMultipleWorkflows(v bool)`

SetEnabledActiveMultipleWorkflows sets EnabledActiveMultipleWorkflows field to given value.

### HasEnabledActiveMultipleWorkflows

`func (o *BTActiveWorkflowInfo) HasEnabledActiveMultipleWorkflows() bool`

HasEnabledActiveMultipleWorkflows returns a boolean if a field has been set.

### GetObsoletionWorkflow

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflow() BTPublishedWorkflowInfo`

GetObsoletionWorkflow returns the ObsoletionWorkflow field if non-nil, zero value otherwise.

### GetObsoletionWorkflowOk

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowOk() (*BTPublishedWorkflowInfo, bool)`

GetObsoletionWorkflowOk returns a tuple with the ObsoletionWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionWorkflow

`func (o *BTActiveWorkflowInfo) SetObsoletionWorkflow(v BTPublishedWorkflowInfo)`

SetObsoletionWorkflow sets ObsoletionWorkflow field to given value.

### HasObsoletionWorkflow

`func (o *BTActiveWorkflowInfo) HasObsoletionWorkflow() bool`

HasObsoletionWorkflow returns a boolean if a field has been set.

### GetObsoletionWorkflowId

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowId() string`

GetObsoletionWorkflowId returns the ObsoletionWorkflowId field if non-nil, zero value otherwise.

### GetObsoletionWorkflowIdOk

`func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowIdOk() (*string, bool)`

GetObsoletionWorkflowIdOk returns a tuple with the ObsoletionWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoletionWorkflowId

`func (o *BTActiveWorkflowInfo) SetObsoletionWorkflowId(v string)`

SetObsoletionWorkflowId sets ObsoletionWorkflowId field to given value.

### HasObsoletionWorkflowId

`func (o *BTActiveWorkflowInfo) HasObsoletionWorkflowId() bool`

HasObsoletionWorkflowId returns a boolean if a field has been set.

### GetPickableWorkflows

`func (o *BTActiveWorkflowInfo) GetPickableWorkflows() []BTPublishedWorkflowInfo`

GetPickableWorkflows returns the PickableWorkflows field if non-nil, zero value otherwise.

### GetPickableWorkflowsOk

`func (o *BTActiveWorkflowInfo) GetPickableWorkflowsOk() (*[]BTPublishedWorkflowInfo, bool)`

GetPickableWorkflowsOk returns a tuple with the PickableWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickableWorkflows

`func (o *BTActiveWorkflowInfo) SetPickableWorkflows(v []BTPublishedWorkflowInfo)`

SetPickableWorkflows sets PickableWorkflows field to given value.

### HasPickableWorkflows

`func (o *BTActiveWorkflowInfo) HasPickableWorkflows() bool`

HasPickableWorkflows returns a boolean if a field has been set.

### GetReleaseWorkflow

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflow() BTPublishedWorkflowInfo`

GetReleaseWorkflow returns the ReleaseWorkflow field if non-nil, zero value otherwise.

### GetReleaseWorkflowOk

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowOk() (*BTPublishedWorkflowInfo, bool)`

GetReleaseWorkflowOk returns a tuple with the ReleaseWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseWorkflow

`func (o *BTActiveWorkflowInfo) SetReleaseWorkflow(v BTPublishedWorkflowInfo)`

SetReleaseWorkflow sets ReleaseWorkflow field to given value.

### HasReleaseWorkflow

`func (o *BTActiveWorkflowInfo) HasReleaseWorkflow() bool`

HasReleaseWorkflow returns a boolean if a field has been set.

### GetReleaseWorkflowId

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowId() string`

GetReleaseWorkflowId returns the ReleaseWorkflowId field if non-nil, zero value otherwise.

### GetReleaseWorkflowIdOk

`func (o *BTActiveWorkflowInfo) GetReleaseWorkflowIdOk() (*string, bool)`

GetReleaseWorkflowIdOk returns a tuple with the ReleaseWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseWorkflowId

`func (o *BTActiveWorkflowInfo) SetReleaseWorkflowId(v string)`

SetReleaseWorkflowId sets ReleaseWorkflowId field to given value.

### HasReleaseWorkflowId

`func (o *BTActiveWorkflowInfo) HasReleaseWorkflowId() bool`

HasReleaseWorkflowId returns a boolean if a field has been set.

### GetUsingAutoPartNumbering

`func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumbering() bool`

GetUsingAutoPartNumbering returns the UsingAutoPartNumbering field if non-nil, zero value otherwise.

### GetUsingAutoPartNumberingOk

`func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingOk() (*bool, bool)`

GetUsingAutoPartNumberingOk returns a tuple with the UsingAutoPartNumbering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingAutoPartNumbering

`func (o *BTActiveWorkflowInfo) SetUsingAutoPartNumbering(v bool)`

SetUsingAutoPartNumbering sets UsingAutoPartNumbering field to given value.

### HasUsingAutoPartNumbering

`func (o *BTActiveWorkflowInfo) HasUsingAutoPartNumbering() bool`

HasUsingAutoPartNumbering returns a boolean if a field has been set.

### GetUsingAutoPartNumberingScheme

`func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingScheme() bool`

GetUsingAutoPartNumberingScheme returns the UsingAutoPartNumberingScheme field if non-nil, zero value otherwise.

### GetUsingAutoPartNumberingSchemeOk

`func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingSchemeOk() (*bool, bool)`

GetUsingAutoPartNumberingSchemeOk returns a tuple with the UsingAutoPartNumberingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingAutoPartNumberingScheme

`func (o *BTActiveWorkflowInfo) SetUsingAutoPartNumberingScheme(v bool)`

SetUsingAutoPartNumberingScheme sets UsingAutoPartNumberingScheme field to given value.

### HasUsingAutoPartNumberingScheme

`func (o *BTActiveWorkflowInfo) HasUsingAutoPartNumberingScheme() bool`

HasUsingAutoPartNumberingScheme returns a boolean if a field has been set.

### GetUsingManagedWorkflow

`func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflow() bool`

GetUsingManagedWorkflow returns the UsingManagedWorkflow field if non-nil, zero value otherwise.

### GetUsingManagedWorkflowOk

`func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflowOk() (*bool, bool)`

GetUsingManagedWorkflowOk returns a tuple with the UsingManagedWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingManagedWorkflow

`func (o *BTActiveWorkflowInfo) SetUsingManagedWorkflow(v bool)`

SetUsingManagedWorkflow sets UsingManagedWorkflow field to given value.

### HasUsingManagedWorkflow

`func (o *BTActiveWorkflowInfo) HasUsingManagedWorkflow() bool`

HasUsingManagedWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


