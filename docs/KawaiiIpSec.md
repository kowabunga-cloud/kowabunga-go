# KawaiiIpSec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Kawaii IPsec connection ID (auto-generated). | [optional] 
**Name** | **string** | The Kawaii IPsec connection name. | 
**Description** | Pointer to **string** | The Kawaii IPsec connection description. | [optional] 
**Ip** | Pointer to **string** | The Kawaii IPsec connection IPSec IP. | [optional] 
**RemoteIp** | **string** | The Kawaii IPsec connection remote peer VPN Gateway. | 
**RemoteSubnet** | **string** | The Kawaii IPsec connection remote subnet. | 
**PreSharedKey** | **string** | The Kawaii IPsec connection pre-shared key(PSK). | 
**DpdTimeoutAction** | Pointer to **string** | The Kawaii IPsec connection Dead Peer Detection Action (clear,restart or trap). | [optional] [default to "restart"]
**DpdTimeout** | Pointer to **string** | The Kawaii IPsec connection Dead Peer Detection Timeout. | [optional] [default to "240s"]
**StartAction** | Pointer to **string** | The Kawaii IPsec connection start action (none, start, trap). | [optional] [default to "start"]
**RekeyTime** | Pointer to **string** | The Kawaii IPsec connection rekey time. Default is 2h. | [optional] [default to "2h"]
**Phase1Lifetime** | Pointer to **string** | The Kawaii IPsec connection Lifetime for phase 1 negociation. Default is 1h. | [optional] [default to "1h"]
**Phase1DhGroupNumber** | **int64** | The Kawaii IPsec connection phase 1 Diffie Hellman IANA algorithm. | 
**Phase1IntegrityAlgorithm** | **string** | The Kawaii IPsec connection phase 1 integrity algorithm.. | 
**Phase1EncryptionAlgorithm** | **string** | The Kawaii IPsec connection phase 1 encryption algorithm.. | 
**Phase2Lifetime** | Pointer to **string** | The Kawaii IPsec connection Lifetime for phase 2 negociation. Default is 1h. | [optional] [default to "1h"]
**Phase2DhGroupNumber** | **int64** | The Kawaii IPsec connection phase 2 Diffie Hellman IANA algorithm. | 
**Phase2IntegrityAlgorithm** | **string** | The Kawaii IPsec connection phase 2 integrity algorithm.. | 
**Phase2EncryptionAlgorithm** | **string** | The Kawaii IPsec connection phase 2 encryption algorithm.. | 
**Firewall** | Pointer to [**KawaiiFirewall**](KawaiiFirewall.md) |  | [optional] 

## Methods

### NewKawaiiIpSec

`func NewKawaiiIpSec(name string, remoteIp string, remoteSubnet string, preSharedKey string, phase1DhGroupNumber int64, phase1IntegrityAlgorithm string, phase1EncryptionAlgorithm string, phase2DhGroupNumber int64, phase2IntegrityAlgorithm string, phase2EncryptionAlgorithm string, ) *KawaiiIpSec`

NewKawaiiIpSec instantiates a new KawaiiIpSec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKawaiiIpSecWithDefaults

`func NewKawaiiIpSecWithDefaults() *KawaiiIpSec`

NewKawaiiIpSecWithDefaults instantiates a new KawaiiIpSec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KawaiiIpSec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KawaiiIpSec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KawaiiIpSec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KawaiiIpSec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KawaiiIpSec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KawaiiIpSec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KawaiiIpSec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KawaiiIpSec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KawaiiIpSec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KawaiiIpSec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KawaiiIpSec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIp

`func (o *KawaiiIpSec) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *KawaiiIpSec) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *KawaiiIpSec) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *KawaiiIpSec) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetRemoteIp

`func (o *KawaiiIpSec) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *KawaiiIpSec) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *KawaiiIpSec) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetRemoteSubnet

`func (o *KawaiiIpSec) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *KawaiiIpSec) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *KawaiiIpSec) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.


### GetPreSharedKey

`func (o *KawaiiIpSec) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *KawaiiIpSec) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *KawaiiIpSec) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.


### GetDpdTimeoutAction

`func (o *KawaiiIpSec) GetDpdTimeoutAction() string`

