package storage

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

type DBRecord struct {
    Id        int
    Task      string
    Completed bool
}

func InitStorage(storagePath string) (*sql.DB, error) {
    const OP = "storage.db_helper.InitStorage"
    db, err := sql.Open("sqlite3", storagePath)

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("%s: %w", OP, err)
    }

    if err = createInitTable(db); err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
    return db, nil
}

func createInitTable(db *sql.DB) error {
    const OP = "storage.db_helper.createInitTable"
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            Id INTEGER PRIMARY KEY,
            Task TEXT,
            completed BOOl
        )
    `)

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
    return nil
}

func ListRecords(db *sql.DB) []DBRecord {
    const OP = "storage.db_helper.ListRecords"
    records, err := db.Query("SELECT * FROM tasks")
    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
    defer records.Close()

    var result []DBRecord

    for records.Next() {
        var id int
        var task string
        var completed bool
        err = records.Scan(&id, &task, &completed)
        if err != nil {
            log.Fatalf("%s: %w", OP, err)
        }
        result = append(result, DBRecord{
            Id:        id,
            Task:      task,
            Completed: completed,
        })
    }
    return result
}

func AddRecord(task string, db *sql.DB) {
    const OP = "storage.db_helper.AddRecord"
    stmt, err := db.Prepare("INSERT INTO tasks(Task, completed) VALUES(?, ?)")

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }

    defer stmt.Close()

    _, err = stmt.Exec(task, false)

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
}

func RemoveRecord(id int, db *sql.DB) bool {
    const OP = "storage.db_helper.RemoveRecord"

    if !checkExists(id, OP, db) {
        return false
    }

    stmt, err := db.Prepare("DELETE FROM tasks WHERE Id = ?")

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(id)

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
    return true
}

func SetComplete(id int, completed bool, db *sql.DB) bool {
    const OP = "storage.db_helper.SetComplete"

    if !checkExists(id, OP, db) {
        return false
    }

    stmt, err := db.Prepare("UPDATE tasks SET completed = ? WHERE id = ?")

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }

    _, err = stmt.Exec(completed, id)

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
    return true
}

func SetTask(id int, task string, db *sql.DB) bool {
    const OP = "storage.db_helper.SetTask"

    if !checkExists(id, OP, db) {
        return false
    }

    stmt, err := db.Prepare("UPDATE tasks SET task = ? WHERE id = ?")

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }

    _, err = stmt.Exec(task, id)

    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }
    return true
}

func checkExists(id int, OP string, db *sql.DB) bool {
    records, err := db.Query("SELECT * FROM tasks WHERE Id = ?", id)
    if err != nil {
        log.Fatalf("%s: %w", OP, err)
    }

    defer records.Close()

    if !records.Next() {
        return false
    }

    return true
}
