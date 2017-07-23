# \OrganizationsApi

All URIs are relative to *https://app.quantimo.do/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrganizationsOrganizationIdUsersPost**](OrganizationsApi.md#V1OrganizationsOrganizationIdUsersPost) | **Post** /v1/organizations/{organizationId}/users | Get user tokens for existing users, create new users


# **V1OrganizationsOrganizationIdUsersPost**
> UserTokenSuccessfulResponse V1OrganizationsOrganizationIdUsersPost($organizationId, $body, $accessToken, $userId)

Get user tokens for existing users, create new users

Get user tokens for existing users, create new users


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32**| Organization ID | 
 **body** | [**UserTokenRequest**](UserTokenRequest.md)| Provides organization token and user ID | 
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 

### Return type

[**UserTokenSuccessfulResponse**](UserTokenSuccessfulResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

