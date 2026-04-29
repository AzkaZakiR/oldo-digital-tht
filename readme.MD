# Golang API

## Live URL APLIKASI

`https://oldodigital.azkazk11.my.id/`

# 📖 Deskripsi Project

Project ini adalah RESTful API sederhana yang dibangun menggunakan bahasa Go (Golang). Aplikasi ini menyediakan fitur CRUD (Create, Read, Update, Delete) untuk mengelola data user, paket data dan juga transaksi menggunakan framework web Fiber, ORM GORM, dan database PostgreSQL.

# 🛠 Tech Stack

- Golang – Bahasa pemrograman utama
- Fiber – Framework web
- GORM – ORM untuk database
- PostgreSQL – Database relasional
- Air / Fresh – Live reload saat development

## 🚀 Cara Menjalankan Project

### 1. Clone Project

```bash
git clone https://github.com/AzkaZakiR/oldo-digital-tht.git
cd oldo-digital-tht
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Setup ENV

create file .env di root project

```bash
DB_HOST=localhost
DB_USER=postgres
DB_NAME=oldo_digital(sesuaikan)
DB_PORT=5432
DB_TIMEZONE=Asia/Jakarta
DB_PASSWORD=
```

## 🗄 Cara Menjalankan Database

Pastikan PostgreSQL sudah terinstall dan berjalan.

Buat database:

```SQL
CREATE DATABASE yourdb;
```

## ▶️ Cara Run Aplikasi

```bash
go run main.go
```

Atau jika pakai Live Reload

```bash
air
```

Aplikasi akan running di

```bash
http://localhost:3000
```

## 📡 Contoh Penggunaan Endpoint

### 1. Register user baru

|        |                                 |
| ------ | :-----------------------------: |
| URL    | http://localhost:3000/api/users |
| METHOD |              POST               |
| Input  |              JSON               |

untuk update tinggal mengganti method menjadi patch dan tambah /:id di belakang

```json
{
   "name": "",
   "email": "",
   "phoneNumber": "",
   "password": ""
}
```

### 2. Get All users

|        |                                 |
| ------ | :-----------------------------: |
| URL    | http://localhost:3000/api/users |
| METHOD |               GET               |
| Input  |                -                |

### 3. Get user by Id

|        |                                     |
| ------ | :---------------------------------: |
| URL    | http://localhost:3000/api/users/:id |
| METHOD |                 GET                 |
| Input  |                  -                  |

untuk delete tinggal mengganti method menjadi delete

### 4. Create Data plan (paket data)

|        |                                    |
| ------ | :--------------------------------: |
| URL    | http://localhost:3000/api/dataplan |
| METHOD |                POST                |
| Input  |                JSON                |

untuk update tinggal mengganti method menjadi patch dan tambah /:id di belakang

```json
{
   "name": "Paket 50 GB",
   "price": 50000,
   "quota": 50,
   "active_period": 30,
   "is_active": true
}
```

### 5. Get Data plans

|        |                                    |
| ------ | :--------------------------------: |
| URL    | http://localhost:3000/api/dataplan |
| METHOD |                GET                 |
| Input  |                 -                  |

### 6. Get Data plan by id

|        |                                        |
| ------ | :------------------------------------: |
| URL    | http://localhost:3000/api/dataplan/:id |
| METHOD |                  GET                   |
| Input  |                   -                    |

untuk delete tinggal mengganti method menjadi delete

### 7. create Transaksi

|        |                                        |
| ------ | :------------------------------------: |
| URL    | http://localhost:3000/api/transactions |
| METHOD |                  POST                  |
| Input  |                  JSON                  |

```json
{
   "user_id": 45,
   "data_plan_id": 2
}
```

### 8. Get Transactions

|        |                                        |
| ------ | :------------------------------------: |
| URL    | http://localhost:3000/api/transactions |
| METHOD |                  GET                   |
| Input  |                   -                    |

### 9. Get Transactions by id

|        |                                            |
| ------ | :----------------------------------------: |
| URL    | http://localhost:3000/api/transactions/:id |
| METHOD |                    GET                     |
| Input  |                     -                      |
