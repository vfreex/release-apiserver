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
from openshift_release.models.com_github_vfreex_release_apiserver_pkg_apis_art_v1alpha1_build_spec import ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec  # noqa: E501
from openshift_release.rest import ApiException

class TestComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec(unittest.TestCase):
    """ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = openshift_release.models.com_github_vfreex_release_apiserver_pkg_apis_art_v1alpha1_build_spec.ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec()  # noqa: E501
        if include_optional :
            return ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec(
                brew = openshift_release.models.com/github/vfreex/release_apiserver/pkg/apis/art/v1alpha1/brew_build.com.github.vfreex.release-apiserver.pkg.apis.art.v1alpha1.BrewBuild(
                    instance = '0', 
                    name = '0', 
                    nvr = '0', 
                    release = '0', 
                    target = '0', 
                    task_number = 56, 
                    version = '0', ), 
                jenkins = openshift_release.models.com/github/vfreex/release_apiserver/pkg/apis/art/v1alpha1/build_jenkins_info.com.github.vfreex.release-apiserver.pkg.apis.art.v1alpha1.BuildJenkinsInfo(
                    build_url = '0', ), 
                type = '0'
            )
        else :
            return ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec(
                brew = openshift_release.models.com/github/vfreex/release_apiserver/pkg/apis/art/v1alpha1/brew_build.com.github.vfreex.release-apiserver.pkg.apis.art.v1alpha1.BrewBuild(
                    instance = '0', 
                    name = '0', 
                    nvr = '0', 
                    release = '0', 
                    target = '0', 
                    task_number = 56, 
                    version = '0', ),
                type = '0',
        )

    def testComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec(self):
        """Test ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()