package main

import (
	"fmt"
	"math"
	"time"
)

// ==================== KONSTANTA WARNA ====================

const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

// ==================== STRUCTS ====================

type SleepRecord struct {
	Date     string
	BedTime  string
	WakeTime string
	Duration float64
}

type HealthRecord struct {
	Date         string
	WaterGlass   int
	WaterBottle  float64
	TotalWaterML float64
	ExerciseMins int
	ExerciseType string
	Mood         string
	MoodScore    int
}

// ==================== DATA (disimpan di memory) ====================

var sleepData [100]SleepRecord
var healthData [100]HealthRecord
var sleepCount int = 0
var healthCount int = 0

// ==================== HELPER: INPUT ====================

func readString(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

func readInt(prompt string) (int, bool) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)

	// cek manual apakah semua karakter adalah digit
	if len(input) == 0 {
		fmt.Println(Red + "  ⚠️  ERROR: Input tidak boleh kosong! Harap masukkan angka." + Reset)
		return 0, false
	}

	start := 0
	if input[0] == '-' {
		start = 1
	}

	for i := start; i < len(input); i++ {
		if input[i] < '0' || input[i] > '9' {
			fmt.Printf(Red+"  ⚠️  ERROR: Input '%s' bukan angka! Harap masukkan angka yang valid.\n"+Reset, input)
			return 0, false
		}
	}

	// konversi manual string ke int
	hasil := 0
	negatif := false
	if input[0] == '-' {
		negatif = true
		input = input[1:]
	}
	for i := 0; i < len(input); i++ {
		hasil = hasil*10 + int(input[i]-'0')
	}
	if negatif {
		hasil = -hasil
	}
	return hasil, true
}

func readFloat(prompt string) (float64, bool) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)

	if len(input) == 0 {
		fmt.Println(Red + "  ⚠️  ERROR: Input tidak boleh kosong! Harap masukkan angka." + Reset)
		return 0, false
	}

	titikCount := 0
	start := 0
	if input[0] == '-' {
		start = 1
	}
	for i := start; i < len(input); i++ {
		if input[i] == '.' {
			titikCount++
			if titikCount > 1 {
				fmt.Printf(Red+"  ⚠️  ERROR: Input '%s' bukan angka! Harap masukkan angka yang valid.\n"+Reset, input)
				return 0, false
			}
		} else if input[i] < '0' || input[i] > '9' {
			fmt.Printf(Red+"  ⚠️  ERROR: Input '%s' bukan angka! Harap masukkan angka yang valid.\n"+Reset, input)
			return 0, false
		}
	}

	// konversi manual string ke float64
	var hasil float64 = 0
	var desimal float64 = 0
	var divider float64 = 1
	bagianDesimal := false
	negatif := false
	idx := 0
	if input[0] == '-' {
		negatif = true
		idx = 1
	}
	for i := idx; i < len(input); i++ {
		if input[i] == '.' {
			bagianDesimal = true
		} else {
			digit := float64(input[i] - '0')
			if bagianDesimal {
				divider *= 10
				desimal += digit / divider
			} else {
				hasil = hasil*10 + digit
			}
		}
	}
	hasil += desimal
	if negatif {
		hasil = -hasil
	}
	return hasil, true
}

// ==================== HELPER: TAMPILAN ====================

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func cetakGaris(karakter string, panjang int) {
	for i := 0; i < panjang; i++ {
		fmt.Print(karakter)
	}
	fmt.Println()
}

func cetakHeader(judul string) {
	cetakGaris("═", 57)
	fmt.Printf("%s%s  %s%s\n", Bold, Cyan, judul, Reset)
	cetakGaris("═", 57)
}

func ulangKarakter(karakter string, n int) string {
	hasil := ""
	for i := 0; i < n; i++ {
		hasil += karakter
	}
	return hasil
}

func padKanan(teks string, lebar int) string {
	panjang := len([]rune(teks))
	if panjang >= lebar {
		return teks
	}
	return teks + ulangKarakter(" ", lebar-panjang)
}

func intToStr(n int) string {
	if n == 0 {
		return "0"
	}
	negatif := false
	if n < 0 {
		negatif = true
		n = -n
	}
	hasil := ""
	for n > 0 {
		digit := n % 10
		hasil = string(rune('0'+digit)) + hasil
		n /= 10
	}
	if negatif {
		hasil = "-" + hasil
	}
	return hasil
}

