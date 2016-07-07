# \PairsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1PairsCsvGet**](PairsApi.md#V1PairsCsvGet) | **Get** /v1/pairsCsv | Get pairs
[**V1PairsGet**](PairsApi.md#V1PairsGet) | **Get** /v1/pairs | Get pairs


# **V1PairsCsvGet**
> []Pairs V1PairsCsvGet($cause, $effect, $accessToken, $causeSource, $causeUnit, $delay, $duration, $effectSource, $effectUnit, $endTime, $startTime, $limit, $offset, $sort)

Get pairs

Pairs cause measurements with effect measurements grouped over the duration of action after the onset delay.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cause** | **string**| Original variable name for the explanatory or independent variable | 
 **effect** | **string**| Original variable name for the outcome or dependent variable | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **causeSource** | **string**| Name of data source that the cause measurements should come from | [optional] 
 **causeUnit** | **string**| Abbreviated name for the unit cause measurements to be returned in | [optional] 
 **delay** | **string**| Delay before onset of action (in seconds) from the cause variable settings. | [optional] 
 **duration** | **string**| Duration of action (in seconds) from the cause variable settings. | [optional] 
 **effectSource** | **string**| Name of data source that the effectmeasurements should come from | [optional] 
 **effectUnit** | **string**| Abbreviated name for the unit effect measurements to be returned in | [optional] 
 **endTime** | **string**| The most recent date (in epoch time) for which we should return measurements | [optional] 
 **startTime** | **string**| The earliest date (in epoch time) for which we should return measurements | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**[]Pairs**](Pairs.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1PairsGet**
> []Pairs V1PairsGet($cause, $effect, $accessToken, $causeSource, $causeUnit, $delay, $duration, $effectSource, $effectUnit, $endTime, $startTime, $limit, $offset, $sort)

Get pairs

Pairs cause measurements with effect measurements grouped over the duration of action after the onset delay.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cause** | **string**| Original variable name for the explanatory or independent variable | 
 **effect** | **string**| Original variable name for the outcome or dependent variable | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **causeSource** | **string**| Name of data source that the cause measurements should come from | [optional] 
 **causeUnit** | **string**| Abbreviated name for the unit cause measurements to be returned in | [optional] 
 **delay** | **string**| Delay before onset of action (in seconds) from the cause variable settings. | [optional] 
 **duration** | **string**| Duration of action (in seconds) from the cause variable settings. | [optional] 
 **effectSource** | **string**| Name of data source that the effectmeasurements should come from | [optional] 
 **effectUnit** | **string**| Abbreviated name for the unit effect measurements to be returned in | [optional] 
 **endTime** | **string**| The most recent date (in epoch time) for which we should return measurements | [optional] 
 **startTime** | **string**| The earliest date (in epoch time) for which we should return measurements | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. | [optional] 
 **offset** | **int32**| Now suppose you wanted to show results 11-20. You&#39;d set the offset to 10 and the limit to 10. | [optional] 
 **sort** | **int32**| Sort by given field. If the field is prefixed with &#x60;-, it will sort in descending order. | [optional] 

### Return type

[**[]Pairs**](Pairs.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

