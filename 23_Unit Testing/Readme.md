# Resume Materi Unit Testing

## Software Testing
Suatu proses menganalisis item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang dibutuhkan (yaitu, cacat) dan untuk mengevaluasi fitur item perangkat lunak.

#### Tujuan Dari Testing
<ul>
<li>Mencegah Regresi</li>
<li>Tingkat Keyakinan dalam Refactoring</li>
<li>Tingkatkan Code Design</li>
<li>Dokumentasi kode</li>
<li>Penskalaan Dokumentasi Tim</li>
</ul>

#### Level dari Testing
<ul>
<li>UI : End To End Testing</li>
<li>Integration</li>
<li>Unit Testing</li>
</ul>

## Framework
Berdasarkan bahasa pemrogramanmu, kamu harus memilih framework unit testing. Framework menyediakan tools, dan kebutuhan struktur untuk menjalankan testing secara efisiensi. Beberapa framework testing untuk go ada testify.

## Mocking
Sebuah unit testing harus tidak bergantung pada code lain seperti third party api. Kode diharuskan independen. Berikut beberapa yang tidak perlu di testing adalah :
<ul>
<li>third party module</li>
<li>Database</li>
<li>Third party api atau eksternal sistem</li>
<li>Obje class bahwa kamu melakukan test di tempat lain</li>
</ul>

Maka dari itu dibutuhkan mook objek untuk melakukan fake objek atau data tiruan untuk simulasi output.

## Coverage
Alat cakupan menganalisis kode aplikasi saat pengujian sedang berjalan. Beberapa coverage report yang dihasilkan bisa berupa format cli, xml, html dan lcov.