# BTExportTessellatedFacesResponse898

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bodies** | Pointer to [**[]BTExportTessellatedBody3398**](BTExportTessellatedBody3398.md) |  | [optional] 
**BodiesInfo** | Pointer to [**BTExportModelBodiesResponse734**](BTExportModelBodiesResponse734.md) |  | [optional] 
**BtType** | Pointer to **string** |  | [optional] 
**CombineCompositePartConstituents** | Pointer to **bool** |  | [optional] 
**DisplayData** | Pointer to [**BTPartStudioDisplayData346**](BTPartStudioDisplayData346.md) |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**ErrorEnum** | Pointer to [**GBTErrorStringEnum**](GBTErrorStringEnum.md) |  | [optional] 
**FacetPoints** | Pointer to [**[]BTVector3d389**](BTVector3d389.md) |  | [optional] 
**FullElementId** | Pointer to [**BTFullElementId756**](BTFullElementId756.md) |  | [optional] 
**OutputFaceAppearances** | Pointer to **bool** |  | [optional] 
**OutputSeparateFaceNodes** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTExportTessellatedFacesResponse898

`func NewBTExportTessellatedFacesResponse898() *BTExportTessellatedFacesResponse898`

NewBTExportTessellatedFacesResponse898 instantiates a new BTExportTessellatedFacesResponse898 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTExportTessellatedFacesResponse898WithDefaults

`func NewBTExportTessellatedFacesResponse898WithDefaults() *BTExportTessellatedFacesResponse898`

NewBTExportTessellatedFacesResponse898WithDefaults instantiates a new BTExportTessellatedFacesResponse898 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodies

`func (o *BTExportTessellatedFacesResponse898) GetBodies() []BTExportTessellatedBody3398`

GetBodies returns the Bodies field if non-nil, zero value otherwise.

### GetBodiesOk

`func (o *BTExportTessellatedFacesResponse898) GetBodiesOk() (*[]BTExportTessellatedBody3398, bool)`

GetBodiesOk returns a tuple with the Bodies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodies

`func (o *BTExportTessellatedFacesResponse898) SetBodies(v []BTExportTessellatedBody3398)`

SetBodies sets Bodies field to given value.

### HasBodies

`func (o *BTExportTessellatedFacesResponse898) HasBodies() bool`

HasBodies returns a boolean if a field has been set.

### GetBodiesInfo

`func (o *BTExportTessellatedFacesResponse898) GetBodiesInfo() BTExportModelBodiesResponse734`

GetBodiesInfo returns the BodiesInfo field if non-nil, zero value otherwise.

### GetBodiesInfoOk

`func (o *BTExportTessellatedFacesResponse898) GetBodiesInfoOk() (*BTExportModelBodiesResponse734, bool)`

GetBodiesInfoOk returns a tuple with the BodiesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodiesInfo

`func (o *BTExportTessellatedFacesResponse898) SetBodiesInfo(v BTExportModelBodiesResponse734)`

SetBodiesInfo sets BodiesInfo field to given value.

### HasBodiesInfo

`func (o *BTExportTessellatedFacesResponse898) HasBodiesInfo() bool`

HasBodiesInfo returns a boolean if a field has been set.

### GetBtType

