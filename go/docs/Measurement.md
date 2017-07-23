# Measurement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariableName** | **string** | Name of the variable for which we are creating the measurement records | [default to null]
**SourceName** | **string** | Application or device used to record the measurement values | [default to null]
**StartTimeString** | **string** | Start Time for the measurement event in UTC ISO 8601 \&quot;YYYY-MM-DDThh:mm:ss\&quot; | [default to null]
**StartTimeEpoch** | **int32** | Seconds between the start of the event measured and 1970 (Unix timestamp) | [optional] [default to null]
**HumanTime** | [**HumanTime**](HumanTime.md) |  | [optional] [default to null]
**Value** | **float64** | Converted measurement value in requested unit | [default to null]
**OriginalValue** | **int32** | Original value as originally submitted | [optional] [default to null]
**OriginalunitAbbreviatedName** | **string** | Original Unit of measurement as originally submitted | [optional] [default to null]
**UnitAbbreviatedName** | **string** | Abbreviated name for the unit of measurement | [default to null]
**Note** | **string** | Note of measurement | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


