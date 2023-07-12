# GoRestApiExample
A simple CRUD json rest api using Go, Gin.

# Features
    - Structured project folders
    - Sperate handler interface and database interface
    - JWT token authentication with access and refresh token with token revokation mechanism using redis
    - Strong custom Password Validation
    - Log to file


## .env file
m_db_username=your own choices
m_db_password=your own choices
m_db_dbname=your own choices
postgres="user=your own choices password=your own choices dbname=your own choices sslmode=disable"
test="user=your own choices password=your own choices dbname=your own choices sslmode=disable"
PORT=3000
jwt_secret=your secret key