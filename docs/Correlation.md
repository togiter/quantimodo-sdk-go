# Correlation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationCoefficient** | **float32** | Pearson correlation coefficient between cause and effect measurements | [default to null]
**Cause** | **string** | ORIGINAL variable name of the cause variable for which the user desires correlations. | [default to null]
**OriginalCause** | **string** | original name of the cause. | [optional] [default to null]
**Effect** | **string** | ORIGINAL variable name of the effect variable for which the user desires correlations. | [default to null]
**OriginalEffect** | **string** | effect variable original name. | [optional] [default to null]
**OnsetDelay** | **float64** | User estimated or default time after cause measurement before a perceivable effect is observed | [default to null]
**DurationOfAction** | **float32** | Time over which the cause is expected to produce a perceivable effect following the onset delay | [default to null]
**NumberOfPairs** | **float32** | Number of points that went into the correlation calculation | [default to null]
**EffectSize** | **string** | Magnitude of the effects of a cause indicating whether it&#39;s practically meaningful. | [optional] [default to null]
**StatisticalSignificance** | **string** | A function of the effect size and sample size | [optional] [default to null]
**Timestamp** | **float32** | Time at which correlation was calculated | [default to null]
**ReverseCorrelation** | **float32** | Correlation when cause and effect are reversed. For any causal relationship, the forward correlation should exceed the reverse correlation. | [optional] [default to null]
**CausalityFactor** | **float32** |  | [optional] [default to null]
**CauseCategory** | **string** | Variable category of the cause variable. | [optional] [default to null]
**EffectCategory** | **string** | Variable category of the effect variable. | [optional] [default to null]
**ValuePredictingHighOutcome** | **float32** | cause value that predicts an above average effect value (in default unit for cause variable) | [optional] [default to null]
**ValuePredictingLowOutcome** | **float32** | cause value that predicts a below average effect value (in default unit for cause variable) | [optional] [default to null]
**OptimalPearsonProduct** | **float32** | Optimal Pearson Product | [optional] [default to null]
**AverageVote** | **float32** | Average Vote | [optional] [default to null]
**UserVote** | **float32** | User Vote | [optional] [default to null]
**CauseUnit** | **string** | Unit of the predictor variable | [optional] [default to null]
**CauseUnitId** | **int32** | Unit Id of the predictor variable | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


