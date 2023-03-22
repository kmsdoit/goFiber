package bookmark

import (
	"goFiber/main/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateBookmarkService(name string, url string) (models.Bookmark, error) {
	var newBookmark = models.Bookmark{Name: name, Url: url}

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return newBookmark, err
	}
	db.Create(&models.Bookmark{Name: name, Url: url})

	return newBookmark, nil
}

func GetAllBookmarkService() ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		return bookmarks, err
	}

	db.Find(&bookmarks)

	return bookmarks, nil
}
