package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const KAPASITAS_MAKSIMAL_TUGAS = 100

type Tugas struct {
	MataKuliah string
	Kategori   string
	Deadline   string
	Jam        string
	Status     bool
}

type KumpulanDataTugas [KAPASITAS_MAKSIMAL_TUGAS]Tugas

var daftarTugasUtama KumpulanDataTugas
var running bool = true 

func main() {
	var pilihanMenu int
	pembacaInput := bufio.NewReader(os.Stdin)

	for running {
		fmt.Println("\n=== APLIKASI MANAJEMEN TUGAS HARIAN (KHUSUS JUNI 2026) ===")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Cari Tugas")
		fmt.Println("3. Ubah Tugas")
		fmt.Println("4. Hapus Tugas")
		fmt.Println("5. Urutkan Tugas")
		fmt.Println("6. Ubah Status Tugas")
		fmt.Println("7. Tampilkan Semua Tugas")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scanf("%d\n", &pilihanMenu)

		switch pilihanMenu {
		case 1:
			tambahTugas(pembacaInput)
		case 2:
			cariTugas()
		case 3:
			ubahTugas(pembacaInput)
		case 4:
			hapusTugas()
		case 5:
			urutkanTugas()
		case 6:
			ubahStatusTugas()
		case 7:
			tampilkanSemuaTugas()
		case 0:
			fmt.Println("Terima kasih! Tetap semangat menyelesaikan tugasmu.")
			running = false 
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi!")
		}
	}
}

func hitungJumlahTugas() int {
	jumlahTugasTersimpan := 0
	for indeks := 0; indeks < KAPASITAS_MAKSIMAL_TUGAS; indeks++ {
		if daftarTugasUtama[indeks].MataKuliah != "" {
			jumlahTugasTersimpan++
		}
	}
	return jumlahTugasTersimpan
}

func validasiTanggal(tanggalInput string) bool {
	if len(tanggalInput) != 10 || tanggalInput[4] != '-' || tanggalInput[7] != '-' {
		return false
	}

	var tahun, bulan, hari int
	if _, err := fmt.Sscanf(tanggalInput, "%4d-%2d-%2d", &tahun, &bulan, &hari); err != nil {
		return false
	}

	if tahun != 2026 || bulan != 6 {
		fmt.Println("⚠️ Batasan: Aplikasi ini hanya menerima input untuk bulan Juni 2026!")
		return false
	}

	if hari < 1 || hari > 30 {
		return false
	}

	return true
}

func validasiJam(jamInput string) bool {
	if len(jamInput) != 5 || jamInput[2] != ':' {
		return false
	}

	var jam, menit int
	fmt.Sscanf(jamInput, "%2d:%2d", &jam, &menit)
	if jam < 0 || jam > 23 || menit < 0 || menit > 59 {
		return false
	}
	return true
}

func hitungSisaHari(tanggalDeadline string) int {
	var tahun, bulan, hari int
	fmt.Sscanf(tanggalDeadline, "%4d-%2d-%2d", &tahun, &bulan, &hari)

	waktuSekarang := time.Now()
	tanggalHariIni := waktuSekarang.Day()

	return hari - tanggalHariIni
}

func berikanFeedback(tugasBaru Tugas) {
	fmt.Println("\nData berhasil diproses!")
	sisaHari := hitungSisaHari(tugasBaru.Deadline)

	if sisaHari < 0 {
		fmt.Println("⚠️ Peringatan: Deadline tugas ini sudah lewat di bulan ini!")
	} else if sisaHari <= 2 {
		fmt.Printf("⚠️ Bahaya: Tugas %s mendekati deadline (%d hari lagi)! Prioritaskan tugas ini.\n", tugasBaru.MataKuliah, sisaHari)
	} else {
		fmt.Printf("💬 Info: Waktu pengerjaan masih aman (%d hari lagi). Cicil sedikit demi sedikit!\n", sisaHari)
	}

	if strings.EqualFold(tugasBaru.Kategori, "Tubes") || strings.EqualFold(tugasBaru.Kategori, "UAS") {
		fmt.Println("💡 Rekomendasi: Kategori 'Tubes/UAS' membutuhkan fokus tinggi. Jangan ditunda sampai malam terakhir!")
	}
}