func floatToStr(f float64, desimal int) string {
	if desimal == 0 {
		return intToStr(int(math.Round(f)))
	}
	bulat := int(f)
	mult := math.Pow(10, float64(desimal))
	fracPart := int(math.Round((f - float64(bulat)) * mult))
	if fracPart < 0 {
		fracPart = -fracPart
	}
	fracStr := intToStr(fracPart)
	// padding nol di depan jika perlu
	for len(fracStr) < desimal {
		fracStr = "0" + fracStr
	}
	return intToStr(bulat) + "." + fracStr
}

// ==================== HELPER: TANGGAL ====================

func hariIni() string {
	t := time.Now()
	tahun := intToStr(t.Year())
	bulan := intToStr(int(t.Month()))
	if len(bulan) < 2 {
		bulan = "0" + bulan
	}
	hari := intToStr(t.Day())
	if len(hari) < 2 {
		hari = "0" + hari
	}
	return tahun + "-" + bulan + "-" + hari
}

func selisihHari(tanggal string) float64 {
	// parse tanggal format YYYY-MM-DD secara manual
	if len(tanggal) != 10 {
		return 999
	}
	tahun := int(tanggal[0]-'0')*1000 + int(tanggal[1]-'0')*100 +
		int(tanggal[2]-'0')*10 + int(tanggal[3]-'0')
	bulan := int(tanggal[5]-'0')*10 + int(tanggal[6]-'0')
	hari := int(tanggal[8]-'0')*10 + int(tanggal[9]-'0')

	loc := time.Now().Location()
	tgl := time.Date(tahun, time.Month(bulan), hari, 0, 0, 0, 0, loc)
	selisih := time.Since(tgl).Hours() / 24
	return selisih
}

// ==================== HELPER: DURASI TIDUR ====================

func hitungDurasiTidur(bedTime string, wakeTime string) (float64, bool) {
	if len(bedTime) != 5 || len(wakeTime) != 5 {
		fmt.Println(Red + "  ⚠️  ERROR: Format waktu harus HH:MM (contoh: 22:30)" + Reset)
		return 0, false
	}
	if bedTime[2] != ':' || wakeTime[2] != ':' {
		fmt.Println(Red + "  ⚠️  ERROR: Format waktu harus HH:MM (contoh: 22:30)" + Reset)
		return 0, false
	}

	jamTidur := int(bedTime[0]-'0')*10 + int(bedTime[1]-'0')
	menitTidur := int(bedTime[3]-'0')*10 + int(bedTime[4]-'0')
	jamBangun := int(wakeTime[0]-'0')*10 + int(wakeTime[1]-'0')
	menitBangun := int(wakeTime[3]-'0')*10 + int(wakeTime[4]-'0')

	if jamTidur > 23 || menitTidur > 59 || jamBangun > 23 || menitBangun > 59 {
		fmt.Println(Red + "  ⚠️  ERROR: Jam harus 00-23, menit harus 00-59" + Reset)
		return 0, false
	}

	totalMenitTidur := jamTidur*60 + menitTidur
	totalMenitBangun := jamBangun*60 + menitBangun

	selisihMenit := totalMenitBangun - totalMenitTidur
	if selisihMenit < 0 {
		selisihMenit += 24 * 60 // melewati tengah malam
	}

	durasi := float64(selisihMenit) / 60.0
	return math.Round(durasi*100) / 100, true
}

// ==================== BAR VISUAL ====================

func sleepBar(jam float64) string {
	filled := int(math.Round(jam / 9 * 10))
	if filled > 10 {
		filled = 10
	}
	bar := "[" + ulangKarakter("█", filled) + ulangKarakter("░", 10-filled) + "]"
	if jam >= 7 {
		return Green + bar + Reset
	} else if jam >= 6 {
		return Yellow + bar + Reset
	}
	return Red + bar + Reset
}

func moodBar(skor float64) string {
	filled := int(math.Round(skor / 5 * 10))
	if filled > 10 {
		filled = 10
	}
	bar := "[" + ulangKarakter("♥", filled) + ulangKarakter("·", 10-filled) + "]"
	if skor >= 4 {
		return Green + bar + Reset
	} else if skor >= 3 {
		return Yellow + bar + Reset
	}
	return Red + bar + Reset
}

// ==================== EVALUASI ====================

func evalTidur(jam float64) {
	switch {
	case jam >= 7 && jam <= 9:
		fmt.Printf("  %s😴 Tidur BAIK (ideal 7-9 jam terpenuhi)%s\n", Green, Reset)
	case jam >= 6 && jam < 7:
		fmt.Printf("  %s😐 Tidur KURANG, coba 30 menit lebih awal%s\n", Yellow, Reset)
	case jam > 9:
		fmt.Printf("  %s😪 Tidur BERLEBIH, bisa menyebabkan lesu%s\n", Yellow, Reset)
	default:
		fmt.Printf("  %s😵 Tidur SANGAT KURANG (<6 jam), berbahaya!%s\n", Red, Reset)
	}
}

