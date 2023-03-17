package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Name       string
	Address    string
	Occupation string
	Reason     string
}

func (b Biodata) ShowBiodata() {
	fmt.Printf(
		"Name:\t\t\t\t%s\nAddress:\t\t\t%s\nOccupation:\t\t\t%s\nReason Choose Golang Class:\t%s\n",
		b.Name,
		b.Address,
		b.Occupation,
		b.Reason,
	)
}

var FriendBio = map[string]Biodata{
	"1": {
		Name:       "Mochamad",
		Address:    "Jl.Jombang-Mojokerto No.1",
		Occupation: "Programmer",
		Reason:     "Ingin belajar lebih banyak tentang Go",
	},
	"2": {
		Name:       "Syahrul",
		Address:    "Jl.Jombang-Mojokerto No.2",
		Occupation: "Freelancer",
		Reason:     "Ingin tau tentang microservice",
	},
	"3": {
		Name:       "Samsudin",
		Address:    "Jl.Jombang-Mojokerto No.3",
		Occupation: "Karyawan",
		Reason:     "Mencari nafkah",
	},
	"4": {
		Name:       "Hayolo",
		Address:    "Jl.Jombang-Mojokerto No.4",
		Occupation: "Mahasiswa",
		Reason:     "Mencari ijazah",
	},
}

func main() {
	args := os.Args[:]
	if len(args) < 2 {
		fmt.Println("Please insert one argument as biodata index!")
		return
	}

	index := args[1]
	val, exist := FriendBio[index]
	if exist {
		val.ShowBiodata()
		return
	}

	fmt.Printf("No data in index %s\n", index)
}
