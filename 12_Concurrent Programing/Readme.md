# Resume Materi Concurrency in Golang

## Sequential Program
Dalam sequential rpogram, sebelum menjalankan sebuah task baru, harus selesai task sebelumnya. Jadi sequential akan memproses beberapa taks secara bergantian, seperti dimulai dari video,website kemudian image. Sehingga akan memakan cukup waktu banyak apabila terdapat banyak proses atau task

## Parallel Program
Dalam parallel program, multiple task bisa mengeksekusi secara serentak. Tetapi parallel program membutuhkan mesin mult-core. Seperti contoh apabila terdapat 3 task seperti mereload vidio, web dan image, makan task akan dijalankan secara bebarengan sehingga membutuhkan waktu yang lebih sedikit untuk menyelesaikan program.

## Concurrent Program
Dalam concurrent program, multi task bisa di eksekusi secara independen dan bisa muncul serentak. Concurrent program tidak membutuhkan mesin multi core

## Goroutines
Goroutines adalah fungsi atau method yang berjalan bersamaan (independen) dengan function atau method lain.
Biaya untuk membuat sebuah Goroutine kecil jika dibandingkan dengan sebuah thread. Sebuah thread adalah proses yang ringan, atau dengan kata lain, sebuah thread adalah sebuah unit yang mana mengeksekusi code dibawah program.

#### Contoh Penggunaan Goroutine
```
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello world")
}

func main() {
	go hello()
	fmt.Println("finis")
	time.Sleep(1 * time.Second)
}

```

Untuk membuat goroutine pada function atau method hanya dengan menambahkan go pada depan nama variabel, kemudian untuk dapat melihat hasilnya bisa dengan menambahkan time.sleep agar program menunggu proses selesai function goroutine. Dalam sebuah program bisa terdapat lebih dari satu goroutine.

## GOMAXPROCS
Gomaxprocs digunakan untuk mengontrol jumlah dari operating systemthread yang tersedia untuk mengeksekusi Goroutine untuk sebuah go program tertentu. Berikut adalah command untuk menjalankan gomaxprocs ( time GOMAXPROCS=1 go run main.go)

## Channel & Select
#### Channel
Channel adalah sebuah obejk komunikasi digunakan goroutine agar bisa berkomunikasi dengan yang lain.
```
package main

import "fmt"
import "runtime"

func main() {
    runtime.GOMAXPROCS(2)

    var messages = make(chan string)

    var sayHelloTo = func(who string) {
        var data = fmt.Sprintf("hello %s", who)
        messages <- data
    }

    go sayHelloTo("john wick")
    go sayHelloTo("ethan hunt")
    go sayHelloTo("jason bourne")

    var message1 = <-messages
    fmt.Println(message1)

    var message2 = <-messages
    fmt.Println(message2)

    var message3 = <-messages
    fmt.Println(message3)
}
```

#### Unbuffred and Buffered channels
Unbuffered channels memblokir goroutine pengirim hingga ada penerima terkait yang siap menerima nilai yang dikirim. Ini berarti bahwa data dijamin akan diterima sesuai urutan pengirimannya, dan sinkronisasi dibangun ke dalam channel. Sedangkan Buffered channels, di sisi lain, dapat menampung nilai dalam jumlah terbatas (ditentukan oleh ukuran buffer), dan hanya akan memblokir goroutine pengirim saat buffer penuh. Ini dapat memungkinkan beberapa konkurensi tambahan, tetapi membutuhkan pertimbangan yang cermat untuk menghindari kebuntuan dan masalah sinkronisasi lainnya.

## Select
Select membuat goroutine lebih mudah untuk mengendalikan komunikasi data sampai satu atau banyak channels

## Race Condition
Race condition dimana 2 thread mengakses memory di waktu yang sama, satu dari yang menulis. Race condition terjadi karena tidak tersinkronisasi akses ke berbagi memory. untuk mengatasi hal tersebut bisa menggunakan WaitGroup, Channels Blocking dan Mutex.

## Blocking dengan Waithgroups
Cara paling mudah untuk menyelesaikan data race, adalah memblokir akses baca hingga operasi tulis selesai.

## Channel Blocking
Ini memungkina goroutine kita untuk melakukan sinkronisasi satu sama lain untuk sesaat.

## Mutex
Jika kami tidak peduli dengan urutan baca dan tulis, kami hanya meminta agar urutan tersebut tidak terjadi secara bersamaan.

Jika ini terdengar seperti kasus penggunaan Anda, maka kami harus mempertimbangkan untuk menggunakan mutex.
