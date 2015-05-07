package sgo

const (
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
	DEL = "DELETE"
)

// BankAccount.Type
const (
	BankAccountBusinessChecking		= "BUSINESS_CHECKING"
	BankAccountChecking						= "CHECKING"
	BankAccountInvestment					= "INVESTMENT"
	BankAccountLoan								= "LOAN"
	BankAccountSavings						= "SAVINGS"
	BankAccountOther							= "OTHER"
)

// Discount.Type
const (
	DiscountFixed									= "FIXED"
	DiscountVariablePercentage		= "VARIABLE_PERCENTAGE"
	DiscountVariableAmount				= "VARIABLE_AMOUNT"
)

// Employee.Status
const (
	EmployeeStatusActive					= "ACTIVE"
	EmployeeStatusInactive				= "INACTIVE"
)

// EmployeeRole.Permission
const (
	PermissionAccessSalesHistory				= "REGISTER_ACCESS_SALES_HISTORY"
	PermissionAppleRestrictedDiscounts	= "REGISTER_APPLY_RESTRICTED_DISCOUNTS"
	PermissionChangeSettings						= "REGISTER_CHANGE_SETTINGS"
	PermissionEditItem									= "REGISTER_EDIT_ITEM"
	PermissionIssueRefunds							= "REGISTER_ISSUE_REFUNDS"
	PermissionOpenCashDrawerOutsideSale	= "REGISTER_OPEN_CASH_DRAWER_OUTSIDE_SALE"
	PermissionViewSummaryReports				= "REGISTER_VIEW_SUMMARY_REPORTS"
)

// Fee.AdjustmentType
const (
	FeeAdjustmentTypeTax = "TAX"
)

// Fee.CalculationPhase
const (
	FeeCalculationPhaseSubtotal		= "FEE_SUBTOTAL_PHASE"
	FeeCalculationPhaseTotal			= "FEE_TOTAL_PHASE"
	FeeCalculationPhaseOther			= "OTHER"
)

// Fee.InclusionType
const (
	FeeInclusionTypeAdditive			= "ADDITIVE"
	FeeInclusionTypeInclusive			= "INCLUSIVE"
)

// Fee.Type
const (
	FeeTypeCAGST 							= "CA_GST"
	FeeTypeHST 								= "CA_HST"
	FeeTypeCAPEIPST						= "CA_PEI_PST"
	FeeTypeCAPST 							= "CA_PST"
	FeeTypeCAQST 							= "CA_QST"
	FeeTypeJPConsumptionTax 	= "JP_CONSUMPTION_TAX"
	FeeTypeUSSalesTax 				= "US_SALES_TAX"
	FeeTypeOther 							= "OTHER"
)

// InventoryAdjustmentType
const (
	InventoryAdjustmentTypeSale 						= "SALE	"
	InventoryAdjustmentTypeReceiveStock 		= "RECEIVE_STOCK"
	InventoryAdjustmentTypeManualAdjust 		= "MANUAL_ADJUST"	
)

// InventoryAlertType
const (
	InventoryAlertTypeLowQuantity 					= "LOW_QUANTITY"
	InventoryAlertTypeNone 									= "NONE"
)

// Item.Color
const (
	ItemColorGray 					 = "9da2a6" //	A gray color.
	ItemColorLightGreen			 = "4ab200" //	A lighter green color.
	ItemColorDarkGreen			 = "0b8000" //	A darker green color.
	ItemColorLightBlue			 = "13b1bf" //	A lighter blue color.
	ItemColorDarkBlue				 = "2952cc" //	A darker blue color.
	ItemColorPurple					 = "a82ee5" //	A purple color.
	ItemColorLightRed				 = "e5457a" //	A lighter red color.
	ItemColorDarkRed				 = "b21212" //	A dark red color.
	ItemColorGold						 = "e5BF00" //	A gold color.
	ItemColorBrown					 = "593c00" //	A brown color.
)

// Item.Visibility
const(
	ItemVisibilityPublic  		= "PUBLIC"
	ItemVisibilityPrivate 		= "PRIVATE"
)

// ItemVariation.PricingType
const (
	ItemVariationPricingTypeFixed 		 = "FIXED_PRICING"
	ItemVariationPricingTypeVariable	 = "VARIABLE_PRICING"
)

// ListOrder
const (
	ListOrderASC 		= "ASC"
	ListOrderDESC 	= "DESC"
)

// Money.CurrencyCode
const (
	CAD = "CAD"
	JPY = "JPY"
	USD = "USD"
)

// Merchant.BusinessType
const (
)

