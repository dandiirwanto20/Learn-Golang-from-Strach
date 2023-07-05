package helpers

import "fmt"

func ShowFriendData(friendList []Friend, no int) {
	found := false
	for _, friend := range friendList {
		if friend.No == no {
			found = true
			fmt.Println("Data Teman Kelas:")
			fmt.Println("No: ", friend.No)
			fmt.Println("Nama: ", friend.Nama)
			fmt.Println("Alamat: ", friend.Alamat)
			fmt.Println("Pekerjaan: ", friend.Pekerjaan)
			fmt.Println("Alasan: ", friend.Alasan)
			break
		}
	}

	if !found {
		fmt.Println("Tidak ada teman dengan nomor ", no)
	}

}
