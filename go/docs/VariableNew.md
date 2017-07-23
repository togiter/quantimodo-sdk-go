# VariableNew

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User-defined variable display name. | [default to null]
**Category** | **string** | Variable category like Mood, Sleep, Physical Activity, Treatment, Symptom, etc. | [default to null]
**Unit** | **string** | Abbreviated name of the default unit for the variable | [default to null]
**CombinationOperation** | **string** | Way to aggregate measurements over time. Options are \&quot;MEAN\&quot; or \&quot;SUM\&quot;. SUM should be used for things like minutes of exercise.  If you use MEAN for exercise, then a person might exercise more minutes in one day but add separate measurements that were smaller.  So when we are doing correlational analysis, we would think that the person exercised less that day even though they exercised more.  Conversely, we must use MEAN for things such as ratings which cannot be SUMMED. | [default to null]
**Parent** | **string** | Parent | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


