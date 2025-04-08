# KawaiiNetIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | **[]string** | The Kawaii global public gateways virtual IP addresses (read-only). | 
**Private** | **[]string** | The Kawaii global private gateways virtual IP addresses (read-only). | 
**Zones** | [**[]KawaiiNetIpZone**](KawaiiNetIpZone.md) | The Kawaii per-zone list of Kowabunga virtual IP addresses. | 

## Methods

### NewKawaiiNetIp

`func NewKawaiiNetIp(public []string, private []string, zones []KawaiiNetIpZone, ) *KawaiiNetIp`

NewKawaiiNetIp instantiates a new KawaiiNetIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiNetIpWithDefaults

`func NewKawaiiNetIpWithDefaults() *KawaiiNetIp`

NewKawaiiNetIpWithDefaults instantiates a new KawaiiNetIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *KawaiiNetIp) GetPublic() []string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *KawaiiNetIp) GetPublicOk() (*[]string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *KawaiiNetIp) SetPublic(v []string)`

SetPublic sets Public field to given value.


### GetPrivate

`func (o *KawaiiNetIp) GetPrivate() []string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *KawaiiNetIp) GetPrivateOk() (*[]string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *KawaiiNetIp) SetPrivate(v []string)`

SetPrivate sets Private field to given value.


### GetZones

`func (o *KawaiiNetIp) GetZones() []KawaiiNetIpZone`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *KawaiiNetIp) GetZonesOk() (*[]KawaiiNetIpZone, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *KawaiiNetIp) SetZones(v []KawaiiNetIpZone)`

SetZones sets Zones field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


