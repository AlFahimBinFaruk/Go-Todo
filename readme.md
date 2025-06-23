### Run DB migration
```cmd
go run ./migrate/migrate.go
```

### Run the application
1. With CompileDaemon for hot reload.
```cmd
CompileDaemon -command="./go-todo"
```
2. with go cli
```cmd
go run server.go
```
