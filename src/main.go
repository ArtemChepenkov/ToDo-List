package main

import (
    "bufio"
    "database/sql"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "to_do_list/storage"
)

type DBRecord storage.DBRecord

func main() {
    db, err := storage.InitStorage("storage/db/storage.db") // make config read
    if err != nil {
        log.Fatalf("failed to init db: %w", err)
    }
    scanner := bufio.NewScanner(os.Stdin)
    userInteraction(scanner, db)
}

func userInteraction(scanner *bufio.Scanner, db *sql.DB) {
    fmt.Println("You can use commands:")
    fmt.Println("1. list (shows your tasks)")
    fmt.Println("2. add")
    fmt.Println("3. complete <id>")
    fmt.Println("4. uncomplete <id>")
    fmt.Println("5. remove <id>")
    fmt.Println("6. change <id>")

    for {
        fmt.Print("> ")
        scanner.Scan()
        option := strings.Fields(strings.ToLower(scanner.Text()))
        switch option[0] {
        case "list":
            records := storage.ListRecords(db)
            maxLen := 0
            for _, val := range records {
                if len(val.Task) > maxLen {
                    maxLen = len(val.Task)
                }
            }
            maxLen++
            for _, val := range records {
                var completed string
                if val.Completed == false {
                    completed = "❌"
                } else {
                    completed = "✅"
                }
                spaces := maxLen - len(val.Task)
                fmt.Printf("%d) ", val.Id)
                fmt.Print(val.Task)
                for i := 0; i < spaces; i++ {
                    fmt.Print(" ")
                }
                fmt.Println(completed)
            }

        case "add":
            fmt.Print("Type your task> ")
            scanner.Scan()
            task := scanner.Text()
            storage.AddRecord(task, db)

        case "complete":
            if len(option) <= 1 {
                fmt.Println("Correct syntax: remove <id>")
                continue
            }
            id, _ := strconv.Atoi(option[1])
            res := storage.SetComplete(id, true, db)
            if !res {
                fmt.Println("Wrong id")
            }

        case "uncomplete":
            if len(option) <= 1 {
                fmt.Println("Correct syntax: remove <id>")
                continue
            }
            id, _ := strconv.Atoi(option[1])
            res := storage.SetComplete(id, false, db)
            if !res {
                fmt.Println("Wrong id")
            }

        case "remove":
            if len(option) <= 1 {
                fmt.Println("Correct syntax: remove <id>")
                continue
            }
            id, _ := strconv.Atoi(option[1])
            res := storage.RemoveRecord(id, db)
            if !res {
                fmt.Println("Wrong id")
            }

        case "change":
            fmt.Print("Type your new task> ")
            id, _ := strconv.Atoi(option[1])
            scanner.Scan()
            task := scanner.Text()
            res := storage.SetTask(id, task, db)
            if !res {
                fmt.Println("Wrong id")
            }
        }
    }
}
