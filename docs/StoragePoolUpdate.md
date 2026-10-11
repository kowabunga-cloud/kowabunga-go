# StoragePoolUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The storage pool name. | [optional] 
**Description** | Pointer to **string** | The storage pool description. | [optional] 
**Pool** | Pointer to **string** | Ceph pool name. | [optional] 
**CephAddress** | Pointer to **string** | Ceph Monitor(s) address or FQDN. | [optional] 
**CephPort** | Pointer to **int32** | Ceph Monitor(s) port. | [optional] 
**CephSecretUuid** | Pointer to **string** | The libvirt secret UUID for CephX authentication. | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) | Cost associated to the storage pool. | [optional] 
**Agents** | Pointer to **[]string** | a list of existing remote agents managing the storage pool. | [optional] 

## Methods

### NewStoragePoolUpdate

`func NewStoragePoolUpdate() *StoragePoolUpdate`

NewStoragePoolUpdate instantiates a new StoragePoolUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePoolUpdateWithDefaults

`func NewStoragePoolUpdateWithDefaults() *StoragePoolUpdate`

NewStoragePoolUpdateWithDefaults instantiates a new StoragePoolUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StoragePoolUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePoolUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePoolUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePoolUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StoragePoolUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoragePoolUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoragePoolUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoragePoolUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPool

`func (o *StoragePoolUpdate) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *StoragePoolUpdate) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *StoragePoolUpdate) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *StoragePoolUpdate) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetCephAddress

`func (o *StoragePoolUpdate) GetCephAddress() string`

GetCephAddress returns the CephAddress field if non-nil, zero value otherwise.

### GetCephAddressOk

`func (o *StoragePoolUpdate) GetCephAddressOk() (*string, bool)`

GetCephAddressOk returns a tuple with the CephAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephAddress

`func (o *StoragePoolUpdate) SetCephAddress(v string)`

SetCephAddress sets CephAddress field to given value.

### HasCephAddress

`func (o *StoragePoolUpdate) HasCephAddress() bool`

HasCephAddress returns a boolean if a field has been set.

### GetCephPort

`func (o *StoragePoolUpdate) GetCephPort() int32`

GetCephPort returns the CephPort field if non-nil, zero value otherwise.

### GetCephPortOk

`func (o *StoragePoolUpdate) GetCephPortOk() (*int32, bool)`

GetCephPortOk returns a tuple with the CephPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephPort

`func (o *StoragePoolUpdate) SetCephPort(v int32)`

SetCephPort sets CephPort field to given value.

### HasCephPort

`func (o *StoragePoolUpdate) HasCephPort() bool`

HasCephPort returns a boolean if a field has been set.

### GetCephSecretUuid

`func (o *StoragePoolUpdate) GetCephSecretUuid() string`

GetCephSecretUuid returns the CephSecretUuid field if non-nil, zero value otherwise.

### GetCephSecretUuidOk

`func (o *StoragePoolUpdate) GetCephSecretUuidOk() (*string, bool)`

GetCephSecretUuidOk returns a tuple with the CephSecretUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephSecretUuid

`func (o *StoragePoolUpdate) SetCephSecretUuid(v string)`

SetCephSecretUuid sets CephSecretUuid field to given value.

### HasCephSecretUuid

`func (o *StoragePoolUpdate) HasCephSecretUuid() bool`

HasCephSecretUuid returns a boolean if a field has been set.

### GetCost

`func (o *StoragePoolUpdate) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *StoragePoolUpdate) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *StoragePoolUpdate) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *StoragePoolUpdate) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetAgents

`func (o *StoragePoolUpdate) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *StoragePoolUpdate) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *StoragePoolUpdate) SetAgents(v []string)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *StoragePoolUpdate) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


