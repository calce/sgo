[![Build Status](https://travis-ci.org/calce/sgo.svg)](https://travis-ci.org/calce/sgo)
# sgo
Golang [Square Connect API](https://connect.squareup.com)

## Install
```
$ go get github.com/calce/sgo
```

### Example Usage
```go
package main

import (
	"github.com/calce/sgo"
	"github.com/calce/sgo/payments"
)

func main() {

	sgo.Token = "Secret"
	sgo.MerchantId = "me"
	
	payments, _ := payments.List(nil)
	
	for payments.HasNext() {
		payment, _ := payments.Next()
	}

}
```