func evalKesehatan(air float64, olahraga int) {
	if air >= 2000 {
		fmt.Printf("  %-22s: %s✅ CUKUP (%.1f L)%s\n", "Konsumsi air", Green, air/1000, Reset)
	} else if air >= 1500 {
		fmt.Printf("  %-22s: %s⚠️  HAMPIR CUKUP (%.1f L)%s\n", "Konsumsi air", Yellow, air/1000, Reset)
	} else {
		fmt.Printf("  %-22s: %s❌ KURANG (%.1f L)%s\n", "Konsumsi air", Red, air/1000, Reset)
	}

	if olahraga >= 30 {
		fmt.Printf("  %-22s: %s✅ BAIK (%d menit)%s\n", "Olahraga", Green, olahraga, Reset)
	} else if olahraga > 0 {
		fmt.Printf("  %-22s: %s⚠️  KURANG (%d menit)%s\n", "Olahraga", Yellow, olahraga, Reset)
	} else {
		fmt.Printf("  %-22s: %s❌ Tidak olahraga%s\n", "Olahraga", Red, Reset)
	}
}

// ==================== MODUL TIDUR ====================

func inputTidur() {
	clearScreen()
	cetakHeader("🌙  PENCATATAN TIDUR")
	fmt.Printf("%sTanggal:%s %s\n\n", Bold, Reset, hariIni())

	jamTidur := readString("  Jam tidur  (HH:MM, contoh 22:30) : ")
	jamBangun := readString("  Jam bangun (HH:MM, contoh 06:00) : ")

	durasi, ok := hitungDurasiTidur(jamTidur, jamBangun)
	if !ok {
		readString("\n  Tekan Enter untuk kembali...")
		return
	}

	tanggal := hariIni()
	record := SleepRecord{Date: tanggal, BedTime: jamTidur, WakeTime: jamBangun, Duration: durasi}

	// cek apakah sudah ada data hari ini → timpa
	diperbarui := false
	for i := 0; i < sleepCount; i++ {
		if sleepData[i].Date == tanggal {
			sleepData[i] = record
			diperbarui = true
			break
		}
	}
	if !diperbarui && sleepCount < 100 {
		sleepData[sleepCount] = record
		sleepCount++
	}

	fmt.Println()
	fmt.Printf(Green+"  ✅ Data tidur tersimpan! Durasi: %s jam\n"+Reset, floatToStr(durasi, 2))
	fmt.Println()
	evalTidur(durasi)
	readString("\n  Tekan Enter untuk kembali...")
}

// ==================== MODUL KESEHATAN ====================

