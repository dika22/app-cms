# Auth Service

## 🧑‍💻 How Run Auth Service

```bash
cd nama-project atau auth-service

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