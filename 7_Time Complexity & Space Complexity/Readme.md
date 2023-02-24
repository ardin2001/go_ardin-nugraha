# Resume Time Complexity & Space Complexity

## Time Complexity
Penggunaan dari kompleksitas waktu membuat memudahkan untuk memperkirakan waktu berjalan sebuah program. Kompleksitas bisa dilihat dari jumlah maksimum pada operasi primitif yang dapat dijalankan suatu program.

## Apa Itu Operasi Dominan ?
![operasi dominan] (https://github.com/ardin2001/go_ardin-nugraha/blob/1ef44448edcceee28b17df07fabd2707734a3556/7_Time%20Complexity%20&%20Space%20Complexity/screenshots/materi/no1.jpg)
Pada gambar diatas, operasi pada baris ke 4 adalah dominan dan akan dijalankan n waktu. Kompleksitas mendeskripsikan dalam notasi Big-O: dalam kasus O(n) - kompleksitas linier.

## Jenis-jenis Complexcity
Berikut beberapa jenis-jenis complexcity, diantaranya :
<ul>
<li>Constant Time - O(1)</li>
Pada jenis complexcity ini yaitu selalu ada jumlah operasi yang tetap. Maksutnya yaitu eksekusi program yang dilakukan pada program ini berjalan selama satu kali.
<li>Linear Time - O(n)</li>
Pada jenis complexcity ini adalah kode program akan dieksekusi sebanyak jumlah dari n input. Semakin besar parameter n, maka eksekusi program akan semakin besar.
<li>Linear Time - O(n+m)</li>
Pada jenis complexcity ini adalah kode program akan dieksekusi sebanyak dari jumlah parameter n+m. Semakin besar parameter m dan n, maka eksekusi program akan semakin besar.
<li>Logarithmic Time - O(log n)</li>
Pada jenis complexcity ini berdasarkan setengah dari jumlah n setiap iterasi dari perulangan.
<li>Quadratic Time - O(n^2)</li>
Pada jenis complexcity ini program akan dieksekusi sebanyak n^2. Program jenis ini biasanya menggunakan nesteed loop.
</ul>

## Exponential and Factorial Time
Perlu diketahui bahwa ada jenis lain dari kompleksitas waktu seperti factorial time O(n!) dan Exponential time O(2^n). Algoritma kompleksitas seperti itu hanya dapat menyelesaikan masalah untuk nilai n yang sangat kecil, karena akan memakan waktu terlalu lama untuk eksekusi nilai n yang lebih besar.

## Space Complexcity
Batas memory memberikan informasi tentang kompleksitas ruang. Anda dapat memperkirakan jumlah variabel yang dapat anda deklarasikan pada program anda. singkatnya, jika anda mempunyai jumlah variabel yang konstan, anda juga mempunyai kompleksitas ruang yang konstan: dalam notasi Big-O ini adalah O(1). Jika anda membutuhkan mendeklarasikan array dengan n elemen, anda memiliki kompleksitas ruang yang linier- )(n).
