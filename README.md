# Mongorm

Welcome to **Mongorm** – your new best friend in the MongoDB ORM world, crafted with love and a sprinkle of humor by Tahcin Ul Karim (a.k.a. Mycin), a student at Notre Dame College and a passionate programmer who’s been coding for the past 5 years. Whether you're a Go enthusiast or just someone looking to simplify MongoDB interactions, Mongorm is here to make your life easier. 

## 📜 Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Getting Started](#getting-started)
  - [Initialization](#initialization)
  - [Creating Documents](#creating-documents)
  - [Reading Documents](#reading-documents)
  - [Updating Documents](#updating-documents)
  - [Deleting Documents](#deleting-documents)
- [Contributing](#contributing)
- [License](#license)

## 📝 Introduction

Mongorm is an ORM (Object-Relational Mapping) tool for MongoDB written in Go. Inspired by [Abhishek Ranjan’s blog post](https://medium.com/@abhishekranjandev/building-a-gorm-like-orm-for-mongodb-with-golang-9812d43e2b78), this project brings a GORM-like experience to MongoDB. With Mongorm, you’ll be able to seamlessly interact with MongoDB collections using familiar Go idioms.

## 💾 Installation

To get started with Mongorm, you need to install it via Go modules. Here’s how:

```bash
go get github.com/mmycin/mongorm
```

## 🚀 Getting Started

### Initialization

First things first, you need to initialize a connection to your MongoDB database. Here’s a sample code snippet to get you started:

```go
package main

import (
    "context"
    "fmt"

    "github.com/mmycin/mongorm"
    "github.com/mmycin/mongorm/utils"
)

func main() {
    // Returns the client object of type `*mongo.Client` and error. If you want, you can handle the client onject if you want or leave it
    _, err := mongorm.Initialize("mongodb+srv://username:password@cluster0.mongodb.net/", "testdb")
    utils.HandleError(err)
    fmt.Println("MongoDB connected successfully!")
}
```

Replace `"mongodb+srv://username:password@cluster0.mongodb.net/"` with your MongoDB connection string and `"testdb"` with your database name.

### Creating Documents

To create a new document, use the `CreateOne` function:

```go
package main

import (
    "fmt"

    "github.com/mmycin/mongorm"
    "github.com/mmycin/mongorm/model"
    "github.com/mmycin/mongorm/utils"
)

type User struct {
	model.BaseModel `json:"-"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
}

func main() {
    err := mongorm.Initialize("mongodb+srv://username:password@cluster0.mongodb.net/", "testdb")
    utils.HandleError(err)

    user := User {
        Name:  "John Doe",
        Email: "john@example.com",
        Phone: "01234567800"
    }

    err = mongorm.CreateOne("users", &user)
    utils.HandleError(err)
}
```

### Reading Documents

To read all documents or find specific ones, use the `ReadAll` function:

```go
package main

import (
    "fmt"

    "github.com/mmycin/mongorm"
    "github.com/mmycin/mongorm/model"
    "github.com/mmycin/mongorm/utils"
)

type User struct {
	model.BaseModel `json:"-"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
}

func main() {
    err := mongorm.Initialize("mongodb+srv://username:password@cluster0.mongodb.net/", "testdb")
    utils.HandleError(err)

    var users []User
    err = mongorm.ReadAll("users", &users)
    utils.HandleError(err)

    fmt.Println("Users in the collection:")
    for _, u := range users {
        fmt.Printf("Name: %s, Email: %s\n", u.Name, u.Email)
    }
}
```

### Updating Documents

To update an existing document:

```go
package main

import (
    "context"
    "fmt"

    "github.com/mmycin/mongorm"
    "github.com/mmycin/mongorm/model"
    "github.com/mmycin/mongorm/utils"
)

func main() {
    err := mongorm.Initialize("mongodb+srv://username:password@cluster0.mongodb.net/", "testdb")
    utils.HandleError(err)

    filter := utils.Json{"name": "John Doe"}
    update := utils.Json{"email": "john.doe@example.com"}

    err = mongorm.Update("users", filter, update)
    utils.HandleError(err)

    fmt.Println("User updated successfully!")
}
```

### Deleting Documents

To delete a single document:

```go
package main

import (
    "context"
    "fmt"

    "github.com/mmycin/mongorm"
    "github.com/mmycin/mongorm/utils"
)

func main() {
    err := mongorm.Initialize("mongodb+srv://username:password@cluster0.mongodb.net/", "testdb")
    utils.HandleError(err)

    filter := utils.Json{"name": "John Doe"}
    err = mongorm.DeleteOne("users", filter)
    utils.HandleError(err)

    fmt.Println("User deleted successfully!")
}
```

## 🤝 Contributing

Got ideas or found a bug? We’d love to hear from you! Check out our [Contributing Guidelines](CONTRIBUTING.md) for more information on how to get involved.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

