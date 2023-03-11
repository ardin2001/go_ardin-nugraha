```
package main
// Kendaraan : nama struct/objek sebaiknya menggunakan huruf besar di awal nama variabel
type kendaraan struct {
    // totalRoda : menggunakan huruf kapital pada kata kedua, agar lebih mudah dibaca
	totalroda       int
    // kecepatanPerJam : menggunakan huruf kapital pada kata kedua dan ketiga, agar lebih mudah dibaca
	kecepatanperjam int
}
// Mobil : nama struct/objek sebaiknya menggunakan huruf besar di awal nama variabel
type mobil struct {
    // Kendaraan : sesuai nama stuct diatas
	kendaraan
}
// (m *Mobil) : sesuai nama stuct diatas
func (m *mobil) berjalan() {
    // tambahKecepatan :  menggunakan huruf kapital pada kata kedua, agar lebih mudah dibaca
	m.tambahkecepatan(10)
}
// Mobil : sesuai nama stuct diatas
// tambahKecepatan, kecepatanBaru : menggunakan huruf kapital pada kata kedua, agar lebih mudah dibaca
func (m *mobil) tambahkecepatan(kecepatanbaru int) {
    // kecepatanPerJam, kecepatanBaru menggunakan : huruf kapital pada kata kedua dan ketiga, agar lebih mudah dibaca
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
    // mobilCepat : menggunakan huruf kapital pada kata kedua, agar lebih mudah dibaca
	mobilcepat := mobil{}
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()
    // mobilLamban : menggunakan huruf kapital pada kata kedua, agar lebih mudah dibaca
	mobillamban := mobil{}
	mobillamban.berjalan()
}
```