func tambahTugas(pembacaInput *bufio.Reader) {
	jumlahTugasSaatIni := hitungJumlahTugas()
	if jumlahTugasSaatIni >= KAPASITAS_MAKSIMAL_TUGAS {
		fmt.Println("Daftar tugas sudah penuh.")
		return
	}

	fmt.Print("Masukkan Mata Kuliah: ")
	mataKuliahInput, _ := pembacaInput.ReadString('\n')
	mataKuliahInput = strings.TrimSpace(mataKuliahInput)

	fmt.Print("Masukkan Kategori (Tubes/UAS/Kuis/Lainnya): ")
	kategoriInput, _ := pembacaInput.ReadString('\n')
	kategoriInput = strings.TrimSpace(kategoriInput)

	var tanggalDeadlineInput string
	apakahTanggalValid := false
	for !apakahTanggalValid {
		fmt.Print("Masukkan Deadline (YYYY-MM-DD): ")
		fmt.Scanf("%s\n", &tanggalDeadlineInput)
		if validasiTanggal(tanggalDeadlineInput) {
			apakahTanggalValid = true
		} else {
			fmt.Println("Format tanggal salah, tidak ada, atau di luar Juni 2026. Coba lagi.")
		}
	}

	var jamPengumpulanInput string
	apakahJamValid := false
	for !apakahJamValid {
		fmt.Print("Masukkan Jam Pengumpulan (HH:MM): ")
		fmt.Scanf("%s\n", &jamPengumpulanInput)
		if validasiJam(jamPengumpulanInput) {
			apakahJamValid = true
		} else {
			fmt.Println("Format jam salah. Coba lagi.")
		}
	}

	daftarTugasUtama[jumlahTugasSaatIni] = Tugas{
		MataKuliah: mataKuliahInput,
		Kategori:   kategoriInput,
		Deadline:   tanggalDeadlineInput,
		Jam:        jamPengumpulanInput,
		Status:     false,
	}

	berikanFeedback(daftarTugasUtama[jumlahTugasSaatIni])
}

func cetakSatuTugas(tugasYgDicetak Tugas, nomorUrut int) {
	statusString := "Belum Selesai"
	if tugasYgDicetak.Status {
		statusString = "Selesai"
	}
	fmt.Printf("%d. [%s] %s | Deadline: %s Pukul %s | Status: %s\n",
		nomorUrut+1, tugasYgDicetak.Kategori, tugasYgDicetak.MataKuliah, tugasYgDicetak.Deadline, tugasYgDicetak.Jam, statusString)
}

func tampilkanSemuaTugas() {
	totalTugas := hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	fmt.Println("\n=== DAFTAR SEMUA TUGAS ===")
	for indeks := 0; indeks < totalTugas; indeks++ {
		cetakSatuTugas(daftarTugasUtama[indeks], indeks)
	}
}

