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
    Type   string  // "deposit" یا "withdraw"
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
// متد Withdraw با محدودیت روزانه
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
// تابع Interactive Menu
// -------------------------
func main() {
    acc := BankAccount{
        Owner:      "Mojtaba",
        Balance:    1000,
        DailyLimit: 500,
    }

    for {
        fmt.Println("\n--- Bank Account Menu ---")
        fmt.Println("1. Show Balance")
        fmt.Println("2. Deposit")
        fmt.Println("3. Withdraw")
        fmt.Println("4. Show Transactions")
        fmt.Println("5. Exit")
        fmt.Print("Choose an option: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            fmt.Printf("Current Balance: %.2f\n", acc.GetBalance())
        case 2:
            fmt.Print("Enter deposit amount: ")
            var amount float64
            fmt.Scan(&amount)
            acc.Deposit(amount)
        case 3:
            fmt.Print("Enter withdrawal amount: ")
            var amount float64
            fmt.Scan(&amount)
            err := acc.Withdraw(amount)
            if err != nil {
                fmt.Println("Error:", err)
            }
        case 4:
            acc.PrintTransactions()
        case 5:
            fmt.Println("Exiting... Goodbye!")
            return
        default:
            fmt.Println("Invalid option, try again.")
        }
    }
}
