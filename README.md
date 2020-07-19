# Covid-API
Corona API for Palembang City (Fake Data)

Base URL    : `http://classico.id:1103/`

Example     : `http://classico.id:1103/covid/`

### Routes
  * `GET` `covid/` : `Mengakses seluruh data Covid-19 Palembang`
  * `GET` `covid/{id}` : `Mengakses data Covid-19 Palembang dengan id`
  * `GET` `info/covid/` : `Mengakses seluruh data Info Covid-19`
  * `GET` `covid/page/{id}` : `Mengakses data Covid-19 Palembang dengan Pagination`
  * `POST` `covid/` : `Mengakses seluruh data Covid-19 Palembang`
  * `PUT` `covid/{id}` : `Mengedit data Covid-19 Palembang`
  * `DELETE` `covid/{id}` : `Menghapus data Covid-19 Palembang`

### Tutorial API
  * Clone Repository ini dengan cara `git clone https://github.com/calvinfm/Covid-GO-API.git`
  * Atau bisa dengan mendownload menjadi `.zip` file
  * Simpan / Extract Folder kedalam `C:\Users\{User}\go`
  * Buka Aplikasi `JetBrains GOLAND` atau `Visual Studio Code` 
  * Buka terminal dan mengarahkan ke folder tempat API disimpan
  * Run Script dengan command `go run Main.go`
  * Server akan berjalan lalu akses di `127.0.0.1:{port}`


### Tools
  * IDE
    * [Goland](https://www.jetbrains.com/go/)
    * [Visual Studio Code](https://code.visualstudio.com/download)
  * Database
    * [Postgresql](https://www.postgresql.org/download/)

