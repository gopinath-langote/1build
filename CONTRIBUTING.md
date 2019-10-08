# Contribution guidelines

First of all: Thank you! We really appreciate your efforts to make 1build better ❤️

Create or find an issue you would like to implement:
-   issues with the label `help wanted` are ready to be picked up and implemented
-   review the existing issues and ensure that you are not going to create a duplicate
-   create your issue and wait for comments from maintainers
-   once the issue is discussed and assigned to you – feel free to implement

## Developing 1build

1.  Prepare project (Install GoLang 1.9.x)

    ```sh
    git clone https://github.com/gopinath-langote/1build
    cd 1build
    
    go build 
    ```

2.  Make sure that all the existing tests are passed, extend tests if needed
    ```sh
    go test -v -cover github.com/gopinath-langote/1build/testing -run . 
    ```
    
    -   Alternatively `install` 1build from releases to get `1build` configuration for this project

3.  Project uses major library to build app - [cobra](https://github.com/spf13/cobra)

4.  Project uses - [go modules](https://github.com/golang/go/wiki/Modules) for dependency management.

5.  Update necessary documents if needed – Readme etc.

6.  Submit pull request

7.  Make sure all the checks are passing

8.  Wait for maintainers to review the code

9.  Thanks for you contribution :smile:
