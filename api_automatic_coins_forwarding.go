/*
CryptoAPIs

Crypto APIs is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2021-03-20
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AutomaticCoinsForwardingApiService AutomaticCoinsForwardingApi service
type AutomaticCoinsForwardingApiService service

type ApiCreateAutomaticCoinsForwardingRequest struct {
	ctx context.Context
	ApiService *AutomaticCoinsForwardingApiService
	blockchain string
	network string
	context *string
	createAutomaticCoinsForwardingRB *CreateAutomaticCoinsForwardingRB
}

// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user.
func (r ApiCreateAutomaticCoinsForwardingRequest) Context(context string) ApiCreateAutomaticCoinsForwardingRequest {
	r.context = &context
	return r
}

func (r ApiCreateAutomaticCoinsForwardingRequest) CreateAutomaticCoinsForwardingRB(createAutomaticCoinsForwardingRB CreateAutomaticCoinsForwardingRB) ApiCreateAutomaticCoinsForwardingRequest {
	r.createAutomaticCoinsForwardingRB = &createAutomaticCoinsForwardingRB
	return r
}

func (r ApiCreateAutomaticCoinsForwardingRequest) Execute() (*CreateAutomaticCoinsForwardingR, *http.Response, error) {
	return r.ApiService.CreateAutomaticCoinsForwardingExecute(r)
}

/*
CreateAutomaticCoinsForwarding Create Automatic Coins Forwarding

Through this endpoint customers can set up an automatic forwarding function specifically for coins (**not** tokens). They can have a `toAddress` which is essentially the main address and the destination for the automatic coins forwarding. 

There is also a `minimumTransferAmount` which only when reached will then trigger the forwarding. Through this the customer can save from fees.

Moreover, `feePriority` can be also set,  which defines how quickly to move the coins once they are received. The higher priority, the larger the fee will be. It can be "SLOW", "STANDARD" or "FAST".

The response of this endpoint contains an attribute `fromAddress` which can be used as a deposit address. Any funds received by this address will be automatically forwarded to `toAddress` based on what the customer has set for the automation.

For this automatic forwarding the customer can set a callback subscription.

{warning}The subscription will work for all incoming transactions until it is deleted. There is no need to do that for every transaction.{/warning}

{note}This endpoint generates a new `fromAddress` each time.{/note}

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
 @param network Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
 @return ApiCreateAutomaticCoinsForwardingRequest
*/
func (a *AutomaticCoinsForwardingApiService) CreateAutomaticCoinsForwarding(ctx context.Context, blockchain string, network string) ApiCreateAutomaticCoinsForwardingRequest {
	return ApiCreateAutomaticCoinsForwardingRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
		network: network,
	}
}

