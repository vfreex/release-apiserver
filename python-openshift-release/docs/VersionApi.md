# openshift_release.VersionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_code_version**](VersionApi.md#get_code_version) | **GET** /version/ | 


# **get_code_version**
> IoK8sApimachineryPkgVersionInfo get_code_version()



get the code version

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
    api_instance = openshift_release.VersionApi(api_client)
    
    try:
        api_response = api_instance.get_code_version()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VersionApi->get_code_version: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**IoK8sApimachineryPkgVersionInfo**](IoK8sApimachineryPkgVersionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

