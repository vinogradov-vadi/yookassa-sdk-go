// Package yoorefund describes all the necessary entities for working with YooMoney Refunds.
package yoorefund

import yoocommon "github.com/vinogradov-vadi/yookassa-sdk-go/yookassa/common"

// The Deal within which the payment is being carried out.
type Deal struct {
	// Deal ID. Taken from the linked payment.
	ID string `json:"id,omitempty"`

	// Information about money distribution.
	RefundSettlements []yoocommon.Settlement `json:"refund_settlements,omitempty"`
}
