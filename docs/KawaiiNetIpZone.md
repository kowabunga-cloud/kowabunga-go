# KawaiiNetIpZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **string** | The Kawaii zone name (read-only). | 
**Public** | **string** | The Kawaii zone gateway public virtual IP (read-only). | 
**Private** | **string** | The Kawaii zone gateway private virtual IP (read-only). | 

## Methods

### NewKawaiiNetIpZone

`func NewKawaiiNetIpZone(zone string, public string, private string, ) *KawaiiNetIpZone`

NewKawaiiNetIpZone instantiates a new KawaiiNetIpZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiNetIpZoneWithDefaults

`func NewKawaiiNetIpZoneWithDefaults() *KawaiiNetIpZone`

NewKawaiiNetIpZoneWithDefaults instantiates a new KawaiiNetIpZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *KawaiiNetIpZone) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *KawaiiNetIpZone) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *KawaiiNetIpZone) SetZone(v string)`

SetZone sets Zone field to given value.


### GetPublic

`func (o *KawaiiNetIpZone) GetPublic() string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *KawaiiNetIpZone) GetPublicOk() (*string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *KawaiiNetIpZone) SetPublic(v string)`

SetPublic sets Public field to given value.


### GetPrivate

`func (o *KawaiiNetIpZone) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *KawaiiNetIpZone) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *KawaiiNetIpZone) SetPrivate(v string)`

SetPrivate sets Private field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


