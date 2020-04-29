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
from openshift_release.models.com_github_vfreex_release_apiserver_pkg_apis_art_v1alpha1_advisory_spec import ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec  # noqa: E501
from openshift_release.rest import ApiException

class TestComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec(unittest.TestCase):
    """ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.com_github_vfreex_release_apiserver_pkg_apis_art_v1alpha1_advisory_spec.ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec()  # noqa: E501
        if include_optional :
            return ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec(
                impetus = '0', 
                instance = '0', 
                number = 56, 
                release_name = '0'
            )
        else :
            return ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec(
        )

    def testComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec(self):
        """Test ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1AdvisorySpec"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()