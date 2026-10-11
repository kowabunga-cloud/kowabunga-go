# KawaiiUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kawaii name, auto-generated from the project name. | [optional] 
**Description** | Pointer to **string** | The Kawaii description. | [optional] 
**Firewall** | Pointer to [**KawaiiFirewall**](KawaiiFirewall.md) | The Kawaii firewall settings from/to public Internet. | [optional] 
**Dnat** | Pointer to [**[]KawaiiDNatRule**](KawaiiDNatRule.md) | The Kawaii list of NAT forwarding entries. Kawaii will forward public Internet traffic from all public virtual IPs to requested private subnet IP addresses. | [optional] 

## Methods

### NewKawaiiUpdate

`func NewKawaiiUpdate() *KawaiiUpdate`

NewKawaiiUpdate instantiates a new KawaiiUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiUpdateWithDefaults

`func NewKawaiiUpdateWithDefaults() *KawaiiUpdate`

NewKawaiiUpdateWithDefaults instantiates a new KawaiiUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KawaiiUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KawaiiUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KawaiiUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KawaiiUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KawaiiUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KawaiiUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KawaiiUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KawaiiUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirewall

`func (o *KawaiiUpdate) GetFirewall() KawaiiFirewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *KawaiiUpdate) GetFirewallOk() (*KawaiiFirewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *KawaiiUpdate) SetFirewall(v KawaiiFirewall)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *KawaiiUpdate) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetDnat

`func (o *KawaiiUpdate) GetDnat() []KawaiiDNatRule`

GetDnat returns the Dnat field if non-nil, zero value otherwise.

### GetDnatOk

`func (o *KawaiiUpdate) GetDnatOk() (*[]KawaiiDNatRule, bool)`

GetDnatOk returns a tuple with the Dnat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnat

`func (o *KawaiiUpdate) SetDnat(v []KawaiiDNatRule)`

SetDnat sets Dnat field to given value.

### HasDnat

`func (o *KawaiiUpdate) HasDnat() bool`

HasDnat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


