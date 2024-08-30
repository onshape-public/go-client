# BTAppDrawingViewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociativityChangeId** | Pointer to **string** |  | [optional] 
**BomReferenceId** | Pointer to **string** |  | [optional] 
**BrokenOutBBoxes** | Pointer to [**map[string]BTBoundingBox1052**](BTBoundingBox1052.md) |  | [optional] 
**BrokenOutEndConditions** | Pointer to [**map[string]BTBrokenOutEndCondition1107**](BTBrokenOutEndCondition1107.md) |  | [optional] 
**BrokenOutPointNumbers** | Pointer to **[]int32** |  | [optional] 
**ChangeId** | Pointer to **string** |  | [optional] 
**ComputeIntersection** | Pointer to **bool** |  | [optional] 
**CutPoint** | Pointer to **[]float64** |  | [optional] 
**DepthSectionEndCondition** | Pointer to [**BTBrokenOutEndCondition1107**](BTBrokenOutEndCondition1107.md) |  | [optional] 
**DisplayStateId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** | &#x60;0: OK (healthy) | 1: INFO | 2: WARNING | 3: ERROR (dangling or view generation call failed) | 4: UNKNOWN&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 
**ExplodedViewId** | Pointer to **string** |  | [optional] 
**HasSecondaryViewDefinition** | Pointer to **bool** |  | [optional] 
**HiddenLines** | Pointer to **string** |  | [optional] 
**IgnoreFaultyParts** | Pointer to **bool** |  | [optional] 
**IncludeHiddenInstances** | Pointer to **bool** |  | [optional] 
**IncludeSurfaces** | Pointer to **bool** |  | [optional] 
**IncludeWires** | Pointer to **bool** |  | [optional] 
**IsAlignedSection** | Pointer to **bool** |  | [optional] 
**IsBrokenOutSection** | Pointer to **bool** |  | [optional] 
**IsCopiedView** | Pointer to **bool** |  | [optional] 
**IsCropView** | Pointer to **bool** |  | [optional] 
**IsPartialSection** | Pointer to **bool** |  | [optional] 
**IsSectionOfAlignedSection** | Pointer to **bool** |  | [optional] 
**IsSectionOfSectionOnBase** | Pointer to **bool** |  | [optional] 
**IsSurface** | Pointer to **bool** |  | [optional] 
**ModelReferenceId** | Pointer to **string** |  | [optional] 
**ModificationId** | Pointer to **string** |  | [optional] 
**NamedPositionId** | Pointer to **string** |  | [optional] 
**OccurrenceOrQueryToGeometryProperties** | Pointer to [**map[string]BTAppElementViewGeometryProperties1100**](BTAppElementViewGeometryProperties1100.md) |  | [optional] 
**OffsetSectionPoints** | Pointer to **[]float64** |  | [optional] 
**ParentChangeId** | Pointer to **string** |  | [optional] 
**ParentViewId** | Pointer to **string** |  | [optional] 
**Perspective** | Pointer to **bool** |  | [optional] 
**ProjectionAngle** | Pointer to **string** |  | [optional] 
**QualityOption** | Pointer to **int32** |  | [optional] 
**RenderSketches** | Pointer to **bool** |  | [optional] 
**SectionId** | Pointer to **string** |  | [optional] 
**ShowAutoCenterlines** | Pointer to **bool** |  | [optional] 
**ShowAutoCentermarks** | Pointer to **bool** |  | [optional] 
**ShowCutGeomOnly** | Pointer to **bool** |  | [optional] 
**ShowTangentLines** | Pointer to **bool** |  | [optional] 
**ShowThreads** | Pointer to **bool** |  | [optional] 
**ShowViewingPlane** | Pointer to **bool** |  | [optional] 
**SimplificationOption** | Pointer to **int32** |  | [optional] 
**SimplificationThreshold** | Pointer to **float64** |  | [optional] 
**UseParentViewSectionData** | Pointer to **bool** |  | [optional] 
**ViewDirection** | Pointer to **[]float64** |  | [optional] 
**ViewId** | Pointer to **string** |  | [optional] 
**ViewMatrix** | Pointer to **[]float64** |  | [optional] 
**ViewVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTAppDrawingViewInfo

`func NewBTAppDrawingViewInfo() *BTAppDrawingViewInfo`

NewBTAppDrawingViewInfo instantiates a new BTAppDrawingViewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppDrawingViewInfoWithDefaults

`func NewBTAppDrawingViewInfoWithDefaults() *BTAppDrawingViewInfo`

NewBTAppDrawingViewInfoWithDefaults instantiates a new BTAppDrawingViewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociativityChangeId

`func (o *BTAppDrawingViewInfo) GetAssociativityChangeId() string`

GetAssociativityChangeId returns the AssociativityChangeId field if non-nil, zero value otherwise.

### GetAssociativityChangeIdOk

`func (o *BTAppDrawingViewInfo) GetAssociativityChangeIdOk() (*string, bool)`

GetAssociativityChangeIdOk returns a tuple with the AssociativityChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociativityChangeId

`func (o *BTAppDrawingViewInfo) SetAssociativityChangeId(v string)`

SetAssociativityChangeId sets AssociativityChangeId field to given value.

### HasAssociativityChangeId

