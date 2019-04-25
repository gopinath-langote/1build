import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name='1build',
    version='0.0.4',
    scripts=['1build'],
    packages=setuptools.find_packages(),
    install_requires=['ruamel.yaml>=0.15.94'],
    setup_requires=['ruamel.yaml>=0.15.94'],
    url='https://github.com/gopinath-langote/1build',
    license='MIT',
    author='Gopinath Langote',
    author_email='gopinathlangote11@gmail.com',
    description='1 build command for all your build tools for each project.',
    long_description=long_description,
    long_description_content_type="text/markdown",
    classifiers=[
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
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
        "Topic :: Terminals",
        "Topic :: Software Development :: Build Tools"
    ],
)
