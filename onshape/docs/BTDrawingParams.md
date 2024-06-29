# BTDrawingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Border** | Pointer to **bool** | Set to &#x60;true&#x60; to include a border in the drawing. | [optional] [default to false]
**ComputeIntersection** | Pointer to **bool** | Set to &#x60;true&#x60; to compute and display virtual edges (curves drawn where parts intersect). Leave as &#x60;false&#x60; to improve performance. | [optional] [default to false]
**DecimalSeparator** | Pointer to **string** | &#x60;PERIOD&#x60; | &#x60;COMMA&#x60; | [optional] [default to "PERIOD"]
**DisplayStateId** | Pointer to **string** | Apply this display state&#39;s properties to the drawing. | [optional] 
**DocumentId** | Pointer to **string** | The document in which to create the drawing. If used, this value must match the document ID (&#x60;did&#x60;) value provided in the URL. | [optional] 
**DocumentMicroversionId** | Pointer to **string** | Create a drawing of a part or assembly from this microversion. | [optional] 
**DrawingName** | Pointer to **string** | Provide a name for the drawing. | [optional] 
**ElementConfiguration** | Pointer to **string** | Apply this configuration from the source element to the drawing. | [optional] 
**ElementId** | Pointer to **string** | The id of the element in which to perform the operation. | [optional] 
**ElementMicroversionId** | Pointer to **string** | The id of the element microversion in which to perform the operation. | [optional] 
**ExplosionId** | Pointer to **string** | Apply this exploded view to the drawing. | [optional] 
**ExternalDocumentId** | Pointer to **string** | Create a drawing of an element from this external document. | [optional] 
**ExternalDocumentVersionId** | Pointer to **string** | Create a drawing of an element from this external document version. | [optional] 
**HiddenLines** | Pointer to [**BTDrawingHiddenLineOption**](BTDrawingHiddenLineOption.md) |  | [optional] 
**IncludeSurfaces** | Pointer to **bool** | Set to &#x60;true&#x60; to include surfaces in the drawing. | [optional] [default to false]
**IncludeWires** | Pointer to **bool** | Set to &#x60;true&#x60; to include wires in the drawing. | [optional] [default to false]
**IsFlattenedPart** | Pointer to **bool** | Set to &#x60;true&#x60; if creating a drawing from a flattened part. | [optional] [default to false]
**IsSketchOnly** | Pointer to **bool** | Set to &#x60;true&#x60; if creating a drawing of a sketch. | [optional] [default to false]
**IsSurface** | Pointer to **bool** | Set to &#x60;true&#x60; if creating a drawing from a surface. | [optional] [default to false]
**Language** | Pointer to **string** | Set the language for the drawing. Accepts any ISO 639-1 standard language code. | [optional] [default to "en-us"]
**Location** | Pointer to [**BTElementLocationParams**](BTElementLocationParams.md) |  | [optional] 
**ModelType** | Pointer to **string** | The type of model to include in the drawing: &#x60;partstudio&#x60; | &#x60;assembly&#x60; | [optional] 
**NamedPositionId** | Pointer to **string** | Apply this named view to the drawing. | [optional] 
**NumberHorizontalZones** | Pointer to **int32** | The number of horizontal zones to include in the drawing&#39;s graphics area. | [optional] 
**NumberVerticalZones** | Pointer to **int32** | The number of vertical zones to include in the drawing&#39;s graphics area. | [optional] 
**PartId** | Pointer to **string** | Include this part in the drawing. | [optional] 
**PartNumber** | Pointer to **string** | Include this part in the drawing. | [optional] 
**PartQuery** | Pointer to **string** | Include all parts found by the query in the drawing. | [optional] 
**Projection** | Pointer to **string** | Apply this projection to the drawing. | [optional] 
**PureSketch** | Pointer to **bool** | Set to &#x60;true&#x60; if creating the drawing of an empty sketch. | [optional] [default to false]
**QualityOption** | Pointer to **string** | &#x60;BEST_PERFORMANCE&#x60; | &#x60;BEST_QUALITY&#x60; | &#x60;BALANCED&#x60; | &#x60;ADAPTIVE&#x60; | [optional] 
**ReferenceType** | Pointer to **int32** | Specify the type of element to create the drawing from. &#x60;0: UNKNOWN&#x60; | &#x60;1: PARTSTUDIO&#x60; | &#x60;2: ASSEMBLY&#x60; | &#x60;3: PART&#x60; | &#x60;4: FLATTENED_PART&#x60; | &#x60;5: COMPOSITE_PART&#x60; | &#x60;6: MESH_PART&#x60; | &#x60;7: SURFACE&#x60; | &#x60;8: SKETCH&#x60; | &#x60;9: CURVE&#x60; | [optional] 
**ReferenceTypeEnum** | Pointer to [**GBTAppElementReferenceType**](GBTAppElementReferenceType.md) |  | [optional] 
**Revision** | Pointer to **string** | Create the drawing from this specific revision. | [optional] 
**ShowCutGeomOnly** | Pointer to **bool** | Set to &#x60;true&#x60; to show only cut geometry in the drawing. | [optional] [default to false]
**SimplificationOption** | Pointer to **string** | &#x60;NONE&#x60; | &#x60;ABSOLUTE&#x60; | &#x60;RATIO_TO_MODEL&#x60; | &#x60;RATIO_TO_BODY&#x60; | &#x60;AUTOMATIC&#x60; | [optional] 
**SimplificationThreshold** | Pointer to **float64** | &#x60;NONE&#x60; | &#x60;UNKNOWN&#x60; | &#x60;SMOOTH&#x60; | &#x60;DRAFTING&#x60; | [optional] 
**Size** | Pointer to **string** | Provide a size for the drawing. | [optional] 
**SketchIds** | Pointer to **[]string** | Include these sketches in the drawing. | [optional] 
**Standard** | Pointer to **string** | Provide the Standard to use in the drawing. | [optional] 
**StartZones** | Pointer to **string** | The zone in which to start the drawing. | [optional] 
**TemplateArgs** | Pointer to **[]string** | Provide any additional arguments for the template being used for this drawing. | [optional] 
**TemplateDocumentId** | Pointer to **string** | Apply the template from this document to the drawing. | [optional] 
**TemplateElementId** | Pointer to **string** | Apply the template from this element to the drawing. | [optional] 
**TemplateName** | Pointer to **string** | Apply this template to the drawing. | [optional] 
**TemplateVersionId** | Pointer to **string** | Apply the template from this version to the drawing. | [optional] 
**TemplateWorkspaceId** | Pointer to **string** | Apply the template from this workspace to the drawing. | [optional] 
**Titleblock** | Pointer to **bool** | Set to &#x60;true&#x60; to include a title block in the drawing. | [optional] [default to false]
**Units** | Pointer to **string** | Units for the element: &#x60;METER&#x60; | &#x60;CENTIMETER&#x60; | &#x60;MILLIMETER&#x60; | &#x60;INCH&#x60; | &#x60;FOOT&#x60; | &#x60;YARD&#x60; | [optional] 
**Views** | Pointer to **string** | Add these views to the drawing. | [optional] 
**WorkspaceId** | Pointer to **string** | Create a drawing of a part or assembly from this workspace. | [optional] 

## Methods

### NewBTDrawingParams

`func NewBTDrawingParams() *BTDrawingParams`

NewBTDrawingParams instantiates a new BTDrawingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDrawingParamsWithDefaults

`func NewBTDrawingParamsWithDefaults() *BTDrawingParams`

NewBTDrawingParamsWithDefaults instantiates a new BTDrawingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorder

`func (o *BTDrawingParams) GetBorder() bool`

GetBorder returns the Border field if non-nil, zero value otherwise.

### GetBorderOk

`func (o *BTDrawingParams) GetBorderOk() (*bool, bool)`

GetBorderOk returns a tuple with the Border field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorder

`func (o *BTDrawingParams) SetBorder(v bool)`

SetBorder sets Border field to given value.

### HasBorder

`func (o *BTDrawingParams) HasBorder() bool`

HasBorder returns a boolean if a field has been set.

### GetComputeIntersection

`func (o *BTDrawingParams) GetComputeIntersection() bool`

GetComputeIntersection returns the ComputeIntersection field if non-nil, zero value otherwise.

### GetComputeIntersectionOk

`func (o *BTDrawingParams) GetComputeIntersectionOk() (*bool, bool)`

GetComputeIntersectionOk returns a tuple with the ComputeIntersection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeIntersection

`func (o *BTDrawingParams) SetComputeIntersection(v bool)`

SetComputeIntersection sets ComputeIntersection field to given value.

### HasComputeIntersection

`func (o *BTDrawingParams) HasComputeIntersection() bool`

HasComputeIntersection returns a boolean if a field has been set.

### GetDecimalSeparator

`func (o *BTDrawingParams) GetDecimalSeparator() string`

GetDecimalSeparator returns the DecimalSeparator field if non-nil, zero value otherwise.

### GetDecimalSeparatorOk

`func (o *BTDrawingParams) GetDecimalSeparatorOk() (*string, bool)`

GetDecimalSeparatorOk returns a tuple with the DecimalSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalSeparator

`func (o *BTDrawingParams) SetDecimalSeparator(v string)`

SetDecimalSeparator sets DecimalSeparator field to given value.

### HasDecimalSeparator

`func (o *BTDrawingParams) HasDecimalSeparator() bool`

HasDecimalSeparator returns a boolean if a field has been set.

### GetDisplayStateId

`func (o *BTDrawingParams) GetDisplayStateId() string`

GetDisplayStateId returns the DisplayStateId field if non-nil, zero value otherwise.

### GetDisplayStateIdOk

`func (o *BTDrawingParams) GetDisplayStateIdOk() (*string, bool)`

GetDisplayStateIdOk returns a tuple with the DisplayStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStateId

`func (o *BTDrawingParams) SetDisplayStateId(v string)`

SetDisplayStateId sets DisplayStateId field to given value.

### HasDisplayStateId

`func (o *BTDrawingParams) HasDisplayStateId() bool`

HasDisplayStateId returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTDrawingParams) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTDrawingParams) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTDrawingParams) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTDrawingParams) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentMicroversionId

`func (o *BTDrawingParams) GetDocumentMicroversionId() string`

GetDocumentMicroversionId returns the DocumentMicroversionId field if non-nil, zero value otherwise.

### GetDocumentMicroversionIdOk

`func (o *BTDrawingParams) GetDocumentMicroversionIdOk() (*string, bool)`

GetDocumentMicroversionIdOk returns a tuple with the DocumentMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversionId

`func (o *BTDrawingParams) SetDocumentMicroversionId(v string)`

SetDocumentMicroversionId sets DocumentMicroversionId field to given value.

### HasDocumentMicroversionId

`func (o *BTDrawingParams) HasDocumentMicroversionId() bool`

HasDocumentMicroversionId returns a boolean if a field has been set.

### GetDrawingName

`func (o *BTDrawingParams) GetDrawingName() string`

GetDrawingName returns the DrawingName field if non-nil, zero value otherwise.

### GetDrawingNameOk

`func (o *BTDrawingParams) GetDrawingNameOk() (*string, bool)`

GetDrawingNameOk returns a tuple with the DrawingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawingName

`func (o *BTDrawingParams) SetDrawingName(v string)`

SetDrawingName sets DrawingName field to given value.

### HasDrawingName

`func (o *BTDrawingParams) HasDrawingName() bool`

