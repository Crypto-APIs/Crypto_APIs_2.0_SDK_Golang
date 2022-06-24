# ListWalletTransactionsRINonFungibleTokensInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvertedAmount** | **string** | Defines the tokens&#39; converted amount value. | 
**ExchangeRateUnit** | **string** | Represents token&#39;s exchange rate unit. | 
**Name** | **string** | Defines the token&#39;s name as a string. | 
**Recipient** | **string** | Defines the address to which the recipient receives the transferred tokens. | 
**Sender** | **string** | Defines the address from which the sender transfers tokens. | 
**Symbol** | **string** | Defines the symbol of the non-fungible tokens. | 
**TokenId** | **string** | Represents tokens&#39; unique identifier. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewListWalletTransactionsRINonFungibleTokensInner

`func NewListWalletTransactionsRINonFungibleTokensInner(convertedAmount string, exchangeRateUnit string, name string, recipient string, sender string, symbol string, tokenId string, type_ string, ) *ListWalletTransactionsRINonFungibleTokensInner`

NewListWalletTransactionsRINonFungibleTokensInner instantiates a new ListWalletTransactionsRINonFungibleTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWalletTransactionsRINonFungibleTokensInnerWithDefaults

`func NewListWalletTransactionsRINonFungibleTokensInnerWithDefaults() *ListWalletTransactionsRINonFungibleTokensInner`

NewListWalletTransactionsRINonFungibleTokensInnerWithDefaults instantiates a new ListWalletTransactionsRINonFungibleTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvertedAmount

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetConvertedAmount() string`

GetConvertedAmount returns the ConvertedAmount field if non-nil, zero value otherwise.

### GetConvertedAmountOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetConvertedAmountOk() (*string, bool)`

GetConvertedAmountOk returns a tuple with the ConvertedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedAmount

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetConvertedAmount(v string)`

SetConvertedAmount sets ConvertedAmount field to given value.


### GetExchangeRateUnit

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetExchangeRateUnit() string`

GetExchangeRateUnit returns the ExchangeRateUnit field if non-nil, zero value otherwise.

### GetExchangeRateUnitOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetExchangeRateUnitOk() (*string, bool)`

GetExchangeRateUnitOk returns a tuple with the ExchangeRateUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateUnit

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetExchangeRateUnit(v string)`

SetExchangeRateUnit sets ExchangeRateUnit field to given value.


### GetName

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetName(v string)`

SetName sets Name field to given value.


### GetRecipient

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSender

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSymbol

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenId

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetType

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListWalletTransactionsRINonFungibleTokensInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListWalletTransactionsRINonFungibleTokensInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


