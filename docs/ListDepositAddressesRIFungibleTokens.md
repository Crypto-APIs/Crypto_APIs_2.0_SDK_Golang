# ListDepositAddressesRIFungibleTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the fungible tokens. | 
**Identifier** | **string** | Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract. | 
**Name** | **string** | Defines the token&#39;s name as a string. | 
**Symbol** | **string** | Defines the symbol of the fungible tokens. | 
**TokenDecimals** | **int32** | Defines the decimals of the token, i.e. the number of digits that come after the decimal coma of the token. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewListDepositAddressesRIFungibleTokens

`func NewListDepositAddressesRIFungibleTokens(amount string, identifier string, name string, symbol string, tokenDecimals int32, type_ string, ) *ListDepositAddressesRIFungibleTokens`

NewListDepositAddressesRIFungibleTokens instantiates a new ListDepositAddressesRIFungibleTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDepositAddressesRIFungibleTokensWithDefaults

`func NewListDepositAddressesRIFungibleTokensWithDefaults() *ListDepositAddressesRIFungibleTokens`

NewListDepositAddressesRIFungibleTokensWithDefaults instantiates a new ListDepositAddressesRIFungibleTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListDepositAddressesRIFungibleTokens) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListDepositAddressesRIFungibleTokens) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListDepositAddressesRIFungibleTokens) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIdentifier

`func (o *ListDepositAddressesRIFungibleTokens) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ListDepositAddressesRIFungibleTokens) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ListDepositAddressesRIFungibleTokens) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *ListDepositAddressesRIFungibleTokens) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDepositAddressesRIFungibleTokens) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDepositAddressesRIFungibleTokens) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *ListDepositAddressesRIFungibleTokens) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListDepositAddressesRIFungibleTokens) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListDepositAddressesRIFungibleTokens) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenDecimals

`func (o *ListDepositAddressesRIFungibleTokens) GetTokenDecimals() int32`

GetTokenDecimals returns the TokenDecimals field if non-nil, zero value otherwise.

### GetTokenDecimalsOk

`func (o *ListDepositAddressesRIFungibleTokens) GetTokenDecimalsOk() (*int32, bool)`

GetTokenDecimalsOk returns a tuple with the TokenDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenDecimals

`func (o *ListDepositAddressesRIFungibleTokens) SetTokenDecimals(v int32)`

SetTokenDecimals sets TokenDecimals field to given value.


### GetType

`func (o *ListDepositAddressesRIFungibleTokens) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListDepositAddressesRIFungibleTokens) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListDepositAddressesRIFungibleTokens) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


