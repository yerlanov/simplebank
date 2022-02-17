package api

import (
	"github.com/stretchr/testify/require"
	db "github.com/yerlanov/simplebank/db/sqlc"
	"github.com/yerlanov/simplebank/util"
	"testing"
)

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}
