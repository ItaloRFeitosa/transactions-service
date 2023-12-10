# Use integer to store amount as cents in database
## Context and Problem Statement
* How we store money amount in database?
* How we handle multiple currencies?
* How we ensure precision in amount operations?
## Considered Options
* Store as Double
* Store as Decimal
* Store as Integer
## Decision Outcome
Chosen option: store amount as integer, because it's easier to ensure precision in operations and to accept multiple currencies in the future.

## Pros and Cons of the Options
###  Store as Integer
* Good, because integers are safe to do operations, like, sum, subtraction, multiply and divide, without have the problem of rounding or interpolation.
* Good, because it can support other currencies, just needing to define the length of decimal part. (At the moment just 2 digits decimal is being accepted)
* Bad, because it needs to handle in the application level an abstraction to handle the decimal and currency logic. Can be handle by third party lib:
    * https://github.com/shopspring/decimal
    * https://pkg.go.dev/github.com/Rhymond/go-money
###  Store as Double
* Good, because it can be handle out of the box in application level.
* Good, because it's easy to work with multiple currencies.
* Bad, because loose precision when doing decimal operations at application level. ref: https://go.dev/play/p/TQBd4yJe6B
###  Store as Decimal
* Good, because it can be handle out of the box in application level.
* Bad, because we have to add extra decimal fixed points to support multiple currencies, and adds more complexity to application level.
* Bad, because loose precision when doing decimal operations at application level. ref: https://go.dev/play/p/TQBd4yJe6B