HasDrawingName returns a boolean if a field has been set.

### GetElementConfiguration

`func (o *BTDrawingParams) GetElementConfiguration() string`

GetElementConfiguration returns the ElementConfiguration field if non-nil, zero value otherwise.

### GetElementConfigurationOk

`func (o *BTDrawingParams) GetElementConfigurationOk() (*string, bool)`

GetElementConfigurationOk returns a tuple with the ElementConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementConfiguration

`func (o *BTDrawingParams) SetElementConfiguration(v string)`

SetElementConfiguration sets ElementConfiguration field to given value.

### HasElementConfiguration

`func (o *BTDrawingParams) HasElementConfiguration() bool`

HasElementConfiguration returns a boolean if a field has been set.

### GetElementId

`func (o *BTDrawingParams) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTDrawingParams) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTDrawingParams) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTDrawingParams) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementMicroversionId

`func (o *BTDrawingParams) GetElementMicroversionId() string`

GetElementMicroversionId returns the ElementMicroversionId field if non-nil, zero value otherwise.

### GetElementMicroversionIdOk

`func (o *BTDrawingParams) GetElementMicroversionIdOk() (*string, bool)`

GetElementMicroversionIdOk returns a tuple with the ElementMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementMicroversionId

`func (o *BTDrawingParams) SetElementMicroversionId(v string)`

SetElementMicroversionId sets ElementMicroversionId field to given value.

### HasElementMicroversionId

`func (o *BTDrawingParams) HasElementMicroversionId() bool`

HasElementMicroversionId returns a boolean if a field has been set.

### GetExplosionId

`func (o *BTDrawingParams) GetExplosionId() string`

GetExplosionId returns the ExplosionId field if non-nil, zero value otherwise.

### GetExplosionIdOk

`func (o *BTDrawingParams) GetExplosionIdOk() (*string, bool)`

GetExplosionIdOk returns a tuple with the ExplosionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplosionId

`func (o *BTDrawingParams) SetExplosionId(v string)`

SetExplosionId sets ExplosionId field to given value.

### HasExplosionId

`func (o *BTDrawingParams) HasExplosionId() bool`

HasExplosionId returns a boolean if a field has been set.

### GetExternalDocumentId

`func (o *BTDrawingParams) GetExternalDocumentId() string`

GetExternalDocumentId returns the ExternalDocumentId field if non-nil, zero value otherwise.

### GetExternalDocumentIdOk

`func (o *BTDrawingParams) GetExternalDocumentIdOk() (*string, bool)`

GetExternalDocumentIdOk returns a tuple with the ExternalDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentId

`func (o *BTDrawingParams) SetExternalDocumentId(v string)`

SetExternalDocumentId sets ExternalDocumentId field to given value.

### HasExternalDocumentId

`func (o *BTDrawingParams) HasExternalDocumentId() bool`

HasExternalDocumentId returns a boolean if a field has been set.

### GetExternalDocumentVersionId

`func (o *BTDrawingParams) GetExternalDocumentVersionId() string`

GetExternalDocumentVersionId returns the ExternalDocumentVersionId field if non-nil, zero value otherwise.

### GetExternalDocumentVersionIdOk

`func (o *BTDrawingParams) GetExternalDocumentVersionIdOk() (*string, bool)`

GetExternalDocumentVersionIdOk returns a tuple with the ExternalDocumentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDocumentVersionId

`func (o *BTDrawingParams) SetExternalDocumentVersionId(v string)`

SetExternalDocumentVersionId sets ExternalDocumentVersionId field to given value.

### HasExternalDocumentVersionId

`func (o *BTDrawingParams) HasExternalDocumentVersionId() bool`

HasExternalDocumentVersionId returns a boolean if a field has been set.

### GetHiddenLines

`func (o *BTDrawingParams) GetHiddenLines() BTDrawingHiddenLineOption`

GetHiddenLines returns the HiddenLines field if non-nil, zero value otherwise.

### GetHiddenLinesOk

`func (o *BTDrawingParams) GetHiddenLinesOk() (*BTDrawingHiddenLineOption, bool)`

GetHiddenLinesOk returns a tuple with the HiddenLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenLines

`func (o *BTDrawingParams) SetHiddenLines(v BTDrawingHiddenLineOption)`

SetHiddenLines sets HiddenLines field to given value.

### HasHiddenLines

`func (o *BTDrawingParams) HasHiddenLines() bool`

HasHiddenLines returns a boolean if a field has been set.

### GetIncludeSurfaces

`func (o *BTDrawingParams) GetIncludeSurfaces() bool`

GetIncludeSurfaces returns the IncludeSurfaces field if non-nil, zero value otherwise.

### GetIncludeSurfacesOk

`func (o *BTDrawingParams) GetIncludeSurfacesOk() (*bool, bool)`

GetIncludeSurfacesOk returns a tuple with the IncludeSurfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSurfaces

`func (o *BTDrawingParams) SetIncludeSurfaces(v bool)`

SetIncludeSurfaces sets IncludeSurfaces field to given value.

### HasIncludeSurfaces

`func (o *BTDrawingParams) HasIncludeSurfaces() bool`

HasIncludeSurfaces returns a boolean if a field has been set.

### GetIncludeWires

`func (o *BTDrawingParams) GetIncludeWires() bool`

GetIncludeWires returns the IncludeWires field if non-nil, zero value otherwise.

### GetIncludeWiresOk

`func (o *BTDrawingParams) GetIncludeWiresOk() (*bool, bool)`

GetIncludeWiresOk returns a tuple with the IncludeWires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWires

`func (o *BTDrawingParams) SetIncludeWires(v bool)`

SetIncludeWires sets IncludeWires field to given value.

### HasIncludeWires

`func (o *BTDrawingParams) HasIncludeWires() bool`

HasIncludeWires returns a boolean if a field has been set.

### GetIsFlattenedPart

`func (o *BTDrawingParams) GetIsFlattenedPart() bool`

GetIsFlattenedPart returns the IsFlattenedPart field if non-nil, zero value otherwise.

### GetIsFlattenedPartOk

`func (o *BTDrawingParams) GetIsFlattenedPartOk() (*bool, bool)`

GetIsFlattenedPartOk returns a tuple with the IsFlattenedPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattenedPart

`func (o *BTDrawingParams) SetIsFlattenedPart(v bool)`

SetIsFlattenedPart sets IsFlattenedPart field to given value.

### HasIsFlattenedPart

`func (o *BTDrawingParams) HasIsFlattenedPart() bool`

HasIsFlattenedPart returns a boolean if a field has been set.

### GetIsSketchOnly

`func (o *BTDrawingParams) GetIsSketchOnly() bool`

GetIsSketchOnly returns the IsSketchOnly field if non-nil, zero value otherwise.

### GetIsSketchOnlyOk

`func (o *BTDrawingParams) GetIsSketchOnlyOk() (*bool, bool)`

GetIsSketchOnlyOk returns a tuple with the IsSketchOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSketchOnly

`func (o *BTDrawingParams) SetIsSketchOnly(v bool)`

SetIsSketchOnly sets IsSketchOnly field to given value.

### HasIsSketchOnly

`func (o *BTDrawingParams) HasIsSketchOnly() bool`

HasIsSketchOnly returns a boolean if a field has been set.

### GetIsSurface

`func (o *BTDrawingParams) GetIsSurface() bool`

GetIsSurface returns the IsSurface field if non-nil, zero value otherwise.

### GetIsSurfaceOk

`func (o *BTDrawingParams) GetIsSurfaceOk() (*bool, bool)`

GetIsSurfaceOk returns a tuple with the IsSurface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurface

`func (o *BTDrawingParams) SetIsSurface(v bool)`

SetIsSurface sets IsSurface field to given value.

### HasIsSurface

`func (o *BTDrawingParams) HasIsSurface() bool`

HasIsSurface returns a boolean if a field has been set.

### GetLanguage

`func (o *BTDrawingParams) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BTDrawingParams) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BTDrawingParams) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BTDrawingParams) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLocation

