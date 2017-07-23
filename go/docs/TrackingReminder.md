# TrackingReminder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [optional] [default to null]
**ClientId** | **string** | clientId | [optional] [default to null]
**UserId** | **int32** | ID of User | [optional] [default to null]
**VariableId** | **int32** | Id for the variable to be tracked | [default to null]
**DefaultValue** | **float32** | Default value to use for the measurement when tracking | [default to null]
**ReminderStartTime** | **string** | Earliest time of day at which reminders should appear in UTC HH:MM:SS format | [optional] [default to null]
**ReminderEndTime** | **string** | Latest time of day at which reminders should appear in UTC HH:MM:SS format | [optional] [default to null]
**ReminderSound** | **string** | String identifier for the sound to accompany the reminder | [optional] [default to null]
**ReminderFrequency** | **int32** | Number of seconds between one reminder and the next | [default to null]
**PopUp** | **bool** | True if the reminders should appear as a popup notification | [optional] [default to null]
**Sms** | **bool** | True if the reminders should be delivered via SMS | [optional] [default to null]
**Email** | **bool** | True if the reminders should be delivered via email | [optional] [default to null]
**NotificationBar** | **bool** | True if the reminders should appear in the notification bar | [optional] [default to null]
**LatestTrackingReminderNotificationReminderTime** | [**time.Time**](time.Time.md) | UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot;  timestamp for the reminder time of the latest tracking reminder notification that has been pre-emptively generated in the database | [optional] [default to null]
**LastTracked** | [**time.Time**](time.Time.md) | UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot;  timestamp for the last time a measurement was received for this user and variable | [optional] [default to null]
**StartTrackingDate** | **string** | Earliest date on which the user should be reminded to track in YYYY-MM-DD format | [optional] [default to null]
**StopTrackingDate** | **string** | Latest date on which the user should be reminded to track in YYYY-MM-DD format | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the record in the database was last updated. Use UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot;  datetime format. Time zone should be UTC and not local. | [optional] [default to null]
**VariableName** | **string** | Name of the variable to be used when sending measurements | [optional] [default to null]
**VariableCategoryName** | **string** | Name of the variable category to be used when sending measurements | [optional] [default to null]
**UnitAbbreviatedName** | **string** | Abbreviated name of the unit to be used when sending measurements | [optional] [default to null]
**CombinationOperation** | **string** | The way multiple measurements are aggregated over time | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


