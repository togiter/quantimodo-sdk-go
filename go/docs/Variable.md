# Variable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Variable ID | [optional] [default to null]
**Name** | **string** | User-defined variable display name. | [default to null]
**Category** | **string** | Variable category like Mood, Sleep, Physical Activity, Treatment, Symptom, etc. | [default to null]
**UnitAbbreviatedName** | **string** | Abbreviated name of the default unit for the variable | [default to null]
**AbbreviatedUnitId** | **int32** | Id of the default unit for the variable | [default to null]
**Sources** | **string** | Comma-separated list of source names to limit variables to those sources | [default to null]
**MinimumAllowedValue** | **float64** | The minimum allowed value for measurements. While you can record a value below this minimum, it will be excluded from the correlation analysis. | [default to null]
**MaximumAllowedValue** | **float64** | The maximum allowed value for measurements. While you can record a value above this maximum, it will be excluded from the correlation analysis. | [default to null]
**CombinationOperation** | **string** | Way to aggregate measurements over time. Options are \&quot;MEAN\&quot; or \&quot;SUM\&quot;. SUM should be used for things like minutes of exercise.  If you use MEAN for exercise, then a person might exercise more minutes in one day but add separate measurements that were smaller.  So when we are doing correlational analysis, we would think that the person exercised less that day even though they exercised more.  Conversely, we must use MEAN for things such as ratings which cannot be SUMMED. | [default to null]
**FillingValue** | **float64** | When it comes to analysis to determine the effects of this variable, knowing when it did not occur is as important as knowing when it did occur. For example, if you are tracking a medication, it is important to know when you did not take it, but you do not have to log zero values for all the days when you haven&#39;t taken it. Hence, you can specify a filling value (typically 0) to insert whenever data is missing. | [default to null]
**JoinWith** | **string** | The Variable this Variable should be joined with. If the variable is joined with some other variable then it is not shown to user in the list of variables. | [default to null]
**JoinedVariables** | [**[]Variable**](Variable.md) | Array of Variables that are joined with this Variable | [default to null]
**Parent** | **int32** | Id of the parent variable if this variable has any parent | [default to null]
**SubVariables** | [**[]Variable**](Variable.md) | Array of Variables that are sub variables to this Variable | [default to null]
**OnsetDelay** | **int32** | The amount of time in seconds that elapses after the predictor/stimulus event before the outcome as perceived by a self-tracker is known as the “onset delay”. For example, the “onset delay” between the time a person takes an aspirin (predictor/stimulus event) and the time a person perceives a change in their headache severity (outcome) is approximately 30 minutes. | [default to null]
**DurationOfAction** | **int32** | The amount of time over which a predictor/stimulus event can exert an observable influence on an outcome variable’s value. For instance, aspirin (stimulus/predictor) typically decreases headache severity for approximately four hours (duration of action) following the onset delay. | [default to null]
**EarliestMeasurementTime** | **int32** | Earliest measurement time | [default to null]
**LatestMeasurementTime** | **int32** | Latest measurement time | [default to null]
**Updated** | **int32** | When this variable or its settings were last updated | [default to null]
**CauseOnly** | **int32** | A value of 1 indicates that this variable is generally a cause in a causal relationship.  An example of a causeOnly variable would be a variable such as Cloud Cover which would generally not be influenced by the behaviour of the user. | [default to null]
**NumberOfCorrelations** | **int32** | Number of correlations | [default to null]
**Outcome** | **int32** | Outcome variables (those with &#x60;outcome&#x60; &#x3D;&#x3D; 1) are variables for which a human would generally want to identify the influencing factors. These include symptoms of illness, physique, mood, cognitive performance, etc.  Generally correlation calculations are only performed on outcome variables. | [default to null]
**RawMeasurementsAtLastAnalysis** | **int32** | The number of measurements that a given user had for this variable the last time a correlation calculation was performed. Generally correlation values are only updated once the current number of measurements for a variable is more than 10% greater than the rawMeasurementsAtLastAnalysis.  This avoids a computationally-demanding recalculation when there&#39;s not enough new data to make a significant difference in the correlation. | [default to null]
**NumberOfRawMeasurements** | **int32** | Number of measurements | [default to null]
**LastUnit** | **string** | Last unit | [default to null]
**LastValue** | **int32** | Last value | [default to null]
**MostCommonValue** | **int32** | Most common value | [default to null]
**MostCommonUnit** | **string** | Most common unit | [default to null]
**LastSource** | **int32** | Last source | [default to null]
**ImageUrl** | **string** |  | [optional] [default to null]
**IonIcon** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


