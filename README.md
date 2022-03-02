# Hacktiv8 Golang Course - Final Project
This is a GitHub Repo for Hactiv8 Golang Course - Final Project

## Database Setup
Modifikasi path dsn sesuai dengan konfigurasi masing-masing (folder database -> file database-config.go)

```
db_user = // <- gunakan credentials masing-masing
db_pass = // <- gunakan credentials masing-masing
db_host = // <- gunakan credentials masing-masing
db_port = // <- gunakan credentials masing-masing
db_name = // <- gunakan credentials masing-masing
```

```
dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",db_user, db_pass, db_host, db_port, db_name)
```

## Postman API Setup
Import file Postman API yang ada di GitHub Repo sebagai collection di local machine anda
