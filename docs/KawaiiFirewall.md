# KawaiiFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingress** | Pointer to [**[]KawaiiFirewallIngressRule**](KawaiiFirewallIngressRule.md) | The Kawaii public firewall list of ingress rules. Kawaii default policy is to drop all incoming traffic, including ICMP. Specified ruleset will be explicitly accepted. | [optional] 
**EgressPolicy** | Pointer to **string** | The default public traffic egress policy. | [optional] [default to "accept"]
**Egress** | Pointer to [**[]KawaiiFirewallEgressRule**](KawaiiFirewallEgressRule.md) | The Kawaii public firewall list of egress rules. Kawaii default policy is to accept all outgoing traffic, including ICMP. Specified ruleset will be explicitly dropped if egress_policy is set to accept, and explicitly accepted if egress policy is set to drop.. | [optional] 

## Methods

### NewKawaiiFirewall

`func NewKawaiiFirewall() *KawaiiFirewall`

NewKawaiiFirewall instantiates a new KawaiiFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiFirewallWithDefaults

`func NewKawaiiFirewallWithDefaults() *KawaiiFirewall`

NewKawaiiFirewallWithDefaults instantiates a new KawaiiFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngress

`func (o *KawaiiFirewall) GetIngress() []KawaiiFirewallIngressRule`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *KawaiiFirewall) GetIngressOk() (*[]KawaiiFirewallIngressRule, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *KawaiiFirewall) SetIngress(v []KawaiiFirewallIngressRule)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *KawaiiFirewall) HasIngress() bool`

HasIngress returns a boolean if a field has been set.

### GetEgressPolicy

`func (o *KawaiiFirewall) GetEgressPolicy() string`

GetEgressPolicy returns the EgressPolicy field if non-nil, zero value otherwise.

### GetEgressPolicyOk

`func (o *KawaiiFirewall) GetEgressPolicyOk() (*string, bool)`

GetEgressPolicyOk returns a tuple with the EgressPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressPolicy

`func (o *KawaiiFirewall) SetEgressPolicy(v string)`

SetEgressPolicy sets EgressPolicy field to given value.

### HasEgressPolicy

`func (o *KawaiiFirewall) HasEgressPolicy() bool`

HasEgressPolicy returns a boolean if a field has been set.

### GetEgress

`func (o *KawaiiFirewall) GetEgress() []KawaiiFirewallEgressRule`

GetEgress returns the Egress field if non-nil, zero value otherwise.

### GetEgressOk

`func (o *KawaiiFirewall) GetEgressOk() (*[]KawaiiFirewallEgressRule, bool)`

GetEgressOk returns a tuple with the Egress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgress

`func (o *KawaiiFirewall) SetEgress(v []KawaiiFirewallEgressRule)`

SetEgress sets Egress field to given value.

### HasEgress

`func (o *KawaiiFirewall) HasEgress() bool`

HasEgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


