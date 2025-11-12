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
