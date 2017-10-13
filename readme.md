## Traceable error for the go lang
Each error contains information where it is occurred and prints this information.  
Usage example  
```go
package mytest  

import (
    "errors"
    "github.com/steptosky/go-traceable-error/errt"
)  

func myFunc() error {
    err := errors.New("test")
    if err != nil {
        return errt.NewFrom(err)
    }
    return nil
}
```

Result example:
```
    Trace: Error_test.go:21
        Error: err3
    Trace: Error_test.go:20
        Error: err2
    Trace: Error_test.go:19
        Error: err1
        Error: err0  // original go error, does not have trace information
```
```
    Trace: Error_test.go:36
        Error: err3
    Trace: Error_test.go:35
        Error: err2
    Trace: Error_test.go:34
        Error: err1
    Trace: Error_test.go:33
        Error: err0
```
Source printing is disabled
```
    Error: err3
    Error: err2
    Error: err1
    Error: err0
```

## Installation
With go
```
go get github.com/steptosky/go-traceable-error/errt
```
With [go dep](https://github.com/golang/dep)
```
dep ensure -add github.com/steptosky/go-traceable-error/errt
```