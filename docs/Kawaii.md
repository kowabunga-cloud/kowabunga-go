# Kawaii

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kawaii ID (auto-generated). | [optional] 
**Name** | Pointer to **string** | The Kawaii name. | [optional] 
**Description** | Pointer to **string** | The Kawaii description. | [optional] 
**Netip** | Pointer to [**KawaiiNetIp**](KawaiiNetIp.md) |  | [optional] 
**Firewall** | Pointer to [**KawaiiFirewall**](KawaiiFirewall.md) |  | [optional] 
**Dnat** | Pointer to [**[]KawaiiDNatRule**](KawaiiDNatRule.md) | The Kawaii list of NAT forwarding entries. Kawaii will forward public Internet traffic from all public virtual IPs to requested private subnet IP addresses. | [optional] 
**VpcPeerings** | Pointer to [**[]KawaiiVpcPeering**](KawaiiVpcPeering.md) | The Kawaii list of Kowabunga private VPC subnet peering entries. | [optional] 
**IpsecConnections** | Pointer to [**[]KawaiiIpSec**](KawaiiIpSec.md) | The Kawaii list of Kowabunga IPsec connections. | [optional] 

## Methods

### NewKawaii

`func NewKawaii() *Kawaii`

NewKawaii instantiates a new Kawaii object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiWithDefaults

`func NewKawaiiWithDefaults() *Kawaii`

NewKawaiiWithDefaults instantiates a new Kawaii object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Kawaii) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Kawaii) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Kawaii) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Kawaii) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Kawaii) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Kawaii) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Kawaii) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Kawaii) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Kawaii) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Kawaii) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Kawaii) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Kawaii) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNetip

`func (o *Kawaii) GetNetip() KawaiiNetIp`

GetNetip returns the Netip field if non-nil, zero value otherwise.

### GetNetipOk

`func (o *Kawaii) GetNetipOk() (*KawaiiNetIp, bool)`

GetNetipOk returns a tuple with the Netip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetip

`func (o *Kawaii) SetNetip(v KawaiiNetIp)`

SetNetip sets Netip field to given value.

### HasNetip

`func (o *Kawaii) HasNetip() bool`

HasNetip returns a boolean if a field has been set.

### GetFirewall

`func (o *Kawaii) GetFirewall() KawaiiFirewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *Kawaii) GetFirewallOk() (*KawaiiFirewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *Kawaii) SetFirewall(v KawaiiFirewall)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *Kawaii) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetDnat

`func (o *Kawaii) GetDnat() []KawaiiDNatRule`

GetDnat returns the Dnat field if non-nil, zero value otherwise.

### GetDnatOk

`func (o *Kawaii) GetDnatOk() (*[]KawaiiDNatRule, bool)`

GetDnatOk returns a tuple with the Dnat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnat

`func (o *Kawaii) SetDnat(v []KawaiiDNatRule)`

SetDnat sets Dnat field to given value.

### HasDnat

`func (o *Kawaii) HasDnat() bool`

HasDnat returns a boolean if a field has been set.

### GetVpcPeerings

`func (o *Kawaii) GetVpcPeerings() []KawaiiVpcPeering`

GetVpcPeerings returns the VpcPeerings field if non-nil, zero value otherwise.

### GetVpcPeeringsOk

`func (o *Kawaii) GetVpcPeeringsOk() (*[]KawaiiVpcPeering, bool)`

GetVpcPeeringsOk returns a tuple with the VpcPeerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeerings

`func (o *Kawaii) SetVpcPeerings(v []KawaiiVpcPeering)`

SetVpcPeerings sets VpcPeerings field to given value.

### HasVpcPeerings

`func (o *Kawaii) HasVpcPeerings() bool`

HasVpcPeerings returns a boolean if a field has been set.

### GetIpsecConnections

`func (o *Kawaii) GetIpsecConnections() []KawaiiIpSec`

GetIpsecConnections returns the IpsecConnections field if non-nil, zero value otherwise.

### GetIpsecConnectionsOk

`func (o *Kawaii) GetIpsecConnectionsOk() (*[]KawaiiIpSec, bool)`

GetIpsecConnectionsOk returns a tuple with the IpsecConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecConnections

`func (o *Kawaii) SetIpsecConnections(v []KawaiiIpSec)`

SetIpsecConnections sets IpsecConnections field to given value.

### HasIpsecConnections

`func (o *Kawaii) HasIpsecConnections() bool`

HasIpsecConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


