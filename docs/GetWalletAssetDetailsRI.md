# GetWalletAssetDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmedBalance** | [**GetWalletAssetDetailsRIConfirmedBalance**](GetWalletAssetDetailsRIConfirmedBalance.md) |  | 
**DepositAddressesCount** | **int32** | Specifies the count of deposit addresses in the Wallet. | 
**FungibleTokens** | [**[]GetWalletAssetDetailsRIFungibleTokensInner**](GetWalletAssetDetailsRIFungibleTokensInner.md) | Represents fungible tokens&#39;es detailed information | 
**Name** | **string** | Defines the name of the Wallet given to it by the user. | 
**NonFungibleTokens** | [**[]GetWalletAssetDetailsRINonFungibleTokensInner**](GetWalletAssetDetailsRINonFungibleTokensInner.md) | Represents non-fungible tokens&#39;es detailed information. | 
**RecievedConfirmedAmount** | [**GetWalletAssetDetailsRIRecievedConfirmedAmount**](GetWalletAssetDetailsRIRecievedConfirmedAmount.md) |  | 
**SentConfirmedAmount** | [**GetWalletAssetDetailsRISentConfirmedAmount**](GetWalletAssetDetailsRISentConfirmedAmount.md) |  | 

## Methods

### NewGetWalletAssetDetailsRI

`func NewGetWalletAssetDetailsRI(confirmedBalance GetWalletAssetDetailsRIConfirmedBalance, depositAddressesCount int32, fungibleTokens []GetWalletAssetDetailsRIFungibleTokensInner, name string, nonFungibleTokens []GetWalletAssetDetailsRINonFungibleTokensInner, recievedConfirmedAmount GetWalletAssetDetailsRIRecievedConfirmedAmount, sentConfirmedAmount GetWalletAssetDetailsRISentConfirmedAmount, ) *GetWalletAssetDetailsRI`

NewGetWalletAssetDetailsRI instantiates a new GetWalletAssetDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWalletAssetDetailsRIWithDefaults

`func NewGetWalletAssetDetailsRIWithDefaults() *GetWalletAssetDetailsRI`

NewGetWalletAssetDetailsRIWithDefaults instantiates a new GetWalletAssetDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmedBalance

`func (o *GetWalletAssetDetailsRI) GetConfirmedBalance() GetWalletAssetDetailsRIConfirmedBalance`

GetConfirmedBalance returns the ConfirmedBalance field if non-nil, zero value otherwise.

### GetConfirmedBalanceOk

`func (o *GetWalletAssetDetailsRI) GetConfirmedBalanceOk() (*GetWalletAssetDetailsRIConfirmedBalance, bool)`

GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBalance

`func (o *GetWalletAssetDetailsRI) SetConfirmedBalance(v GetWalletAssetDetailsRIConfirmedBalance)`

SetConfirmedBalance sets ConfirmedBalance field to given value.


### GetDepositAddressesCount

`func (o *GetWalletAssetDetailsRI) GetDepositAddressesCount() int32`

GetDepositAddressesCount returns the DepositAddressesCount field if non-nil, zero value otherwise.

### GetDepositAddressesCountOk

`func (o *GetWalletAssetDetailsRI) GetDepositAddressesCountOk() (*int32, bool)`

GetDepositAddressesCountOk returns a tuple with the DepositAddressesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositAddressesCount

`func (o *GetWalletAssetDetailsRI) SetDepositAddressesCount(v int32)`

SetDepositAddressesCount sets DepositAddressesCount field to given value.


### GetFungibleTokens

`func (o *GetWalletAssetDetailsRI) GetFungibleTokens() []GetWalletAssetDetailsRIFungibleTokensInner`

GetFungibleTokens returns the FungibleTokens field if non-nil, zero value otherwise.

### GetFungibleTokensOk

`func (o *GetWalletAssetDetailsRI) GetFungibleTokensOk() (*[]GetWalletAssetDetailsRIFungibleTokensInner, bool)`

GetFungibleTokensOk returns a tuple with the FungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFungibleTokens

`func (o *GetWalletAssetDetailsRI) SetFungibleTokens(v []GetWalletAssetDetailsRIFungibleTokensInner)`

SetFungibleTokens sets FungibleTokens field to given value.


### GetName

`func (o *GetWalletAssetDetailsRI) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetWalletAssetDetailsRI) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetWalletAssetDetailsRI) SetName(v string)`

SetName sets Name field to given value.


### GetNonFungibleTokens

`func (o *GetWalletAssetDetailsRI) GetNonFungibleTokens() []GetWalletAssetDetailsRINonFungibleTokensInner`

GetNonFungibleTokens returns the NonFungibleTokens field if non-nil, zero value otherwise.

### GetNonFungibleTokensOk

`func (o *GetWalletAssetDetailsRI) GetNonFungibleTokensOk() (*[]GetWalletAssetDetailsRINonFungibleTokensInner, bool)`

GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFungibleTokens

`func (o *GetWalletAssetDetailsRI) SetNonFungibleTokens(v []GetWalletAssetDetailsRINonFungibleTokensInner)`

SetNonFungibleTokens sets NonFungibleTokens field to given value.


### GetRecievedConfirmedAmount

`func (o *GetWalletAssetDetailsRI) GetRecievedConfirmedAmount() GetWalletAssetDetailsRIRecievedConfirmedAmount`

GetRecievedConfirmedAmount returns the RecievedConfirmedAmount field if non-nil, zero value otherwise.

### GetRecievedConfirmedAmountOk

`func (o *GetWalletAssetDetailsRI) GetRecievedConfirmedAmountOk() (*GetWalletAssetDetailsRIRecievedConfirmedAmount, bool)`

GetRecievedConfirmedAmountOk returns a tuple with the RecievedConfirmedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecievedConfirmedAmount

`func (o *GetWalletAssetDetailsRI) SetRecievedConfirmedAmount(v GetWalletAssetDetailsRIRecievedConfirmedAmount)`

SetRecievedConfirmedAmount sets RecievedConfirmedAmount field to given value.


### GetSentConfirmedAmount

`func (o *GetWalletAssetDetailsRI) GetSentConfirmedAmount() GetWalletAssetDetailsRISentConfirmedAmount`

GetSentConfirmedAmount returns the SentConfirmedAmount field if non-nil, zero value otherwise.

### GetSentConfirmedAmountOk

`func (o *GetWalletAssetDetailsRI) GetSentConfirmedAmountOk() (*GetWalletAssetDetailsRISentConfirmedAmount, bool)`

GetSentConfirmedAmountOk returns a tuple with the SentConfirmedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentConfirmedAmount

`func (o *GetWalletAssetDetailsRI) SetSentConfirmedAmount(v GetWalletAssetDetailsRISentConfirmedAmount)`

SetSentConfirmedAmount sets SentConfirmedAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


