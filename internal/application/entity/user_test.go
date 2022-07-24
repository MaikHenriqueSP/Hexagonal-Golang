package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Run("User creation", func(t *testing.T) {
		expectedEmail := "email@email.com"
		expectedFullName := "First Name Last Name"
		user := NewUser(expectedEmail, expectedFullName)
		
		assert.NotNil(t, user)
		assert.Equal(t, expectedEmail, user.email)
		assert.Equal(t, expectedFullName, user.fullName)		
	})
}

