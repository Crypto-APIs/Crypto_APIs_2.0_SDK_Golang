# GetWalletAssetDetailsRIFungibleTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmedAmount** | **string** | Defines the amount of the fungible tokens. | 
**Identifier** | **string** | Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract. | 
**Symbol** | **string** | Defines the symbol of the fungible tokens. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewGetWalletAssetDetailsRIFungibleTokens

`func NewGetWalletAssetDetailsRIFungibleTokens(confirmedAmount string, identifier string, symbol string, type_ string, ) *GetWalletAssetDetailsRIFungibleTokens`

NewGetWalletAssetDetailsRIFungibleTokens instantiates a new GetWalletAssetDetailsRIFungibleTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletAssetDetailsRIFungibleTokensWithDefaults

`func NewGetWalletAssetDetailsRIFungibleTokensWithDefaults() *GetWalletAssetDetailsRIFungibleTokens`

NewGetWalletAssetDetailsRIFungibleTokensWithDefaults instantiates a new GetWalletAssetDetailsRIFungibleTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmedAmount

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetConfirmedAmount() string`

GetConfirmedAmount returns the ConfirmedAmount field if non-nil, zero value otherwise.

### GetConfirmedAmountOk

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetConfirmedAmountOk() (*string, bool)`

GetConfirmedAmountOk returns a tuple with the ConfirmedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAmount

`func (o *GetWalletAssetDetailsRIFungibleTokens) SetConfirmedAmount(v string)`

SetConfirmedAmount sets ConfirmedAmount field to given value.


### GetIdentifier

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GetWalletAssetDetailsRIFungibleTokens) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetSymbol

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetWalletAssetDetailsRIFungibleTokens) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetWalletAssetDetailsRIFungibleTokens) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetWalletAssetDetailsRIFungibleTokens) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


