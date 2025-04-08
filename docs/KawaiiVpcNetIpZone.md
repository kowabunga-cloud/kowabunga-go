# KawaiiVpcNetIpZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | The Kawaii zone name (read-only). | 
**Private** | **string** | The Kawaii zone gateway private IP address in VPC peered subnet  (read-only). | 

## Methods

### NewKawaiiVpcNetIpZone

`func NewKawaiiVpcNetIpZone(zone string, private string, ) *KawaiiVpcNetIpZone`

NewKawaiiVpcNetIpZone instantiates a new KawaiiVpcNetIpZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiVpcNetIpZoneWithDefaults

`func NewKawaiiVpcNetIpZoneWithDefaults() *KawaiiVpcNetIpZone`

NewKawaiiVpcNetIpZoneWithDefaults instantiates a new KawaiiVpcNetIpZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *KawaiiVpcNetIpZone) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *KawaiiVpcNetIpZone) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *KawaiiVpcNetIpZone) SetZone(v string)`

SetZone sets Zone field to given value.


### GetPrivate

`func (o *KawaiiVpcNetIpZone) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *KawaiiVpcNetIpZone) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *KawaiiVpcNetIpZone) SetPrivate(v string)`

SetPrivate sets Private field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