func cariTugas() {
	totalTugas := hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	var pilihanMetodeCari int
	fmt.Println("\nCari Berdasarkan:")
	fmt.Println("1. Kategori (Sequential Search)")
	fmt.Println("2. Deadline (Binary Search)")
	fmt.Println("3. Tanggal & Jam Spesifik (Sequential Search)")
	fmt.Print("Pilihan: ")
	fmt.Scanf("%d\n", &pilihanMetodeCari)

	if pilihanMetodeCari == 1 {
		var kategoriDicari string
		fmt.Print("Masukkan kategori yang dicari: ")
		fmt.Scanf("%s\n", &kategoriDicari)
		apakahDitemukan := false

		for indeks := 0; indeks < totalTugas; indeks++ {
			if strings.EqualFold(daftarTugasUtama[indeks].Kategori, kategoriDicari) {
				cetakSatuTugas(daftarTugasUtama[indeks], indeks)
				apakahDitemukan = true
			}
		}
		if !apakahDitemukan {
			fmt.Println("Data tugas dengan kategori tersebut tidak ditemukan.")
		}

	} else if pilihanMetodeCari == 2 {
		selectionSortTugas(true) 
		var tanggalDicari string
		fmt.Print("Masukkan tanggal deadline yang dicari (YYYY-MM-DD): ")
		fmt.Scanf("%s\n", &tanggalDicari)

		posisiKiri := 0
		posisiKanan := totalTugas - 1
		indeksHasilTemuan := -1

		for posisiKiri <= posisiKanan && indeksHasilTemuan == -1 {
			titikTengah := (posisiKiri + posisiKanan) / 2
			if daftarTugasUtama[titikTengah].Deadline == tanggalDicari {
				indeksHasilTemuan = titikTengah
			} else if daftarTugasUtama[titikTengah].Deadline < tanggalDicari {
				posisiKiri = titikTengah + 1
			} else {
				posisiKanan = titikTengah - 1
			}
		}

		if indeksHasilTemuan != -1 {
			fmt.Println("\nData ditemukan:")
			for indeks := 0; indeks < totalTugas; indeks++ {
				if daftarTugasUtama[indeks].Deadline == tanggalDicari {
					cetakSatuTugas(daftarTugasUtama[indeks], indeks)
				}
			}
		} else {
			fmt.Println("Data dengan deadline tersebut tidak ditemukan.")
		}

	} else if pilihanMetodeCari == 3 {
		var tanggalDicari, jamDicari string
		fmt.Print("Masukkan Tanggal (YYYY-MM-DD): ")
		fmt.Scanf("%s\n", &tanggalDicari)
		fmt.Print("Masukkan Jam (HH:MM): ")
		fmt.Scanf("%s\n", &jamDicari)

		apakahDitemukan := false
		for indeks := 0; indeks < totalTugas; indeks++ {
			if daftarTugasUtama[indeks].Deadline == tanggalDicari && daftarTugasUtama[indeks].Jam == jamDicari {
				cetakSatuTugas(daftarTugasUtama[indeks], indeks)
				apakahDitemukan = true
			}
		}
		if !apakahDitemukan {
			fmt.Println("Data dengan kombinasi tanggal & jam tersebut tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan pencarian tidak valid.")
	}
}

func ubahTugas(pembacaInput *bufio.Reader) {
	totalTugas := hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Tidak ada tugas yang bisa diubah.")
		return
	}

	tampilkanSemuaTugas()

	var nomorTugasPilihan int
	fmt.Print("Masukkan nomor tugas yang ingin diubah: ")
	fmt.Scanf("%d\n", &nomorTugasPilihan)
	indeksPilihan := nomorTugasPilihan - 1

	if indeksPilihan < 0 || indeksPilihan >= totalTugas {
		fmt.Println("Nomor tidak valid.")
		return
	}

	fmt.Print("Masukkan Mata Kuliah baru: ")
	mataKuliahBaru, _ := pembacaInput.ReadString('\n')
	mataKuliahBaru = strings.TrimSpace(mataKuliahBaru)

	fmt.Print("Masukkan Kategori baru: ")
	kategoriBaru, _ := pembacaInput.ReadString('\n')
	kategoriBaru = strings.TrimSpace(kategoriBaru)

	var tanggalDeadlineBaru, jamPengumpulanBaru string
	fmt.Print("Masukkan Deadline baru (YYYY-MM-DD): ")
	fmt.Scanf("%s\n", &tanggalDeadlineBaru)
	fmt.Print("Masukkan Jam Pengumpulan baru (HH:MM): ")
	fmt.Scanf("%s\n", &jamPengumpulanBaru)

	if validasiTanggal(tanggalDeadlineBaru) && validasiJam(jamPengumpulanBaru) {
		daftarTugasUtama[indeksPilihan].MataKuliah = mataKuliahBaru
		daftarTugasUtama[indeksPilihan].Kategori = kategoriBaru
		daftarTugasUtama[indeksPilihan].Deadline = tanggalDeadlineBaru
		daftarTugasUtama[indeksPilihan].Jam = jamPengumpulanBaru
		fmt.Println("Data tugas berhasil diperbarui!")
	} else {
		fmt.Println("Perubahan gagal. Pastikan tanggal berada di rentang Juni 2026.")
	}
}