// Merchant.AccountType
const (
	MerchantAccountTypeBusiness = "BUSINESS"
	MerchantAccountTypeLocation = "LOCATION"
)

// Merchant.AccountCapability
const (
	MerchantAccountCapabilityCreditCardProcessing = "CREDIT_CARD_PROCESSING"
)

// ModifierList.SelectionType
const (
	ModifierListSelectionTypeSingle 		= "ModifierListSelectionType"
	ModifierListSelectionTypeMultiple 	= "ModifierListSelectionType"
)

// OAuth.Permission
const (
	OAuthPermissionMerchantProfileRead 			= "MERCHANT_PROFILE_READ"
	OAuthPermissionPaymentsRead 						= "PAYMENTS_READ"
	OAuthPermissionPaymentsWrite 						= "PAYMENTS_WRITE"
	OAuthPermissionSettlementsRead 					= "SETTLEMENTS_READ"
	OAuthPermissionBankAccountsRead 				= "BANK_ACCOUNTS_READ"
	OAuthPermissionItemsRead 								= "ITEMS_READ"
	OAuthPermissionItemsWrite 							= "ITEMS_WRITE"
	OAuthPermissionOrdersRead 							= "ORDERS_READ"
	OAuthPermissionOrdersWrite 							= "ORDERS_WRITE"
	OAuthPermissionEmployeesRead 						= "EMPLOYEES_READ"
	OAuthPermissionEmployeesWrite 					= "EMPLOYEES_WRITE"
	OAuthPermissionTimeCardsRead 						= "TIMECARDS_READ"
	OAuthPermissionTimeCardsWrite 					= "TIMECARS_WRITE"

)

// Order.Action
const (
	OrderActionComplete = "COMPLETE"
	OrderActionCancel = "CANCEL"
	OrderActionRefund = "REFUND"
)

// Order.State
const (
	OrderStatePending 			= "PENDING"
	OrderStateOpen 					= "OPEN"
	OrderStateCompleted 		= "COMPLETED"
	OrderStateCancelled 		= "CANCELED"
	OrderStateRefunded 			= "REFUNDED"
	OrderStateRejected 			= "REJECTED"
)

// OrderHistoryEntry.ActionType
const (
	OrderHistoryEntryActionTypeOrderPlaced 				= "ORDER_PLACED"
	OrderHistoryEntryActionTypeDeclined 					= "DECLINED"
	OrderHistoryEntryActionTypePaymentReceived 		= "PAYMENT_RECEIVED"
	OrderHistoryEntryActionTypeCanceled 					= "CANCELED"
	OrderHistoryEntryActionTypeCompleted 					= "COMPLETED"
	OrderHistoryEntryActionTypeRefunded 					= "REFUNDED"
	OrderHistoryEntryActionTypeExpired 						= "EXPIRED"
)

// PaymentItemization.Type
const (
	PaymentItemizationTypeItem 									= "ITEM"
	PaymentItemizationTypeCustomAmount 					= "CUSTOM_AMOUNT"
	PaymentItemizationTypeGiftCardActivation 		= "GIFT_CARD_ACTIVATION"
	PaymentItemizationTypeGiftCardReloaded 			= "GIFT_CARD_RELOAD"
	PaymentItemizationTypeGiftCardUnknown 			= "GIFT_CARD_UNKNOWN"
	PaymentItemizationTypeOther 								= "OTHER"
)

// Refund.Type
const (
	RefundTypeFull 		= "FULL"
	RefundTypePartial = "PARTIAL"
)

// PageCell.Type
const (
	PageCellTypeItem = "ITEM"
	PageCellTypeDiscount = "DISCOUNT"
	PageCellTypeCategory = "CATEGORY"
	PageCellTypePlaceHolder = "PLACEHOLDER"
)

// PageCell.PlaceholderType
const (
	PageCellPlaceholderTypeAllItems = "ALL_ITEMS"
	PageCellPlaceholderTypeDiscountsCategory = "DISCOUNTS_CATEGORY"
	PageCellPlaceholderTypeRewardsFinder = "REWARDS_FINDER"
)

// RequestMethod
const (
	RequestMethodDelete = "DELETE"
	RequestMethodGet = "GET"
	RequestMethodPost = "POST"
	RequestMethodPut = "PUT"
)

