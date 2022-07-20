# BTWorkspacePartParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appearance** | Pointer to [**BTPartAppearanceParams**](BTPartAppearanceParams.md) |  | [optional] 
**ApplyUpdateToAllConfigurations** | Pointer to **bool** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**CustomProperties** | Pointer to [**[]BTNameValuePair**](BTNameValuePair.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Material** | Pointer to [**BTMaterialParams**](BTMaterialParams.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to [**BTPSOIdentity2741**](BTPSOIdentity2741.md) |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**ProductLine** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Title1** | Pointer to **string** |  | [optional] 
**Title2** | Pointer to **string** |  | [optional] 
**Title3** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWorkspacePartParams

`func NewBTWorkspacePartParams() *BTWorkspacePartParams`

NewBTWorkspacePartParams instantiates a new BTWorkspacePartParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkspacePartParamsWithDefaults

`func NewBTWorkspacePartParamsWithDefaults() *BTWorkspacePartParams`

NewBTWorkspacePartParamsWithDefaults instantiates a new BTWorkspacePartParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearance

`func (o *BTWorkspacePartParams) GetAppearance() BTPartAppearanceParams`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *BTWorkspacePartParams) GetAppearanceOk() (*BTPartAppearanceParams, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *BTWorkspacePartParams) SetAppearance(v BTPartAppearanceParams)`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *BTWorkspacePartParams) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### GetApplyUpdateToAllConfigurations

`func (o *BTWorkspacePartParams) GetApplyUpdateToAllConfigurations() bool`

GetApplyUpdateToAllConfigurations returns the ApplyUpdateToAllConfigurations field if non-nil, zero value otherwise.

### GetApplyUpdateToAllConfigurationsOk

`func (o *BTWorkspacePartParams) GetApplyUpdateToAllConfigurationsOk() (*bool, bool)`

GetApplyUpdateToAllConfigurationsOk returns a tuple with the ApplyUpdateToAllConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyUpdateToAllConfigurations

`func (o *BTWorkspacePartParams) SetApplyUpdateToAllConfigurations(v bool)`

SetApplyUpdateToAllConfigurations sets ApplyUpdateToAllConfigurations field to given value.

### HasApplyUpdateToAllConfigurations

`func (o *BTWorkspacePartParams) HasApplyUpdateToAllConfigurations() bool`

HasApplyUpdateToAllConfigurations returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTWorkspacePartParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTWorkspacePartParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTWorkspacePartParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTWorkspacePartParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetConnectionId

`func (o *BTWorkspacePartParams) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BTWorkspacePartParams) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BTWorkspacePartParams) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *BTWorkspacePartParams) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCustomProperties

`func (o *BTWorkspacePartParams) GetCustomProperties() []BTNameValuePair`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BTWorkspacePartParams) GetCustomPropertiesOk() (*[]BTNameValuePair, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BTWorkspacePartParams) SetCustomProperties(v []BTNameValuePair)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BTWorkspacePartParams) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *BTWorkspacePartParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTWorkspacePartParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTWorkspacePartParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTWorkspacePartParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *BTWorkspacePartParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTWorkspacePartParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTWorkspacePartParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTWorkspacePartParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetMaterial

`func (o *BTWorkspacePartParams) GetMaterial() BTMaterialParams`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *BTWorkspacePartParams) GetMaterialOk() (*BTMaterialParams, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *BTWorkspacePartParams) SetMaterial(v BTMaterialParams)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *BTWorkspacePartParams) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetName

`func (o *BTWorkspacePartParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTWorkspacePartParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTWorkspacePartParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTWorkspacePartParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartId

`func (o *BTWorkspacePartParams) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTWorkspacePartParams) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTWorkspacePartParams) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTWorkspacePartParams) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTWorkspacePartParams) GetPartIdentity() BTPSOIdentity2741`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTWorkspacePartParams) GetPartIdentityOk() (*BTPSOIdentity2741, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTWorkspacePartParams) SetPartIdentity(v BTPSOIdentity2741)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTWorkspacePartParams) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTWorkspacePartParams) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTWorkspacePartParams) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTWorkspacePartParams) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTWorkspacePartParams) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProductLine

`func (o *BTWorkspacePartParams) GetProductLine() string`

GetProductLine returns the ProductLine field if non-nil, zero value otherwise.

### GetProductLineOk

`func (o *BTWorkspacePartParams) GetProductLineOk() (*string, bool)`

GetProductLineOk returns a tuple with the ProductLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLine

`func (o *BTWorkspacePartParams) SetProductLine(v string)`

SetProductLine sets ProductLine field to given value.

### HasProductLine

`func (o *BTWorkspacePartParams) HasProductLine() bool`

HasProductLine returns a boolean if a field has been set.

### GetProject

`func (o *BTWorkspacePartParams) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BTWorkspacePartParams) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BTWorkspacePartParams) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BTWorkspacePartParams) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevision

`func (o *BTWorkspacePartParams) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTWorkspacePartParams) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTWorkspacePartParams) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTWorkspacePartParams) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTitle1

`func (o *BTWorkspacePartParams) GetTitle1() string`

GetTitle1 returns the Title1 field if non-nil, zero value otherwise.

### GetTitle1Ok

`func (o *BTWorkspacePartParams) GetTitle1Ok() (*string, bool)`

GetTitle1Ok returns a tuple with the Title1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle1

`func (o *BTWorkspacePartParams) SetTitle1(v string)`

SetTitle1 sets Title1 field to given value.

### HasTitle1

`func (o *BTWorkspacePartParams) HasTitle1() bool`

HasTitle1 returns a boolean if a field has been set.

### GetTitle2

`func (o *BTWorkspacePartParams) GetTitle2() string`

GetTitle2 returns the Title2 field if non-nil, zero value otherwise.

### GetTitle2Ok

`func (o *BTWorkspacePartParams) GetTitle2Ok() (*string, bool)`

GetTitle2Ok returns a tuple with the Title2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle2

`func (o *BTWorkspacePartParams) SetTitle2(v string)`

SetTitle2 sets Title2 field to given value.

### HasTitle2

`func (o *BTWorkspacePartParams) HasTitle2() bool`

HasTitle2 returns a boolean if a field has been set.

### GetTitle3

`func (o *BTWorkspacePartParams) GetTitle3() string`

GetTitle3 returns the Title3 field if non-nil, zero value otherwise.

### GetTitle3Ok

`func (o *BTWorkspacePartParams) GetTitle3Ok() (*string, bool)`

GetTitle3Ok returns a tuple with the Title3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle3

`func (o *BTWorkspacePartParams) SetTitle3(v string)`

SetTitle3 sets Title3 field to given value.

### HasTitle3

`func (o *BTWorkspacePartParams) HasTitle3() bool`

HasTitle3 returns a boolean if a field has been set.

### GetVendor

`func (o *BTWorkspacePartParams) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BTWorkspacePartParams) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BTWorkspacePartParams) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *BTWorkspacePartParams) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


