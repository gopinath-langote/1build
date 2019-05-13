![Logo](docs/assets/1build-logo.png)


[![PyPi Version](https://img.shields.io/pypi/v/1build.svg)](https://pypi.org/project/1build/)
[![Build Status](https://travis-ci.org/gopinath-langote/1build.svg?branch=master)](https://travis-ci.org/gopinath-langote/1build)
[![Code Coverage](https://img.shields.io/codecov/c/gh/gopinath-langote/1build.svg)](https://codecov.io/gh/gopinath-langote/1build)
[![Requirements Status](https://requires.io/github/gopinath-langote/1build/requirements.svg?branch=master)](https://requires.io/github/gopinath-langote/1build/requirements/?branch=master)
[![Requirements Status](https://img.shields.io/pypi/dm/1build.svg)](https://pypi.org/project/1build)

1build is an automation tool that arms you with the convenience to configure project-local command line aliases – and then 
run the commands quickly and easily. It is particularly helpful when you deal with multiple projects and switch between 
them all the time. It is often the fact that different projects use different build tools and have different environment 
requirements – and then switching from project to another is becoming increasingly cumbersome. That is where 1build comes 
into play.

With 1build you can create simple and easy memorable command aliases for commonly used project commands such as build, 
test, run or anything else. These aliases will have a project-local scope which means that they will be accessible only 
within the project directory. This way you can unify all your projects to build with the same simple command disregarding 
of what build tool they use. It will remove the hassle of remembering all those commands improving the mental focus for 
the things that actually matter.

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

- create project configuration file in the project folder with name `1build.yaml`

- Example of `1build.yaml` for JVM maven project:
    ```yaml
    project: Sample JVM Project Name
    commands:
      - build: mvn clean package
      - lint: mvn antrun:run@ktlint-format
      - test: mvn clean test
    ```

### Running 1build for the above sample project:

- building the project
  ```console
  1build build
  ```

- fix the coding guidelinges lint
  ```console
  1build lint
  ```

### Using `before` and `after` commands
Consider that your project `X` requires `Java 11` and the other project requires `Java 8`. It is a headache to always 
remember to switch the java version. What you want is to switch to `Java 11` automatically when you build the project 
`X` and switch it back to `Java 8` when the build is complete. Another example – a project requires `Docker` to be up 
and running or you need to clean up the database after running a test harness.

This is where `before` & `after` commands are useful. These commands are both optional – 
you can use one of them, both or neither.

#### Examples:
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

3. Clean up database after some commans
     ```yaml
    project: Containerized Project
    after: ./clean_database.sh
    commands:
      - build: ./gradlew clean 
    ```
