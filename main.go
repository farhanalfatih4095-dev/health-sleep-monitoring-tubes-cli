// package main

// import "fmt"

// // ==================== KONSTANTA WARNA ====================

// const (
// 	Reset  = "\033[0m"
// 	Bold   = "\033[1m"
// 	Red    = "\033[31m"
// 	Green  = "\033[32m"
// 	Yellow = "\033[33m"
// 	Blue   = "\033[34m"
// 	Cyan   = "\033[36m"
// )

// // ==================== STRUCTS ====================

// type SleepRecord struct {
// 	Date     string
// 	BedTime  string
// 	WakeTime string
// 	Duration float64
// }

// type HealthRecord struct {
// 	Date         string
// 	WaterGlass   int
// 	WaterBottle  float64
// 	TotalWaterML float64
// 	ExerciseMins int
// 	ExerciseType string
// 	Mood         string
// 	MoodScore    int
// }

// // ==================== DATA GLOBAL ====================

// var sleepData [100]SleepRecord
// var healthData [100]HealthRecord
// var sleepCount int = 0
// var healthCount int = 0

// // tanggal sesi disimpan saat program pertama kali dijalankan
// var tanggalSesi string = ""

// // ==================== HELPER MATEMATIKA MANUAL ====================
// // Tidak menggunakan package math — semua dihitung sendiri.

// // roundFloat membulatkan float64 ke n desimal.
// func roundFloat(f float64, desimal int) float64 {
// 	mult := 1.0
// 	for i := 0; i < desimal; i++ {
// 		mult *= 10
// 	}
// 	if f >= 0 {
// 		return float64(int(f*mult+0.5)) / mult
// 	}
// 	return float64(int(f*mult-0.5)) / mult
// }

// // absFloat mengembalikan nilai absolut float64.
// func absFloat(f float64) float64 {
// 	if f < 0 {
// 		return -f
// 	}
// 	return f
// }

// // ==================== HELPER TAMPILAN ====================

// func clearScreen() {
// 	fmt.Print("\033[H\033[2J")
// }

// func cetakGaris(karakter string, panjang int) {
// 	for i := 0; i < panjang; i++ {
// 		fmt.Print(karakter)
// 	}
// 	fmt.Println()
// }

// func cetakHeader(judul string) {
// 	cetakGaris("═", 57)
// 	fmt.Printf("%s%s  %s%s\n", Bold, Cyan, judul, Reset)
// 	cetakGaris("═", 57)
// }

// // ==================== HELPER STRING ====================

// func ulangKarakter(karakter string, n int) string {
// 	hasil := ""
// 	for i := 0; i < n; i++ {
// 		hasil += karakter
// 	}
// 	return hasil
// }

// func padKanan(teks string, lebar int) string {
// 	panjang := len([]rune(teks))
// 	if panjang >= lebar {
// 		return teks
// 	}
// 	return teks + ulangKarakter(" ", lebar-panjang)
// }

// func intToStr(n int) string {
// 	if n == 0 {
// 		return "0"
// 	}
// 	negatif := false
// 	if n < 0 {
// 		negatif = true
// 		n = -n
// 	}
// 	hasil := ""
// 	for n > 0 {
// 		digit := n % 10
// 		hasil = string(rune('0'+digit)) + hasil
// 		n /= 10
// 	}
// 	if negatif {
// 		hasil = "-" + hasil
// 	}
// 	return hasil
// }

// // floatToStr mengubah float64 ke string dengan jumlah desimal tertentu.
// func floatToStr(f float64, desimal int) string {
// 	f = roundFloat(f, desimal)
// 	if desimal == 0 {
// 		return intToStr(int(f))
// 	}
// 	bulat := int(f)
// 	if f < 0 {
// 		bulat = int(f)
// 	}

// 	mult := 1.0
// 	for i := 0; i < desimal; i++ {
// 		mult *= 10
// 	}

// 	fracPart := int(roundFloat(absFloat(f-float64(bulat))*mult, 0))
// 	fracStr := intToStr(fracPart)
// 	for len(fracStr) < desimal {
// 		fracStr = "0" + fracStr
// 	}
// 	return intToStr(bulat) + "." + fracStr
// }

// // ==================== HELPER OUTPUT ====================

// func printError(msg string) {
// 	fmt.Println(Red + "  ⚠️  ERROR: " + msg + Reset)
// }

// func printSuccess(msg string) {
// 	fmt.Println(Green + "  ✅ " + msg + Reset)
// }

// // ==================== BAR VISUAL ====================

// func sleepBar(jam float64) string {
// 	filled := int(roundFloat(jam/9*10, 0))
// 	if filled > 10 {
// 		filled = 10
// 	}
// 	bar := "[" + ulangKarakter("█", filled) + ulangKarakter("░", 10-filled) + "]"
// 	if jam >= 7 {
// 		return Green + bar + Reset
// 	} else if jam >= 6 {
// 		return Yellow + bar + Reset
// 	}
// 	return Red + bar + Reset
// }

// func moodBar(skor float64) string {
// 	filled := int(roundFloat(skor/5*10, 0))
// 	if filled > 10 {
// 		filled = 10
// 	}
// 	bar := "[" + ulangKarakter("♥", filled) + ulangKarakter("·", 10-filled) + "]"
// 	if skor >= 4 {
// 		return Green + bar + Reset
// 	} else if skor >= 3 {
// 		return Yellow + bar + Reset
// 	}
// 	return Red + bar + Reset
// }

