# GetOmniTransactionDetailsByTransactionIDTxidE401

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Specifies an error code, e.g. error 404. | 
**Message** | **string** | Specifies the message of the error, i.e. why the error was returned, e.g. error 404 stands for “not found”. | 
**Details** | Pointer to [**[]BannedIpAddressDetailsInner**](BannedIpAddressDetailsInner.md) |  | [optional] 

## Methods

### NewGetOmniTransactionDetailsByTransactionIDTxidE401

`func NewGetOmniTransactionDetailsByTransactionIDTxidE401(code string, message string, ) *GetOmniTransactionDetailsByTransactionIDTxidE401`

NewGetOmniTransactionDetailsByTransactionIDTxidE401 instantiates a new GetOmniTransactionDetailsByTransactionIDTxidE401 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOmniTransactionDetailsByTransactionIDTxidE401WithDefaults

`func NewGetOmniTransactionDetailsByTransactionIDTxidE401WithDefaults() *GetOmniTransactionDetailsByTransactionIDTxidE401`

NewGetOmniTransactionDetailsByTransactionIDTxidE401WithDefaults instantiates a new GetOmniTransactionDetailsByTransactionIDTxidE401 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) GetDetails() []BannedIpAddressDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) GetDetailsOk() (*[]BannedIpAddressDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) SetDetails(v []BannedIpAddressDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetOmniTransactionDetailsByTransactionIDTxidE401) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


