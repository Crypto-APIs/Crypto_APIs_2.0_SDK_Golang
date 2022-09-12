# ListDepositAddressesRIFungibleTokensInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the fungible tokens. | 
**Identifier** | **string** | Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract. | 
**Name** | **string** | Defines the token&#39;s name as a string. | 
**Symbol** | **string** | Defines the symbol of the fungible tokens. | 
**TokenDecimals** | **int64** | Defines the decimals of the token, i.e. the number of digits that come after the decimal coma of the token. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewListDepositAddressesRIFungibleTokensInner

`func NewListDepositAddressesRIFungibleTokensInner(amount string, identifier string, name string, symbol string, tokenDecimals int64, type_ string, ) *ListDepositAddressesRIFungibleTokensInner`

NewListDepositAddressesRIFungibleTokensInner instantiates a new ListDepositAddressesRIFungibleTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDepositAddressesRIFungibleTokensInnerWithDefaults

`func NewListDepositAddressesRIFungibleTokensInnerWithDefaults() *ListDepositAddressesRIFungibleTokensInner`

NewListDepositAddressesRIFungibleTokensInnerWithDefaults instantiates a new ListDepositAddressesRIFungibleTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListDepositAddressesRIFungibleTokensInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListDepositAddressesRIFungibleTokensInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListDepositAddressesRIFungibleTokensInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIdentifier

`func (o *ListDepositAddressesRIFungibleTokensInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ListDepositAddressesRIFungibleTokensInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ListDepositAddressesRIFungibleTokensInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *ListDepositAddressesRIFungibleTokensInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDepositAddressesRIFungibleTokensInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDepositAddressesRIFungibleTokensInner) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *ListDepositAddressesRIFungibleTokensInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListDepositAddressesRIFungibleTokensInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListDepositAddressesRIFungibleTokensInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenDecimals

`func (o *ListDepositAddressesRIFungibleTokensInner) GetTokenDecimals() int64`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *ListDepositAddressesRIFungibleTokensInner) GetTokenDecimalsOk() (*int64, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *ListDepositAddressesRIFungibleTokensInner) SetTokenDecimals(v int64)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetType

`func (o *ListDepositAddressesRIFungibleTokensInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListDepositAddressesRIFungibleTokensInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListDepositAddressesRIFungibleTokensInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


