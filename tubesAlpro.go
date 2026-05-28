package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// ==================== STRUCTS ====================

type SleepRecord struct {
	Date     string  `json:"date"`
	BedTime  string  `json:"bed_time"`
	WakeTime string  `json:"wake_time"`
	Duration float64 `json:"duration_hours"`
}

type HealthRecord struct {
	Date         string  `json:"date"`
	WaterGlass   int     `json:"water_glass"`
	WaterBottle  float64 `json:"water_bottle"`
	TotalWaterML float64 `json:"total_water_ml"`
	ExerciseMins int     `json:"exercise_mins"`
	ExerciseType string  `json:"exercise_type"`
	Mood         string  `json:"mood"`
	MoodScore    int     `json:"mood_score"`
}

type AppData struct {
	SleepRecords  []SleepRecord  `json:"sleep_records"`
	HealthRecords []HealthRecord `json:"health_records"`
}

// ==================== GLOBALS ====================

var data AppData
var scanner = bufio.NewScanner(os.Stdin)

const dataFile = "health_data.json"

// Warna ANSI
const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	White  = "\033[97m"
	BgBlue = "\033[44m"
)

// ==================== HELPERS ====================

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readInt(prompt string) (int, error) {
	input := readLine(prompt)
	val, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("⚠️  ERROR: Input '%s' bukan angka! Harap masukkan angka yang valid", input)
	}
	return val, nil
}

func readFloat(prompt string) (float64, error) {
	input := readLine(prompt)
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("⚠️  ERROR: Input '%s' bukan angka! Harap masukkan angka yang valid", input)
	}
	return val, nil
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printLine(char string, length int) {
	fmt.Println(strings.Repeat(char, length))
}

func printHeader(title string) {
	printLine("═", 55)
	fmt.Printf("%s%s  %s%s\n", Bold, Cyan, title, Reset)
	printLine("═", 55)
}

func printSuccess(msg string) {
	fmt.Printf("%s✅  %s%s\n", Green, msg, Reset)
}

func printError(msg string) {
	fmt.Printf("%s%s%s\n", Red, msg, Reset)
}

func printInfo(msg string) {
	fmt.Printf("%s ℹ  %s%s\n", Blue, msg, Reset)
}

func today() string {
	return time.Now().Format("2006-01-02")
}

func parseDateToTime(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// ==================== DATA PERSISTENCE ====================

func loadData() {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		data = AppData{}
		return
	}
	json.Unmarshal(file, &data)
}

func saveData() error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, bytes, 0644)
}

// ==================== SLEEP MODULE ====================

func parseSleepDuration(bedTime, wakeTime string) (float64, error) {
	layout := "15:04"
	bed, err := time.Parse(layout, bedTime)
	if err != nil {
		return 0, fmt.Errorf("format waktu tidur tidak valid, gunakan HH:MM (contoh: 22:30)")
	}
	wake, err := time.Parse(layout, wakeTime)
	if err != nil {
		return 0, fmt.Errorf("format waktu bangun tidak valid, gunakan HH:MM (contoh: 06:00)")
	}

	diff := wake.Sub(bed).Hours()
	if diff < 0 {
		diff += 24 // tidur melewati tengah malam
	}
	return math.Round(diff*100) / 100, nil
}

func inputSleep() {
	clearScreen()
	printHeader("🌙  PENCATATAN TIDUR")

	dateStr := today()
	fmt.Printf("%sTanggal:%s %s\n\n", Bold, Reset, dateStr)

	bedTime := readLine("  Jam tidur (HH:MM, contoh 22:30) : ")
	wakeTime := readLine("  Jam bangun (HH:MM, contoh 06:00): ")

	duration, err := parseSleepDuration(bedTime, wakeTime)
	if err != nil {
		printError("  ⚠️  ERROR: " + err.Error())
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}

	record := SleepRecord{
		Date:     dateStr,
		BedTime:  bedTime,
		WakeTime: wakeTime,
		Duration: duration,
	}

	// Cek apakah sudah ada record hari ini, kalau ada timpa
	updated := false
	for i, r := range data.SleepRecords {
		if r.Date == dateStr {
			data.SleepRecords[i] = record
			updated = true
			break
		}
	}
	if !updated {
		data.SleepRecords = append(data.SleepRecords, record)
	}

	if err := saveData(); err != nil {
		printError("  Gagal menyimpan data: " + err.Error())
	} else {
		fmt.Println()
		printSuccess(fmt.Sprintf("Data tidur tersimpan! Durasi tidur: %.2f jam", duration))
		evaluateSleep(duration)
	}

	readLine("\n  Tekan Enter untuk kembali...")
}

