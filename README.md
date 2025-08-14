# app-cms
Deskripsi singkat tentang proyek Golang ini.

## 🚀 Fitur

- Create Article [status : draft]
- Update Article [status : publish]
- Get Articles
- Get Article by ID

## Implementation
- Ratelimit -> handle by ip per request second
- Singleflight -> jika ada banyak request and cache expired, 
  hanya akan ada 1 request yang melakukan query ke DB dan write ke cache

## 🛠️ Teknologi

Service ini didevelop dengan:

- [Go](https://golang.org/) versi 1.23
- Modul Go (`go mod`)
- Database: PostgreSQL
- Cache: Redis
- Tambahan: Docker

## 🧑‍💻 How Run News Service

```bash
# clone repository new
git clone https://github.com/dika22/app-cms.git
cd nama-project

# Important
Status Article 1: draft 2: publish 3: archived

# set .env
cp -R .env.copy to .env
create name db
# how migrate
make migrate

# generate swagger
swag init or make swag

# Cara menjalankan http 
make http-serve

# Cara menjalankan worker
make start-worker

# how run unit test
make test


# how run swagger port sesuaikan dengan yang di .env
http://localhost:8001/swagger/index.html
note : sesuaikan alamat url
```


## 🧑‍💻 How Run Auth Service

```bash
cd nama-project atau auth-service

# important
Role : 1 SuperAdmin, 2 : Editor 3 : Writer 4 : user biasa

# set .env
cp -R .env.copy to .env
create name db
# how migrate
make migrate

# generate swagger
swag init or make swag

# Cara menjalankan http 
make http-serve

# Cara menjalankan worker
make start-worker

# how run unit test
make test


# how run swagger port sesuaikan dengan yang di .env
http://localhost:3000/swagger/index.html
note : sesuaikan alamat url
```
# Noted :
```bash 
- Seharusnya untuk GET implementasi REDIS 
```