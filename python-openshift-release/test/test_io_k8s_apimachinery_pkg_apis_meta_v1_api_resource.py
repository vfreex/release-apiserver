# coding: utf-8

"""
    Api

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: v0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import openshift_release
from openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_api_resource import IoK8sApimachineryPkgApisMetaV1APIResource  # noqa: E501
from openshift_release.rest import ApiException

class TestIoK8sApimachineryPkgApisMetaV1APIResource(unittest.TestCase):
    """IoK8sApimachineryPkgApisMetaV1APIResource unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test IoK8sApimachineryPkgApisMetaV1APIResource
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_api_resource.IoK8sApimachineryPkgApisMetaV1APIResource()  # noqa: E501
        if include_optional :
            return IoK8sApimachineryPkgApisMetaV1APIResource(
                categories = [
                    '0'
                    ], 
                group = '0', 
                kind = '0', 
                name = '0', 
                namespaced = True, 
                short_names = [
                    '0'
                    ], 
                singular_name = '0', 
                storage_version_hash = '0', 
                verbs = [
                    '0'
                    ], 
                version = '0'
            )
        else :
            return IoK8sApimachineryPkgApisMetaV1APIResource(
                kind = '0',
                name = '0',
                namespaced = True,
                singular_name = '0',
                verbs = [
                    '0'
                    ],
        )

    def testIoK8sApimachineryPkgApisMetaV1APIResource(self):
        """Test IoK8sApimachineryPkgApisMetaV1APIResource"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
