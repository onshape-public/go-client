# BTElementMergeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchPointElementName** | Pointer to **string** |  | [optional] 
**BranchPointElementPath** | Pointer to **[]string** |  | [optional] 
**DependentElementMergeInfo** | Pointer to [**BTElementMergeInfo**](BTElementMergeInfo.md) |  | [optional] 
**ElementDataType** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 
**Mergeable** | Pointer to **bool** |  | [optional] 
**SourceElementName** | Pointer to **string** |  | [optional] 
**SourceElementPath** | Pointer to **[]string** |  | [optional] 
**SourceElementStatus** | Pointer to **string** |  | [optional] 
**SourceModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**SourceModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**TargetElementName** | Pointer to **string** |  | [optional] 
**TargetElementPath** | Pointer to **[]string** |  | [optional] 
**TargetElementStatus** | Pointer to **string** |  | [optional] 
**TargetModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**TargetModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 

## Methods

### NewBTElementMergeInfo

`func NewBTElementMergeInfo() *BTElementMergeInfo`

NewBTElementMergeInfo instantiates a new BTElementMergeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTElementMergeInfoWithDefaults

`func NewBTElementMergeInfoWithDefaults() *BTElementMergeInfo`

NewBTElementMergeInfoWithDefaults instantiates a new BTElementMergeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchPointElementName

`func (o *BTElementMergeInfo) GetBranchPointElementName() string`

GetBranchPointElementName returns the BranchPointElementName field if non-nil, zero value otherwise.

### GetBranchPointElementNameOk

`func (o *BTElementMergeInfo) GetBranchPointElementNameOk() (*string, bool)`

GetBranchPointElementNameOk returns a tuple with the BranchPointElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchPointElementName

`func (o *BTElementMergeInfo) SetBranchPointElementName(v string)`

SetBranchPointElementName sets BranchPointElementName field to given value.

### HasBranchPointElementName

`func (o *BTElementMergeInfo) HasBranchPointElementName() bool`

HasBranchPointElementName returns a boolean if a field has been set.

### GetBranchPointElementPath

`func (o *BTElementMergeInfo) GetBranchPointElementPath() []string`

GetBranchPointElementPath returns the BranchPointElementPath field if non-nil, zero value otherwise.

### GetBranchPointElementPathOk

`func (o *BTElementMergeInfo) GetBranchPointElementPathOk() (*[]string, bool)`

GetBranchPointElementPathOk returns a tuple with the BranchPointElementPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchPointElementPath

`func (o *BTElementMergeInfo) SetBranchPointElementPath(v []string)`

SetBranchPointElementPath sets BranchPointElementPath field to given value.

### HasBranchPointElementPath

`func (o *BTElementMergeInfo) HasBranchPointElementPath() bool`

HasBranchPointElementPath returns a boolean if a field has been set.

### GetDependentElementMergeInfo

`func (o *BTElementMergeInfo) GetDependentElementMergeInfo() BTElementMergeInfo`

GetDependentElementMergeInfo returns the DependentElementMergeInfo field if non-nil, zero value otherwise.

### GetDependentElementMergeInfoOk

`func (o *BTElementMergeInfo) GetDependentElementMergeInfoOk() (*BTElementMergeInfo, bool)`

GetDependentElementMergeInfoOk returns a tuple with the DependentElementMergeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentElementMergeInfo

`func (o *BTElementMergeInfo) SetDependentElementMergeInfo(v BTElementMergeInfo)`

SetDependentElementMergeInfo sets DependentElementMergeInfo field to given value.

### HasDependentElementMergeInfo

`func (o *BTElementMergeInfo) HasDependentElementMergeInfo() bool`

HasDependentElementMergeInfo returns a boolean if a field has been set.

### GetElementDataType

`func (o *BTElementMergeInfo) GetElementDataType() string`

GetElementDataType returns the ElementDataType field if non-nil, zero value otherwise.

### GetElementDataTypeOk

`func (o *BTElementMergeInfo) GetElementDataTypeOk() (*string, bool)`

GetElementDataTypeOk returns a tuple with the ElementDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementDataType

`func (o *BTElementMergeInfo) SetElementDataType(v string)`

SetElementDataType sets ElementDataType field to given value.

### HasElementDataType

`func (o *BTElementMergeInfo) HasElementDataType() bool`

HasElementDataType returns a boolean if a field has been set.

### GetElementId

`func (o *BTElementMergeInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTElementMergeInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTElementMergeInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTElementMergeInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementType

`func (o *BTElementMergeInfo) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTElementMergeInfo) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTElementMergeInfo) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTElementMergeInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetMergeable

`func (o *BTElementMergeInfo) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *BTElementMergeInfo) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *BTElementMergeInfo) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.

### HasMergeable

`func (o *BTElementMergeInfo) HasMergeable() bool`

HasMergeable returns a boolean if a field has been set.

### GetSourceElementName

`func (o *BTElementMergeInfo) GetSourceElementName() string`

GetSourceElementName returns the SourceElementName field if non-nil, zero value otherwise.

### GetSourceElementNameOk

`func (o *BTElementMergeInfo) GetSourceElementNameOk() (*string, bool)`

GetSourceElementNameOk returns a tuple with the SourceElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceElementName

`func (o *BTElementMergeInfo) SetSourceElementName(v string)`

SetSourceElementName sets SourceElementName field to given value.

### HasSourceElementName

`func (o *BTElementMergeInfo) HasSourceElementName() bool`

HasSourceElementName returns a boolean if a field has been set.

### GetSourceElementPath

`func (o *BTElementMergeInfo) GetSourceElementPath() []string`

GetSourceElementPath returns the SourceElementPath field if non-nil, zero value otherwise.

### GetSourceElementPathOk

`func (o *BTElementMergeInfo) GetSourceElementPathOk() (*[]string, bool)`

GetSourceElementPathOk returns a tuple with the SourceElementPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceElementPath

`func (o *BTElementMergeInfo) SetSourceElementPath(v []string)`

SetSourceElementPath sets SourceElementPath field to given value.

### HasSourceElementPath

`func (o *BTElementMergeInfo) HasSourceElementPath() bool`

HasSourceElementPath returns a boolean if a field has been set.

### GetSourceElementStatus

`func (o *BTElementMergeInfo) GetSourceElementStatus() string`

GetSourceElementStatus returns the SourceElementStatus field if non-nil, zero value otherwise.

### GetSourceElementStatusOk

`func (o *BTElementMergeInfo) GetSourceElementStatusOk() (*string, bool)`

GetSourceElementStatusOk returns a tuple with the SourceElementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceElementStatus

`func (o *BTElementMergeInfo) SetSourceElementStatus(v string)`

SetSourceElementStatus sets SourceElementStatus field to given value.

### HasSourceElementStatus

`func (o *BTElementMergeInfo) HasSourceElementStatus() bool`

HasSourceElementStatus returns a boolean if a field has been set.

### GetSourceModifiedAt

`func (o *BTElementMergeInfo) GetSourceModifiedAt() JSONTime`

GetSourceModifiedAt returns the SourceModifiedAt field if non-nil, zero value otherwise.

### GetSourceModifiedAtOk

`func (o *BTElementMergeInfo) GetSourceModifiedAtOk() (*JSONTime, bool)`

GetSourceModifiedAtOk returns a tuple with the SourceModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceModifiedAt

`func (o *BTElementMergeInfo) SetSourceModifiedAt(v JSONTime)`

SetSourceModifiedAt sets SourceModifiedAt field to given value.

### HasSourceModifiedAt

`func (o *BTElementMergeInfo) HasSourceModifiedAt() bool`

HasSourceModifiedAt returns a boolean if a field has been set.

### GetSourceModifiedBy

`func (o *BTElementMergeInfo) GetSourceModifiedBy() BTUserBasicSummaryInfo`

GetSourceModifiedBy returns the SourceModifiedBy field if non-nil, zero value otherwise.

### GetSourceModifiedByOk

`func (o *BTElementMergeInfo) GetSourceModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetSourceModifiedByOk returns a tuple with the SourceModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceModifiedBy

`func (o *BTElementMergeInfo) SetSourceModifiedBy(v BTUserBasicSummaryInfo)`

SetSourceModifiedBy sets SourceModifiedBy field to given value.

### HasSourceModifiedBy

`func (o *BTElementMergeInfo) HasSourceModifiedBy() bool`

HasSourceModifiedBy returns a boolean if a field has been set.

### GetTargetElementName

`func (o *BTElementMergeInfo) GetTargetElementName() string`

GetTargetElementName returns the TargetElementName field if non-nil, zero value otherwise.

### GetTargetElementNameOk

`func (o *BTElementMergeInfo) GetTargetElementNameOk() (*string, bool)`

GetTargetElementNameOk returns a tuple with the TargetElementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetElementName

`func (o *BTElementMergeInfo) SetTargetElementName(v string)`

SetTargetElementName sets TargetElementName field to given value.

### HasTargetElementName

`func (o *BTElementMergeInfo) HasTargetElementName() bool`

HasTargetElementName returns a boolean if a field has been set.

### GetTargetElementPath

`func (o *BTElementMergeInfo) GetTargetElementPath() []string`

GetTargetElementPath returns the TargetElementPath field if non-nil, zero value otherwise.

### GetTargetElementPathOk

`func (o *BTElementMergeInfo) GetTargetElementPathOk() (*[]string, bool)`

GetTargetElementPathOk returns a tuple with the TargetElementPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetElementPath

`func (o *BTElementMergeInfo) SetTargetElementPath(v []string)`

SetTargetElementPath sets TargetElementPath field to given value.

### HasTargetElementPath

`func (o *BTElementMergeInfo) HasTargetElementPath() bool`

HasTargetElementPath returns a boolean if a field has been set.

### GetTargetElementStatus

`func (o *BTElementMergeInfo) GetTargetElementStatus() string`

GetTargetElementStatus returns the TargetElementStatus field if non-nil, zero value otherwise.

### GetTargetElementStatusOk

`func (o *BTElementMergeInfo) GetTargetElementStatusOk() (*string, bool)`

GetTargetElementStatusOk returns a tuple with the TargetElementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetElementStatus

`func (o *BTElementMergeInfo) SetTargetElementStatus(v string)`

SetTargetElementStatus sets TargetElementStatus field to given value.

### HasTargetElementStatus

`func (o *BTElementMergeInfo) HasTargetElementStatus() bool`

HasTargetElementStatus returns a boolean if a field has been set.

### GetTargetModifiedAt

`func (o *BTElementMergeInfo) GetTargetModifiedAt() JSONTime`

GetTargetModifiedAt returns the TargetModifiedAt field if non-nil, zero value otherwise.

### GetTargetModifiedAtOk

`func (o *BTElementMergeInfo) GetTargetModifiedAtOk() (*JSONTime, bool)`

GetTargetModifiedAtOk returns a tuple with the TargetModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetModifiedAt

`func (o *BTElementMergeInfo) SetTargetModifiedAt(v JSONTime)`

SetTargetModifiedAt sets TargetModifiedAt field to given value.

### HasTargetModifiedAt

`func (o *BTElementMergeInfo) HasTargetModifiedAt() bool`

HasTargetModifiedAt returns a boolean if a field has been set.

### GetTargetModifiedBy

`func (o *BTElementMergeInfo) GetTargetModifiedBy() BTUserBasicSummaryInfo`

GetTargetModifiedBy returns the TargetModifiedBy field if non-nil, zero value otherwise.

### GetTargetModifiedByOk

`func (o *BTElementMergeInfo) GetTargetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetTargetModifiedByOk returns a tuple with the TargetModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetModifiedBy

`func (o *BTElementMergeInfo) SetTargetModifiedBy(v BTUserBasicSummaryInfo)`

SetTargetModifiedBy sets TargetModifiedBy field to given value.

### HasTargetModifiedBy

`func (o *BTElementMergeInfo) HasTargetModifiedBy() bool`

HasTargetModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


