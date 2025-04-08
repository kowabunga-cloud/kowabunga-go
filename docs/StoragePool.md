# StoragePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The storage pool ID (auto-generated). | [optional] 
**Name** | **string** | The storage pool name. | 
**Description** | Pointer to **string** | The storage pool description. | [optional] 
**Pool** | **string** | Ceph pool name. | 
**CephAddress** | Pointer to **string** | Ceph Monitor(s) address or FQDN. | [optional] [default to "localhost"]
**CephPort** | Pointer to **int64** | Ceph Monitor(s) port (default 3300). | [optional] [default to 3300]
**CephSecretUuid** | Pointer to **string** | The libvirt secret UUID for CephX authentication. | [optional] 
**Cost** | Pointer to [**Cost**](Cost.md) |  | [optional] 
**Agents** | **[]string** | a list of existing remote agents managing the storage pool. | 

## Methods

### NewStoragePool

`func NewStoragePool(name string, pool string, agents []string, ) *StoragePool`

NewStoragePool instantiates a new StoragePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePoolWithDefaults

`func NewStoragePoolWithDefaults() *StoragePool`

NewStoragePoolWithDefaults instantiates a new StoragePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoragePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoragePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoragePool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoragePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StoragePool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StoragePool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoragePool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoragePool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoragePool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPool

`func (o *StoragePool) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *StoragePool) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *StoragePool) SetPool(v string)`

SetPool sets Pool field to given value.


### GetCephAddress

`func (o *StoragePool) GetCephAddress() string`

GetCephAddress returns the CephAddress field if non-nil, zero value otherwise.

### GetCephAddressOk

`func (o *StoragePool) GetCephAddressOk() (*string, bool)`

GetCephAddressOk returns a tuple with the CephAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephAddress

`func (o *StoragePool) SetCephAddress(v string)`

SetCephAddress sets CephAddress field to given value.

### HasCephAddress

`func (o *StoragePool) HasCephAddress() bool`

HasCephAddress returns a boolean if a field has been set.

### GetCephPort

`func (o *StoragePool) GetCephPort() int64`

GetCephPort returns the CephPort field if non-nil, zero value otherwise.

### GetCephPortOk

`func (o *StoragePool) GetCephPortOk() (*int64, bool)`

GetCephPortOk returns a tuple with the CephPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephPort

`func (o *StoragePool) SetCephPort(v int64)`

SetCephPort sets CephPort field to given value.

### HasCephPort

`func (o *StoragePool) HasCephPort() bool`

HasCephPort returns a boolean if a field has been set.

### GetCephSecretUuid

`func (o *StoragePool) GetCephSecretUuid() string`

GetCephSecretUuid returns the CephSecretUuid field if non-nil, zero value otherwise.

### GetCephSecretUuidOk

`func (o *StoragePool) GetCephSecretUuidOk() (*string, bool)`

GetCephSecretUuidOk returns a tuple with the CephSecretUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephSecretUuid

`func (o *StoragePool) SetCephSecretUuid(v string)`

SetCephSecretUuid sets CephSecretUuid field to given value.

### HasCephSecretUuid

`func (o *StoragePool) HasCephSecretUuid() bool`

HasCephSecretUuid returns a boolean if a field has been set.

### GetCost

`func (o *StoragePool) GetCost() Cost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *StoragePool) GetCostOk() (*Cost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *StoragePool) SetCost(v Cost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *StoragePool) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetAgents

`func (o *StoragePool) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *StoragePool) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *StoragePool) SetAgents(v []string)`

SetAgents sets Agents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


