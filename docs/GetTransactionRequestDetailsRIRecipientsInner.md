# GetTransactionRequestDetailsRIRecipientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address which receives this transaction. In UTXO-based protocols like Bitcoin there could be several senders while in account-based protocols like Ethereum there is always only one recipient. | 
**AddressTag** | Pointer to **int32** | Defines a specific Tag that is an additional XRP address feature. It helps identify a transaction recipient beyond a wallet address. The tag that was encoded into the x-Address along with the Classic Address. | [optional] 
**Amount** | **string** | Represents the amount received to this address. | 
**ClassicAddress** | Pointer to **string** | Represents the public address, which is a compressed and shortened form of a public key. A classic address is shown when the destination address is an x-Address. | [optional] 
**Unit** | **string** | Defines the unit of the amount. | 

## Methods

### NewGetTransactionRequestDetailsRIRecipientsInner

`func NewGetTransactionRequestDetailsRIRecipientsInner(address string, amount string, unit string, ) *GetTransactionRequestDetailsRIRecipientsInner`

NewGetTransactionRequestDetailsRIRecipientsInner instantiates a new GetTransactionRequestDetailsRIRecipientsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionRequestDetailsRIRecipientsInnerWithDefaults

`func NewGetTransactionRequestDetailsRIRecipientsInnerWithDefaults() *GetTransactionRequestDetailsRIRecipientsInner`

NewGetTransactionRequestDetailsRIRecipientsInnerWithDefaults instantiates a new GetTransactionRequestDetailsRIRecipientsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetTransactionRequestDetailsRIRecipientsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressTag

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetAddressTag() int32`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetAddressTagOk() (*int32, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *GetTransactionRequestDetailsRIRecipientsInner) SetAddressTag(v int32)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *GetTransactionRequestDetailsRIRecipientsInner) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

### GetAmount

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetTransactionRequestDetailsRIRecipientsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClassicAddress

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetClassicAddress() string`

GetClassicAddress returns the ClassicAddress field if non-nil, zero value otherwise.

### GetClassicAddressOk

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetClassicAddressOk() (*string, bool)`

GetClassicAddressOk returns a tuple with the ClassicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicAddress

`func (o *GetTransactionRequestDetailsRIRecipientsInner) SetClassicAddress(v string)`

SetClassicAddress sets ClassicAddress field to given value.

### HasClassicAddress

`func (o *GetTransactionRequestDetailsRIRecipientsInner) HasClassicAddress() bool`

HasClassicAddress returns a boolean if a field has been set.

### GetUnit

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetTransactionRequestDetailsRIRecipientsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetTransactionRequestDetailsRIRecipientsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


