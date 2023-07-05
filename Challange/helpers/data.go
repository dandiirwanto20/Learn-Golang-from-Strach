package helpers

func GetFriendList() []Friend {
	teman1 := Friend{No: 1, Nama: "Dandi", Pekerjaan: "IT Dev", Alamat: "Rembang", Alasan: "Ingin Belajar Golang"}
	teman2 := Friend{No: 2, Nama: "Dian", Pekerjaan: "Finance", Alamat: "Solo", Alasan: "Ingin Belajar Finance"}
	teman3 := Friend{No: 3, Nama: "Ucing", Pekerjaan: "Marketing", Alamat: "Semarang", Alasan: "Ingin Belajar Marketing"}

	friendList := []Friend{teman1, teman2, teman3}

	return friendList
}
