# Resume Materi Middleware

## Middleware
middleware adalah entitas yang menghubungkan ke pemrosesan permintaan/respons server. Jadi middleware bisa diibaratkan sebagai layer atau pelindung dari seseorang yang akan masuk ke sisi server website guna mengecek apakah user tersebut layak untuk mengakses data tersebut.

#### Implementasi Middleware
<ul>
<li>Logging Middleware</li>
<li>Session Middleware</li>
<li>Auth Middleware</li>
<li>Custom Middleware</li>
</ul>

## Setup Middleware Echo

#### Echo#Pre()
Dieksekusi sebelum request proses route
<ul>
<li>HTTPSRedirect</li>
<li>Rewrite</li>
<li>AddTraillingslash</li>
</ul>

#### Echo#Use()
Dieksekusi sesudah request proses route
<ul>
<li>Logger</li>
<li>BodyLimit</li>
<li>JWTAuth</li>
<li>BasicAuth</li>
<li>CORS</li>
</ul>

## Logger
Digunakan untuk mencatat setiap transaksi yang masuk atau request yang masuk

## Auth Middleware
Adalah salah satu autentikasi yang digunakan untuk mengamankan suatu data. Beberapa jenis autentikasi diantaranya ada basic authentication dan JWT.

#### Basic Authentication
Adalah teknik request http autentikasi, dimana method membutuhkan informasi username dan password untuk dimasukkan kedalam request header.

#### JWT middleware
WT adalah singkatan dari JSON Web Token yaitu sebuah JSON Object yang digunakan untuk aktifitas transfer informasi antar platform. JSON Web Token dapat berfungsi untuk sistem otentikasi dan juga untuk pertukaran informasi. Token ini terdiri dari header, payload dan signature.