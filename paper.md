---
title: '1build: Frictionless way of managing project-specific commands'
tags:
  - bash
  - developer-tools
  - productivity
  - command-line
  - go
  - awesome-go
authors:
  - name: Gopinath Langote
    orcid: 0000-0000-0000-0000
    affiliation: 1
  - name: Deepak Ahire
    orcid: 0000-0002-9174-0797
    affiliation: 1
affiliations:
  - name: Walchand College of Engineering, Sangli
    index: 1
date: April 2020
bibliography: paper.bib
---

# Summary

<br>

<h1 align="center">
  <br>
  <a href="https://github.com/gopinath-langote/1build">
    <img src="https://github.com/gopinath-langote/1build/blob/master/docs/assets/1build-logo.png?raw=true" alt="1build" width="500"></a>
  <br>
</h1>

<br>

1build is an automation tool that arms you with the convenience to configure project-local command line aliases – and then
run the commands quickly and easily. It is particularly helpful when you deal with multiple projects and switch between
them all the time. It is often the fact that different projects use different build tools and have different environment
requirements – and then switching from one project to another is becoming increasingly cumbersome. That is where 1build comes
into play.

With 1build you can create simple and easily memorable command aliases for commonly used project commands such as build,
test, run or anything else. These aliases will have a project-local scope which means that they will be accessible only
within the project directory. This way you can unify all your projects to build with the same simple command disregarding
of what build tool they use. It will remove the hassle of remembering all those commands improving the mental focus for
the things that actually matter.

# Install

### Homebrew

```console
brew install gopinath-langote/one-build/one-build
```

### Manual

1.  Download and install binary from [the latest release](https://github.com/gopinath-langote/1build/releases/latest)
2.  Recommended: add `1build` executable to your `$PATH`.

# Usage

### Configuration
-   Create `1build.yaml` configuration file by
    ```console
    1build init --name <your_project_name>
    ```

-   Edit file according to project command list, Example of `1build.yaml` for node project:
    ```yaml
    project: Sample Web App
    commands:
      - build: npm run build
      - test: npm run test
    ```

### Running 1build for the above sample project

-   building the project
```console
  1build build
  ```

-   fix the coding guidelines lint and run tests (executing more than one commands at once)
```console
  1build lint test
  ```

### Set new or update existing configuration

-   Set new command configuration for `lint` to `"eslint server.js"`
```console
   1build set lint "eslint server.js"
   ```
### Remove/Unset existing configuration

-   Unset command configuration for `lint`
```console
   1build unset lint
   ```

-   To unset multiple commands at once
```console
  1build unset lint test build
  ```

### Using `before` and `after` commands
Consider that your project requires some environment variables to set before running any
commands and you want to clean up those after running commands. It is a headache to always
remember to set those environment variables. What you want is to set env variables automatically
when you run the command in the project and remove those when the command is complete.
Another example – a project requires `Docker` to be up
and running or you need to clean up the database after running a test harness.

This is where `before` & `after` commands are useful. These commands are both optional – 
you can use one of them, both or neither.

#### Examples
1.  Setting env variables and cleaning those up
    ```console
    1build set before 'export VARNAME="my value"'
    1build set after "unset VARNAME"
    ```
  
    Configuration with before and after setup
    
    ```yaml
    project: Sample Web App
    before: export VARNAME="my value"
    after: unset VARNAME
    commands:
       - build: npm run build
    ```

2.  Ensure that `Docker` is up and running
    ```console
    1build set before "./docker_run.sh"
    ```

3.  Clean up database after some commands
    ```console
    1build set after "./clean_database.sh"
    ```

4.  Remove `before` and `after` commands
    ```console
    1build unset before after
    ```

### Using `--quiet` or `-q` flag
Sometimes you choose to not see all logs/output of your command and just see success or failure as the outcome.
So using `--quiet` or `-q` flag to 1build command execution will result in just executing the command
but not showing the entire output of it, only shows SUCCESS/FAILURE as result of command execution.
```console
  1build lint test --quiet
  ```
  OR
```console
  1build lint test -q
  ```

See `1build --help` for command usages.

# Acknowledgements

I acknowledge JetBrains 1Password for supporting the project as the official sponsor and the efforts of the official contributors of 1build.

# References
