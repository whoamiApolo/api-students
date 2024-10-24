# API Students
API Restful to manage students

# Routes:
- GET /students - List all students
- POST /students - Create a student
- GET /students/:id - Show a student
- PUT /students/:id - Update a student
- DELETE /students/:id - Delete a student

# Struct Student:
- Name string
- CPF int
- Email string
- Age int
- Active bool


# Curl

- curl localhost:8081/students