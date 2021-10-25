# GetTransactionRequestDetailsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | **string** | Defines an optional note for additional details. | 
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**FeePriority** | **string** | Defines the priority for the fee, if it is \&quot;slow\&quot;, \&quot;standard\&quot; or \&quot;fast\&quot;. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**Recipients** | [**[]GetTransactionRequestDetailsRIRecipients**](GetTransactionRequestDetailsRIRecipients.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**TotalTransactionAmount** | **string** | Defines the total transaction amount. | 
**TransactionRequestStatus** | **string** | Defines the status of the transaction request, e.g. pending. | 
**TransactionType** | **string** | Defines the transaction type, if it is for coins or tokens. | 
**Unit** | **string** | Defines the unit of the amount. | 
**WalletId** | **string** | Defines the unique ID of the Wallet. | 

## Methods

### NewGetTransactionRequestDetailsRI

`func NewGetTransactionRequestDetailsRI(additionalDetails string, blockchain string, feePriority string, network string, recipients []GetTransactionRequestDetailsRIRecipients, totalTransactionAmount string, transactionRequestStatus string, transactionType string, unit string, walletId string, ) *GetTransactionRequestDetailsRI`

NewGetTransactionRequestDetailsRI instantiates a new GetTransactionRequestDetailsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionRequestDetailsRIWithDefaults

`func NewGetTransactionRequestDetailsRIWithDefaults() *GetTransactionRequestDetailsRI`

NewGetTransactionRequestDetailsRIWithDefaults instantiates a new GetTransactionRequestDetailsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *GetTransactionRequestDetailsRI) GetAdditionalDetails() string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *GetTransactionRequestDetailsRI) GetAdditionalDetailsOk() (*string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *GetTransactionRequestDetailsRI) SetAdditionalDetails(v string)`

SetAdditionalDetails sets AdditionalDetails field to given value.


### GetBlockchain

`func (o *GetTransactionRequestDetailsRI) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *GetTransactionRequestDetailsRI) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *GetTransactionRequestDetailsRI) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetFeePriority

`func (o *GetTransactionRequestDetailsRI) GetFeePriority() string`

GetFeePriority returns the FeePriority field if non-nil, zero value otherwise.

### GetFeePriorityOk

`func (o *GetTransactionRequestDetailsRI) GetFeePriorityOk() (*string, bool)`

GetFeePriorityOk returns a tuple with the FeePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePriority

`func (o *GetTransactionRequestDetailsRI) SetFeePriority(v string)`

SetFeePriority sets FeePriority field to given value.


### GetNetwork

`func (o *GetTransactionRequestDetailsRI) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetTransactionRequestDetailsRI) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetTransactionRequestDetailsRI) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetRecipients

`func (o *GetTransactionRequestDetailsRI) GetRecipients() []GetTransactionRequestDetailsRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetTransactionRequestDetailsRI) GetRecipientsOk() (*[]GetTransactionRequestDetailsRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetTransactionRequestDetailsRI) SetRecipients(v []GetTransactionRequestDetailsRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetTotalTransactionAmount

`func (o *GetTransactionRequestDetailsRI) GetTotalTransactionAmount() string`

GetTotalTransactionAmount returns the TotalTransactionAmount field if non-nil, zero value otherwise.

### GetTotalTransactionAmountOk

`func (o *GetTransactionRequestDetailsRI) GetTotalTransactionAmountOk() (*string, bool)`

GetTotalTransactionAmountOk returns a tuple with the TotalTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactionAmount

`func (o *GetTransactionRequestDetailsRI) SetTotalTransactionAmount(v string)`

SetTotalTransactionAmount sets TotalTransactionAmount field to given value.


### GetTransactionRequestStatus

`func (o *GetTransactionRequestDetailsRI) GetTransactionRequestStatus() string`

GetTransactionRequestStatus returns the TransactionRequestStatus field if non-nil, zero value otherwise.

### GetTransactionRequestStatusOk

`func (o *GetTransactionRequestDetailsRI) GetTransactionRequestStatusOk() (*string, bool)`

GetTransactionRequestStatusOk returns a tuple with the TransactionRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRequestStatus

`func (o *GetTransactionRequestDetailsRI) SetTransactionRequestStatus(v string)`

SetTransactionRequestStatus sets TransactionRequestStatus field to given value.


### GetTransactionType

`func (o *GetTransactionRequestDetailsRI) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *GetTransactionRequestDetailsRI) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *GetTransactionRequestDetailsRI) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetUnit

`func (o *GetTransactionRequestDetailsRI) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetTransactionRequestDetailsRI) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetTransactionRequestDetailsRI) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetWalletId

`func (o *GetTransactionRequestDetailsRI) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *GetTransactionRequestDetailsRI) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *GetTransactionRequestDetailsRI) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


