# Save Transaction Flow
This flow describe how a transaction can be saved and associated to an account. Many transactions can be associated to an account.
## Sequence Diagram
```mermaid
sequenceDiagram
    participant Client
    participant API
    participant Database
    Client->>API: POST /transactions
    Note right of Client: account_id: int<br/>operation_type_id: int<br/>amount: int
    API->>API: check transaction business rules
    alt if transaction has inconsistent state
        API-->>Client: 422 Unprocessable Entity
        Note right of Client: error: [wrong operation_type_id,<br/> amount must be negative]
    end
    API->>Database: select from 'accounts' where account_id
    alt if account not exists
        Database-->>API: account record not found
        API-->>Client: 422 Unprocessable Entity
        Note right of Client: error: [account not exists]
    end
    API->>Database: insert transaction into 'transactions' table
    Database-->>API: inserted succesfully
    API-->>Client: 201 Created
    Note right of Client: account_id: int
```
## Requirements
* A transaction can only be saved if given account exists.
* Account ID, Operation Type ID and Amount fields are required.
* There are four type of operation types:

    | Operation Type  | ID | Rules |
    |---|---|---|
    | Purchase | 1 | amount must be negative |
    | Purchase In Installments | 2 | amount must be negative |
    | Withdraw | 3 | amount must be negative |
    | Payment | 4 | amount must be positive |