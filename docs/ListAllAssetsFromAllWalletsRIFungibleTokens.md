# ListAllAssetsFromAllWalletsRIFungibleTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the fungible tokens. | 
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Identifier** | **string** | Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**Symbol** | **string** | Defines the symbol of the fungible tokens. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewListAllAssetsFromAllWalletsRIFungibleTokens

`func NewListAllAssetsFromAllWalletsRIFungibleTokens(amount string, blockchain string, identifier string, network string, symbol string, type_ string, ) *ListAllAssetsFromAllWalletsRIFungibleTokens`

NewListAllAssetsFromAllWalletsRIFungibleTokens instantiates a new ListAllAssetsFromAllWalletsRIFungibleTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllAssetsFromAllWalletsRIFungibleTokensWithDefaults

`func NewListAllAssetsFromAllWalletsRIFungibleTokensWithDefaults() *ListAllAssetsFromAllWalletsRIFungibleTokens`

NewListAllAssetsFromAllWalletsRIFungibleTokensWithDefaults instantiates a new ListAllAssetsFromAllWalletsRIFungibleTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBlockchain

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetIdentifier

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetNetwork

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetSymbol

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAllAssetsFromAllWalletsRIFungibleTokens) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


