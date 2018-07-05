# \UsersApi

All URIs are relative to *https://api.stackexchange.com/2.2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMe**](UsersApi.md#GetMe) | **Get** /me | 
[**GetMyReputation**](UsersApi.md#GetMyReputation) | **Get** /me/reputation | 
[**GetMyReputationHistory**](UsersApi.md#GetMyReputationHistory) | **Get** /me/reputation-history | 
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /users | 
[**GetUsersReputation**](UsersApi.md#GetUsersReputation) | **Get** /users/{userIds}/reputation | 
[**GetUsersReputationHistory**](UsersApi.md#GetUsersReputationHistory) | **Get** /users/{userIds}/reputation-history | 


# **GetMe**
> UsersResponse GetMe(ctx, site)


Returns the user associated with the passed access_token.  This method returns a [user](https://api.stackexchange.com/docs/types/user).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | [default to stackoverflow]

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyReputation**
> ReputationResponse GetMyReputation(ctx, )


Returns the reputation changed for the user associated with the given access_token. This method returns a list of [reputation changes](https://api.stackexchange.com/docs/types/reputation).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReputationResponse**](ReputationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMyReputationHistory**
> ReputationHistoryResponse GetMyReputationHistory(ctx, optional)


Returns user's public reputation history.  This method returns a list of [reputation_history](https://api.stackexchange.com/docs/types/reputation-history).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMyReputationHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMyReputationHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**ReputationHistoryResponse**](ReputationHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> UsersResponse GetUsers(ctx, site, optional)


Returns all users on a site.  This method returns a list of users.  The sorts accepted by this method operate on the following fields of the user object:  reputation – reputation creation – creation_date name – display_name modified – last_modified_date  reputation is the default sort.  It is possible to create moderately complex queries using sort, min, max, fromdate, and todate. The `inname` parameter lets consumers filter the results down to just those users with a certain substring in their display name. For example, `inname=kevin` will return all users with both users named simply \"Kevin\" or those with Kevin as one of (or part of) their names; such as \"Kevin Montrose\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **site** | **string**|  | [default to stackoverflow]
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inname** | **optional.String**|  | 
 **fromdate** | **optional.Int32**|  | 
 **todate** | **optional.Int32**|  | 
 **min** | **optional.Int32**|  | 
 **max** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **order** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersReputation**
> ReputationResponse GetUsersReputation(ctx, userIds, site, optional)


Gets a subset of the reputation changes for users in {ids}. Reputation changes are intentionally scrubbed of some data to make it difficult to correlate votes on particular posts with user reputation changes. That being said, this method returns enough data for reasonable display of reputation trends. {ids} can contain up to 100 semicolon delimited ids. To find ids programmatically look for user_id on user or shallow_user objects. This method returns a list of reputation objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userIds** | **string**|  | 
  **site** | **string**|  | [default to stackoverflow]
 **optional** | ***GetUsersReputationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsersReputationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromdate** | **optional.Int32**|  | 
 **todate** | **optional.Int32**|  | 
 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**ReputationResponse**](ReputationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsersReputationHistory**
> ReputationHistoryResponse GetUsersReputationHistory(ctx, userIds, site, optional)


Returns users' public reputation history. This method returns a list of reputation_history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userIds** | **string**|  | 
  **site** | **string**|  | [default to stackoverflow]
 **optional** | ***GetUsersReputationHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUsersReputationHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**|  | 
 **pagesize** | **optional.Int32**|  | 

### Return type

[**ReputationHistoryResponse**](ReputationHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

