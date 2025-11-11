package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"example.com/paymentapp/payments"
	"example.com/paymentapp/processor"
)

func main() {
	// Setup logging
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Payment processing application started")

	rand.Seed(time.Now().UnixNano())

	// Create context with timeout (can be cancelled if needed)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Optional: Uncomment to test cancellation
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	log.Println("Cancelling payment processing...")
	// 	cancel()
	// }()

	// create processors
	bank := &payments.BankCard{Holder: "Ali", CardNumber: "1111-2222-3333-4444", Limit: 200.0}
	wallet := &payments.Wallet{UserID: "user123", Balance: 80.0}
	crypto := &payments.Crypto{Address: "0xABCDEF", Balance: 50.0}

	processors := []payments.PaymentProcessor{bank, wallet, crypto, bank, wallet}
	amounts := []float64{45.0, 70.0, 30.0, 120.0, 20.0}

	fmt.Println("Processing payments...")
	processor.HandlePayments(ctx, processors, amounts)

	fmt.Println("\nAll payments processed!")
	log.Println("Payment processing application finished")
}
