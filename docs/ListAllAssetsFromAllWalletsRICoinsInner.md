# ListAllAssetsFromAllWalletsRICoinsInner

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

### NewListAllAssetsFromAllWalletsRICoinsInner

`func NewListAllAssetsFromAllWalletsRICoinsInner(blockchain string, confirmedBalance string, network string, totalReceived string, totalSpent string, unit string, ) *ListAllAssetsFromAllWalletsRICoinsInner`

NewListAllAssetsFromAllWalletsRICoinsInner instantiates a new ListAllAssetsFromAllWalletsRICoinsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllAssetsFromAllWalletsRICoinsInnerWithDefaults

`func NewListAllAssetsFromAllWalletsRICoinsInnerWithDefaults() *ListAllAssetsFromAllWalletsRICoinsInner`

NewListAllAssetsFromAllWalletsRICoinsInnerWithDefaults instantiates a new ListAllAssetsFromAllWalletsRICoinsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockchain

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetConfirmedBalance

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetConfirmedBalance() string`

GetConfirmedBalance returns the ConfirmedBalance field if non-nil, zero value otherwise.

### GetConfirmedBalanceOk

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetConfirmedBalanceOk() (*string, bool)`

GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBalance

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) SetConfirmedBalance(v string)`

SetConfirmedBalance sets ConfirmedBalance field to given value.


### GetNetwork

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetTotalReceived

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetTotalReceived() string`

GetTotalReceived returns the TotalReceived field if non-nil, zero value otherwise.

### GetTotalReceivedOk

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetTotalReceivedOk() (*string, bool)`

GetTotalReceivedOk returns a tuple with the TotalReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReceived

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) SetTotalReceived(v string)`

SetTotalReceived sets TotalReceived field to given value.


### GetTotalSpent

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetTotalSpent() string`

GetTotalSpent returns the TotalSpent field if non-nil, zero value otherwise.

### GetTotalSpentOk

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetTotalSpentOk() (*string, bool)`

GetTotalSpentOk returns a tuple with the TotalSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpent

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) SetTotalSpent(v string)`

SetTotalSpent sets TotalSpent field to given value.


### GetUnit

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ListAllAssetsFromAllWalletsRICoinsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


