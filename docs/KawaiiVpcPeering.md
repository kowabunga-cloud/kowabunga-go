# KawaiiVpcPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | **string** | Kowabunga Subnet ID to be peered with (subnet local IP addresses will be automatically assigned to Kawaii instances).. | 
**Policy** | Pointer to **string** | The default VPC traffic forwarding policy. | [optional] [default to "drop"]
**Ingress** | Pointer to [**[]KawaiiVpcForwardRule**](KawaiiVpcForwardRule.md) | The firewall list of forwarding ingress rules from VPC peered subnet. ICMP traffic is always accepted. The specified ruleset will be explicitly accepted if drop is the default policy (useless otherwise). | [optional] 
**Egress** | Pointer to [**[]KawaiiVpcForwardRule**](KawaiiVpcForwardRule.md) | The firewall list of forwarding egress rules to VPC peered subnet. ICMP traffic is always accepted. The specified ruleset will be explicitly accepted if drop is the default policy (useless otherwise). | [optional] 
**Netip** | Pointer to [**[]KawaiiVpcNetIpZone**](KawaiiVpcNetIpZone.md) | The per-zone auto-assigned private IPs in peered subnet (read-only). | [optional] 

## Methods

### NewKawaiiVpcPeering

`func NewKawaiiVpcPeering(subnet string, ) *KawaiiVpcPeering`

NewKawaiiVpcPeering instantiates a new KawaiiVpcPeering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiVpcPeeringWithDefaults

`func NewKawaiiVpcPeeringWithDefaults() *KawaiiVpcPeering`

NewKawaiiVpcPeeringWithDefaults instantiates a new KawaiiVpcPeering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *KawaiiVpcPeering) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *KawaiiVpcPeering) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *KawaiiVpcPeering) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.


### GetPolicy

`func (o *KawaiiVpcPeering) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *KawaiiVpcPeering) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *KawaiiVpcPeering) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *KawaiiVpcPeering) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetIngress

`func (o *KawaiiVpcPeering) GetIngress() []KawaiiVpcForwardRule`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *KawaiiVpcPeering) GetIngressOk() (*[]KawaiiVpcForwardRule, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *KawaiiVpcPeering) SetIngress(v []KawaiiVpcForwardRule)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *KawaiiVpcPeering) HasIngress() bool`

HasIngress returns a boolean if a field has been set.

### GetEgress

`func (o *KawaiiVpcPeering) GetEgress() []KawaiiVpcForwardRule`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *KawaiiVpcPeering) GetEgressOk() (*[]KawaiiVpcForwardRule, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *KawaiiVpcPeering) SetEgress(v []KawaiiVpcForwardRule)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *KawaiiVpcPeering) HasEgress() bool`

HasEgress returns a boolean if a field has been set.

### GetNetip

`func (o *KawaiiVpcPeering) GetNetip() []KawaiiVpcNetIpZone`

GetNetip returns the Netip field if non-nil, zero value otherwise.

### GetNetipOk

`func (o *KawaiiVpcPeering) GetNetipOk() (*[]KawaiiVpcNetIpZone, bool)`

GetNetipOk returns a tuple with the Netip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetip

`func (o *KawaiiVpcPeering) SetNetip(v []KawaiiVpcNetIpZone)`

SetNetip sets Netip field to given value.

### HasNetip

`func (o *KawaiiVpcPeering) HasNetip() bool`

HasNetip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


