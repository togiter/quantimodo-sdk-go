# \PairsApi

All URIs are relative to *https://app.quantimo.do/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PairsCsvGet**](PairsApi.md#V1PairsCsvGet) | **Get** /v1/pairsCsv | Get pairs
[**V1PairsGet**](PairsApi.md#V1PairsGet) | **Get** /v1/pairs | Get pairs


# **V1PairsCsvGet**
> []Pairs V1PairsCsvGet($cause, $effect, $accessToken, $userId, $causeSource, $causeUnit, $delay, $duration, $effectSource, $effectUnit, $endTime, $startTime, $limit, $offset, $sort)

Get pairs

Pairs cause measurements with effect measurements grouped over the duration of action after the onset delay.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cause** | **string**| Original variable name for the explanatory or independent variable | 
 **effect** | **string**| Original variable name for the outcome or dependent variable | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **causeSource** | **string**| Name of data source that the cause measurements should come from | [optional] 
 **causeUnit** | **string**| Abbreviated name for the unit cause measurements to be returned in | [optional] 
 **delay** | **string**| The amount of time in seconds that elapses after the predictor/stimulus event before the outcome as perceived by a self-tracker is known as the “onset delay”. For example, the “onset delay” between the time a person takes an aspirin (predictor/stimulus event) and the time a person perceives a change in their headache severity (outcome) is approximately 30 minutes. | [optional] 
 **duration** | **string**| The amount of time over which a predictor/stimulus event can exert an observable influence on an outcome variable’s value. For instance, aspirin (stimulus/predictor) typically decreases headache severity for approximately four hours (duration of action) following the onset delay. | [optional] 
 **effectSource** | **string**| Name of data source that the effectmeasurements should come from | [optional] 
 **effectUnit** | **string**| Abbreviated name for the unit effect measurements to be returned in | [optional] 
 **endTime** | **string**| The most recent date (in epoch time) for which we should return measurements | [optional] 
 **startTime** | **string**| The earliest date (in epoch time) for which we should return measurements | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Since the maximum limit is 200 records, to get more than that you&#39;ll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the &#x60;limit&#x60; and &#x60;offset&#x60; query parameters.  For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters, &#x60;imit&#x3D;20&amp;offset&#x3D;60&#x60;. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**[]Pairs**](Pairs.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PairsGet**
> []Pairs V1PairsGet($cause, $effect, $accessToken, $userId, $causeSource, $causeUnit, $delay, $duration, $effectSource, $effectUnit, $endTime, $startTime, $limit, $offset, $sort)

Get pairs

Pairs cause measurements with effect measurements grouped over the duration of action after the onset delay.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cause** | **string**| Original variable name for the explanatory or independent variable | 
 **effect** | **string**| Original variable name for the outcome or dependent variable | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **causeSource** | **string**| Name of data source that the cause measurements should come from | [optional] 
 **causeUnit** | **string**| Abbreviated name for the unit cause measurements to be returned in | [optional] 
 **delay** | **string**| The amount of time in seconds that elapses after the predictor/stimulus event before the outcome as perceived by a self-tracker is known as the “onset delay”. For example, the “onset delay” between the time a person takes an aspirin (predictor/stimulus event) and the time a person perceives a change in their headache severity (outcome) is approximately 30 minutes. | [optional] 
 **duration** | **string**| The amount of time over which a predictor/stimulus event can exert an observable influence on an outcome variable’s value. For instance, aspirin (stimulus/predictor) typically decreases headache severity for approximately four hours (duration of action) following the onset delay. | [optional] 
 **effectSource** | **string**| Name of data source that the effectmeasurements should come from | [optional] 
 **effectUnit** | **string**| Abbreviated name for the unit effect measurements to be returned in | [optional] 
 **endTime** | **string**| The most recent date (in epoch time) for which we should return measurements | [optional] 
 **startTime** | **string**| The earliest date (in epoch time) for which we should return measurements | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Since the maximum limit is 200 records, to get more than that you&#39;ll have to make multiple API calls and page through the results. To retrieve all the data, you can iterate through data by using the &#x60;limit&#x60; and &#x60;offset&#x60; query parameters.  For example, if you want to retrieve data from 61-80 then you can use a query with the following parameters, &#x60;imit&#x3D;20&amp;offset&#x3D;60&#x60;. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**[]Pairs**](Pairs.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

