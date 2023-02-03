package mysql

import (
	"log"
	"testing"
)

func TestDatabaseOp(t *testing.T) {
	InitDB()
	err := CreateLike(123, 321, 111)
	if err != nil {
		log.Fatal(err)
	}
	id, findErr := FindIDinLike(321, 111)
	if findErr != nil {
		log.Fatal(findErr)
	}
	//fmt.Println("keyid:", ID)
	Del := DelLike(id)
	if Del != nil {
		log.Fatal("删除失败")
	}
}
