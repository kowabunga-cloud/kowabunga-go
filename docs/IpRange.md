# IpRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**First** | **string** | The range&#39;s first IP address. | 
**Last** | **string** | The range&#39;s last IP address. | 

## Methods

### NewIpRange

`func NewIpRange(first string, last string, ) *IpRange`

NewIpRange instantiates a new IpRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRangeWithDefaults

`func NewIpRangeWithDefaults() *IpRange`

NewIpRangeWithDefaults instantiates a new IpRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirst

`func (o *IpRange) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *IpRange) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *IpRange) SetFirst(v string)`

SetFirst sets First field to given value.


### GetLast

`func (o *IpRange) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *IpRange) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *IpRange) SetLast(v string)`

SetLast sets Last field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


