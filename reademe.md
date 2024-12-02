# Jwt Auth

### Pakages

```bash
    #JWT: a Go library for JSON Web Token implementation.
   go  get github.com/gin-gonic/gin

   #DotEnv: a Go dependency for managing environment variables.
   go  get github.com/golang-jwt/jwt/v4

   #Gorm: an ORM (Object Relational Mapper) for Golang.
   go  get github.com/joho/godotenv

   #Go Crypto: a cryptography library in Go that will be used for password encryption and decryption.
   go  get golang.org/x/crypto

   #Driver for postgres
   go  get gorm.io/driver/postgres

   # gorm for type orm
   go  get gorm.io/gorm

   #CompileDaemon: a Go package facilitating automatic app reloading upon saving changes
   go  get github.com/githubnemo/CompileDaemon
```

### Run Application:

```bash
    docker compose up
    go run main.go
```

```bash
go run migrate/migrate.go
```

### Migrate Tables

```bash
go run migrate/migrate.go
```

[Generate Token](https://github.com/golang-jwt/jwt)

### Drop Tables

```bash
go run scripts/drop.table.go
```

### Truncate Tables

```bash
go run scripts/truncate.table.go
```

### Make an Super Admin

```bash
    go run scripts/makeSuperAdmin.go
```
