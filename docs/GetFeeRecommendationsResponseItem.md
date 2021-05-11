# GetFeeRecommendationsResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | **string** | Defines the unit of the fee, e.g. BTC. | 
**Fast** | **string** | Defines the fee priority as \&quot;FAST\&quot;. It works per byte, for UTXO-based protocols like Bitcoin, or per gas price, for account-based protocols like Ethereum. These are calculated based on Mempool. | 
**Slow** | **string** | Defines the fee priority as \&quot;SLOW\&quot;. It works per byte, for UTXO-based protocols like Bitcoin, or per gas price, for account-based protocols like Ethereum. These are calculated based on Mempool. | 
**Standard** | **string** | Defines the fee priority as \&quot;STANDARD\&quot;. It works per byte, for UTXO-based protocols like Bitcoin, or per gas price, for account-based protocols like Ethereum. These are calculated based on Mempool. | 

## Methods

### NewGetFeeRecommendationsResponseItem

`func NewGetFeeRecommendationsResponseItem(unit string, fast string, slow string, standard string, ) *GetFeeRecommendationsResponseItem`

NewGetFeeRecommendationsResponseItem instantiates a new GetFeeRecommendationsResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeeRecommendationsResponseItemWithDefaults

`func NewGetFeeRecommendationsResponseItemWithDefaults() *GetFeeRecommendationsResponseItem`

NewGetFeeRecommendationsResponseItemWithDefaults instantiates a new GetFeeRecommendationsResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *GetFeeRecommendationsResponseItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GetFeeRecommendationsResponseItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GetFeeRecommendationsResponseItem) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetFast

`func (o *GetFeeRecommendationsResponseItem) GetFast() string`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *GetFeeRecommendationsResponseItem) GetFastOk() (*string, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *GetFeeRecommendationsResponseItem) SetFast(v string)`

SetFast sets Fast field to given value.


### GetSlow

`func (o *GetFeeRecommendationsResponseItem) GetSlow() string`

GetSlow returns the Slow field if non-nil, zero value otherwise.

### GetSlowOk

`func (o *GetFeeRecommendationsResponseItem) GetSlowOk() (*string, bool)`

GetSlowOk returns a tuple with the Slow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlow

`func (o *GetFeeRecommendationsResponseItem) SetSlow(v string)`

SetSlow sets Slow field to given value.


### GetStandard

`func (o *GetFeeRecommendationsResponseItem) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *GetFeeRecommendationsResponseItem) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *GetFeeRecommendationsResponseItem) SetStandard(v string)`

SetStandard sets Standard field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


