package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
)

type Note struct {
    ID   int    `json:"id"`
    Text string `json:"text"`
}

const noteFile = "notes.json"

// خواندن یادداشت‌ها
func loadNotes() ([]Note, error) {
    if _, err := os.Stat(noteFile); os.IsNotExist(err) {
        return []Note{}, nil
    }
    data, err := ioutil.ReadFile(noteFile)
    if err != nil {
        return nil, err
    }
    if len(data) == 0 {
        return []Note{}, nil
    }
    var notes []Note
    err = json.Unmarshal(data, &notes)
    return notes, err
}

// ذخیره یادداشت‌ها
func saveNotes(notes []Note) error {
    data, err := json.MarshalIndent(notes, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(noteFile, data, 0644)
}

func main() {
    addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    listCmd := flag.NewFlagSet("list", flag.ExitOnError)
    delCmd := flag.NewFlagSet("del", flag.ExitOnError)
    addText := addCmd.String("text", "", "متن یادداشت جدید")

    if len(os.Args) < 2 {
        fmt.Println("Usage: notemgr [add|list|del]")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "add":
        addCmd.Parse(os.Args[2:])
        if *addText == "" {
            fmt.Println("لطفا با --text متن یادداشت را وارد کنید.")
            os.Exit(1)
        }
        notes, _ := loadNotes()
        newNote := Note{ID: len(notes) + 1, Text: *addText}
        notes = append(notes, newNote)
        saveNotes(notes)
        fmt.Println("✅ یادداشت اضافه شد:", newNote.Text)

    case "list":
        listCmd.Parse(os.Args[2:])
        notes, _ := loadNotes()
        if len(notes) == 0 {
            fmt.Println("⛔ هیچ یادداشتی وجود ندارد.")
            return
        }
        fmt.Println("📋 لیست یادداشت‌ها:")
        for _, n := range notes {
            fmt.Printf("%d. %s\n", n.ID, n.Text)
        }

    case "del":
        delCmd.Parse(os.Args[2:])
        if delCmd.NArg() == 0 {
            fmt.Println("لطفا شماره یادداشت را وارد کنید. مثل: notemgr del 2")
            return
        }
        id := delCmd.Arg(0)
        notes, _ := loadNotes()
        var newNotes []Note
        for _, n := range notes {
            if fmt.Sprintf("%d", n.ID) != id {
                newNotes = append(newNotes, n)
            }
        }
        saveNotes(newNotes)
        fmt.Println("🗑️ یادداشت حذف شد:", id)

    default:
        fmt.Println("Unknown command:", os.Args[1])
        fmt.Println("Available: add, list, del")
    }
}
