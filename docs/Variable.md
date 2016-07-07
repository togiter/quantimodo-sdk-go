# Variable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Variable ID | [optional] [default to null]
**Name** | **string** | User-defined variable display name. | [default to null]
**OriginalName** | **string** | Name used when the variable was originally created in the &#x60;variables&#x60; table. | [default to null]
**Category** | **string** | Variable category like Mood, Sleep, Physical Activity, Treatment, Symptom, etc. | [default to null]
**AbbreviatedUnitName** | **string** | Abbreviated name of the default unit for the variable | [default to null]
**AbbreviatedUnitId** | **int32** | Id of the default unit for the variable | [default to null]
**Sources** | **string** | Comma-separated list of source names to limit variables to those sources | [default to null]
**MinimumValue** | **float64** | Minimum reasonable value for this variable (uses default unit) | [default to null]
**MaximumValue** | **float64** | Maximum reasonable value for this variable (uses default unit) | [default to null]
**CombinationOperation** | **string** | Way to aggregate measurements over time. Options are \&quot;MEAN\&quot; or \&quot;SUM\&quot;.  SUM should be used for things like minutes of exercise.  If you use MEAN for exercise, then a person might exercise more minutes in one day but add separate measurements that were smaller.  So when we are doing correlational analysis, we would think that the person exercised less that day even though they exercised more.  Conversely, we must use MEAN for things such as ratings which cannot be SUMMED. | [default to null]
**FillingValue** | **float64** | Value for replacing null measurements | [default to null]
**JoinWith** | **string** | The Variable this Variable should be joined with. If the variable is joined with some other variable then it is not shown to user in the list of variables. | [default to null]
**JoinedVariables** | [**[]Variable**](Variable.md) | Array of Variables that are joined with this Variable | [default to null]
**Parent** | **int32** | Id of the parent variable if this variable has any parent | [default to null]
**SubVariables** | [**[]Variable**](Variable.md) | Array of Variables that are sub variables to this Variable | [default to null]
**OnsetDelay** | **int32** | How long it takes for a measurement in this variable to take effect | [default to null]
**DurationOfAction** | **int32** | How long the effect of a measurement in this variable lasts | [default to null]
**EarliestMeasurementTime** | **int32** | Earliest measurement time | [default to null]
**LatestMeasurementTime** | **int32** | Latest measurement time | [default to null]
**Updated** | **int32** | When this variable or its settings were last updated | [default to null]
**CauseOnly** | **int32** | A value of 1 indicates that this variable is generally a cause in a causal relationship.  An example of a causeOnly variable would be a variable such as Cloud Cover which would generally not be influenced by the behaviour of the user. | [default to null]
**NumberOfCorrelations** | **int32** | Number of correlations | [default to null]
**Outcome** | **int32** | Outcome variables (those with &#x60;outcome&#x60; &#x3D;&#x3D; 1) are variables for which a human would generally want to identify the influencing factors.  These include symptoms of illness, physique, mood, cognitive performance, etc.  Generally correlation calculations are only performed on outcome variables. | [default to null]
**MeasurementsAtLastAnalysis** | **int32** | The number of measurements that a given user had for this variable the last time a correlation calculation was performed. Generally correlation values are only updated once the current number of measurements for a variable is more than 10% greater than the measurementsAtLastAnalysis.  This avoids a computationally-demanding recalculation when there&#39;s not enough new data to make a significant difference in the correlation. | [default to null]
**NumberOfMeasurements** | **int32** | Number of measurements | [default to null]
**LastUnit** | **string** | Last unit | [default to null]
**LastValue** | **int32** | Last value | [default to null]
**MostCommonValue** | **int32** | Most common value | [default to null]
**MostCommonUnit** | **string** | Most common unit | [default to null]
**LastSource** | **int32** | Last source | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


