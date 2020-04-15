# openshift_release.ApisApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_api_versions**](ApisApi.md#get_api_versions) | **GET** /apis/ | 


# **get_api_versions**
> IoK8sApimachineryPkgApisMetaV1APIGroupList get_api_versions()



get available API versions

### Example

```python
from __future__ import print_function
import time
import openshift_release
from openshift_release.rest import ApiException
from pprint import pprint

# Enter a context with an instance of the API client
with openshift_release.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = openshift_release.ApisApi(api_client)
    
    try:
        api_response = api_instance.get_api_versions()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling ApisApi->get_api_versions: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IoK8sApimachineryPkgApisMetaV1APIGroupList**](IoK8sApimachineryPkgApisMetaV1APIGroupList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