`func (o *BTAppDrawingViewInfo) HasAssociativityChangeId() bool`

HasAssociativityChangeId returns a boolean if a field has been set.

### GetBomReferenceId

`func (o *BTAppDrawingViewInfo) GetBomReferenceId() string`

GetBomReferenceId returns the BomReferenceId field if non-nil, zero value otherwise.

### GetBomReferenceIdOk

`func (o *BTAppDrawingViewInfo) GetBomReferenceIdOk() (*string, bool)`

GetBomReferenceIdOk returns a tuple with the BomReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomReferenceId

`func (o *BTAppDrawingViewInfo) SetBomReferenceId(v string)`

SetBomReferenceId sets BomReferenceId field to given value.

### HasBomReferenceId

`func (o *BTAppDrawingViewInfo) HasBomReferenceId() bool`

HasBomReferenceId returns a boolean if a field has been set.

### GetBrokenOutBBoxes

`func (o *BTAppDrawingViewInfo) GetBrokenOutBBoxes() map[string]BTBoundingBox1052`

GetBrokenOutBBoxes returns the BrokenOutBBoxes field if non-nil, zero value otherwise.

### GetBrokenOutBBoxesOk

`func (o *BTAppDrawingViewInfo) GetBrokenOutBBoxesOk() (*map[string]BTBoundingBox1052, bool)`

GetBrokenOutBBoxesOk returns a tuple with the BrokenOutBBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenOutBBoxes

`func (o *BTAppDrawingViewInfo) SetBrokenOutBBoxes(v map[string]BTBoundingBox1052)`

SetBrokenOutBBoxes sets BrokenOutBBoxes field to given value.

### HasBrokenOutBBoxes

`func (o *BTAppDrawingViewInfo) HasBrokenOutBBoxes() bool`

HasBrokenOutBBoxes returns a boolean if a field has been set.

### GetBrokenOutEndConditions

`func (o *BTAppDrawingViewInfo) GetBrokenOutEndConditions() map[string]BTBrokenOutEndCondition1107`

GetBrokenOutEndConditions returns the BrokenOutEndConditions field if non-nil, zero value otherwise.

### GetBrokenOutEndConditionsOk

`func (o *BTAppDrawingViewInfo) GetBrokenOutEndConditionsOk() (*map[string]BTBrokenOutEndCondition1107, bool)`

GetBrokenOutEndConditionsOk returns a tuple with the BrokenOutEndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenOutEndConditions

`func (o *BTAppDrawingViewInfo) SetBrokenOutEndConditions(v map[string]BTBrokenOutEndCondition1107)`

SetBrokenOutEndConditions sets BrokenOutEndConditions field to given value.

### HasBrokenOutEndConditions

`func (o *BTAppDrawingViewInfo) HasBrokenOutEndConditions() bool`

HasBrokenOutEndConditions returns a boolean if a field has been set.

### GetBrokenOutPointNumbers

`func (o *BTAppDrawingViewInfo) GetBrokenOutPointNumbers() []int32`

GetBrokenOutPointNumbers returns the BrokenOutPointNumbers field if non-nil, zero value otherwise.

### GetBrokenOutPointNumbersOk

`func (o *BTAppDrawingViewInfo) GetBrokenOutPointNumbersOk() (*[]int32, bool)`

GetBrokenOutPointNumbersOk returns a tuple with the BrokenOutPointNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenOutPointNumbers

`func (o *BTAppDrawingViewInfo) SetBrokenOutPointNumbers(v []int32)`

SetBrokenOutPointNumbers sets BrokenOutPointNumbers field to given value.

### HasBrokenOutPointNumbers

`func (o *BTAppDrawingViewInfo) HasBrokenOutPointNumbers() bool`

HasBrokenOutPointNumbers returns a boolean if a field has been set.

### GetChangeId

