package main

import(
	"fmt"
	rqrcode "github.com/skip2/go-qrcode"	
	wqrcode "github.com/tuotoo/qrcode"
	"log"
	"os"
)

//CreateQRCode creates QRCode
func CreateQRCode(){
	err := rqrcode.WriteFile("https://www.gmail.com", rqrcode.Medium, 256, "qr.png")
	if err != nil{
		log.Fatal(err)
	}	
	fmt.Println("Qr code created!")
}

func ReadQRCode(){
	fi, err := os.Open("qr.png")
	if err != nil{
		log.Fatal(err)	
	}
	defer fi.Close()
	qrmatrix, err := wqrcode.Decode(fi)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(qrmatrix.Content)
}

func main(){
	CreateQRCode()
	ReadQRCode()
	fmt.Println("Success!")
}