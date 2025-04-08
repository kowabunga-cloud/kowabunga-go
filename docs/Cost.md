# Cost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float32** | The unit price information. | 
**Currency** | **string** | The associated currency. | 

## Methods

### NewCost

`func NewCost(price float32, currency string, ) *Cost`

NewCost instantiates a new Cost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostWithDefaults

`func NewCostWithDefaults() *Cost`

NewCostWithDefaults instantiates a new Cost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *Cost) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Cost) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Cost) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *Cost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Cost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Cost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