`func (o *BTDrawingParams) GetLocation() BTElementLocationParams`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BTDrawingParams) GetLocationOk() (*BTElementLocationParams, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BTDrawingParams) SetLocation(v BTElementLocationParams)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BTDrawingParams) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModelType

`func (o *BTDrawingParams) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *BTDrawingParams) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *BTDrawingParams) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *BTDrawingParams) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetNamedPositionId

`func (o *BTDrawingParams) GetNamedPositionId() string`

GetNamedPositionId returns the NamedPositionId field if non-nil, zero value otherwise.

### GetNamedPositionIdOk

`func (o *BTDrawingParams) GetNamedPositionIdOk() (*string, bool)`

GetNamedPositionIdOk returns a tuple with the NamedPositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedPositionId

`func (o *BTDrawingParams) SetNamedPositionId(v string)`

SetNamedPositionId sets NamedPositionId field to given value.

### HasNamedPositionId

`func (o *BTDrawingParams) HasNamedPositionId() bool`

HasNamedPositionId returns a boolean if a field has been set.

### GetNumberHorizontalZones

`func (o *BTDrawingParams) GetNumberHorizontalZones() int32`

GetNumberHorizontalZones returns the NumberHorizontalZones field if non-nil, zero value otherwise.

### GetNumberHorizontalZonesOk

`func (o *BTDrawingParams) GetNumberHorizontalZonesOk() (*int32, bool)`

GetNumberHorizontalZonesOk returns a tuple with the NumberHorizontalZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberHorizontalZones

`func (o *BTDrawingParams) SetNumberHorizontalZones(v int32)`

SetNumberHorizontalZones sets NumberHorizontalZones field to given value.

### HasNumberHorizontalZones

`func (o *BTDrawingParams) HasNumberHorizontalZones() bool`

HasNumberHorizontalZones returns a boolean if a field has been set.

### GetNumberVerticalZones

`func (o *BTDrawingParams) GetNumberVerticalZones() int32`

GetNumberVerticalZones returns the NumberVerticalZones field if non-nil, zero value otherwise.

### GetNumberVerticalZonesOk

`func (o *BTDrawingParams) GetNumberVerticalZonesOk() (*int32, bool)`

GetNumberVerticalZonesOk returns a tuple with the NumberVerticalZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberVerticalZones

`func (o *BTDrawingParams) SetNumberVerticalZones(v int32)`

SetNumberVerticalZones sets NumberVerticalZones field to given value.

### HasNumberVerticalZones

`func (o *BTDrawingParams) HasNumberVerticalZones() bool`

HasNumberVerticalZones returns a boolean if a field has been set.

### GetPartId

`func (o *BTDrawingParams) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTDrawingParams) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTDrawingParams) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTDrawingParams) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTDrawingParams) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTDrawingParams) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTDrawingParams) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTDrawingParams) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPartQuery

`func (o *BTDrawingParams) GetPartQuery() string`

GetPartQuery returns the PartQuery field if non-nil, zero value otherwise.

### GetPartQueryOk

`func (o *BTDrawingParams) GetPartQueryOk() (*string, bool)`

GetPartQueryOk returns a tuple with the PartQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartQuery

`func (o *BTDrawingParams) SetPartQuery(v string)`

SetPartQuery sets PartQuery field to given value.

### HasPartQuery

`func (o *BTDrawingParams) HasPartQuery() bool`

HasPartQuery returns a boolean if a field has been set.

### GetProjection

`func (o *BTDrawingParams) GetProjection() string`

GetProjection returns the Projection field if non-nil, zero value otherwise.

### GetProjectionOk

`func (o *BTDrawingParams) GetProjectionOk() (*string, bool)`

GetProjectionOk returns a tuple with the Projection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjection

`func (o *BTDrawingParams) SetProjection(v string)`

SetProjection sets Projection field to given value.

### HasProjection

`func (o *BTDrawingParams) HasProjection() bool`

HasProjection returns a boolean if a field has been set.

### GetPureSketch

`func (o *BTDrawingParams) GetPureSketch() bool`

GetPureSketch returns the PureSketch field if non-nil, zero value otherwise.

### GetPureSketchOk

