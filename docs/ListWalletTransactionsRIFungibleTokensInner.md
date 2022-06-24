# ListWalletTransactionsRIFungibleTokensInner

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

### NewListWalletTransactionsRIFungibleTokensInner

`func NewListWalletTransactionsRIFungibleTokensInner(amount string, convertedAmount string, exchangeRateUnit string, name string, recipient string, sender string, symbol string, tokenDecimals int32, type_ string, ) *ListWalletTransactionsRIFungibleTokensInner`

NewListWalletTransactionsRIFungibleTokensInner instantiates a new ListWalletTransactionsRIFungibleTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRIFungibleTokensInnerWithDefaults

`func NewListWalletTransactionsRIFungibleTokensInnerWithDefaults() *ListWalletTransactionsRIFungibleTokensInner`

NewListWalletTransactionsRIFungibleTokensInnerWithDefaults instantiates a new ListWalletTransactionsRIFungibleTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetConvertedAmount

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetConvertedAmount() string`

GetConvertedAmount returns the ConvertedAmount field if non-nil, zero value otherwise.

### GetConvertedAmountOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetConvertedAmountOk() (*string, bool)`

GetConvertedAmountOk returns a tuple with the ConvertedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAmount

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetConvertedAmount(v string)`

SetConvertedAmount sets ConvertedAmount field to given value.


### GetExchangeRateUnit

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetExchangeRateUnit() string`

GetExchangeRateUnit returns the ExchangeRateUnit field if non-nil, zero value otherwise.

### GetExchangeRateUnitOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetExchangeRateUnitOk() (*string, bool)`

GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUnit

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetExchangeRateUnit(v string)`

SetExchangeRateUnit sets ExchangeRateUnit field to given value.


### GetName

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetName(v string)`

SetName sets Name field to given value.


### GetRecipient

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSymbol

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenDecimals

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetTokenDecimals() int32`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetTokenDecimalsOk() (*int32, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetTokenDecimals(v int32)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetType

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListWalletTransactionsRIFungibleTokensInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListWalletTransactionsRIFungibleTokensInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


