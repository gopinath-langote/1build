import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()
setuptools.setup(
    name='1build',
    version='0.0.3',
    scripts=['1build'],
    license="MIT License",
    author="Gopinath Langote",
    install_requires=[
        'ruamel.yaml>=0.15.94'
    ],
    author_email="gopinathlangote11@gmail.com",
    description="Unified build command for all project using underlying different build tools.",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/gopinath-langote/1build",
    packages=[],
    classifiers=[
        "Programming Language :: Python :: 2.7",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.0",
        "Programming Language :: Python :: 3.1",
        "Programming Language :: Python :: 3.2",
        "Programming Language :: Python :: 3.3",
        "Programming Language :: Python :: 3.4",
        "Programming Language :: Python :: 3.5",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "License :: OSI Approved :: MIT License",
        "Topic :: Software Development :: Build Tools",
        "Intended Audience :: Developers",
        "Operating System :: OS Independent",
    ],
)
