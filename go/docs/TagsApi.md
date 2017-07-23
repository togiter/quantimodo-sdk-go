# \TagsApi

All URIs are relative to *https://app.quantimo.do/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UserTagsDeletePost**](TagsApi.md#V1UserTagsDeletePost) | **Post** /v1/userTags/delete | Delete user tag or ingredient
[**V1UserTagsPost**](TagsApi.md#V1UserTagsPost) | **Post** /v1/userTags | Post or update user tags or ingredients


# **V1UserTagsDeletePost**
> CommonResponse V1UserTagsDeletePost($taggedVariableId, $tagVariableId)

Delete user tag or ingredient

Delete previously created user tags or ingredients.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taggedVariableId** | **int32**| This is the id of the variable being tagged with an ingredient or something. | 
 **tagVariableId** | **int32**| This is the id of the ingredient variable whose value is determined based on the value of the tagged variable. | 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UserTagsPost**
> CommonResponse V1UserTagsPost($body, $accessToken, $userId)

Post or update user tags or ingredients

This endpoint allows users to tag foods with their ingredients.  This information will then be used to infer the user intake of the different ingredients by just entering the foods. The inferred intake levels will then be used to determine the effects of different nutrients on the user during analysis.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserTag**](UserTag.md)| Contains the new user tag data | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

[**CommonResponse**](CommonResponse.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

