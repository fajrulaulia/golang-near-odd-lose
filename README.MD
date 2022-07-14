# Axiata Digital Services Indonesia Technical Test(Software Engineer Backend)

untuk memenuhi soal no 1, menggunakan golang dan unit test


# Kebutuhan
- go dengan versi(Rekomendasi) v1.16.x keatas

# Setup
- masuk ke folder project yang sudah diekstrak 
ketikkan 
```bash
$ go mod tidy
```


# Menjalankan
disini saya menggunakan unit test dimana ada asserstion untuk menguji coba, bisa dilihat di `axiata_test.go`

- ketikkan untuk menjalankan test
    ```bash
    $ go test
    ```

- untuk lebih detail 
    ```bash
    $ go test -v
    ```
- untuk mengetahui persentase kode yang di kover dan output code yang di coverage
    ```bash
    $ go test -coverprofile cover.out && go tool cover -html=cover.out
    ```

# Coder
- Fajrul Aulia
    - E-Mail: devel.fajrul.aulia@gmail.com
    - github: https://github.com/fajrulaulia?tab=repositories