[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d5133cd1145e49729e90515a8cb1bcde)](https://app.codacy.com/app/kamiazya/go-dispatcher?utm_source=github.com&utm_medium=referral&utm_content=kamiazya/go-dispatcher&utm_campaign=badger)
# Dispatcher [![GoDoc](https://godoc.org/github.com/kamiazya/go-dispatcher?status.svg)](https://godoc.org/github.com/kamiazya/go-dispatcher) [![Build Status](https://travis-ci.org/kamiazya/go-dispatcher.svg?branch=master)](https://travis-ci.org/kamiazya/go-dispatcher) [![codecov.io](https://codecov.io/github/kamiazya/go-dispatcher/coverage.svg?branch=master)](https://codecov.io/github/kamiazya/go-dispatcher?branch=master) [![Maintainability](https://api.codeclimate.com/v1/badges/d53905c52749161e8345/maintainability)](https://codeclimate.com/github/kamiazya/go-dispatcher/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/d53905c52749161e8345/test_coverage)](https://codeclimate.com/github/kamiazya/go-dispatcher/test_coverage)

Package dispatcher implements a customizable Job Queue Dispatcher.

## Description

This package is a simple Job Queue of the Dispatcher-Worker model.

Implemented with API design that emphasizes Dispatcher's ease of creation and behavior customization.

## Features

- Simple Job-Queue
- Custom Dispatcher behavior

## Usage

Please install Package and include it in code.

```go
package main

import "github.com/kamiazya/go-dispatcher"

```

### New

You can get Dispatcher whth some Options like this.

```golang

// New Dispatcher
d, err := dispatcher.New(
    dispatcher.MaxWorker(2),
    dispatcher.MaxRetry(2),
)
if err != nil {
    // do something.
}

d.Start()

d.Dispatch(func() error {
    // do something.
    return nil
})

// wait for all tasks done
d.Wait()

```

### From Config

You can get Dispatcher whth Config like this.

```go

// default config
c := dispatcher.DafaultConfig()

// set value to config
c.MaxQueue = 10

// generate Dispatcher from config
d, _ := dispatcher.GenerateFromConfig(*c)

d.Start()

d.Dispatch(func() error {
    // do something.
    fmt.Println("this is the pettern of ")
    return nil
})

// Stop Dispatcher after all tasks done.
d.Stop(false)
```

## Lifecycle

![lifecycle](./lifecycle.png)

## Installation

```bash
go get github.com/kamiazya/go-dispatcher
```