`func (o *BTAppDrawingViewInfo) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *BTAppDrawingViewInfo) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *BTAppDrawingViewInfo) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *BTAppDrawingViewInfo) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetComputeIntersection

`func (o *BTAppDrawingViewInfo) GetComputeIntersection() bool`

GetComputeIntersection returns the ComputeIntersection field if non-nil, zero value otherwise.

### GetComputeIntersectionOk

`func (o *BTAppDrawingViewInfo) GetComputeIntersectionOk() (*bool, bool)`

GetComputeIntersectionOk returns a tuple with the ComputeIntersection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeIntersection

`func (o *BTAppDrawingViewInfo) SetComputeIntersection(v bool)`

SetComputeIntersection sets ComputeIntersection field to given value.

### HasComputeIntersection

`func (o *BTAppDrawingViewInfo) HasComputeIntersection() bool`

HasComputeIntersection returns a boolean if a field has been set.

### GetCutPoint

`func (o *BTAppDrawingViewInfo) GetCutPoint() []float64`

GetCutPoint returns the CutPoint field if non-nil, zero value otherwise.

### GetCutPointOk

`func (o *BTAppDrawingViewInfo) GetCutPointOk() (*[]float64, bool)`

GetCutPointOk returns a tuple with the CutPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutPoint

`func (o *BTAppDrawingViewInfo) SetCutPoint(v []float64)`

SetCutPoint sets CutPoint field to given value.

### HasCutPoint

`func (o *BTAppDrawingViewInfo) HasCutPoint() bool`

HasCutPoint returns a boolean if a field has been set.

### GetDepthSectionEndCondition

`func (o *BTAppDrawingViewInfo) GetDepthSectionEndCondition() BTBrokenOutEndCondition1107`

GetDepthSectionEndCondition returns the DepthSectionEndCondition field if non-nil, zero value otherwise.

### GetDepthSectionEndConditionOk

`func (o *BTAppDrawingViewInfo) GetDepthSectionEndConditionOk() (*BTBrokenOutEndCondition1107, bool)`

GetDepthSectionEndConditionOk returns a tuple with the DepthSectionEndCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepthSectionEndCondition

`func (o *BTAppDrawingViewInfo) SetDepthSectionEndCondition(v BTBrokenOutEndCondition1107)`

SetDepthSectionEndCondition sets DepthSectionEndCondition field to given value.

### HasDepthSectionEndCondition

`func (o *BTAppDrawingViewInfo) HasDepthSectionEndCondition() bool`

HasDepthSectionEndCondition returns a boolean if a field has been set.

### GetDisplayStateId

`func (o *BTAppDrawingViewInfo) GetDisplayStateId() string`

GetDisplayStateId returns the DisplayStateId field if non-nil, zero value otherwise.

### GetDisplayStateIdOk

`func (o *BTAppDrawingViewInfo) GetDisplayStateIdOk() (*string, bool)`

GetDisplayStateIdOk returns a tuple with the DisplayStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStateId

`func (o *BTAppDrawingViewInfo) SetDisplayStateId(v string)`

SetDisplayStateId sets DisplayStateId field to given value.

### HasDisplayStateId

`func (o *BTAppDrawingViewInfo) HasDisplayStateId() bool`

HasDisplayStateId returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppDrawingViewInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppDrawingViewInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppDrawingViewInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppDrawingViewInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppDrawingViewInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppDrawingViewInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppDrawingViewInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppDrawingViewInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppDrawingViewInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppDrawingViewInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppDrawingViewInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppDrawingViewInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### GetExplodedViewId

`func (o *BTAppDrawingViewInfo) GetExplodedViewId() string`

GetExplodedViewId returns the ExplodedViewId field if non-nil, zero value otherwise.

### GetExplodedViewIdOk

`func (o *BTAppDrawingViewInfo) GetExplodedViewIdOk() (*string, bool)`

GetExplodedViewIdOk returns a tuple with the ExplodedViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplodedViewId

`func (o *BTAppDrawingViewInfo) SetExplodedViewId(v string)`

SetExplodedViewId sets ExplodedViewId field to given value.

### HasExplodedViewId

`func (o *BTAppDrawingViewInfo) HasExplodedViewId() bool`

HasExplodedViewId returns a boolean if a field has been set.

### GetHasSecondaryViewDefinition

`func (o *BTAppDrawingViewInfo) GetHasSecondaryViewDefinition() bool`

GetHasSecondaryViewDefinition returns the HasSecondaryViewDefinition field if non-nil, zero value otherwise.

### GetHasSecondaryViewDefinitionOk

`func (o *BTAppDrawingViewInfo) GetHasSecondaryViewDefinitionOk() (*bool, bool)`

GetHasSecondaryViewDefinitionOk returns a tuple with the HasSecondaryViewDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecondaryViewDefinition

`func (o *BTAppDrawingViewInfo) SetHasSecondaryViewDefinition(v bool)`

SetHasSecondaryViewDefinition sets HasSecondaryViewDefinition field to given value.

### HasHasSecondaryViewDefinition

`func (o *BTAppDrawingViewInfo) HasHasSecondaryViewDefinition() bool`

HasHasSecondaryViewDefinition returns a boolean if a field has been set.

### GetHiddenLines

`func (o *BTAppDrawingViewInfo) GetHiddenLines() string`

GetHiddenLines returns the HiddenLines field if non-nil, zero value otherwise.

### GetHiddenLinesOk

`func (o *BTAppDrawingViewInfo) GetHiddenLinesOk() (*string, bool)`

GetHiddenLinesOk returns a tuple with the HiddenLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenLines

`func (o *BTAppDrawingViewInfo) SetHiddenLines(v string)`

SetHiddenLines sets HiddenLines field to given value.

### HasHiddenLines

`func (o *BTAppDrawingViewInfo) HasHiddenLines() bool`

HasHiddenLines returns a boolean if a field has been set.

### GetIgnoreFaultyParts

`func (o *BTAppDrawingViewInfo) GetIgnoreFaultyParts() bool`

GetIgnoreFaultyParts returns the IgnoreFaultyParts field if non-nil, zero value otherwise.

### GetIgnoreFaultyPartsOk

`func (o *BTAppDrawingViewInfo) GetIgnoreFaultyPartsOk() (*bool, bool)`

GetIgnoreFaultyPartsOk returns a tuple with the IgnoreFaultyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFaultyParts

`func (o *BTAppDrawingViewInfo) SetIgnoreFaultyParts(v bool)`

SetIgnoreFaultyParts sets IgnoreFaultyParts field to given value.

### HasIgnoreFaultyParts

`func (o *BTAppDrawingViewInfo) HasIgnoreFaultyParts() bool`

HasIgnoreFaultyParts returns a boolean if a field has been set.

### GetIncludeHiddenInstances

`func (o *BTAppDrawingViewInfo) GetIncludeHiddenInstances() bool`

GetIncludeHiddenInstances returns the IncludeHiddenInstances field if non-nil, zero value otherwise.

### GetIncludeHiddenInstancesOk

`func (o *BTAppDrawingViewInfo) GetIncludeHiddenInstancesOk() (*bool, bool)`

GetIncludeHiddenInstancesOk returns a tuple with the IncludeHiddenInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHiddenInstances

`func (o *BTAppDrawingViewInfo) SetIncludeHiddenInstances(v bool)`

SetIncludeHiddenInstances sets IncludeHiddenInstances field to given value.

### HasIncludeHiddenInstances

`func (o *BTAppDrawingViewInfo) HasIncludeHiddenInstances() bool`

HasIncludeHiddenInstances returns a boolean if a field has been set.

### GetIncludeSurfaces

`func (o *BTAppDrawingViewInfo) GetIncludeSurfaces() bool`

GetIncludeSurfaces returns the IncludeSurfaces field if non-nil, zero value otherwise.

### GetIncludeSurfacesOk

`func (o *BTAppDrawingViewInfo) GetIncludeSurfacesOk() (*bool, bool)`

GetIncludeSurfacesOk returns a tuple with the IncludeSurfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSurfaces

`func (o *BTAppDrawingViewInfo) SetIncludeSurfaces(v bool)`

SetIncludeSurfaces sets IncludeSurfaces field to given value.

### HasIncludeSurfaces

`func (o *BTAppDrawingViewInfo) HasIncludeSurfaces() bool`

HasIncludeSurfaces returns a boolean if a field has been set.

### GetIncludeWires

`func (o *BTAppDrawingViewInfo) GetIncludeWires() bool`

GetIncludeWires returns the IncludeWires field if non-nil, zero value otherwise.

### GetIncludeWiresOk

`func (o *BTAppDrawingViewInfo) GetIncludeWiresOk() (*bool, bool)`

GetIncludeWiresOk returns a tuple with the IncludeWires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWires

`func (o *BTAppDrawingViewInfo) SetIncludeWires(v bool)`

SetIncludeWires sets IncludeWires field to given value.

### HasIncludeWires

`func (o *BTAppDrawingViewInfo) HasIncludeWires() bool`

HasIncludeWires returns a boolean if a field has been set.

### GetIsAlignedSection

`func (o *BTAppDrawingViewInfo) GetIsAlignedSection() bool`

GetIsAlignedSection returns the IsAlignedSection field if non-nil, zero value otherwise.

### GetIsAlignedSectionOk

`func (o *BTAppDrawingViewInfo) GetIsAlignedSectionOk() (*bool, bool)`

GetIsAlignedSectionOk returns a tuple with the IsAlignedSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlignedSection

`func (o *BTAppDrawingViewInfo) SetIsAlignedSection(v bool)`

SetIsAlignedSection sets IsAlignedSection field to given value.

### HasIsAlignedSection

`func (o *BTAppDrawingViewInfo) HasIsAlignedSection() bool`

HasIsAlignedSection returns a boolean if a field has been set.

### GetIsBrokenOutSection

`func (o *BTAppDrawingViewInfo) GetIsBrokenOutSection() bool`

GetIsBrokenOutSection returns the IsBrokenOutSection field if non-nil, zero value otherwise.

### GetIsBrokenOutSectionOk

`func (o *BTAppDrawingViewInfo) GetIsBrokenOutSectionOk() (*bool, bool)`

GetIsBrokenOutSectionOk returns a tuple with the IsBrokenOutSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrokenOutSection

`func (o *BTAppDrawingViewInfo) SetIsBrokenOutSection(v bool)`

SetIsBrokenOutSection sets IsBrokenOutSection field to given value.

### HasIsBrokenOutSection

`func (o *BTAppDrawingViewInfo) HasIsBrokenOutSection() bool`

HasIsBrokenOutSection returns a boolean if a field has been set.

### GetIsCopiedView

`func (o *BTAppDrawingViewInfo) GetIsCopiedView() bool`

GetIsCopiedView returns the IsCopiedView field if non-nil, zero value otherwise.

### GetIsCopiedViewOk

`func (o *BTAppDrawingViewInfo) GetIsCopiedViewOk() (*bool, bool)`

GetIsCopiedViewOk returns a tuple with the IsCopiedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCopiedView

`func (o *BTAppDrawingViewInfo) SetIsCopiedView(v bool)`

SetIsCopiedView sets IsCopiedView field to given value.

### HasIsCopiedView

`func (o *BTAppDrawingViewInfo) HasIsCopiedView() bool`

HasIsCopiedView returns a boolean if a field has been set.

### GetIsCropView

`func (o *BTAppDrawingViewInfo) GetIsCropView() bool`

GetIsCropView returns the IsCropView field if non-nil, zero value otherwise.

### GetIsCropViewOk

`func (o *BTAppDrawingViewInfo) GetIsCropViewOk() (*bool, bool)`

GetIsCropViewOk returns a tuple with the IsCropView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCropView

`func (o *BTAppDrawingViewInfo) SetIsCropView(v bool)`

SetIsCropView sets IsCropView field to given value.

### HasIsCropView

`func (o *BTAppDrawingViewInfo) HasIsCropView() bool`

HasIsCropView returns a boolean if a field has been set.

### GetIsPartialSection

`func (o *BTAppDrawingViewInfo) GetIsPartialSection() bool`

GetIsPartialSection returns the IsPartialSection field if non-nil, zero value otherwise.

### GetIsPartialSectionOk

`func (o *BTAppDrawingViewInfo) GetIsPartialSectionOk() (*bool, bool)`

GetIsPartialSectionOk returns a tuple with the IsPartialSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartialSection

`func (o *BTAppDrawingViewInfo) SetIsPartialSection(v bool)`

SetIsPartialSection sets IsPartialSection field to given value.

### HasIsPartialSection

`func (o *BTAppDrawingViewInfo) HasIsPartialSection() bool`

HasIsPartialSection returns a boolean if a field has been set.

### GetIsSectionOfAlignedSection

`func (o *BTAppDrawingViewInfo) GetIsSectionOfAlignedSection() bool`

GetIsSectionOfAlignedSection returns the IsSectionOfAlignedSection field if non-nil, zero value otherwise.

### GetIsSectionOfAlignedSectionOk

`func (o *BTAppDrawingViewInfo) GetIsSectionOfAlignedSectionOk() (*bool, bool)`

GetIsSectionOfAlignedSectionOk returns a tuple with the IsSectionOfAlignedSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSectionOfAlignedSection

`func (o *BTAppDrawingViewInfo) SetIsSectionOfAlignedSection(v bool)`

SetIsSectionOfAlignedSection sets IsSectionOfAlignedSection field to given value.

### HasIsSectionOfAlignedSection

`func (o *BTAppDrawingViewInfo) HasIsSectionOfAlignedSection() bool`

HasIsSectionOfAlignedSection returns a boolean if a field has been set.

### GetIsSectionOfSectionOnBase

`func (o *BTAppDrawingViewInfo) GetIsSectionOfSectionOnBase() bool`

GetIsSectionOfSectionOnBase returns the IsSectionOfSectionOnBase field if non-nil, zero value otherwise.

### GetIsSectionOfSectionOnBaseOk

`func (o *BTAppDrawingViewInfo) GetIsSectionOfSectionOnBaseOk() (*bool, bool)`

GetIsSectionOfSectionOnBaseOk returns a tuple with the IsSectionOfSectionOnBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSectionOfSectionOnBase

`func (o *BTAppDrawingViewInfo) SetIsSectionOfSectionOnBase(v bool)`

SetIsSectionOfSectionOnBase sets IsSectionOfSectionOnBase field to given value.

### HasIsSectionOfSectionOnBase

`func (o *BTAppDrawingViewInfo) HasIsSectionOfSectionOnBase() bool`

HasIsSectionOfSectionOnBase returns a boolean if a field has been set.

### GetIsSurface

`func (o *BTAppDrawingViewInfo) GetIsSurface() bool`

GetIsSurface returns the IsSurface field if non-nil, zero value otherwise.

### GetIsSurfaceOk

`func (o *BTAppDrawingViewInfo) GetIsSurfaceOk() (*bool, bool)`

GetIsSurfaceOk returns a tuple with the IsSurface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurface

`func (o *BTAppDrawingViewInfo) SetIsSurface(v bool)`

SetIsSurface sets IsSurface field to given value.

### HasIsSurface

`func (o *BTAppDrawingViewInfo) HasIsSurface() bool`

HasIsSurface returns a boolean if a field has been set.

### GetModelReferenceId

`func (o *BTAppDrawingViewInfo) GetModelReferenceId() string`

GetModelReferenceId returns the ModelReferenceId field if non-nil, zero value otherwise.

### GetModelReferenceIdOk

`func (o *BTAppDrawingViewInfo) GetModelReferenceIdOk() (*string, bool)`

GetModelReferenceIdOk returns a tuple with the ModelReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelReferenceId

`func (o *BTAppDrawingViewInfo) SetModelReferenceId(v string)`

SetModelReferenceId sets ModelReferenceId field to given value.

### HasModelReferenceId

`func (o *BTAppDrawingViewInfo) HasModelReferenceId() bool`

HasModelReferenceId returns a boolean if a field has been set.

### GetModificationId

`func (o *BTAppDrawingViewInfo) GetModificationId() string`

GetModificationId returns the ModificationId field if non-nil, zero value otherwise.

### GetModificationIdOk

`func (o *BTAppDrawingViewInfo) GetModificationIdOk() (*string, bool)`

GetModificationIdOk returns a tuple with the ModificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationId

`func (o *BTAppDrawingViewInfo) SetModificationId(v string)`

SetModificationId sets ModificationId field to given value.

### HasModificationId

`func (o *BTAppDrawingViewInfo) HasModificationId() bool`

HasModificationId returns a boolean if a field has been set.

### GetNamedPositionId

`func (o *BTAppDrawingViewInfo) GetNamedPositionId() string`

GetNamedPositionId returns the NamedPositionId field if non-nil, zero value otherwise.

### GetNamedPositionIdOk

`func (o *BTAppDrawingViewInfo) GetNamedPositionIdOk() (*string, bool)`

GetNamedPositionIdOk returns a tuple with the NamedPositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedPositionId

`func (o *BTAppDrawingViewInfo) SetNamedPositionId(v string)`

SetNamedPositionId sets NamedPositionId field to given value.

### HasNamedPositionId

`func (o *BTAppDrawingViewInfo) HasNamedPositionId() bool`

HasNamedPositionId returns a boolean if a field has been set.

### GetOccurrenceOrQueryToGeometryProperties

`func (o *BTAppDrawingViewInfo) GetOccurrenceOrQueryToGeometryProperties() map[string]BTAppElementViewGeometryProperties1100`

GetOccurrenceOrQueryToGeometryProperties returns the OccurrenceOrQueryToGeometryProperties field if non-nil, zero value otherwise.

### GetOccurrenceOrQueryToGeometryPropertiesOk

`func (o *BTAppDrawingViewInfo) GetOccurrenceOrQueryToGeometryPropertiesOk() (*map[string]BTAppElementViewGeometryProperties1100, bool)`

GetOccurrenceOrQueryToGeometryPropertiesOk returns a tuple with the OccurrenceOrQueryToGeometryProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceOrQueryToGeometryProperties

`func (o *BTAppDrawingViewInfo) SetOccurrenceOrQueryToGeometryProperties(v map[string]BTAppElementViewGeometryProperties1100)`

SetOccurrenceOrQueryToGeometryProperties sets OccurrenceOrQueryToGeometryProperties field to given value.

### HasOccurrenceOrQueryToGeometryProperties

`func (o *BTAppDrawingViewInfo) HasOccurrenceOrQueryToGeometryProperties() bool`

HasOccurrenceOrQueryToGeometryProperties returns a boolean if a field has been set.

### GetOffsetSectionPoints

`func (o *BTAppDrawingViewInfo) GetOffsetSectionPoints() []float64`

GetOffsetSectionPoints returns the OffsetSectionPoints field if non-nil, zero value otherwise.

### GetOffsetSectionPointsOk

`func (o *BTAppDrawingViewInfo) GetOffsetSectionPointsOk() (*[]float64, bool)`

GetOffsetSectionPointsOk returns a tuple with the OffsetSectionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetSectionPoints

`func (o *BTAppDrawingViewInfo) SetOffsetSectionPoints(v []float64)`

SetOffsetSectionPoints sets OffsetSectionPoints field to given value.

### HasOffsetSectionPoints

`func (o *BTAppDrawingViewInfo) HasOffsetSectionPoints() bool`

HasOffsetSectionPoints returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppDrawingViewInfo) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppDrawingViewInfo) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppDrawingViewInfo) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppDrawingViewInfo) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.

### GetParentViewId

`func (o *BTAppDrawingViewInfo) GetParentViewId() string`

GetParentViewId returns the ParentViewId field if non-nil, zero value otherwise.

### GetParentViewIdOk

`func (o *BTAppDrawingViewInfo) GetParentViewIdOk() (*string, bool)`

GetParentViewIdOk returns a tuple with the ParentViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentViewId

`func (o *BTAppDrawingViewInfo) SetParentViewId(v string)`

SetParentViewId sets ParentViewId field to given value.

### HasParentViewId

`func (o *BTAppDrawingViewInfo) HasParentViewId() bool`

HasParentViewId returns a boolean if a field has been set.

### GetPerspective

`func (o *BTAppDrawingViewInfo) GetPerspective() bool`

GetPerspective returns the Perspective field if non-nil, zero value otherwise.

### GetPerspectiveOk

`func (o *BTAppDrawingViewInfo) GetPerspectiveOk() (*bool, bool)`

GetPerspectiveOk returns a tuple with the Perspective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerspective

`func (o *BTAppDrawingViewInfo) SetPerspective(v bool)`

SetPerspective sets Perspective field to given value.

### HasPerspective

`func (o *BTAppDrawingViewInfo) HasPerspective() bool`

HasPerspective returns a boolean if a field has been set.

### GetProjectionAngle

`func (o *BTAppDrawingViewInfo) GetProjectionAngle() string`

GetProjectionAngle returns the ProjectionAngle field if non-nil, zero value otherwise.

### GetProjectionAngleOk

`func (o *BTAppDrawingViewInfo) GetProjectionAngleOk() (*string, bool)`

GetProjectionAngleOk returns a tuple with the ProjectionAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectionAngle

`func (o *BTAppDrawingViewInfo) SetProjectionAngle(v string)`

SetProjectionAngle sets ProjectionAngle field to given value.

### HasProjectionAngle

`func (o *BTAppDrawingViewInfo) HasProjectionAngle() bool`

HasProjectionAngle returns a boolean if a field has been set.

### GetQualityOption

`func (o *BTAppDrawingViewInfo) GetQualityOption() int32`

GetQualityOption returns the QualityOption field if non-nil, zero value otherwise.

### GetQualityOptionOk

`func (o *BTAppDrawingViewInfo) GetQualityOptionOk() (*int32, bool)`

GetQualityOptionOk returns a tuple with the QualityOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOption

`func (o *BTAppDrawingViewInfo) SetQualityOption(v int32)`

SetQualityOption sets QualityOption field to given value.

### HasQualityOption

`func (o *BTAppDrawingViewInfo) HasQualityOption() bool`

HasQualityOption returns a boolean if a field has been set.

### GetRenderSketches

`func (o *BTAppDrawingViewInfo) GetRenderSketches() bool`

GetRenderSketches returns the RenderSketches field if non-nil, zero value otherwise.

### GetRenderSketchesOk

`func (o *BTAppDrawingViewInfo) GetRenderSketchesOk() (*bool, bool)`

GetRenderSketchesOk returns a tuple with the RenderSketches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderSketches

`func (o *BTAppDrawingViewInfo) SetRenderSketches(v bool)`

SetRenderSketches sets RenderSketches field to given value.

### HasRenderSketches

`func (o *BTAppDrawingViewInfo) HasRenderSketches() bool`

HasRenderSketches returns a boolean if a field has been set.

### GetSectionId

`func (o *BTAppDrawingViewInfo) GetSectionId() string`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *BTAppDrawingViewInfo) GetSectionIdOk() (*string, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *BTAppDrawingViewInfo) SetSectionId(v string)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *BTAppDrawingViewInfo) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetShowAutoCenterlines

`func (o *BTAppDrawingViewInfo) GetShowAutoCenterlines() bool`

GetShowAutoCenterlines returns the ShowAutoCenterlines field if non-nil, zero value otherwise.

### GetShowAutoCenterlinesOk

`func (o *BTAppDrawingViewInfo) GetShowAutoCenterlinesOk() (*bool, bool)`

GetShowAutoCenterlinesOk returns a tuple with the ShowAutoCenterlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAutoCenterlines

`func (o *BTAppDrawingViewInfo) SetShowAutoCenterlines(v bool)`

SetShowAutoCenterlines sets ShowAutoCenterlines field to given value.

### HasShowAutoCenterlines

`func (o *BTAppDrawingViewInfo) HasShowAutoCenterlines() bool`

HasShowAutoCenterlines returns a boolean if a field has been set.

### GetShowAutoCentermarks

`func (o *BTAppDrawingViewInfo) GetShowAutoCentermarks() bool`

GetShowAutoCentermarks returns the ShowAutoCentermarks field if non-nil, zero value otherwise.

### GetShowAutoCentermarksOk

`func (o *BTAppDrawingViewInfo) GetShowAutoCentermarksOk() (*bool, bool)`

GetShowAutoCentermarksOk returns a tuple with the ShowAutoCentermarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAutoCentermarks

`func (o *BTAppDrawingViewInfo) SetShowAutoCentermarks(v bool)`

SetShowAutoCentermarks sets ShowAutoCentermarks field to given value.

### HasShowAutoCentermarks

`func (o *BTAppDrawingViewInfo) HasShowAutoCentermarks() bool`

HasShowAutoCentermarks returns a boolean if a field has been set.

### GetShowCutGeomOnly

`func (o *BTAppDrawingViewInfo) GetShowCutGeomOnly() bool`

GetShowCutGeomOnly returns the ShowCutGeomOnly field if non-nil, zero value otherwise.

### GetShowCutGeomOnlyOk

`func (o *BTAppDrawingViewInfo) GetShowCutGeomOnlyOk() (*bool, bool)`

GetShowCutGeomOnlyOk returns a tuple with the ShowCutGeomOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCutGeomOnly

`func (o *BTAppDrawingViewInfo) SetShowCutGeomOnly(v bool)`

SetShowCutGeomOnly sets ShowCutGeomOnly field to given value.

### HasShowCutGeomOnly

`func (o *BTAppDrawingViewInfo) HasShowCutGeomOnly() bool`

HasShowCutGeomOnly returns a boolean if a field has been set.

### GetShowTangentLines

`func (o *BTAppDrawingViewInfo) GetShowTangentLines() bool`

GetShowTangentLines returns the ShowTangentLines field if non-nil, zero value otherwise.

### GetShowTangentLinesOk

`func (o *BTAppDrawingViewInfo) GetShowTangentLinesOk() (*bool, bool)`

GetShowTangentLinesOk returns a tuple with the ShowTangentLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTangentLines

`func (o *BTAppDrawingViewInfo) SetShowTangentLines(v bool)`

SetShowTangentLines sets ShowTangentLines field to given value.

### HasShowTangentLines

`func (o *BTAppDrawingViewInfo) HasShowTangentLines() bool`

HasShowTangentLines returns a boolean if a field has been set.

### GetShowThreads

`func (o *BTAppDrawingViewInfo) GetShowThreads() bool`

GetShowThreads returns the ShowThreads field if non-nil, zero value otherwise.

### GetShowThreadsOk

`func (o *BTAppDrawingViewInfo) GetShowThreadsOk() (*bool, bool)`

GetShowThreadsOk returns a tuple with the ShowThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowThreads

`func (o *BTAppDrawingViewInfo) SetShowThreads(v bool)`

SetShowThreads sets ShowThreads field to given value.

### HasShowThreads

`func (o *BTAppDrawingViewInfo) HasShowThreads() bool`

HasShowThreads returns a boolean if a field has been set.

### GetShowViewingPlane

`func (o *BTAppDrawingViewInfo) GetShowViewingPlane() bool`

GetShowViewingPlane returns the ShowViewingPlane field if non-nil, zero value otherwise.

### GetShowViewingPlaneOk

`func (o *BTAppDrawingViewInfo) GetShowViewingPlaneOk() (*bool, bool)`

GetShowViewingPlaneOk returns a tuple with the ShowViewingPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowViewingPlane

`func (o *BTAppDrawingViewInfo) SetShowViewingPlane(v bool)`

SetShowViewingPlane sets ShowViewingPlane field to given value.

### HasShowViewingPlane

`func (o *BTAppDrawingViewInfo) HasShowViewingPlane() bool`

HasShowViewingPlane returns a boolean if a field has been set.

### GetSimplificationOption

`func (o *BTAppDrawingViewInfo) GetSimplificationOption() int32`

GetSimplificationOption returns the SimplificationOption field if non-nil, zero value otherwise.

### GetSimplificationOptionOk

`func (o *BTAppDrawingViewInfo) GetSimplificationOptionOk() (*int32, bool)`

GetSimplificationOptionOk returns a tuple with the SimplificationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplificationOption

`func (o *BTAppDrawingViewInfo) SetSimplificationOption(v int32)`

SetSimplificationOption sets SimplificationOption field to given value.

### HasSimplificationOption

`func (o *BTAppDrawingViewInfo) HasSimplificationOption() bool`

HasSimplificationOption returns a boolean if a field has been set.

### GetSimplificationThreshold

`func (o *BTAppDrawingViewInfo) GetSimplificationThreshold() float64`

GetSimplificationThreshold returns the SimplificationThreshold field if non-nil, zero value otherwise.

### GetSimplificationThresholdOk

`func (o *BTAppDrawingViewInfo) GetSimplificationThresholdOk() (*float64, bool)`

GetSimplificationThresholdOk returns a tuple with the SimplificationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplificationThreshold

`func (o *BTAppDrawingViewInfo) SetSimplificationThreshold(v float64)`

SetSimplificationThreshold sets SimplificationThreshold field to given value.

### HasSimplificationThreshold

`func (o *BTAppDrawingViewInfo) HasSimplificationThreshold() bool`

HasSimplificationThreshold returns a boolean if a field has been set.

### GetUseParentViewSectionData

`func (o *BTAppDrawingViewInfo) GetUseParentViewSectionData() bool`

GetUseParentViewSectionData returns the UseParentViewSectionData field if non-nil, zero value otherwise.

### GetUseParentViewSectionDataOk

`func (o *BTAppDrawingViewInfo) GetUseParentViewSectionDataOk() (*bool, bool)`

GetUseParentViewSectionDataOk returns a tuple with the UseParentViewSectionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseParentViewSectionData

`func (o *BTAppDrawingViewInfo) SetUseParentViewSectionData(v bool)`

SetUseParentViewSectionData sets UseParentViewSectionData field to given value.

### HasUseParentViewSectionData

`func (o *BTAppDrawingViewInfo) HasUseParentViewSectionData() bool`

HasUseParentViewSectionData returns a boolean if a field has been set.

### GetViewDirection

`func (o *BTAppDrawingViewInfo) GetViewDirection() []float64`

GetViewDirection returns the ViewDirection field if non-nil, zero value otherwise.

### GetViewDirectionOk

`func (o *BTAppDrawingViewInfo) GetViewDirectionOk() (*[]float64, bool)`

GetViewDirectionOk returns a tuple with the ViewDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDirection

`func (o *BTAppDrawingViewInfo) SetViewDirection(v []float64)`

SetViewDirection sets ViewDirection field to given value.

### HasViewDirection

`func (o *BTAppDrawingViewInfo) HasViewDirection() bool`

HasViewDirection returns a boolean if a field has been set.

### GetViewId

`func (o *BTAppDrawingViewInfo) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *BTAppDrawingViewInfo) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *BTAppDrawingViewInfo) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *BTAppDrawingViewInfo) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewMatrix

