package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	/* buatlah kode utama yang membaca baris pertama (n dan k). kemudian data diisioleh prosedur isiArray(n), dan pencarian oleh fungsi posisi(n,k), dan setelah itu output dicetak.*/
	var n, k, hasil int
	fmt.Scan(&n, &k)
	isiArray(n)
	//hasil = posisi(n, k)
	hasil = posisiA(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
func isiArray(n int) {
	/* I.S. terdefinisi integern, dan sejumlah n data sudah siap pada piranti masukan.
	   F.S. Array data berisi n (<=NMAX) bilangan */
	var i int = 0
	for i < n && i < NMAX {
		fmt.Scan(&data[i])
		i++
	}
}
func posisi(n, k int) int {
	/* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dariposisi 0. Jika tidak ada kembalikan -1 */
	var i int
	i = 0
	for i < n && data[i] != k {
		i++
	}
	if i < n {
		return i
	} else {
		return -1
	}
}

func posisiA(n, k int) int {
	/* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dari
	   posisi 0. Jika tidak ada kembalikan -1 */
	var kiri, kanan, tengah int
	var nilai int

	kiri = 0
	kanan = n - 1
	nilai = -1

	for kiri <= kanan && nilai == -1 {
		tengah = (kiri + kanan) / 2
		if k < data[tengah] {
			kanan = tengah - 1
		} else if k > data[tengah] {
			kiri = tengah + 1
		} else {
			nilai = tengah
		}
	}
	return nilai
}
