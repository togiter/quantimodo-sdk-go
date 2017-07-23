# \RemindersApi

All URIs are relative to *https://app.quantimo.do/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1TrackingReminderNotificationsGet**](RemindersApi.md#V1TrackingReminderNotificationsGet) | **Get** /v1/trackingReminderNotifications | Get specific pending tracking reminders
[**V1TrackingReminderNotificationsSkipPost**](RemindersApi.md#V1TrackingReminderNotificationsSkipPost) | **Post** /v1/trackingReminderNotifications/skip | Skip a pending tracking reminder
[**V1TrackingReminderNotificationsSnoozePost**](RemindersApi.md#V1TrackingReminderNotificationsSnoozePost) | **Post** /v1/trackingReminderNotifications/snooze | Snooze a pending tracking reminder
[**V1TrackingReminderNotificationsTrackPost**](RemindersApi.md#V1TrackingReminderNotificationsTrackPost) | **Post** /v1/trackingReminderNotifications/track | Track a pending tracking reminder
[**V1TrackingRemindersDeletePost**](RemindersApi.md#V1TrackingRemindersDeletePost) | **Post** /v1/trackingReminders/delete | Delete tracking reminder
[**V1TrackingRemindersGet**](RemindersApi.md#V1TrackingRemindersGet) | **Get** /v1/trackingReminders | Get repeating tracking reminder settings
[**V1TrackingRemindersPost**](RemindersApi.md#V1TrackingRemindersPost) | **Post** /v1/trackingReminders | Store a Tracking Reminder


# **V1TrackingReminderNotificationsGet**
> InlineResponse2002 V1TrackingReminderNotificationsGet($accessToken, $userId, $variableCategoryName, $createdAt, $updatedAt, $limit, $offset, $sort)

Get specific pending tracking reminders

Specfic pending reminder instances that still need to be tracked.  


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **variableCategoryName** | **string**| Limit tracking reminder notifications to a specific variable category | [optional] 
 **createdAt** | **string**| When the record was first created. Use UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot;  datetime format. Time zone should be UTC and not local. | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot;  datetime format. Time zone should be UTC and not local. | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TrackingReminderNotificationsSkipPost**
> CommonResponse V1TrackingReminderNotificationsSkipPost($body, $accessToken, $userId)

Skip a pending tracking reminder

Deletes the pending tracking reminder


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TrackingReminderNotificationSkip**](TrackingReminderNotificationSkip.md)| Id of the pending reminder to be skipped or deleted | 
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

# **V1TrackingReminderNotificationsSnoozePost**
> CommonResponse V1TrackingReminderNotificationsSnoozePost($body, $accessToken, $userId)

Snooze a pending tracking reminder

Changes the reminder time to now plus one hour


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TrackingReminderNotificationSnooze**](TrackingReminderNotificationSnooze.md)| Id of the pending reminder to be snoozed | 
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

# **V1TrackingReminderNotificationsTrackPost**
> CommonResponse V1TrackingReminderNotificationsTrackPost($body, $accessToken, $userId)

Track a pending tracking reminder

Adds the default measurement for the pending tracking reminder with the reminder time as the measurment start time


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TrackingReminderNotificationTrack**](TrackingReminderNotificationTrack.md)| Id of the pending reminder to be tracked | 
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

# **V1TrackingRemindersDeletePost**
> CommonResponse V1TrackingRemindersDeletePost($body, $accessToken, $userId)

Delete tracking reminder

Delete previously created tracking reminder


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TrackingReminderDelete**](TrackingReminderDelete.md)| Id of reminder to be deleted | 
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

# **V1TrackingRemindersGet**
> InlineResponse200 V1TrackingRemindersGet($accessToken, $userId, $variableCategoryName, $createdAt, $updatedAt, $limit, $offset, $sort)

Get repeating tracking reminder settings

Users can be reminded to track certain variables at a specified frequency with a default value.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **variableCategoryName** | **string**| Limit tracking reminders to a specific variable category | [optional] 
 **createdAt** | **string**| When the record was first created. Use UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot;  datetime format. Time zone should be UTC and not local. | [optional] 
 **updatedAt** | **string**| When the record was last updated. Use UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot;  datetime format. Time zone should be UTC and not local. | [optional] 
 **limit** | **int32**| The LIMIT is used to limit the number of results returned. So if you have 1000 results, but only want to the first 10, you would set this to 10 and offset to 0. The maximum limit is 200 records. | [optional] 
 **offset** | **int32**| OFFSET says to skip that many rows before beginning to return rows to the client. OFFSET 0 is the same as omitting the OFFSET clause. If both OFFSET and LIMIT appear, then OFFSET rows are skipped before starting to count the LIMIT rows that are returned. | [optional] 
 **sort** | **string**| Sort by given field. If the field is prefixed with &#39;-&#39;, it will sort in descending order. | [optional] 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1TrackingRemindersPost**
> InlineResponse2001 V1TrackingRemindersPost($accessToken, $userId, $body)

Store a Tracking Reminder

This is to enable users to create reminders to track a variable with a default value at a specified frequency


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessToken** | **string**| User&#39;s OAuth2 access token | [optional] 
 **userId** | **int32**| User&#39;s id | [optional] 
 **body** | [**TrackingReminder**](TrackingReminder.md)| TrackingReminder that should be stored | [optional] 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[access_token](../README.md#access_token), [quantimodo_oauth2](../README.md#quantimodo_oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

