package main

import "fmt"

func main() {
    // Deklarasi dan inisialisasi variabel
    var a, b, sum, difference int

    // Input dua bilangan dari pengguna
    fmt.Print("Masukkan bilangan pertama: ")
    fmt.Scanln(&a)
    fmt.Print("Masukkan bilangan kedua: ")
    fmt.Scanln(&b)

    // Menghitung penjumlahan dan pengurangan
    sum = a + b
    difference = a - b

    // Menampilkan hasil penjumlahan dan pengurangan
    fmt.Printf("Penjumlahan %d dan %d adalah %d\n", a, b, sum)
    fmt.Printf("Pengurangan %d dan %d adalah %d\n", a, b, difference)
}