Payment Concurrency Exercise
============================

This is a sample Go project demonstrating:
- interfaces and structs
- goroutines and channels
- error handling
- package separation and project layout

Structure:
- payments/ : contains PaymentProcessor interface and implementations
- processor/ : contains the concurrent handler for payments
- main.go : example usage and entry point

To run:
```bash
cd payment_exercise
go mod tidy
go run ./...
```
