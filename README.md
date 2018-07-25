Save and retrieve data to/from PostgreSQL database using Go.

Based on tutorial here:

[https://www.alexedwards.net/blog/practical-persistence-sql](https://www.alexedwards.net/blog/practical-persistence-sql)

### Terminal outputs

- First run, with `createBook()`.

```
$ go run main.go

******* getAllBooks() *******
978-1503261969, Emma, Jayne Austen, £9.44
978-1505255607, The Time Machine, H. G. Wells, £5.99
978-1503379640, The Prince, Niccolò Machiavelli, £6.99

******* getOneBook() *******
978-1505255607, The Time Machine, H. G. Wells, £5.99

******* createBook() *******
Book new_isbn created successfully (1 row affected)
```

- Second run, without `createBook()`.

```
$ go run main.go

******* getAllBooks() *******
978-1503261969, Emma, Jayne Austen, £9.44
978-1505255607, The Time Machine, H. G. Wells, £5.99
978-1503379640, The Prince, Niccolò Machiavelli, £6.99
new_isbn      , new_title, new_author, £9.99

******* getOneBook() *******
978-1505255607, The Time Machine, H. G. Wells, £5.99
```

### Database

Database name is `go_postgresql`.

### SQL to create and populate table

```
CREATE TABLE books (
  isbn    char(14) NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

ALTER TABLE books ADD PRIMARY KEY (isbn);
```
