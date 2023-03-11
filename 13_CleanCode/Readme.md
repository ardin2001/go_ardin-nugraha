# Resume Materi Clean Code

## Apa itu Clean Code ?
Clean Code adalah istilah untuk kode yang mudah dibaca, difahami dan diubah oleh programmer. Alasan kenapa Clean Code itu penting karena agar memudahkan ketika melakukan kolaborasi bersama sesama developer, memudahkan juga ketika terdapat pengembangan fitur dan cepat ketika melakukan development. Cepat dilakukannya development ini karena code yang mudah dibaca membuat cepat memahami maksut kode dan tidak perlu butuh waktu yang banyak untuk memahami kode.

## Karakteristik Clean Code
<ol>
<li>Mudah dipahami</li>
 b :=12

penamaan nama variabel b menjadi ambigu karena tidak memiliki makna
<li>Mudah dieja dan dicari</li>
const yyyymmdstr = moment().format('YYYY/MM/DD')
	
penamaan nama variabel yyyymmdstr sulit dibaca dan dieja, lebih baik menggunakan nama variabel currentDate
<li>Singkat namun mendeskripsikan konteks</li>
<li>Konsisten</li>
Penggunaan nama	variabel ataupun nama method dan lain-lain harus selalu konsisten, seperti contoh menggunakan tipe camelcase maka seterusnya harus sama.
<li>Hindari penamaan konteks yang tidak perlu</li>
const Car = {
	
 carMake : 'honda',
	
 carModel : 'accord',
	
 carColor : 'blue',
	
}
	
Penambahan nama variabel diatas memiliki penamaan konteks yang perlu dihilangkan, berikut adalah contohnya.
const Car = {
	
 Make : 'honda',
	
 Model : 'accord',
	
 Color : 'blue',
	
}
	
<li>Komentar</li>
Penggunaan komentar pada suatu program tidak perlu setiap baris dikasih komentar, cukup pada beberapa blok kode yang memiliki cukup kesulitan pemahaman saja
<li>Good function</li>
Apabila terdapat function yang memiliki parameter dengan banyak value, lebih baik menggunakan array atau slice atau map sebagai parameter
<li>Gunakan konvensi</li>
Terdapat beberapa style guide yang dapat diikuti agar penulisan code sesuai dengan aturan. Beberapa style code yang cukup terkenal diantaranya ada airnb dan google.
<li>Formatting</li>
Terdapat beberapa saran formatting, diantaranya :
<ul>
	<li>Lebar baris code 80 - 120 karakter.</li>
	<li>Satu class 300 - 500 baris.</li>
	<li>Baris code yang berhubungan saling berdekatan.</li>
	<li>Dekatkan fungsi dengan pemanggilnya.</li>
	<li>Deklarasi variabel berdekatan dengan penggunanya.</li>
	<li>Perhatikan indentasi.</li>
	<li>Menggunakan prettier atau formatter.</li>
</ul>
</ol>

## Clean Code Principle

#### Keep It So Simple (KISS)
Hindari membuat fungsi yang dibuat untuk melakukan A, sekaligus memodifikasi B dan mengecek fungsi C, dst. 

#### Don't Repeat Yourself (DRY)
Duplikasi code terjadi karena sering copy paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.

#### Refactoring
Refactoring adalah proses restrukrisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor.
<ul>
<li>Membuat abstraksi</li>
<li>Memecah kode dengan fungsi/class</li>
<li>Perbaiki penamaan dan lokasi code</li>
<li>Deteksi kode yang memiliki duplikasi</li>
</ul>