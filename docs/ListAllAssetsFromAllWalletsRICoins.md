# ListAllAssetsFromAllWalletsRICoins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**ConfirmedBalance** | **string** | Defines the total balance of the address that is confirmed. It doesn&#39;t include unconfirmed transactions. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**TotalReceived** | **string** | Defines the total amount of all coins received to the address, based on confirmed transactions. | 
**TotalSpent** | **string** | Defines the total amount of all spent by this address coins, based on confirmed transactions. | 
**Unit** | **string** | Represents the unit of the confirmed balance. | 

## Methods

### NewListAllAssetsFromAllWalletsRICoins

`func NewListAllAssetsFromAllWalletsRICoins(blockchain string, confirmedBalance string, network string, totalReceived string, totalSpent string, unit string, ) *ListAllAssetsFromAllWalletsRICoins`

NewListAllAssetsFromAllWalletsRICoins instantiates a new ListAllAssetsFromAllWalletsRICoins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllAssetsFromAllWalletsRICoinsWithDefaults

`func NewListAllAssetsFromAllWalletsRICoinsWithDefaults() *ListAllAssetsFromAllWalletsRICoins`

NewListAllAssetsFromAllWalletsRICoinsWithDefaults instantiates a new ListAllAssetsFromAllWalletsRICoins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *ListAllAssetsFromAllWalletsRICoins) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ListAllAssetsFromAllWalletsRICoins) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ListAllAssetsFromAllWalletsRICoins) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetConfirmedBalance

`func (o *ListAllAssetsFromAllWalletsRICoins) GetConfirmedBalance() string`

GetConfirmedBalance returns the ConfirmedBalance field if non-nil, zero value otherwise.

### GetConfirmedBalanceOk

`func (o *ListAllAssetsFromAllWalletsRICoins) GetConfirmedBalanceOk() (*string, bool)`

GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBalance

`func (o *ListAllAssetsFromAllWalletsRICoins) SetConfirmedBalance(v string)`

SetConfirmedBalance sets ConfirmedBalance field to given value.


### GetNetwork

`func (o *ListAllAssetsFromAllWalletsRICoins) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListAllAssetsFromAllWalletsRICoins) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListAllAssetsFromAllWalletsRICoins) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTotalReceived

`func (o *ListAllAssetsFromAllWalletsRICoins) GetTotalReceived() string`

GetTotalReceived returns the TotalReceived field if non-nil, zero value otherwise.

### GetTotalReceivedOk

`func (o *ListAllAssetsFromAllWalletsRICoins) GetTotalReceivedOk() (*string, bool)`

GetTotalReceivedOk returns a tuple with the TotalReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReceived

`func (o *ListAllAssetsFromAllWalletsRICoins) SetTotalReceived(v string)`

SetTotalReceived sets TotalReceived field to given value.


### GetTotalSpent

`func (o *ListAllAssetsFromAllWalletsRICoins) GetTotalSpent() string`

GetTotalSpent returns the TotalSpent field if non-nil, zero value otherwise.

### GetTotalSpentOk

`func (o *ListAllAssetsFromAllWalletsRICoins) GetTotalSpentOk() (*string, bool)`

GetTotalSpentOk returns a tuple with the TotalSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpent

`func (o *ListAllAssetsFromAllWalletsRICoins) SetTotalSpent(v string)`

SetTotalSpent sets TotalSpent field to given value.


### GetUnit

`func (o *ListAllAssetsFromAllWalletsRICoins) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListAllAssetsFromAllWalletsRICoins) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListAllAssetsFromAllWalletsRICoins) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