`func (o *BTDrawingParams) GetPureSketchOk() (*bool, bool)`

GetPureSketchOk returns a tuple with the PureSketch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureSketch

`func (o *BTDrawingParams) SetPureSketch(v bool)`

SetPureSketch sets PureSketch field to given value.

### HasPureSketch

`func (o *BTDrawingParams) HasPureSketch() bool`

HasPureSketch returns a boolean if a field has been set.

### GetQualityOption

`func (o *BTDrawingParams) GetQualityOption() string`

GetQualityOption returns the QualityOption field if non-nil, zero value otherwise.

### GetQualityOptionOk

`func (o *BTDrawingParams) GetQualityOptionOk() (*string, bool)`

GetQualityOptionOk returns a tuple with the QualityOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOption

`func (o *BTDrawingParams) SetQualityOption(v string)`

SetQualityOption sets QualityOption field to given value.

### HasQualityOption

`func (o *BTDrawingParams) HasQualityOption() bool`

HasQualityOption returns a boolean if a field has been set.

### GetReferenceType

`func (o *BTDrawingParams) GetReferenceType() int32`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *BTDrawingParams) GetReferenceTypeOk() (*int32, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *BTDrawingParams) SetReferenceType(v int32)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *BTDrawingParams) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetReferenceTypeEnum

`func (o *BTDrawingParams) GetReferenceTypeEnum() GBTAppElementReferenceType`

GetReferenceTypeEnum returns the ReferenceTypeEnum field if non-nil, zero value otherwise.

### GetReferenceTypeEnumOk

`func (o *BTDrawingParams) GetReferenceTypeEnumOk() (*GBTAppElementReferenceType, bool)`

GetReferenceTypeEnumOk returns a tuple with the ReferenceTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceTypeEnum

`func (o *BTDrawingParams) SetReferenceTypeEnum(v GBTAppElementReferenceType)`

SetReferenceTypeEnum sets ReferenceTypeEnum field to given value.

### HasReferenceTypeEnum

`func (o *BTDrawingParams) HasReferenceTypeEnum() bool`

HasReferenceTypeEnum returns a boolean if a field has been set.

### GetRevision