func evaluateSleep(hours float64) {
	fmt.Println()
	switch {
	case hours >= 7 && hours <= 9:
		fmt.Printf("  %s😴 Kualitas tidur BAIK! Durasi ideal 7-9 jam terpenuhi.%s\n", Green, Reset)
	case hours >= 6 && hours < 7:
		fmt.Printf("  %s😐 Tidur sedikit KURANG. Coba tidur 30 menit lebih awal.%s\n", Yellow, Reset)
	case hours > 9:
		fmt.Printf("  %s😪 Tidur terlalu LAMA. Tidur berlebih bisa sebabkan lesu.%s\n", Yellow, Reset)
	default:
		fmt.Printf("  %s😵 Tidur sangat KURANG! Kurang dari 6 jam berbahaya bagi kesehatan.%s\n", Red, Reset)
	}
}

func viewSleepStats() {
	clearScreen()
	printHeader("📊  STATISTIK TIDUR - 7 HARI TERAKHIR")

	now := time.Now()
	var weekRecords []SleepRecord
	var totalHours float64

	for _, r := range data.SleepRecords {
		t, err := parseDateToTime(r.Date)
		if err != nil {
			continue
		}
		if now.Sub(t).Hours() <= 7*24 {
			weekRecords = append(weekRecords, r)
			totalHours += r.Duration
		}
	}

	if len(weekRecords) == 0 {
		printInfo("Belum ada data tidur minggu ini.")
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}

	fmt.Printf("\n  %-12s %-10s %-10s %-10s\n", "Tanggal", "Tidur", "Bangun", "Durasi")
	printLine("─", 45)
	for _, r := range weekRecords {
		bar := sleepBar(r.Duration)
		fmt.Printf("  %-12s %-10s %-10s %-6.1f jam %s\n",
			r.Date, r.BedTime, r.WakeTime, r.Duration, bar)
	}
	printLine("─", 45)

	avg := totalHours / float64(len(weekRecords))
	fmt.Printf("\n  %sRata-rata tidur : %.2f jam/malam%s\n", Bold, avg, Reset)
	fmt.Printf("  Total catatan   : %d hari\n", len(weekRecords))
	fmt.Println()
	evaluateSleep(avg)

	readLine("\n  Tekan Enter untuk kembali...")
}

func sleepBar(hours float64) string {
	filled := int(math.Round(hours / 9 * 10))
	if filled > 10 {
		filled = 10
	}
	bar := "[" + strings.Repeat("█", filled) + strings.Repeat("░", 10-filled) + "]"
	if hours >= 7 {
		return Green + bar + Reset
	} else if hours >= 6 {
		return Yellow + bar + Reset
	}
	return Red + bar + Reset
}

// ==================== HEALTH MODULE ====================

func moodOptions() map[string]int {
	return map[string]int{
		"Sangat Bahagia 😄": 5,
		"Bahagia 😊":        4,
		"Biasa Saja 😐":     3,
		"Sedih 😢":          2,
		"Sangat Buruk 😭":   1,
	}
}

