# database
Gorm ORM Database connector for all supported RDMS.

## Installation
```bash
go get github.com/dkuye/helper
```
## Usage
First you need to have the following in your .env file for your project.
```
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=name
DB_USERNAME=user
DB_PASSWORD=password
```

```go
package main

import (
    "fmt"
    "github.com/dkuye/database"
    "github.com/joho/godotenv"
)

func main(){
    // Open .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    // Connect to you database
    db := database.Connect()
    defer db.Close()
}
```

