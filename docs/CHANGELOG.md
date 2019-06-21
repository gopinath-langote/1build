# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## v1.2.0 [DRAFT]


## [v1.1.0](https://github.com/gopinath-langote/1build/releases/tag/v1.1.0) | 2019-06-21 | [Linked Issues](https://github.com/gopinath-langote/1build/milestone/5?closed=1) | [PyPi 1.0.0](https://pypi.org/project/1build/1.1.0/)

### Added
- Show 1build package version using `-v` or `--version`
- Create 1build configuration by `1build --init <project_name>
- Changelog draft
- Pull request template

### Changed
- Issue template to custom template
- `1build` logo to have tagline below it

### Deprecated

### Removed
- Running tests on macOS using Python version 3.6

### Fixed
- Missing test for `configuration file not found`
- Add `.pyc` in `.gitigore`
- Running tests on macOS using Python version 3.7

### Security


## [v1.0.0](https://github.com/gopinath-langote/1build/releases/tag/v1.0.0) | 2019-06-06 | [Linked Issues](https://github.com/gopinath-langote/1build/milestone/4?closed=1) | [PyPi 1.0.0](https://pypi.org/project/1build/1.0.0/)

### Added
- Introduce python argument parser
- `-h` or `--help` - show help message
- `-l` or `--list` - List all the commands available in current directory configuration `1build.yaml` file.
- [Contribution Guide](https://github.com/gopinath-langote/1build/blob/master/CONTRIBUTING.md) for developers.
- Introduce [changelog](https://github.com/gopinath-langote/1build/blob/master/docs/CHANGELOG.md) for the releases


### Changed
- Readme to include Contribution, Versioning, Changelog, Authors, Contributors.
- Readme badges to make center alignment


## [v0.0.5](https://github.com/gopinath-langote/1build/releases/tag/v0.0.5) | 2019-05-27 | [Linked Issues](https://github.com/gopinath-langote/1build/milestone/3?closed=1) | [PyPi 0.0.5](https://pypi.org/project/1build/0.0.5/)

### Added
- Run tests on multiple platforms on CI (Windows, Linux, MacOS)
- Check lint for code - use `pep8`

### Changed
- Move `1build` single script to new module
- Seperation of conerns in place - different python files for each responsibility

### Removed
- Support for `python 2` & lower versions than `py3.5`


## [v0.0.4](https://github.com/gopinath-langote/1build/releases/tag/v0.0.4) | 2019-05-13 | [Linked Issues](https://github.com/gopinath-langote/1build/milestone/2?closed=1) | [PyPi 0.0.4](https://pypi.org/project/1build/0.0.4/)

### Added
- Support for `before` and `after` setup for each command
- Tests in place
- `1build` logo
- Downloads, PyPi veresion badges in Readme
- Publish package documentation
- Add Code Coverage, Safety checks in CI

### Changed
- Test file structure to have different files for different feature tests
- More focused Readme description
- Better logging messages

### Security
- Add safety check to the Github PR & CI (travis)



## [v0.0.2](https://github.com/gopinath-langote/1build/releases/tag/v0.0.2) | 2019-05-01 | [Linked Issues](https://github.com/gopinath-langote/1build/milestone/1?closed=1) | [PyPi 0.0.2](https://pypi.org/project/1build/0.0.2/)

### Added
- Packaging python distriution
- Run 1build with simple project file with commands
- Project Skeleton
- First release of `1build`
- Support py2 & py3
