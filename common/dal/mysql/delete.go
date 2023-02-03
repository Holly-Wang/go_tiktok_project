package mysql

import "fmt"

func DelLike(keyID uint64) error {
	like := Like{KeyID: keyID}
	result := db.Delete(&like) //修改结构体后是硬删除，修改前包含gorm.Model后生成了gorm.deletedat,自动变为软删除
	if result.Error != nil {
		fmt.Println("删除like表数据失败,error: " + result.Error.Error())
		return result.Error
	}
	return nil
}