func inputKesehatan() {
	clearScreen()
	cetakHeader("💊  PENCATATAN KESEHATAN HARIAN")
	tanggal := hariIni()
	fmt.Printf("%sTanggal:%s %s\n", Bold, Reset, tanggal)

	// --- Air ---
	fmt.Printf("\n%s  💧 KONSUMSI AIR PUTIH%s\n", Cyan, Reset)
	cetakGaris("─", 42)

	gelas, ok := readInt("  Jumlah gelas (ukuran 250ml)         : ")
	if !ok {
		readString("\n  Tekan Enter untuk kembali...")
		return
	}
	botol, ok2 := readFloat("  Jumlah botol (ukuran 600ml)         : ")
	if !ok2 {
		readString("\n  Tekan Enter untuk kembali...")
		return
	}
	totalML := float64(gelas)*250 + botol*600

	// --- Olahraga ---
	fmt.Printf("\n%s  🏃 OLAHRAGA%s\n", Cyan, Reset)
	cetakGaris("─", 42)

	menitOlahraga, ok3 := readInt("  Durasi olahraga (menit, 0 jika tidak): ")
	if !ok3 {
		readString("\n  Tekan Enter untuk kembali...")
		return
	}
	jenisOlahraga := ""
	if menitOlahraga > 0 {
		jenisOlahraga = readString("  Jenis olahraga (contoh: Lari, Gym)   : ")
		if jenisOlahraga == "" {
			jenisOlahraga = "Tidak disebutkan"
		}
	}

	// --- Mood ---
	fmt.Printf("\n%s  😊 SUASANA HATI (MOOD)%s\n", Cyan, Reset)
	cetakGaris("─", 42)
	fmt.Println("  1. Sangat Bahagia 😄")
	fmt.Println("  2. Bahagia 😊")
	fmt.Println("  3. Biasa Saja 😐")
	fmt.Println("  4. Sedih 😢")
	fmt.Println("  5. Sangat Buruk 😭")

	pilihanMood, ok4 := readInt("\n  Pilih mood hari ini (1-5)            : ")
	if !ok4 {
		readString("\n  Tekan Enter untuk kembali...")
		return
	}
	if pilihanMood < 1 || pilihanMood > 5 {
		fmt.Println(Red + "  ⚠️  ERROR: Pilihan mood harus antara 1 dan 5!" + Reset)
		readString("\n  Tekan Enter untuk kembali...")
		return
	}

	moodNama := [5]string{
		"Sangat Bahagia 😄",
		"Bahagia 😊",
		"Biasa Saja 😐",
		"Sedih 😢",
		"Sangat Buruk 😭",
	}
	moodSkor := [5]int{5, 4, 3, 2, 1}

	record := HealthRecord{
		Date: tanggal, WaterGlass: gelas, WaterBottle: botol,
		TotalWaterML: totalML, ExerciseMins: menitOlahraga,
		ExerciseType: jenisOlahraga,
		Mood:         moodNama[pilihanMood-1],
		MoodScore:    moodSkor[pilihanMood-1],
	}

	diperbarui := false
	for i := 0; i < healthCount; i++ {
		if healthData[i].Date == tanggal {
			healthData[i] = record
			diperbarui = true
			break
		}
	}
	if !diperbarui && healthCount < 100 {
		healthData[healthCount] = record
		healthCount++
	}

	fmt.Println()
	cetakGaris("─", 42)
	fmt.Println(Green + "  ✅ Data kesehatan berhasil disimpan!" + Reset)
	fmt.Println()
	fmt.Printf("  %-22s: %s ml (%.1f L)\n", "Total air", intToStr(int(totalML)), totalML/1000)
	if menitOlahraga > 0 {
		fmt.Printf("  %-22s: %d menit (%s)\n", "Olahraga", menitOlahraga, jenisOlahraga)
	} else {
		fmt.Printf("  %-22s: Tidak ada hari ini\n", "Olahraga")
	}
	fmt.Printf("  %-22s: %s\n", "Mood", moodNama[pilihanMood-1])
	fmt.Println()
	fmt.Printf("%s  📋 Evaluasi Singkat:%s\n", Bold, Reset)
	cetakGaris("─", 42)
	evalKesehatan(totalML, menitOlahraga)

	readString("\n  Tekan Enter untuk kembali...")
}

// ==================== STATISTIK GABUNGAN ====================

