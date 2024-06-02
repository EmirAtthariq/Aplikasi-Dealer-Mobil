package main

import "fmt"

type car struct {
	nama      string
	tahun     string
	penjualan int
}
type pabrik struct {
	nama         string
	jumlahProduk int
	penjualan    int
	produk       mobil
}

const NMAX int = 10

type pabrikan [NMAX]pabrik
type mobil [NMAX]car

var P pabrikan
var a int = 0 //a = jumlah pabrikan
var exit1 bool = false

func main() {
	for !exit1 {
		Menu()

	}
}
func Menu() {
	var x int // x = input user

	MainMenu()
	fmt.Scan(&x)
	if x == 1 {
		inputData(&a, &P)
	} else if x == 2 {
		ubahData()
	} else if x == 3 {
		cariData()
	} else if x == 4 {
		tampilkanData()
	} else if x == 5 {
		exit(&exit1)
	}
}

func exit(exit1 *bool) {
	var x string
	fmt.Println("---------------------")
	fmt.Print("Apakah anda yakin ingin keluar? Y/N: ")
	fmt.Scan(&x)
	if x == "Y" {
		fmt.Println("---------------------")
		fmt.Println("      Goodbye :)")
		fmt.Println("---------------------")
		*exit1 = true
	}
}

func MainMenu() {
	fmt.Println("---------------------")
	fmt.Println("   Dealer Mobil")
	fmt.Println("	1. Input Data")
	fmt.Println("	2. Ubah Data")
	fmt.Println("	3. Cari Data")
	fmt.Println("	4. Tampilkan Data")
	fmt.Println("	5. Keluar")
	fmt.Println("---------------------")
	fmt.Print("Pilih Menu: ")
}

func inputData(a *int, P *pabrikan) {
	var jumlahMobil, i, x int
	var Pabrikan, cek1 string
	var cek bool
	cek = false
	*a += 1
	x = *a - 1
	for x < NMAX && !cek && *a <= NMAX {
		fmt.Println("---------------------")
		fmt.Print("Masukkan nama pabrikan: ")
		fmt.Scan(&Pabrikan)
		P[x].nama = Pabrikan
		fmt.Println("---------------------")
		fmt.Print("Masukkan jumlah mobil buatan pabrikan tersebut (Max input 10 mobil) : ")
		fmt.Scan(&jumlahMobil)
		if jumlahMobil > NMAX {
			jumlahMobil = NMAX
		}
		P[x].jumlahProduk = jumlahMobil
		fmt.Println("---------------------")
		fmt.Println("Masukkan data mobil dengan urutan Nama, Tahun Keluar, Penjualan Unit : ")
		i = 0
		for i < jumlahMobil {
			fmt.Scan(&P[x].produk[i].nama, &P[x].produk[i].tahun, &P[x].produk[i].penjualan)
			i += 1
		}
		i = 0
		for i < jumlahMobil {
			P[x].penjualan += P[x].produk[i].penjualan
			i += 1
		}
		fmt.Println("---------------------")
		fmt.Print("Apakah anda ingin memasukkan pabrikan lainnya? Y/N: ")
		fmt.Scan(&cek1)
		if cek1 == "N" {
			cek = true
		} else if cek1 == "Y" {
			x += 1
			*a += 1
		}
	}
	if *a >= NMAX {
		*a = NMAX
		fmt.Println("---------------------")
		fmt.Println("Data sudah penuh")
		fmt.Println("Jumlah data pabrikan mobil :", *a)
	}

}

func ubahData() {
	var x int
	menuUbah()
	fmt.Scan(&x)
	if x == 1 {
		hapusData()
	} else if x == 2 {
		gantiData()
	} else if x == 3 {

	}
}

func menuUbah() {
	fmt.Println("---------------------")
	fmt.Println("	1. Hapus Data")
	fmt.Println("	2. Ganti Data")
	fmt.Println("	3. Kembali")
	fmt.Println("---------------------")
	fmt.Print("Pilih Menu: ")
}

func hapusData() {
	var x int
	menuHapus()
	fmt.Scan(&x)
	fmt.Println("---------------------")
	if x == 1 {
		hapusPabrikan(&a, &P)
		ubahData()

	} else if x == 2 {
		hapusMobil(&a, &P)
		ubahData()
	}
}

