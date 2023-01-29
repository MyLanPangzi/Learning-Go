package main

import (
	"encoding/binary"
	"fmt"
	"github.com/MyLanPangzi/package_example/formatter"
	"github.com/MyLanPangzi/package_example/math"
	"github.com/learning-go-book/formatter"
	"github.com/learning-go-book/simpletax"
	simpletaxv2 "github.com/learning-go-book/simpletax/v2"
	"github.com/shopspring/decimal"
	"log"
	"os"

	//"github.com/learning-go-book/package_example/formatter"
	//"github.com/learning-go-book/package_example/math"
	crand "crypto/rand"
	"math/rand"
)

func main() {
	//overrideName()
	//workWithThirtyModule()
	//versions()
	upgrade()
}

func upgrade() {
	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	zip := os.Args[2]
	country := os.Args[3]
	percent, err := simpletaxv2.ForCountryPostalCode(country, zip)
	if err != nil {
		log.Fatal(err)
	}
	total := amount.Add(amount.Mul(percent))
	fmt.Println(total)
}

func versions() {
	amount, _ := decimal.NewFromString(os.Args[1])
	zip := os.Args[2]
	tax, err := simpletax.TaxForZip(zip)
	if err != nil {
		log.Fatal(err)
	}
	total := amount.Add(amount.Mul(tax)).Round(2)
	fmt.Println(total)
}

func workWithThirtyModule() {
	if len(os.Args) < 3 {
		fmt.Println("Need two parameters: amount and percent")
		os.Exit(1)
	}
	amount, err := decimal.NewFromString(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	percent, err := decimal.NewFromString(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	percent = percent.Div(decimal.NewFromInt(100))
	total := amount.Add(amount.Mul(percent)).Round(2)
	fmt.Println(formatter.Space(80, os.Args[1], os.Args[2], total.StringFixed(2)))
}

func overrideName() {
	num := math.Double(2)
	output := print.Formatter(num)
	fmt.Println(output)

	r := seedRand()
	fmt.Println(r.Intn(100))
}
func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}
