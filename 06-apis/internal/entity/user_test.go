package entity

import (
	"testing"

	"github.com/tj/assert"
)

func TestUser(t *testing.T) {
	name := "john"
	email := "john@doe.com"
	password := "password"

	t.Run("create new user", func(t *testing.T) {

		user, err := NewUser(name, email, password)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, email, user.Email)
		assert.NotEmpty(t, user.Password)
	})

	t.Run("compare password", func(t *testing.T) {
		user, _ := NewUser(name, email, password)

		assert.True(t, user.ValidatePassword(password))
		assert.False(t, user.ValidatePassword("wrong-password"))
		assert.NotEqual(t, password, user.Password)
	})
}
