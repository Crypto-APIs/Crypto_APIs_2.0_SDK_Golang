# CreateCoinsTransactionRequestFromAddressRIRecipientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Defines the destination address. | 
**AddressTag** | Pointer to **int32** | Defines a specific Tag that is an additional XRP address feature. It helps identify a transaction recipient beyond a wallet address. The tag that was encoded into the x-Address along with the Source Classic Address. | [optional] 
**Amount** | **string** | Defines the amount sent to the destination address. | 
**ClassicAddress** | Pointer to **string** | Represents the public address, which is a compressed and shortened form of a public key. The classic address is shown when the source address is an x-Address. | [optional] 

## Methods

### NewCreateCoinsTransactionRequestFromAddressRIRecipientsInner

`func NewCreateCoinsTransactionRequestFromAddressRIRecipientsInner(address string, amount string, ) *CreateCoinsTransactionRequestFromAddressRIRecipientsInner`

NewCreateCoinsTransactionRequestFromAddressRIRecipientsInner instantiates a new CreateCoinsTransactionRequestFromAddressRIRecipientsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCoinsTransactionRequestFromAddressRIRecipientsInnerWithDefaults

`func NewCreateCoinsTransactionRequestFromAddressRIRecipientsInnerWithDefaults() *CreateCoinsTransactionRequestFromAddressRIRecipientsInner`

NewCreateCoinsTransactionRequestFromAddressRIRecipientsInnerWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRIRecipientsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressTag

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetAddressTag() int32`

GetAddressTag returns the AddressTag field if non-nil, zero value otherwise.

### GetAddressTagOk

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetAddressTagOk() (*int32, bool)`

GetAddressTagOk returns a tuple with the AddressTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressTag

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) SetAddressTag(v int32)`

SetAddressTag sets AddressTag field to given value.

### HasAddressTag

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) HasAddressTag() bool`

HasAddressTag returns a boolean if a field has been set.

### GetAmount

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClassicAddress

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetClassicAddress() string`

GetClassicAddress returns the ClassicAddress field if non-nil, zero value otherwise.

### GetClassicAddressOk

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) GetClassicAddressOk() (*string, bool)`

GetClassicAddressOk returns a tuple with the ClassicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicAddress

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) SetClassicAddress(v string)`

SetClassicAddress sets ClassicAddress field to given value.

### HasClassicAddress

`func (o *CreateCoinsTransactionRequestFromAddressRIRecipientsInner) HasClassicAddress() bool`

HasClassicAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


