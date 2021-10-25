# ListUnspentTransactionOutputsByAddressRISenders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Amount** | **string** | Represents the total amount sent by this address including the fee. | 

## Methods

### NewListUnspentTransactionOutputsByAddressRISenders

`func NewListUnspentTransactionOutputsByAddressRISenders(address string, amount string, ) *ListUnspentTransactionOutputsByAddressRISenders`

NewListUnspentTransactionOutputsByAddressRISenders instantiates a new ListUnspentTransactionOutputsByAddressRISenders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUnspentTransactionOutputsByAddressRISendersWithDefaults

`func NewListUnspentTransactionOutputsByAddressRISendersWithDefaults() *ListUnspentTransactionOutputsByAddressRISenders`

NewListUnspentTransactionOutputsByAddressRISendersWithDefaults instantiates a new ListUnspentTransactionOutputsByAddressRISenders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListUnspentTransactionOutputsByAddressRISenders) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListUnspentTransactionOutputsByAddressRISenders) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListUnspentTransactionOutputsByAddressRISenders) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *ListUnspentTransactionOutputsByAddressRISenders) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListUnspentTransactionOutputsByAddressRISenders) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListUnspentTransactionOutputsByAddressRISenders) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


