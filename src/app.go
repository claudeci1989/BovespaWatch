package main

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    // the actual postgresql driver. The underscore before the library means that we import pq without side effects. Basically, it means Go will only import the library for its initialization. For pq, the initialization registers pq as a driver for the SQL interface.
)

func main() {
    fmt.Println("The Bovespa Watch")

    db, err := sql.Open("postgres", "user=rodrigo dbname=bovespa-watch sslmode=disable")

    if err != nil {
    }
}