`func (o *BTExportTessellatedFacesResponse898) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTExportTessellatedFacesResponse898) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTExportTessellatedFacesResponse898) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTExportTessellatedFacesResponse898) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCombineCompositePartConstituents

`func (o *BTExportTessellatedFacesResponse898) GetCombineCompositePartConstituents() bool`

GetCombineCompositePartConstituents returns the CombineCompositePartConstituents field if non-nil, zero value otherwise.

### GetCombineCompositePartConstituentsOk

`func (o *BTExportTessellatedFacesResponse898) GetCombineCompositePartConstituentsOk() (*bool, bool)`

GetCombineCompositePartConstituentsOk returns a tuple with the CombineCompositePartConstituents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombineCompositePartConstituents

`func (o *BTExportTessellatedFacesResponse898) SetCombineCompositePartConstituents(v bool)`

SetCombineCompositePartConstituents sets CombineCompositePartConstituents field to given value.

### HasCombineCompositePartConstituents

`func (o *BTExportTessellatedFacesResponse898) HasCombineCompositePartConstituents() bool`

HasCombineCompositePartConstituents returns a boolean if a field has been set.

### GetDisplayData

`func (o *BTExportTessellatedFacesResponse898) GetDisplayData() BTPartStudioDisplayData346`

GetDisplayData returns the DisplayData field if non-nil, zero value otherwise.

### GetDisplayDataOk

`func (o *BTExportTessellatedFacesResponse898) GetDisplayDataOk() (*BTPartStudioDisplayData346, bool)`

GetDisplayDataOk returns a tuple with the DisplayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayData

`func (o *BTExportTessellatedFacesResponse898) SetDisplayData(v BTPartStudioDisplayData346)`

SetDisplayData sets DisplayData field to given value.

### HasDisplayData

`func (o *BTExportTessellatedFacesResponse898) HasDisplayData() bool`

HasDisplayData returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTExportTessellatedFacesResponse898) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTExportTessellatedFacesResponse898) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTExportTessellatedFacesResponse898) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTExportTessellatedFacesResponse898) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTExportTessellatedFacesResponse898) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTExportTessellatedFacesResponse898) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTExportTessellatedFacesResponse898) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTExportTessellatedFacesResponse898) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetErrorEnum

`func (o *BTExportTessellatedFacesResponse898) GetErrorEnum() GBTErrorStringEnum`

GetErrorEnum returns the ErrorEnum field if non-nil, zero value otherwise.

### GetErrorEnumOk

`func (o *BTExportTessellatedFacesResponse898) GetErrorEnumOk() (*GBTErrorStringEnum, bool)`

GetErrorEnumOk returns a tuple with the ErrorEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorEnum

`func (o *BTExportTessellatedFacesResponse898) SetErrorEnum(v GBTErrorStringEnum)`

SetErrorEnum sets ErrorEnum field to given value.

### HasErrorEnum

`func (o *BTExportTessellatedFacesResponse898) HasErrorEnum() bool`

HasErrorEnum returns a boolean if a field has been set.

### GetFacetPoints

`func (o *BTExportTessellatedFacesResponse898) GetFacetPoints() []BTVector3d389`

GetFacetPoints returns the FacetPoints field if non-nil, zero value otherwise.

### GetFacetPointsOk

`func (o *BTExportTessellatedFacesResponse898) GetFacetPointsOk() (*[]BTVector3d389, bool)`

GetFacetPointsOk returns a tuple with the FacetPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetPoints

`func (o *BTExportTessellatedFacesResponse898) SetFacetPoints(v []BTVector3d389)`

SetFacetPoints sets FacetPoints field to given value.

### HasFacetPoints

`func (o *BTExportTessellatedFacesResponse898) HasFacetPoints() bool`

HasFacetPoints returns a boolean if a field has been set.

### GetFullElementId

`func (o *BTExportTessellatedFacesResponse898) GetFullElementId() BTFullElementId756`

GetFullElementId returns the FullElementId field if non-nil, zero value otherwise.

### GetFullElementIdOk

`func (o *BTExportTessellatedFacesResponse898) GetFullElementIdOk() (*BTFullElementId756, bool)`

GetFullElementIdOk returns a tuple with the FullElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullElementId

`func (o *BTExportTessellatedFacesResponse898) SetFullElementId(v BTFullElementId756)`

SetFullElementId sets FullElementId field to given value.

### HasFullElementId

`func (o *BTExportTessellatedFacesResponse898) HasFullElementId() bool`

HasFullElementId returns a boolean if a field has been set.

### GetOutputFaceAppearances

`func (o *BTExportTessellatedFacesResponse898) GetOutputFaceAppearances() bool`

GetOutputFaceAppearances returns the OutputFaceAppearances field if non-nil, zero value otherwise.

### GetOutputFaceAppearancesOk

`func (o *BTExportTessellatedFacesResponse898) GetOutputFaceAppearancesOk() (*bool, bool)`

GetOutputFaceAppearancesOk returns a tuple with the OutputFaceAppearances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFaceAppearances

`func (o *BTExportTessellatedFacesResponse898) SetOutputFaceAppearances(v bool)`

SetOutputFaceAppearances sets OutputFaceAppearances field to given value.

### HasOutputFaceAppearances

`func (o *BTExportTessellatedFacesResponse898) HasOutputFaceAppearances() bool`

HasOutputFaceAppearances returns a boolean if a field has been set.

### GetOutputSeparateFaceNodes

`func (o *BTExportTessellatedFacesResponse898) GetOutputSeparateFaceNodes() bool`

GetOutputSeparateFaceNodes returns the OutputSeparateFaceNodes field if non-nil, zero value otherwise.

### GetOutputSeparateFaceNodesOk

`func (o *BTExportTessellatedFacesResponse898) GetOutputSeparateFaceNodesOk() (*bool, bool)`

GetOutputSeparateFaceNodesOk returns a tuple with the OutputSeparateFaceNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSeparateFaceNodes

`func (o *BTExportTessellatedFacesResponse898) SetOutputSeparateFaceNodes(v bool)`

SetOutputSeparateFaceNodes sets OutputSeparateFaceNodes field to given value.

### HasOutputSeparateFaceNodes

`func (o *BTExportTessellatedFacesResponse898) HasOutputSeparateFaceNodes() bool`

HasOutputSeparateFaceNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


