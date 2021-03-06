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


class ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec(object):
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
        'brew': 'ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild',
        'jenkins': 'ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildJenkinsInfo',
        'type': 'str'
    }

    attribute_map = {
        'brew': 'brew',
        'jenkins': 'jenkins',
        'type': 'type'
    }

    def __init__(self, brew=None, jenkins=None, type=None, local_vars_configuration=None):  # noqa: E501
        """ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._brew = None
        self._jenkins = None
        self._type = None
        self.discriminator = None

        self.brew = brew
        if jenkins is not None:
            self.jenkins = jenkins
        self.type = type

    @property
    def brew(self):
        """Gets the brew of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501


        :return: The brew of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501
        :rtype: ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild
        """
        return self._brew

    @brew.setter
    def brew(self, brew):
        """Sets the brew of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.


        :param brew: The brew of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501
        :type: ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild
        """
        if self.local_vars_configuration.client_side_validation and brew is None:  # noqa: E501
            raise ValueError("Invalid value for `brew`, must not be `None`")  # noqa: E501

        self._brew = brew

    @property
    def jenkins(self):
        """Gets the jenkins of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501


        :return: The jenkins of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501
        :rtype: ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildJenkinsInfo
        """
        return self._jenkins

    @jenkins.setter
    def jenkins(self, jenkins):
        """Sets the jenkins of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.


        :param jenkins: The jenkins of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501
        :type: ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildJenkinsInfo
        """

        self._jenkins = jenkins

    @property
    def type(self):
        """Gets the type of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501


        :return: The type of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501
        :rtype: str
        """
        return self._type

    @type.setter
    def type(self, type):
        """Sets the type of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.


        :param type: The type of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and type is None:  # noqa: E501
            raise ValueError("Invalid value for `type`, must not be `None`")  # noqa: E501

        self._type = type

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
        if not isinstance(other, ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BuildSpec):
            return True

        return self.to_dict() != other.to_dict()
