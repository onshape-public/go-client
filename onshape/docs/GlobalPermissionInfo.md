# GlobalPermissionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessReports** | Pointer to **bool** |  | [optional] 
**AdminEnterprise** | Pointer to **bool** |  | [optional] 
**AllowAppStoreAccess** | Pointer to **bool** |  | [optional] 
**AllowPublicDocumentsAccess** | Pointer to **bool** |  | [optional] 
**ApproveReleases** | Pointer to **bool** |  | [optional] 
**CreateChangeOrders** | Pointer to **bool** |  | [optional] 
**CreateChangeRequests** | Pointer to **bool** |  | [optional] 
**CreateDocumentsInRoot** | Pointer to **bool** |  | [optional] 
**CreateProject** | Pointer to **bool** |  | [optional] 
**CreateReleases** | Pointer to **bool** |  | [optional] 
**CreateTasks** | Pointer to **bool** |  | [optional] 
**DeletePermanently** | Pointer to **bool** |  | [optional] 
**ManageGuestUsers** | Pointer to **bool** |  | [optional] 
**ManageNonGeometricItems** | Pointer to **bool** |  | [optional] 
**ManageRbac** | Pointer to **bool** |  | [optional] 
**ManageStandardContentMetadata** | Pointer to **bool** |  | [optional] 
**ManageUsers** | Pointer to **bool** |  | [optional] 
**ManageWorkflows** | Pointer to **bool** |  | [optional] 
**ShareForAnonymousAccess** | Pointer to **bool** |  | [optional] 
**TransferDocumentsFromEnterprise** | Pointer to **bool** |  | [optional] 
**ViewChangeOrders** | Pointer to **bool** |  | [optional] 
**ViewChangeRequests** | Pointer to **bool** |  | [optional] 

## Methods

### NewGlobalPermissionInfo

`func NewGlobalPermissionInfo() *GlobalPermissionInfo`

NewGlobalPermissionInfo instantiates a new GlobalPermissionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalPermissionInfoWithDefaults

`func NewGlobalPermissionInfoWithDefaults() *GlobalPermissionInfo`

NewGlobalPermissionInfoWithDefaults instantiates a new GlobalPermissionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessReports

`func (o *GlobalPermissionInfo) GetAccessReports() bool`

GetAccessReports returns the AccessReports field if non-nil, zero value otherwise.

### GetAccessReportsOk

`func (o *GlobalPermissionInfo) GetAccessReportsOk() (*bool, bool)`

GetAccessReportsOk returns a tuple with the AccessReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReports

`func (o *GlobalPermissionInfo) SetAccessReports(v bool)`

SetAccessReports sets AccessReports field to given value.

### HasAccessReports

`func (o *GlobalPermissionInfo) HasAccessReports() bool`

HasAccessReports returns a boolean if a field has been set.

### GetAdminEnterprise

`func (o *GlobalPermissionInfo) GetAdminEnterprise() bool`

GetAdminEnterprise returns the AdminEnterprise field if non-nil, zero value otherwise.

### GetAdminEnterpriseOk

`func (o *GlobalPermissionInfo) GetAdminEnterpriseOk() (*bool, bool)`

GetAdminEnterpriseOk returns a tuple with the AdminEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEnterprise

`func (o *GlobalPermissionInfo) SetAdminEnterprise(v bool)`

SetAdminEnterprise sets AdminEnterprise field to given value.

### HasAdminEnterprise

`func (o *GlobalPermissionInfo) HasAdminEnterprise() bool`

HasAdminEnterprise returns a boolean if a field has been set.

### GetAllowAppStoreAccess

`func (o *GlobalPermissionInfo) GetAllowAppStoreAccess() bool`

GetAllowAppStoreAccess returns the AllowAppStoreAccess field if non-nil, zero value otherwise.

### GetAllowAppStoreAccessOk

`func (o *GlobalPermissionInfo) GetAllowAppStoreAccessOk() (*bool, bool)`

GetAllowAppStoreAccessOk returns a tuple with the AllowAppStoreAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAppStoreAccess

`func (o *GlobalPermissionInfo) SetAllowAppStoreAccess(v bool)`

SetAllowAppStoreAccess sets AllowAppStoreAccess field to given value.

### HasAllowAppStoreAccess

`func (o *GlobalPermissionInfo) HasAllowAppStoreAccess() bool`

HasAllowAppStoreAccess returns a boolean if a field has been set.

### GetAllowPublicDocumentsAccess

`func (o *GlobalPermissionInfo) GetAllowPublicDocumentsAccess() bool`

GetAllowPublicDocumentsAccess returns the AllowPublicDocumentsAccess field if non-nil, zero value otherwise.

### GetAllowPublicDocumentsAccessOk

`func (o *GlobalPermissionInfo) GetAllowPublicDocumentsAccessOk() (*bool, bool)`

