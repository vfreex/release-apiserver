# coding: utf-8

"""
    Api

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: v0
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest

import openshift_release
from openshift_release.api.version_api import VersionApi  # noqa: E501
from openshift_release.rest import ApiException


class TestVersionApi(unittest.TestCase):
    """VersionApi unit test stubs"""

    def setUp(self):
        self.api = openshift_release.api.version_api.VersionApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_get_code_version(self):
        """Test case for get_code_version

        """
        pass


if __name__ == '__main__':
    unittest.main()
