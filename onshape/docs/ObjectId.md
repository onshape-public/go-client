# ObjectId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **JSONTime** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewObjectId

`func NewObjectId() *ObjectId`

NewObjectId instantiates a new ObjectId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectIdWithDefaults

`func NewObjectIdWithDefaults() *ObjectId`

NewObjectIdWithDefaults instantiates a new ObjectId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ObjectId) GetDate() JSONTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ObjectId) GetDateOk() (*JSONTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ObjectId) SetDate(v JSONTime)`

SetDate sets Date field to given value.

### HasDate

`func (o *ObjectId) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTimestamp

`func (o *ObjectId) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ObjectId) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ObjectId) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ObjectId) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


