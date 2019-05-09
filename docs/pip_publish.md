# Releasing the distribution to PyPi

## Prerequisites
1. Install `wheel` & `setuptool`:
`python -m pip install --upgrade pip setuptools wheel`
2. Install `twine` – for publishing artifacts:
`pip install twine`

## Creating a distribution
1. Increment the version number in `setup.py`
2. Create the distribution package:
`python setup.py sdist bdist_wheel`

## Uploading the distribution to PyPi 
3. Upload distribution to PyPi:
`twine upload dist/* --verbose`

## References
1. Packaging Python Projects –  
https://packaging.python.org/tutorials/packaging-projects/

2. Python Packaging User Guide – 
https://packaging.python.org/