// // ==================== INPUT HELPER ====================

// func readString(prompt string) string {
// 	fmt.Print(prompt)
// 	var input string
// 	fmt.Scanln(&input)
// 	return input
// }

// func readInt(prompt string) (int, bool) {
// 	fmt.Print(prompt)
// 	var input string
// 	fmt.Scanln(&input)

// 	if len(input) == 0 {
// 		printError("Input tidak boleh kosong!")
// 		return 0, false
// 	}
// 	start := 0
// 	if input[0] == '-' {
// 		start = 1
// 	}
// 	for i := start; i < len(input); i++ {
// 		if input[i] < '0' || input[i] > '9' {
// 			printError("'" + input + "' bukan angka valid.")
// 			return 0, false
// 		}
// 	}
// 	hasil := 0
// 	negatif := false
// 	if input[0] == '-' {
// 		negatif = true
// 		input = input[1:]
// 	}
// 	for i := 0; i < len(input); i++ {
// 		hasil = hasil*10 + int(input[i]-'0')
// 	}
// 	if negatif {
// 		hasil = -hasil
// 	}
// 	return hasil, true
// }

// func readFloat(prompt string) (float64, bool) {
// 	fmt.Print(prompt)
// 	var input string
// 	fmt.Scanln(&input)

// 	if len(input) == 0 {
// 		printError("Input tidak boleh kosong!")
// 		return 0, false
// 	}
// 	titikCount := 0
// 	start := 0
// 	if input[0] == '-' {
// 		start = 1
// 	}
// 	for i := start; i < len(input); i++ {
// 		if input[i] == '.' {
// 			titikCount++
// 			if titikCount > 1 {
// 				printError("'" + input + "' bukan angka valid.")
// 				return 0, false
// 			}
// 		} else if input[i] < '0' || input[i] > '9' {
// 			printError("'" + input + "' bukan angka valid.")
// 			return 0, false
// 		}
// 	}
// 	var hasil float64
// 	var desimal float64
// 	var divider float64 = 1
// 	bagianDesimal := false
// 	negatif := false
// 	idx := 0
// 	if input[0] == '-' {
// 		negatif = true
// 		idx = 1
// 	}
// 	for i := idx; i < len(input); i++ {
// 		if input[i] == '.' {
// 			bagianDesimal = true
// 		} else {
// 			digit := float64(input[i] - '0')
// 			if bagianDesimal {
// 				divider *= 10
// 				desimal += digit / divider
// 			} else {
// 				hasil = hasil*10 + digit
// 			}
// 		}
// 	}
// 	hasil += desimal
// 	if negatif {
// 		hasil = -hasil
// 	}
// 	return hasil, true
// }

// // ==================== HELPER TANGGAL MANUAL ====================
// // Tanpa package time — semua operasi tanggal dihitung secara manual.

// // validasiTanggal memastikan format YYYY-MM-DD dan nilai yang masuk akal.
// func validasiTanggal(tanggal string) bool {
// 	if len(tanggal) != 10 {
// 		return false
// 	}
// 	if tanggal[4] != '-' || tanggal[7] != '-' {
// 		return false
// 	}
// 	for i, ch := range tanggal {
// 		if i == 4 || i == 7 {
// 			continue
// 		}
// 		if ch < '0' || ch > '9' {
// 			return false
// 		}
// 	}
// 	bulan := parseBagianTanggal(tanggal, 5, 2)
// 	hari := parseBagianTanggal(tanggal, 8, 2)
// 	return bulan >= 1 && bulan <= 12 && hari >= 1 && hari <= 31
// }

// // parseBagianTanggal mengambil angka dari string tanggal di posisi tertentu.
// func parseBagianTanggal(tanggal string, mulai int, panjang int) int {
// 	hasil := 0
// 	for i := 0; i < panjang; i++ {
// 		hasil = hasil*10 + int(tanggal[mulai+i]-'0')
// 	}
// 	return hasil
// }

// // tanggalKeHari mengkonversi YYYY-MM-DD ke jumlah hari sejak epoch (1 Jan tahun 1).
// // Digunakan untuk menghitung selisih antar tanggal.
// func tanggalKeHari(tanggal string) int {
// 	tahun := parseBagianTanggal(tanggal, 0, 4)
// 	bulan := parseBagianTanggal(tanggal, 5, 2)
// 	hari := parseBagianTanggal(tanggal, 8, 2)

// 	// Jumlah hari akumulasi per bulan (non-leap year)
// 	hariPerBulan := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// 	// Hitung total hari dari tahun
// 	totalHari := tahun * 365
// 	// Tambahkan hari dari tahun kabisat
// 	totalHari += (tahun - 1) / 4
// 	totalHari -= (tahun - 1) / 100
// 	totalHari += (tahun - 1) / 400

// 	// Tambahkan hari dari bulan-bulan sebelumnya di tahun ini
// 	for m := 1; m < bulan; m++ {
// 		totalHari += hariPerBulan[m]
// 		// Tambah 1 hari untuk Februari jika tahun kabisat
// 		if m == 2 && adaTahunKabisat(tahun) {
// 			totalHari++
// 		}
// 	}

// 	totalHari += hari
// 	return totalHari
// }

// // adaTahunKabisat mengecek apakah tahun tertentu adalah tahun kabisat.
// func adaTahunKabisat(tahun int) bool {
// 	return (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0
// }

