import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()
setuptools.setup(
    name='gopitest',
    version='0.0.1',
    scripts=['1build'],
    author="Gopinath Langote",
    install_requires=['PyYAML>=3.12'],
    author_email="gopinathlangote11@gmail.com",
    description="1build utility",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/gopinath-langote/1build",
    packages=[],
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
)
