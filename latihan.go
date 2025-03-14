package main
import "fmt"

func jumlahDigit(n int) int {
    var hasil, i int
    hasil = 0
    if n > 0{
        i = n%10
        hasil += i
        hasil += jumlahDigit(n/10)
    }
    return hasil
}

func main(){
    var masukan int
    fmt.Scan(&masukan)
    fmt.Println(jumlahDigit(masukan))
}

