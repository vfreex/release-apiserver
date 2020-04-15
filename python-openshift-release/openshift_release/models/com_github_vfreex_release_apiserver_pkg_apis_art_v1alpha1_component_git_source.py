# coding: utf-8

"""
    Api

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: v0
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from openshift_release.configuration import Configuration


class ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'fallback_ref': 'str',
        'path': 'str',
        'ref': 'str',
        'url': 'str'
    }

    attribute_map = {
        'fallback_ref': 'fallbackRef',
        'path': 'path',
        'ref': 'ref',
        'url': 'url'
    }

    def __init__(self, fallback_ref=None, path=None, ref=None, url=None, local_vars_configuration=None):  # noqa: E501
        """ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._fallback_ref = None
        self._path = None
        self._ref = None
        self._url = None
        self.discriminator = None

        if fallback_ref is not None:
            self.fallback_ref = fallback_ref
        if path is not None:
            self.path = path
        self.ref = ref
        self.url = url

    @property
    def fallback_ref(self):
        """Gets the fallback_ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501


        :return: The fallback_ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :rtype: str
        """
        return self._fallback_ref

    @fallback_ref.setter
    def fallback_ref(self, fallback_ref):
        """Sets the fallback_ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.


        :param fallback_ref: The fallback_ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :type: str
        """

        self._fallback_ref = fallback_ref

    @property
    def path(self):
        """Gets the path of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501


        :return: The path of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :rtype: str
        """
        return self._path

    @path.setter
    def path(self, path):
        """Sets the path of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.


        :param path: The path of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :type: str
        """

        self._path = path

    @property
    def ref(self):
        """Gets the ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501


        :return: The ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :rtype: str
        """
        return self._ref

    @ref.setter
    def ref(self, ref):
        """Sets the ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.


        :param ref: The ref of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and ref is None:  # noqa: E501
            raise ValueError("Invalid value for `ref`, must not be `None`")  # noqa: E501

        self._ref = ref

    @property
    def url(self):
        """Gets the url of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501


        :return: The url of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :rtype: str
        """
        return self._url

    @url.setter
    def url(self, url):
        """Sets the url of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.


        :param url: The url of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and url is None:  # noqa: E501
            raise ValueError("Invalid value for `url`, must not be `None`")  # noqa: E501

        self._url = url

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1ComponentGitSource):
            return True

        return self.to_dict() != other.to_dict()
