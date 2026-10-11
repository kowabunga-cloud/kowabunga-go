# RegionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The region name. | [optional] 
**Description** | Pointer to **string** | The region description. | [optional] 
**Domain** | Pointer to **string** | Region domain name (e.g. myregion.kowabunga.acme.com). | [optional] 

## Methods

### NewRegionUpdate

`func NewRegionUpdate() *RegionUpdate`

NewRegionUpdate instantiates a new RegionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionUpdateWithDefaults

`func NewRegionUpdateWithDefaults() *RegionUpdate`

NewRegionUpdateWithDefaults instantiates a new RegionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegionUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegionUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RegionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *RegionUpdate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RegionUpdate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RegionUpdate) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RegionUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


