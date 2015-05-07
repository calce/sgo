package sgo

type PaymentListParams struct {
	BeginTime string											`param:"begin_time"`
	EndTime string												`param:"end_time"`
	Order string													`param:"order"`
	Limit int															`param:"limit"`
}

type SettlementListParams struct {
	BeginTime string											`param:"begin_time"`
	EndTime string												`param:"end_time"`
	Order string													`param:"order"`
	Limit int															`param:"limit"`
	Status string													`param:"status"`	
}

type RefundCreateParams struct {
	PaymentId string											`param:"payment_id"`
	RefundType string											`param:"type"`
	Reason string													`param:"reason"`
	RefundedMoney Money										`param:"refunded_money"`
	RequestIdempotenceKey string					`param:"request_idempotence_key"`
}

type RefundListParams struct {
	BeginTime string											`param:"begin_time"`
	EndTime string												`param:"end_time"`
	Order string													`param:"order"`
	Limit int															`param:"limit"`
}

type OrderListParams struct {
	Limit int															`param:"limit"`
	Order string													`param:"order"`
}

type OrderUpdateParams struct {
	Action string													`param:"action"`
	ShippedTrackingNumber string					`param:"shipped_tracking_number"`
	CompletedNote string									`param:"completed_note"`
	RefundedNote string										`param:"refunded_note"`
	CanceledNote string										`param:"canceled_note"`
}