# SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

from setuptools import setup, find_packages

with open("./README.md") as f:
    readme = f.read()

with open("../VERSION") as f:
    version = f.read()

setup(
    name="onos-api",
    version=version,
    description="ONOS Python API",
    long_description=readme,
    long_description_content_type="text/markdown",
    author="Open Networking Foundation (ONF)",
    packages=find_packages(exclude=("tests",)),
    namespace_packages=[],
    include_package_data=True,
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
    ],
    scripts=[],
    install_requires=[
        "protobuf",
        "grpclib",
        "betterproto",
        "gnmi-proto",
    ],
    python_requires=">=3.6",
    setup_requires=["setuptools>=41.1.0"],
)
