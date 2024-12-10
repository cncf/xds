from setuptools import setup, find_packages

setup(
    name='xds',  
    version='0.1.0', 
    description='xds Protocol Buffer Messages', 
    packages=find_packages(),
    install_requires=[
        'protobuf==5.29.1',  
        'protoc-gen-validate',
    ],
)