GetDpdTimeoutAction returns the DpdTimeoutAction field if non-nil, zero value otherwise.

### GetDpdTimeoutActionOk

`func (o *KawaiiIpSec) GetDpdTimeoutActionOk() (*string, bool)`

GetDpdTimeoutActionOk returns a tuple with the DpdTimeoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeoutAction

`func (o *KawaiiIpSec) SetDpdTimeoutAction(v string)`

SetDpdTimeoutAction sets DpdTimeoutAction field to given value.

### HasDpdTimeoutAction

`func (o *KawaiiIpSec) HasDpdTimeoutAction() bool`

HasDpdTimeoutAction returns a boolean if a field has been set.

### GetDpdTimeout

`func (o *KawaiiIpSec) GetDpdTimeout() string`

GetDpdTimeout returns the DpdTimeout field if non-nil, zero value otherwise.

### GetDpdTimeoutOk

`func (o *KawaiiIpSec) GetDpdTimeoutOk() (*string, bool)`

GetDpdTimeoutOk returns a tuple with the DpdTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdTimeout

`func (o *KawaiiIpSec) SetDpdTimeout(v string)`

SetDpdTimeout sets DpdTimeout field to given value.

### HasDpdTimeout

`func (o *KawaiiIpSec) HasDpdTimeout() bool`

HasDpdTimeout returns a boolean if a field has been set.

### GetStartAction

`func (o *KawaiiIpSec) GetStartAction() string`

GetStartAction returns the StartAction field if non-nil, zero value otherwise.

### GetStartActionOk

`func (o *KawaiiIpSec) GetStartActionOk() (*string, bool)`

GetStartActionOk returns a tuple with the StartAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAction

`func (o *KawaiiIpSec) SetStartAction(v string)`

SetStartAction sets StartAction field to given value.

### HasStartAction

`func (o *KawaiiIpSec) HasStartAction() bool`

HasStartAction returns a boolean if a field has been set.

### GetRekeyTime

`func (o *KawaiiIpSec) GetRekeyTime() string`

GetRekeyTime returns the RekeyTime field if non-nil, zero value otherwise.

### GetRekeyTimeOk

`func (o *KawaiiIpSec) GetRekeyTimeOk() (*string, bool)`

GetRekeyTimeOk returns a tuple with the RekeyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRekeyTime

`func (o *KawaiiIpSec) SetRekeyTime(v string)`

SetRekeyTime sets RekeyTime field to given value.

### HasRekeyTime

`func (o *KawaiiIpSec) HasRekeyTime() bool`

HasRekeyTime returns a boolean if a field has been set.

### GetPhase1Lifetime

`func (o *KawaiiIpSec) GetPhase1Lifetime() string`

GetPhase1Lifetime returns the Phase1Lifetime field if non-nil, zero value otherwise.

### GetPhase1LifetimeOk

`func (o *KawaiiIpSec) GetPhase1LifetimeOk() (*string, bool)`

GetPhase1LifetimeOk returns a tuple with the Phase1Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Lifetime

`func (o *KawaiiIpSec) SetPhase1Lifetime(v string)`

SetPhase1Lifetime sets Phase1Lifetime field to given value.

### HasPhase1Lifetime

`func (o *KawaiiIpSec) HasPhase1Lifetime() bool`

HasPhase1Lifetime returns a boolean if a field has been set.

### GetPhase1DhGroupNumber

`func (o *KawaiiIpSec) GetPhase1DhGroupNumber() int64`

GetPhase1DhGroupNumber returns the Phase1DhGroupNumber field if non-nil, zero value otherwise.

### GetPhase1DhGroupNumberOk

`func (o *KawaiiIpSec) GetPhase1DhGroupNumberOk() (*int64, bool)`

GetPhase1DhGroupNumberOk returns a tuple with the Phase1DhGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1DhGroupNumber

`func (o *KawaiiIpSec) SetPhase1DhGroupNumber(v int64)`

SetPhase1DhGroupNumber sets Phase1DhGroupNumber field to given value.


### GetPhase1IntegrityAlgorithm

`func (o *KawaiiIpSec) GetPhase1IntegrityAlgorithm() string`

GetPhase1IntegrityAlgorithm returns the Phase1IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetPhase1IntegrityAlgorithmOk

