# Materi String, Advance Function , Pointer , Method, Struct and Interface

## Bekerja Dengan String
Beberapa method yang dapat digunakan untuk membantu dalam menyelesaikan beberapa permasalahan dalam string, diantaranya :
<ol>
<li>len</li>
Len adalah salah satu method yang digunakan untuk meghitung panjang nilai pada sebuah string. Berikut adalah contoh implementasi code.

```
sentence := "Hello"
lenSentence := len(sentence)
fmt.Println(lenSentence)
```
<li>Compare String</li>
Compare string adalah cara untuk melakukan perbandingan antar string, dimana hasil dari compare tersebut adalah boolean.

```
str1 := "abc"
str2 := "abd"
fmt.Println(str1==str2)
```
<li>Contains</li>
Contains adalah salah satu method yang digunakan untuk mengecek apakah dalam string mengandung kata substring dimana retun valuenya adalah boolean. Berikut adalah contoh kode programnya.

```
str := "abc"
substr := "abd"
res := strings.Contains(str,substr)
fmt.Println(res)
```
<li>Substring</li>
Substring adalah cara untuk memotong bagian dari string, dimulai dengan [index awal:index akhir].
<li>Replace</li>
Replace adalah salah satu method untuk mereplace kata yang ada pada string. Berikut adalah contoh kode program.

```
s := "hello friends, my name is blablbalba"
sReplace := strings.Replace(s,"blablbalba","putra",-1)
fmt.Println(sReplace)
```
<li>Insert</li>
Insert adalah salah satu cara untuk menggabungkan teks kedalam string.
</ol>

## More Function
#### Variadic Funcion
Variadic Function digunakan untuk menampung banyak parameter untuk dijadikan array, caranya adalah dengan menambahkan ... sebelum tipe data.

```
func getMinMax(numbers ...*int) (min int, max int) {
}
func main() {
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
}
```

## Anonymous Function
Anonymous Function adalah sebuah function yang mana tidak berisi nama. Itu biasa digunakan ketika anda ingin membuat sebuah inline function.

## Closure
Closure adalah jenis spesial dari anonymous function yang mana deklarasi variabel di referensikan diluar dari function itu sendiri.

## Defer Function
Defer function adalah sebuah cara yang digunakan agar function yang diberi defer agar dieksekusi di akhir program.

## Pointer
Pointer adalah sebuah variabel toko penyimpanan alamat dari variabel lain. Pointer memiliki kekuatan untuk mengubah data yang mereka tunjuk.

#### Pointer Declaration
Untuk mendaklarasikan pointer yaitu menggunakan var nama_variabel * variabel_type. Berikut adalah contoh penggunaan pointer. (* ) operator ini digunakan untuk mendeklarasikan pointer, (&) operator ini digunakan untuk mengakses alamat variabel.

```
var name string = "John"
var nameAddress *string = &name
```

## Struct
Sebuah struct adalah tipe yang ditentukan pengguna yang berisi kumpulan dari bidang atau property dari function(method). Untuk mendalarasikan struct yaitu dengan menggunakan keyword (type struct_variable_name struct {field data_type}.
#### Method
Method adalah sebuah function yang melekat pada suatu tipe (bisa berupa struct atau tipe data lainnya). Untuk mendeklarasikan method yaitu dengan keyword (func (receive StructType) MethodName(parameter) (returnType) {}.

## Iterface
Interface adalah sebuah kumpulan tanda tangan  method yang dapat diimplementasikan objek. Karena interface mendeskripsikan perilaku objek. Untuk mendeklarasikan interface menggunakan keyword (type interface_name) interface {method_name return_type}.
