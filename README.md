GO (golang.org)

Motivation

    Efficient compilation (Java)
    Efficient execution (C++)
    Ease of programming (JS / Python)

Language Characteristics

    Strong, static type system
    C-inspired syntax
    Multi-paradigm (procedural or object oriented)
    Garbage collected
    Fully compiled
    Rapid compilation
    Single binary output
    Concurrent by default

Timeline

    2007 - start design
    2009 - publicly announced
    2012 - version 1.0
    2019 - version 1.12
    version 2.0 - ?!?


How Can Go Help? - Philosophy & Values

    Simplicity

        Increment and decrement are statements
        All loops in Go are simple loops

    Network aware and concurrent apps

        net and net/http packages - web server using only standar library
        goroutines - start thousands of concurrent threads with minimal resources
        channels - safely communicate between concurrent tasks

    Out-of-the box experience

        Standard library (string manipulation, data compression, file manipulation, networking api, testing, etc.)
        Go CLI (project initialization, build, code generation, retrieve dependencies, test, profiling, documentation, report language bugs)

    Cross-platform

        Change GOOS / GOARCH environment variables to compile to another platform

    Backward compatibility

        "It is intended that programs written to Go1 will continue compile and run correctly over the lifetime of the specification"

        Exceptions (security, unspecified behavior, spec errors, bugs)

    Go's Primary Use Case

        Really Shines (Web Services, Web Applications, DevOps)
        Task Automation
        GUI / Thick-client
        Machine Learning