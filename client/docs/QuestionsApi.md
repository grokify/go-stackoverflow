# \QuestionsApi

All URIs are relative to *https://api.stackexchange.com/2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeaturedQuestions**](QuestionsApi.md#GetFeaturedQuestions) | **Get** /questions/featured | 
[**GetQuestions**](QuestionsApi.md#GetQuestions) | **Get** /questions | 
[**GetQuestionsWithoutAnswers**](QuestionsApi.md#GetQuestionsWithoutAnswers) | **Get** /questions/no-answers | 
[**GetUnansweredQuestions**](QuestionsApi.md#GetUnansweredQuestions) | **Get** /questions/unanswered | 



## GetFeaturedQuestions

> QuestionsResponse GetFeaturedQuestions(ctx, site, optional)


Gets all the questions on the site. This method allows you make fairly flexible queries across the entire corpus of questions on a site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string**|  | [default to stackoverflow]
 **optional** | ***GetFeaturedQuestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFeaturedQuestionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromdate** | **optional.Int32**|  | 
 **todate** | **optional.Int32**|  | 
 **min** | **optional.Int32**|  | 
 **max** | **optional.Int32**|  | 
 **tagged** | **optional.String**|  | 
 **order** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**QuestionsResponse**](QuestionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuestions

> QuestionsResponse GetQuestions(ctx, site, optional)


Gets all the questions on the site. This method allows you make fairly flexible queries across the entire corpus of questions on a site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string**|  | [default to stackoverflow]
 **optional** | ***GetQuestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetQuestionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromdate** | **optional.Int32**|  | 
 **todate** | **optional.Int32**|  | 
 **min** | **optional.Int32**|  | 
 **max** | **optional.Int32**|  | 
 **tagged** | **optional.String**|  | 
 **order** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**QuestionsResponse**](QuestionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuestionsWithoutAnswers

> QuestionsResponse GetQuestionsWithoutAnswers(ctx, site, optional)


Returns questions which have received no answers. Compare with /questions/unanswered which merely returns questions that the sites consider insufficiently well answered.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string**|  | [default to stackoverflow]
 **optional** | ***GetQuestionsWithoutAnswersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetQuestionsWithoutAnswersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromdate** | **optional.Int32**|  | 
 **todate** | **optional.Int32**|  | 
 **min** | **optional.Int32**|  | 
 **max** | **optional.Int32**|  | 
 **tagged** | **optional.String**|  | 
 **order** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**QuestionsResponse**](QuestionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnansweredQuestions

> QuestionsResponse GetUnansweredQuestions(ctx, site, optional)


Gets all the questions on the site. This method allows you make fairly flexible queries across the entire corpus of questions on a site.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**site** | **string**|  | [default to stackoverflow]
 **optional** | ***GetUnansweredQuestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUnansweredQuestionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromdate** | **optional.Int32**|  | 
 **todate** | **optional.Int32**|  | 
 **min** | **optional.Int32**|  | 
 **max** | **optional.Int32**|  | 
 **tagged** | **optional.String**|  | 
 **order** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**QuestionsResponse**](QuestionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

