//go:generate ffjson $GOFILE
package sgo

type BankAccount struct {
	Id string															`json:"id"`
	MerchantId string											`json:"merchant_id"`
	BankName string												`json:"bank_name"`
	Name string														`json:"name"`
	BankAccountType string								`json:"type"`
	RoutingNumber string									`json:"routing_number"`
	AccountNumber string									`json:"account_number_suffix"`
	CurrencyCode string										`json:"currency_code"`
}

type BatchRequest struct {
	Method string													`json:"method"`	
	RelativePath string										`json:"relative_path"`
	AccessToken string										`json:"access_token"`
	RequestId string											`json:"request_id"`
}

type BatchResponse struct {
	StatusCode float64										`json:"status_code"`
	Headers interface{}										`json:"headers"`
	Body interface{}											`json:"body"`
	RequestId string											`json:"request_id"`
}

type Coordinates struct {
	Latitude float64											`json:"latitude"`
	Longitude float64											`json:"longitude"`
}

type Category struct {
	Id string															`json:"id"`
	Name string														`json:"name"`
}

type Device struct {
	Name string														`json:"name"`
	Id string															`json:"id"`
}

type Discount struct {
	Id string															`json:"id"`
	Name string														`json:"name"`
	Rate string														`json:"rate"`
	AmountMoney Money											`json:"amount_money"`
	DiscountType string										`json:"discount_type"`
	PinRequired bool											`json:"pin_required"`
	Color string													`json:"color"`
}

type Employee struct {
	Id string															`json:"id"`
	FirstName string											`json:"first_name"`
	LastName string												`json:"last_name"`
	RoleIds []string											`json:"role_ids"`
	Status string													`json:"status"`
	ExternalId string											`json:"external_id"`
	CreatedAt string											`json:"created_at"`
	UpdatedAt string											`json:"updated_at"`
}

type EmployeeRole struct {
	Id string															`json:"id"`
	Name string														`json:"name"`
	Permissions []string									`json:"permissions"`
	IsOwner bool													`json:"is_owner"`
	CreatedAt string											`json:"created_at"`
	UpdatedAt string											`json:"updated_at"`
}

type Fee struct {
	Id string															`json:"id"`
	Name string														`json:"name"`
	Rate string														`json:"rate"`
	CalculationPhase string								`json:"calculation_phase"`
	AdjustmentType string									`json:"adjustment_type"`
	AppliesToCustomAmounts bool						`json:"applies_to_custom_amounts"`
	Enabled bool													`json:"enabled"`
	InclusionType string									`json:"inclusion_type"`
	FeeType string												`json:"type"`
}

type GlobalAddress struct {
	AddressLine1 string										`json:"address_line_1"`
	AddressLine2 string										`json:"address_line_2"`
	AddressLine3 string										`json:"address_line_3"`
	AddressLine4 string										`json:"address_line_4"`
	AddressLine5 string										`json:"address_line_5"`
	Locality string												`json:"locality"`
	Sublocality string										`json:"sublocality"`
	Sublocality1 string										`json:"sublocality1"`
	Sublocality2 string										`json:"sublocality2"`
	Sublocality3 string										`json:"sublocality3"`
	Sublocality4 string										`json:"sublocality4"`
	Sublocality5 string										`json:"sublocality5"`
	AdministrativeDistrictLevel1 string		`json:"administrative_district_level_1"`
	AdministrativeDistrictLevel2 string		`json:"administrative_district_level_2"`
	AdministrativeDistrictLevel3 string		`json:"administrative_district_level_3"`
	PostalCode string											`json:"postal_code"`
	CountryCode string										`json:"country_code"`
	AddressCoordinates string							`json:"address_coordinates"`
}

type InventoryEntry struct {
	VariationId string 										`json:"variation_id"`
	QuantityOnHand string									`json:"quantity_on_hand"`
}

type Item struct {
	Id string															`json:"id"`
	Name string														`json:"name"`
	Description string										`json:"description"`
	Abbreviation string										`json:"abbreviation"`
	Color string													`json:"color"`
	Visibility string											`json:"visibility"`
	AvailableOnline string								`json:"available_online"`
	MasterImage ItemImage									`json:"master_image"`
	Category Category											`json:"category"`
	Variations []ItemVariation						`json:"variations"`
	ModifierLists []ModifierList					`json:"modifier_lists"`
	Fees []Fee														`json:"fees"`
	Taxable bool													`json:"taxable"`
}

type ItemImage struct {
	Id string															`json:"id"`
	Url string														`json:"Url"`
}

type ItemVariation struct {
	Id string															`json:"id"`
	Name string														`json:"name"`
	ItemId string													`json:"item_id"`
	Ordinal string												`json:"ordinal"`
	PricingType string										`json:"pricing_type"`
	PriceMoney Money											`json:"price_money"`
	SKU string														`json:"sku"`
	TrackInventory string									`json:"track_inventory"`
	InventoryAlertType string							`json:"inventory_alert_type"`
	InventoryAlertThreshold float64				`json:"inventory_alert_threshold"`
	UserData string												`json:"user_data"`
}