`func (o *BTAppDrawingViewInfo) GetViewMatrix() []float64`

GetViewMatrix returns the ViewMatrix field if non-nil, zero value otherwise.

### GetViewMatrixOk

`func (o *BTAppDrawingViewInfo) GetViewMatrixOk() (*[]float64, bool)`

GetViewMatrixOk returns a tuple with the ViewMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatrix

`func (o *BTAppDrawingViewInfo) SetViewMatrix(v []float64)`

SetViewMatrix sets ViewMatrix field to given value.

### HasViewMatrix

`func (o *BTAppDrawingViewInfo) HasViewMatrix() bool`

HasViewMatrix returns a boolean if a field has been set.

### GetViewVersion

`func (o *BTAppDrawingViewInfo) GetViewVersion() int32`

GetViewVersion returns the ViewVersion field if non-nil, zero value otherwise.

### GetViewVersionOk

`func (o *BTAppDrawingViewInfo) GetViewVersionOk() (*int32, bool)`

GetViewVersionOk returns a tuple with the ViewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewVersion

`func (o *BTAppDrawingViewInfo) SetViewVersion(v int32)`

SetViewVersion sets ViewVersion field to given value.

### HasViewVersion

`func (o *BTAppDrawingViewInfo) HasViewVersion() bool`

HasViewVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


