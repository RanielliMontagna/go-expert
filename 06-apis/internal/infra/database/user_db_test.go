package database

import (
	"goexpert/apis/internal/entity"
	"testing"

	"github.com/tj/assert"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func TestUserDB(t *testing.T) {
	t.Run("create new user", func(t *testing.T) {

		db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})

		if err != nil {
			t.Error(err)
		}

		db.AutoMigrate(entity.User{})

		user, _ := entity.NewUser("John Doe", "john@doe.com", "password")
		userDB := NewUser(db)

		err = userDB.Create(user)
		assert.Nil(t, err)

		var userFound entity.User
		err = db.First(&userFound, "id = ?", user.ID).Error
		assert.Nil(t, err)
		assert.Equal(t, user.ID, userFound.ID)
		assert.Equal(t, user.Name, userFound.Name)
		assert.Equal(t, user.Email, userFound.Email)
		assert.NotNil(t, userFound.Password)
	})
}