// // selisihHariDari menghitung selisih hari antara tanggalSesi dan tanggal record.
// func selisihHariDari(tanggal string) int {
// 	if tanggalSesi == "" || !validasiTanggal(tanggal) {
// 		return 999
// 	}
// 	return tanggalKeHari(tanggalSesi) - tanggalKeHari(tanggal)
// }

// // ==================== HELPER DURASI TIDUR ====================

// // validasiJam memastikan format HH:MM dan nilainya masuk akal.
// func validasiJam(jam string) bool {
// 	if len(jam) != 5 || jam[2] != ':' {
// 		return false
// 	}
// 	for i, ch := range jam {
// 		if i == 2 {
// 			continue
// 		}
// 		if ch < '0' || ch > '9' {
// 			return false
// 		}
// 	}
// 	j := int(jam[0]-'0')*10 + int(jam[1]-'0')
// 	m := int(jam[3]-'0')*10 + int(jam[4]-'0')
// 	return j <= 23 && m <= 59
// }

// // hitungDurasiTidur menghitung selisih jam tidur dan bangun dalam jam.
// // Menangani kasus melewati tengah malam (misal tidur 22:00, bangun 06:00).
// func hitungDurasiTidur(bedTime string, wakeTime string) (float64, bool) {
// 	if !validasiJam(bedTime) {
// 		printError("Format jam tidur harus HH:MM (contoh: 22:30), jam 00-23, menit 00-59")
// 		return 0, false
// 	}
// 	if !validasiJam(wakeTime) {
// 		printError("Format jam bangun harus HH:MM (contoh: 06:00), jam 00-23, menit 00-59")
// 		return 0, false
// 	}

// 	menitTidur := int(bedTime[0]-'0')*10*60 + int(bedTime[1]-'0')*60 +
// 		int(bedTime[3]-'0')*10 + int(bedTime[4]-'0')
// 	menitBangun := int(wakeTime[0]-'0')*10*60 + int(wakeTime[1]-'0')*60 +
// 		int(wakeTime[3]-'0')*10 + int(wakeTime[4]-'0')

// 	selisihMenit := menitBangun - menitTidur
// 	if selisihMenit <= 0 {
// 		selisihMenit += 24 * 60 // melewati tengah malam
// 	}

// 	return roundFloat(float64(selisihMenit)/60.0, 2), true
// }

// // ==================== EVALUASI ====================

// func evalTidur(jam float64) {
// 	switch {
// 	case jam >= 7 && jam <= 9:
// 		fmt.Printf("  %s😴 Tidur BAIK (ideal 7-9 jam terpenuhi)%s\n", Green, Reset)
// 	case jam >= 6 && jam < 7:
// 		fmt.Printf("  %s😐 Tidur KURANG, coba 30 menit lebih awal%s\n", Yellow, Reset)
// 	case jam > 9:
// 		fmt.Printf("  %s😪 Tidur BERLEBIH, bisa menyebabkan lesu%s\n", Yellow, Reset)
// 	default:
// 		fmt.Printf("  %s😵 Tidur SANGAT KURANG (<6 jam), berbahaya!%s\n", Red, Reset)
// 	}
// }

// func evalKesehatan(air float64, olahraga int) {
// 	if air >= 2000 {
// 		fmt.Printf("  %-22s: %s✅ CUKUP (%.1f L)%s\n", "Konsumsi air", Green, air/1000, Reset)
// 	} else if air >= 1500 {
// 		fmt.Printf("  %-22s: %s⚠️  HAMPIR CUKUP (%.1f L)%s\n", "Konsumsi air", Yellow, air/1000, Reset)
// 	} else {
// 		fmt.Printf("  %-22s: %s❌ KURANG (%.1f L)%s\n", "Konsumsi air", Red, air/1000, Reset)
// 	}
// 	if olahraga >= 30 {
// 		fmt.Printf("  %-22s: %s✅ BAIK (%d menit)%s\n", "Olahraga", Green, olahraga, Reset)
// 	} else if olahraga > 0 {
// 		fmt.Printf("  %-22s: %s⚠️  KURANG (%d menit)%s\n", "Olahraga", Yellow, olahraga, Reset)
// 	} else {
// 		fmt.Printf("  %-22s: %s❌ Tidak olahraga%s\n", "Olahraga", Red, Reset)
// 	}
// }

// // ==================== MODUL TIDUR ====================

// func inputTidur() {
// 	clearScreen()
// 	cetakHeader("🌙  PENCATATAN TIDUR")
// 	fmt.Printf("%sTanggal sesi:%s %s\n\n", Bold, Reset, tanggalSesi)

// 	jamTidur := readString("  Jam tidur  (HH:MM, contoh 22:30) : ")
// 	jamBangun := readString("  Jam bangun (HH:MM, contoh 06:00) : ")