func hapusPabrikan(a *int, P *pabrikan) {
	var x string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Print("Masukkan Pabrikan yang akan dihapus: ")
	fmt.Scan(&x)
	fmt.Println()
	for i < *a && !cek {
		if P[i].nama == x {
			j = i
			for j < *a-1 {
				P[j] = P[j+1]
				j += 1
			}
			*a = *a - 1
			fmt.Println("---------------------")
			fmt.Println("Pabrikan sudah dihapus")
			cek = true
		}
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Pabrikan tidak ditemukan")
	}

}

func menuHapus() {
	fmt.Println("---------------------")
	fmt.Println("	1. Pabrikan")
	fmt.Println("	2. Mobil")
	fmt.Println("---------------------")
	fmt.Print("Pilih data yang akan dihapus: ")
}

func hapusMobil(a *int, P *pabrikan) {
	var x string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Print("Masukkan mobil yang akan dihapus: ")
	fmt.Scan(&x)
	fmt.Println()
	for i = 0; i < *a; i++ {
		for j < P[i].jumlahProduk && !cek {
			if P[i].produk[j].nama == x {
				for j < P[i].jumlahProduk-1 {
					P[i].produk[j] = P[i].produk[j+1]
					j += 1
				}
				P[i].jumlahProduk = P[i].jumlahProduk - 1
				fmt.Println("---------------------")
				fmt.Println("Mobil sudah dihapus")
				cek = true
			}
			j += 1
		}
		j = 0
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func gantiData() {
	var x int
	menuGanti()
	fmt.Scan(&x)
	if x == 1 {
		gantiPabrikan(&a, &P)
		ubahData()

	} else if x == 2 {
		gantiMobil()
		ubahData()
	}
}

func menuGanti() {
	fmt.Println("---------------------")
	fmt.Println("	1. Pabrikan")
	fmt.Println("	2. Mobil")
	fmt.Println("---------------------")
	fmt.Print("Pilih data yang akan diganti: ")
}

func gantiPabrikan(a *int, P *pabrikan) {
	var x, y string
	var i int
	var cek bool
	i = 0
	cek = false
	fmt.Println("---------------------")
	fmt.Print("Masukkan nama Pabrikan yang akan diganti: ")
	fmt.Scan(&x)
	fmt.Println("---------------------")
	fmt.Print("Masukkan nama Pabrikan yang baru: ")
	fmt.Scan(&y)
	fmt.Println()
	for i < *a && !cek {
		if P[i].nama == x {
			P[i].nama = y
			fmt.Println("---------------------")
			fmt.Println("Pabrikan sudah diganti")
			cek = true
		}
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Pabrikan tidak ditemukan")
	}
}

func gantiMobil() {
	var x int
	menuGantiMobil()
	fmt.Scan(&x)
	fmt.Println("---------------------")
	if x == 1 {
		gantiNamaMobil(&a, &P)
		ubahData()
	} else if x == 2 {
		gantiTahunMobil(&a, &P)
		ubahData()
	} else if x == 3 {
		gantiPenjualanMobil(&a, &P)
		ubahData()
	}
}
func menuGantiMobil() {
	fmt.Println("---------------------")
	fmt.Println("	1. Nama")
	fmt.Println("	2. Tahun Keluar")
	fmt.Println("	3. Penjualan")
	fmt.Println("---------------------")
	fmt.Print("Pilih data yang akan diganti: ")
}

func gantiNamaMobil(a *int, P *pabrikan) {
	var x, y string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Println("---------------------")
	fmt.Print("Masukkan nama mobil yang akan diganti: ")
	fmt.Scan(&x)
	fmt.Println("---------------------")
	fmt.Print("Masukkan nama mobil yang baru: ")
	fmt.Scan(&y)
	fmt.Println()
	for i < *a && !cek {
		for j < P[i].jumlahProduk && !cek {
			if P[i].produk[j].nama == x {
				P[i].produk[j].nama = y
				fmt.Println("---------------------")
				fmt.Println("Nama mobil sudah diganti")
				cek = true
			}
			j += 1
		}
		j = 0
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func gantiTahunMobil(a *int, P *pabrikan) {
	var x, y string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Print("Masukkan nama mobil yang akan diganti: ")
	fmt.Scan(&x)
	fmt.Println("---------------------")
	fmt.Print("Masukkan tahun keluaran mobil yang baru: ")
	fmt.Scan(&y)
	fmt.Println()
	for i < *a && !cek {
		for j < P[i].jumlahProduk && !cek {
			if P[i].produk[j].nama == x {
				P[i].produk[j].tahun = y
				fmt.Println("---------------------")
				fmt.Println("Tahun keluaran mobil sudah diganti")
				cek = true
			}
			j += 1
		}
		j = 0
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func gantiPenjualanMobil(a *int, P *pabrikan) {
	var x string
	var y, i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Println("---------------------")
	fmt.Print("Masukkan nama mobil yang akan diganti: ")
	fmt.Scan(&x)
	fmt.Println("---------------------")
	fmt.Print("Masukkan penjualan unit mobil yang baru: ")
	fmt.Scan(&y)
	fmt.Println()
	for i < *a && !cek {
		for j < P[i].jumlahProduk && !cek {
			if P[i].produk[j].nama == x {
				P[i].produk[j].penjualan = y
				fmt.Println("---------------------")
				fmt.Println("Penjualan unit mobil sudah diganti")
				cek = true
			}
			j += 1
		}
		j = 0
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func cariData() {
	var x int
	menuCariData()
	fmt.Scan(&x)
	fmt.Println("---------------------")
	if x == 1 {
		cariMobil(a, P)
	} else if x == 2 {
		cariPabrikan(a, P)
	}
}

func menuCariData() {
	fmt.Println("---------------------")
	fmt.Println("	1. Nama Mobil")
	fmt.Println("	2. Nama Pabrikan")
	fmt.Println("---------------------")
	fmt.Print("Pilih data yang akan dicari: ")
}

func cariMobil(a int, P pabrikan) {
	var x string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Print("Masukkan nama mobil yang akan dicari: ")
	fmt.Scan(&x)
	fmt.Println()
	for i < a && !cek {
		for j < P[i].jumlahProduk && !cek {
			if P[i].produk[j].nama == x {
				fmt.Println(P[i].produk[j].nama, P[i].produk[j].tahun, P[i].produk[j].penjualan)
				cek = true
			}
			j += 1
		}
		j = 0
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func cariPabrikan(a int, P pabrikan) {
	var x string
	var i, k int
	var cek bool
	i = 0
	k = 0
	cek = false
	fmt.Print("Masukkan nama pabrikan yang akan dicari: ")
	fmt.Scan(&x)
	fmt.Println()
	for i < a && !cek {
		if P[i].nama == x {
			fmt.Println(P[i].nama)
			cek = true
			for k = 0; k < P[i].jumlahProduk; k++ {
				fmt.Println(P[i].produk[k].nama, P[i].produk[k].tahun, P[i].produk[k].penjualan)
			}
		}
		i = i + 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Pabrikan tidak ditemukan")

	}
}

func tampilkanData() {
	var x int
	menuTampilkan()
	fmt.Scan(&x)
	if x == 1 {
		tampilPabrikan()
	} else if x == 2 {
		tampilMobil()
	}
}

func menuTampilkan() {
	fmt.Println("---------------------")
	fmt.Println("	1. Pabrikan")
	fmt.Println("	2. Mobil")
	fmt.Println("---------------------")
	fmt.Print("Pilih data yang akan ditampilkan: ")
}

func tampilPabrikan() {
	var x int
	menuTampilPabrikan()
	fmt.Scan(&x)
	fmt.Println("---------------------")
	if x == 1 {
		pabrikanJumlah(a, P)
	} else if x == 2 {
		pabrikanTertinggi(a, P)
	}
}

func menuTampilPabrikan() {
	fmt.Println("---------------------")
	fmt.Println("	1. Jumlah Mobil")
	fmt.Println("	2. Penjualan Tertinggi")
	fmt.Println("---------------------")
	fmt.Print("Pilih berdasarkan apa pabrikan yang akan ditampilkan: ")
}

func pabrikanJumlah(a int, P pabrikan) {
	var x int
	var i int
	var cek bool
	i = 0
	fmt.Print("Berapa jumlah mobil dari pabrikan yang dicari? : ")
	fmt.Scan(&x)
	fmt.Println("---------------------")
	fmt.Println("Pabrikan dengan jumlah mobil sebanyak", x, ":")
	for i < a {
		if P[i].jumlahProduk == x {
			fmt.Println(P[i].nama)
			cek = true
		}
		i = i + 1
	}
	if cek == false {

		fmt.Println("Tidak ada pabrikan dengan jumlah mobil sebanyak", x)

	}

}

func pabrikanTertinggi(a int, P pabrikan) {
	var pass, idx, i, x int
	type temp [NMAX]pabrik
	var t temp
	pass = 1

	for pass <= a-1 {
		idx = pass - 1
		i = pass
		for i < a {
			if P[idx].penjualan < P[i].penjualan {
				idx = i
			}
			i += 1
		}
		t[0] = P[pass-1]
		P[pass-1] = P[idx]
		P[idx] = t[0]
		pass += 1
	}
	fmt.Print("Berapa pabrikan yang ditampilkan dengan penjualan tertinggi? : ")
	fmt.Scan(&x)
	if x > a {
		x = a
	}
	fmt.Println("---------------------")
	fmt.Println(x, "Pabrikan dengan penjualan tertinggi: ")
	for y := 0; y < x; y++ {
		fmt.Println(P[y].nama, "dengan total penjualan sebanyak", P[y].penjualan, "unit")
	}

}

func tampilMobil() {
	var x int
	menuTampilMobil()
	fmt.Scan(&x)
	fmt.Println("---------------------")
	if x == 1 {
		mobilNama(a, P)
	} else if x == 2 {
		mobilTahun(a, P)
	} else if x == 3 {
		mobilPabrikan(a, P)
	} else if x == 4 {
		mobilPenjualan(a, P)
	}
}

func menuTampilMobil() {
	fmt.Println("---------------------")
	fmt.Println("	1. Nama")
	fmt.Println("	2. Tahun Keluar")
	fmt.Println("	3. Pabrikan")
	fmt.Println("	4. Penjualan Tertinggi")
	fmt.Println("---------------------")
	fmt.Print("Pilih berdasarkan apa mobil yang akan ditampilkan: ")
}

func mobilNama(a int, P pabrikan) {
	var x string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Print("Masukkan nama mobil yang akan ditampilkan: ")
	fmt.Scan(&x)
	fmt.Println()
	for i < a && !cek {
		for j < P[i].jumlahProduk && !cek {
			if P[i].produk[j].nama == x {
				fmt.Println(P[i].produk[j].nama, P[i].produk[j].tahun, P[i].produk[j].penjualan)
				cek = true
			}
			j += 1
		}
		j = 0
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func mobilTahun(a int, P pabrikan) {
	var x string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Print("Masukkan tahun keluar mobil yang akan ditampilkan: ")
	fmt.Scan(&x)
	fmt.Println("---------------------")
	fmt.Println("Mobil yang dikeluarkan di tahun", x)
	for i < a && !cek {
		for j < P[i].jumlahProduk {
			if P[i].produk[j].tahun == x {
				fmt.Println(P[i].produk[j].nama, P[i].produk[j].penjualan)
				cek = true
			}
			j += 1
		}
		j = 0
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func mobilPabrikan(a int, P pabrikan) {
	var x string
	var i, j int
	var cek bool
	i = 0
	j = 0
	cek = false
	fmt.Print("Masukkan pabrikan mobil yang akan ditampilkan: ")
	fmt.Scan(&x)
	fmt.Println("---------------------")
	fmt.Println("Mobil yang diproduksi oleh pabrikan", x)
	for i < a && !cek {
		if P[i].nama == x {
			for j < P[i].jumlahProduk {
				fmt.Println(P[i].produk[j].nama, P[i].produk[j].tahun, P[i].produk[j].penjualan)
				j += 1
				cek = true
			}
		}
		j = 0
		i += 1
	}
	if cek == false {
		fmt.Println("---------------------")
		fmt.Println("Mobil tidak ditemukan")
	}
}

func mobilPenjualan(a int, P pabrikan) {
	type tempo struct {
		nama      string
		penjualan int
	}
	type temp [1000]tempo
	type temp2 [1000]tempo
	var t temp
	var t2 temp2
	var pass, idx, i, x, b, c, d int
	pass = 1
	d = 0
	for b = 0; b < a; b++ {
		for c = 0; c < P[b].jumlahProduk; c++ {
			t[d].nama = P[b].produk[c].nama
			t[d].penjualan = P[b].produk[c].penjualan
			d += 1
		}
	}
	for pass <= d-1 {
		idx = pass - 1
		i = pass
		for i < d {
			if t[idx].penjualan < t[i].penjualan {
				idx = i
			}
			i += 1
		}
		t2[0] = t[pass-1]
		t[pass-1] = t[idx]
		t[idx] = t2[0]
		pass += 1
	}
	fmt.Print("Berapa mobil yang ditampilkan dengan penjualan tertinggi? : ")
	fmt.Scan(&x)
	if x > d {
		x = d
	}
	fmt.Println("---------------------")
	fmt.Println(x, "Mobil dengan penjualan tertinggi: ")
	for y := 0; y < x; y++ {
		fmt.Println(t[y].nama, "dengan total penjualan sebanyak", t[y].penjualan, "unit")
	}
}
