# Open Account Flow
This flow describe how an card holder (client) can open account. Each account can have many transactions associated.
## Sequence Diagram
```mermaid
sequenceDiagram
    participant Client
    participant API
    participant Database
    Client->>API: POST /accounts
    Note right of Client: document_type: string<br/>document_number: string
    API->>API: validate document number to given document type
    alt if document number is invalid
        API-->>Client: 422 Unprocessable Entity
        Note right of Client: error: [invalid document number]
    end
    API->>Database: insert into 'accounts' table
    Database-->>API: inserted succesfully
    API-->>Client: 201 Created
    Note right of Client: account_id: int
```
## Requirements
* Just CPF, CNPJ and NINO document types are allowed.
* Document Type and Number are required fields