// 	durasi, ok := hitungDurasiTidur(jamTidur, jamBangun)
// 	if !ok {
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	record := SleepRecord{
// 		Date:     tanggalSesi,
// 		BedTime:  jamTidur,
// 		WakeTime: jamBangun,
// 		Duration: durasi,
// 	}

// 	diperbarui := false
// 	for i := 0; i < sleepCount; i++ {
// 		if sleepData[i].Date == tanggalSesi {
// 			sleepData[i] = record
// 			diperbarui = true
// 			break
// 		}
// 	}
// 	if !diperbarui && sleepCount < 100 {
// 		sleepData[sleepCount] = record
// 		sleepCount++
// 	}

// 	fmt.Println()
// 	printSuccess("Data tidur tersimpan! Durasi: " + floatToStr(durasi, 2) + " jam")
// 	fmt.Println()
// 	evalTidur(durasi)
// 	readString("\n  Tekan Enter untuk kembali...")
// }

// func hapusTidur() {
// 	clearScreen()
// 	cetakHeader("🗑️   HAPUS RIWAYAT TIDUR")

// 	if sleepCount == 0 {
// 		fmt.Println(Blue + "  ℹ  Belum ada data tidur yang tersimpan." + Reset)
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	fmt.Printf("\n  %-4s  %-12s  %-7s  %-7s  %s\n", "No.", "Tanggal", "Tidur", "Bangun", "Durasi")
// 	cetakGaris("─", 57)
// 	for i := 0; i < sleepCount; i++ {
// 		r := sleepData[i]
// 		fmt.Printf("  %-4s  %-12s  %-7s  %-7s  %s jam\n",
// 			intToStr(i+1), r.Date, r.BedTime, r.WakeTime, floatToStr(r.Duration, 2))
// 	}
// 	cetakGaris("─", 57)

// 	nomor, ok := readInt("\n  Masukkan nomor yang ingin dihapus (0 = batal) : ")
// 	if !ok {
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}
// 	if nomor == 0 {
// 		fmt.Println("  Operasi dibatalkan.")
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}
// 	if nomor < 1 || nomor > sleepCount {
// 		printError("Nomor tidak valid!")
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	idx := nomor - 1
// 	tanggalHapus := sleepData[idx].Date

// 	konfirmasi := readString("  Hapus data tanggal " + tanggalHapus + "? (y/n) : ")
// 	if konfirmasi != "y" && konfirmasi != "Y" {
// 		fmt.Println("  Operasi dibatalkan.")
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	// shift left: geser semua elemen setelah idx ke kiri satu posisi
// 	for i := idx; i < sleepCount-1; i++ {
// 		sleepData[i] = sleepData[i+1]
// 	}
// 	sleepData[sleepCount-1] = SleepRecord{}
// 	sleepCount--

// 	printSuccess("Data tidur tanggal " + tanggalHapus + " berhasil dihapus!")
// 	readString("\n  Tekan Enter untuk kembali...")
// }

// func menuTidur() {
// 	for {
// 		clearScreen()
// 		cetakGaris("═", 42)
// 		fmt.Printf("%s%s  🌙 MENU TIDUR%s\n", Bold, Cyan, Reset)
// 		cetakGaris("═", 42)
// 		fmt.Println()
// 		fmt.Printf("  %s╔══ PILIHAN ════════════════════╗%s\n", Blue, Reset)
// 		fmt.Printf("  %s║%s  1. ➕  Catat Tidur Hari Ini  %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s║%s  2. 🗑️   Hapus Riwayat Tidur  %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s║%s  0. 🔙  Kembali               %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s╚═══════════════════════════════╝%s\n", Blue, Reset)

// 		pilihan := readString("\n  Pilih (0-2): ")
// 		switch pilihan {
// 		case "1":
// 			inputTidur()
// 		case "2":
// 			hapusTidur()
// 		case "0":
// 			return
// 		default:
// 			printError("Pilihan tidak valid!")
// 			readString("  Tekan Enter untuk melanjutkan...")
// 		}
// 	}
// }

// // ==================== MODUL KESEHATAN ====================

// func inputKesehatan() {
// 	clearScreen()
// 	cetakHeader("💊  PENCATATAN KESEHATAN HARIAN")
// 	fmt.Printf("%sTanggal sesi:%s %s\n", Bold, Reset, tanggalSesi)

// 	fmt.Printf("\n%s  💧 KONSUMSI AIR PUTIH%s\n", Cyan, Reset)
// 	cetakGaris("─", 42)

// 	gelas, ok := readInt("  Jumlah gelas (ukuran 250ml)          : ")
// 	if !ok {
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}
// 	botol, ok2 := readFloat("  Jumlah botol (ukuran 600ml)          : ")
// 	if !ok2 {
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}
// 	totalML := float64(gelas)*250 + botol*600

// 	fmt.Printf("\n%s  🏃 OLAHRAGA%s\n", Cyan, Reset)
// 	cetakGaris("─", 42)

// 	menitOlahraga, ok3 := readInt("  Durasi olahraga (menit, 0 jika tidak) : ")
// 	if !ok3 {
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}
// 	jenisOlahraga := ""
// 	if menitOlahraga > 0 {
// 		jenisOlahraga = readString("  Jenis olahraga (contoh: Lari, Gym)    : ")
// 		if jenisOlahraga == "" {
// 			jenisOlahraga = "Tidak disebutkan"
// 		}
// 	}

// 	fmt.Printf("\n%s  😊 SUASANA HATI (MOOD)%s\n", Cyan, Reset)
// 	cetakGaris("─", 42)
// 	fmt.Println("  1. Sangat Bahagia 😄")
// 	fmt.Println("  2. Bahagia 😊")
// 	fmt.Println("  3. Biasa Saja 😐")
// 	fmt.Println("  4. Sedih 😢")
// 	fmt.Println("  5. Sangat Buruk 😭")