// Execute executes the request
//  @return CreateAutomaticCoinsForwardingR
func (a *AutomaticCoinsForwardingApiService) CreateAutomaticCoinsForwardingExecute(r ApiCreateAutomaticCoinsForwardingRequest) (*CreateAutomaticCoinsForwardingR, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateAutomaticCoinsForwardingR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutomaticCoinsForwardingApiService.CreateAutomaticCoinsForwarding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/blockchain-automations/{blockchain}/{network}/coins-forwarding/automations"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", url.PathEscape(parameterToString(r.network, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createAutomaticCoinsForwardingRB
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CreateAutomaticCoinsForwarding400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v CreateAutomaticCoinsForwarding401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ConvertBitcoinCashAddress402Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v CreateAutomaticCoinsForwarding403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetXRPRippleTransactionDetailsByTransactionID404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v CreateAutomaticCoinsForwarding409Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ConvertBitcoinCashAddress415Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ConvertBitcoinCashAddress422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ConvertBitcoinCashAddress429Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConvertBitcoinCashAddress500Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteAutomaticCoinsForwardingRequest struct {
	ctx context.Context
	ApiService *AutomaticCoinsForwardingApiService
	blockchain string
	network string
	referenceId string
	context *string
}

// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user.
func (r ApiDeleteAutomaticCoinsForwardingRequest) Context(context string) ApiDeleteAutomaticCoinsForwardingRequest {
	r.context = &context
	return r
}

func (r ApiDeleteAutomaticCoinsForwardingRequest) Execute() (*DeleteAutomaticCoinsForwardingR, *http.Response, error) {
	return r.ApiService.DeleteAutomaticCoinsForwardingExecute(r)
}

/*
DeleteAutomaticCoinsForwarding Delete Automatic Coins Forwarding

Through this endpoint customers can delete a forwarding function they have set for **coins** (**not** tokens).

By setting a `fromAddress` and a `toAddress`, and specifying the amount, coins can be transferred between addresses. 

A `feePriority` will be returned which represents the fee priority of the automation whether it is "SLOW", "STANDARD" OR "FAST".

{warning}The subscription will work for all incoming transactions until it is deleted. There is no need to do that for every transaction.{/warning}

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
 @param network Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
 @param referenceId Represents a unique ID used to reference the specific callback subscription.
 @return ApiDeleteAutomaticCoinsForwardingRequest
*/
func (a *AutomaticCoinsForwardingApiService) DeleteAutomaticCoinsForwarding(ctx context.Context, blockchain string, network string, referenceId string) ApiDeleteAutomaticCoinsForwardingRequest {
	return ApiDeleteAutomaticCoinsForwardingRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
		network: network,
		referenceId: referenceId,
	}
}

// Execute executes the request
//  @return DeleteAutomaticCoinsForwardingR
func (a *AutomaticCoinsForwardingApiService) DeleteAutomaticCoinsForwardingExecute(r ApiDeleteAutomaticCoinsForwardingRequest) (*DeleteAutomaticCoinsForwardingR, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteAutomaticCoinsForwardingR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutomaticCoinsForwardingApiService.DeleteAutomaticCoinsForwarding")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/blockchain-automations/{blockchain}/{network}/coins-forwarding/automations/{referenceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", url.PathEscape(parameterToString(r.network, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"referenceId"+"}", url.PathEscape(parameterToString(r.referenceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v DeleteAutomaticCoinsForwarding400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v DeleteAutomaticCoinsForwarding401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ConvertBitcoinCashAddress402Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v DeleteAutomaticCoinsForwarding403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetXRPRippleTransactionDetailsByTransactionID404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ConvertBitcoinCashAddress409Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ConvertBitcoinCashAddress415Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ConvertBitcoinCashAddress422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ConvertBitcoinCashAddress429Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConvertBitcoinCashAddress500Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListCoinsForwardingAutomationsRequest struct {
	ctx context.Context
	ApiService *AutomaticCoinsForwardingApiService
	blockchain string
	network string
	context *string
	limit *int64
	offset *int64
}

// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user.
func (r ApiListCoinsForwardingAutomationsRequest) Context(context string) ApiListCoinsForwardingAutomationsRequest {
	r.context = &context
	return r
}

// Defines how many items should be returned in the response per page basis.
func (r ApiListCoinsForwardingAutomationsRequest) Limit(limit int64) ApiListCoinsForwardingAutomationsRequest {
	r.limit = &limit
	return r
}

// The starting index of the response items, i.e. where the response should start listing the returned items.
func (r ApiListCoinsForwardingAutomationsRequest) Offset(offset int64) ApiListCoinsForwardingAutomationsRequest {
	r.offset = &offset
	return r
}

func (r ApiListCoinsForwardingAutomationsRequest) Execute() (*ListCoinsForwardingAutomationsR, *http.Response, error) {
	return r.ApiService.ListCoinsForwardingAutomationsExecute(r)
}

/*
ListCoinsForwardingAutomations List Coins Forwarding Automations

Through this endpoint customers can list all of their **coins** forwarding automations (**not** tokens).

Customers can set up automatic forwarding functions for coins by setting a `fromAddress` and a `toAddress`, and specifying the amount that can be transferred between addresses. 

A `feePriority` will be returned which represents the fee priority of the automation whether it is "SLOW", "STANDARD" OR "FAST".

{warning}The subscription will work for all transactions until it is deleted. There is no need to do that for every transaction.{/warning}

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param blockchain Represents the specific blockchain protocol name, e.g. Ethereum, Bitcoin, etc.
 @param network Represents the name of the blockchain network used; blockchain networks are usually identical as technology and software, but they differ in data, e.g. - \"mainnet\" is the live network with actual data while networks like \"testnet\", \"ropsten\" are test networks.
 @return ApiListCoinsForwardingAutomationsRequest
*/
func (a *AutomaticCoinsForwardingApiService) ListCoinsForwardingAutomations(ctx context.Context, blockchain string, network string) ApiListCoinsForwardingAutomationsRequest {
	return ApiListCoinsForwardingAutomationsRequest{
		ApiService: a,
		ctx: ctx,
		blockchain: blockchain,
		network: network,
	}
}

// Execute executes the request
//  @return ListCoinsForwardingAutomationsR
func (a *AutomaticCoinsForwardingApiService) ListCoinsForwardingAutomationsExecute(r ApiListCoinsForwardingAutomationsRequest) (*ListCoinsForwardingAutomationsR, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCoinsForwardingAutomationsR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutomaticCoinsForwardingApiService.ListCoinsForwardingAutomations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/blockchain-automations/{blockchain}/{network}/coins-forwarding/automations"
	localVarPath = strings.Replace(localVarPath, "{"+"blockchain"+"}", url.PathEscape(parameterToString(r.blockchain, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"network"+"}", url.PathEscape(parameterToString(r.network, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ListCoinsForwardingAutomations400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ListCoinsForwardingAutomations401Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ConvertBitcoinCashAddress402Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ListCoinsForwardingAutomations403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GetXRPRippleTransactionDetailsByTransactionID404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ConvertBitcoinCashAddress409Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ConvertBitcoinCashAddress415Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ConvertBitcoinCashAddress422Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ConvertBitcoinCashAddress429Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConvertBitcoinCashAddress500Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
