POST http://localhost:8000/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Ash",
  "last_name": "Ketchum",
  "age": 10,
  "email": "ash.ketchum@pokemon.com",
  "avatar": "ash-ketchum"
}

HTTP 201
[Captures]
ash_id: jsonpath "$.account.id"


POST http://localhost:8000/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Misty",
  "last_name": "Waterflower",
  "age": 12,
  "email": "misty.waterflower@pokemon.com",
  "avatar": "misty-waterflower"
}

HTTP 201
[Captures]
misty_id: jsonpath "$.account.id"


POST http://localhost:8000/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Professor",
  "last_name": "Oak",
  "age": 65,
  "email": "professor.oak@pokemon.com",
  "avatar": "professor-oak"
}

HTTP 201
[Captures]
oak_id: jsonpath "$.account.id"


POST http://localhost:8000/api/v1/accounts
Content-Type: application/json
{
  "first_name": "Nurse",
  "last_name": "Joy",
  "age": 24,
  "email": "nurse.joy@pokemon.com",
  "avatar": "nurse-joy"
}

HTTP 201
[Captures]
joy_id: jsonpath "$.account.id"


POST http://localhost:8000/api/v1/doctors
Content-Type: application/json
{
  "account_id": "{{oak_id}}"
}

HTTP 201
[Captures]
doctor_id: jsonpath "$.doctor.id"


POST http://localhost:8000/api/v1/doctors/assignments
Content-Type: application/json
{
  "account_id": "{{ash_id}}",
  "doctor_id": "{{doctor_id}}"
}

HTTP 201
[Captures]
oak_to_ash_assignment_id: jsonpath "$.assignment.id"


POST http://localhost:8000/api/v1/nurses
Content-Type: application/json
{
  "account_id": "{{joy_id}}"
}

HTTP 201
[Captures]
nurse_id: jsonpath "$.nurse.id"


POST http://localhost:8000/api/v1/hospitals
Content-Type: application/json
{
  "name": "Pokemon Center"
}

HTTP 201
[Captures]
hospital_id: jsonpath "$.hospital.id"


POST http://localhost:8000/api/v1/nurses/assignments
Content-Type: application/json
{
  "nurse_id": "{{nurse_id}}",
  "hospital_id": "{{hospital_id}}"
}

HTTP 201
[Captures]
joy_to_pokemon_center_assignment_id: jsonpath "$.assignment.id"


POST http://localhost:8000/api/v1/hospitals/assignments
Content-Type: application/json
{
  "account_id": "{{ash_id}}",
  "hospital_id": "{{hospital_id}}"
}

HTTP 201
[Captures]
ash_to_pokemon_center_assignment_id: jsonpath "$.assignment.id"


POST http://localhost:8000/api/v1/medical-records
Content-Type: application/json
{
  "account_id": "{{ash_id}}",
  "title": "Ash Annual Checkup",
  "description": "Routine annual physical examination for Ash Ketchum."
}

HTTP 201
[Captures]
ash_medical_record_id: jsonpath "$.medical_record.id"


POST http://localhost:8000/api/v1/medical-records
Content-Type: application/json
{
  "account_id": "{{misty_id}}",
  "title": "Misty Water Exposure Assessment",
  "description": "Assessment following extended exposure to Water-type Pokemon."
}

HTTP 201
[Captures]
misty_medical_record_1_id: jsonpath "$.medical_record.id"


POST http://localhost:8000/api/v1/medical-records
Content-Type: application/json
{
  "account_id": "{{misty_id}}",
  "title": "Misty Starmie Incident",
  "description": "Follow-up visit after Starmie Psybeam-related dizziness."
}

HTTP 201
[Captures]
misty_medical_record_2_id: jsonpath "$.medical_record.id"
