# SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

from setuptools import setup, find_packages


PACKAGES = find_packages("python/onos")


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
    packages=[f"onos_api.{pkg}" for pkg in PACKAGES],
    package_dir={
        f"onos_api.{pkg}": f"python/onos/{pkg}".replace(".", "/") for pkg in PACKAGES
    },
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
    install_requires=[
        "betterproto>=2.0.0b3,<3",
        "gnmi-proto",
    ],
    python_requires=">=3.6",
    setup_requires=["setuptools>=41.1.0"],
)
