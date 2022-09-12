# GetTransactionDetailsByTransactionIDFromCallbackRIBST

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Defines the amount of the transaction. | 
**BandwidthUsed** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed**](GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed.md) |  | 
**Contract** | **string** | Represents the specific transaction contract. | 
**EnergyUsed** | [**GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed**](GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed.md) |  | 
**HasInternalTransactions** | **bool** | Defines if the transaction includes internal transactions (true) or not (false). | 
**HasTokenTransfers** | **string** | Defines if the transaction includes token transfers (true) or not (false). | 
**Input** | **string** | Represents the transaction&#39;s input value. | 
**Recipients** | **string** | Represents the recipient address. | 
**Senders** | **string** | Represents the sender address. | 
**TransactionStatus** | **string** | Represents the status of this transaction. | 

## Methods

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBST

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBST(amount string, bandwidthUsed GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed, contract string, energyUsed GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed, hasInternalTransactions bool, hasTokenTransfers string, input string, recipients string, senders string, transactionStatus string, ) *GetTransactionDetailsByTransactionIDFromCallbackRIBST`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBST instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBST object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionDetailsByTransactionIDFromCallbackRIBSTWithDefaults

`func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSTWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBST`

NewGetTransactionDetailsByTransactionIDFromCallbackRIBSTWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBST object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetBandwidthUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetBandwidthUsed() GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed`

GetBandwidthUsed returns the BandwidthUsed field if non-nil, zero value otherwise.

### GetBandwidthUsedOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetBandwidthUsedOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed, bool)`

GetBandwidthUsedOk returns a tuple with the BandwidthUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetBandwidthUsed(v GetTransactionDetailsByTransactionIDFromCallbackRIBSTBandwidthUsed)`

SetBandwidthUsed sets BandwidthUsed field to given value.


### GetContract

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetContract(v string)`

SetContract sets Contract field to given value.


### GetEnergyUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetEnergyUsed() GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed`

GetEnergyUsed returns the EnergyUsed field if non-nil, zero value otherwise.

### GetEnergyUsedOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetEnergyUsedOk() (*GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed, bool)`

GetEnergyUsedOk returns a tuple with the EnergyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyUsed

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetEnergyUsed(v GetTransactionDetailsByTransactionIDFromCallbackRIBSTEnergyUsed)`

SetEnergyUsed sets EnergyUsed field to given value.


### GetHasInternalTransactions

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasInternalTransactions() bool`

GetHasInternalTransactions returns the HasInternalTransactions field if non-nil, zero value otherwise.

### GetHasInternalTransactionsOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasInternalTransactionsOk() (*bool, bool)`

GetHasInternalTransactionsOk returns a tuple with the HasInternalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInternalTransactions

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetHasInternalTransactions(v bool)`

SetHasInternalTransactions sets HasInternalTransactions field to given value.


### GetHasTokenTransfers

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasTokenTransfers() string`

GetHasTokenTransfers returns the HasTokenTransfers field if non-nil, zero value otherwise.

### GetHasTokenTransfersOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetHasTokenTransfersOk() (*string, bool)`

GetHasTokenTransfersOk returns a tuple with the HasTokenTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTokenTransfers

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetHasTokenTransfers(v string)`

SetHasTokenTransfers sets HasTokenTransfers field to given value.


### GetInput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetInput(v string)`

SetInput sets Input field to given value.


### GetRecipients

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetRecipients() string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetRecipientsOk() (*string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetRecipients(v string)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetSenders() string`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetSendersOk() (*string, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetSenders(v string)`

SetSenders sets Senders field to given value.


### GetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBST) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


