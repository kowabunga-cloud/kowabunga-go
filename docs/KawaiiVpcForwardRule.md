# KawaiiVpcForwardRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | The transport layer protocol to forward public traffic to. | [optional] [default to "tcp"]
**Ports** | **string** | The port (or list of ports) to forward public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005). | 

## Methods

### NewKawaiiVpcForwardRule

`func NewKawaiiVpcForwardRule(ports string, ) *KawaiiVpcForwardRule`

NewKawaiiVpcForwardRule instantiates a new KawaiiVpcForwardRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiVpcForwardRuleWithDefaults

`func NewKawaiiVpcForwardRuleWithDefaults() *KawaiiVpcForwardRule`

NewKawaiiVpcForwardRuleWithDefaults instantiates a new KawaiiVpcForwardRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *KawaiiVpcForwardRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KawaiiVpcForwardRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KawaiiVpcForwardRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KawaiiVpcForwardRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPorts

`func (o *KawaiiVpcForwardRule) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KawaiiVpcForwardRule) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KawaiiVpcForwardRule) SetPorts(v string)`

SetPorts sets Ports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


