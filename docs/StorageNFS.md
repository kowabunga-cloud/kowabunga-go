# StorageNFS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The NFS storage ID (auto-generated). | [optional] 
**Name** | **string** | The NFS storage name. | 
**Description** | Pointer to **string** | The NFS storage description. | [optional] 
**Endpoint** | **string** | The associated NFS endpoint FQDN. | 
**Fs** | Pointer to **string** | The underlying associated Ceph volume name. | [optional] [default to "nfs"]
**Backends** | Pointer to **[]string** | List of NFS Ganesha API server IP addresses. | [optional] 
**Port** | Pointer to **int64** | NFS Ganesha API server port (default 54934). | [optional] [default to 54934]

## Methods

### NewStorageNFS

`func NewStorageNFS(name string, endpoint string, ) *StorageNFS`

NewStorageNFS instantiates a new StorageNFS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNFSWithDefaults

`func NewStorageNFSWithDefaults() *StorageNFS`

NewStorageNFSWithDefaults instantiates a new StorageNFS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageNFS) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageNFS) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageNFS) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StorageNFS) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StorageNFS) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNFS) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNFS) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StorageNFS) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageNFS) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageNFS) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageNFS) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *StorageNFS) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *StorageNFS) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *StorageNFS) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetFs

`func (o *StorageNFS) GetFs() string`

GetFs returns the Fs field if non-nil, zero value otherwise.

### GetFsOk

`func (o *StorageNFS) GetFsOk() (*string, bool)`

GetFsOk returns a tuple with the Fs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFs

`func (o *StorageNFS) SetFs(v string)`

SetFs sets Fs field to given value.

### HasFs

`func (o *StorageNFS) HasFs() bool`

HasFs returns a boolean if a field has been set.

### GetBackends

`func (o *StorageNFS) GetBackends() []string`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *StorageNFS) GetBackendsOk() (*[]string, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *StorageNFS) SetBackends(v []string)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *StorageNFS) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetPort

`func (o *StorageNFS) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StorageNFS) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StorageNFS) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *StorageNFS) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


