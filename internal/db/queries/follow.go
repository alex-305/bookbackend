package queries

import (
	"github.com/alex-305/bookbackend/internal/models"
	"gorm.io/gorm"
)

func ChangeFollowerCount(followedID models.UserID, g *gorm.DB, change int8) *gorm.DB {
	return g.Model(&models.User{}).Where(("userid = ?"), followedID).Update("followercount", gorm.Expr("followercount+?", change))
}

func ChangeFollowingCount(followerID models.UserID, g *gorm.DB, change int8) *gorm.DB {
	return g.Model(&models.User{}).Where(("userid = ?"), followerID).Update("followingcount", gorm.Expr("followingcount+?", change))
}
