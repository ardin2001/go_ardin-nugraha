# Resume Materi Clean and Hexagonal Architecture

##  Sebelum Migrasi
The goal ketika kita menggunakan clean code adalah kode kita lebih modular, scallabel, dan maintainable.

<ul>
<li>Modular dalam artian kita bisa dengan mudah mengganti dipendensi satu ke dipendensi lain</li>
<li>Scallabel dalam artian kita dapat dengan mudah menambahkan feature baru dan lain sebagainya</li>
<li>Maintainable dalam artian kita dapat dengan mudah memperbaiki issue bilamana terdapat issue pada kode kita</li>
</ul>

## Opsi Migrasi
Ada 3 opsi bila kita mau melakukan migrasi arsitektur kode dari mvc ke clean code:

<ul>
<li>Pertahankan desain sekarang dengan memisahkan dependensi.</li>
<li>Pertahankan desainnya tetapi kita pindahkan kodenya kedalam suatu layer.</li>
<li>Ubah desainnya dan pisahkan dependensi.</li>
</ul>

Pada contoh kali ini kita akan memakai cara kedua, dimana kita akan memisahkan kode kita ke dalam layer layer yaitu usecase dan repository.

## Struktur Code
Berikut struktur kode yang akan kita pakai nantinya,
<ul>
<li>Controller: berisi kode yang berhubungan langsung ke user (interface layer).</li>
<li>Repository: berisi kode yang berhubungan langsung ke database(lapisan antarmuka).</li>
<li>Usecase: berisi bisnis logik yang dipakai.</li>
</ul>
