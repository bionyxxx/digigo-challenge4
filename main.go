package main

import (
	"fmt"
	"os"
	"strconv"
)

// buat struct baru isinya field: no, name, address, occupation, reason
type student struct {
	no         int
	name       string
	address    string
	occupation string
	reason     string
}

// buat struct baru: nama field bebas tipe datanya slice student
type classroom struct {
	students []student
}

// buat interface isinya:
// getName, getAddress, getOccupation, getReason
type biodata interface {
	getName(no int) string
	getAddress(no int) string
	getOccupation(no int) string
	getReason(no int) string
}

// implement method2 interface ke struct list student
func (c classroom) getName(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.name
		}
	}
	return ""
}

func (c classroom) getAddress(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.address
		}
	}
	return ""
}

func (c classroom) getOccupation(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.occupation
		}
	}
	return ""
}

func (c classroom) getReason(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.reason
		}
	}
	return ""
}

func print(i int, b biodata) {
	fmt.Println("Nama\t: ", b.getName(i))
	fmt.Println("Alamat\t: ", b.getAddress(i))
	fmt.Println("Pekerjaan\t: ", b.getOccupation(i))
	fmt.Println("Alasan\t: ", b.getReason(i))
}

// os.Args -> buat ambil argument di command line

func main() {
	student_id := os.Args[1]

	// convert student_id pake strconv.Atoi
	i, _ := strconv.Atoi(student_id)

	// print(i)

	///inisialisasi var classroom
	var cr classroom

	//inisialisasi slice student
	var murid = []student{
		{
			no:         1,
			name:       "Dutta Fachrezy",
			address:    "Depok",
			occupation: "Mahasiswa",
			reason:     "Upgrade Skill",
		},
		{
			no:         2,
			name:       "Rivan",
			address:    "Cikarang",
			occupation: "Mahasiswa",
			reason:     "Learning something new",
		},
	}

	//murid diassign ke cr.students
	cr.students = murid

	// call func print
	print(i, cr)
}
