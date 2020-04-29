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
from openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_object_meta import IoK8sApimachineryPkgApisMetaV1ObjectMeta  # noqa: E501
from openshift_release.rest import ApiException

class TestIoK8sApimachineryPkgApisMetaV1ObjectMeta(unittest.TestCase):
    """IoK8sApimachineryPkgApisMetaV1ObjectMeta unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test IoK8sApimachineryPkgApisMetaV1ObjectMeta
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.io_k8s_apimachinery_pkg_apis_meta_v1_object_meta.IoK8sApimachineryPkgApisMetaV1ObjectMeta()  # noqa: E501
        if include_optional :
            return IoK8sApimachineryPkgApisMetaV1ObjectMeta(
                annotations = {
                    'key' : '0'
                    }, 
                cluster_name = '0', 
                creation_timestamp = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                deletion_grace_period_seconds = 56, 
                deletion_timestamp = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                finalizers = [
                    '0'
                    ], 
                generate_name = '0', 
                generation = 56, 
                labels = {
                    'key' : '0'
                    }, 
                managed_fields = [
                    openshift_release.models.io/k8s/apimachinery/pkg/apis/meta/v1/managed_fields_entry.io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry(
                        api_version = '0', 
                        fields_type = '0', 
                        fields_v1 = openshift_release.models.io/k8s/apimachinery/pkg/apis/meta/v1/fields_v1.io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1(), 
                        manager = '0', 
                        operation = '0', 
                        time = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), )
                    ], 
                name = '0', 
                namespace = '0', 
                owner_references = [
                    openshift_release.models.io/k8s/apimachinery/pkg/apis/meta/v1/owner_reference.io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference(
                        api_version = '0', 
                        block_owner_deletion = True, 
                        controller = True, 
                        kind = '0', 
                        name = '0', 
                        uid = '0', )
                    ], 
                resource_version = '0', 
                self_link = '0', 
                uid = '0'
            )
        else :
            return IoK8sApimachineryPkgApisMetaV1ObjectMeta(
        )

    def testIoK8sApimachineryPkgApisMetaV1ObjectMeta(self):
        """Test IoK8sApimachineryPkgApisMetaV1ObjectMeta"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()