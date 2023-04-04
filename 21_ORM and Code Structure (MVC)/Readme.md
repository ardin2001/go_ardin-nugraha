# Resume Materi ORM and Code Structure

## ORM
Pemetaan objek-relasional (alat pemetaan ORM, O/RM, dan O/R) dalam ilmu komputer adalah teknik pemrograman untuk mengubah data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek. ORM biasanya digunakan pada database berbasis RDBMS.

## ORM Pada GO
Pada materi ini mempelajari mengenai ORM pada golang yaitu Gorm. Kemudian mempelajari mengenai database migration. membuat database berdasarkan struktur model. mengenai instalasi dan konfigurasi gorm dengan database tertentu bisa mengunjungi website resmi gorm.

## Struktur Project
Menggunakan konsep MVC dimana project dibagi berdasarkan tugasnya, seperti routing, config, controllers, dan model.

<ul>
<li>Routing : bertugas untuk mengatur route-route pembagian request dan response kepada user</li>
<li>Config : bertugas untuk melakukan konfigurasi seperti konfigurasi pada database</li>
<li>Controllers: bertugas untuk mengontrol antara user dengan model</li>
<li>Models : bertugas sebagai representasi objek pada database</li>
</ul>

