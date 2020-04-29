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
from openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_managed_fields_entry import IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry  # noqa: E501
from openshift_release.rest import ApiException

class TestIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry(unittest.TestCase):
    """IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_managed_fields_entry.IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry()  # noqa: E501
        if include_optional :
            return IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry(
                api_version = '0', 
                fields_type = '0', 
                fields_v1 = openshift_release.models.io/k8s/apimachinery/pkg/apis/meta/v1/fields_v1.io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1(), 
                manager = '0', 
                operation = '0', 
                time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else :
            return IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry(
        )

    def testIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry(self):
        """Test IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()