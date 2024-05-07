// Create only DATABASE and TABLE, other operations will be handled by Golang

CREATE DATABASE stocksdb;

\c stocksdb;

CREATE TABLE stocks (
    stockid SERIAL PRIMARY KEY,
    name TEXT,
    price INT,
    company TEXT
);

SELECT * FROM stocks;