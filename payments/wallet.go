package payments

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Wallet simulates a digital wallet
type Wallet struct {
	UserID  string
	Balance float64
}

func (w *Wallet) Process(ctx context.Context, amount float64) (string, error) {
	log.Printf("[Wallet] Processing payment of %.2f for user %s", amount, w.UserID)

	// Check for cancellation before processing
	select {
	case <-ctx.Done():
		log.Printf("[Wallet] Payment cancelled: %v", ctx.Err())
		return "", ctx.Err()
	default:
	}

	// simulate delay with context cancellation support
	delay := time.Duration(rand.Intn(500)+50) * time.Millisecond
	select {
	case <-time.After(delay):
	case <-ctx.Done():
		log.Printf("[Wallet] Payment cancelled during delay: %v", ctx.Err())
		return "", ctx.Err()
	}

	// Check for cancellation before service call
	select {
	case <-ctx.Done():
		log.Printf("[Wallet] Payment cancelled: %v", ctx.Err())
		return "", ctx.Err()
	default:
	}

	// 25% chance of transient error
	if rand.Intn(4) == 0 {
		err := errors.New("wallet service temporarily unavailable")
		log.Printf("[Wallet] Payment failed: %v", err)
		return "", err
	}

	if amount > w.Balance {
		err := errors.New("insufficient wallet balance")
		log.Printf("[Wallet] Payment failed: %v", err)
		return "", err
	}

	w.Balance -= amount
	msg := fmt.Sprintf("✅ [Wallet] Payment of %.2f succeeded (remaining balance: %.2f)", amount, w.Balance)
	log.Printf("[Wallet] Payment succeeded: remaining balance %.2f", w.Balance)
	return msg, nil
}
