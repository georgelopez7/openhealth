POST http://localhost:/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Ash",
  "last_name": "Ketchum",
  "age": 10,
  "email": "ash.ketchum@pokemon.com"
}

HTTP 201
[Captures]
ash_id: jsonpath "$.account.id"


POST http://localhost:/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Misty",
  "last_name": "Waterflower",
  "age": 12,
  "email": "misty.waterflower@pokemon.com"
}

HTTP 201
[Captures]
misty_id: jsonpath "$.account.id"


POST http://localhost:/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Professor",
  "last_name": "Oak",
  "age": 65,
  "email": "professor.oak@pokemon.com"
}

HTTP 201
[Captures]
oak_id: jsonpath "$.account.id"


POST http://localhost:/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Nurse",
  "last_name": "Joy",
  "age": 24,
  "email": "nurse.joy@pokemon.com"
}

HTTP 201
[Captures]
joy_id: jsonpath "$.account.id"
