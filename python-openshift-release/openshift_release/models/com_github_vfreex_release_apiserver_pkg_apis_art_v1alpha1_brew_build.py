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


class ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild(object):
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
        'instance': 'str',
        'name': 'str',
        'nvr': 'str',
        'release': 'str',
        'target': 'str',
        'task_number': 'int',
        'version': 'str'
    }

    attribute_map = {
        'instance': 'instance',
        'name': 'name',
        'nvr': 'nvr',
        'release': 'release',
        'target': 'target',
        'task_number': 'taskNumber',
        'version': 'version'
    }

    def __init__(self, instance=None, name=None, nvr=None, release=None, target=None, task_number=None, version=None, local_vars_configuration=None):  # noqa: E501
        """ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._instance = None
        self._name = None
        self._nvr = None
        self._release = None
        self._target = None
        self._task_number = None
        self._version = None
        self.discriminator = None

        if instance is not None:
            self.instance = instance
        self.name = name
        if nvr is not None:
            self.nvr = nvr
        self.release = release
        if target is not None:
            self.target = target
        if task_number is not None:
            self.task_number = task_number
        self.version = version

    @property
    def instance(self):
        """Gets the instance of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501


        :return: The instance of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :rtype: str
        """
        return self._instance

    @instance.setter
    def instance(self, instance):
        """Sets the instance of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.


        :param instance: The instance of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :type: str
        """

        self._instance = instance

    @property
    def name(self):
        """Gets the name of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501


        :return: The name of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.


        :param name: The name of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and name is None:  # noqa: E501
            raise ValueError("Invalid value for `name`, must not be `None`")  # noqa: E501

        self._name = name

    @property
    def nvr(self):
        """Gets the nvr of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501


        :return: The nvr of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :rtype: str
        """
        return self._nvr

    @nvr.setter
    def nvr(self, nvr):
        """Sets the nvr of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.


        :param nvr: The nvr of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :type: str
        """

        self._nvr = nvr

    @property
    def release(self):
        """Gets the release of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501


        :return: The release of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :rtype: str
        """
        return self._release

    @release.setter
    def release(self, release):
        """Sets the release of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.


        :param release: The release of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and release is None:  # noqa: E501
            raise ValueError("Invalid value for `release`, must not be `None`")  # noqa: E501

        self._release = release

    @property
    def target(self):
        """Gets the target of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501


        :return: The target of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :rtype: str
        """
        return self._target

    @target.setter
    def target(self, target):
        """Sets the target of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.


        :param target: The target of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :type: str
        """

        self._target = target

    @property
    def task_number(self):
        """Gets the task_number of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501


        :return: The task_number of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :rtype: int
        """
        return self._task_number

    @task_number.setter
    def task_number(self, task_number):
        """Sets the task_number of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.


        :param task_number: The task_number of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :type: int
        """

        self._task_number = task_number

    @property
    def version(self):
        """Gets the version of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501


        :return: The version of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :rtype: str
        """
        return self._version

    @version.setter
    def version(self, version):
        """Sets the version of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.


        :param version: The version of this ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild.  # noqa: E501
        :type: str
        """
        if self.local_vars_configuration.client_side_validation and version is None:  # noqa: E501
            raise ValueError("Invalid value for `version`, must not be `None`")  # noqa: E501

        self._version = version

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
        if not isinstance(other, ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, ComGithubVfreexReleaseApiserverPkgApisArtV1alpha1BrewBuild):
            return True

        return self.to_dict() != other.to_dict()