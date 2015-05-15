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
![alt text](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "Logo Title Text 1")
|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Payments          |  ✔   |    ✔     |        |        |        |![](https://github.com/adam-p/markdown-here/raw/master/src/common/images/icon48.png "80%")|
| Settlements       |      |          |        |        |        |![](http://bar.calce.co/0) |
| Refunds           |      |          |        |        |        |![](http://bar.calce.co/0)|
| Orders            |      |          |        |        |        |![](http://bar.calce.co/0)|
| Merchant          |      |          |        |        |        |![](http://bar.calce.co/0)|
| Bank Accounts     |      |          |        |        |        |![](http://bar.calce.co/0)|


Item Management

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Items             |      |          |        |        |        |![](http://bar.calce.co/0)|
| Variations        |      |          |        |        |        |![](http://bar.calce.co/0) |
| Inventory         |      |          |        |        |        |![](http://bar.calce.co/0)|
| Modifier Lists    |      |          |        |        |        |![](http://bar.calce.co/0)|
| Merchant          |      |          |        |        |        |![](http://bar.calce.co/0)|
| Modifier Options  |      |          |        |        |        |![](http://bar.calce.co/0)|
| Categories        |      |          |        |        |        |![](http://bar.calce.co/0)|
| Discounts         |      |          |        |        |        |![](http://bar.calce.co/0)|
| Fees              |      |          |        |        |        |![](http://bar.calce.co/0)|
| Pages             |      |          |        |        |        |![](http://bar.calce.co/0)|
| Cells             |      |          |        |        |        |![](http://bar.calce.co/0)|

Batching

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Submit Batch      |      |          |        |        |        |![](http://bar.calce.co/0)|

API Webhooks

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| API Webhooks      |      |          |        |        |        |![](http://bar.calce.co/0)|

Subscription Management

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Subscriptions     |      |          |        |        |        |![](http://bar.calce.co/0)|
| Plans             |      |          |        |        |        |![](http://bar.calce.co/0)|

Multi-Location Management

|                   | List | Retrieve | Delete | Update | Upload | Tests |
|-------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Locations         |      |          |        |        |        |![](http://bar.calce.co/0)|

Employee Management

|                     | List | Retrieve | Delete | Update | Upload | Tests |
|---------------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Employees           |      |          |        |        |        |![](http://bar.calce.co/0)|
| Roles               |      |          |        |        |        |![](http://bar.calce.co/0) |
| Timecards           |      |          |        |        |        |![](http://bar.calce.co/0)|
| Cash Drawer Shifts  |      |          |        |        |        |![](http://bar.calce.co/0)|
