# ListWalletTransactionsRIFungibleTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the fungible tokens. | 
**ConvertedAmount** | **string** | Defines the tokens&#39; converted amount value. | 
**ExchangeRateUnit** | **string** | Represents token&#39;s exchange rate unit. | 
**Name** | **string** | Defines the token&#39;s name as a string. | 
**Recipient** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one | 
**Sender** | **string** | Represents the address which sends this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one sender. | 
**Symbol** | **string** | Defines the symbol of the fungible tokens. | 
**TokenDecimals** | **int32** | Defines the decimals of the token, i.e. the number of digits that come after the decimal coma of the token. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewListWalletTransactionsRIFungibleTokens

`func NewListWalletTransactionsRIFungibleTokens(amount string, convertedAmount string, exchangeRateUnit string, name string, recipient string, sender string, symbol string, tokenDecimals int32, type_ string, ) *ListWalletTransactionsRIFungibleTokens`

NewListWalletTransactionsRIFungibleTokens instantiates a new ListWalletTransactionsRIFungibleTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIFungibleTokensWithDefaults

`func NewListWalletTransactionsRIFungibleTokensWithDefaults() *ListWalletTransactionsRIFungibleTokens`

NewListWalletTransactionsRIFungibleTokensWithDefaults instantiates a new ListWalletTransactionsRIFungibleTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListWalletTransactionsRIFungibleTokens) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRIFungibleTokens) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetConvertedAmount

`func (o *ListWalletTransactionsRIFungibleTokens) GetConvertedAmount() string`

GetConvertedAmount returns the ConvertedAmount field if non-nil, zero value otherwise.

### GetConvertedAmountOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetConvertedAmountOk() (*string, bool)`

GetConvertedAmountOk returns a tuple with the ConvertedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAmount

`func (o *ListWalletTransactionsRIFungibleTokens) SetConvertedAmount(v string)`

SetConvertedAmount sets ConvertedAmount field to given value.


### GetExchangeRateUnit

`func (o *ListWalletTransactionsRIFungibleTokens) GetExchangeRateUnit() string`

GetExchangeRateUnit returns the ExchangeRateUnit field if non-nil, zero value otherwise.

### GetExchangeRateUnitOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetExchangeRateUnitOk() (*string, bool)`

GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUnit

`func (o *ListWalletTransactionsRIFungibleTokens) SetExchangeRateUnit(v string)`

SetExchangeRateUnit sets ExchangeRateUnit field to given value.


### GetName

`func (o *ListWalletTransactionsRIFungibleTokens) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWalletTransactionsRIFungibleTokens) SetName(v string)`

SetName sets Name field to given value.


### GetRecipient

`func (o *ListWalletTransactionsRIFungibleTokens) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ListWalletTransactionsRIFungibleTokens) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *ListWalletTransactionsRIFungibleTokens) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListWalletTransactionsRIFungibleTokens) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSymbol

`func (o *ListWalletTransactionsRIFungibleTokens) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListWalletTransactionsRIFungibleTokens) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenDecimals

`func (o *ListWalletTransactionsRIFungibleTokens) GetTokenDecimals() int32`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetTokenDecimalsOk() (*int32, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *ListWalletTransactionsRIFungibleTokens) SetTokenDecimals(v int32)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetType

`func (o *ListWalletTransactionsRIFungibleTokens) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListWalletTransactionsRIFungibleTokens) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListWalletTransactionsRIFungibleTokens) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


