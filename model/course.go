package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	id       int          `gorm:"type:bigint;size:100;primary_key;auto_increment:true"`
	P_xml     string      `gorm:"type:TEXT;"    json:"p_xml"`
	Retcode    int        `gorm:"type:NUMERIC"  json:"retcode"`
	Rettext string        `gorm:"type:VARCHAR(255);"  json:"rettext"`
}

type Result struct {
	retcode int
	rettext string
}

func LoadCourse(pxml string, db *gorm.DB) (Result, error) {
	var results = Result{}
	fmt.Println(pxml)
	err:=db.Debug().Exec("call load_market_course(?, @retcode, @rettext);",pxml).Error
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}
	error:=db.Debug().Raw("SELECT @retcode AS retcode, @rettext AS rettext;").Scan(&results).Error
	if error != nil {
		fmt.Println(error)
		panic(error)
	}
	return results, nil
}