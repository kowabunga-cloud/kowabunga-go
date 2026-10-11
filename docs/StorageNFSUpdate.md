# StorageNFSUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The NFS storage name. | [optional] 
**Description** | Pointer to **string** | The NFS storage description. | [optional] 
**Endpoint** | Pointer to **string** | The associated NFS endpoint FQDN. | [optional] 
**Fs** | Pointer to **string** | The underlying associated Ceph volume name. | [optional] 
**Backends** | Pointer to **[]string** | List of NFS Ganesha API server IP addresses. | [optional] 
**Port** | Pointer to **int32** | NFS Ganesha API server port. | [optional] 

## Methods

### NewStorageNFSUpdate

`func NewStorageNFSUpdate() *StorageNFSUpdate`

NewStorageNFSUpdate instantiates a new StorageNFSUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNFSUpdateWithDefaults

`func NewStorageNFSUpdateWithDefaults() *StorageNFSUpdate`

NewStorageNFSUpdateWithDefaults instantiates a new StorageNFSUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageNFSUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNFSUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNFSUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNFSUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StorageNFSUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageNFSUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageNFSUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageNFSUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *StorageNFSUpdate) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *StorageNFSUpdate) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *StorageNFSUpdate) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *StorageNFSUpdate) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetFs

`func (o *StorageNFSUpdate) GetFs() string`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *StorageNFSUpdate) GetFsOk() (*string, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *StorageNFSUpdate) SetFs(v string)`

SetFs sets Fs field to given value.

### HasFs

`func (o *StorageNFSUpdate) HasFs() bool`

HasFs returns a boolean if a field has been set.

### GetBackends

`func (o *StorageNFSUpdate) GetBackends() []string`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *StorageNFSUpdate) GetBackendsOk() (*[]string, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *StorageNFSUpdate) SetBackends(v []string)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *StorageNFSUpdate) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetPort

`func (o *StorageNFSUpdate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StorageNFSUpdate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StorageNFSUpdate) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *StorageNFSUpdate) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


