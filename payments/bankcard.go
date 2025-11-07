package payments

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// BankCard simulates a payment via bank card.
type BankCard struct {
	Holder     string
	CardNumber string
	Limit      float64 // available limit / balance for simplicity
}

func (b *BankCard) Process(ctx context.Context, amount float64) (string, error) {
	log.Printf("[BankCard] Processing payment of %.2f for %s", amount, b.Holder)

	// Check for cancellation before processing
	select {
	case <-ctx.Done():
		log.Printf("[BankCard] Payment cancelled: %v", ctx.Err())
		return "", ctx.Err()
	default:
	}

	// simulate variable delay with context cancellation support
	delay := time.Duration(rand.Intn(400)+100) * time.Millisecond
	select {
	case <-time.After(delay):
	case <-ctx.Done():
		log.Printf("[BankCard] Payment cancelled during delay: %v", ctx.Err())
		return "", ctx.Err()
	}

	// Check for cancellation before network call
	select {
	case <-ctx.Done():
		log.Printf("[BankCard] Payment cancelled: %v", ctx.Err())
		return "", ctx.Err()
	default:
	}

	// simulate random network failure
	if rand.Intn(15) == 0 {
		err := errors.New("network error while processing bank card")
		log.Printf("[BankCard] Payment failed: %v", err)
		return "", err
	}

	if amount > b.Limit {
		err := errors.New("insufficient card limit")
		log.Printf("[BankCard] Payment failed: %v", err)
		return "", err
	}

	// deduct
	b.Limit -= amount
	msg := fmt.Sprintf("✅ [BankCard] Payment of %.2f succeeded (remaining limit: %.2f)", amount, b.Limit)
	log.Printf("[BankCard] Payment succeeded: remaining limit %.2f", b.Limit)
	return msg, nil
}