`func (o *BTDrawingParams) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTDrawingParams) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTDrawingParams) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTDrawingParams) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetShowCutGeomOnly

`func (o *BTDrawingParams) GetShowCutGeomOnly() bool`

GetShowCutGeomOnly returns the ShowCutGeomOnly field if non-nil, zero value otherwise.

### GetShowCutGeomOnlyOk

`func (o *BTDrawingParams) GetShowCutGeomOnlyOk() (*bool, bool)`

GetShowCutGeomOnlyOk returns a tuple with the ShowCutGeomOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCutGeomOnly

`func (o *BTDrawingParams) SetShowCutGeomOnly(v bool)`

SetShowCutGeomOnly sets ShowCutGeomOnly field to given value.

### HasShowCutGeomOnly

`func (o *BTDrawingParams) HasShowCutGeomOnly() bool`

HasShowCutGeomOnly returns a boolean if a field has been set.

### GetSimplificationOption

`func (o *BTDrawingParams) GetSimplificationOption() string`

GetSimplificationOption returns the SimplificationOption field if non-nil, zero value otherwise.

### GetSimplificationOptionOk

`func (o *BTDrawingParams) GetSimplificationOptionOk() (*string, bool)`

GetSimplificationOptionOk returns a tuple with the SimplificationOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplificationOption

`func (o *BTDrawingParams) SetSimplificationOption(v string)`

SetSimplificationOption sets SimplificationOption field to given value.

### HasSimplificationOption

`func (o *BTDrawingParams) HasSimplificationOption() bool`

HasSimplificationOption returns a boolean if a field has been set.

### GetSimplificationThreshold

`func (o *BTDrawingParams) GetSimplificationThreshold() float64`

GetSimplificationThreshold returns the SimplificationThreshold field if non-nil, zero value otherwise.

### GetSimplificationThresholdOk

`func (o *BTDrawingParams) GetSimplificationThresholdOk() (*float64, bool)`

GetSimplificationThresholdOk returns a tuple with the SimplificationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplificationThreshold

`func (o *BTDrawingParams) SetSimplificationThreshold(v float64)`

SetSimplificationThreshold sets SimplificationThreshold field to given value.

### HasSimplificationThreshold

`func (o *BTDrawingParams) HasSimplificationThreshold() bool`

HasSimplificationThreshold returns a boolean if a field has been set.

### GetSize

`func (o *BTDrawingParams) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BTDrawingParams) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BTDrawingParams) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *BTDrawingParams) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSketchIds

`func (o *BTDrawingParams) GetSketchIds() []string`

GetSketchIds returns the SketchIds field if non-nil, zero value otherwise.

### GetSketchIdsOk

`func (o *BTDrawingParams) GetSketchIdsOk() (*[]string, bool)`

GetSketchIdsOk returns a tuple with the SketchIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSketchIds

`func (o *BTDrawingParams) SetSketchIds(v []string)`

SetSketchIds sets SketchIds field to given value.

### HasSketchIds

`func (o *BTDrawingParams) HasSketchIds() bool`

HasSketchIds returns a boolean if a field has been set.

### GetStandard

