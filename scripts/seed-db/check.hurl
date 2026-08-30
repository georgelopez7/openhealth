GET http://openhealth-api:8000/api/v1/accounts
Authorization: Bearer {{API_AUTH_TOKEN}}
HTTP 200
[Asserts]
jsonpath "$.accounts" count == 0