type Merchant struct {
	Id string																			`json:"id"`
	Name string																		`json:"name"`
	Email string																	`json:"email"`
	AccountType string														`json:"account_type"`
	AccountCapabilities []string									`json:"account_capabilities"`	
	CountryCode string														`json:"country_code"`
	LanguageCode string														`json:"language_code"`
	CurrencyCode string														`json:"currency_code"`
	BusinessName string														`json:"business_name"`
	BusinessAddress GlobalAddress									`json:"business_address"`
	BusinessPhone PhoneNumber											`json:"business_phone"`
	BusinessType string														`json:"business_type"`
	ShippingAddress GlobalAddress									`json:"shipping_address"`
	LocationDetails MerchantLocationDetails				`json:"location_details"`
	MarketUrl string															`json:"market_url"`
}

type MerchantLocationDetails struct {
	Nickname string																`json:"nickname"`
}

type ModifierList struct {
	Id string																			`json:"id"`
	Name string																		`json:"name"`
	SelectionType string													`json:"selection_type"`
	Modifier_options []ModifierOption							`json:"modifier_options"`
}


type ModifierOption struct {
	Id string																			`json:"id"`
	Name string																		`json:"name"`
	PriceMoney Money															`json:"price_money"`
	OnByDefault string														`json:"on_by_default"`
	Ordinal string																`json:"ordinal"`
	ModifierListId string													`json:"modifier_list_id"`
}

type Money struct {
	Amount float64																`json:"amount"`
	CurrencyCode string														`json:"currency_code"`
}

type Order struct {
	Id string																			`json:"id"`
	State string																	`json:"state"`
	BuyerEmail string															`json:"buyer_email"`
	RecipientName string													`json:"recipient_name"`
	RecipientPhoneNumber string										`json:"recipient_phone_number"`
	ShippingAddress GlobalAddress									`json:"shipping_address"`
	SubtotalMoney Money														`json:"subtotal_money"`
	TotalShippingMoney Money											`json:"total_shipping_money"`
	TotalTaxMoney Money														`json:"total_tax_money"`
	TotalPriceMoney Money													`json:"total_price_money"`
	TotalDiscountMoney Money											`json:"total_discount_money"`
	CreatedAt string															`json:"created_at"`
	UpdatedAt string															`json:"updated_at"`
	ExpiresAt string															`json:"expires_at"`
	PaymentId string															`json:"payment_id"`
	BuyerNote string															`json:"buyer_note"`
	CompletedNote string													`json:"completed_note"`
	RefundedNote string														`json:"refunded_note"`
	CanceledNote string														`json:"canceled_note"`
	Tender Tender																	`json:"tender"`
	OrderHistory []OrderHistoryEntry							`json:"order_history"`
	PromoCode string															`json:"promo_code"`
	BtcReceiveAddress string											`json:"btc_receive_address"`
	BtcPriceSatoshi float64												`json:"btc_price_satoshi"`
}

type OrderHistoryEntry struct {
	Action string																	`json:"action"`
	CreatedAt string																`json:"created_at"`
}

type Page struct {
	Id string																			`json:"id"`
	Name string																		`json:"name"`
	PageIndex float64															`json:"page_index"`
	Cells []PageCell															`json:"cells"`	
}

type PageCell struct {
	PageId string																	`json:"page_id"`
	Row float64																		`json:"row"`
	Column float64																`json:"column"`
	ObjectType string															`json:"object_type"`
	ObjectId string																`json:"object_id"`
	PlaceHolderType string												`json:"placeholder_type"`
}

type Payment struct {
	Id string																			`json:"id"`
	MerchantId string															`json:"merchant_id"`
	CreatedAt string															`json:"created_at"`
	CreatorId string															`json:"creator_id"`
	PaymentDevice Device													`json:"device"`
	PaymentUrl string															`json:"payment_url"`
	ReceiptUrl string															`json:"receipt_url"`
	InclusiveTaxMoney Money												`json:"inclusive_tax_money"`
	AdditiveTaxMoney Money												`json:"additive_tax_money"`
	TaxMoney Money																`json:"tax_money"`
	TipMoney Money																`json:"tip_money"`
	DiscountMoney Money														`json:"discount_money"`
	TotalCollectedMoney Money											`json:"total_collected_money"`
	ProcessingFeeMoney Money											`json:"processing_fee_money"`
	NetTotalMoney Money														`json:"net_total_money"`
	RefundedMoney Money														`json:"refunded_money"`
	InclusiveTax []PaymentTax											`json:"inclusive_tax"`
	AdditiveTax []PaymentTax											`json:"additive_tax"`
	Tender []Tender																`json:"tender"`
	Refunds []Refund															`json:"refunds"`
	Itemizations []PaymentItemization							`json:"itemizations"`
}

type PaymentDiscount struct {
	Name string																		`json:"name"`
	AppliedMoney Money														`json:"applied_money"`
	DiscountId string															`json:"discount_id"`	
}

type PaymentExtensions struct {
	Metadata map[string]string										`json:"metadata"`
}

