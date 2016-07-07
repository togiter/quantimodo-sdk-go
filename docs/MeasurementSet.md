# MeasurementSet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measurements** | [**[]ValueObject**](ValueObject.md) | Array of timestamps, values, and optional notes | [default to null]
**VariableName** | **string** | ORIGINAL name of the variable for which we are creating the measurement records | [default to null]
**SourceName** | **string** | Name of the application or device used to record the measurement values | [default to null]
**VariableCategoryName** | **string** | Variable category name | [optional] [default to null]
**CombinationOperation** | **string** | Way to aggregate measurements over time. Options are \&quot;MEAN\&quot; or \&quot;SUM\&quot;.  SUM should be used for things like minutes of exercise.  If you use MEAN for exercise, then a person might exercise more minutes in one day but add separate measurements that were smaller.  So when we are doing correlational analysis, we would think that the person exercised less that day even though they exercised more.  Conversely, we must use MEAN for things such as ratings which cannot be SUMMED. | [optional] [default to null]
**AbbreviatedUnitName** | **string** | Unit of measurement | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


