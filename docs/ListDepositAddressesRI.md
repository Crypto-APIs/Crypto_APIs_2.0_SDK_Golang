# ListDepositAddressesRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Specifies the specific address&#39;s unique string value. | 
**ConfirmedBalance** | [**ListDepositAddressesRIConfirmedBalance**](ListDepositAddressesRIConfirmedBalance.md) |  | 
**CreatedTimestamp** | **int32** | Defines the specific UNIX time when the deposit address was created. | 
**FungibleTokens** | [**[]ListDepositAddressesRIFungibleTokens**](ListDepositAddressesRIFungibleTokens.md) | Represents fungible tokens&#39;es detailed information | 
**Index** | **string** | Represents the index of the address in the wallet. | 
**Label** | **string** | Represents a custom tag that customers can set up for their Wallets and addresses. E.g. custom label named \&quot;Special addresses\&quot;. | 
**NonFungibleTokens** | [**[]ListDepositAddressesRINonFungibleTokens**](ListDepositAddressesRINonFungibleTokens.md) | Represents non-fungible tokens&#39;es detailed information. | 

## Methods

### NewListDepositAddressesRI

`func NewListDepositAddressesRI(address string, confirmedBalance ListDepositAddressesRIConfirmedBalance, createdTimestamp int32, fungibleTokens []ListDepositAddressesRIFungibleTokens, index string, label string, nonFungibleTokens []ListDepositAddressesRINonFungibleTokens, ) *ListDepositAddressesRI`

NewListDepositAddressesRI instantiates a new ListDepositAddressesRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDepositAddressesRIWithDefaults

`func NewListDepositAddressesRIWithDefaults() *ListDepositAddressesRI`

NewListDepositAddressesRIWithDefaults instantiates a new ListDepositAddressesRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ListDepositAddressesRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ListDepositAddressesRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ListDepositAddressesRI) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetConfirmedBalance

`func (o *ListDepositAddressesRI) GetConfirmedBalance() ListDepositAddressesRIConfirmedBalance`

GetConfirmedBalance returns the ConfirmedBalance field if non-nil, zero value otherwise.

### GetConfirmedBalanceOk

`func (o *ListDepositAddressesRI) GetConfirmedBalanceOk() (*ListDepositAddressesRIConfirmedBalance, bool)`

GetConfirmedBalanceOk returns a tuple with the ConfirmedBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBalance

`func (o *ListDepositAddressesRI) SetConfirmedBalance(v ListDepositAddressesRIConfirmedBalance)`

SetConfirmedBalance sets ConfirmedBalance field to given value.


### GetCreatedTimestamp

`func (o *ListDepositAddressesRI) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ListDepositAddressesRI) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ListDepositAddressesRI) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetFungibleTokens

`func (o *ListDepositAddressesRI) GetFungibleTokens() []ListDepositAddressesRIFungibleTokens`

GetFungibleTokens returns the FungibleTokens field if non-nil, zero value otherwise.

### GetFungibleTokensOk

`func (o *ListDepositAddressesRI) GetFungibleTokensOk() (*[]ListDepositAddressesRIFungibleTokens, bool)`

GetFungibleTokensOk returns a tuple with the FungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFungibleTokens

`func (o *ListDepositAddressesRI) SetFungibleTokens(v []ListDepositAddressesRIFungibleTokens)`

SetFungibleTokens sets FungibleTokens field to given value.


### GetIndex

`func (o *ListDepositAddressesRI) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ListDepositAddressesRI) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ListDepositAddressesRI) SetIndex(v string)`

SetIndex sets Index field to given value.


### GetLabel

`func (o *ListDepositAddressesRI) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListDepositAddressesRI) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListDepositAddressesRI) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNonFungibleTokens

`func (o *ListDepositAddressesRI) GetNonFungibleTokens() []ListDepositAddressesRINonFungibleTokens`

GetNonFungibleTokens returns the NonFungibleTokens field if non-nil, zero value otherwise.

### GetNonFungibleTokensOk

`func (o *ListDepositAddressesRI) GetNonFungibleTokensOk() (*[]ListDepositAddressesRINonFungibleTokens, bool)`

GetNonFungibleTokensOk returns a tuple with the NonFungibleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonFungibleTokens

`func (o *ListDepositAddressesRI) SetNonFungibleTokens(v []ListDepositAddressesRINonFungibleTokens)`

SetNonFungibleTokens sets NonFungibleTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


