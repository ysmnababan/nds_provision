# Simple CRUD Program using PostgreSQL and Gin Gonic Framework

This is a simple CRUD application built with the Gin Gonic Framework and PostgreSQL.

## How to Run This Program

1. **Create Your Local Database**
   - Set up a PostgreSQL database on your local machine.

2. **Create the `.env` File**
   - Create a `.env` file in the project root directory with the necessary PostgreSQL configuration details.
   - Example:
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=your_username
     DB_PASSWORD=your_password
     DB_NAME=your_database_name
     ```

3. **Data Seeding (Optional)**
   - If you need to seed the database with initial data, use the `dataseed` function in the `main.go` file.

4. **Import Postman Collection**
   - Import the Postman collection file `nds_basic_crud.postman_collection.json` to test the endpoints.

5. **Run the Program**
   - Execute the `main.go` file on your local machine:
     ```bash
     go run main.go
     ```

6. **Test the Endpoints**
   - Use Postman to test the available CRUD endpoints.