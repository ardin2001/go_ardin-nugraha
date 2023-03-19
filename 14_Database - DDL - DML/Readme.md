# Resume Materi Database Schema & Data Definiton Language

## Database
Database adalah sekumpulan data yang terorganisir. Biasanya data berisi beberapa data atau informasi kemudian disimpan didalam 1 tempat tersendiri. Database teridir dari dua jenis yaitu database relationship dan database relationship (no sql).

#### Database Relationship
Database retationship terdiri dari beberapa jenis relasi, ada one to one, one to many dan many to many. Berikut adalah contoh masing-masing relasi:

One to one relationship
<ol>
<li>1 user hanya memiliki 1 foto profil.</li>
<li>Satu siswa hanya memiliki satu nomor induk siswa (NIS), dan satu NIS hanya bisa dimiliki oleh satu siswa.</li>
<li>Satu akun email hanya bisa dimiliki oleh satu orang, dan satu orang hanya bisa memiliki satu akun email.</li>
</ol>
One to Many relationship
<ol>
<li>Satu produsen (manufacturer) dapat memiliki banyak produk, tetapi setiap produk hanya diproduksi oleh satu produsen.</li>
<li>Satu kategori barang (category) dapat memiliki banyak produk, tetapi setiap produk hanya termasuk dalam satu kategori.</li>
<li>Satu perusahaan (company) dapat memiliki banyak karyawan (employee), tetapi setiap karyawan hanya bekerja untuk satu perusahaan.</li>
</ol>
Many to Many relationship
<ol>
<li>SHubungan antara banyak mahasiswa dengan banyak mata kuliah. Satu mahasiswa bisa mengambil banyak mata kuliah, dan satu mata kuliah bisa diambil oleh banyak mahasiswa.</li>
<li>Hubungan antara banyak penulis dengan banyak buku. Banyak penulis bisa menulis banyak buku, dan banyak buku bisa ditulis oleh banyak penulis.</li>
<li>Hubungan antara banyak pasien dengan banyak dokter. Satu pasien bisa memiliki banyak dokter, dan satu dokter bisa memiliki banyak pasien.</li>
</ol>
Untuk mengimplementasikan pembuatasan database relationship ini. Bisa menggunakan beberapa software Relational Database Management System, diantaranya adalah mysql.

## Jenis Perintah SQL
Jenis-jenis perintah sql dibagi menjadi 3, yang pertama ada Data Definiton Language, Data ManipulationLanguage dan Data Control Language.

#### Data Definiton Language
Untuk pengertian DDL dan DML sendiri yaitu, jika DDL (Data Defition Language) adalah kumpulan perintah SQL yang dapat anda gunakan untuk mengelola, mengubah struktur datatype dari objek pada database seperti index, table, trigger, view dan lain sebagainya. Beberapa query yang biasa digunakan dalam DDL diantaranya :
<ul>
<li>CREATE DATABASE database_name</li>
<li>USE database_name</li>
<li>CREATE TABLE</li>
<li>DROP TABLE</li>
<li>RENAME TABLE</li>
</ul>
Kemudian untuk membuat schema table bisa menggunakan query berikut :

```
CREATE TABLE operator(
    id VARCHAR (20) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR (50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
```

Kemudian untuk memodifikasi skema table dapat membuat query berikut :

```
ALTER TABLE table_name ADD COLUMN column_name data_type;
```

## Data Manipulation Language
Data Manipulation Language atau biasa disingkat DML adalah perintah yang digunakan untuk memanipulasi data dalam table dari suatu database. Untuk statement yang bisa digunakan ada insert, select, update dan delete.

#### INSERT
Insert biasa digunakan untuk memasukkan data kedalam kolom table, berikut adalah contoh querynya :
```
INSERT INTO nama_tabel (kolom1, kolom2, kolom3)
VALUES ('nilai_kolom1', 'nilai_kolom2', 'nilai_kolom3');

```
#### SELECT
Perintah select biasa digunakan untuk menampilkan kolom dalam tabel, berikut adalah contoh querynya :
```
SELECT nama_kolom1,nama_kolom2 FROM mahasiswa;
```

#### UPDATE
Perintah query update biasanya digunakan untuk mengupdate data dalam table, berikut adalah contoh querynya :
```
UPDATE nama_tabel
SET kolom1 = 'nilai_baru1', kolom2 = 'nilai_baru2'
WHERE kondisi;
```

#### DELETE
Perintah delete biasa digunakan untuk menghapus data dalam table, berikut adalah contoh querynya :
```
DELETE FROM students
WHERE name = 'John Doe';
```

## DML Statement
DML statement biasanya digunakan untuk menampilkan data tetapi dengan menambahkan kondisi atau statement tertentu. Diantaranta ada statement like/between, and/or, order by dan limit.

#### LIKE / BETWEEN
```
SELECT * FROM nama_tabel
WHERE nama_kolom LIKE 'A%';
SELECT * FROM nama_tabel
WHERE nama_kolom BETWEEN 10 AND 20;
```

#### AND / OR
```
SELECT * FROM nama_tabel
WHERE nama_kolom BETWEEN 10 AND 20
OR nama_kolom2 = 'nilai_tertentu';
```

#### ORDER BY
```
SELECT * FROM nama_tabel
ORDER BY nama_kolom ASC, nama_kolom2 ASC;
```

#### LIMIT
```
SELECT * FROM nama_tabel
LIMIT 5;
```
