# ListAllAssetsFromAllWalletsRINonFungibleTokensInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**Identifier** | **string** | Defines the specific token identifier. For Bitcoin-based transactions it should be the propertyId and for Ethereum-based transactions - the contract. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**Symbol** | **string** | Defines the symbol of the non-fungible tokens. | 
**TokenId** | **string** | Represents tokens&#39; unique identifier. | 
**Type** | **string** | Defines the specific token type. | 

## Methods

### NewListAllAssetsFromAllWalletsRINonFungibleTokensInner

`func NewListAllAssetsFromAllWalletsRINonFungibleTokensInner(blockchain string, identifier string, network string, symbol string, tokenId string, type_ string, ) *ListAllAssetsFromAllWalletsRINonFungibleTokensInner`

NewListAllAssetsFromAllWalletsRINonFungibleTokensInner instantiates a new ListAllAssetsFromAllWalletsRINonFungibleTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllAssetsFromAllWalletsRINonFungibleTokensInnerWithDefaults

`func NewListAllAssetsFromAllWalletsRINonFungibleTokensInnerWithDefaults() *ListAllAssetsFromAllWalletsRINonFungibleTokensInner`

NewListAllAssetsFromAllWalletsRINonFungibleTokensInnerWithDefaults instantiates a new ListAllAssetsFromAllWalletsRINonFungibleTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetIdentifier

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetNetwork

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetSymbol

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenId

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetType

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAllAssetsFromAllWalletsRINonFungibleTokensInner) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


