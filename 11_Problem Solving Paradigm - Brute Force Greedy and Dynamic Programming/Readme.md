# Resume Materi Problem Solving Paradigm - Brute Force Greedy and Dynamic Programming

Pada outline materi ini, membahas mengenai apa itu paradigma menyesaikan masalah, Brute Force, Divide & Conquer, Greedy dan Dynamic Programming.

## Apa itu Problem Solving Paradigm ?
Problem Solving Paradigm adalah pendekatan yang umum yang digunakan untuk menyelesaikan masalah seperti Complete Search, Divide and Conquer, pendekatan Greedy, Dynamic Programming. Setiap masalah perlu diselesaikan dengan pendekatan yang sesuai.

## Complete Search
<ul>
<li>Complete Search atau yang biasa dikenal dengan nama Brute Force adalah sebuah metode untuk menyelesaikan masalah  dengan melintasi seluruh ruang pencarian untuk mendapatkan solusi yang diperlukan.</li>
<li>Brute force biasa terjadi ketika tidak  ada algoritma lain yang tersedia.</li>
<li>Brute force adalaah algoritma tertua</li>
<li>Brute force mudah diimplementasikan dan mudah digunakan untuk pemula</li>
</ul>
Salah satu contoh penyelesaian masalah menggunakan brute force bisa menggunakan pencarian, find max atau min dari sebuah kumpulan nilai dan masih banyak lagi.

## Divide and Conquer
Divide and Conquer adalah paradigma pemecahan masalah dimana suatu masalah dibuat lebih sederhana, dengan membaginya menjadi bagian -bagian yang lebih kecil dan kemudian menggabungkan lagi bagian tersebut.
Berikut adalah langkah-langkahnya :
<ul>
<li>Divide : membagi masalah yang besar menjadi masalah yang lebih kecil, biasanya pembagiannya dibagi setengah sampai kebagian paling terkecil</li>
<li>Conquer : ketika masalah sudah cukup kecil untuk diselesaikan, langsung selesaikan</li>
<li>Combine : jika dibutuhkan maka perlu menggabungkan solusi dari masalah-masalah yang lebih kecil menjadi solusi untuk masalah yang besar.</li>
</ul>
Salah satu contoh implementasi dari Divide and Conquer adalah algoritma binary search, dimana tiap tiap element dipecah menjadi element paling kecil untuk menyelesaikan malsah. Menggunakan binary search lebih efisien dari pada menggunakan sequential search, karena binary membagi setiap iterasi dengan 2 pada masing-masing data.

## Greedy
Sebuah algoritma dikatakan greedy jika membuat pilihan optimal lokal pada setiap langkah dengan harapan pada akhirnya mencapai solusi optimal global. Salah satu contoh kasus sederhana yang menerapkan greedy adalah coin change. Algoritma Greedy terkenal lainnya ada Huffman Coding, Activity Selection, Dijkstra algorithm dan lain-lain.

## Dynamic Programming
Dynamic Programming addalah teknik algoritma untuk memecahkan masalah optimasi dengan memecahnya menjadi submasalah yang lebih sederhana dan memanfaatkan fakta bahwa solusi optimal untuk keseruhan masalah tergantung pada solusi optimal untuk submasalah.

#### Karakterisik dari Dynamic Programming
<ol>
<li>Overlapping Subproblems</li>
Submasalah adalah versi yang lebih kecil dari masalah aslinya. Masalah apa pun memiliki submasalah yang tumpang tindih jika menemukan solusinya melibatkan penyelesaian submasalah yang sama beberapa kali.
<li>Optimal Substructure Property</li>
Setiap masalah memiliki sifat substruktur optimal jika solusi optimal keseluruhannya dapat dibangun dari solusi optimal submasalahnya.
</ol>