// SettlementEntry.Type
const (
	SettlementEntryTypeAdjustment 											= "ADJUSTMENT"
	SettlementEntryTypeBalanceCharge 										= "BALANCE_CHARGE"
	SettlementEntryTypeCharge 													= "CHARGE"
	SettlementEntryTypeFreeProcessing 									= "FREE_PROCESSING"
	SettlementEntryTypeHoldAdjustment 									= "HOLD_ADJUSTMENT"
	SettlementEntryTypeRedeptionCode 										= "REDEMPTION_CODE"
	SettlementEntryTypeRefund 													= "REFUND"
	SettlementEntryTypeReturnedPayout 									= "RETURNED_PAYOUT"
	SettlementEntryTypeSquareCapitalAdvance 						= "SQUARE_CAPITAL_ADVANCE"
	SettlementEntryTypeSquareCapitalPayment 						= "SQUARE_CAPITAL_PAYMENT"
	SettlementEntryTypeSquareCapitalReversedPayment 		= "SQUARE_CAPITAL_REVERSED_PAYMENT"
	SettlementEntryTypeOther 														= "OTHER"
	SettlementEntryTypeIncentedPayment 									= "INCENTED_PAYMENT"
	SettlementEntryTypeReturnedACHEntry 								= "RETURNED_ACH_ENTRY"
	SettlementEntryTypeReturnedSquare275 								= "RETURNED_SQUARE_275"
	SettlementEntryTypeSquare275 												= "SQUARE_275"
)

// Settlement.Status
const (
	SettlementStatusFailed 		= "FAILED"
	SettlementStatusSent 			= "SENT"
)

// Subscription.PaymentMethod
const (
	SubscriptionPaymentMethodCard 									= "CARD"
	SubscriptionPaymentMethodPayoutAdjustment 			= "PAYOUT_ADJUSTMENT"
	SubscriptionPaymentMethodOther 									= "OTHER"
)

// Subscription.Status
const (
	SubscriptionStatusActive 						= "ACTIVE"
	SubscriptionStatusInGrade 					= "IN_GRACE"
	SubscriptionStatusDelinquent 				= "DELINQUENT"
	SubscriptionStatusCanceled 					= "CANCELED"
	SubscriptionStatusTerminated 				= "TERMINATED"
	SubscriptionStatusOther 						= "OTHER"
)

// SubscriptionFee.Status
const (
	SubscriptionFeeStatusPending 				= "PENDING"
	SubscriptionFeeStatusPaymentFailed 	= "PAYMENT_FAILED"
	SubscriptionFeeStatusPaid 					= "PAID"
	SubscriptionFeeStatusCanceled 			= "CANCELED"
	SubscriptionFeeStatusOther 					= "OTHER"
)

// Tender.CardBrand
const (
	TenderCardBrandUnknown 							= "UNKNOWN"
	TenderCardBrandVisa 								= "VISA"
	TenderCardBrandMasterCard 					= "MASTER_CARD"
	TenderCardBrandAmericanExpress 			= "AMERICAN_EXPRESS"
	TenderCardBrandDiscover 						= "DISCOVER"
	TenderCardBrandDiscoverDiners 			= "DISCOVER_DINERS"
	TenderCardBrandJCB 									= "JCB"
)

// Tender.EntryMethod
const (
	TenderEntryMethodManual 						= "MANUAL"
	TenderEntryMethodScanned 						= "SCANNED"
	TenderEntryMethodSquareCash 				= "SQUARE_CASH"
	TenderEntryMethodSquareWallet 			= "SQUARE_WALLET"
	TenderEntryMethodSwiped 						= "SWIPED"
	TenderEntryMethodWebForm 						= "WEB_FORM"
	TenderEntryMethodOther 							= "OTHER"
)

// Tender.Type
const (
	TenderTypeCreditCard 					= "CREDIT_CARD"
	TenderTypeCash 								= "CASH"
	TenderTypeThirdPartyCard 			= "THIRD_PARTY_CARD"
	TenderTypeNoSale 							= "NO_SALE"
	TenderTypeSquareWallet 				= "SQUARE_WALLET"
	TenderTypeSquareGiftCard 			= "SQUARE_GIFT_CARD"
	TenderTypeUnknown 						= "UNKNOWN"
	TenderTypeOther 							= "OTHER"
)

// TimecardEvent.Type
const (
	TimecardEventTypeCreate 						= "API_CREATE"
	TimecardEventTypeApiEdit 						= "API_EDIT"
	TimecardEventTypeClockIn 						= "CLOCKIN"
	TimecardEventTypeClockOut 					= "CLOCKOUT"
	TimecardEventTypeSupervisorClose 		= "SUPERVISOR_CLOSE"
	TimecardEventTypeEdit 							= "EDIT"
)

// WebhookEventType
const (
	WebhookEventTypePaymentUpdated 			= "PAYMENT_UPDATED"
	WebhookEventTypeTimeCardUpdated 		= "TIMECARD_UPDATED"	
)