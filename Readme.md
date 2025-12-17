# Online Cinema Ticketing System

## Deskripsi
Online Cinema Ticketing System adalah platform pembelian tiket bioskop secara online
yang dirancang untuk skala nasional. Sistem ini memungkinkan customer membeli tiket
kapan saja dengan jaminan kursi tidak dapat digunakan oleh orang lain, serta mendukung
refund dan pembatalan jadwal oleh pihak bioskop secara otomatis.

---

## Tujuan Sistem
- Mempermudah pembelian tiket bioskop secara online
- Menjamin tidak terjadi double booking kursi
- Mendukung transaksi bersamaan dalam jumlah besar
- Menyediakan sistem pencatatan tiket yang akurat
- Menangani proses refund dan pembatalan secara otomatis

---

## Chart 1 – Flow Pembelian Tiket

Chart berikut menggambarkan alur utama pembelian tiket bioskop oleh customer,
mulai dari membuka aplikasi hingga e-ticket diterima.

![Chart 1 - Flow Pembelian Tiket](docs/chart-1.png)

---

## Chart 2 – Flow Refund & Pembatalan

Chart berikut menggambarkan alur sistem ketika terjadi pembatalan jadwal
oleh pihak bioskop dan proses refund kepada customer.

![Chart 2 - Flow Refund dan Pembatalan](docs/chart-2.png)

---

## Penjelasan Singkat Sistem

### Sistem Pemilihan Kursi
- Kursi memiliki status: `AVAILABLE`, `LOCKED`, dan `SOLD`
- Kursi akan dikunci sementara saat dipilih customer
- Jika pembayaran berhasil, kursi menjadi `SOLD`
- Jika gagal atau waktu habis, kursi kembali `AVAILABLE`

### Pencatatan dan Restok Tiket
- Setiap transaksi berhasil akan tercatat sebagai tiket terjual
- Restok tiket terjadi otomatis jika transaksi gagal atau dibatalkan
- Sistem selalu memperbarui status kursi secara real-time

### Refund dan Pembatalan
- Bioskop dapat membatalkan jadwal tayang
- Sistem akan memproses refund otomatis
- Kursi dikembalikan ke stok
- Customer menerima notifikasi

---

## Catatan
Dokumentasi ini dibuat sebagai gambaran rancangan sistem
dan dapat dikembangkan lebih lanjut ke tahap implementasi teknis.

---


Dokumentasi Database
https://dbdiagram.io/d/Technical-Test-MKP-6915a79b6735e111709ef61f