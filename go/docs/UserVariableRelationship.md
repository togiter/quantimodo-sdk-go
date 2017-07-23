# UserVariableRelationship

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id | [optional] [default to null]
**ConfidenceLevel** | **string** | Our confidence that a consistent predictive relationship exists based on the amount of evidence, reproducibility, and other factors | [default to null]
**ConfidenceScore** | **float32** | A quantitative representation of our confidence that a consistent predictive relationship exists based on the amount of evidence, reproducibility, and other factors | [default to null]
**Direction** | **string** | Direction is positive if higher predictor values generally precede higher outcome values. Direction is negative if higher predictor values generally precede lower outcome values. | [default to null]
**DurationOfAction** | **int32** | The amount of time over which a predictor/stimulus event can exert an observable influence on an outcome variable’s value. For instance, aspirin (stimulus/predictor) typically decreases headache severity for approximately four hours (duration of action) following the onset delay. | [default to null]
**ErrorMessage** | **string** | error_message | [optional] [default to null]
**OnsetDelay** | **int32** | The amount of time in seconds that elapses after the predictor/stimulus event before the outcome as perceived by a self-tracker is known as the “onset delay”. For example, the “onset delay” between the time a person takes an aspirin (predictor/stimulus event) and the time a person perceives a change in their headache severity (outcome) is approximately 30 minutes. | [optional] [default to null]
**OutcomeVariableId** | **int32** | Variable ID for the outcome variable | [default to null]
**PredictorVariableId** | **int32** | Variable ID for the predictor variable | [default to null]
**PredictorUnitId** | **int32** | ID for default unit of the predictor variable | [default to null]
**SinnRank** | **float32** | A value representative of the relevance of this predictor relative to other predictors of this outcome.  Usually used for relevancy sorting. | [default to null]
**StrengthLevel** | **string** | Can be weak, medium, or strong based on the size of the effect which the predictor appears to have on the outcome relative to other variable relationship strength scores. | [default to null]
**StrengthScore** | **float32** | A value represented to the size of the effect which the predictor appears to have on the outcome. | [default to null]
**UserId** | **int32** | userId | [optional] [default to null]
**Vote** | **string** | vote | [optional] [default to null]
**ValuePredictingHighOutcome** | **float32** | Value for the predictor variable (in it&#39;s default unit) which typically precedes an above average outcome value | [default to null]
**ValuePredictingLowOutcome** | **float32** | Value for the predictor variable (in it&#39;s default unit) which typically precedes a below average outcome value | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


