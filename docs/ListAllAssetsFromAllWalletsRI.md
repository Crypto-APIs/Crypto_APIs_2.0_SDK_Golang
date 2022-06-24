# ListAllAssetsFromAllWalletsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coins** | [**[]ListAllAssetsFromAllWalletsRICoinsInner**](ListAllAssetsFromAllWalletsRICoinsInner.md) |  | 
**FungibleTokens** | [**[]ListAllAssetsFromAllWalletsRIFungibleTokensInner**](ListAllAssetsFromAllWalletsRIFungibleTokensInner.md) | Represents fungible tokens&#39;es detailed information | 
**NonFungibleTokens** | [**[]ListAllAssetsFromAllWalletsRINonFungibleTokensInner**](ListAllAssetsFromAllWalletsRINonFungibleTokensInner.md) | Represents non-fungible tokens&#39;es detailed information. | 
**WalletId** | **string** | Defines the unique ID of the Wallet. | 
**WalletName** | **string** | Represents the name of the wallet. | 

## Methods

### NewListAllAssetsFromAllWalletsRI

`func NewListAllAssetsFromAllWalletsRI(coins []ListAllAssetsFromAllWalletsRICoinsInner, fungibleTokens []ListAllAssetsFromAllWalletsRIFungibleTokensInner, nonFungibleTokens []ListAllAssetsFromAllWalletsRINonFungibleTokensInner, walletId string, walletName string, ) *ListAllAssetsFromAllWalletsRI`

NewListAllAssetsFromAllWalletsRI instantiates a new ListAllAssetsFromAllWalletsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllAssetsFromAllWalletsRIWithDefaults

`func NewListAllAssetsFromAllWalletsRIWithDefaults() *ListAllAssetsFromAllWalletsRI`

NewListAllAssetsFromAllWalletsRIWithDefaults instantiates a new ListAllAssetsFromAllWalletsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoins

`func (o *ListAllAssetsFromAllWalletsRI) GetCoins() []ListAllAssetsFromAllWalletsRICoinsInner`

GetCoins returns the Coins field if non-nil, zero value otherwise.

### GetCoinsOk

`func (o *ListAllAssetsFromAllWalletsRI) GetCoinsOk() (*[]ListAllAssetsFromAllWalletsRICoinsInner, bool)`

GetCoinsOk returns a tuple with the Coins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoins

`func (o *ListAllAssetsFromAllWalletsRI) SetCoins(v []ListAllAssetsFromAllWalletsRICoinsInner)`

SetCoins sets Coins field to given value.


### GetFungibleTokens

`func (o *ListAllAssetsFromAllWalletsRI) GetFungibleTokens() []ListAllAssetsFromAllWalletsRIFungibleTokensInner`

GetFungibleTokens returns the FungibleTokens field if non-nil, zero value otherwise.

### GetFungibleTokensOk

`func (o *ListAllAssetsFromAllWalletsRI) GetFungibleTokensOk() (*[]ListAllAssetsFromAllWalletsRIFungibleTokensInner, bool)`

GetFungibleTokensOk returns a tuple with the FungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFungibleTokens

`func (o *ListAllAssetsFromAllWalletsRI) SetFungibleTokens(v []ListAllAssetsFromAllWalletsRIFungibleTokensInner)`

SetFungibleTokens sets FungibleTokens field to given value.


### GetNonFungibleTokens

`func (o *ListAllAssetsFromAllWalletsRI) GetNonFungibleTokens() []ListAllAssetsFromAllWalletsRINonFungibleTokensInner`

GetNonFungibleTokens returns the NonFungibleTokens field if non-nil, zero value otherwise.

### GetNonFungibleTokensOk

`func (o *ListAllAssetsFromAllWalletsRI) GetNonFungibleTokensOk() (*[]ListAllAssetsFromAllWalletsRINonFungibleTokensInner, bool)`

GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFungibleTokens

`func (o *ListAllAssetsFromAllWalletsRI) SetNonFungibleTokens(v []ListAllAssetsFromAllWalletsRINonFungibleTokensInner)`

SetNonFungibleTokens sets NonFungibleTokens field to given value.


### GetWalletId

`func (o *ListAllAssetsFromAllWalletsRI) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ListAllAssetsFromAllWalletsRI) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ListAllAssetsFromAllWalletsRI) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetWalletName

`func (o *ListAllAssetsFromAllWalletsRI) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *ListAllAssetsFromAllWalletsRI) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *ListAllAssetsFromAllWalletsRI) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


