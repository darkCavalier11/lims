#lims
lims short for Library Information Management System contains all the functionality to manage a 
real world library functionalities. Contains 4 tables used to manage different data

```sql
CREATE TABLE Book (
  "book_id" varchar(1200) PRIMARY KEY,
  "isbn" varchar(1200) NOT NULL,
  "title" varchar(1200) NOT NULL,
  "subtitle" text,
  "author" varchar(1200) NOT NULL,
  "published" varchar(1200) NOT NULL,
  "publisher" varchar(1200) NOT NULL,
  "pages" int,
  "description" text,
  "image_url" varchar(1200),
  "reviews" int,
  "rating" float
);

CREATE TABLE RegUser (
  "user_id" varchar(200) PRIMARY KEY,
  "first_name" varchar(200) NOT NULL,
  "last_name" varchar(200),
  "gender" varchar(200) NOT NULL,
  "email" varchar(200) UNIQUE NOT NULL,
  "password" varchar(200) NOT NULL,
  "is_admin" boolean NOT NULL
);

CREATE TABLE Review
(
    "review_id" varchar(200) PRIMARY KEY,
    "user_id" varchar(200) NOT NULL,
    "book_id" varchar(200) NOT NULL,
    "comment" text,
    "rating" int,
    "date" varchar(200),
    "edited" boolean
);

CREATE TABLE BookIssue
(
    "issue_id" varchar(200) PRIMARY KEY,
    "user_id" varchar(200) NOT NULL,
    "book_id" varchar(200) NOT NULL,
    "issue_date" varchar(200),
    "return_date" varchar(200),
    "returned" boolean
);

ALTER TABLE Review ADD FOREIGN KEY ("user_id") REFERENCES RegUser ("user_id");
ALTER TABLE Review ADD FOREIGN KEY ("book_id") REFERENCES Book ("book_id");

ALTER TABLE BookIssue ADD FOREIGN KEY ("user_id") REFERENCES RegUser ("user_id");
ALTER TABLE BookIssue ADD FOREIGN KEY ("book_id") REFERENCES Book ("book_id");


```

Few basic functionalities are already implemented and others can be implemented or can be built upon this.
All the implemented functionalities include
```
    GetBookById
    AddBook
    DeleteBook
    SearchBook
    GetReviewsOfBook
    IssueBook
    GetIssueById
    GetBooksIssuedByUser
    Connect
    AddEditDeleteReview
    AddAndDeleteUser
    SearchUser
```

## file structure and functions

`Connect()` method of `db/db.go` invoked inside `main.go` use to connect to the local or cloud db
and instantiate a global variable `Lib` which can used by external user
to interact with the db instance based on the functionality 
available. <br> Inside `db` these files contains code that
directly interact with the database which in turn provides the methods
outlined above. Any new method should be added here and the test as well.
