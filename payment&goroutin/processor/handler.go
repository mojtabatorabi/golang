package processor

import (
	"context"
	"fmt"
	"log"
	"sync"

	"example.com/paymentapp/payments"
)

// HandlePayments runs payments concurrently and collects results via channel.
// The context can be used for cancellation and timeout.
func HandlePayments(ctx context.Context, processors []payments.PaymentProcessor, amounts []float64) {
	if len(processors) != len(amounts) {
		log.Printf("ERROR: processors and amounts length mismatch (processors: %d, amounts: %d)", len(processors), len(amounts))
		fmt.Println("processors and amounts length mismatch")
		return
	}

	log.Printf("Starting payment processing for %d payments", len(processors))

	var wg sync.WaitGroup
	results := make(chan string, len(processors))

	for i, p := range processors {
		// Check if context is already cancelled
		select {
		case <-ctx.Done():
			log.Printf("Payment processing cancelled, stopping at payment %d/%d", i, len(processors))
			close(results)
			return
		default:
		}

		wg.Add(1)
		go func(proc payments.PaymentProcessor, amt float64, idx int) {
			defer wg.Done()

			// Check context before processing
			select {
			case <-ctx.Done():
				log.Printf("Payment %d cancelled before processing", idx)
				results <- fmt.Sprintf("❌ [%T] Payment of %.2f cancelled: %v", proc, amt, ctx.Err())
				return
			default:
			}

			msg, err := proc.Process(ctx, amt)
			if err != nil {
				if err == context.Canceled || err == context.DeadlineExceeded {
					log.Printf("Payment %d cancelled or timed out: %v", idx, err)
				} else {
					log.Printf("Payment %d failed: %v", idx, err)
				}
				results <- fmt.Sprintf("❌ [%T] Payment of %.2f failed: %v", proc, amt, err)
			} else {
				log.Printf("Payment %d completed successfully", idx)
				results <- msg
			}
		}(p, amounts[i], i)
	}

	// Wait for all goroutines or context cancellation
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Printf("All payment goroutines completed")
	case <-ctx.Done():
		log.Printf("Payment processing cancelled, waiting for goroutines to finish")
		wg.Wait()
	}

	close(results)

	log.Printf("Collecting payment results")
	for r := range results {
		fmt.Println(r)
	}
	log.Printf("Payment processing completed")
}