GetAllowPublicDocumentsAccessOk returns a tuple with the AllowPublicDocumentsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPublicDocumentsAccess

`func (o *GlobalPermissionInfo) SetAllowPublicDocumentsAccess(v bool)`

SetAllowPublicDocumentsAccess sets AllowPublicDocumentsAccess field to given value.

### HasAllowPublicDocumentsAccess

`func (o *GlobalPermissionInfo) HasAllowPublicDocumentsAccess() bool`

HasAllowPublicDocumentsAccess returns a boolean if a field has been set.

### GetApproveReleases

`func (o *GlobalPermissionInfo) GetApproveReleases() bool`

GetApproveReleases returns the ApproveReleases field if non-nil, zero value otherwise.

### GetApproveReleasesOk

`func (o *GlobalPermissionInfo) GetApproveReleasesOk() (*bool, bool)`

GetApproveReleasesOk returns a tuple with the ApproveReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveReleases

`func (o *GlobalPermissionInfo) SetApproveReleases(v bool)`

SetApproveReleases sets ApproveReleases field to given value.

### HasApproveReleases

`func (o *GlobalPermissionInfo) HasApproveReleases() bool`

HasApproveReleases returns a boolean if a field has been set.

### GetCreateChangeOrders

`func (o *GlobalPermissionInfo) GetCreateChangeOrders() bool`

GetCreateChangeOrders returns the CreateChangeOrders field if non-nil, zero value otherwise.

### GetCreateChangeOrdersOk

`func (o *GlobalPermissionInfo) GetCreateChangeOrdersOk() (*bool, bool)`

GetCreateChangeOrdersOk returns a tuple with the CreateChangeOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateChangeOrders

`func (o *GlobalPermissionInfo) SetCreateChangeOrders(v bool)`

SetCreateChangeOrders sets CreateChangeOrders field to given value.

### HasCreateChangeOrders

`func (o *GlobalPermissionInfo) HasCreateChangeOrders() bool`

HasCreateChangeOrders returns a boolean if a field has been set.

### GetCreateChangeRequests

`func (o *GlobalPermissionInfo) GetCreateChangeRequests() bool`

GetCreateChangeRequests returns the CreateChangeRequests field if non-nil, zero value otherwise.

### GetCreateChangeRequestsOk

`func (o *GlobalPermissionInfo) GetCreateChangeRequestsOk() (*bool, bool)`

GetCreateChangeRequestsOk returns a tuple with the CreateChangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateChangeRequests

`func (o *GlobalPermissionInfo) SetCreateChangeRequests(v bool)`

SetCreateChangeRequests sets CreateChangeRequests field to given value.

### HasCreateChangeRequests

`func (o *GlobalPermissionInfo) HasCreateChangeRequests() bool`

HasCreateChangeRequests returns a boolean if a field has been set.

### GetCreateDocumentsInRoot

`func (o *GlobalPermissionInfo) GetCreateDocumentsInRoot() bool`

GetCreateDocumentsInRoot returns the CreateDocumentsInRoot field if non-nil, zero value otherwise.

### GetCreateDocumentsInRootOk

`func (o *GlobalPermissionInfo) GetCreateDocumentsInRootOk() (*bool, bool)`

GetCreateDocumentsInRootOk returns a tuple with the CreateDocumentsInRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDocumentsInRoot

`func (o *GlobalPermissionInfo) SetCreateDocumentsInRoot(v bool)`

SetCreateDocumentsInRoot sets CreateDocumentsInRoot field to given value.

### HasCreateDocumentsInRoot

`func (o *GlobalPermissionInfo) HasCreateDocumentsInRoot() bool`

HasCreateDocumentsInRoot returns a boolean if a field has been set.

### GetCreateProject

`func (o *GlobalPermissionInfo) GetCreateProject() bool`

GetCreateProject returns the CreateProject field if non-nil, zero value otherwise.

### GetCreateProjectOk

`func (o *GlobalPermissionInfo) GetCreateProjectOk() (*bool, bool)`

GetCreateProjectOk returns a tuple with the CreateProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProject

`func (o *GlobalPermissionInfo) SetCreateProject(v bool)`

SetCreateProject sets CreateProject field to given value.

### HasCreateProject

`func (o *GlobalPermissionInfo) HasCreateProject() bool`

HasCreateProject returns a boolean if a field has been set.

### GetCreateReleases

`func (o *GlobalPermissionInfo) GetCreateReleases() bool`

GetCreateReleases returns the CreateReleases field if non-nil, zero value otherwise.

### GetCreateReleasesOk

`func (o *GlobalPermissionInfo) GetCreateReleasesOk() (*bool, bool)`

GetCreateReleasesOk returns a tuple with the CreateReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateReleases

