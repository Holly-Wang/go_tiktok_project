package mysql

import (
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

func DelLike(keyID int64) error {
	like := Like{KeyID: keyID}
	// 修改结构体后是硬删除，修改前包含gorm.Model后生成了gorm.deletedat,自动变为软删除
	if err := db.Delete(&like).Error; err != nil {
		logs.Error("删除like表数据失败, err: ", err)
		return err
	}
	return nil
}

func DelComment(keyID int64) error {
	comment := Comment{CommentID: keyID}
	if err := db.Delete(&comment).Error; err != nil {
		logs.Error("删除comment表数据失败, err: ", err)
		return err
	}
	return nil
}
