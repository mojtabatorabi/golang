package main

import (
    "fmt"
    "errors"
    "sync"
    "time"
)

// -------------------------
// ساختار تراکنش
// -------------------------
type Transaction struct {
    Time   time.Time
    Type   string
    Amount float64
}

// -------------------------
// ساختار حساب بانکی
// -------------------------
type BankAccount struct {
    Owner          string
    Balance        float64
    Transactions   []Transaction
    DailyLimit     float64
    DailyWithdrawn float64
    Mutex          sync.Mutex
}

// -------------------------
// متد Deposit
// -------------------------
func (b *BankAccount) Deposit(amount float64) {
    if amount <= 0 {
        fmt.Println("Deposit amount must be positive")
        return
    }

    b.Mutex.Lock()
    defer b.Mutex.Unlock()

    b.Balance += amount
    b.Transactions = append(b.Transactions, Transaction{
        Time:   time.Now(),
        Type:   "deposit",
        Amount: amount,
    })

    fmt.Printf("Deposited %.2f, new balance: %.2f\n", amount, b.Balance)
}

// -------------------------
// متد Withdraw
// -------------------------
func (b *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }

    b.Mutex.Lock()
    defer b.Mutex.Unlock()

    if b.DailyWithdrawn+amount > b.DailyLimit {
        return errors.New("daily withdrawal limit exceeded")
    }

    if amount > b.Balance {
        return errors.New("insufficient funds")
    }

    b.Balance -= amount
    b.DailyWithdrawn += amount
    b.Transactions = append(b.Transactions, Transaction{
        Time:   time.Now(),
        Type:   "withdraw",
        Amount: amount,
    })

    fmt.Printf("Withdrew %.2f, new balance: %.2f\n", amount, b.Balance)
    return nil
}

// -------------------------
// متد GetBalance
// -------------------------
func (b *BankAccount) GetBalance() float64 {
    b.Mutex.Lock()
    defer b.Mutex.Unlock()
    return b.Balance
}

// -------------------------
// متد PrintTransactions
// -------------------------
func (b *BankAccount) PrintTransactions() {
    b.Mutex.Lock()
    defer b.Mutex.Unlock()

    fmt.Println("Transaction History:")
    for _, t := range b.Transactions {
        fmt.Printf("%s: %s %.2f\n", t.Time.Format("2006-01-02 15:04:05"), t.Type, t.Amount)
    }
}

// -------------------------
// تابع main concurrent
// -------------------------
func main() {
    acc := BankAccount{
        Owner:      "Mojtaba",
        Balance:    1000,
        DailyLimit: 1000,
    }

    var wg sync.WaitGroup

    // ۵ goroutine واریز
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(amount float64) {
            defer wg.Done()
            acc.Deposit(amount)
        }(100)
    }

    // ۵ goroutine برداشت
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(amount float64) {
            defer wg.Done()
            err := acc.Withdraw(amount)
            if err != nil {
                fmt.Println("Error:", err)
            }
        }(150)
    }

    wg.Wait() // منتظر می‌مانیم همه goroutine‌ها تمام شوند

    fmt.Printf("\nFinal Balance: %.2f\n", acc.GetBalance())
    acc.PrintTransactions()
}
