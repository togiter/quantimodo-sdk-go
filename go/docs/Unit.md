# Unit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unit name | [default to null]
**AbbreviatedName** | **string** | Unit abbreviation | [default to null]
**Category** | **string** | Unit category | [default to null]
**MinimumAllowedValue** | **float64** | The minimum allowed value for measurements. While you can record a value below this minimum, it will be excluded from the correlation analysis. | [optional] [default to null]
**MaximumAllowedValue** | **float64** | The maximum allowed value for measurements. While you can record a value above this maximum, it will be excluded from the correlation analysis. | [optional] [default to null]
**ConversionSteps** | [**[]ConversionStep**](ConversionStep.md) | Conversion steps list | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


