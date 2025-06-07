# BTBExportAdvancedParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyExportParams** | Pointer to [**BTBAssemblyExportParams**](BTBAssemblyExportParams.md) |  | [optional] 
**Configuration** | Pointer to **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [optional] 
**ElementIds** | Pointer to **[]string** | An array of element ids for multi-element export. | [optional] 
**EvaluateExportRule** | Pointer to **bool** | Set to &#x60;true&#x60; to evaluate the export rule for the given &#x60;formatName&#x60; and to include an &#x60;exportRuleFileName&#x60; value in the response. | [optional] [default to false]
**IgnoreExportRulesForContents** | Pointer to **bool** | For multiple elements export, use &#x60;true&#x60; if export rule shouldn&#39;t be applied for all elements. | [optional] [default to false]
**LinkDocumentId** | Pointer to **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [optional] 
**LinkDocumentWorkspaceId** | Pointer to **string** | The id of the workspace through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [optional] 
**PartIds** | Pointer to **string** | IDs of the parts to retrieve. Use comma-separated IDs for multiple parts (example: partIds&#x3D;JHK,JHD). | [optional] 
**PartsExportFilter** | Pointer to [**BTPartsExportFilter4308**](BTPartsExportFilter4308.md) |  | [optional] 

## Methods

### NewBTBExportAdvancedParams

`func NewBTBExportAdvancedParams() *BTBExportAdvancedParams`

NewBTBExportAdvancedParams instantiates a new BTBExportAdvancedParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBExportAdvancedParamsWithDefaults

`func NewBTBExportAdvancedParamsWithDefaults() *BTBExportAdvancedParams`

NewBTBExportAdvancedParamsWithDefaults instantiates a new BTBExportAdvancedParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyExportParams

`func (o *BTBExportAdvancedParams) GetAssemblyExportParams() BTBAssemblyExportParams`

GetAssemblyExportParams returns the AssemblyExportParams field if non-nil, zero value otherwise.

### GetAssemblyExportParamsOk

`func (o *BTBExportAdvancedParams) GetAssemblyExportParamsOk() (*BTBAssemblyExportParams, bool)`

GetAssemblyExportParamsOk returns a tuple with the AssemblyExportParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyExportParams

`func (o *BTBExportAdvancedParams) SetAssemblyExportParams(v BTBAssemblyExportParams)`

SetAssemblyExportParams sets AssemblyExportParams field to given value.

### HasAssemblyExportParams

`func (o *BTBExportAdvancedParams) HasAssemblyExportParams() bool`

HasAssemblyExportParams returns a boolean if a field has been set.

### GetConfiguration

`func (o *BTBExportAdvancedParams) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BTBExportAdvancedParams) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BTBExportAdvancedParams) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *BTBExportAdvancedParams) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetElementIds

`func (o *BTBExportAdvancedParams) GetElementIds() []string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *BTBExportAdvancedParams) GetElementIdsOk() (*[]string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *BTBExportAdvancedParams) SetElementIds(v []string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *BTBExportAdvancedParams) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### GetEvaluateExportRule

`func (o *BTBExportAdvancedParams) GetEvaluateExportRule() bool`

GetEvaluateExportRule returns the EvaluateExportRule field if non-nil, zero value otherwise.

### GetEvaluateExportRuleOk

`func (o *BTBExportAdvancedParams) GetEvaluateExportRuleOk() (*bool, bool)`

GetEvaluateExportRuleOk returns a tuple with the EvaluateExportRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluateExportRule

`func (o *BTBExportAdvancedParams) SetEvaluateExportRule(v bool)`

SetEvaluateExportRule sets EvaluateExportRule field to given value.

### HasEvaluateExportRule

`func (o *BTBExportAdvancedParams) HasEvaluateExportRule() bool`

HasEvaluateExportRule returns a boolean if a field has been set.

### GetIgnoreExportRulesForContents

`func (o *BTBExportAdvancedParams) GetIgnoreExportRulesForContents() bool`

