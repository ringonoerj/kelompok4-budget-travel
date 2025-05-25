package main

import (
	"fmt"
	"strings"
)

// ===============================
// Tipe Bentukan dan Variabel Global
// ===============================

const MAX = 100

// Struct untuk menyimpan data pengeluaran
type Pengeluaran struct {
	Kategori   string
	Jumlah     int
	Keterangan string
}

// Array utama dan variabel global yang diperbolehkan
var dataPengeluaran [MAX]Pengeluaran
var totalData int = 0
var totalBudget int // Diinput oleh user saat program dimulai

// ===============================
// Subprogram: CRUD Pengeluaran
// ===============================

// Menambahkan pengeluaran ke array
func tambahPengeluaran(kategori string, jumlah int, keterangan string) {
	if totalData < MAX {
		dataPengeluaran[totalData] = Pengeluaran{kategori, jumlah, keterangan}
		totalData++
	} else {
		fmt.Println("Data pengeluaran sudah penuh.")
	}
}

// Mengedit data pengeluaran di indeks tertentu
func editPengeluaran(index int, jumlah int, keterangan string) {
	if index >= 0 && index < totalData {
		dataPengeluaran[index].Jumlah = jumlah
		dataPengeluaran[index].Keterangan = keterangan
	} else {
		fmt.Println("Index tidak valid.")
	}
}

// Menghapus data pengeluaran dari array
func hapusPengeluaran(index int) {
	if index >= 0 && index < totalData {
		for i := index; i < totalData-1; i++ {
			dataPengeluaran[i] = dataPengeluaran[i+1]
		}
		totalData--
	} else {
		fmt.Println("Index tidak valid.")
	}
}

// ===============================
// Subprogram: Tampilan dan Laporan
// ===============================

// Menampilkan seluruh data pengeluaran
func tampilkanPengeluaran() {
	fmt.Println("Daftar Pengeluaran:")
	for i := 0; i < totalData; i++ {
		fmt.Printf("%d. [%s] Rp%d - %s\n", i+1, dataPengeluaran[i].Kategori, dataPengeluaran[i].Jumlah, dataPengeluaran[i].Keterangan)
	}
}

// Menghitung total pengeluaran dari seluruh data
func totalPengeluaran() int {
	total := 0
	for i := 0; i < totalData; i++ {
		total += dataPengeluaran[i].Jumlah
	}
	return total
}

// Memberikan saran penghematan berdasarkan selisih budget
func saranPenghematan() {
	total := totalPengeluaran()
	fmt.Printf("Total Pengeluaran: Rp%d\n", total)
	fmt.Printf("Total Budget: Rp%d\n", totalBudget)
	if total > totalBudget {
		fmt.Printf("Pengeluaran melebihi budget sebesar Rp%d.\n", total-totalBudget)
		fmt.Println("Saran: Kurangi pengeluaran yang tidak terlalu penting seperti hiburan, makanan & oleh-oleh yang mahal.")
	} else {
		fmt.Printf("Masih dalam anggaran. Sisa budget: Rp%d\n", totalBudget-total)
	}
}

// Menampilkan laporan per kategori
func laporanKategori() {
	kategoriMap := map[string]int{}
	for i := 0; i < totalData; i++ {
		kategoriMap[dataPengeluaran[i].Kategori] += dataPengeluaran[i].Jumlah
	}
	fmt.Println("Laporan per Kategori:")
	for k, v := range kategoriMap {
		fmt.Printf("- %s: Rp%d\n", k, v)
	}
}

// ===============================
// Subprogram: Pencarian Data
// ===============================

