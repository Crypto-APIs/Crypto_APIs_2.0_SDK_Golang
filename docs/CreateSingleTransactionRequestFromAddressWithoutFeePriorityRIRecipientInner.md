# CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Defines the destination address. | 
**Amount** | **string** | Defines the amount sent to the destination address. | 
**ClassicAddress** | Pointer to **string** | Represents the public address, which is a compressed and shortened form of a public key. The classic address is shown when the source address is an x-Address. | [optional] 
**Unit** | **string** | Defines the unit of the recieved amount for the address. | 

## Methods

### NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner

`func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner(address string, amount string, unit string, ) *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner`

NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInnerWithDefaults

`func NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInnerWithDefaults() *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner`

NewCreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInnerWithDefaults instantiates a new CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetClassicAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetClassicAddress() string`

GetClassicAddress returns the ClassicAddress field if non-nil, zero value otherwise.

### GetClassicAddressOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetClassicAddressOk() (*string, bool)`

GetClassicAddressOk returns a tuple with the ClassicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) SetClassicAddress(v string)`

SetClassicAddress sets ClassicAddress field to given value.

### HasClassicAddress

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) HasClassicAddress() bool`

HasClassicAddress returns a boolean if a field has been set.

### GetUnit

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CreateSingleTransactionRequestFromAddressWithoutFeePriorityRIRecipientInner) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


