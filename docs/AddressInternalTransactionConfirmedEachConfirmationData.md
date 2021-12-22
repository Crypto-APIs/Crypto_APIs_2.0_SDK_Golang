# AddressInternalTransactionConfirmedEachConfirmationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** | Represents the Crypto APIs 2.0 product which sends the callback. | 
**Event** | **string** | Defines the specific event, for which a callback subscription is set. | 
**Item** | [**AddressInternalTransactionConfirmedEachConfirmationDataItem**](AddressInternalTransactionConfirmedEachConfirmationDataItem.md) |  | 

## Methods

### NewAddressInternalTransactionConfirmedEachConfirmationData

`func NewAddressInternalTransactionConfirmedEachConfirmationData(product string, event string, item AddressInternalTransactionConfirmedEachConfirmationDataItem, ) *AddressInternalTransactionConfirmedEachConfirmationData`

NewAddressInternalTransactionConfirmedEachConfirmationData instantiates a new AddressInternalTransactionConfirmedEachConfirmationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressInternalTransactionConfirmedEachConfirmationDataWithDefaults

`func NewAddressInternalTransactionConfirmedEachConfirmationDataWithDefaults() *AddressInternalTransactionConfirmedEachConfirmationData`

NewAddressInternalTransactionConfirmedEachConfirmationDataWithDefaults instantiates a new AddressInternalTransactionConfirmedEachConfirmationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetEvent

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetItem

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) GetItem() AddressInternalTransactionConfirmedEachConfirmationDataItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) GetItemOk() (*AddressInternalTransactionConfirmedEachConfirmationDataItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AddressInternalTransactionConfirmedEachConfirmationData) SetItem(v AddressInternalTransactionConfirmedEachConfirmationDataItem)`

SetItem sets Item field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


