# AddressTokensTransactionUnconfirmedErc20

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Specifies the name of the token. | 
**Symbol** | **string** | Specifies an identifier of the token, where up to five alphanumeric characters can be used for it. | 
**Decimals** | Pointer to **string** | Defines how many decimals can be used to break the token. | [optional] 
**Amount** | **string** | Defines the amount of tokens sent with the transaction that is pending confirmation. | 
**ContractAddress** | **string** | Defines the address of the contract. | 

## Methods

### NewAddressTokensTransactionUnconfirmedErc20

`func NewAddressTokensTransactionUnconfirmedErc20(name string, symbol string, amount string, contractAddress string, ) *AddressTokensTransactionUnconfirmedErc20`

NewAddressTokensTransactionUnconfirmedErc20 instantiates a new AddressTokensTransactionUnconfirmedErc20 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTokensTransactionUnconfirmedErc20WithDefaults

`func NewAddressTokensTransactionUnconfirmedErc20WithDefaults() *AddressTokensTransactionUnconfirmedErc20`

NewAddressTokensTransactionUnconfirmedErc20WithDefaults instantiates a new AddressTokensTransactionUnconfirmedErc20 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddressTokensTransactionUnconfirmedErc20) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressTokensTransactionUnconfirmedErc20) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressTokensTransactionUnconfirmedErc20) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *AddressTokensTransactionUnconfirmedErc20) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AddressTokensTransactionUnconfirmedErc20) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AddressTokensTransactionUnconfirmedErc20) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *AddressTokensTransactionUnconfirmedErc20) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *AddressTokensTransactionUnconfirmedErc20) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *AddressTokensTransactionUnconfirmedErc20) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *AddressTokensTransactionUnconfirmedErc20) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetAmount

`func (o *AddressTokensTransactionUnconfirmedErc20) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTokensTransactionUnconfirmedErc20) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTokensTransactionUnconfirmedErc20) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetContractAddress

`func (o *AddressTokensTransactionUnconfirmedErc20) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *AddressTokensTransactionUnconfirmedErc20) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *AddressTokensTransactionUnconfirmedErc20) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


