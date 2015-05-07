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
### Progress
* Transaction Management
	* [X] Payments
	* [ ] Settlements
	* [ ] Refunds
	* [ ] Orders
	* [ ] Merchant
	* [ ] Bank Accounts

* Item Management
	* [ ] Items
	* [ ] Variations
	* [ ] Inventory
	* [ ] Modifier Lists
	* [ ] Modifier Options
	* [ ] Categories
	* [ ] Discounts
	* [ ] Fees
	* [ ] Pages
	* [ ] Cells

* Batching
	* [ ] Submit Batch

* API Webhooks
	* [ ] List
	* [ ] Update

* Subscription Management
	* [ ] Subscriptions
	* [ ] Plans

* Multi-Location Management
	* [ ] List Locations
	* [ ] Retrieve Business

* Employee Management
	* [ ] Employees
	* [ ] Roles
	* [ ] Timecards
	* [ ] Cash Drawer Shifts