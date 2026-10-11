# KawaiiIpSecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The Kawaii IPsec connection name. | [optional] 
**Description** | Pointer to **string** | The Kawaii IPsec connection description. | [optional] 
**RemoteIp** | Pointer to **string** | The Kawaii IPsec connection remote peer VPN gateway IP address. | [optional] 
**RemoteSubnet** | Pointer to **string** | The Kawaii IPsec connection remote subnet. | [optional] 
**PreSharedKey** | Pointer to **string** | The Kawaii IPsec connection pre-shared key (PSK). | [optional] 
**DpdTimeoutAction** | Pointer to **string** | The Kawaii IPsec connection Dead Peer Detection (DPD) timeout action. | [optional] 
**DpdTimeout** | Pointer to **string** | The Kawaii IPsec connection Dead Peer Detection (DPD) timeout. | [optional] 
**StartAction** | Pointer to **string** | The Kawaii IPsec connection start action. | [optional] 
**RekeyTime** | Pointer to **string** | The Kawaii IPsec connection rekey time. | [optional] 
**Phase1Lifetime** | Pointer to **string** | The Kawaii IPsec connection lifetime for phase 1 negotiation. | [optional] 
**Phase1DhGroupNumber** | Pointer to **int32** | The Kawaii IPsec connection phase 1 Diffie Hellman IANA algorithm. | [optional] 
**Phase1IntegrityAlgorithm** | Pointer to **string** | The Kawaii IPsec connection phase 1 integrity algorithm. | [optional] 
**Phase1EncryptionAlgorithm** | Pointer to **string** | The Kawaii IPsec connection phase 1 encryption algorithm. | [optional] 
**Phase2Lifetime** | Pointer to **string** | The Kawaii IPsec connection lifetime for phase 2 negotiation. | [optional] 
**Phase2DhGroupNumber** | Pointer to **int32** | The Kawaii IPsec connection phase 2 Diffie Hellman IANA algorithm. | [optional] 
**Phase2IntegrityAlgorithm** | Pointer to **string** | The Kawaii IPsec connection phase 2 integrity algorithm. | [optional] 
**Phase2EncryptionAlgorithm** | Pointer to **string** | The Kawaii IPsec connection phase 2 encryption algorithm. | [optional] 
**Firewall** | Pointer to [**KawaiiIpSecFirewall**](KawaiiIpSecFirewall.md) | The Kawaii IPsec connection firewall settings on the IPsec interface. | [optional] 

## Methods

### NewKawaiiIpSecUpdate

`func NewKawaiiIpSecUpdate() *KawaiiIpSecUpdate`

NewKawaiiIpSecUpdate instantiates a new KawaiiIpSecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiIpSecUpdateWithDefaults

`func NewKawaiiIpSecUpdateWithDefaults() *KawaiiIpSecUpdate`

NewKawaiiIpSecUpdateWithDefaults instantiates a new KawaiiIpSecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KawaiiIpSecUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KawaiiIpSecUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KawaiiIpSecUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KawaiiIpSecUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *KawaiiIpSecUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KawaiiIpSecUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KawaiiIpSecUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KawaiiIpSecUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRemoteIp

`func (o *KawaiiIpSecUpdate) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *KawaiiIpSecUpdate) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *KawaiiIpSecUpdate) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.

### HasRemoteIp

`func (o *KawaiiIpSecUpdate) HasRemoteIp() bool`

HasRemoteIp returns a boolean if a field has been set.

### GetRemoteSubnet

`func (o *KawaiiIpSecUpdate) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *KawaiiIpSecUpdate) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *KawaiiIpSecUpdate) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.

### HasRemoteSubnet

`func (o *KawaiiIpSecUpdate) HasRemoteSubnet() bool`

HasRemoteSubnet returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *KawaiiIpSecUpdate) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *KawaiiIpSecUpdate) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *KawaiiIpSecUpdate) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *KawaiiIpSecUpdate) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetDpdTimeoutAction

`func (o *KawaiiIpSecUpdate) GetDpdTimeoutAction() string`

GetDpdTimeoutAction returns the DpdTimeoutAction field if non-nil, zero value otherwise.

### GetDpdTimeoutActionOk

`func (o *KawaiiIpSecUpdate) GetDpdTimeoutActionOk() (*string, bool)`

GetDpdTimeoutActionOk returns a tuple with the DpdTimeoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeoutAction

`func (o *KawaiiIpSecUpdate) SetDpdTimeoutAction(v string)`

SetDpdTimeoutAction sets DpdTimeoutAction field to given value.

### HasDpdTimeoutAction

