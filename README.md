# Building Society Savings Account Interest Calculator

## Introduction

This package calculates a Building Society's savings account applicable interest.
The interest calculated is based on:

    1. The balance of the account
    2. The band the balance belongs to

The table below represents the different bands and their applicable interest rates:

    | Range               | Interest Rate |
    | ------------------- | ------------- |
    | < $1,000            | 1%            |
    | $1,000 - < $5,000   | 1.5%          |
    | $5,000 - < $10,000  | 2%            |
    | $10,000 - < $50,000 | 2.5%          |
    | $50,000 +           | 3%            |

E.g. A balance of $1,001 would pay interest at 1.5% resulting in an interest rate of $15.02.

## Code Structure

The package is structured using a thin DDD-inspired architecture as demostrated below:

    |- pkg
        |- account
            |- domain.go # This represents our domain layer (Datastructures that drive the package e.g InterestRatingTable)
            |- interest_rate.go # This represents our usecase (business logice) layer i.e where our interest calculation logic lies
    |- go.mod
    |- go.sum
    |- README.md
    |- server.go # A thin `gorilla mux` server

## How to Run the package

### Prerequisiutes

You need to have the following before running the package:

    - Go 1.21+
    - Curl

### Run the server

1. To run the server, open your terminal and run the following command(s):

   ```bash
   $ go run server.go
   |
   ```

   The server run on `http://localhost:8000` on your local machine

2. Using any application/tool of your choice, call the endpoint to calculate the interest by supplying your balance as a query parameter valus as illustrated below

   ```bash
   $ curl http://localhost:8000/interest/{balance}
   ```

3. Calculate interest of any given balance

   ```bash
   $ curl http://localhost:8000/interest/1000
   {"interest":"15"}
   ```

### Run Tests

The package is covered by behaviour, table driven tests which can be run using the command:

```bash
$ go test ./...
```
