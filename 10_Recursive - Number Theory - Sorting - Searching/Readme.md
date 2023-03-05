# Resume Materi Recursive, Number Theory, Sorting, Searching

## Outline
Pada materi kali ini, topik yang dibahas mengenai rekursif, number theory, searching dan sorting. Untuk implementasi mengenai source kode program menggunakan bahasa pemrograman java.

## Rekursif
Rekursif adalah keadaan dimana suatu fungsi memecahkan masalah dengan memanggil dirinya sendiri. Dimana terdapat kondiri yang akan menghentikan program bila telah mencapai base case.

#### Mengapa Kita Membutuhkan Rekursif ?
<ol>
<li>Dengan menggunakan rekursif, banyak masalah lebih mudah dipecahkan dan mempersingkat kode saat menggunakan pendekatan rekursif.</li>
<li>Pada dasarnya, baik menggunakan metode iteratif(for loop) dan menggunakan rekursif sama-sama menggunakan metode yang berulang.</li>
<li>Namun terkadang solusi iteratif untuk suatu masalah sangat sulit dipikirkan dan memerlukan teknik khusus.</li>
<li>Dengan solusi rekursif, mungkin lebih mudah untuk melihat dan merancang jalur penyelesaian.</li>
</ol>

#### Strategi Rekursif
Ada dua hal yang perlu dipikirkan saat menggunakan strategi rekursif.
<ul>
<li>Base Case</li>
Dalam fungsi rekursif, base case tidak melakukan sebuah pemanggilan rekursif lagi, melainkan bertujuan untuk memberhentikan proses rekursif.
<li>Recurrence Relations.</li>
Pada bagian ini biasanya berisi beberapa operasi dan pemanggilan fungsi rekursif kembali atau memanggil dirinya sendiri.
</ul>

#### Contoh Masalah : Faktorial
```
package main
import (
 "fmt"
)
func getFactorial(n int) int {
 if n == 0 || n == 1 {
  return 1
 }
 return n * getFactorial(n-1)
}
func main() {
 val := getFactorial(5)
 fmt.Println(val)
}
```
Pada contoh diatas adalah salah satu bentuk penyelesaian masalah faktorial menggunakan metode rekursif, dimana suatu program akan memanggil fungsi dirinya sendiri sampai dimana suatu keadaan mememuhi kondisi base case, dimana pada contoh diatas base case terletak pada return 1.

## Number Theory
Number Theory adalah salah satu cbang matematika yang mempelajari bilangan bulat. Ada banyak topik dalam seperti bilangan prima, faktor persekutuan terbesar, faktor pesekutuan terkecil, faktorial, dan bilangan prima, serta masih banyak lagi.

#### Berikut adalah contoh source code bilangan prima
```
package main
import (
   "fmt"
   "math"
)

func checkPrimeNumber(num int) {
   if num < 2 {
      fmt.Println("Number must be greater than 2.")
      return
   }
   sq_root := int(math.Sqrt(float64(num)))
   for i:=2; i<=sq_root; i++{
      if num % i == 0 {
         fmt.Println("Non Prime Number")
         return
      }
   }
   fmt.Println("Prime Number")
   return
}

func main(){
   checkPrimeNumber(0)
   checkPrimeNumber(2)
   checkPrimeNumber(13)
   checkPrimeNumber(152)
```


## Searching
Searching adalah proses untuk menemukan sebuah posisi nilai didalam sebuah list nilai. Beberapa contoh searching algorithm seperti linier search dan binary search.

#### Berikut ini adalah contoh dari linier search
```
package main

//import package yang ingin digunakan
import "fmt"

//penamaan function tergantung developer
func linearsort(list []int, key int) int {
	//proses perulangan dari data array items
	for _, data := range list {
		if data == key {
			fmt.Println("angka yang anda cari tersedia")
			return key
		}
	}
	fmt.Println("angka yang anda cari tidak tersedia")
	return key
}
func main() {
	//variabel items menampung data array
	list := []int{50, 27, 60, 70, 80, 95, 100, 120, 150}
	fmt.Println(linearsort(list, 80))
}

```

#### Berikut adalah contoh dari Binary Search
```
package main
import "fmt"

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high{
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		}else{
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}


func main(){
	items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
 ```
 
## Sorting
Sorting adalah proses menyusun data dengan urutan tertentu. Biasanya diurutkan berdasarkan nilai elemen. Terdapat beberapa sorting algorithm diantaranya selection sort, bubble sort, counting sort dan merge sort.

#### Berikut adalah contoh algoritma bubble sort
```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
 
func bubblesort(items []int) {
    var (
        n = len(items)
        sorted = false
    )
    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
                items[i+1], items[i] = items[i], items[i+1]
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }
}
```