func lihatStatistik() {
	clearScreen()
	cetakHeader("📊  STATISTIK LENGKAP - 7 HARI TERAKHIR")

	const L = 57 // lebar garis, konsisten dengan header

	// ════ BAGIAN TIDUR ════
	fmt.Printf("\n  %s🌙 STATISTIK TIDUR%s\n", Cyan, Reset)
	cetakGaris("─", L)

	var totalJam float64
	jumlahTidur := 0

	for i := 0; i < sleepCount; i++ {
		if selisihHari(sleepData[i].Date) <= 7 {
			jumlahTidur++
			totalJam += sleepData[i].Duration
		}
	}

	if jumlahTidur == 0 {
		fmt.Println(Blue + "  ℹ  Belum ada data tidur minggu ini." + Reset)
	} else {
		fmt.Printf("  %-12s  %-7s  %-7s  %-9s  Bar\n", "Tanggal", "Tidur", "Bangun", "Durasi")
		cetakGaris("─", L)
		for i := 0; i < sleepCount; i++ {
			r := sleepData[i]
			if selisihHari(r.Date) <= 7 {
				fmt.Printf("  %-12s  %-7s  %-7s  %s jam   %s\n",
					r.Date, r.BedTime, r.WakeTime,
					padKanan(floatToStr(r.Duration, 1), 4),
					sleepBar(r.Duration))
			}
		}
		cetakGaris("─", L)
		rataJam := totalJam / float64(jumlahTidur)
		fmt.Printf("  %-32s: %s%s jam/malam%s\n", "Rata-rata tidur", Bold, floatToStr(rataJam, 2), Reset)
		fmt.Printf("  %-32s: %d hari\n", "Total catatan", jumlahTidur)
		fmt.Printf("  %-32s: ", "Evaluasi")
		evalTidur(rataJam)
	}

	// ════ BAGIAN KESEHATAN ════
	fmt.Printf("\n  %s💊 STATISTIK KESEHATAN%s\n", Cyan, Reset)
	cetakGaris("─", L)

	var totalAir float64
	var totalOlahraga int
	var totalMood int
	jumlahHealth := 0

	for i := 0; i < healthCount; i++ {
		if selisihHari(healthData[i].Date) <= 7 {
			jumlahHealth++
			totalAir += healthData[i].TotalWaterML
			totalOlahraga += healthData[i].ExerciseMins
			totalMood += healthData[i].MoodScore
		}
	}

	if jumlahHealth == 0 {
		fmt.Println(Blue + "  ℹ  Belum ada data kesehatan minggu ini." + Reset)
	} else {
		fmt.Printf("  %-12s  %-10s  %-10s  %s\n", "Tanggal", "Air (ml)", "Olahraga", "Mood")
		cetakGaris("─", L)
		for i := 0; i < healthCount; i++ {
			r := healthData[i]
			if selisihHari(r.Date) <= 7 {
				ol := intToStr(r.ExerciseMins) + " mnt"
				if r.ExerciseMins == 0 {
					ol = "─"
				}
				fmt.Printf("  %-12s  %-10s  %-10s  %s\n",
					r.Date, intToStr(int(r.TotalWaterML)), ol, r.Mood)
			}
		}
		cetakGaris("─", L)
		rataAir := totalAir / float64(jumlahHealth)
		rataOlahraga := totalOlahraga / jumlahHealth
		rataMood := float64(totalMood) / float64(jumlahHealth)

		fmt.Printf("  %-32s: %s ml (%.1f L)\n", "Rata-rata air", intToStr(int(rataAir)), rataAir/1000)
		fmt.Printf("  %-32s: %d menit/hari\n", "Rata-rata olahraga", rataOlahraga)
		fmt.Printf("  %-32s: %s/5.0 %s\n", "Rata-rata mood", floatToStr(rataMood, 1), moodBar(rataMood))
		fmt.Printf("  %-32s: %d hari\n", "Total catatan", jumlahHealth)
		cetakGaris("─", L)
		fmt.Printf("%s  💡 Evaluasi Singkat:%s\n", Bold, Reset)
		evalKesehatan(rataAir, rataOlahraga)
		if rataMood >= 4 {
			fmt.Printf("  %-22s: %s✅ Mood POSITIF%s\n", "Suasana hati", Green, Reset)
		} else if rataMood >= 3 {
			fmt.Printf("  %-22s: %s😐 Mood BIASA%s\n", "Suasana hati", Yellow, Reset)
		} else {
			fmt.Printf("  %-22s: %s😢 Mood KURANG BAIK%s\n", "Suasana hati", Red, Reset)
		}
	}

	readString("\n  Tekan Enter untuk kembali...")
}

// ==================== tabel menu ====================

func mainMenu() {
	for {
		clearScreen()
		cetakGaris("═", 40)
		fmt.Printf("%s%s  🩺 KESEHATAN & POLA TIDUR  🌙%s\n", Bold, Cyan, Reset)
		cetakGaris("═", 40)
		t := time.Now()
		fmt.Printf("  %sTanggal:%s %s\n\n", Bold, Reset, t.Format("Monday, 02 January 2006"))

		fmt.Printf("  %s╔══ MENU PENCATATAN ════════════╗%s\n", Blue, Reset)
		fmt.Printf("  %s║%s  1. 🌙  Catat Tidur           %s║%s\n", Blue, Reset, Blue, Reset)
		fmt.Printf("  %s║%s  2. 💊  Catat Kesehatan Harian%s║%s\n", Blue, Reset, Blue, Reset)
		fmt.Printf("  %s╠══ MENU STATISTIK ═════════════╣%s\n", Blue, Reset)
		fmt.Printf("  %s║%s  3. 📊  Lihat Statistik       %s║%s\n", Blue, Reset, Blue, Reset)
		fmt.Printf("  %s╠═══════════════════════════════╣%s\n", Blue, Reset)
		fmt.Printf("  %s║%s  0. 🚪  Keluar                %s║%s\n", Blue, Reset, Blue, Reset)
		fmt.Printf("  %s╚═══════════════════════════════╝%s\n", Blue, Reset)

		pilihan := readString("\n  Pilih menu (0-3): ")

		switch pilihan {
		case "1":
			inputTidur()
		case "2":
			inputKesehatan()
		case "3":
			lihatStatistik()
		case "0":
			clearScreen()
			fmt.Printf("\n  %s👋 Terima kasih! Jaga kesehatan Anda selalu!%s\n\n", Green, Reset)
			return
		default:
			fmt.Println(Red + "  ⚠️  Pilihan tidak valid! Masukkan angka 0-3." + Reset)
			readString("  Tekan Enter untuk melanjutkan...")
		}
	}
}

// ==================== main func ====================

func main() {
	mainMenu()
}
