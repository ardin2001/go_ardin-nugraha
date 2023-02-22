# Resume Version Control and Branch Management (Git)

## Apa itu Versioning?
Versioning adalah mengatur versi dari sebuah source code program. Mengatur yang dimaksut adalah mengatur dan memanajemen sebah source code mulai dari awal sampai beberapa perubahan atau development sampai benar benar program tahap akhir.

Terdapat beberapa tools untuk membantu versioning, diantaranya :
<ul>
<li>Version Control System</li>
<li>Source Code Manager</li>
<li>Revision Control System</li>
</ul>

## Version Control System (VCS)
Berikut beberapa contoh pada versioning VCS, diantaranya :
<ol>
<li>Single User</li>
Untuk pengembangan Version Control System yang pertama yaitu Single User. Disini masih berada pada local repository. Beberapa software yang digunakan ada SCCS dan RCS.
<li>Centralized</li>
Untuk pengembangan Version Control System yang kedua yaitu Centralized. Dimana metode penyimpanannnya sudah menggunakan repository server dengan bantuan internet. Beberapa contoh software yang digunakan diantaranya ada CVS, Perforce, Subversion dan Microsoft Team Foundation Server.
<li>Distributed</li>
Pada metode pengembangan Version Control System yang terakhir ini sudah menggunakan metode penyimpangan repository di server maupun di local. Beberapa aplikasi diantaranya ada Git, Bazaar dan Mercurial
</ol>

## GIT
Git adalah salah satu Version Control System yang digunakan oleh para developer untuk membantu dalam dalam proses pengembangan dan pembuatan software bersama-sama sehingga.

## The Staging Area
Ekosistem mengenai The Staging Area terdiri dari tiga bagian, diantaranya :

<ul>
<li>Working Directory</li>
Pada bagian ini, project masih terletak pada local repository. Untuk dapat dapat di letakkan pada repository server kita harus melakukan git add terlebih dahulu.
<li>Staging Area</li>
Setelah project dilakukan git add, maka project akan masuk ke bagian Staging Area. Dimana untuk membedakan dan memperlihatkan bahwa project akan di upload ke repository server.
<li>Repository</li>
Kemudian lakukan git commit agar project dapat ditempatkan di repository server.
</ul>

## Command Git
Terdapat beberapa command pada git yang digunakan untuk membantu dalam memanajemen versioning. Berikut beberapa command penting yang bisa digunakan :
<ul>
<li>git init : menginisialisasi sebuah project repository.</li>
<li>git clone : untuk melakukan cloning repository server ke local repository.</li>
<li>git status : untuk mengecek status local repository.</li>
<li>git add : untuk menambahkan file atau folder yang ingin di tambahkan ke server repository.</li>
<li>git commit : digunakan untuk melakukan commit project.</li>
<li>git diff : digunakan untuk melihat file yang mana terdapat perubahan.</li>
<li>git branch : digunakan untuk melihat daftar branch.</li>
<li>git branch [nama branch] : digunakan untuk membuat branch baru.</li>
<li>git checkout : digunakan untuk berpindah antar branch.</li>
<li>git reset : digunakan untuk mereset branch ke commit sebelumnya.</li>
<li>git pull : digunakan untuk mendownload isi file terbaru dari server repository.</li>
<li>git push : digunakan untuk mengupload project local repository ke server repository</li>
<li>git merge : digunakan untuk menggabungkan branch satu dengan yang lainnya.</li>
</ul>

# File .gitignore
File jenis ini biasa digunakan untuk membantu developer agar ketika melakukan push ke server repository beberapa file pada local repository tidak perlu di upload. Jenis file ini tergantung dari kemauan developer, biasanya beberapa file yang tidak perlu ikut dimasukkan ke server repository seperti .env, node_modules(pada javascript) dan beberapa package atau lain-lain.

## Pull request
Pull request digunakan untuk meminta izin menggabungkan branch lain ke branch main atau master.

## Workflow Collaboration
Terdapat banyak contoh workflow collaboration yang dapat diterapkan dalam project kita. Workflow Collaboration dapat diterapkan sesuai dengan aturan atau kesepakatan antar tim. Bisa dengan hanya satu branch main/master, atau dua branch main dengan dev dan menggunakan lebih dari dua branch lain.