type PaymentItemDetail struct {
	CategoryName string														`json:"category_name"`
	SKU string																		`json:"sku"`
	IconColor RGBAColor														`json:"icon_color"`
	ItemId string																	`json:"item_id"`
	ItemVariationId string												`json:"item_variation_id"`
}

type PaymentItemization struct {
	Name string																		`json:"name"`
	Quantity string																`json:"quantity"`
	ItemizationType string												`json:"itemization_type"`
	ItemDetail PaymentItemDetail									`json:"item_detail"`
  Notes string																	`json:"notes"`
  ItemVariationName string											`json:"item_variation_name"`
  TotalMoney Money															`json:"total_money"`
  SingleQuantityMoney Money											`json:"single_quantity_money"`
  GrossSalesMoney Money													`json:"gross_sales_money"`
  DiscountMoney Money														`json:"discount_money"`
  NetSalesMoney Money														`json:"net_sales_money"`
  Taxes []PaymentTax														`json:"taxes"`
  Discounts []PaymentDiscount										`json:"discounts"`
  Modifiers []PaymentModifier										`json:"modifiers"`
}

type PaymentModifier struct {
	Name string																		`json:"name"`
	AppliedMoney Money														`json:"applied_money"`
	ModifierOptionId string												`json:"modifier_option_id"`	
}

type PaymentTax struct {
	Name string																		`json:"name"`
	AppliedMoney Money														`json:"applied_money"`
	Rate string																		`json:"rate"`
	InclusionType string													`json:"inclusion_type"`
	FeeId string																	`json:"fee_id"`
}

type PhoneNumber struct {
	CallingCode string														`json:"calling_code"`
	Number string																	`json:"number"`	
}

type Refund struct {
	RefundType string															`json:"type"`
	Reason string																	`json:"reason"`
	RefundedMoney Money														`json:"refunded_money"`
	CreatedAt string															`json:"created_at"`
	ProcessedAt string														`json:"processed_at"`
	PaymentId string															`json:"payment_id"`
}

type Settlement struct {
	Id string																			`json:"id"`
	Status string																	`json:"status"`
	InitiatedAt string														`json:"initiated_at"`
	BankAccountId string													`json:"bank_account_id"`
	TotalMoney Money															`json:"total_money"`
	Entries []SettlementEntry											`json:"entries"`
}

type SettlementEntry struct {
	SettlementEntryType string										`json:"type"`
	PaymentId string															`json:"payment_id"`
	AmountMoney Money															`json:"amount_money"`
	FeeMoney Money																`json:"fee_money"`
}

type Subscription struct {
	Id string																			`json:"id"`
	MerchantId string															`json:"merchant_id"`
	PlanId string																	`json:"plan_id"`
	Status string																	`json:"status"`
	PaymentMethod string													`json:"payment_method"`
	FeeBaseMoney Money														`json:"fee_base_money"`
	ServiceStartDate string												`json:"service_start_date"`
	Fee []SubscriptionFee													`json:"fees"`
}

type SubscriptionFee struct {
	FeeDate string																`json:"fee_date"`
	FeeStatus string															`json:"fee_status"`
	FeeBaseMoney Money														`json:"fee_base_money"`
	FeeTaxMoney Money															`json:"fee_tax_money"`
	FeeTotalMoney Money														`json:"fee_total_money"`
}

type SubscriptionPlan struct {
	Id string																			`json:"id"`
	Name string																		`json:"name"`	
	CountryCode string														`json:"name"`	
	FeeBaseMoney Money														`json:"fee_base_money"`	
}

type Tender struct {
	Id string																			`json:"id"`
	TenderType string															`json:"type"`
	Name string																		`json:"name"`	
	ReceiptUrl string															`json:"receipt_url"`
	CardBrand string															`json:"card_brand"`
	PanSuffix string															`json:"pan_suffix"`
	EntryMethod string														`json:"entry_method"`
	PaymentNote string														`json:"payment_note"`
	TotalMoney Money															`json:"total_money"`
	TenderedMoney Money														`json:"tendered_money"`
	ChargeBackMoney Money													`json:"change_back_money"`
	RefundedMoney Money														`json:"refunded_money"`
}

type Timecard struct {
	Id string																			`json:"id"`	
	EmployeeId string															`json:"employee_id"`
	ClockInTime string														`json:"clockin_time"`
	ClockOutTime string														`json:"clockout_time"`
	ClockInLocationId string											`json:"clockin_location_id"`
	ClockOutLocationId string											`json:"clockout_location_id"`
	CreatedAt string															`json:"created_at"`
	UpdatedAt string															`json:"updated_at"`
}

type TimecardEvent struct {
	Id string																			`json:"id"`	
	EventType string															`json:"event_type"`
	ClockInTime string														`json:"clockin_time"`
	ClockOutTime string														`json:"clockout_time"`
	CreatedAt string															`json:"created_at"`	
}

// not in docs
type RGBAColor struct {
	A float64																			`json:"a"`
	R float64																			`json:"r"`
	G float64																			`json:"g"`
	B float64																			`json:"b"`
}