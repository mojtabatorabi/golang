package main

import (
    "fmt"
    "sync"
)

func main() {
    acc := BankAccount{
        Owner:      "Mojtaba",
        Balance:    1000,
        DailyLimit: 500,
    }

    var wg sync.WaitGroup

    // ۵ goroutine که واریز انجام می‌دهند
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(amount float64) {
            defer wg.Done()
            acc.Deposit(amount)
        }(100)
    }

    // ۵ goroutine که برداشت انجام می‌دهند
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(amount float64) {
            defer wg.Done()
            err := acc.Withdraw(amount)
            if err != nil {
                fmt.Println("Error:", err)
            }
        }(50)
    }

    wg.Wait() // منتظر می‌مانیم همه goroutine‌ها تمام شوند

    fmt.Printf("Final Balance: %.2f\n", acc.GetBalance())
    acc.PrintTransactions()
}
