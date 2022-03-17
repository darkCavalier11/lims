DROP DATABASE IF EXISTS lims;

CREATE DATABASE lims;

\c lims;

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
  "image_url" varchar(1200)
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
    "rating" int
);

CREATE TABLE BookIssue
(
    "issue_id" varchar(200) PRIMARY KEY,
    "user_id" varchar(200) NOT NULL,
    "book_id" varchar(200) NOT NULL,
    "issue_date" date,
    "return_date" date
);

ALTER TABLE Review ADD FOREIGN KEY ("user_id") REFERENCES RegUser ("user_id");
ALTER TABLE Review ADD FOREIGN KEY ("book_id") REFERENCES Book ("book_id");

ALTER TABLE BookIssue ADD FOREIGN KEY ("user_id") REFERENCES RegUser ("user_id");
ALTER TABLE BookIssue ADD FOREIGN KEY ("book_id") REFERENCES Book ("book_id");
