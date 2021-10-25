# AddressTokensTransactionConfirmedEachConfirmationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Specifies the name of the token. | 
**Symbol** | **string** | Specifies an identifier of the token, where up to five alphanumeric characters can be used for it. | 
**Decimals** | Pointer to **string** | Defines how many decimals can be used to break the token. | [optional] 
**Amount** | **string** | Defines the amount of tokens sent with the confirmed transaction. | 
**ContractAddress** | **string** | Defines the address of the contract. | 
**TokenId** | **string** | Specifies the ID of the token. | 
**PropertyId** | **string** | Defines the ID of the property for Omni Layer. | 
**TransactionType** | **string** | Defines the type of the transaction. | 
**CreatedByTransactionId** | **string** | The transaction ID used to create the token. | 

## Methods

### NewAddressTokensTransactionConfirmedEachConfirmationToken

`func NewAddressTokensTransactionConfirmedEachConfirmationToken(name string, symbol string, amount string, contractAddress string, tokenId string, propertyId string, transactionType string, createdByTransactionId string, ) *AddressTokensTransactionConfirmedEachConfirmationToken`

NewAddressTokensTransactionConfirmedEachConfirmationToken instantiates a new AddressTokensTransactionConfirmedEachConfirmationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTokensTransactionConfirmedEachConfirmationTokenWithDefaults

`func NewAddressTokensTransactionConfirmedEachConfirmationTokenWithDefaults() *AddressTokensTransactionConfirmedEachConfirmationToken`

NewAddressTokensTransactionConfirmedEachConfirmationTokenWithDefaults instantiates a new AddressTokensTransactionConfirmedEachConfirmationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetAmount

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetContractAddress

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetTokenId

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetPropertyId

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetPropertyId() string`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetPropertyIdOk() (*string, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetPropertyId(v string)`

SetPropertyId sets PropertyId field to given value.


### GetTransactionType

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetCreatedByTransactionId

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetCreatedByTransactionId() string`

GetCreatedByTransactionId returns the CreatedByTransactionId field if non-nil, zero value otherwise.

### GetCreatedByTransactionIdOk

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) GetCreatedByTransactionIdOk() (*string, bool)`

GetCreatedByTransactionIdOk returns a tuple with the CreatedByTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByTransactionId

`func (o *AddressTokensTransactionConfirmedEachConfirmationToken) SetCreatedByTransactionId(v string)`

SetCreatedByTransactionId sets CreatedByTransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


