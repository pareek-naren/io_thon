# IO Go SDK

## Introduction ##

IO Customer Registration wrapper provides access to IO Identity Management Platform API.

IO is an Identity Management Platform that simplifies user registration while securing data. IO Platform simplifies and secures your user registration process, increases conversion with Social Login that combines major social platforms, and offers a solution with Traditional Customer Registration. You can gather a wealth of user profile data from Social Login or Traditional Customer Registration. 

With our golang sdk you can easily implement Standard login/registration, Social login and can also implement Single sign on in your golang application.

Please visit [here](http://www.linktothedocumentation.io/) for more information.

## Installation

To install, run:
`go get github.com/pareek-narendra/io_thon/ciam_thon`

Import the package:

`import "github.com/pareek-narendra/io_thon/ciam_thon"`

Install all package dependencies by running `go get ./...` in the root folder of this SDK.  

## Usage

Take a peek:

Before making any API calls, the IO API client must be initialized with your IO API key and API secret.

Sample code:

```
cfg := lr.Config{
    ApiKey:    <your API key>,
    ApiSecret: <your API secret>,
}

lrclient, err := lr.NewIO(&cfg)

if err != nil {
    // handle error
}
```