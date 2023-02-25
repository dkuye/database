# database
Gorm ORM RDMS Database connector for mysql, sqlserver & postgres.

## Installation
```bash
go get github.com/dkuye/database
```
## Usage
First you need to have the following in your .env file for your project.
```
DB_CONNECTION=[ mysql | sqlserver | postgres ]
DB_HOST=127.0.0.1
DB_PORT=1234
DB_DATABASE=myName
DB_USERNAME=myUser
DB_PASSWORD=myPassword
```

Additional optional paramaters to a connection:

```
TIME_ZONE=Africa/Lagos
```

MySQL specifics
```
CHAR_SET=utf8mb4
PARSE_TIME=True
```

PostgreSQL specifics
```
SSL_MODE=[ disable | allow | prefer | require | verify-ca | verify-full ]
```


```go
package main

import (
    "github.com/dkuye/database"
    "github.com/joho/godotenv"
    "log"
)

func main(){
    // Open .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    // Initial database connection
    database.Init()
    // Connect to you database
    db := database.DBConnect
    // user gorm ORM query
    db.Exec("[your query]")
}
```

