#!/usr/bin/env python

# SPDX-FileCopyrightText: © 2020 Open Networking Foundation <support@opennetworking.org>
# SPDX-License-Identifier: Apache-2.0

from setuptools import find_packages, setup


PACKAGES = find_packages("python")


def version():
    with open("VERSION") as f:
        return f.read()


setup(
    name="onos-api",
    version=version(),
    description="ONOS API for Python",
    author="Open Networking Foundation and Partners",
    author_email="support@opennetworking.org",
    packages=[f"onos_api.{pkg}" for pkg in PACKAGES],
    package_dir={f"onos_api.{pkg}": f"python/{pkg}".replace(".", "/") for pkg in PACKAGES},
    install_requires=["betterproto>=2.0.0b3,<3.0"],
    classifiers=[
        "License :: OSI Approved :: Apache Software License",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
    ],
)


