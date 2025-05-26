# Posting API
API ini digunakan untuk mengelola dan memproses sebuah postingan.

## Dokumentasi 
dokumentasi di publish dengan [postman](https://documenter.getpostman.com/view/26786172/2sB2qcCgRb)

## Fitur
- autentikasi login dan register
- autorisasi pada path tertentu
- membuat, mengambil, menghapus, dan memperbarui postingan
- menyukai dan batal menyukai postingan
- memberi komentar terhadap postingan

## Tech Stack
- Go
- Gorm
- Echo
- PostgreSQL
- Testify

## Pattern
Project pattern yang sering saya gunakan adalah layered pattern dengan struktur controller - service - repository. Alasan nya dikarenakan dengan pattern seperti ini sudah cukup untuk memisahkan tanggung jawab, menjadikan kode menjadi terstruktur, dan mudah maintenance.
- controller : digunakan sebagai layer yang mengatur request dan response dari client
- service : tempat dimana business logic disimpan
- repository : digunakan untuk berinteraksi dengan database