// 	pilihanMood, ok4 := readInt("\n  Pilih mood hari ini (1-5)             : ")
// 	if !ok4 {
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}
// 	if pilihanMood < 1 || pilihanMood > 5 {
// 		printError("Pilihan mood harus antara 1 dan 5!")
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	moodNama := [5]string{"Sangat Bahagia 😄", "Bahagia 😊", "Biasa Saja 😐", "Sedih 😢", "Sangat Buruk 😭"}
// 	moodSkor := [5]int{5, 4, 3, 2, 1}

// 	record := HealthRecord{
// 		Date:         tanggalSesi,
// 		WaterGlass:   gelas,
// 		WaterBottle:  botol,
// 		TotalWaterML: totalML,
// 		ExerciseMins: menitOlahraga,
// 		ExerciseType: jenisOlahraga,
// 		Mood:         moodNama[pilihanMood-1],
// 		MoodScore:    moodSkor[pilihanMood-1],
// 	}

// 	diperbarui := false
// 	for i := 0; i < healthCount; i++ {
// 		if healthData[i].Date == tanggalSesi {
// 			healthData[i] = record
// 			diperbarui = true
// 			break
// 		}
// 	}
// 	if !diperbarui && healthCount < 100 {
// 		healthData[healthCount] = record
// 		healthCount++
// 	}

// 	fmt.Println()
// 	cetakGaris("─", 42)
// 	printSuccess("Data kesehatan berhasil disimpan!")
// 	fmt.Println()
// 	fmt.Printf("  %-22s: %s ml (%.1f L)\n", "Total air", intToStr(int(totalML)), totalML/1000)
// 	if menitOlahraga > 0 {
// 		fmt.Printf("  %-22s: %d menit (%s)\n", "Olahraga", menitOlahraga, jenisOlahraga)
// 	} else {
// 		fmt.Printf("  %-22s: Tidak ada hari ini\n", "Olahraga")
// 	}
// 	fmt.Printf("  %-22s: %s\n", "Mood", moodNama[pilihanMood-1])
// 	fmt.Println()
// 	fmt.Printf("%s  📋 Evaluasi Singkat:%s\n", Bold, Reset)
// 	cetakGaris("─", 42)
// 	evalKesehatan(totalML, menitOlahraga)
// 	readString("\n  Tekan Enter untuk kembali...")
// }

// // ==================== SEQUENTIAL SEARCH ====================
// // Mencari record dari index 0 hingga sleepCount-1 satu per satu.
// // Time complexity: O(n) — cocok untuk data kecil atau data tidak terurut.
// // Return: index jika ditemukan, -1 jika tidak ada.

// func sequentialSearch(target string) int {
// 	for i := 0; i < sleepCount; i++ {
// 		if sleepData[i].Date == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// // ==================== BINARY SEARCH ====================
// // Mencari dengan membagi rentang pencarian di tengah setiap iterasi.
// // SYARAT: data HARUS sudah terurut ascending berdasarkan Date.
// // Time complexity: O(log n) — jauh lebih cepat untuk data besar yang terurut.
// // Return: index jika ditemukan, -1 jika tidak ada.

// func binarySearch(data [100]SleepRecord, count int, target string) int {
// 	low := 0
// 	high := count - 1
// 	for low <= high {
// 		mid := (low + high) / 2
// 		if data[mid].Date == target {
// 			return mid
// 		} else if data[mid].Date < target {
// 			low = mid + 1 // target di sisi kanan (tanggal lebih baru)
// 		} else {
// 			high = mid - 1 // target di sisi kiri (tanggal lebih lama)
// 		}
// 	}
// 	return -1
// }

// // copyAndSortByDate membuat salinan sleepData yang sudah diurutkan,
// // agar data asli tidak berubah saat binary search dilakukan.
// func copyAndSortByDate() ([100]SleepRecord, int) {
// 	var temp [100]SleepRecord
// 	for i := 0; i < sleepCount; i++ {
// 		temp[i] = sleepData[i]
// 	}
// 	// insertion sort ascending by Date
// 	for i := 1; i < sleepCount; i++ {
// 		key := temp[i]
// 		j := i - 1
// 		for j >= 0 && temp[j].Date > key.Date {
// 			temp[j+1] = temp[j]
// 			j--
// 		}
// 		temp[j+1] = key
// 	}
// 	return temp, sleepCount
// }

// func menuCariTidur() {
// 	clearScreen()
// 	cetakHeader("🔍  CARI RIWAYAT TIDUR")

// 	if sleepCount == 0 {
// 		fmt.Println(Blue + "  ℹ  Belum ada data tidur yang tersimpan." + Reset)
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	fmt.Printf("\n  %sMetode Pencarian:%s\n", Bold, Reset)
// 	fmt.Println("  1. Sequential Search  (urut satu per satu, O(n))")
// 	fmt.Println("  2. Binary Search      (bagi dua, O(log n), data diurutkan dulu)")
// 	fmt.Println("  0. Batal")

// 	pilihan := readString("\n  Pilih metode (0-2): ")
// 	if pilihan == "0" {
// 		return
// 	}
// 	if pilihan != "1" && pilihan != "2" {
// 		printError("Pilihan tidak valid!")
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	fmt.Printf("\n  Masukkan tanggal %s(format YYYY-MM-DD, contoh: 2025-06-15)%s\n", Yellow, Reset)
// 	target := readString("  Tanggal : ")

// 	if !validasiTanggal(target) {
// 		printError("Format tanggal harus YYYY-MM-DD (contoh: 2025-06-15)")
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	var idx int
// 	var metode string