// Pencarian data berdasarkan kategori (Sequential Search)
func sequentialSearch(kategori string) {
	fmt.Println("Hasil Pencarian Sequential:")
	found := false
	for i := 0; i < totalData; i++ {
		if strings.EqualFold(dataPengeluaran[i].Kategori, kategori) {
			fmt.Printf("%d. [%s] Rp%d - %s\n", i+1, dataPengeluaran[i].Kategori, dataPengeluaran[i].Jumlah, dataPengeluaran[i].Keterangan)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Pencarian data berdasarkan kategori (Binary Search) — data perlu diurutkan dulu
func binarySearchKategori(kategori string) int {
	insertionSortKategoriAsc() // Binary search perlu data terurut
	low := 0
	high := totalData - 1

	for low <= high {
		mid := (low + high) / 2
		if strings.EqualFold(dataPengeluaran[mid].Kategori, kategori) {
			return mid
		} else if strings.ToLower(dataPengeluaran[mid].Kategori) < strings.ToLower(kategori) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// ===============================
// Subprogram: Pengurutan Data
// ===============================

// Selection Sort berdasarkan jumlah (ascending)
func selectionSortJumlahAsc() {
	for i := 0; i < totalData-1; i++ {
		minIdx := i
		for j := i + 1; j < totalData; j++ {
			if dataPengeluaran[j].Jumlah < dataPengeluaran[minIdx].Jumlah {
				minIdx = j
			}
		}
		dataPengeluaran[i], dataPengeluaran[minIdx] = dataPengeluaran[minIdx], dataPengeluaran[i]
	}
}

// Insertion Sort berdasarkan kategori (ascending)
func insertionSortKategoriAsc() {
	for i := 1; i < totalData; i++ {
		temp := dataPengeluaran[i]
		j := i - 1
		for j >= 0 && strings.ToLower(dataPengeluaran[j].Kategori) > strings.ToLower(temp.Kategori) {
			dataPengeluaran[j+1] = dataPengeluaran[j]
			j--
		}
		dataPengeluaran[j+1] = temp
	}
}

// ===============================
// Fungsi Utama (Main Program)
// ===============================

func main() {
	fmt.Println("*******************************************************")
	fmt.Println("***            Aplikasi Budget Travel               ***")
	fmt.Println("***            Created by Kelompok 4                ***")
	fmt.Println("***     Muhammad Alfin Ramadhan (103042400021)      ***")
	fmt.Println("***        Ringo Noer Junaedy (103042400012)        ***")
	fmt.Println("***         Algoritma Pemrograman 2 - 2425          ***")
	fmt.Println("***                                                 ***")
	fmt.Println("*******************************************************")

	fmt.Print("Masukkan total budget untuk perjalanan Anda (Rp): ")
	fmt.Scan(&totalBudget)

	var pilihan int

	for {
		fmt.Println("\nMenu Utama:")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Tampilkan Pengeluaran")
		fmt.Println("3. Edit Pengeluaran")
		fmt.Println("4. Hapus Pengeluaran")
		fmt.Println("5. Cari (Sequential Search)")
		fmt.Println("6. Cari (Binary Search)")
		fmt.Println("7. Urutkan Jumlah (Selection Sort Ascending)")
		fmt.Println("8. Urutkan Kategori (Insertion Sort Ascending)")
		fmt.Println("9. Laporan & Saran Budget")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			for {
				fmt.Println("\nPilih Kategori:")
				fmt.Println("1. Transportasi")
				fmt.Println("2. Akomodasi")
				fmt.Println("3. Makanan")
				fmt.Println("4. Hiburan")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Print("Pilihan kategori: ")
				var kat int
				fmt.Scan(&kat)
				if kat == 5 {
					break
				}
				kategori := ""
				switch kat {
				case 1:
					kategori = "Transportasi"
				case 2:
					kategori = "Akomodasi"
				case 3:
					kategori = "Makanan"
				case 4:
					kategori = "Hiburan"
				default:
					fmt.Println("Kategori tidak valid.")
					continue
				}
				var jumlah int
				var keterangan string
				fmt.Print("Masukkan jumlah: ")
				fmt.Scan(&jumlah)
				fmt.Print("Masukkan keterangan: ")
				fmt.Scan(&keterangan)
				tambahPengeluaran(kategori, jumlah, keterangan)
				break
			}
		case 2:
			tampilkanPengeluaran()
		case 3:
			var idx, jumlah int
			var keterangan string
			fmt.Print("Masukkan index pengeluaran (0-based): ")
			fmt.Scan(&idx)
			fmt.Print("Jumlah baru: ")
			fmt.Scan(&jumlah)
			fmt.Print("Keterangan baru: ")
			fmt.Scan(&keterangan)
			editPengeluaran(idx, jumlah, keterangan)
		case 4:
			var idx int
			fmt.Print("Masukkan index pengeluaran (0-based): ")
			fmt.Scan(&idx)
			hapusPengeluaran(idx)
		case 5:
			var kategori string
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&kategori)
			sequentialSearch(kategori)
		case 6:
			var kategori string
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&kategori)
			idx := binarySearchKategori(kategori)
			if idx != -1 {
				fmt.Printf("Ditemukan: [%s] Rp%d - %s\n", dataPengeluaran[idx].Kategori, dataPengeluaran[idx].Jumlah, dataPengeluaran[idx].Keterangan)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		case 7:
			selectionSortJumlahAsc()
			fmt.Println("Data diurutkan berdasarkan jumlah (ascending).")
		case 8:
			insertionSortKategoriAsc()
			fmt.Println("Data diurutkan berdasarkan kategori (ascending).")
		case 9:
			tampilkanPengeluaran()
			laporanKategori()
			saranPenghematan()
		case 10:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
