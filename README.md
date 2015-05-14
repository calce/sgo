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

|              | List | Retrieve | Delete | Update | Upload | Tests |
|--------------|:----:|:--------:|:------:|:------:|:------:|:-----:|
| Payments     |  [X] |    [X]   |  [ ]   |   [ ]  |        |    ![](http://progressed.io/bar/20)   |


* Transaction Management
	* [X] ![](http://progressed.io/bar/20) Payments
	* [ ] ![](http://progressed.io/bar/0) Settlements
	* [ ] ![](http://progressed.io/bar/0) Refunds
	* [ ] ![](http://progressed.io/bar/0) Orders
	* [ ] ![](http://progressed.io/bar/0) Merchant
	* [ ] ![](http://progressed.io/bar/0) Bank Accounts

* Item Management
	* [ ] ![](http://progressed.io/bar/0) Items
	* [ ] ![](http://progressed.io/bar/0) Variations
	* [ ] ![](http://progressed.io/bar/0) Inventory
	* [ ] ![](http://progressed.io/bar/0) Modifier Lists
	* [ ] ![](http://progressed.io/bar/0) Modifier Options
	* [ ] ![](http://progressed.io/bar/0) Categories
	* [ ] ![](http://progressed.io/bar/0) Discounts
	* [ ] ![](http://progressed.io/bar/0) Fees
	* [ ] ![](http://progressed.io/bar/0) Pages
	* [ ] ![](http://progressed.io/bar/0) Cells

* Batching
	* [ ] ![](http://progressed.io/bar/0) Submit Batch

* API Webhooks
	* [ ] ![](http://progressed.io/bar/0) List
	* [ ] ![](http://progressed.io/bar/0) Update

* Subscription Management
	* [ ] ![](http://progressed.io/bar/0) Subscriptions
	* [ ] ![](http://progressed.io/bar/0) Plans

* Multi-Location Management
	* [ ] ![](http://progressed.io/bar/0) List Locations
	* [ ] ![](http://progressed.io/bar/0) Retrieve Business

* Employee Management
	* [ ] ![](http://progressed.io/bar/0) Employees
	* [ ] ![](http://progressed.io/bar/0) Roles
	* [ ] ![](http://progressed.io/bar/0) Timecards
	* [ ] ![](http://progressed.io/bar/0) Cash Drawer Shifts