func hapusTugas() {
	totalTugas := hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Tidak ada tugas untuk dihapus.")
		return
	}

	tampilkanSemuaTugas()

	var nomorTugasPilihan int
	fmt.Print("Masukkan nomor tugas yang ingin dihapus: ")
	fmt.Scanf("%d\n", &nomorTugasPilihan)
	indeksPilihan := nomorTugasPilihan - 1

	if indeksPilihan < 0 || indeksPilihan >= totalTugas {
		fmt.Println("Nomor tidak valid.")
		return
	}

	var teksKonfirmasi string
	fmt.Printf("Apakah Anda yakin menghapus tugas %s? (y/n): ", daftarTugasUtama[indeksPilihan].MataKuliah)
	fmt.Scanf("%s\n", &teksKonfirmasi)

	if teksKonfirmasi == "y" || teksKonfirmasi == "Y" {
		for indeks := indeksPilihan; indeks < totalTugas-1; indeks++ {
			daftarTugasUtama[indeks] = daftarTugasUtama[indeks+1]
		}
		daftarTugasUtama[totalTugas-1] = Tugas{}
		fmt.Println("Tugas berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

func urutkanTugas() {
	totalTugas := hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Tidak ada data untuk diurutkan.")
		return
	}

	var pilihanMetode, pilihanKriteria int
	fmt.Println("\nPilih Metode Pengurutan:")
	fmt.Println("1. Selection Sort (Berdasarkan Deadline)")
	fmt.Println("2. Insertion Sort (Berdasarkan Status)")
	fmt.Print("Pilihan: ")
	fmt.Scanf("%d\n", &pilihanMetode)

	if pilihanMetode == 1 {
		fmt.Println("Urutkan Deadline secara:")
		fmt.Println("1. Terdekat (Ascending)")
		fmt.Println("2. Terlama (Descending)")
		fmt.Print("Pilihan: ")
		fmt.Scanf("%d\n", &pilihanKriteria)
		selectionSortTugas(pilihanKriteria == 1)
		fmt.Println("Data berhasil diurutkan berdasarkan Deadline Juni 2026!")
	} else if pilihanMetode == 2 {
		fmt.Println("Urutkan Status secara:")
		fmt.Println("1. Belum Selesai -> Selesai (Ascending)")
		fmt.Println("2. Selesai -> Belum Selesai (Descending)")
		fmt.Print("Pilihan: ")
		fmt.Scanf("%d\n", &pilihanKriteria)
		insertionSortTugas(pilihanKriteria == 1)
		fmt.Println("Data berhasil diurutkan berdasarkan Status!")
	} else {
		fmt.Println("Metode tidak tersedia.")
		return
	}

	tampilkanSemuaTugas()
}

func selectionSortTugas(apakahAscending bool) {
	totalTugas := hitungJumlahTugas()
	for indeksLuar := 0; indeksLuar < totalTugas-1; indeksLuar++ {
		indeksTerpilih := indeksLuar
		for indeksDalam := indeksLuar + 1; indeksDalam < totalTugas; indeksDalam++ {
			if apakahAscending {
				if daftarTugasUtama[indeksDalam].Deadline < daftarTugasUtama[indeksTerpilih].Deadline {
					indeksTerpilih = indeksDalam
				}
			} else {
				if daftarTugasUtama[indeksDalam].Deadline > daftarTugasUtama[indeksTerpilih].Deadline {
					indeksTerpilih = indeksDalam
				}
			}
		}

		daftarTugasUtama[indeksLuar], daftarTugasUtama[indeksTerpilih] = daftarTugasUtama[indeksTerpilih], daftarTugasUtama[indeksLuar]
	}
}

func insertionSortTugas(apakahAscending bool) {
	totalTugas := hitungJumlahTugas()
	for indeksLuar := 1; indeksLuar < totalTugas; indeksLuar++ {
		dataPenyimpanSementara := daftarTugasUtama[indeksLuar]
		indeksDalam := indeksLuar - 1

		nilaiBobotIndeksDalam := 0
		nilaiBobotSementara := 0

		loopTerus := indeksDalam >= 0
		for loopTerus {
			if daftarTugasUtama[indeksDalam].Status {
				nilaiBobotIndeksDalam = 1
			} else {
				nilaiBobotIndeksDalam = 0
			}
			if dataPenyimpanSementara.Status {
				nilaiBobotSementara = 1
			} else {
				nilaiBobotSementara = 0
			}

			apakahMemenuhiKondisiTukar := false
			if apakahAscending {
				apakahMemenuhiKondisiTukar = nilaiBobotIndeksDalam > nilaiBobotSementara
			} else {
				apakahMemenuhiKondisiTukar = nilaiBobotIndeksDalam < nilaiBobotSementara
			}

			if apakahMemenuhiKondisiTukar {
				daftarTugasUtama[indeksDalam+1] = daftarTugasUtama[indeksDalam]
				indeksDalam--
				loopTerus = indeksDalam >= 0
			} else {
				loopTerus = false 
			}
		}
		daftarTugasUtama[indeksDalam+1] = dataPenyimpanSementara
	}
}

func ubahStatusTugas() {
	totalTugas := hitungJumlahTugas()
	if totalTugas == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	tampilkanSemuaTugas()

	var nomorTugasPilihan, pilihanStatus int
	fmt.Print("Masukkan nomor tugas yang ingin diubah statusnya: ")
	fmt.Scanf("%d\n", &nomorTugasPilihan)
	indeksPilihan := nomorTugasPilihan - 1

	if indeksPilihan < 0 || indeksPilihan >= totalTugas {
		fmt.Println("Nomor tidak valid.")
		return
	}

	fmt.Println("Ubah status menjadi:")
	fmt.Println("1. Selesai")
	fmt.Println("2. Belum Selesai")
	fmt.Print("Pilihan: ")
	fmt.Scanf("%d\n", &pilihanStatus)

	if pilihanStatus == 1 {
		daftarTugasUtama[indeksPilihan].Status = true
		fmt.Println("🎉 Mantap! Tugas telah ditandai SELESAI.")
	} else if pilihanStatus == 2 {
		daftarTugasUtama[indeksPilihan].Status = false
		fmt.Println("Status diatur kembali ke BELUM SELESAI.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
