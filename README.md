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
### Progress / Tests

Transaction Management

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Payments          |  ✔   |    ✔     |        |        |        |![](http://progressed.io/bar/20)|
| Settlements       |      |          |        |        |        |![](http://progressed.io/bar/0) |
| Refunds           |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Orders            |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|
| Merchant          |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|
| Bank Accounts     |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|



Item Management

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Items             |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Variations        |      |          |        |        |        |![](http://progressed.io/bar/0) |
| Inventory         |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Modifier Lists    |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|
| Merchant          |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|
| Modifier Options  |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|
| Categories        |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Discounts         |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Fees              |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Pages             |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Cells             |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|

Batching

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Submit Batch      |      |          |        |        |        |![](http://progressed.io/bar/20)|

API Webhooks

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| API Webhooks      |      |          |        |        |        |![](http://progressed.io/bar/20)|

Subscription Management

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Subscriptions     |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Plans             |      |          |        |        |        |![](http://progressed.io/bar/20)|

Multi-Location Management
|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Location          |      |          |        |        |        |![](http://progressed.io/bar/20)|

Employee Management

|                     | List | Retrieve | Delete | Update | Upload | Tests |
|---------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Employees           |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Roles               |      |          |        |        |        |![](http://progressed.io/bar/0) |
| Timecards           |      |          |        |        |        |![](http://progressed.io/bar/20)|
| Cash Drawer Shifts  |   ︎   |          |        |        |        |![](http://progressed.io/bar/20)|
