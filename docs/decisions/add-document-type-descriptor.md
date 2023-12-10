# Add document type descriptor to account data
## Context and Problem Statement
* How to handle/validate multi country document number with ease?
## Considered Options
* Add descriptor field
* Use regular expressions only
## Decision Outcome
Chosen Option: Add descriptor field, because it's straightforward to add new document types and its validations.
```json
{
    "document_type": "NINO",
    "document_number": "BZ960351B"
}
```