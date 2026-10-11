# KawaiiIpSecFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingress** | Pointer to [**[]KawaiiFirewallIngressRule**](KawaiiFirewallIngressRule.md) | The Kawaii IPsec connection firewall list of ingress rules. Default policy is to drop all incoming traffic from the remote peer, including ICMP. Specified ruleset will be explicitly accepted. | [optional] 

## Methods

### NewKawaiiIpSecFirewall

`func NewKawaiiIpSecFirewall() *KawaiiIpSecFirewall`

NewKawaiiIpSecFirewall instantiates a new KawaiiIpSecFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiIpSecFirewallWithDefaults

`func NewKawaiiIpSecFirewallWithDefaults() *KawaiiIpSecFirewall`

NewKawaiiIpSecFirewallWithDefaults instantiates a new KawaiiIpSecFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngress

`func (o *KawaiiIpSecFirewall) GetIngress() []KawaiiFirewallIngressRule`

GetIngress returns the Ingress field if non-nil, zero value otherwise.

### GetIngressOk

`func (o *KawaiiIpSecFirewall) GetIngressOk() (*[]KawaiiFirewallIngressRule, bool)`

GetIngressOk returns a tuple with the Ingress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngress

`func (o *KawaiiIpSecFirewall) SetIngress(v []KawaiiFirewallIngressRule)`

SetIngress sets Ingress field to given value.

### HasIngress

`func (o *KawaiiIpSecFirewall) HasIngress() bool`

HasIngress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