`func (o *BTDrawingParams) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *BTDrawingParams) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *BTDrawingParams) SetStandard(v string)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *BTDrawingParams) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### GetStartZones

`func (o *BTDrawingParams) GetStartZones() string`

GetStartZones returns the StartZones field if non-nil, zero value otherwise.

### GetStartZonesOk

`func (o *BTDrawingParams) GetStartZonesOk() (*string, bool)`

GetStartZonesOk returns a tuple with the StartZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartZones

`func (o *BTDrawingParams) SetStartZones(v string)`

SetStartZones sets StartZones field to given value.

### HasStartZones

`func (o *BTDrawingParams) HasStartZones() bool`

HasStartZones returns a boolean if a field has been set.

### GetTemplateArgs

`func (o *BTDrawingParams) GetTemplateArgs() []string`

GetTemplateArgs returns the TemplateArgs field if non-nil, zero value otherwise.

### GetTemplateArgsOk

`func (o *BTDrawingParams) GetTemplateArgsOk() (*[]string, bool)`

GetTemplateArgsOk returns a tuple with the TemplateArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateArgs

`func (o *BTDrawingParams) SetTemplateArgs(v []string)`

SetTemplateArgs sets TemplateArgs field to given value.

### HasTemplateArgs

`func (o *BTDrawingParams) HasTemplateArgs() bool`

HasTemplateArgs returns a boolean if a field has been set.

### GetTemplateDocumentId

`func (o *BTDrawingParams) GetTemplateDocumentId() string`

GetTemplateDocumentId returns the TemplateDocumentId field if non-nil, zero value otherwise.

### GetTemplateDocumentIdOk

`func (o *BTDrawingParams) GetTemplateDocumentIdOk() (*string, bool)`

GetTemplateDocumentIdOk returns a tuple with the TemplateDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDocumentId

`func (o *BTDrawingParams) SetTemplateDocumentId(v string)`

SetTemplateDocumentId sets TemplateDocumentId field to given value.

### HasTemplateDocumentId

`func (o *BTDrawingParams) HasTemplateDocumentId() bool`

HasTemplateDocumentId returns a boolean if a field has been set.

### GetTemplateElementId

`func (o *BTDrawingParams) GetTemplateElementId() string`

GetTemplateElementId returns the TemplateElementId field if non-nil, zero value otherwise.

### GetTemplateElementIdOk

`func (o *BTDrawingParams) GetTemplateElementIdOk() (*string, bool)`

GetTemplateElementIdOk returns a tuple with the TemplateElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateElementId

`func (o *BTDrawingParams) SetTemplateElementId(v string)`

SetTemplateElementId sets TemplateElementId field to given value.

### HasTemplateElementId

`func (o *BTDrawingParams) HasTemplateElementId() bool`

HasTemplateElementId returns a boolean if a field has been set.

### GetTemplateName

`func (o *BTDrawingParams) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *BTDrawingParams) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *BTDrawingParams) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *BTDrawingParams) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateVersionId

`func (o *BTDrawingParams) GetTemplateVersionId() string`

GetTemplateVersionId returns the TemplateVersionId field if non-nil, zero value otherwise.

### GetTemplateVersionIdOk

`func (o *BTDrawingParams) GetTemplateVersionIdOk() (*string, bool)`

GetTemplateVersionIdOk returns a tuple with the TemplateVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateVersionId

`func (o *BTDrawingParams) SetTemplateVersionId(v string)`

SetTemplateVersionId sets TemplateVersionId field to given value.

### HasTemplateVersionId

`func (o *BTDrawingParams) HasTemplateVersionId() bool`

HasTemplateVersionId returns a boolean if a field has been set.

### GetTemplateWorkspaceId

`func (o *BTDrawingParams) GetTemplateWorkspaceId() string`

GetTemplateWorkspaceId returns the TemplateWorkspaceId field if non-nil, zero value otherwise.

### GetTemplateWorkspaceIdOk

`func (o *BTDrawingParams) GetTemplateWorkspaceIdOk() (*string, bool)`

GetTemplateWorkspaceIdOk returns a tuple with the TemplateWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateWorkspaceId

`func (o *BTDrawingParams) SetTemplateWorkspaceId(v string)`

SetTemplateWorkspaceId sets TemplateWorkspaceId field to given value.

### HasTemplateWorkspaceId

`func (o *BTDrawingParams) HasTemplateWorkspaceId() bool`

HasTemplateWorkspaceId returns a boolean if a field has been set.

### GetTitleblock

`func (o *BTDrawingParams) GetTitleblock() bool`

GetTitleblock returns the Titleblock field if non-nil, zero value otherwise.

### GetTitleblockOk

`func (o *BTDrawingParams) GetTitleblockOk() (*bool, bool)`

GetTitleblockOk returns a tuple with the Titleblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleblock

`func (o *BTDrawingParams) SetTitleblock(v bool)`

SetTitleblock sets Titleblock field to given value.

### HasTitleblock

`func (o *BTDrawingParams) HasTitleblock() bool`

HasTitleblock returns a boolean if a field has been set.

### GetUnits

`func (o *BTDrawingParams) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *BTDrawingParams) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *BTDrawingParams) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *BTDrawingParams) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetViews

`func (o *BTDrawingParams) GetViews() string`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *BTDrawingParams) GetViewsOk() (*string, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *BTDrawingParams) SetViews(v string)`

SetViews sets Views field to given value.

### HasViews

`func (o *BTDrawingParams) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTDrawingParams) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTDrawingParams) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTDrawingParams) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTDrawingParams) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


