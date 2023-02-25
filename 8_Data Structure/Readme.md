# Data Structure

## Apa itu Array ?
Array adalah sebuah struktur data yang berisi sekelompok elemen, dapat berisis satu jenis variabel dengan ukuran alokasi tetap. Tipe data yang berbeda bisa ditangani sebagai elemen dalam array seperti number, string dan boolean.
#### Deklarasi Array
Untuk mendeklarasikan array anda butuh untuk spesifik nilai dari elemen itu. Diawali dengan kurung siku, dan diikuti oleh tipe data dari element array. Berikut adalah contoh deklarasi array :
- var <nama_variabel> [ukuran_array] <tipe_data>

#### Menambahkan dan Mengakses Elemen Array
```
package main

import "fmt"

func main() {
	var countries [2]string

	countries[0] = "Jakarta"
	countries[1] = "Bandung"

	fmt.Println(countries[0])
	fmt.Println(countries[1])
}
```
Pada contoh program diatas, untuk menambah elemen array bisa menggunakan <nama_variabel> [indeks_array] = <data array>, kemudian untuk mengakses elemen array bisa menggunakan <nama_variabel>[indeks_array].

#### Inisialisasi Dengan Array Literal
```
package main

import "fmt"

func main() {
	odd_number := [5] int {1,2,3,4,5}

	fmt.Println(odd_number)
	fmt.Println(odd_number[1])
}
```
Untuk menginisialisasi array literal bisa dengan menggunakan <nama_variabel> := [jumlah_array} <tipe_data> {isi _data_array}

#### Menyimpulkan Panjang Dari Array
```
package main

import (
  "fmt"
  "reflect"
)

func main() {
	odd_number := [...] int {1,2,3,4,5}

	fmt.Println(reflect.Valueof(odd_number).Kind())
	fmt.Println(len(odd_number))
}
```
Terdapat beberapa package yang dapat membantu kita untuk mengerjakan sebuah source code, diantaranya package reflect. Terdapat beberapa method dari package reflect yang bisa digunakan seperti Valueof, Typeof dan lain-lain untuk mengetahui jenis data dari variabel dan method len yang digunakan untuk mengetahui panjang dari array.

## Apa itu Slice ?
Slice adalah salah satu struktur data yang mengandung elemen grup, dapat berisi satu jenis variabel (seperti array) tetapi mempunyai alokasi ukuran yang dinamis.
#### Cara membuat slice dari array
```
var primes = [5] int{1,2,3,4}
var part_primes []int = primes[1:4]
```
#### Slice declaration
var even_numbers []int
var add_number = []int{1,2,3,4}
numbers := [] int{1,2,3,4,5}
var primes = make([]int,5,10)

## Map
Map adalah salah satu struktur data yang menyimpan data dalam bentuk pasangan kunci dan nilai dan kuncinya unik.
```
package main

import "fmt"

func main() {
	var salary = map[string]int{}
	fmt.Println(salary)

	var salary_a = map[string]int{"ardin": 175, "nugraha": 174}
	fmt.Println(salary_a)

	salary_b := map[string]int{}
	fmt.Println(salary_b)

	var salary_c = make(map[string]int)
	salary_c["doe"] = 4777
	fmt.Println(salary_c)
}
```
Pada contoh kode diatas terdapat beberapa cara implementasi untuk mendeklarasikan map, mulai dari long declaration, long declaration with value, short declaration dan using make.

## Funcion
Sebuah function adalah bagian kode yang dipanggil dengan nama. Fungsi adalah cara mudah untuk membagi kode anda menjadi blok-blok yang berguna. Function berguna untuk clean code, rapi dan modular.
#### Deklarasi Function
Berikut beberapa cara untuk mendeklarasikan function, diantaranya :
<ul>
<li>func name_function() {statement}</li>
<li>funct name_function() type_return {statement}</li>
<li>funct name_function(parameter) type_return{statement}</li>
</ul>