`func (o *GlobalPermissionInfo) SetCreateReleases(v bool)`

SetCreateReleases sets CreateReleases field to given value.

### HasCreateReleases

`func (o *GlobalPermissionInfo) HasCreateReleases() bool`

HasCreateReleases returns a boolean if a field has been set.

### GetCreateTasks

`func (o *GlobalPermissionInfo) GetCreateTasks() bool`

GetCreateTasks returns the CreateTasks field if non-nil, zero value otherwise.

### GetCreateTasksOk

`func (o *GlobalPermissionInfo) GetCreateTasksOk() (*bool, bool)`

GetCreateTasksOk returns a tuple with the CreateTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTasks

`func (o *GlobalPermissionInfo) SetCreateTasks(v bool)`

SetCreateTasks sets CreateTasks field to given value.

### HasCreateTasks

`func (o *GlobalPermissionInfo) HasCreateTasks() bool`

HasCreateTasks returns a boolean if a field has been set.

### GetDeletePermanently

`func (o *GlobalPermissionInfo) GetDeletePermanently() bool`

GetDeletePermanently returns the DeletePermanently field if non-nil, zero value otherwise.

### GetDeletePermanentlyOk

`func (o *GlobalPermissionInfo) GetDeletePermanentlyOk() (*bool, bool)`

GetDeletePermanentlyOk returns a tuple with the DeletePermanently field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePermanently

`func (o *GlobalPermissionInfo) SetDeletePermanently(v bool)`

SetDeletePermanently sets DeletePermanently field to given value.

### HasDeletePermanently

`func (o *GlobalPermissionInfo) HasDeletePermanently() bool`

HasDeletePermanently returns a boolean if a field has been set.

### GetManageGuestUsers

`func (o *GlobalPermissionInfo) GetManageGuestUsers() bool`

GetManageGuestUsers returns the ManageGuestUsers field if non-nil, zero value otherwise.

### GetManageGuestUsersOk

`func (o *GlobalPermissionInfo) GetManageGuestUsersOk() (*bool, bool)`

GetManageGuestUsersOk returns a tuple with the ManageGuestUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageGuestUsers

`func (o *GlobalPermissionInfo) SetManageGuestUsers(v bool)`

SetManageGuestUsers sets ManageGuestUsers field to given value.

### HasManageGuestUsers

`func (o *GlobalPermissionInfo) HasManageGuestUsers() bool`

HasManageGuestUsers returns a boolean if a field has been set.

### GetManageNonGeometricItems

`func (o *GlobalPermissionInfo) GetManageNonGeometricItems() bool`

GetManageNonGeometricItems returns the ManageNonGeometricItems field if non-nil, zero value otherwise.

### GetManageNonGeometricItemsOk

`func (o *GlobalPermissionInfo) GetManageNonGeometricItemsOk() (*bool, bool)`

GetManageNonGeometricItemsOk returns a tuple with the ManageNonGeometricItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageNonGeometricItems

`func (o *GlobalPermissionInfo) SetManageNonGeometricItems(v bool)`

SetManageNonGeometricItems sets ManageNonGeometricItems field to given value.

### HasManageNonGeometricItems

`func (o *GlobalPermissionInfo) HasManageNonGeometricItems() bool`

HasManageNonGeometricItems returns a boolean if a field has been set.

### GetManageRbac

`func (o *GlobalPermissionInfo) GetManageRbac() bool`

GetManageRbac returns the ManageRbac field if non-nil, zero value otherwise.

### GetManageRbacOk

`func (o *GlobalPermissionInfo) GetManageRbacOk() (*bool, bool)`

GetManageRbacOk returns a tuple with the ManageRbac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageRbac

`func (o *GlobalPermissionInfo) SetManageRbac(v bool)`

SetManageRbac sets ManageRbac field to given value.

### HasManageRbac

`func (o *GlobalPermissionInfo) HasManageRbac() bool`

HasManageRbac returns a boolean if a field has been set.

### GetManageStandardContentMetadata

`func (o *GlobalPermissionInfo) GetManageStandardContentMetadata() bool`

GetManageStandardContentMetadata returns the ManageStandardContentMetadata field if non-nil, zero value otherwise.

### GetManageStandardContentMetadataOk

`func (o *GlobalPermissionInfo) GetManageStandardContentMetadataOk() (*bool, bool)`

GetManageStandardContentMetadataOk returns a tuple with the ManageStandardContentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageStandardContentMetadata

`func (o *GlobalPermissionInfo) SetManageStandardContentMetadata(v bool)`

SetManageStandardContentMetadata sets ManageStandardContentMetadata field to given value.

### HasManageStandardContentMetadata

`func (o *GlobalPermissionInfo) HasManageStandardContentMetadata() bool`

HasManageStandardContentMetadata returns a boolean if a field has been set.

