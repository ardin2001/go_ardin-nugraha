# Resume Materi Intro Echo Golang

## Third Party Library Golang
terdapat beberapa third party library pada golang yang dapat memudahkan user dalam mengembangkan aplikasi baik berbasis restfull api. Diantaranya ada echo, Go Kit, Logrus, gorm.io dan cobra.

## Apa itu echo ?
Echo adalah framework bahasa golang untuk pengembangan aplikasi web baik berbasis REST API. Echo memiliki performance yang tinggi, dapat diperluas, minimalis dan go web framework.

#### Keuntungan dari menggunakan echo framework
<ul>
<li>Router optimal</li>
<li>Rendering data</li>
<li>Data binding</li>
<li>Middleware</li>
<li>Scalable</li>
</ul>

#### Echo adalah sebuah kerangka kerja minimalis
Tidak ada driver database atau ORM; Tidak ada perancah folder yang disediakan, tentukan struktur Anda sendiri!: Mesin template: https://echo.labstack.com/guide/templates

## Setting up echo
Diawali dengan membuat directori, kemudian go mod ini dan yang terakhir go get github.com/labstack/echo/v4. Kemudan melakukan routing terhadap user yang ingin melakukan request agar mendapatkan response. Response ini dapat disesuaikan.