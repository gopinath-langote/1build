![Logo](docs/assets/1build-logo.png)


[![PyPi Version](https://img.shields.io/pypi/v/1build.svg)](https://pypi.org/project/1build/)
[![Build Status](https://travis-ci.org/gopinath-langote/1build.svg?branch=master)](https://travis-ci.org/gopinath-langote/1build)
[![Code Coverage](https://img.shields.io/codecov/c/gh/gopinath-langote/1build.svg)](https://codecov.io/gh/gopinath-langote/1build)
[![Requirements Status](https://requires.io/github/gopinath-langote/1build/requirements.svg?branch=master)](https://requires.io/github/gopinath-langote/1build/requirements/?branch=master)
[![Requirements Status](https://img.shields.io/pypi/dm/1build.svg)](https://pypi.org/project/1build)

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

Running 1build for the above sample project:

- building the project
```bash
1build build
```

- reformat the code lint
```bash
1build lint
```

# Using `before` and `after` commands
Consider that your project `X` requires `Java 11` and the other project requires `Java 8`. It is a headache to always remember to switch the java version. What you want is to switch to `Java 11` automatically when you build the project `X` and switch it back to `Java 8` when the build is complete. Another example – a project requires `Docker` to be up and running or you need to clean up the database after running a test harness.

This is where `before` & `after` commands are useful. These commands are both optional – you can use one of them, both or neither.

### Examples:
1. Switching to `Java 11` and then back to `Java 8`
```yaml
project: Sample JVM Project Name
before: ./switch_to_java_11.sh
after: ./switch_to_java_8.sh
commands:
  - build: mvn clean package
```

2. Ensure that `Docker` is up and running
```yaml
project: Containerized Project
before: ./docker_run.sh
commands:
  - build: ./gradlew clean 
```

3. Clean up database on exit
 ```yaml
project: Containerized Project
after: ./clean_database.sh
commands:
  - build: ./gradlew clean 
```
