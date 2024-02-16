package main

import (
	"GoSql/config"
	"GoSql/users"
	"fmt"
)

func main() {
	// input database
	database := config.InitMysql()
	// migrate database
	config.Migrate(database)
	var input int
	for input <= 5 {
		fmt.Println(" Pilih menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Lupa Password")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&input)
		// menu login
		if input == 1 {
			var isRunning bool = true
			for isRunning {
				var hp string
				var password string
				fmt.Println("Masukkan HP")
				fmt.Scanln(&hp)
				fmt.Println("Masukkan Password")
				fmt.Scanln(&password)
				loggedIn, err := users.Login(database, hp, password)
				if err == nil {
					fmt.Println("Selamat Datang,", loggedIn.Nama)
					if loggedIn != (users.Tbl_user{}) {
						//  login berhasil
						var menuInput int
						for menuInput <= 10 {
							fmt.Println("Pilih menu:")
							fmt.Printf("Saldo anda : %.2f\n", loggedIn.Balance)
							fmt.Println("1. Tambah User")
							fmt.Println("2. Lihat User")
							fmt.Println("3. Edit User")
							fmt.Println("4. Hapus User")
							fmt.Println("5. TopUp Saldo")
							fmt.Println("6. Transfer")
							fmt.Println("7. History Transfer")
							fmt.Println("8. History TopUp")
							fmt.Println("9. Cari User")
							fmt.Println("99. Keluar")
							fmt.Print("Masukkan pilihan: ")
							fmt.Scanln(&menuInput)

							switch menuInput {
							case 1:
								// form untuk menambahkan users
								var newUser users.Tbl_user
								fmt.Print("Masukkan nama : ")
								fmt.Scanln(&newUser.Nama)
								fmt.Print("Masukkan nomor HP : ")
								fmt.Scanln(&newUser.HP)
								fmt.Print("Masukkan password : ")
								fmt.Scanln(&newUser.Password)
								fmt.Print("Masukkan alamat : ")
								fmt.Scanln(&newUser.Alamat)

								success, err := users.TambahUsers(database, newUser)
								if err != nil {
									fmt.Println("Gagal menambahkan barang:", err)
								} else if success {
									fmt.Println("--Barang berhasil ditambahkan--")
								} else {
									fmt.Println("Gagal menambahkan barang")
								}
							case 2:

								// tampilkan daftar users
								daftarBarang, err := users.LihatUsers(database, loggedIn.ID)
								if err != nil {
									fmt.Println("Gagal mendapatkan daftar user:", err)
								} else {
									fmt.Println("Daftar Users:")
									for i, users := range daftarBarang {
										fmt.Printf("%d. Nama: %s, Password: %s, Alamat: %s\n", i+1, users.Nama, users.Password, users.Alamat)
									}
								}
							case 3:
								// edit form users
								var idEdit uint
								fmt.Print("Masukan id yg ingin diedit : ")
								fmt.Scanln(&idEdit)
								var hp string
								fmt.Print("Masukan no telp : ")
								fmt.Scanln(&hp)
								var namaUser string
								fmt.Print("Masukan nama user : ")
								fmt.Scanln(&namaUser)
								var password string
								fmt.Print("Masukan password : ")
								fmt.Scanln(&password)
								var alamat string
								fmt.Print("Masukan alamat  : ")
								fmt.Scanln(&alamat)
								updateData := map[string]interface{}{
									"HP":       hp,
									"Nama":     namaUser,
									"Password": password,
									"Alamat":   alamat,
								}
								success, err := users.EditUsers(database, idEdit, updateData)
								if err != nil {
									fmt.Println("Gagal mengedit users : ", err)
								} else if success {
									fmt.Println("Users berhasil diubah")
								} else {
									fmt.Println("Tidak ada perubahan pada user")
								}
							case 4:
							    	var hpToDelete string
							    		fmt.Print("Masukkan nomor HP pengguna yang ingin dihapus: ")
							    		fmt.Scanln(&hpToDelete)
							
							    	success, err := users.HapusUsers(database, hpToDelete)
							    	if err != nil {
							        	fmt.Println("Gagal menghapus pengguna:", err)
							    	} else if success {
							        	fmt.Println("Pengguna berhasil dihapus.")
							   	 } else {
							        fmt.Println("Pengguna tidak ditemukan.")
							    }

							case 5:
								fmt.Print("Masukan jumlah TopUp : ")
								var jumlahTopUp float64
								fmt.Scanln(&jumlahTopUp)

								err := users.TopUp(int(loggedIn.ID), jumlahTopUp, database)
								if err != nil {
									fmt.Println("gagal melakukan TopUp:", err)
								} else {
									fmt.Println("TopUp berhasil dilakukan")
								}
							case 6:
							case 7:
							case 8:
								history, err := users.HistoryTopUp(int(loggedIn.ID), database)
								if err != nil {
									fmt.Println("gagal melihat history TopUp ", err)
								}
								fmt.Println("TopUp History : ")
								for _, entry := range history {
									fmt.Printf("ID: %d, Balance: %2.f\n", entry.ID, entry.Balance)
								}
							case 9:
								// Cari users berdasarkan id
								var cari string
								fmt.Print("Masukan no Hp Users : ")
								fmt.Scanln(&cari)

								hasilPencarian, err := users.CariUsers(database, cari)
								if err != nil {
									fmt.Println("Gagal melakukan pencarian : ", err)
								} else if len(hasilPencarian) > 0 {
									fmt.Println("Hasil pencarian : ")
									for i, users := range hasilPencarian {
										fmt.Printf("%d. Nama: %s, Password: %s, Alamat: %s\n", i+1, users.Nama, users.Password, users.Alamat)
									}

								} else {
									fmt.Println("Tidak ada kriteria barang")
								}
							case 99:
								fmt.Println("Keluar")
							default:
								fmt.Println("Pilihan tidak valid")
							}
						}
					}
					isRunning = false
				} else {
					var inputExit string
					fmt.Println("Ketik 'EXIT' untuk kembali ke menu sebelumnya")
					fmt.Scanln(&inputExit)
					if inputExit == "EXIT" {
						isRunning = false
					}
				}
			}
			// kalo sukses welcome, kalo gagal isi lagi
		} else if input == 2 {
			// menu register
			var newUser users.Tbl_user
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan nomor HP : ")
			fmt.Scanln(&newUser.HP)
			fmt.Print("Masukkan password : ")
			fmt.Scanln(&newUser.Password)
			fmt.Print("Masukkan alamat : ")
			fmt.Scanln(&newUser.Alamat)
			success, err := users.Register(database, newUser)
			if err != nil {
				fmt.Println("terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
			}

			if success {
				fmt.Println("--selamat anda telah terdaftar--")

			}
		} else if input == 3 {
			// lupa kata sandi
			var hp string
			fmt.Print("Masukkan nomor HP Anda: ")
			fmt.Scanln(&hp)

			var currentPassword string
			fmt.Print("Masukkan kata sandi saat ini: ")
			fmt.Scanln(&currentPassword)

			var newPassword string
			fmt.Print("Masukkan kata sandi baru: ")
			fmt.Scanln(&newPassword)

			// Verifikasi login pengguna
			user, err := users.Login(database, hp, currentPassword)
			if err != nil {
				fmt.Println("Terjadi kesalahan:", err)
				return
			}
			success, err := user.GantiPassword(database, newPassword)
			if err != nil {
				fmt.Println("Terjadi kesalahan:", err)
				return
			}

			if success {
				fmt.Println("--Password berhasil diubah--")
			} else {
				fmt.Println("Gagal mengubah password.")
			}
		}
	}
	fmt.Println("Terimakasih  telah bertransaksi !!!")

}
