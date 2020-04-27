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
    orcid: 0000-0002-1558-0816
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

![1build Official Logo](docs/assets/1build-logo.png)

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

# Statement Of Need

With a rapid development of different languages and frameworks, setting up your project or use case specific environments,  for example,  Machine Learning, can be a tricky task. If you’ve never set up something like that before, you might spend hours fiddling with different commands trying to get the thing to work. [1build](https://github.com/gopinath-langote/1build) is the key! 

With [1build](https://github.com/gopinath-langote/1build), once you setup the environment for your research projects, you’ll be able to focus right down into the use case and never have to worry about installing packages ever again.

[1build](https://github.com/gopinath-langote/1build) also takes care of programming language setup, creating a virtual environment, and automating the installation of librarires, for example, Machine Learning and Deep Learning Libraries. 
  
In addition to this, [1build](https://github.com/gopinath-langote/1build) has been developed to a high degree of best practice in research software development [@jime2017], and is thoroughly documented: https://1build.gitbook.io/1build/. The documentation has been well formatted with the aim of easy learning and accessibilty. Furthermore, [1build](https://github.com/gopinath-langote/1build) is automatically tested using example, integration, unit tests with code coverage and [A+ on Go Report Card](https://goreportcard.com/report/github.com/gopinath-langote/1build). The current version of [1build](https://github.com/gopinath-langote/1build) has been archived on [Zenodo](https://zenodo.org).

With its active community of developers, timely announcement of the updates and releases on [Twitter](https://twitter.com/GopinathLangote), it is used by many undergraduate and graduate students around the world.

# Acknowledgements

I acknowledge JetBrains 1Password for supporting the project as the official sponsor and the efforts of the official contributors of 1build.

# References
