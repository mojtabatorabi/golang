package payments

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Crypto simulates a cryptocurrency payment
type Crypto struct {
	Address string
	Balance float64
}

func (c *Crypto) Process(ctx context.Context, amount float64) (string, error) {
	log.Printf("[Crypto] Processing payment of %.2f from address %s", amount, c.Address)

	// Check for cancellation before processing
	select {
	case <-ctx.Done():
		log.Printf("[Crypto] Payment cancelled: %v", ctx.Err())
		return "", ctx.Err()
	default:
	}

	// simulate delay with context cancellation support
	delay := time.Duration(rand.Intn(700)+150) * time.Millisecond
	select {
	case <-time.After(delay):
	case <-ctx.Done():
		log.Printf("[Crypto] Payment cancelled during delay: %v", ctx.Err())
		return "", ctx.Err()
	}

	// Check for cancellation before network call
	select {
	case <-ctx.Done():
		log.Printf("[Crypto] Payment cancelled: %v", ctx.Err())
		return "", ctx.Err()
	default:
	}

	// higher chance of transient failure
	if rand.Intn(6) == 0 {
		err := errors.New("crypto network congestion")
		log.Printf("[Crypto] Payment failed: %v", err)
		return "", err
	}

	if amount > c.Balance {
		err := errors.New("insufficient crypto balance")
		log.Printf("[Crypto] Payment failed: %v", err)
		return "", err
	}

	// simulate gas fee: small random deduction
	fee := float64(rand.Intn(3)) * 0.5
	total := amount + fee
	if total > c.Balance {
		err := errors.New("insufficient crypto balance after fees")
		log.Printf("[Crypto] Payment failed: %v", err)
		return "", err
	}

	c.Balance -= total
	msg := fmt.Sprintf("✅ [Crypto] Payment of %.2f succeeded (fee %.2f, remaining %.2f)", amount, fee, c.Balance)
	log.Printf("[Crypto] Payment succeeded: fee %.2f, remaining balance %.2f", fee, c.Balance)
	return msg, nil
}
