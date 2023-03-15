package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var friends = []Person{
	Person{
		nama:      "Denny",
		alamat:    "Karanganyar",
		pekerjaan: "BackEnd Developer",
		alasan:    "Mau belajar Golang",
	},
	Person{
		nama:      "Winalton",
		alamat:    "Tangerang",
		pekerjaan: "IT Consultant",
		alasan:    "Mau beli Iphone",
	},
	Person{
		nama:      "Jovi",
		alamat:    "Medit",
		pekerjaan: "FrontEnd Developer",
		alasan:    "Penasaran",
	},
	Person{
		nama:      "Jessie",
		alamat:    "Rawa Buaya",
		pekerjaan: "UI/UX Designer",
		alasan:    "Explore lebih lanjut",
	},
	Person{
		nama:      "Steven",
		alamat:    "Rawa Buaya",
		pekerjaan: "Software Engineer",
		alasan:    "Mau beli Macbook",
	},
	Person{
		nama:      "Stinky",
		alamat:    "Telukgong",
		pekerjaan: "Psikiater",
		alasan:    "Mau belajar hal baru",
	},
	Person{
		nama:      "Verry",
		alamat:    "Mangga Besar",
		pekerjaan: "Chef",
		alasan:    "Penasaran dengan Go",
	},
	Person{
		nama:      "Metta",
		alamat:    "Kejayaan",
		pekerjaan: "Marketing",
		alasan:    "Mau lulus",
	},
	Person{
		nama:      "Caryn",
		alamat:    "Pangeran Jayakarta",
		pekerjaan: "System Analyst",
		alasan:    "Mau beli Macbook",
	},
	Person{
		nama:      "Erik",
		alamat:    "Mangga Besar",
		pekerjaan: "Software Engineer",
		alasan:    "Menambah ilmu",
	},
}

func (p Person) Print(id int) string {
	return fmt.Sprintf("ID : %d\nnama : %s\nalamat : %s\npekerjaan : %s\nalasan : %s", id, p.nama, p.alamat, p.pekerjaan, p.alasan)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Silahkan masukkan index data pada argumen pemanggilan program")
	} else {
		arg := os.Args[1]
		intArg, _ := strconv.Atoi(arg)

		if intArg >= 10 || intArg < 0 {
			fmt.Printf("Tidak terdapat data untuk indeks %s", arg)
		} else {
			fmt.Println(friends[intArg].Print(intArg))
		}
	}
}
