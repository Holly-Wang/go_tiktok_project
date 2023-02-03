package mysql

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB()
	err := CreateLike(123, 321, 111)
	if err != nil {
		log.Fatal(err)
	}
	ID, find_err := FindIDinLike(321, 111)
	if find_err != nil {
		log.Fatal(find_err)
	}
	fmt.Println("keyid:", ID)
	Del := DelLike(ID)
	if Del != nil {
		log.Fatal("删除失败")
	}
	//UpdateLikeCount(123) //测试数据已删除
	os.Exit(m.Run())
}