func inputHealth() {
	clearScreen()
	printHeader("💊  PENCATATAN KESEHATAN HARIAN")

	dateStr := today()
	fmt.Printf("%sTanggal:%s %s\n", Bold, Reset, dateStr)

	// --- Air Putih ---
	fmt.Printf("\n%s💧 KONSUMSI AIR PUTIH%s\n", Cyan, Reset)
	printLine("─", 35)

	glass, err := readInt("  Jumlah gelas (ukuran 250ml)  : ")
	if err != nil {
		printError("  " + err.Error())
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}

	bottle, err := readFloat("  Jumlah botol (ukuran 600ml)  : ")
	if err != nil {
		printError("  " + err.Error())
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}

	totalML := float64(glass)*250 + bottle*600

	// --- Olahraga ---
	fmt.Printf("\n%s🏃 OLAHRAGA%s\n", Cyan, Reset)
	printLine("─", 35)

	exerciseMins, err := readInt("  Durasi olahraga (menit, 0 jika tidak): ")
	if err != nil {
		printError("  " + err.Error())
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}

	exerciseType := ""
	if exerciseMins > 0 {
		exerciseType = readLine("  Jenis olahraga (contoh: Lari, Gym)    : ")
		if exerciseType == "" {
			exerciseType = "Tidak disebutkan"
		}
	}

	// --- Mood ---
	fmt.Printf("\n%s😊 SUASANA HATI (MOOD)%s\n", Cyan, Reset)
	printLine("─", 35)
	moods := []string{
		"Sangat Bahagia 😄",
		"Bahagia 😊",
		"Biasa Saja 😐",
		"Sedih 😢",
		"Sangat Buruk 😭",
	}
	moodScores := []int{5, 4, 3, 2, 1}

	for i, m := range moods {
		fmt.Printf("  %d. %s\n", i+1, m)
	}

	moodIdx, err := readInt("\n  Pilih mood hari ini (1-5): ")
	if err != nil {
		printError("  " + err.Error())
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}
	if moodIdx < 1 || moodIdx > 5 {
		printError("  ⚠️  ERROR: Pilihan mood harus antara 1 dan 5!")
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}

	selectedMood := moods[moodIdx-1]
	moodScore := moodScores[moodIdx-1]

	record := HealthRecord{
		Date:         dateStr,
		WaterGlass:   glass,
		WaterBottle:  bottle,
		TotalWaterML: totalML,
		ExerciseMins: exerciseMins,
		ExerciseType: exerciseType,
		Mood:         selectedMood,
		MoodScore:    moodScore,
	}

	updated := false
	for i, r := range data.HealthRecords {
		if r.Date == dateStr {
			data.HealthRecords[i] = record
			updated = true
			break
		}
	}
	if !updated {
		data.HealthRecords = append(data.HealthRecords, record)
	}

	if err := saveData(); err != nil {
		printError("  Gagal menyimpan data!")
	} else {
		fmt.Println()
		printLine("─", 45)
		printSuccess("Data kesehatan berhasil disimpan!")
		fmt.Println()
		fmt.Printf("  💧 Total air    : %.0f ml (%.1f L)\n", totalML, totalML/1000)
		if exerciseMins > 0 {
			fmt.Printf("  🏃 Olahraga     : %d menit (%s)\n", exerciseMins, exerciseType)
		} else {
			fmt.Printf("  🏃 Olahraga     : Tidak ada hari ini\n")
		}
		fmt.Printf("  😊 Mood         : %s\n", selectedMood)
		fmt.Println()
		evaluateHealth(totalML, exerciseMins)
	}

	readLine("\n  Tekan Enter untuk kembali...")
}

func evaluateHealth(waterML float64, exerciseMins int) {
	fmt.Printf("%s📋 EVALUASI SINGKAT:%s\n", Bold, Reset)
	printLine("─", 35)

	// Evaluasi air
	if waterML >= 2000 {
		fmt.Printf("  %s✅ Konsumsi air CUKUP (≥2L). Tubuh terhidrasi baik!%s\n", Green, Reset)
	} else if waterML >= 1500 {
		fmt.Printf("  %s⚠️  Konsumsi air HAMPIR CUKUP. Tambah 1-2 gelas lagi.%s\n", Yellow, Reset)
	} else {
		fmt.Printf("  %s❌ Konsumsi air KURANG! Target minimal 2 liter/hari.%s\n", Red, Reset)
	}

	// Evaluasi olahraga
	if exerciseMins >= 30 {
		fmt.Printf("  %s✅ Durasi olahraga BAIK (≥30 menit). Pertahankan!%s\n", Green, Reset)
	} else if exerciseMins > 0 {
		fmt.Printf("  %s⚠️  Olahraga hanya %d menit. WHO rekomendasikan 30 menit/hari.%s\n", Yellow, exerciseMins, Reset)
	} else {
		fmt.Printf("  %s❌ Tidak olahraga hari ini. Coba gerak ringan 15-30 menit!%s\n", Red, Reset)
	}
}

