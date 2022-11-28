# API Endpoints

## POST /api/v1/user
Request body in JSON format:
```json
{
    "name": "User name"
}
```
Response body in JSON format:
```json
{
    "ID": "3fc17c00-3510-4fd6-9f4e-605ef1e98005",
    "Name": "Mike",
    "Balance": 1000,
    "VerificationStatus": false
}
```

## POST /api/v1/transaction
Request body in JSON format:
```json
{
    "sender": "3fc17c00-3510-4fd6-9f4e-605ef1e98005",
    "receiver": "3fc17c00-3510-4fd6-9f4e-605ef1e98005",
    "amount": 100
}
```

## GET /api/v1/users
Response body in JSON:
```json
[
    {
        "ID": "608fc271-20f4-4da8-8eef-e20657558142",
        "Name": "Mark",
        "Balance": 1500,
        "VerificationStatus": true
    },
    {
        "ID": "3fc17c00-3510-4fd6-9f4e-605ef1e98005",
        "Name": "Max",
        "Balance": 500,
        "VerificationStatus": true
    }
]
```