`func (o *KawaiiIpSecUpdate) HasDpdTimeoutAction() bool`

HasDpdTimeoutAction returns a boolean if a field has been set.

### GetDpdTimeout

`func (o *KawaiiIpSecUpdate) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *KawaiiIpSecUpdate) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *KawaiiIpSecUpdate) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.

### HasDpdTimeout

`func (o *KawaiiIpSecUpdate) HasDpdTimeout() bool`

HasDpdTimeout returns a boolean if a field has been set.

### GetStartAction

`func (o *KawaiiIpSecUpdate) GetStartAction() string`

GetStartAction returns the StartAction field if non-nil, zero value otherwise.

### GetStartActionOk

`func (o *KawaiiIpSecUpdate) GetStartActionOk() (*string, bool)`

GetStartActionOk returns a tuple with the StartAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAction

`func (o *KawaiiIpSecUpdate) SetStartAction(v string)`

SetStartAction sets StartAction field to given value.

### HasStartAction

`func (o *KawaiiIpSecUpdate) HasStartAction() bool`

HasStartAction returns a boolean if a field has been set.

### GetRekeyTime

`func (o *KawaiiIpSecUpdate) GetRekeyTime() string`

GetRekeyTime returns the RekeyTime field if non-nil, zero value otherwise.

### GetRekeyTimeOk

`func (o *KawaiiIpSecUpdate) GetRekeyTimeOk() (*string, bool)`

GetRekeyTimeOk returns a tuple with the RekeyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyTime

`func (o *KawaiiIpSecUpdate) SetRekeyTime(v string)`

SetRekeyTime sets RekeyTime field to given value.

### HasRekeyTime

`func (o *KawaiiIpSecUpdate) HasRekeyTime() bool`

HasRekeyTime returns a boolean if a field has been set.

### GetPhase1Lifetime

`func (o *KawaiiIpSecUpdate) GetPhase1Lifetime() string`

GetPhase1Lifetime returns the Phase1Lifetime field if non-nil, zero value otherwise.

### GetPhase1LifetimeOk

`func (o *KawaiiIpSecUpdate) GetPhase1LifetimeOk() (*string, bool)`

GetPhase1LifetimeOk returns a tuple with the Phase1Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Lifetime

`func (o *KawaiiIpSecUpdate) SetPhase1Lifetime(v string)`

SetPhase1Lifetime sets Phase1Lifetime field to given value.

### HasPhase1Lifetime

`func (o *KawaiiIpSecUpdate) HasPhase1Lifetime() bool`

HasPhase1Lifetime returns a boolean if a field has been set.

### GetPhase1DhGroupNumber

`func (o *KawaiiIpSecUpdate) GetPhase1DhGroupNumber() int32`

GetPhase1DhGroupNumber returns the Phase1DhGroupNumber field if non-nil, zero value otherwise.

### GetPhase1DhGroupNumberOk

`func (o *KawaiiIpSecUpdate) GetPhase1DhGroupNumberOk() (*int32, bool)`

GetPhase1DhGroupNumberOk returns a tuple with the Phase1DhGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DhGroupNumber

`func (o *KawaiiIpSecUpdate) SetPhase1DhGroupNumber(v int32)`

SetPhase1DhGroupNumber sets Phase1DhGroupNumber field to given value.

### HasPhase1DhGroupNumber

`func (o *KawaiiIpSecUpdate) HasPhase1DhGroupNumber() bool`

HasPhase1DhGroupNumber returns a boolean if a field has been set.

### GetPhase1IntegrityAlgorithm

`func (o *KawaiiIpSecUpdate) GetPhase1IntegrityAlgorithm() string`

GetPhase1IntegrityAlgorithm returns the Phase1IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetPhase1IntegrityAlgorithmOk

`func (o *KawaiiIpSecUpdate) GetPhase1IntegrityAlgorithmOk() (*string, bool)`

GetPhase1IntegrityAlgorithmOk returns a tuple with the Phase1IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1IntegrityAlgorithm

`func (o *KawaiiIpSecUpdate) SetPhase1IntegrityAlgorithm(v string)`

SetPhase1IntegrityAlgorithm sets Phase1IntegrityAlgorithm field to given value.

### HasPhase1IntegrityAlgorithm

`func (o *KawaiiIpSecUpdate) HasPhase1IntegrityAlgorithm() bool`

HasPhase1IntegrityAlgorithm returns a boolean if a field has been set.

### GetPhase1EncryptionAlgorithm

`func (o *KawaiiIpSecUpdate) GetPhase1EncryptionAlgorithm() string`

GetPhase1EncryptionAlgorithm returns the Phase1EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetPhase1EncryptionAlgorithmOk

`func (o *KawaiiIpSecUpdate) GetPhase1EncryptionAlgorithmOk() (*string, bool)`

GetPhase1EncryptionAlgorithmOk returns a tuple with the Phase1EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1EncryptionAlgorithm

`func (o *KawaiiIpSecUpdate) SetPhase1EncryptionAlgorithm(v string)`

SetPhase1EncryptionAlgorithm sets Phase1EncryptionAlgorithm field to given value.

### HasPhase1EncryptionAlgorithm

`func (o *KawaiiIpSecUpdate) HasPhase1EncryptionAlgorithm() bool`

HasPhase1EncryptionAlgorithm returns a boolean if a field has been set.

### GetPhase2Lifetime

`func (o *KawaiiIpSecUpdate) GetPhase2Lifetime() string`

GetPhase2Lifetime returns the Phase2Lifetime field if non-nil, zero value otherwise.

### GetPhase2LifetimeOk

`func (o *KawaiiIpSecUpdate) GetPhase2LifetimeOk() (*string, bool)`

GetPhase2LifetimeOk returns a tuple with the Phase2Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Lifetime

`func (o *KawaiiIpSecUpdate) SetPhase2Lifetime(v string)`

SetPhase2Lifetime sets Phase2Lifetime field to given value.

### HasPhase2Lifetime

`func (o *KawaiiIpSecUpdate) HasPhase2Lifetime() bool`

HasPhase2Lifetime returns a boolean if a field has been set.

### GetPhase2DhGroupNumber

`func (o *KawaiiIpSecUpdate) GetPhase2DhGroupNumber() int32`

GetPhase2DhGroupNumber returns the Phase2DhGroupNumber field if non-nil, zero value otherwise.

### GetPhase2DhGroupNumberOk

`func (o *KawaiiIpSecUpdate) GetPhase2DhGroupNumberOk() (*int32, bool)`

GetPhase2DhGroupNumberOk returns a tuple with the Phase2DhGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2DhGroupNumber

`func (o *KawaiiIpSecUpdate) SetPhase2DhGroupNumber(v int32)`

SetPhase2DhGroupNumber sets Phase2DhGroupNumber field to given value.

### HasPhase2DhGroupNumber

`func (o *KawaiiIpSecUpdate) HasPhase2DhGroupNumber() bool`

HasPhase2DhGroupNumber returns a boolean if a field has been set.

### GetPhase2IntegrityAlgorithm

`func (o *KawaiiIpSecUpdate) GetPhase2IntegrityAlgorithm() string`

GetPhase2IntegrityAlgorithm returns the Phase2IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetPhase2IntegrityAlgorithmOk

`func (o *KawaiiIpSecUpdate) GetPhase2IntegrityAlgorithmOk() (*string, bool)`

GetPhase2IntegrityAlgorithmOk returns a tuple with the Phase2IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2IntegrityAlgorithm

`func (o *KawaiiIpSecUpdate) SetPhase2IntegrityAlgorithm(v string)`

SetPhase2IntegrityAlgorithm sets Phase2IntegrityAlgorithm field to given value.

### HasPhase2IntegrityAlgorithm

`func (o *KawaiiIpSecUpdate) HasPhase2IntegrityAlgorithm() bool`

HasPhase2IntegrityAlgorithm returns a boolean if a field has been set.

### GetPhase2EncryptionAlgorithm

`func (o *KawaiiIpSecUpdate) GetPhase2EncryptionAlgorithm() string`

GetPhase2EncryptionAlgorithm returns the Phase2EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetPhase2EncryptionAlgorithmOk

`func (o *KawaiiIpSecUpdate) GetPhase2EncryptionAlgorithmOk() (*string, bool)`

GetPhase2EncryptionAlgorithmOk returns a tuple with the Phase2EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2EncryptionAlgorithm

`func (o *KawaiiIpSecUpdate) SetPhase2EncryptionAlgorithm(v string)`

SetPhase2EncryptionAlgorithm sets Phase2EncryptionAlgorithm field to given value.

### HasPhase2EncryptionAlgorithm

`func (o *KawaiiIpSecUpdate) HasPhase2EncryptionAlgorithm() bool`

HasPhase2EncryptionAlgorithm returns a boolean if a field has been set.

### GetFirewall

`func (o *KawaiiIpSecUpdate) GetFirewall() KawaiiIpSecFirewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *KawaiiIpSecUpdate) GetFirewallOk() (*KawaiiIpSecFirewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *KawaiiIpSecUpdate) SetFirewall(v KawaiiIpSecFirewall)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *KawaiiIpSecUpdate) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


