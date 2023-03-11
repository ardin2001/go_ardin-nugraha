<ol>
<li>Berapa banyak kekurangan dalam penulisan kode tersebut?</li> 12 kekurangan

<li>Bagian mana saja terjadi kekurangan tersebut?</li>
<ul>
<li>baris3 : type user struct {</li>
<li>baris5 : username int</li>
<li>baris6 : password int</li>
<li>baris9 : type userservice struct {</li>
<li>baris10 : t []user</li>
<li>baris13 : func (u userservice) getallusers() []user {</li>
<li>baris14 : return u.t</li>
<li>baris17 : func (u userservice) getuserbyid(id int) user {</li>
<li>baris18 : for _, r := range u.t {</li>
<li>baris19 : if id == r.id {</li>
<li>baris 20 : return r</li>
</ul>

<li>Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.</li>

```
package main
// User : untuk pembuatan struct/objek lebih baik menggunakan huruf besar pada karakter pertama
type user struct {
	id       int
  // username string : untuk tipe data username lebih baik menggunakan tipe data string, karena username tidak mungkin terbuat dari angka saja
	username int
  // password string : untuk tipe data username lebih baik menggunakan tipe data string, karena password tidak mungkin terbuat dari angka saja. Dengan menggunakan kombinasi tidak hanya berupa int saja makan akan susah di tebak
	password int
}
// UserService : karena nama struct terdiri dari 2 kata, akan lebih mudah dibaca apabila nama variabel struct adalah UserService
type userservice struct {
// users []user : penamaan variabel t yang kurang jelas, sehingga lebih baik menggunakan nama variabel yang lebih bermakna seperti users
	t []user
}
// getAllUsers : nama method menggunakan tipe lowercase sehingga sulit dibaca, akan lebih baik jika diganti dengan nama method getAllUsers
func (u userservice) getallusers() []user {
// return value sulit dimengerti karena nama variabel t yang kurang jelas, sehingga perlu diubah sesuai solusi baris10
	return u.t
}
// getUserById : nama method menggunakan tipe lowercase sehingga sulit dibaca, akan lebih baik jika diganti dengan nama method getAllUsers
func (u userservice) getuserbyid(id int) user {
// penamaan variabel yang kurang baik (r), nama variabel harus sesuai dan bermakna
// kemudian harus dijelaskan berupa komentar apabila ada suatu operasi dalam satu blok
	for _, r := range u.t {
  // penamaan variabel r yang kurang jelas, membuat pembacaan perbandingan menjadi ambigu
		if id == r.id {
    // hasil return tidak begitu jelas karena nama variabel r tidak jelas
			return r
		}
	}

	return user{}
}
```
</ol>
