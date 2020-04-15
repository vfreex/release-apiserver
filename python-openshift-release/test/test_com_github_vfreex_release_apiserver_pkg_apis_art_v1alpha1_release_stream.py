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
from openshift_release.models.com_github_vfreex_release_apiserver_pkg_apis_art_v1alpha1_release_stream import ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream  # noqa: E501
from openshift_release.rest import ApiException

class TestComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream(unittest.TestCase):
    """ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.com_github_vfreex_release_apiserver_pkg_apis_art_v1alpha1_release_stream.ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream()  # noqa: E501
        if include_optional :
            return ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream(
                api_version = '0', 
                kind = '0', 
                metadata = openshift_release.models.io/k8s/apimachinery/pkg/apis/meta/v1/object_meta.io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta(
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
                    uid = '0', ), 
                spec = openshift_release.models.com/github/vfreex/release_apiserver/pkg/apis/art/v1alpha1/release_stream_spec.com.github.vfreex.release-apiserver.pkg.apis.art.v1alpha1.ReleaseStreamSpec(
                    ocp_build_data = openshift_release.models.com/github/vfreex/release_apiserver/pkg/apis/art/v1alpha1/release_stream_ocp_build_data.com.github.vfreex.release-apiserver.pkg.apis.art.v1alpha1.ReleaseStreamOcpBuildData(
                        git = openshift_release.models.com/github/vfreex/release_apiserver/pkg/apis/art/v1alpha1/ocp_build_data_git_source.com.github.vfreex.release-apiserver.pkg.apis.art.v1alpha1.OcpBuildDataGitSource(
                            path = '0', 
                            ref = '0', 
                            url = '0', ), ), ), 
                status = None
            )
        else :
            return ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream(
        )

    def testComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream(self):
        """Test ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ReleaseStream"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