// 	if pilihan == "1" {
// 		metode = "Sequential Search"
// 		idx = sequentialSearch(target)
// 	} else {
// 		metode = "Binary Search"
// 		sortedData, sortedCount := copyAndSortByDate()
// 		found := binarySearch(sortedData, sortedCount, target)
// 		if found == -1 {
// 			idx = -1
// 		} else {
// 			// ambil index di data asli agar tampilan konsisten
// 			idx = sequentialSearch(target)
// 		}
// 	}

// 	fmt.Println()
// 	cetakGaris("─", 57)
// 	fmt.Printf("  Metode  : %s%s%s\n", Cyan, metode, Reset)
// 	fmt.Printf("  Target  : %s\n", target)
// 	cetakGaris("─", 57)

// 	if idx == -1 {
// 		fmt.Printf(Red+"  ❌ Data tanggal %s tidak ditemukan.\n"+Reset, target)
// 	} else {
// 		r := sleepData[idx]
// 		fmt.Println(Green + "  ✅ Data ditemukan!" + Reset)
// 		fmt.Println()
// 		fmt.Printf("  %-20s: %s\n", "Tanggal", r.Date)
// 		fmt.Printf("  %-20s: %s\n", "Jam Tidur", r.BedTime)
// 		fmt.Printf("  %-20s: %s\n", "Jam Bangun", r.WakeTime)
// 		fmt.Printf("  %-20s: %s jam\n", "Durasi", floatToStr(r.Duration, 2))
// 		fmt.Printf("  %-20s:\n", "Status")
// 		evalTidur(r.Duration)
// 	}

// 	readString("\n  Tekan Enter untuk kembali...")
// }

// // ==================== SELECTION SORT ====================
// // Cara kerja: temukan elemen terkecil dari sisa array, tukar ke posisi sekarang.
// // Setiap iterasi outer loop, tepat 1 elemen sudah di posisi finalnya.
// // Time complexity: O(n²) — selalu, tidak bergantung kondisi data awal.

// func selectionSortByDuration() {
// 	for i := 0; i < sleepCount-1; i++ {
// 		minIdx := i
// 		for j := i + 1; j < sleepCount; j++ {
// 			if sleepData[j].Duration < sleepData[minIdx].Duration {
// 				minIdx = j
// 			}
// 		}
// 		if minIdx != i {
// 			sleepData[i], sleepData[minIdx] = sleepData[minIdx], sleepData[i]
// 		}
// 	}
// }

// func selectionSortByDate() {
// 	for i := 0; i < sleepCount-1; i++ {
// 		minIdx := i
// 		for j := i + 1; j < sleepCount; j++ {
// 			if sleepData[j].Date < sleepData[minIdx].Date {
// 				minIdx = j
// 			}
// 		}
// 		if minIdx != i {
// 			sleepData[i], sleepData[minIdx] = sleepData[minIdx], sleepData[i]
// 		}
// 	}
// }

// // ==================== INSERTION SORT ====================
// // Cara kerja: ambil satu elemen, sisipkan ke posisi tepat di bagian yang sudah terurut.
// // Seperti menyusun kartu remi — tangan kiri sudah rapi, ambil kartu baru, sisip di tempat yang benar.
// // Time complexity: O(n²) worst case, O(n) best case (data hampir terurut).

// func insertionSortByDuration() {
// 	for i := 1; i < sleepCount; i++ {
// 		key := sleepData[i]
// 		j := i - 1
// 		for j >= 0 && sleepData[j].Duration > key.Duration {
// 			sleepData[j+1] = sleepData[j]
// 			j--
// 		}
// 		sleepData[j+1] = key
// 	}
// }

// func insertionSortByDate() {
// 	for i := 1; i < sleepCount; i++ {
// 		key := sleepData[i]
// 		j := i - 1
// 		for j >= 0 && sleepData[j].Date > key.Date {
// 			sleepData[j+1] = sleepData[j]
// 			j--
// 		}
// 		sleepData[j+1] = key
// 	}
// }

// func menuUrutkanTidur() {
// 	clearScreen()
// 	cetakHeader("📋  URUTKAN RIWAYAT TIDUR")

// 	if sleepCount == 0 {
// 		fmt.Println(Blue + "  ℹ  Belum ada data tidur yang tersimpan." + Reset)
// 		readString("\n  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	fmt.Printf("\n  %sUrutkan berdasarkan:%s\n", Bold, Reset)
// 	fmt.Println("  1. Durasi tidur  (terpendek → terpanjang)")
// 	fmt.Println("  2. Tanggal       (terlama → terbaru)")

// 	kriteria := readString("\n  Pilih kriteria (1-2): ")
// 	if kriteria != "1" && kriteria != "2" {
// 		printError("Pilihan tidak valid!")
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	fmt.Printf("\n  %sAlgoritma pengurutan:%s\n", Bold, Reset)
// 	fmt.Println("  1. Selection Sort  (selalu O(n²), cocok data kecil)")
// 	fmt.Println("  2. Insertion Sort  (O(n) best case, efisien jika hampir terurut)")

// 	algo := readString("\n  Pilih algoritma (1-2): ")
// 	if algo != "1" && algo != "2" {
// 		printError("Pilihan tidak valid!")
// 		readString("  Tekan Enter untuk kembali...")
// 		return
// 	}

