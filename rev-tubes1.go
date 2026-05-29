package main

import "fmt"

type Ide struct {
	ID       int
	Judul    string
	Kategori string
	Upvote   int
}

var daftarIde []Ide

func garis() {
	fmt.Println("==========================================")
}

func menu() {
	garis()
	fmt.Println("     APLIKASI MANAJEMEN IDE STARTUP")
	garis()

	fmt.Println("1. Tambah Ide")
	fmt.Println("2. Tampilkan Ide")
	fmt.Println("3. Upvote Ide")
	fmt.Println("4. Sequential Search")
	fmt.Println("5. Binary Search")
	fmt.Println("6. Selection Sort")
	fmt.Println("7. Insertion Sort")
	fmt.Println("0. Keluar")

	garis()
	fmt.Print("Pilih menu: ")
}

// ================= TAMBAH IDE =================

func tambahIde() {

	var judul, kategori string

	fmt.Print("Masukkan Judul ide     : ")
	fmt.Scanln(&judul)

	fmt.Print("Masukkan Kategori ide  : ")
	fmt.Scanln(&kategori)

	if judul == "" || kategori == "" {

		fmt.Println("Input tidak boleh kosong.")
		return
	}

	data := Ide{
		ID:       len(daftarIde) + 1,
		Judul:    judul,
		Kategori: kategori,
		Upvote:   0,
	}

	daftarIde = append(daftarIde, data)

	fmt.Println("Ide berhasil ditambahkan.")
}

// ================= TAMPIL =================

func tampilkanIde() {

	if len(daftarIde) == 0 {

		fmt.Println("Belum ada data ide.")
		return
	}

	garis()

	fmt.Printf("%-5s %-20s %-15s %-10s\n",
		"ID", "Judul", "Kategori", "Upvote")

	garis()

	for _, ide := range daftarIde {

		fmt.Printf("%-5d %-20s %-15s %-10d\n",
			ide.ID,
			ide.Judul,
			ide.Kategori,
			ide.Upvote)
	}

	garis()
}

// ================= UPVOTE =================

func upvoteIde() {

	var id int

	fmt.Print("Masukkan ID ide: ")
	fmt.Scanln(&id)

	for i := range daftarIde {

		if daftarIde[i].ID == id {

			daftarIde[i].Upvote++

			fmt.Println("Upvote berhasil.")
			return
		}
	}

	fmt.Println("ID tidak ditemukan.")
}

// ================= SEARCH =================

// Sequential Search
func sequentialSearch() {

	var cari string
	found := false

	fmt.Print("Masukkan judul ide: ")
	fmt.Scanln(&cari)

	for _, ide := range daftarIde {

		if ide.Judul == cari {

			fmt.Println("\nData ditemukan:")
			fmt.Println("ID  :", ide.ID)
			fmt.Println("Judul :", ide.Judul)
			fmt.Println("Kategori :", ide.Kategori)
			fmt.Println("Upvote :", ide.Upvote )


			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan.")
	}
}

// Binary Search
func binarySearch() {

	sortJudul()

	var cari string

	fmt.Print("Cari judul ide: ")
	fmt.Scanln(&cari)

	left := 0
	right := len(daftarIde) - 1

	for left <= right {

		mid := (left + right) / 2

		if daftarIde[mid].Judul == cari {

			fmt.Println("\nData ditemukan:")
			fmt.Println("ID  :", daftarIde[mid].ID)
			fmt.Println("Judul :", daftarIde[mid].Judul)
			fmt.Println("Kategori :", daftarIde[mid].Kategori)
			fmt.Println("Upvote :", daftarIde[mid].Upvote )
			
			return

		} else if daftarIde[mid].Judul < cari {

			left = mid + 1

		} else {

			right = mid - 1
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

// ================= SELECTION SORT =================

func selectionSort() {

	n := len(daftarIde)

	for i := 0; i < n-1; i++ {

		max := i

		for j := i + 1; j < n; j++ {

			if daftarIde[j].Upvote > daftarIde[max].Upvote {
				max = j
			}
		}

		daftarIde[i], daftarIde[max] =
			daftarIde[max], daftarIde[i]
	}

	fmt.Println("Data berhasil diurutkan menggunakan Selection Sort.")
}

// ================= INSERTION SORT =================
func insertionSort() {

	for i := 1; i < len(daftarIde); i++ {

		temp := daftarIde[i]
		j := i - 1

		for j >= 0 && daftarIde[j].Upvote < temp.Upvote {

			daftarIde[j+1] = daftarIde[j]
			j--
		}

		daftarIde[j+1] = temp
	}

	fmt.Println("Data berhasil diurutkan menggunakan Insertion Sort.")
}

// ================= SORT JUDUL =================
func sortJudul() {

	for i := 0; i < len(daftarIde)-1; i++ {

		for j := i + 1; j < len(daftarIde); j++ {

			if daftarIde[i].Judul > daftarIde[j].Judul {

				daftarIde[i], daftarIde[j] =
					daftarIde[j], daftarIde[i]
			}
		}
	}
}

// ================= MAIN =================

func main() {

	daftarIde = append(daftarIde,
		Ide{1, "EcoTrack", "Lingkungan", 10},
		Ide{2, "FoodHub", "Kuliner", 7},
		Ide{3, "HealthSync", "Kesehatan", 15},
	)

	var pilih int

	for {

		menu()
		fmt.Scanln(&pilih)

		switch pilih {

		case 1:
			tambahIde()

		case 2:
			tampilkanIde()

		case 3:
			upvoteIde()

		case 4:
			sequentialSearch()

		case 5:
			binarySearch()

		case 6:
			selectionSort()
			tampilkanIde()

		case 7:
			insertionSort()
			tampilkanIde()

		case 0:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Menu tidak tersedia.")
		}

		fmt.Println()
	}
}