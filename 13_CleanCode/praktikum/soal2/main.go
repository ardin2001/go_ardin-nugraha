package main

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) berjalan() {
	m.tambahKecepatan(10)
}

func (m *Mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := Mobil{}
	mobilLamban.berjalan()
}
