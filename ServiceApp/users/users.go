package users

import "gorm.io/gorm"

type Tbl_users struct {
	HP       string
	Nama     string
	Password string
	Alamat   string
}

func (u *Tbl_users) GantiPassword(connection *gorm.DB, newPassword string) (bool, error) {
	query := connection.Table("Tbl_users").Where("hp = ?", u.HP).Update("password", newPassword)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func Register(connection *gorm.DB, newUser Tbl_users) (bool, error) {
	query := connection.Create(&newUser)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func Login(connection *gorm.DB, hp string, password string) (Tbl_users, error) {
	var result Tbl_users
	err := connection.Where("hp = ? AND password = ?", hp, password).First(&result).Error
	if err != nil {
		return Tbl_users{}, err
	}

	return result, nil
}
