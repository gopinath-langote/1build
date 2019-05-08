![Logo](docs/assets/1build-logo.png)

---

[![PyPi Version](https://img.shields.io/pypi/v/1build.svg)](https://pypi.org/project/1build/)
[![Build Status](https://travis-ci.org/gopinath-langote/1build.svg?branch=master)](https://travis-ci.org/gopinath-langote/1build)
[![Code Coverage](https://img.shields.io/codecov/c/gh/gopinath-langote/1build.svg)](https://codecov.io/gh/gopinath-langote/1build)
[![Requirements Status](https://requires.io/github/gopinath-langote/1build/requirements.svg?branch=master)](https://requires.io/github/gopinath-langote/1build/requirements/?branch=master)

A simple way to unify build commands across all your projects running original building steps with all their specifics under the hood. It is agnostic to the underlying build tool, and environment preparatory steps are supported too.

## Why?

Imagine you are a microservices developer switching around multiple projects written in different languages and built with different building processes. Instead of reading README files all the time learning the specifics of how to run every one of them – you can just capture the configuration once and then use a single unified build command. 

With the support of preparatory and clean up steps – you can include various environment preparations and have them run as part of the build.

## Install

```bash
pip install 1build
```

or

```bash
pip3 install 1build
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

Running 1build for above sample project:

- building the project
```bash
1build build
```

- reformat the code lint
```bash
1build lint
```

# `before` and `after` setup
Consider one of your project `X` need `java 11` and all other needs `java 8`.
You want to switch to `java 11` when you work on project `X` and switch back to `java 8` when done with command.
Or one other project needs docker to be up before running a command or you need to clean database after running a command.
That can be achieved by using `before` & `after` configuration. 

**NOTE: both `before` and `after` commands are optional**
### Examples:
1. Project which need `java 11` 
```yaml
project: Sample JVM Project Name
before: ./switch_to_java_11.sh
after: ./switch_to_java_8.sh
commands:
  - build: mvn clean package
  - lint: mvn antrun:run@ktlint-format
```

2. Project which need docker to be up
```yaml
project: Containerized Project
before: ./docker_run.sh
commands:
  - build: ./gradlew clean 
  - test: ./gradlew test
```

3. Project which need Switching directory
 ```yaml
project: Containerized Project
after: ./clean_database.sh
commands:
  - build: ./gradlew clean 
  - test: ./gradlew test
```