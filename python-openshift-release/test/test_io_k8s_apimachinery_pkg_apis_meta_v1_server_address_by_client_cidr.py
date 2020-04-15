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
from openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_server_address_by_client_cidr import IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR  # noqa: E501
from openshift_release.rest import ApiException

class TestIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR(unittest.TestCase):
    """IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_server_address_by_client_cidr.IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR()  # noqa: E501
        if include_optional :
            return IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR(
                client_cidr = '0', 
                server_address = '0'
            )
        else :
            return IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR(
                client_cidr = '0',
                server_address = '0',
        )

    def testIoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR(self):
        """Test IoK8sApimachineryPkgApisMetaV1ServerAddressByClientCIDR"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