// 	var namaAlgo, namaKriteria string
// 	if algo == "1" {
// 		namaAlgo = "Selection Sort"
// 		if kriteria == "1" {
// 			namaKriteria = "Durasi (ascending)"
// 			selectionSortByDuration()
// 		} else {
// 			namaKriteria = "Tanggal (ascending)"
// 			selectionSortByDate()
// 		}
// 	} else {
// 		namaAlgo = "Insertion Sort"
// 		if kriteria == "1" {
// 			namaKriteria = "Durasi (ascending)"
// 			insertionSortByDuration()
// 		} else {
// 			namaKriteria = "Tanggal (ascending)"
// 			insertionSortByDate()
// 		}
// 	}

// 	fmt.Println()
// 	cetakGaris("─", 57)
// 	fmt.Printf("  Algoritma : %s%s%s\n", Cyan, namaAlgo, Reset)
// 	fmt.Printf("  Kriteria  : %s%s%s\n", Cyan, namaKriteria, Reset)
// 	cetakGaris("─", 57)
// 	fmt.Printf("  %-4s  %-12s  %-7s  %-7s  %-8s  Bar\n", "No.", "Tanggal", "Tidur", "Bangun", "Durasi")
// 	cetakGaris("─", 57)
// 	for i := 0; i < sleepCount; i++ {
// 		r := sleepData[i]
// 		fmt.Printf("  %-4s  %-12s  %-7s  %-7s  %s jam   %s\n",
// 			intToStr(i+1), r.Date, r.BedTime, r.WakeTime,
// 			padKanan(floatToStr(r.Duration, 1), 4),
// 			sleepBar(r.Duration))
// 	}
// 	cetakGaris("─", 57)
// 	printSuccess("Data berhasil diurutkan dengan " + namaAlgo + "!")
// 	fmt.Println(Yellow + "  ⚠️  Urutan data di memori telah berubah permanen." + Reset)
// 	readString("\n  Tekan Enter untuk kembali...")
// }

// // ==================== STATISTIK ====================

// func lihatStatistik() {
// 	clearScreen()
// 	cetakHeader("📊  STATISTIK LENGKAP - 7 HARI TERAKHIR")
// 	const L = 57

// 	fmt.Printf("\n  %s🌙 STATISTIK TIDUR%s\n", Cyan, Reset)
// 	cetakGaris("─", L)

// 	var totalJam float64
// 	jumlahTidur := 0
// 	for i := 0; i < sleepCount; i++ {
// 		if selisihHariDari(sleepData[i].Date) >= 0 && selisihHariDari(sleepData[i].Date) <= 7 {
// 			jumlahTidur++
// 			totalJam += sleepData[i].Duration
// 		}
// 	}

// 	if jumlahTidur == 0 {
// 		fmt.Println(Blue + "  ℹ  Belum ada data tidur dalam 7 hari ini." + Reset)
// 	} else {
// 		fmt.Printf("  %-12s  %-7s  %-7s  %-9s  Bar\n", "Tanggal", "Tidur", "Bangun", "Durasi")
// 		cetakGaris("─", L)
// 		for i := 0; i < sleepCount; i++ {
// 			r := sleepData[i]
// 			selisih := selisihHariDari(r.Date)
// 			if selisih >= 0 && selisih <= 7 {
// 				fmt.Printf("  %-12s  %-7s  %-7s  %s jam   %s\n",
// 					r.Date, r.BedTime, r.WakeTime,
// 					padKanan(floatToStr(r.Duration, 1), 4),
// 					sleepBar(r.Duration))
// 			}
// 		}
// 		cetakGaris("─", L)
// 		rataJam := totalJam / float64(jumlahTidur)
// 		fmt.Printf("  %-32s: %s%s jam/malam%s\n", "Rata-rata tidur", Bold, floatToStr(rataJam, 2), Reset)
// 		fmt.Printf("  %-32s: %d hari\n", "Total catatan", jumlahTidur)
// 		fmt.Printf("  %-32s: ", "Evaluasi")
// 		evalTidur(rataJam)
// 	}

// 	fmt.Printf("\n  %s💊 STATISTIK KESEHATAN%s\n", Cyan, Reset)
// 	cetakGaris("─", L)

// 	var totalAir float64
// 	var totalOlahraga int
// 	var totalMood int
// 	jumlahHealth := 0
// 	for i := 0; i < healthCount; i++ {
// 		selisih := selisihHariDari(healthData[i].Date)
// 		if selisih >= 0 && selisih <= 7 {
// 			jumlahHealth++
// 			totalAir += healthData[i].TotalWaterML
// 			totalOlahraga += healthData[i].ExerciseMins
// 			totalMood += healthData[i].MoodScore
// 		}
// 	}

