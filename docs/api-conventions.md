# API Conventions

## Sequence Diagram
```mermaid
sequenceDiagram
    participant Client
    participant API
    participant Database
    Client->>API: makes http requests
    alt if contract is broken
        API-->>Client: 400 Bad Request
    end
    alt if business rule has failed
        API-->>Client: 422 Unprocessable Entity
    end
        API->>Database: does database operation
    alt if resource was created successfully
        Database-->>API: database response
        API-->>Client: 201 Created
    end
    alt if resource was found
        Database-->>API: database response
        API-->>Client: 200 OK
    end
    alt if resource was not found
        Database-->>API: database response
        API-->>Client: 404 Not Found
    end
    alt if resource was updated/deleted successfully
        Database-->>API: database response
        API-->>Client: 204 No Content
    end
        alt if database operation fails
        Database-->>API: failed operation
        API->>API: log error
        API-->>Client: 500 Internal Server Error
    end
```

## Success Response
```json
{
    "data": {}
}
```

## Error Response
```json
{
    "error": {
        "type": "string",
        "code": "string",
        "message": "string",
        "details": ["string"]
    }
}
```