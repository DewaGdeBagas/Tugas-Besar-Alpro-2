package main

import "fmt"

// Struktur untuk mencatat detail servis
type Servis struct {
	Tanggal    string
	JenisRusak string
}

// Struktur utama kendaraan dengan kapasitas riwayat servis 10
type Kendaraan struct {
	PlatNomor    string
	Pemilik      string
	Tahun        int
	Riwayat      [10]Servis
	JumlahServis int
}

// Array utama untuk menampung data kendaraan dengan kapasitas 20
var daftarKendaraan [20]Kendaraan 
var totalKendaraan int

func cariSequential(plat string) int {
	var i int
	for i = 0; i < totalKendaraan; i = i + 1 {
		if daftarKendaraan[i].PlatNomor == plat {
			return i
		}
	}
	return -1
}

// Binary Search berdasarkan Plat Nomor (Syarat: Data harus urut Plat)
func cariBinary(plat string) int {
	var low, high, mid int
	low = 0
	high = totalKendaraan - 1
	for low <= high {
		mid = (low + high) / 2
		if daftarKendaraan[mid].PlatNomor == plat {
			return mid
		} else if daftarKendaraan[mid].PlatNomor < plat {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Fungsi Helper: Mengurutkan kendaraan berdasarkan Plat Nomor (Ascending)
// Wajib dipanggil sebelum melakukan cariBinary Plat Nomor
func urutPlatNomor() {
	var i, j int
	var key Kendaraan
	for i = 1; i < totalKendaraan; i = i + 1 {
		key = daftarKendaraan[i]
		j = i - 1
		for j >= 0 && daftarKendaraan[j].PlatNomor > key.PlatNomor {
			daftarKendaraan[j+1] = daftarKendaraan[j]
			j = j - 1
		}
		daftarKendaraan[j+1] = key
	}
}

func urutTahunSelection() {
	var i, j, idxMin int
	for i = 0; i < totalKendaraan-1; i = i + 1 {
		idxMin = i
		for j = i + 1; j < totalKendaraan; j = j + 1 {
			if daftarKendaraan[j].Tahun < daftarKendaraan[idxMin].Tahun {
				idxMin = j
			}
		}
		daftarKendaraan[i], daftarKendaraan[idxMin] = daftarKendaraan[idxMin], daftarKendaraan[i]
	}
}

func urutTahunInsertion() {
	var i, j int
	var key Kendaraan
	for i = 1; i < totalKendaraan; i = i + 1 {
		key = daftarKendaraan[i]
		j = i - 1
		for j >= 0 && daftarKendaraan[j].Tahun < key.Tahun {
			daftarKendaraan[j+1] = daftarKendaraan[j]
			j = j - 1
		}
		daftarKendaraan[j+1] = key
	}
}

// Binary Search untuk Tahun (Bisa mengembalikan banyak data yang sama)
// Mengembalikan indeks pertama yang ditemukan, atau -1 jika tidak ada
func cariBinaryTahun(targetTahun int) int {
	var low, high, mid int
	low = 0
	high = totalKendaraan - 1

	for low <= high {
		mid = (low + high) / 2
		if daftarKendaraan[mid].Tahun == targetTahun {
			return mid // Menemukan salah satu data yang cocok
		} else if daftarKendaraan[mid].Tahun < targetTahun {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func tambahKendaraan() {
	if totalKendaraan < 20 {
		fmt.Println("\n--- TAMBAH DATA KENDARAAN ---")
		fmt.Print("Masukkan Plat Nomor : ")
		fmt.Scan(&daftarKendaraan[totalKendaraan].PlatNomor)
		fmt.Print("Masukkan Nama Pemilik: ")
		fmt.Scan(&daftarKendaraan[totalKendaraan].Pemilik)
		fmt.Print("Masukkan Tahun       : ")
		fmt.Scan(&daftarKendaraan[totalKendaraan].Tahun)
		daftarKendaraan[totalKendaraan].JumlahServis = 0

		totalKendaraan = totalKendaraan + 1
		fmt.Println("Data kendaraan berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas kendaraan penuh! Maksimal hanya 20 kendaraan.")
	}
}

func ubahKendaraan() {
	var plat string
	var idx int
	fmt.Println("\n--- UBAH DATA KENDARAAN ---")
	fmt.Print("Masukkan Plat Nomor yang ingin diubah: ")
	fmt.Scan(&plat)

	idx = cariSequential(plat)
	if idx != -1 {
		fmt.Print("Masukkan Nama Pemilik Baru: ")
		fmt.Scan(&daftarKendaraan[idx].Pemilik)
		fmt.Print("Masukkan Tahun Baru       : ")
		fmt.Scan(&daftarKendaraan[idx].Tahun)
		fmt.Println("Data kendaraan berhasil diperbarui!")
	} else {
		fmt.Println("Data kendaraan tidak ditemukan!")
	}
}

func hapusKendaraan() {
	var plat string
	var idx, k int
	fmt.Println("\n--- HAPUS DATA KENDARAAN ---")
	fmt.Print("Masukkan Plat Nomor yang ingin dihapus: ")
	fmt.Scan(&plat)

	idx = cariSequential(plat)
	if idx != -1 {
		if idx == totalKendaraan-1 {
			totalKendaraan = totalKendaraan - 1
		} else {
			for k = idx; k < totalKendaraan-1; k = k + 1 {
				daftarKendaraan[k] = daftarKendaraan[k+1]
			}
			totalKendaraan = totalKendaraan - 1
		}
		fmt.Println("Data kendaraan (dan riwayat servisnya) berhasil dihapus!")
	} else {
		fmt.Println("Data kendaraan tidak ditemukan!")
	}
}

func tambahServis() {
	var plat, kategori, level string
	var idx, idxServis int
	fmt.Println("\n--- TAMBAH RIWAYAT SERVIS & DIAGNOSIS ---")
	fmt.Print("Masukkan Plat Nomor Kendaraan: ")
	fmt.Scan(&plat)

	idx = cariSequential(plat)
	if idx != -1 {
		idxServis = daftarKendaraan[idx].JumlahServis
		if idxServis < 10 {
			fmt.Print("Masukkan Tanggal (DD-MM-YYYY): ")
			fmt.Scan(&daftarKendaraan[idx].Riwayat[idxServis].Tanggal)

			fmt.Print("Masukkan Kategori Kerusakan (interior/eksterior): ")
			fmt.Scan(&kategori)
			fmt.Print("Masukkan Level Kerusakan (low/mid/high): ")
			fmt.Scan(&level)

			daftarKendaraan[idx].Riwayat[idxServis].JenisRusak = kategori + " " + level

			fmt.Println("\n--- Hasil Diagnosis ---")
			if kategori == "interior" {
				if level == "low" {
					fmt.Println("Saran Servis: Pembersihan noda jok dan karpet.")
				} else if level == "mid" {
					fmt.Println("Saran Servis: Perbaikan door trim atau switch jendela.")
				} else if level == "high" {
					fmt.Println("Saran Servis: Reset kelistrikan/dashboard atau re-trim total.")
				}
			} else if kategori == "eksterior" {
				if level == "low" {
					fmt.Println("Saran Servis: Poles baret halus pada body.")
				} else if level == "mid" {
					fmt.Println("Saran Servis: Ketok magic skala ringan atau cat ulang per panel.")
				} else if level == "high" {
					fmt.Println("Saran Servis: Perbaikan struktur bumper atau bodi penyok parah.")
				}
			}

			daftarKendaraan[idx].JumlahServis = daftarKendaraan[idx].JumlahServis + 1
			fmt.Println("Riwayat servis berhasil dicatat ke sistem!")
		} else {
			fmt.Println("Riwayat servis kendaraan ini sudah penuh! Maksimal 10 servis.")
		}
	} else {
		fmt.Println("Data kendaraan tidak ditemukan!")
	}
}

func tampilkanSemua() {
	var i, j, pilSort int

	fmt.Println("\n--- URUTKAN KENDARAAN ---")
	fmt.Println("1. Ascending (Berdasarkan Tahun - Selection Sort)")
	fmt.Println("2. Descending (Berdasarkan Tahun - Insertion Sort)")
	fmt.Print("Pilih metode pengurutan (1 atau 2): ")
	fmt.Scan(&pilSort)

	if pilSort == 1 {
		urutTahunSelection()
		fmt.Println("\nBerhasil diurutkan secara Ascending (Selection Sort)")
	} else if pilSort == 2 {
		urutTahunInsertion()
		fmt.Println("\nBerhasil diurutkan secara Descending (Insertion Sort)")
	} else {
		fmt.Println("\nPilihan tidak valid, data akan ditampilkan sesuai urutan input.")
	}

	fmt.Println("\n--- DAFTAR SEMUA KENDARAAN ---")
	if totalKendaraan == 0 {
		fmt.Println("Belum ada data kendaraan terdaftar.")
		return
	}

	for i = 0; i < totalKendaraan; i = i + 1 {
		fmt.Printf("%d. Plat: %s | Pemilik: %s | Tahun: %d\n", i+1, daftarKendaraan[i].PlatNomor, daftarKendaraan[i].Pemilik, daftarKendaraan[i].Tahun)
		if daftarKendaraan[i].JumlahServis != 0 {
			fmt.Println("   Riwayat Servis:")
			for j = 0; j < daftarKendaraan[i].JumlahServis; j = j + 1 {
				fmt.Printf("   - %s | %s\n", daftarKendaraan[i].Riwayat[j].Tanggal, daftarKendaraan[i].Riwayat[j].JenisRusak)
			}
		} else {
			fmt.Println("   (Belum memiliki riwayat servis)")
		}
		fmt.Println("-----------------------------------------")
	}
}

// FUNGSI OPSI CARI (BERDASARKAN PLAT / TAHUN)
func cariKendaraan() {
	var pilihanCari, j, idx, targetTahun, kiri, kanan, hitung int
	var targetPlat string

	fmt.Println("\n--- MENU CARI DATA KENDARAAN ---")
	fmt.Println("1. Cari Berdasarkan Plat Nomor (Menampilkan 1 Data)")
	fmt.Println("2. Cari Berdasarkan Tahun (Menampilkan Banyak Data)")
	fmt.Print("Pilih opsi pencarian (1-2): ")
	fmt.Scan(&pilihanCari)

	if pilihanCari == 1 {
		// OPSI 1: CARI BERDASARKAN PLAT
		fmt.Print("Masukkan Plat Nomor: ")
		fmt.Scan(&targetPlat)

		// WAJIB: Urutkan dulu berdasarkan plat nomor agar binary search valid
		urutPlatNomor()
		idx = cariBinary(targetPlat)

		if idx != -1 {
			fmt.Println("\n[Data Ditemukan - Hasil Binary Search]")
			fmt.Println("Plat Nomor :", daftarKendaraan[idx].PlatNomor)
			fmt.Println("Pemilik    :", daftarKendaraan[idx].Pemilik)
			fmt.Println("Tahun      :", daftarKendaraan[idx].Tahun)
			if daftarKendaraan[idx].JumlahServis != 0 {
				fmt.Println("Riwayat Servis:")
				for j = 0; j < daftarKendaraan[idx].JumlahServis; j = j + 1 {
					fmt.Printf("- %s | %s\n", daftarKendaraan[idx].Riwayat[j].Tanggal, daftarKendaraan[idx].Riwayat[j].JenisRusak)
				}
			}
		} else {
			fmt.Println("Data kendaraan dengan plat tersebut tidak ditemukan.")
		}

	} else if pilihanCari == 2 {
		// OPSI 2: CARI BERDASARKAN TAHUN
		fmt.Print("Masukkan Tahun Kendaraan: ")
		fmt.Scan(&targetTahun)

		// WAJIB: Urutkan dulu secara Ascending berdasarkan Tahun agar binary search valid
		urutTahunSelection()
		idx = cariBinaryTahun(targetTahun)

		if idx != -1 {
			fmt.Printf("\n--- DATA KENDARAAN TAHUN %d ---\n", targetTahun)
			hitung = 0

			// 1. Sisir ke arah kiri (indeks yang lebih kecil) untuk mencari tahun yang sama
			kiri = idx
			for kiri >= 0 && daftarKendaraan[kiri].Tahun == targetTahun {
				kiri = kiri - 1
			}
			// Kembalikan posisi awal karena loop berhenti saat tahun tidak sama
			kiri = kiri + 1

			// 2. Cetak semua data dari batas kiri ke kanan yang bertahun sama
			kanan = kiri
			for kanan < totalKendaraan && daftarKendaraan[kanan].Tahun == targetTahun {
				hitung = hitung + 1
				fmt.Printf("%d. Plat: %s | Pemilik: %s | Tahun: %d\n", hitung, daftarKendaraan[kanan].PlatNomor, daftarKendaraan[kanan].Pemilik, daftarKendaraan[kanan].Tahun)

				if daftarKendaraan[kanan].JumlahServis != 0 {
					fmt.Println("   Riwayat Servis:")
					for j = 0; j < daftarKendaraan[kanan].JumlahServis; j = j + 1 {
						fmt.Printf("   - %s | %s\n", daftarKendaraan[kanan].Riwayat[j].Tanggal, daftarKendaraan[kanan].Riwayat[j].JenisRusak)
					}
				}
				fmt.Println("-----------------------------------------")
				kanan = kanan + 1
			}
		} else {
			fmt.Printf("Tidak ada kendaraan buatan tahun %d.\n", targetTahun)
		}

	} else {
		fmt.Println("Pilihan menu pencarian tidak valid!")
	}
}

func main() {
	var pilihan int
	pilihan = 0

	for pilihan != 7 {
		if pilihan == 0 {
			fmt.Println("\n===== SISTEM MANAJEMEN SERVIS =====")
			fmt.Println("1. Tambah Data Kendaraan")
			fmt.Println("2. Tampilkan Semua Kendaraan")
			fmt.Println("3. Cari Data Kendaraan")
			fmt.Println("4. Ubah Data Kendaraan")
			fmt.Println("5. Hapus Data Kendaraan")
			fmt.Println("6. Tambah Riwayat Servis (Diagnosis)")
			fmt.Println("7. Keluar")
			fmt.Print("Pilih menu (1-7): ")
			fmt.Scan(&pilihan)
		} else if pilihan == 1 {
			tambahKendaraan()
			pilihan = 0
		} else if pilihan == 2 {
			tampilkanSemua()
			pilihan = 0
		} else if pilihan == 3 {
			cariKendaraan()
			pilihan = 0
		} else if pilihan == 4 {
			ubahKendaraan()
			pilihan = 0
		} else if pilihan == 5 {
			hapusKendaraan()
			pilihan = 0
		} else if pilihan == 6 {
			tambahServis()
			pilihan = 0
		} else if pilihan == 7 {
			fmt.Println("Keluar dari program. Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
			pilihan = 0
		}
	}
}