// 	if jumlahHealth == 0 {
// 		fmt.Println(Blue + "  ℹ  Belum ada data kesehatan dalam 7 hari ini." + Reset)
// 	} else {
// 		fmt.Printf("  %-12s  %-10s  %-10s  %s\n", "Tanggal", "Air (ml)", "Olahraga", "Mood")
// 		cetakGaris("─", L)
// 		for i := 0; i < healthCount; i++ {
// 			r := healthData[i]
// 			selisih := selisihHariDari(r.Date)
// 			if selisih >= 0 && selisih <= 7 {
// 				ol := intToStr(r.ExerciseMins) + " mnt"
// 				if r.ExerciseMins == 0 {
// 					ol = "─"
// 				}
// 				fmt.Printf("  %-12s  %-10s  %-10s  %s\n",
// 					r.Date, intToStr(int(r.TotalWaterML)), ol, r.Mood)
// 			}
// 		}
// 		cetakGaris("─", L)
// 		rataAir := totalAir / float64(jumlahHealth)
// 		rataOlahraga := totalOlahraga / jumlahHealth
// 		rataMood := float64(totalMood) / float64(jumlahHealth)
// 		fmt.Printf("  %-32s: %s ml (%.1f L)\n", "Rata-rata air", intToStr(int(rataAir)), rataAir/1000)
// 		fmt.Printf("  %-32s: %d menit/hari\n", "Rata-rata olahraga", rataOlahraga)
// 		fmt.Printf("  %-32s: %s/5.0 %s\n", "Rata-rata mood", floatToStr(rataMood, 1), moodBar(rataMood))
// 		fmt.Printf("  %-32s: %d hari\n", "Total catatan", jumlahHealth)
// 		cetakGaris("─", L)
// 		fmt.Printf("%s  💡 Evaluasi Singkat:%s\n", Bold, Reset)
// 		evalKesehatan(rataAir, rataOlahraga)
// 		if rataMood >= 4 {
// 			fmt.Printf("  %-22s: %s✅ Mood POSITIF%s\n", "Suasana hati", Green, Reset)
// 		} else if rataMood >= 3 {
// 			fmt.Printf("  %-22s: %s😐 Mood BIASA%s\n", "Suasana hati", Yellow, Reset)
// 		} else {
// 			fmt.Printf("  %-22s: %s😢 Mood KURANG BAIK%s\n", "Suasana hati", Red, Reset)
// 		}
// 	}

// 	readString("\n  Tekan Enter untuk kembali...")
// }

// // ==================== INISIALISASI SESI ====================
// // Karena tidak menggunakan package time, tanggal diinput manual oleh user
// // satu kali di awal sesi. Semua pencatatan hari ini menggunakan tanggal ini.

// func inputTanggalSesi() {
// 	clearScreen()
// 	cetakGaris("═", 57)
// 	fmt.Printf("%s%s  🩺 KESEHATAN & POLA TIDUR  🌙%s\n", Bold, Cyan, Reset)
// 	cetakGaris("═", 57)
// 	fmt.Println()
// 	fmt.Printf("  %sSelamat datang!%s\n", Bold, Reset)
// 	fmt.Println("  Masukkan tanggal hari ini untuk memulai sesi.")
// 	fmt.Printf("  %s(format YYYY-MM-DD, contoh: 2025-06-28)%s\n\n", Yellow, Reset)

// 	for {
// 		tanggal := readString("  Tanggal hari ini : ")
// 		if validasiTanggal(tanggal) {
// 			tanggalSesi = tanggal
// 			printSuccess("Sesi dimulai untuk tanggal " + tanggalSesi)
// 			readString("\n  Tekan Enter untuk masuk ke menu...")
// 			return
// 		}
// 		printError("Format tidak valid. Gunakan YYYY-MM-DD (contoh: 2025-06-28)")
// 	}
// }

// // ==================== MENU UTAMA ====================

// func mainMenu() {
// 	for {
// 		clearScreen()
// 		cetakGaris("═", 44)
// 		fmt.Printf("%s%s  🩺 KESEHATAN & POLA TIDUR  🌙%s\n", Bold, Cyan, Reset)
// 		cetakGaris("═", 44)
// 		fmt.Printf("  %sTanggal sesi:%s %s\n\n", Bold, Reset, tanggalSesi)

// 		fmt.Printf("  %s╔══ PENCATATAN ══════════════════════╗%s\n", Blue, Reset)
// 		fmt.Printf("  %s║%s  1. 🌙  Tidur (Catat / Hapus)     %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s║%s  2. 💊  Catat Kesehatan Harian    %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s╠══ PENCARIAN & PENGURUTAN ══════════╣%s\n", Blue, Reset)
// 		fmt.Printf("  %s║%s  3. 🔍  Cari Riwayat Tidur        %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s║%s  4. 📋  Urutkan Riwayat Tidur     %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s╠══ LAPORAN ═════════════════════════╣%s\n", Blue, Reset)
// 		fmt.Printf("  %s║%s  5. 📊  Lihat Statistik           %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s╠════════════════════════════════════╣%s\n", Blue, Reset)
// 		fmt.Printf("  %s║%s  0. 🚪  Keluar                    %s║%s\n", Blue, Reset, Blue, Reset)
// 		fmt.Printf("  %s╚════════════════════════════════════╝%s\n", Blue, Reset)

// 		pilihan := readString("\n  Pilih menu (0-5): ")
// 		switch pilihan {
// 		case "1":
// 			menuTidur()
// 		case "2":
// 			inputKesehatan()
// 		case "3":
// 			menuCariTidur()
// 		case "4":
// 			menuUrutkanTidur()
// 		case "5":
// 			lihatStatistik()
// 		case "0":
// 			clearScreen()
// 			fmt.Printf("\n  %s👋 Terima kasih! Jaga kesehatan Anda selalu!%s\n\n", Green, Reset)
// 			return
// 		default:
// 			printError("Pilihan tidak valid! Masukkan angka 0-5.")
// 			readString("  Tekan Enter untuk melanjutkan...")
// 		}
// 	}
// }

// // ==================== MAIN ====================

// func main() {
// 	inputTanggalSesi()
// 	mainMenu()
// }