`func (o *KawaiiIpSec) GetPhase1IntegrityAlgorithmOk() (*string, bool)`

GetPhase1IntegrityAlgorithmOk returns a tuple with the Phase1IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1IntegrityAlgorithm

`func (o *KawaiiIpSec) SetPhase1IntegrityAlgorithm(v string)`

SetPhase1IntegrityAlgorithm sets Phase1IntegrityAlgorithm field to given value.


### GetPhase1EncryptionAlgorithm

`func (o *KawaiiIpSec) GetPhase1EncryptionAlgorithm() string`

GetPhase1EncryptionAlgorithm returns the Phase1EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetPhase1EncryptionAlgorithmOk

`func (o *KawaiiIpSec) GetPhase1EncryptionAlgorithmOk() (*string, bool)`

GetPhase1EncryptionAlgorithmOk returns a tuple with the Phase1EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1EncryptionAlgorithm

`func (o *KawaiiIpSec) SetPhase1EncryptionAlgorithm(v string)`

SetPhase1EncryptionAlgorithm sets Phase1EncryptionAlgorithm field to given value.


### GetPhase2Lifetime

`func (o *KawaiiIpSec) GetPhase2Lifetime() string`

GetPhase2Lifetime returns the Phase2Lifetime field if non-nil, zero value otherwise.

### GetPhase2LifetimeOk

`func (o *KawaiiIpSec) GetPhase2LifetimeOk() (*string, bool)`

GetPhase2LifetimeOk returns a tuple with the Phase2Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Lifetime

`func (o *KawaiiIpSec) SetPhase2Lifetime(v string)`

SetPhase2Lifetime sets Phase2Lifetime field to given value.

### HasPhase2Lifetime

`func (o *KawaiiIpSec) HasPhase2Lifetime() bool`

HasPhase2Lifetime returns a boolean if a field has been set.

### GetPhase2DhGroupNumber

`func (o *KawaiiIpSec) GetPhase2DhGroupNumber() int64`

GetPhase2DhGroupNumber returns the Phase2DhGroupNumber field if non-nil, zero value otherwise.

### GetPhase2DhGroupNumberOk

`func (o *KawaiiIpSec) GetPhase2DhGroupNumberOk() (*int64, bool)`

GetPhase2DhGroupNumberOk returns a tuple with the Phase2DhGroupNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2DhGroupNumber

`func (o *KawaiiIpSec) SetPhase2DhGroupNumber(v int64)`

SetPhase2DhGroupNumber sets Phase2DhGroupNumber field to given value.


### GetPhase2IntegrityAlgorithm

`func (o *KawaiiIpSec) GetPhase2IntegrityAlgorithm() string`

GetPhase2IntegrityAlgorithm returns the Phase2IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetPhase2IntegrityAlgorithmOk

`func (o *KawaiiIpSec) GetPhase2IntegrityAlgorithmOk() (*string, bool)`

GetPhase2IntegrityAlgorithmOk returns a tuple with the Phase2IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2IntegrityAlgorithm

`func (o *KawaiiIpSec) SetPhase2IntegrityAlgorithm(v string)`

SetPhase2IntegrityAlgorithm sets Phase2IntegrityAlgorithm field to given value.


### GetPhase2EncryptionAlgorithm

`func (o *KawaiiIpSec) GetPhase2EncryptionAlgorithm() string`

GetPhase2EncryptionAlgorithm returns the Phase2EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetPhase2EncryptionAlgorithmOk

`func (o *KawaiiIpSec) GetPhase2EncryptionAlgorithmOk() (*string, bool)`

GetPhase2EncryptionAlgorithmOk returns a tuple with the Phase2EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2EncryptionAlgorithm

`func (o *KawaiiIpSec) SetPhase2EncryptionAlgorithm(v string)`

SetPhase2EncryptionAlgorithm sets Phase2EncryptionAlgorithm field to given value.


### GetFirewall

`func (o *KawaiiIpSec) GetFirewall() KawaiiFirewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *KawaiiIpSec) GetFirewallOk() (*KawaiiFirewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *KawaiiIpSec) SetFirewall(v KawaiiFirewall)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *KawaiiIpSec) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


