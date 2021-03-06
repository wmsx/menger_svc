package models


type ThumbUpPost struct {
	Model
	MengerId int64 `gorm:"type:bigint;not null"`
	PostId   int64 `gorm:"type:bigint;not null"`
}


func AddThumbUp(mengerId, postId int64) error {
	thumbUpPost := ThumbUpPost{
		MengerId: mengerId,
		PostId:   postId,
	}
	if err := db.Create(&thumbUpPost).Error; err != nil {
		return err
	}
	return nil
}

func GetThumbUpPostIds(mengerId int64, pageNum, pageSize int32) (postIds []int64) {
	var id int64
	db.Table("t_thumb_up_post").
		Select("id").
		Where("menger_id = ?", mengerId).
		Offset(int(pageNum * pageSize)).
		Limit(1).
		Scan(&id)

	db.Table("t_thumb_up_post").
		Select("post_id").
		Where("id >= ?", id).
		Limit(int(pageSize)).
		Scan(&postIds)
	return
}
