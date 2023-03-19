# Resume Materi Join - Union - Agregasi - Subquery - Function (DBMS)

## Join
Sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel. beberapa standar join sql diantaranya ada inner join, left join dan right join.

#### Inner Join
Inner join mengembalikan baris-baris dari dua tabel atau lebih yang memenuhi syarat. Berikut adalah contoh query inner join:
```
SELECT tabel1.name, tabel1.age, tabel2.city, tabel2.country
FROM tabel1
INNER JOIN tabel2 ON tabel1.id = tabel2.id;
```

#### Left Join
Left join akan mengembalikan seluruh baris dari tabel disebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel sebelah kanan yang memenuhi kondisi join. Berikut adalah contoh query left join:
```
SELECT tabel1.name, tabel1.age, tabel2.city, tabel2.country
FROM tabel1
LEFT JOIN tabel2 ON tabel1.id = tabel2.id;
```

#### Right Join
akan mengembalikan seluruh baris dari tabel disebelah kanan yang dikenai kondisi ON dan hanya baris dari tabel sebelah kiri yang memenuhi kondisi join. Berikut adalah contoh query right join:
```
SELECT tabel1.name, tabel1.age, tabel2.city, tabel2.country
FROM tabel1
RIGHT JOIN tabel2 ON tabel1.id = tabel2.id;
```

## Union
Perintah union biasa digunakan untuk menggabungkan 2 atau lebih query select, yang perlu diperhatikan adalah jumlah field yang harus dikeluarkan atau dipanggil harus sama.
```
SELECT id, name, age FROM tabel1
UNION ALL
SELECT id, name, age FROM tabel2;

```

## Fungsi agregate
Fungsi agregate digunakan untuk mengelompokkan beberapa baris untuk membentuk nilai ringkasan tunggal. Beberapa fungsi agregat sql yang sering digunakan diantaranya ada min, max, sum, avg, count dan having.
<ul>
<li>min : mencari nilai minimum</li>
<li>max : mencari nilai maksimum</li>
<li>avg : mencari nilai rata-rata</li>
<li>sum : mencari jumlah nilai suatu kolom</li>
<li>count : mendapatkan jumlah nilai dari sebuah kolom</li>
<li>having : digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi agregat</li>
</ul>

## Subquery
Subquery atau inner query atau nested query adalah query didalam query sql lain. Sebuah subquery digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. Berikut adalah contoh query subquery :
```
SELECT nama_kota, jumlah_pelanggan
FROM (SELECT kota_id, COUNT(*) as jumlah_pelanggan
      FROM tabel_pelanggan
      GROUP BY kota_id) as subquery
JOIN tabel_kota ON tabel_kota.id = subquery.kota_id
```

## Function
Function adalah sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya. Berikut adalah contoh query function :
```
CREATE FUNCTION hitung_total_harga (harga_asli INT, diskon INT) 
RETURNS INT
BEGIN
   DECLARE harga_diskon INT;
   SET harga_diskon = harga_asli - (harga_asli * diskon / 100);
   RETURN harga_diskon;
END;
```

## Trigger
Trigger biasanya digunakan untuk mentrigger proses query lain, baik sesudah atau sebelum proses dijalankan. Berikut adalah contoh penggunaan trigger :
```
CREATE TRIGGER nama_trigger 
BEFORE/AFTER INSERT/UPDATE/DELETE ON nama_tabel 
FOR EACH ROW
BEGIN
   -- kode trigger di sini
END;
```