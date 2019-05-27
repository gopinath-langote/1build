import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()
setuptools.setup(
    name='1build',
    version='0.0.4',
    scripts=['1build'],
    license="MIT License",
    author="Gopinath Langote",
    install_requires=[
        'ruamel.yaml>=0.15.94'
    ],
    author_email="gopinathlangote11@gmail.com",
    description="Frictionless way of managing project-specific commands.",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/gopinath-langote/1build",
    packages=['onebuild'],
    classifiers=[
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "License :: OSI Approved :: MIT License",
        "Topic :: Software Development :: Build Tools",
        "Intended Audience :: Developers",
        "Operating System :: OS Independent",
    ],
)