### GetManageUsers

`func (o *GlobalPermissionInfo) GetManageUsers() bool`

GetManageUsers returns the ManageUsers field if non-nil, zero value otherwise.

### GetManageUsersOk

`func (o *GlobalPermissionInfo) GetManageUsersOk() (*bool, bool)`

GetManageUsersOk returns a tuple with the ManageUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageUsers

`func (o *GlobalPermissionInfo) SetManageUsers(v bool)`

SetManageUsers sets ManageUsers field to given value.

### HasManageUsers

`func (o *GlobalPermissionInfo) HasManageUsers() bool`

HasManageUsers returns a boolean if a field has been set.

### GetManageWorkflows

`func (o *GlobalPermissionInfo) GetManageWorkflows() bool`

GetManageWorkflows returns the ManageWorkflows field if non-nil, zero value otherwise.

### GetManageWorkflowsOk

`func (o *GlobalPermissionInfo) GetManageWorkflowsOk() (*bool, bool)`

GetManageWorkflowsOk returns a tuple with the ManageWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageWorkflows

`func (o *GlobalPermissionInfo) SetManageWorkflows(v bool)`

SetManageWorkflows sets ManageWorkflows field to given value.

### HasManageWorkflows

`func (o *GlobalPermissionInfo) HasManageWorkflows() bool`

HasManageWorkflows returns a boolean if a field has been set.

### GetShareForAnonymousAccess

`func (o *GlobalPermissionInfo) GetShareForAnonymousAccess() bool`

GetShareForAnonymousAccess returns the ShareForAnonymousAccess field if non-nil, zero value otherwise.

### GetShareForAnonymousAccessOk

`func (o *GlobalPermissionInfo) GetShareForAnonymousAccessOk() (*bool, bool)`

GetShareForAnonymousAccessOk returns a tuple with the ShareForAnonymousAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareForAnonymousAccess

`func (o *GlobalPermissionInfo) SetShareForAnonymousAccess(v bool)`

SetShareForAnonymousAccess sets ShareForAnonymousAccess field to given value.

### HasShareForAnonymousAccess

`func (o *GlobalPermissionInfo) HasShareForAnonymousAccess() bool`

HasShareForAnonymousAccess returns a boolean if a field has been set.

### GetTransferDocumentsFromEnterprise

`func (o *GlobalPermissionInfo) GetTransferDocumentsFromEnterprise() bool`

GetTransferDocumentsFromEnterprise returns the TransferDocumentsFromEnterprise field if non-nil, zero value otherwise.

### GetTransferDocumentsFromEnterpriseOk

`func (o *GlobalPermissionInfo) GetTransferDocumentsFromEnterpriseOk() (*bool, bool)`

GetTransferDocumentsFromEnterpriseOk returns a tuple with the TransferDocumentsFromEnterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferDocumentsFromEnterprise

`func (o *GlobalPermissionInfo) SetTransferDocumentsFromEnterprise(v bool)`

SetTransferDocumentsFromEnterprise sets TransferDocumentsFromEnterprise field to given value.

### HasTransferDocumentsFromEnterprise

`func (o *GlobalPermissionInfo) HasTransferDocumentsFromEnterprise() bool`

HasTransferDocumentsFromEnterprise returns a boolean if a field has been set.

### GetViewChangeOrders

`func (o *GlobalPermissionInfo) GetViewChangeOrders() bool`

GetViewChangeOrders returns the ViewChangeOrders field if non-nil, zero value otherwise.

### GetViewChangeOrdersOk

`func (o *GlobalPermissionInfo) GetViewChangeOrdersOk() (*bool, bool)`

GetViewChangeOrdersOk returns a tuple with the ViewChangeOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewChangeOrders

`func (o *GlobalPermissionInfo) SetViewChangeOrders(v bool)`

SetViewChangeOrders sets ViewChangeOrders field to given value.

### HasViewChangeOrders

`func (o *GlobalPermissionInfo) HasViewChangeOrders() bool`

HasViewChangeOrders returns a boolean if a field has been set.

### GetViewChangeRequests

`func (o *GlobalPermissionInfo) GetViewChangeRequests() bool`

GetViewChangeRequests returns the ViewChangeRequests field if non-nil, zero value otherwise.

### GetViewChangeRequestsOk

`func (o *GlobalPermissionInfo) GetViewChangeRequestsOk() (*bool, bool)`

GetViewChangeRequestsOk returns a tuple with the ViewChangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewChangeRequests

`func (o *GlobalPermissionInfo) SetViewChangeRequests(v bool)`

SetViewChangeRequests sets ViewChangeRequests field to given value.

### HasViewChangeRequests

`func (o *GlobalPermissionInfo) HasViewChangeRequests() bool`

HasViewChangeRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


