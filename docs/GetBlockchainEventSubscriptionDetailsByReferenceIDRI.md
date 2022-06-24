# GetBlockchainEventSubscriptionDetailsByReferenceIDRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Represents the address of the transaction. | [optional] 
**Blockchain** | **string** | Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc. | 
**CallbackSecretKey** | Pointer to **string** | Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security). | [optional] 
**CallbackUrl** | **string** | Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs. | 
**ConfirmationsCount** | Pointer to **int32** | Represents the number of confirmations, i.e. the amount of blocks that have been built on top of this block. | [optional] 
**CreatedTimestamp** | **int32** | Defines the specific time/date when the subscription was created in Unix Timestamp. | 
**DeactivationReasons** | Pointer to [**[]ListBlockchainEventsSubscriptionsRIDeactivationReasonsInner**](ListBlockchainEventsSubscriptionsRIDeactivationReasonsInner.md) | Represents the deactivation reason details, available when a blockchain event subscription has status isActive - false. | [optional] 
**EventType** | **string** | Defines the type of the specific event available for the customer to subscribe to for callback notification. | 
**IsActive** | **bool** | Defines whether the subscription is active or not. Set as boolean. | 
**Network** | **string** | Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \&quot;mainnet\&quot; is the live network with actual data while networks like \&quot;testnet\&quot;, \&quot;ropsten\&quot; are test networks. | 
**ReferenceId** | **string** | Represents a unique ID used to reference the specific callback subscription. | 
**TransactionId** | Pointer to **string** | Represents the unique identification string that defines the transaction. | [optional] 

## Methods

### NewGetBlockchainEventSubscriptionDetailsByReferenceIDRI

`func NewGetBlockchainEventSubscriptionDetailsByReferenceIDRI(blockchain string, callbackUrl string, createdTimestamp int32, eventType string, isActive bool, network string, referenceId string, ) *GetBlockchainEventSubscriptionDetailsByReferenceIDRI`

NewGetBlockchainEventSubscriptionDetailsByReferenceIDRI instantiates a new GetBlockchainEventSubscriptionDetailsByReferenceIDRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlockchainEventSubscriptionDetailsByReferenceIDRIWithDefaults

`func NewGetBlockchainEventSubscriptionDetailsByReferenceIDRIWithDefaults() *GetBlockchainEventSubscriptionDetailsByReferenceIDRI`

NewGetBlockchainEventSubscriptionDetailsByReferenceIDRIWithDefaults instantiates a new GetBlockchainEventSubscriptionDetailsByReferenceIDRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBlockchain

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetCallbackSecretKey

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetCallbackSecretKey() string`

GetCallbackSecretKey returns the CallbackSecretKey field if non-nil, zero value otherwise.

### GetCallbackSecretKeyOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetCallbackSecretKeyOk() (*string, bool)`

GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackSecretKey

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetCallbackSecretKey(v string)`

SetCallbackSecretKey sets CallbackSecretKey field to given value.

### HasCallbackSecretKey

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) HasCallbackSecretKey() bool`

HasCallbackSecretKey returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetConfirmationsCount

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetConfirmationsCount() int32`

GetConfirmationsCount returns the ConfirmationsCount field if non-nil, zero value otherwise.

### GetConfirmationsCountOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetConfirmationsCountOk() (*int32, bool)`

GetConfirmationsCountOk returns a tuple with the ConfirmationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationsCount

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetConfirmationsCount(v int32)`

SetConfirmationsCount sets ConfirmationsCount field to given value.

### HasConfirmationsCount

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) HasConfirmationsCount() bool`

HasConfirmationsCount returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetDeactivationReasons

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetDeactivationReasons() []ListBlockchainEventsSubscriptionsRIDeactivationReasonsInner`

GetDeactivationReasons returns the DeactivationReasons field if non-nil, zero value otherwise.

### GetDeactivationReasonsOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetDeactivationReasonsOk() (*[]ListBlockchainEventsSubscriptionsRIDeactivationReasonsInner, bool)`

GetDeactivationReasonsOk returns a tuple with the DeactivationReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivationReasons

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetDeactivationReasons(v []ListBlockchainEventsSubscriptionsRIDeactivationReasonsInner)`

SetDeactivationReasons sets DeactivationReasons field to given value.

### HasDeactivationReasons

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) HasDeactivationReasons() bool`

HasDeactivationReasons returns a boolean if a field has been set.

### GetEventType

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetIsActive

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetNetwork

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetReferenceId

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetTransactionId

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetBlockchainEventSubscriptionDetailsByReferenceIDRI) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


