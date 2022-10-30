package main

import (
	"fmt"
)

func minimum(a []int) (index int, min int) {
	minval := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < minval {
			minval = a[i]
			index = i
		}

	}

	return index, minval
}

func maximum(a []int) (index int, min int) {
	minval := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > minval {
			minval = a[i]
			index = i
		}

	}

	return index, minval
}

// function ini ngereturn index yang nilainya ==10000
func sepuluhrb(a []int) (b []int) {

	for i := 0; i < len(a); i++ {
		if a[i] == 10000 {
			b = append(b, i)

		}

	}

	return b
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Inisialisasi data struct untuk harga barang
type barang struct {
	nama_barang  string
	harga_barang int
}

func main() {
	var list_barang [10]barang
	list_nama_barang := [10]string{"Benih Lele", "Pakan lele cap menara", "Probiotik A", "Probiotik Nila B", "Pakan Nila", "Benih Nila1", "Cupang", "Benih Nila 2", "Benih Cupang", "Probiotik B"}
	list_harga_barang := [10]int{50000, 25000, 75000, 10000, 20000, 20000, 5000, 30000, 10000, 10000}
	for i := 0; i < 10; i++ {
		list_barang[i].nama_barang = list_nama_barang[i]
		list_barang[i].harga_barang = list_harga_barang[i]
		//fmt.Print(list_barang[i].nama_barang, "-", list_barang[i].harga_barang, "\n")
	}
	//Algoritma agar petani dapat membeli barang paling banyak dengan jumlah uang
	//100000
	//Akan dicari barang dengan harga paling kecil,
	//lalu di sum dan mencari harga yang lebih besar
	//sampai harga total menjadi 100000
	index := 0
	sum := 0
	list_harga_barang_test := list_harga_barang
	var list_harga_beli [10]int
	var list_barang_beli [10]string
	for i := 0; i < 10; i++ {
		min_harga := list_harga_barang_test[0]
		for j := 0; j < len(list_harga_barang_test); j++ {
			if list_harga_barang_test[j] < min_harga {
				min_harga = list_harga_barang_test[j]
				index = j

			}

		}
		//menjadikan nilai minimum menjadi nilai maksimum dari array
		list_harga_barang_test[index] = 800000
		fmt.Println(list_harga_barang_test)
		list_harga_beli[i] = min_harga
		list_barang_beli[i] = list_nama_barang[index]
		sum = sum + list_harga_beli[i]
		//fmt.Println(sum)
		if sum == 100000 {
			break
		}
		index = 0
	}
	//fmt.Print(list_harga_beli)
	//fmt.Print(list_barang_beli)
	fmt.Println("Total produk dengan harga dibawah Rp 100.000 : ")
	fmt.Printf("Harga total %d\n", sum)
	for i := 0; i < 7; i++ {
		fmt.Printf(list_barang_beli[i]+" - %d\n", list_harga_beli[i])
	}
	fmt.Println("-------------------------------------------")
	//menjalankan function untuk menampilkan barang dengan harga 100000
	fmt.Println("Daftar produk dengan harga Rp 10.000")
	idx_array10rb := sepuluhrb(list_harga_barang[:])
	for i := 0; i < len(idx_array10rb); i++ {
		fmt.Printf("%s - %d\n", list_nama_barang[idx_array10rb[i]], list_harga_barang[idx_array10rb[i]])
	}
	fmt.Println("-------------------------------------------")
	//function u/ harga paling murah dan paling mahal
	idx_murah, harga_murah := minimum(list_harga_barang[:])
	idx_mahal, harga_mahal := maximum(list_harga_barang[:])
	fmt.Printf("Daftar produk termurah: %s Rp %d\n", list_nama_barang[idx_murah], harga_murah)
	fmt.Printf("Daftar produk termahal: %s Rp %d\n", list_nama_barang[idx_mahal], harga_mahal)

}
