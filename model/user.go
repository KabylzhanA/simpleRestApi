package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"regexp"
	"strings"
	"time"
)

type JsonBirthDate time.Time
type User struct {
	gorm.Model
	id       int           `gorm:"type:bigint;size:100;primary_key;auto_increment:true"`
	Name     string        `gorm:"size:100;not null"               json:"name"`
	Email    string        `gorm:"type:varchar(100);unique_index"  json:"email"`
	Birthday JsonBirthDate `gorm:"type:date;"               json:"birthday"`
}

func (j JsonBirthDate) Value() (driver.Value, error) {
	return time.Time(j).Format("2006-01-02"), nil
}
func (j *JsonBirthDate) Scan(value interface{}) error {
	if value == nil {
		*j = JsonBirthDate(time.Time{})
		return nil
	}
	if val, err := driver.DefaultParameterConverter.ConvertValue(value); err == nil {
		if v, ok := val.(time.Time); ok {
			*j = JsonBirthDate(v)
			return nil
		}
	}
	return errors.New("failed to scan JsonBirthDate")
}
func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (j JsonBirthDate) IsZero() bool {
	t := time.Time(j)
	return t.IsZero()
}
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Birthday.IsZero() {
		return errors.New("birthdate is required")
	}
	if checkmail(u.Email) == false {
		return errors.New("Invalid Email")
	}
	return nil

}

func checkmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
func GetAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	if err := db.Debug().Table("users").Find(&users).Error; err != nil {
		return &[]User{}, err
	}
	return &users, nil
}
func (u *User) GetUser(db *gorm.DB) (*User, error) {
	account := &User{}
	if err := db.Debug().Table("users").Where("email = ?", u.Email).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
func (u *User) DeleteUser(db *gorm.DB) (bool, error) {
	if &u.id != nil {
		db.Delete(&u)
		return true, nil
	}
	return false, errors.New("Please, provide an id")
}

func (u *User) UpdateUser(id int, db *gorm.DB) (*User, error) {
	if err := db.Debug().Table("users").Where("id = ?", id).Updates(User{
		Name:     u.Name,
		Email:    u.Email,
		Birthday: u.Birthday}).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}
func GetUserById(id int, db *gorm.DB) (*User, error) {
	user := &User{}
	if err := db.Debug().Table("users").Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) String() string {
	return fmt.Sprintf("name %s,  email %s, birthdate %s", u.Name, u.Email, u.Birthday)
}
