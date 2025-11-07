package payments

import "context"

// PaymentProcessor defines a common behavior for payment methods.
type PaymentProcessor interface {
	// Process attempts to charge the given amount.
	// Returns a success message or an error.
	// The context can be used for cancellation and timeout.
	Process(ctx context.Context, amount float64) (string, error)
}
