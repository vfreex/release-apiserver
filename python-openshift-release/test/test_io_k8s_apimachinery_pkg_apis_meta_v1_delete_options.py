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
from openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_delete_options import IoK8sApimachineryPkgApisMetaV1DeleteOptions  # noqa: E501
from openshift_release.rest import ApiException

class TestIoK8sApimachineryPkgApisMetaV1DeleteOptions(unittest.TestCase):
    """IoK8sApimachineryPkgApisMetaV1DeleteOptions unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test IoK8sApimachineryPkgApisMetaV1DeleteOptions
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_delete_options.IoK8sApimachineryPkgApisMetaV1DeleteOptions()  # noqa: E501
        if include_optional :
            return IoK8sApimachineryPkgApisMetaV1DeleteOptions(
                api_version = '0', 
                dry_run = [
                    '0'
                    ], 
                grace_period_seconds = 56, 
                kind = '0', 
                orphan_dependents = True, 
                preconditions = openshift_release.models.io/k8s/apimachinery/pkg/apis/meta/v1/preconditions.io.k8s.apimachinery.pkg.apis.meta.v1.Preconditions(
                    resource_version = '0', 
                    uid = '0', ), 
                propagation_policy = '0'
            )
        else :
            return IoK8sApimachineryPkgApisMetaV1DeleteOptions(
        )

    def testIoK8sApimachineryPkgApisMetaV1DeleteOptions(self):
        """Test IoK8sApimachineryPkgApisMetaV1DeleteOptions"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
