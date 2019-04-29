# 1build

[![Build Status](https://travis-ci.org/gopinath-langote/1build.svg?branch=master)](https://travis-ci.org/gopinath-langote/1build)

A simple way to unify build commands across all your projects running original building steps with all their specifics under the hood. It is agnostic to the underlying build tool, and environment preparatory steps are supported too.

## Why?

Imagine you are a microservices developer switching around multiple projects written in different languages and built with different building processes. Instead of reading README files all the time learning the specifics of how to run every one of them – you can just capture the configuration once and then use a single unified build command. 

With the support of preparatory and clean up steps – you can include various environment preparations and have them run as part of the build.

## Install

```bash
pip install -i https://test.pypi.org/simple/ --extra-index-url https://pypi.python.org/simple gopinathr1
```

or

```bash
pip3 install -i https://test.pypi.org/simple/ --extra-index-url https://pypi.python.org/simple gopinathr1
```

## Usage

### Configuration

- сreate project configuration file in the project folder
- file name: `1build.yaml`

Example of `1build.yaml` for JVM maven project:
```yaml
project: Sample JVM Project Name
commands:
  - build: mvn clean package
  - lint: mvn antrun:run@ktlint-format
```

Running 1build

- building the project
```bash
1build build
```

- reformat the code with kotlin-lint
```bash
1build lint
```