func viewHealthStats() {
	clearScreen()
	printHeader("📈  RINGKASAN KESEHATAN - 7 HARI TERAKHIR")

	now := time.Now()
	var weekRecords []HealthRecord
	var totalWater float64
	var totalExercise int
	var totalMoodScore int

	for _, r := range data.HealthRecords {
		t, err := parseDateToTime(r.Date)
		if err != nil {
			continue
		}
		if now.Sub(t).Hours() <= 7*24 {
			weekRecords = append(weekRecords, r)
			totalWater += r.TotalWaterML
			totalExercise += r.ExerciseMins
			totalMoodScore += r.MoodScore
		}
	}

	if len(weekRecords) == 0 {
		printInfo("Belum ada data kesehatan minggu ini.")
		readLine("\n  Tekan Enter untuk kembali...")
		return
	}

	n := len(weekRecords)
	avgWater := totalWater / float64(n)
	avgExercise := totalExercise / n
	avgMood := float64(totalMoodScore) / float64(n)

	fmt.Printf("\n  %-12s %-10s %-12s %-10s\n", "Tanggal", "Air (ml)", "Olahraga", "Mood")
	printLine("─", 50)
	for _, r := range weekRecords {
		ex := fmt.Sprintf("%d mnt", r.ExerciseMins)
		if r.ExerciseMins == 0 {
			ex = "─"
		}
		fmt.Printf("  %-12s %-10.0f %-12s %s\n",
			r.Date, r.TotalWaterML, ex, r.Mood)
	}
	printLine("─", 50)

	fmt.Printf("\n  %s📊 RATA-RATA MINGGUAN%s\n", Bold, Reset)
	printLine("─", 35)
	fmt.Printf("  💧 Air          : %.0f ml/hari (%.1f L)\n", avgWater, avgWater/1000)
	fmt.Printf("  🏃 Olahraga     : %d menit/hari\n", avgExercise)
	fmt.Printf("  😊 Mood         : %.1f/5.0 %s\n", avgMood, moodBar(avgMood))
	fmt.Printf("  📅 Total catatan: %d hari\n", n)

	fmt.Printf("\n  %s💡 EVALUASI SINGKAT:%s\n", Bold, Reset)
	printLine("─", 35)

	if avgWater >= 2000 {
		fmt.Printf("  %s✅ Rata-rata hidrasi BAIK minggu ini!%s\n", Green, Reset)
	} else {
		fmt.Printf("  %s❌ Rata-rata hidrasi KURANG. Perbanyak minum air putih.%s\n", Red, Reset)
	}

	if avgExercise >= 30 {
		fmt.Printf("  %s✅ Aktivitas fisik KONSISTEN minggu ini!%s\n", Green, Reset)
	} else if avgExercise > 0 {
		fmt.Printf("  %s⚠️  Aktivitas fisik ADA tapi belum konsisten.%s\n", Yellow, Reset)
	} else {
		fmt.Printf("  %s❌ Hampir tidak ada aktivitas fisik minggu ini.%s\n", Red, Reset)
	}

	if avgMood >= 4 {
		fmt.Printf("  %s✅ Suasana hati POSITIF minggu ini! Pertahankan.%s\n", Green, Reset)
	} else if avgMood >= 3 {
		fmt.Printf("  %s😐 Suasana hati BIASA minggu ini.%s\n", Yellow, Reset)
	} else {
		fmt.Printf("  %s😢 Suasana hati KURANG BAIK. Luangkan waktu untuk diri sendiri.%s\n", Red, Reset)
	}

	readLine("\n  Tekan Enter untuk kembali...")
}

func moodBar(score float64) string {
	filled := int(math.Round(score / 5 * 10))
	bar := "[" + strings.Repeat("♥", filled) + strings.Repeat("·", 10-filled) + "]"
	if score >= 4 {
		return Green + bar + Reset
	} else if score >= 3 {
		return Yellow + bar + Reset
	}
	return Red + bar + Reset
}

// ==================== MENU ====================

func mainMenu() {
	for {
		clearScreen()
		printLine("═", 55)
		fmt.Printf("%s%s   🩺  APLIKASI KESEHATAN & POLA TIDUR  🌙%s\n", Bold, Cyan, Reset)
		printLine("═", 55)
		fmt.Printf("  %sTanggal:%s %s\n\n", Bold, Reset, time.Now().Format("Monday, 02 January 2006"))

		fmt.Printf("  %s╔══ MENU PENCATATAN ══╗%s\n", Blue, Reset)
		fmt.Println("  ║  1. 🌙  Catat Tidur            ║")
		fmt.Println("  ║  2. 💊  Catat Kesehatan Harian  ║")
		fmt.Printf("  %s╠══ MENU STATISTIK ════╣%s\n", Blue, Reset)
		fmt.Println("  ║  3. 📊  Statistik Tidur         ║")
		fmt.Println("  ║  4. 📈  Ringkasan Kesehatan      ║")
		fmt.Printf("  %s╠══════════════════════╣%s\n", Blue, Reset)
		fmt.Println("  ║  0. 🚪  Keluar                  ║")
		fmt.Printf("  %s╚══════════════════════╝%s\n", Blue, Reset)

		choice := readLine("\n  Pilih menu (0-4): ")

		switch choice {
		case "1":
			inputSleep()
		case "2":
			inputHealth()
		case "3":
			viewSleepStats()
		case "4":
			viewHealthStats()
		case "0":
			clearScreen()
			fmt.Printf("\n  %s👋 Terima kasih! Jaga kesehatan Anda selalu!%s\n\n", Green, Reset)
			os.Exit(0)
		default:
			printError("  ⚠️  Pilihan tidak valid! Masukkan angka 0-4.")
			readLine("  Tekan Enter untuk melanjutkan...")
		}
	}
}

// ==================== MAIN ====================

func main() {
	loadData()
	mainMenu()
}
