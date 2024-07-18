curl -X PUT http://localhost:8080/update -H "Content-Type: application/json" -d '{
    "id": "user-id-from-registration-response",
    "name": "testuser",
    "password": "newtestpassword",
    "score": 10
}'

