package main 
inmport "fmt"

type lapangan struct {
	ID int 
	nama string 
	harga float64
	aktif bool 
}

type penyewa struct {
	id int
	nama string
	telepon string 

}

type jadwal struct {
	id int
	idLapangan int
	idPenyewa int
	tanggal string
	waktuMulai string
	waktuSelesai string
	total float64
}

var daftarLapangan []lapangan
var daftarPenyewa []penyewa
var daftarJadwal []jadwal

var idLapanganCounter int = 1
var idPenyewaCounter int = 1
var idJadwalCounter int = 1

func tambahlapangan()