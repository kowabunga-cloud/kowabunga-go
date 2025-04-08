# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The image template ID (auto-generated). | [optional] 
**Name** | **string** | The image template name. | 
**Description** | Pointer to **string** | The image template description. | [optional] 
**Os** | Pointer to **string** | Type of operating system if OS kind (useful to determine cloud-init parameters for instance). | [optional] [default to "linux"]
**Source** | **string** | HTTP(s) source URL of the KVM-ready OS image. | 

## Methods

### NewTemplate

`func NewTemplate(name string, source string, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Template) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Template) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Template) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Template) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Template) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Template) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Template) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Template) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Template) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOs

`func (o *Template) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Template) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Template) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Template) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetSource

`func (o *Template) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Template) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Template) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


