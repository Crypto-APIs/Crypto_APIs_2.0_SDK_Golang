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


// ExchangeRatesApiService ExchangeRatesApi service
type ExchangeRatesApiService service

type ApiGetExchangeRateByAssetSymbolsRequest struct {
	ctx context.Context
	ApiService *ExchangeRatesApiService
	fromAssetSymbol string
	toAssetSymbol string
	context *string
	calculationTimestamp *int32
}

// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user.
func (r ApiGetExchangeRateByAssetSymbolsRequest) Context(context string) ApiGetExchangeRateByAssetSymbolsRequest {
	r.context = &context
	return r
}

// Defines the time of the market data used to calculate the exchange rate in UNIX Timestamp. Oldest possible timestamp is 30 days.
func (r ApiGetExchangeRateByAssetSymbolsRequest) CalculationTimestamp(calculationTimestamp int32) ApiGetExchangeRateByAssetSymbolsRequest {
	r.calculationTimestamp = &calculationTimestamp
	return r
}

func (r ApiGetExchangeRateByAssetSymbolsRequest) Execute() (*GetExchangeRateByAssetSymbolsR, *http.Response, error) {
	return r.ApiService.GetExchangeRateByAssetSymbolsExecute(r)
}

/*
GetExchangeRateByAssetSymbols Get Exchange Rate By Asset Symbols

Through this endpoint customers can obtain exchange rates by asset symbols. The process represents the exchange rate value of a single unit of one asset versus another one. Data is provided per unique asset symbol.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fromAssetSymbol Defines the base asset symbol to get a rate for.
 @param toAssetSymbol Defines the relation asset symbol in which the base asset rate will be displayed.
 @return ApiGetExchangeRateByAssetSymbolsRequest
*/
func (a *ExchangeRatesApiService) GetExchangeRateByAssetSymbols(ctx context.Context, fromAssetSymbol string, toAssetSymbol string) ApiGetExchangeRateByAssetSymbolsRequest {
	return ApiGetExchangeRateByAssetSymbolsRequest{
		ApiService: a,
		ctx: ctx,
		fromAssetSymbol: fromAssetSymbol,
		toAssetSymbol: toAssetSymbol,
	}
}

// Execute executes the request
//  @return GetExchangeRateByAssetSymbolsR
func (a *ExchangeRatesApiService) GetExchangeRateByAssetSymbolsExecute(r ApiGetExchangeRateByAssetSymbolsRequest) (*GetExchangeRateByAssetSymbolsR, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetExchangeRateByAssetSymbolsR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeRatesApiService.GetExchangeRateByAssetSymbols")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market-data/exchange-rates/by-symbols/{fromAssetSymbol}/{toAssetSymbol}"
	localVarPath = strings.Replace(localVarPath, "{"+"fromAssetSymbol"+"}", url.PathEscape(parameterToString(r.fromAssetSymbol, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toAssetSymbol"+"}", url.PathEscape(parameterToString(r.toAssetSymbol, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	if r.calculationTimestamp != nil {
		localVarQueryParams.Add("calculationTimestamp", parameterToString(*r.calculationTimestamp, ""))
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
			var v GetExchangeRateByAssetSymbols400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v GetExchangeRateByAssetSymbols401Response
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
			var v GetExchangeRateByAssetSymbols403Response
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
			var v GetExchangeRateByAssetSymbols422Response
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

type ApiGetExchangeRateByAssetsIDsRequest struct {
	ctx context.Context
	ApiService *ExchangeRatesApiService
	fromAssetId string
	toAssetId string
	context *string
	calculationTimestamp *int32
}

// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user.
func (r ApiGetExchangeRateByAssetsIDsRequest) Context(context string) ApiGetExchangeRateByAssetsIDsRequest {
	r.context = &context
	return r
}

// Defines the time of the market data used to calculate the exchange rate in UNIX Timestamp. Oldest possible timestamp is 30 days.
func (r ApiGetExchangeRateByAssetsIDsRequest) CalculationTimestamp(calculationTimestamp int32) ApiGetExchangeRateByAssetsIDsRequest {
	r.calculationTimestamp = &calculationTimestamp
	return r
}

func (r ApiGetExchangeRateByAssetsIDsRequest) Execute() (*GetExchangeRateByAssetsIDsR, *http.Response, error) {
	return r.ApiService.GetExchangeRateByAssetsIDsExecute(r)
}

/*
GetExchangeRateByAssetsIDs Get Exchange Rate By Assets IDs

Through this endpoint customers can obtain exchange rates by asset IDs. The process represents the exchange rate value of a single unit of one asset versus another one. Data is provided per unique asset Reference ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fromAssetId Defines the base asset Reference ID to get a rate for.
 @param toAssetId Defines the relation asset Reference ID in which the base asset rate will be displayed.
 @return ApiGetExchangeRateByAssetsIDsRequest
*/
func (a *ExchangeRatesApiService) GetExchangeRateByAssetsIDs(ctx context.Context, fromAssetId string, toAssetId string) ApiGetExchangeRateByAssetsIDsRequest {
	return ApiGetExchangeRateByAssetsIDsRequest{
		ApiService: a,
		ctx: ctx,
		fromAssetId: fromAssetId,
		toAssetId: toAssetId,
	}
}

// Execute executes the request
//  @return GetExchangeRateByAssetsIDsR
func (a *ExchangeRatesApiService) GetExchangeRateByAssetsIDsExecute(r ApiGetExchangeRateByAssetsIDsRequest) (*GetExchangeRateByAssetsIDsR, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetExchangeRateByAssetsIDsR
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeRatesApiService.GetExchangeRateByAssetsIDs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market-data/exchange-rates/by-asset-ids/{fromAssetId}/{toAssetId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fromAssetId"+"}", url.PathEscape(parameterToString(r.fromAssetId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toAssetId"+"}", url.PathEscape(parameterToString(r.toAssetId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.context != nil {
		localVarQueryParams.Add("context", parameterToString(*r.context, ""))
	}
	if r.calculationTimestamp != nil {
		localVarQueryParams.Add("calculationTimestamp", parameterToString(*r.calculationTimestamp, ""))
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
			var v GetExchangeRateByAssetsIDs400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v GetExchangeRateByAssetsIDs401Response
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
			var v GetExchangeRateByAssetsIDs403Response
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
			var v GetExchangeRateByAssetsIDs422Response
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
