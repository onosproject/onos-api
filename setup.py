#!/usr/bin/env python3

# Copyright 2021-present Open Networking Foundation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from setuptools import setup, find_packages


PACKAGES = find_packages("python")


def readme():
    with open("./README.md") as f:
        return f.read()


def version():
    with open("VERSION") as f:
        return f.read()


setup(
    name="onos-api",
    version=version(),
    description="ONOS Python API",
    long_description=readme(),
    long_description_content_type="text/markdown",
    author="Open Networking Foundation (ONF) and Partners",
    author_email="support@opennetworking.org",
    packages=[pkg.replace("onos", "onos_api") for pkg in PACKAGES],
    package_dir={
        pkg.replace("onos", "onos_api"): f"python/{pkg}".replace(".", "/")
        for pkg in PACKAGES
    },
    package_data={"onos_api": ["py.typed"]},
    keywords=["onos-api", "python", "protobuf", "gnmi"],
    license="Apache License v2.0",
    url="https://github.com/onosproject/onos-api/",
    download_url="https://github.com/onosproject/onos-api/",
    classifiers=[
        "Development Status :: 4 - Beta",
        "Environment :: Other Environment",
        "Intended Audience :: Developers",
        "Operating System :: OS Independent",
        "Programming Language :: Python",
        "Programming Language :: Python :: 3",
        "Topic :: Utilities",
        "License :: OSI Approved :: Apache Software License",
        "Typing :: Typed",
    ],
    install_requires=[
        "betterproto>=2.0.0b3,<3",
        "gnmi-proto>=0.1.0a4,<1",
    ],
    python_requires=">=3.6",
    setup_requires=["setuptools>=41.1.0"],
)
