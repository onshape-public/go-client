# BTLazilyParsedFeatureScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to [**BTUiFeatureStudioChecksum2438**](BTUiFeatureStudioChecksum2438.md) |  | [optional] 
**LanguageVersion** | Pointer to **int32** |  | [optional] 
**Lines** | Pointer to **map[string]interface{}** |  | [optional] 
**Model** | Pointer to [**BTMModel141**](BTMModel141.md) |  | [optional] 
**Module** | Pointer to [**BTPModule234**](BTPModule234.md) |  | [optional] 
**ModuleId** | Pointer to [**BTPModuleId235**](BTPModuleId235.md) |  | [optional] 
**NoticeModuleIds** | Pointer to [**BTPModuleId235**](BTPModuleId235.md) |  | [optional] 
**ParentLanguageVersion** | Pointer to **int32** |  | [optional] 
**References** | Pointer to [**map[string]BTMicroversionId366**](BTMicroversionId366.md) |  | [optional] 
**SizeInKBEstimate** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 

## Methods

### NewBTLazilyParsedFeatureScript

`func NewBTLazilyParsedFeatureScript() *BTLazilyParsedFeatureScript`

NewBTLazilyParsedFeatureScript instantiates a new BTLazilyParsedFeatureScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTLazilyParsedFeatureScriptWithDefaults

`func NewBTLazilyParsedFeatureScriptWithDefaults() *BTLazilyParsedFeatureScript`

NewBTLazilyParsedFeatureScriptWithDefaults instantiates a new BTLazilyParsedFeatureScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *BTLazilyParsedFeatureScript) GetChecksum() BTUiFeatureStudioChecksum2438`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *BTLazilyParsedFeatureScript) GetChecksumOk() (*BTUiFeatureStudioChecksum2438, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *BTLazilyParsedFeatureScript) SetChecksum(v BTUiFeatureStudioChecksum2438)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *BTLazilyParsedFeatureScript) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetLanguageVersion

`func (o *BTLazilyParsedFeatureScript) GetLanguageVersion() int32`

GetLanguageVersion returns the LanguageVersion field if non-nil, zero value otherwise.

### GetLanguageVersionOk

`func (o *BTLazilyParsedFeatureScript) GetLanguageVersionOk() (*int32, bool)`

GetLanguageVersionOk returns a tuple with the LanguageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageVersion

`func (o *BTLazilyParsedFeatureScript) SetLanguageVersion(v int32)`

SetLanguageVersion sets LanguageVersion field to given value.

### HasLanguageVersion

`func (o *BTLazilyParsedFeatureScript) HasLanguageVersion() bool`

HasLanguageVersion returns a boolean if a field has been set.

### GetLines

`func (o *BTLazilyParsedFeatureScript) GetLines() map[string]interface{}`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *BTLazilyParsedFeatureScript) GetLinesOk() (*map[string]interface{}, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *BTLazilyParsedFeatureScript) SetLines(v map[string]interface{})`

SetLines sets Lines field to given value.

### HasLines

`func (o *BTLazilyParsedFeatureScript) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetModel

`func (o *BTLazilyParsedFeatureScript) GetModel() BTMModel141`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BTLazilyParsedFeatureScript) GetModelOk() (*BTMModel141, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BTLazilyParsedFeatureScript) SetModel(v BTMModel141)`

SetModel sets Model field to given value.

### HasModel

`func (o *BTLazilyParsedFeatureScript) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModule

`func (o *BTLazilyParsedFeatureScript) GetModule() BTPModule234`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *BTLazilyParsedFeatureScript) GetModuleOk() (*BTPModule234, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *BTLazilyParsedFeatureScript) SetModule(v BTPModule234)`

SetModule sets Module field to given value.

### HasModule

`func (o *BTLazilyParsedFeatureScript) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetModuleId

`func (o *BTLazilyParsedFeatureScript) GetModuleId() BTPModuleId235`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *BTLazilyParsedFeatureScript) GetModuleIdOk() (*BTPModuleId235, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *BTLazilyParsedFeatureScript) SetModuleId(v BTPModuleId235)`

SetModuleId sets ModuleId field to given value.

### HasModuleId

`func (o *BTLazilyParsedFeatureScript) HasModuleId() bool`

HasModuleId returns a boolean if a field has been set.

### GetNoticeModuleIds

`func (o *BTLazilyParsedFeatureScript) GetNoticeModuleIds() BTPModuleId235`

GetNoticeModuleIds returns the NoticeModuleIds field if non-nil, zero value otherwise.

### GetNoticeModuleIdsOk

`func (o *BTLazilyParsedFeatureScript) GetNoticeModuleIdsOk() (*BTPModuleId235, bool)`

GetNoticeModuleIdsOk returns a tuple with the NoticeModuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeModuleIds

`func (o *BTLazilyParsedFeatureScript) SetNoticeModuleIds(v BTPModuleId235)`

SetNoticeModuleIds sets NoticeModuleIds field to given value.

### HasNoticeModuleIds

`func (o *BTLazilyParsedFeatureScript) HasNoticeModuleIds() bool`

HasNoticeModuleIds returns a boolean if a field has been set.

### GetParentLanguageVersion

`func (o *BTLazilyParsedFeatureScript) GetParentLanguageVersion() int32`

GetParentLanguageVersion returns the ParentLanguageVersion field if non-nil, zero value otherwise.

### GetParentLanguageVersionOk

`func (o *BTLazilyParsedFeatureScript) GetParentLanguageVersionOk() (*int32, bool)`

GetParentLanguageVersionOk returns a tuple with the ParentLanguageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLanguageVersion

`func (o *BTLazilyParsedFeatureScript) SetParentLanguageVersion(v int32)`

SetParentLanguageVersion sets ParentLanguageVersion field to given value.

### HasParentLanguageVersion

`func (o *BTLazilyParsedFeatureScript) HasParentLanguageVersion() bool`

HasParentLanguageVersion returns a boolean if a field has been set.

### GetReferences

`func (o *BTLazilyParsedFeatureScript) GetReferences() map[string]BTMicroversionId366`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *BTLazilyParsedFeatureScript) GetReferencesOk() (*map[string]BTMicroversionId366, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *BTLazilyParsedFeatureScript) SetReferences(v map[string]BTMicroversionId366)`

SetReferences sets References field to given value.

### HasReferences

`func (o *BTLazilyParsedFeatureScript) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSizeInKBEstimate

`func (o *BTLazilyParsedFeatureScript) GetSizeInKBEstimate() int32`

GetSizeInKBEstimate returns the SizeInKBEstimate field if non-nil, zero value otherwise.

### GetSizeInKBEstimateOk

`func (o *BTLazilyParsedFeatureScript) GetSizeInKBEstimateOk() (*int32, bool)`

GetSizeInKBEstimateOk returns a tuple with the SizeInKBEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInKBEstimate

`func (o *BTLazilyParsedFeatureScript) SetSizeInKBEstimate(v int32)`

SetSizeInKBEstimate sets SizeInKBEstimate field to given value.

### HasSizeInKBEstimate

`func (o *BTLazilyParsedFeatureScript) HasSizeInKBEstimate() bool`

HasSizeInKBEstimate returns a boolean if a field has been set.

### GetSource

`func (o *BTLazilyParsedFeatureScript) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BTLazilyParsedFeatureScript) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BTLazilyParsedFeatureScript) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BTLazilyParsedFeatureScript) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


