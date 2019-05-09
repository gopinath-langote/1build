# Publishing artifact to PyPi

## Requirements:
1. Install `wheel` & `setuptool` by command
`python -m pip install --upgrade pip setuptools wheel`
2. Install twine to publish artifact
`pip install twine`

## Creating distribution
1. Change version in `setup.py` to newer version
2. Create source distribution
`python setup.py sdist bdist_wheel`

## Uploading distribution to PyPi 
3. Upload distribution to PyP.
`twine upload dist/* --verbose`

## References
1. Packaging Python Projects –  
https://packaging.python.org/tutorials/packaging-projects/

2. Python Packaging User Guide – 
https://packaging.python.org/
  