GetIgnoreExportRulesForContents returns the IgnoreExportRulesForContents field if non-nil, zero value otherwise.

### GetIgnoreExportRulesForContentsOk

`func (o *BTBExportAdvancedParams) GetIgnoreExportRulesForContentsOk() (*bool, bool)`

GetIgnoreExportRulesForContentsOk returns a tuple with the IgnoreExportRulesForContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreExportRulesForContents

`func (o *BTBExportAdvancedParams) SetIgnoreExportRulesForContents(v bool)`

SetIgnoreExportRulesForContents sets IgnoreExportRulesForContents field to given value.

### HasIgnoreExportRulesForContents

`func (o *BTBExportAdvancedParams) HasIgnoreExportRulesForContents() bool`

HasIgnoreExportRulesForContents returns a boolean if a field has been set.

### GetLinkDocumentId

`func (o *BTBExportAdvancedParams) GetLinkDocumentId() string`

GetLinkDocumentId returns the LinkDocumentId field if non-nil, zero value otherwise.

### GetLinkDocumentIdOk

`func (o *BTBExportAdvancedParams) GetLinkDocumentIdOk() (*string, bool)`

GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentId

`func (o *BTBExportAdvancedParams) SetLinkDocumentId(v string)`

SetLinkDocumentId sets LinkDocumentId field to given value.

### HasLinkDocumentId

`func (o *BTBExportAdvancedParams) HasLinkDocumentId() bool`

HasLinkDocumentId returns a boolean if a field has been set.

### GetLinkDocumentWorkspaceId

`func (o *BTBExportAdvancedParams) GetLinkDocumentWorkspaceId() string`

GetLinkDocumentWorkspaceId returns the LinkDocumentWorkspaceId field if non-nil, zero value otherwise.

### GetLinkDocumentWorkspaceIdOk

`func (o *BTBExportAdvancedParams) GetLinkDocumentWorkspaceIdOk() (*string, bool)`

GetLinkDocumentWorkspaceIdOk returns a tuple with the LinkDocumentWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkDocumentWorkspaceId

`func (o *BTBExportAdvancedParams) SetLinkDocumentWorkspaceId(v string)`

SetLinkDocumentWorkspaceId sets LinkDocumentWorkspaceId field to given value.

### HasLinkDocumentWorkspaceId

`func (o *BTBExportAdvancedParams) HasLinkDocumentWorkspaceId() bool`

HasLinkDocumentWorkspaceId returns a boolean if a field has been set.

### GetPartIds

`func (o *BTBExportAdvancedParams) GetPartIds() string`

GetPartIds returns the PartIds field if non-nil, zero value otherwise.

### GetPartIdsOk

`func (o *BTBExportAdvancedParams) GetPartIdsOk() (*string, bool)`

GetPartIdsOk returns a tuple with the PartIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIds

`func (o *BTBExportAdvancedParams) SetPartIds(v string)`

SetPartIds sets PartIds field to given value.

### HasPartIds

`func (o *BTBExportAdvancedParams) HasPartIds() bool`

HasPartIds returns a boolean if a field has been set.

### GetPartsExportFilter

`func (o *BTBExportAdvancedParams) GetPartsExportFilter() BTPartsExportFilter4308`

GetPartsExportFilter returns the PartsExportFilter field if non-nil, zero value otherwise.

### GetPartsExportFilterOk

`func (o *BTBExportAdvancedParams) GetPartsExportFilterOk() (*BTPartsExportFilter4308, bool)`

GetPartsExportFilterOk returns a tuple with the PartsExportFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartsExportFilter

`func (o *BTBExportAdvancedParams) SetPartsExportFilter(v BTPartsExportFilter4308)`

SetPartsExportFilter sets PartsExportFilter field to given value.

### HasPartsExportFilter

`func (o *BTBExportAdvancedParams) HasPartsExportFilter() bool`

HasPartsExportFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


