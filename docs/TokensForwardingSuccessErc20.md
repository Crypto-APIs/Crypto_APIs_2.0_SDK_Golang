# TokensForwardingSuccessErc20

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Specifies the name of the token. | 
**Symbol** | **string** | Specifies an identifier of the token, where up to five alphanumeric characters can be used for it. | 
**Decimals** | Pointer to **string** | Defines how many decimals can be used to break the token. | [optional] 
**Amount** | **string** | Defines the amount of tokens sent with the confirmed transaction. | 
**ContractAddress** | **string** | Defines the address of the contract. | 

## Methods

### NewTokensForwardingSuccessErc20

`func NewTokensForwardingSuccessErc20(name string, symbol string, amount string, contractAddress string, ) *TokensForwardingSuccessErc20`

NewTokensForwardingSuccessErc20 instantiates a new TokensForwardingSuccessErc20 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensForwardingSuccessErc20WithDefaults

`func NewTokensForwardingSuccessErc20WithDefaults() *TokensForwardingSuccessErc20`

NewTokensForwardingSuccessErc20WithDefaults instantiates a new TokensForwardingSuccessErc20 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TokensForwardingSuccessErc20) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokensForwardingSuccessErc20) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokensForwardingSuccessErc20) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokensForwardingSuccessErc20) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokensForwardingSuccessErc20) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokensForwardingSuccessErc20) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *TokensForwardingSuccessErc20) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokensForwardingSuccessErc20) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokensForwardingSuccessErc20) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokensForwardingSuccessErc20) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetAmount

`func (o *TokensForwardingSuccessErc20) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokensForwardingSuccessErc20) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokensForwardingSuccessErc20) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetContractAddress

`func (o *TokensForwardingSuccessErc20) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TokensForwardingSuccessErc20) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TokensForwardingSuccessErc20) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


