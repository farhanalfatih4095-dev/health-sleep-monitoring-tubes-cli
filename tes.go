package main
import "fmt"
const MAKS = 100
type Tidur struct {
	Tanggal   string
	JamTidur  string
	JamBangun string
	DurasiJam float64
}
var data [MAKS]Tidur
var n int

// FUNCTION Helper (Logic & Data Calculation)
func baca(prompt string) string {
	fmt.Print(prompt)
	var s string
	fmt.Scanln(&s)
	return s
}
func bacaAngka(prompt string) int {
	fmt.Print(prompt)
	var x int
	fmt.Scan(&x)
	return x
}
func garis(panjang int) {
	for i := 0; i < panjang; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
func header(judul string) {
	garis(50)
	fmt.Printf("  %s\n", judul)
	garis(50)
}
func tampilTabel() {
	fmt.Printf("  %-4s %-12s %-7s %-7s %-8s %s\n", "No", "Tanggal", "Tidur", "Bangun", "Durasi", "Saran")
	garis(60)
	for i := 0; i < n; i++ {
		r := data[i]
		fmt.Printf("  %-4d %-12s %-7s %-7s %-8.2f %s\n",
			i+1, r.Tanggal, r.JamTidur, r.JamBangun, r.DurasiJam, saranTidur(r.DurasiJam))
	}
}

// FUNCTION UTILITY (Input/Output & UI)
func parseWaktu(s string) int {
	if len(s) != 5 || s[2] != ':' {
		return -1
	}
	jam := int(s[0]-'0')*10 + int(s[1]-'0')
	menit := int(s[3]-'0')*10 + int(s[4]-'0')
	if jam > 23 || menit > 59 {
		return -1
	}
	return jam*60 + menit
}
func hitungDurasi(tidur, bangun string) (float64, bool) {
	t := parseWaktu(tidur)
	b := parseWaktu(bangun)
	if t < 0 || b < 0 {
		fmt.Println("  Format waktu harus HH:MM (contoh: 06:30)")
		return 0, false
	}
	selisih := b - t
	if selisih <= 0 {
		selisih += 1440
	}
	return float64(selisih) / 60.0, true
}
func hariKe(tanggal string) int {
	if len(tanggal) != 10 {
		return -1
	}
	y := int(tanggal[0]-'0')*1000 + int(tanggal[1]-'0')*100 + int(tanggal[2]-'0')*10 + int(tanggal[3]-'0')
	m := int(tanggal[5]-'0')*10 + int(tanggal[6]-'0')
	d := int(tanggal[8]-'0')*10 + int(tanggal[9]-'0')
	if m <= 2 {
		y--
		m += 12
	}
	return 365*y + y/4 - y/100 + y/400 + (153*(m-3)+2)/5 + d
}
func saranTidur(jam float64) string {
	if jam >= 7 && jam <= 9 {
		return "BAIK (ideal 7-9 jam)"
	} else if jam >= 6 {
		return "KURANG, coba tidur lebih awal"
	} else if jam > 9 {
		return "BERLEBIH, bisa menyebabkan lesu"
	}
	return "SANGAT KURANG (<6 jam)"
}

// FUNCTION CRUD (CREATE, READ, UPDATE, DELETE)
func tambahData() {
	if n >= MAKS {
		fmt.Println("  Data penuh!")
		return
	}
	header("TAMBAH DATA TIDUR")
	tgl := baca("  Tanggal (YYYY-MM-DD) : ")
	tidur := baca("  Jam tidur  (HH:MM)   : ")
	bangun := baca("  Jam bangun (HH:MM)   : ")
	durasi, ok := hitungDurasi(tidur, bangun)
	if !ok {
		baca("\n  Tekan Enter...")
		return
	}
	data[n] = Tidur{tgl, tidur, bangun, durasi}
	n++
	fmt.Printf("\n  Tersimpan! Durasi: %.2f jam\n", durasi)
	fmt.Printf("  Saran: %s\n", saranTidur(durasi))
	baca("\n  Tekan Enter...")
}
func ubahData() {
	if n == 0 {
		fmt.Println("  Belum ada data.")
		baca("\n  Tekan Enter...")
		return
	}
	header("UBAH DATA TIDUR")
	tampilTabel()
	idx := bacaAngka("  Nomor data yang diubah: ") - 1
	if idx < 0 || idx >= n {
		fmt.Println("  Nomor tidak valid!")
		baca("\n  Tekan Enter...")
		return
	}
	tgl := baca("  Tanggal baru (YYYY-MM-DD) : ")
	tidur := baca("  Jam tidur baru  (HH:MM)   : ")
	bangun := baca("  Jam bangun baru (HH:MM)   : ")
	durasi, ok := hitungDurasi(tidur, bangun)
	if !ok {
		baca("\n  Tekan Enter...")
		return
	}
	data[idx] = Tidur{tgl, tidur, bangun, durasi}
	fmt.Println("  Data berhasil diubah!")
	baca("\n  Tekan Enter...")
}
func hapusData() {
	if n == 0 {
		fmt.Println("  Belum ada data.")
		baca("\n  Tekan Enter...")
		return
	}
	header("HAPUS DATA TIDUR")
	tampilTabel()
	idx := bacaAngka("  Nomor data yang dihapus: ") - 1
	if idx < 0 || idx >= n {
		fmt.Println("  Nomor tidak valid!")
		baca("\n  Tekan Enter...")
		return
	}
	for i := idx; i < n-1; i++ {
		data[i] = data[i+1]
	}
	n--
	fmt.Println("  Data berhasil dihapus!")
	baca("\n  Tekan Enter...")
}

// FUNCITION SEARCHING
func sequentialSearch(T [MAKS]Tidur, n int, tgl string) int {
	for i := 0; i < n; i++ {
		if T[i].Tanggal == tgl {
			return i
		}
	}
	return -1
}
func binarySearch(T [MAKS]Tidur, n int, tgl string) int {
	lo, hi := 0, n-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if T[mid].Tanggal == tgl {
			return mid
		} else if T[mid].Tanggal < tgl {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// FUNCITION SORTING
func selectionSortDurasi(T *[MAKS]Tidur, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j].DurasiJam < T[minIdx].DurasiJam {
				minIdx = j
			}
		}
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}
func insertionSortTanggal(T *[MAKS]Tidur, n int) {
	for i := 1; i < n; i++ {
		key := T[i]
		j := i - 1
		for j >= 0 && T[j].Tanggal > key.Tanggal {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = key
	}
}

// Function Menu Execution
func cariData() {
	if n == 0 {
		fmt.Println("  Belum ada data.")
		baca("\n  Tekan Enter...")
		return
	}
	header("CARI DATA BERDASARKAN TANGGAL")
	tgl := baca("  Masukkan tanggal (YYYY-MM-DD): ")
	idxSeq := sequentialSearch(data, n, tgl)
	fmt.Printf("\n  Sequential Search: ")
	if idxSeq >= 0 {
		fmt.Printf("Ditemukan di posisi %d\n", idxSeq+1)
	} else {
		fmt.Println("Tidak ditemukan")
	}
	var temp [MAKS]Tidur
	for i := 0; i < n; i++ {
		temp[i] = data[i]
	}
	insertionSortTanggal(&temp, n)
	idxBin := binarySearch(temp, n, tgl)
	fmt.Printf("  Binary Search    : ")
	if idxBin >= 0 {
		fmt.Printf("Ditemukan di posisi %d (data terurut)\n", idxBin+1)
	} else {
		fmt.Println("Tidak ditemukan")
	}
	if idxSeq >= 0 {
		r := data[idxSeq]
		fmt.Printf("\n  Detail: %s | Tidur %s -> Bangun %s | Durasi %.2f jam\n",
			r.Tanggal, r.JamTidur, r.JamBangun, r.DurasiJam)
		fmt.Printf("  Saran : %s\n", saranTidur(r.DurasiJam))
	}
	baca("\n  Tekan Enter...")
}
func sortData() {
	if n == 0 {
		fmt.Println("  Belum ada data.")
		baca("\n  Tekan Enter...")
		return
	}
	header("URUTKAN DATA TIDUR")
	fmt.Println("  1. Berdasarkan Durasi  (Selection Sort)")
	fmt.Println("  2. Berdasarkan Tanggal (Insertion Sort)")
	pilihan := baca("\n  Pilih (1/2): ")
	if pilihan == "1" {
		selectionSortDurasi(&data, n)
		fmt.Println("\n  Data diurutkan berdasarkan durasi!")
	} else if pilihan == "2" {
		insertionSortTanggal(&data, n)
		fmt.Println("\n  Data diurutkan berdasarkan tanggal!")
	} else {
		fmt.Println("  Pilihan tidak valid!")
		baca("\n  Tekan Enter...")
		return
	}
	fmt.Println()
	tampilTabel()
	baca("\n  Tekan Enter...")
}
func laporanMingguan() {
	header("LAPORAN 7 HARI TERAKHIR")
	if n == 0 {
		fmt.Println("  Belum ada data.")
		baca("\n  Tekan Enter...")
		return
	}
	tglRef := baca("  Masukkan tanggal hari ini (YYYY-MM-DD): ")
	hariRef := hariKe(tglRef)
	if hariRef < 0 {
		fmt.Println("  Format tanggal tidak valid!")
		baca("\n  Tekan Enter...")
		return
	}
	fmt.Printf("\n  %-4s %-12s %-7s %-7s %-8s %s\n", "No", "Tanggal", "Tidur", "Bangun", "Durasi", "Saran")
	garis(60)
	var totalJam float64
	jumlah := 0
	for i := 0; i < n; i++ {
		selisih := hariRef - hariKe(data[i].Tanggal)
		if selisih >= 0 && selisih <= 7 {
			jumlah++
			totalJam += data[i].DurasiJam
			r := data[i]
			fmt.Printf("  %-4d %-12s %-7s %-7s %-8.2f %s\n",
				jumlah, r.Tanggal, r.JamTidur, r.JamBangun, r.DurasiJam, saranTidur(r.DurasiJam))
		}
	}
	if jumlah == 0 {
		fmt.Println("\n  Tidak ada data dalam 7 hari terakhir.")
	} else {
		garis(60)
		rata := totalJam / float64(jumlah)
		fmt.Printf("\n  Rekapitulasi:\n")
		fmt.Printf("  Total catatan      : %d hari\n", jumlah)
		fmt.Printf("  Total durasi tidur : %.2f jam\n", totalJam)
		fmt.Printf("  Rata-rata per malam: %.2f jam\n", rata)
		fmt.Printf("  Evaluasi           : %s\n", saranTidur(rata))
	}
	baca("\n  Tekan Enter...")
}

func main() {
	for {
		fmt.Print("\033[H\033[2J")
		garis(40)
		fmt.Println("  PEMANTAU POLA TIDUR")
		garis(40)
		fmt.Println("  1. Tambah Data Tidur")
		fmt.Println("  2. Ubah Data Tidur")
		fmt.Println("  3. Hapus Data Tidur")
		fmt.Println("  4. Cari Data (Sequential & Binary)")
		fmt.Println("  5. Urutkan Data (Selection & Insertion)")
		fmt.Println("  6. Laporan 7 Hari Terakhir")
		fmt.Println("  0. Keluar")
		pilihan := baca("\n  Pilih menu: ")
		fmt.Print("\033[H\033[2J")
		switch pilihan {
		case "1":
			tambahData()
		case "2":
			ubahData()
		case "3":
			hapusData()
		case "4":
			cariData()
		case "5":
			sortData()
		case "6":
			laporanMingguan()
		case "0":
			fmt.Println("\n  Terima kasih! Jaga pola tidur Anda!")
			return
		default:
			fmt.Println("  Pilihan tidak valid!")
			baca("  Tekan Enter...")
		}
	}
}
