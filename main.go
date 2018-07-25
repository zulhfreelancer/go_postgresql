package main

import (
  _ "github.com/lib/pq"
  "database/sql"
  "fmt"
  "log"
  "strconv"
)

var db *sql.DB




type Book struct {
  isbn  string
  title  string
  author string
  price  float32
}




func init() {
  var err error
  dbUrl := "postgres://postgres:postgres@localhost/go_postgresql?sslmode=disable"
  db, err = sql.Open("postgres", dbUrl)
  if err != nil {
    log.Fatal(err)
  }

  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
}




func main() {
  getAllBooks()
  getOneBook()
  // createBook()
}




func getAllBooks() {
  fmt.Println("******* getAllBooks() *******")

  rows, err := db.Query("SELECT * FROM books")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  bks := make([]*Book, 0)
  for rows.Next() {
    bk := new(Book)
    err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
    if err != nil {
      log.Fatal(err)
    }
    bks = append(bks, bk)
  }
  if err = rows.Err(); err != nil {
    log.Fatal(err)
  }

  for _, bk := range bks {
    fmt.Printf("%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
  }
}




func getOneBook() {
  fmt.Println("******* getOneBook() *******")

  isbn := "978-1505255607" // 978-1505255607
  row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

  bk := new(Book)
  err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
  if err == sql.ErrNoRows {
    fmt.Println("Book not found")
    return
  } else if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}




func createBook() {
  fmt.Println("******* createBook() *******")
  isbn := "new_isbn"
  title := "new_title"
  author := "new_author"

  if isbn == "" || title == "" || author == "" {
    fmt.Println("One of the parameter is empty")
    return
  }
  price, err := strconv.ParseFloat("9.99", 32)
  if err != nil {
    fmt.Println("Failed to parse string to float")
    return
  }

  result, err := db.Exec("INSERT INTO books VALUES($1, $2, $3, $4)", isbn, title, author, price)
  if err != nil {
    fmt.Println("Failed to execute SQL to insert a new book")
    return
  }

  rowsAffected, err := result.RowsAffected()
  if err != nil {
    fmt.Println("Failed to retrieve (x row affected) information")
    return
  }

  fmt.Printf("Book %s created successfully (%d row affected)\n", isbn, rowsAffected)

}
