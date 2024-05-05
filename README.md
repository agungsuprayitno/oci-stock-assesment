# API untuk menghitung Cicilan Pinjaman

## Tentang API

Rest API untuk melakukan perhitungan angsuran.

#### disini saya hanya menggunakan hardcode variabel sebagai berikut:

- Plafon Rp 5.000.000
- Suku Bunga 8 persen pertahun
- Lama Pinjaman 6 bulan


## Requirements

- Golang minimum versi 1.18 (saya menggunakan go version 1.22.2)


## Run API

- Clone dari github

  ```
  git run main.go
  ```

- Download dependencies

  ```
  go mod tidy
  ```

- Buka Browser untuk test API dan copy paste url berikut `http://localhost:8080/hitung-cicilan/`

