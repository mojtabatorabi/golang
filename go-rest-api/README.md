عالی 👌
بیا با هم یه پروژه‌ی ساده‌ی **Go REST API** بسازیم که از ساختار حرفه‌ای
`cmd`, `internal`, `pkg` استفاده می‌کنه.

---

## 📁 ساختار پروژه

```
go-rest-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── user/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── model.go
│   └── database/
│       └── connection.go
├── pkg/
│   └── logger/
│       └── logger.go
├── go.mod
└── go.sum
```

---

## 🔹 فایل `cmd/server/main.go`

```go
package main

import (
	"log"
	"net/http"

	"example.com/go-rest-api/internal/user"
	"example.com/go-rest-api/pkg/logger"
)

func main() {
	// Init logger
	log := logger.New()

	// Setup routes
	http.HandleFunc("/users", user.Handler)

	log.Info("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("Server failed: ", err)
	}
}
```

---

## 🔹 فایل `internal/user/model.go`

```go
package user

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
```

---

## 🔹 فایل `internal/user/service.go`

```go
package user

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func GetAll() []User {
	return users
}
```

---

## 🔹 فایل `internal/user/handler.go`

```go
package user

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	users := GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
```

---

## 🔹 فایل `pkg/logger/logger.go`

```go
package logger

import (
	"log"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(v ...interface{}) {
	log.Println("[INFO]", v)
}

func (l *Logger) Error(v ...interface{}) {
	log.Println("[ERROR]", v)
}
```

---

## 🚀 اجرای پروژه

```bash
go run ./cmd/server
```

خروجی در ترمینال:

```
[INFO] Starting server on :8080
```

و وقتی در مرورگر بری به:

```
http://localhost:8080/users
```

می‌بینی:

```json
[
  {"id":1, "name":"Alice"},
  {"id":2, "name":"Bob"}
]
```

---

آیا می‌خوای نسخه‌ی پیشرفته‌ترش رو هم برات بسازم (با Fiber یا Gin و اتصال به دیتابیس واقعی مثل PostgreSQL)؟


https://chatgpt.com/s/t_69147d7a07b881918284acc57537610b