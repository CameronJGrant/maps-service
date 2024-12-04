# \SubmissionsApi

All URIs are relative to *https://submissions.strafes.net/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**action_submission_publish**](SubmissionsApi.md#action_submission_publish) | **PATCH** /submissions/{SubmissionID}/status/publish | Role Validator changes status from Publishing -> Published
[**action_submission_reject**](SubmissionsApi.md#action_submission_reject) | **PATCH** /submissions/{SubmissionID}/status/reject | Role Reviewer changes status from Submitted -> Rejected
[**action_submission_request_changes**](SubmissionsApi.md#action_submission_request_changes) | **PATCH** /submissions/{SubmissionID}/status/request-changes | Role Reviewer changes status from Validated|Accepted|Submitted -> ChangesRequested
[**action_submission_revoke**](SubmissionsApi.md#action_submission_revoke) | **PATCH** /submissions/{SubmissionID}/status/revoke | Role Submitter changes status from Submitted|ChangesRequested -> UnderConstruction
[**action_submission_submit**](SubmissionsApi.md#action_submission_submit) | **PATCH** /submissions/{SubmissionID}/status/submit | Role Submitter changes status from UnderConstruction|ChangesRequested -> Submitted
[**action_submission_trigger_publish**](SubmissionsApi.md#action_submission_trigger_publish) | **PATCH** /submissions/{SubmissionID}/status/trigger-publish | Role Admin changes status from Validated -> Publishing
[**action_submission_trigger_validate**](SubmissionsApi.md#action_submission_trigger_validate) | **PATCH** /submissions/{SubmissionID}/status/trigger-validate | Role Reviewer triggers validation and changes status from Submitted|Accepted -> Validating
[**action_submission_validate**](SubmissionsApi.md#action_submission_validate) | **PATCH** /submissions/{SubmissionID}/status/validate | Role Validator changes status from Validating -> Validated
[**create_submission**](SubmissionsApi.md#create_submission) | **POST** /submissions | Create new submission
[**get_submission**](SubmissionsApi.md#get_submission) | **GET** /submissions/{SubmissionID} | Retrieve map with ID
[**list_submissions**](SubmissionsApi.md#list_submissions) | **GET** /submissions | Get list of submissions
[**patch_submission_completed**](SubmissionsApi.md#patch_submission_completed) | **PATCH** /submissions/{SubmissionID}/completed | Retrieve map with ID
[**patch_submission_model**](SubmissionsApi.md#patch_submission_model) | **PATCH** /submissions/{SubmissionID}/model | Update model following role restrictions



## action_submission_publish

> action_submission_publish(submission_id)
Role Validator changes status from Publishing -> Published

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## action_submission_reject

> action_submission_reject(submission_id)
Role Reviewer changes status from Submitted -> Rejected

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## action_submission_request_changes

> action_submission_request_changes(submission_id)
Role Reviewer changes status from Validated|Accepted|Submitted -> ChangesRequested

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## action_submission_revoke

> action_submission_revoke(submission_id)
Role Submitter changes status from Submitted|ChangesRequested -> UnderConstruction

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## action_submission_submit

> action_submission_submit(submission_id)
Role Submitter changes status from UnderConstruction|ChangesRequested -> Submitted

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## action_submission_trigger_publish

> action_submission_trigger_publish(submission_id)
Role Admin changes status from Validated -> Publishing

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## action_submission_trigger_validate

> action_submission_trigger_validate(submission_id)
Role Reviewer triggers validation and changes status from Submitted|Accepted -> Validating

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## action_submission_validate

> action_submission_validate(submission_id)
Role Validator changes status from Validating -> Validated

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## create_submission

> models::Id create_submission(submission_create)
Create new submission

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_create** | Option<[**SubmissionCreate**](SubmissionCreate.md)> |  |  |

### Return type

[**models::Id**](Id.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_submission

> models::Submission get_submission(submission_id)
Retrieve map with ID

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

[**models::Submission**](Submission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_submissions

> Vec<models::Submission> list_submissions(page, filter)
Get list of submissions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page** | [**Pagination**](.md) |  | [required] |
**filter** | Option<[**SubmissionFilter**](.md)> |  |  |

### Return type

[**Vec<models::Submission>**](Submission.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## patch_submission_completed

> patch_submission_completed(submission_id)
Retrieve map with ID

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## patch_submission_model

> patch_submission_model(submission_id, model_id, version_id)
Update model following role restrictions

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**submission_id** | **i64** |  | [required] |
**model_id** | **i64** |  | [required] |
**version_id** | **i64** |  | [required] |

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

