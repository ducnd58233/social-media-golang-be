package hasher

import (
	"github.com/alexedwards/argon2id"
)

type bcryptHash struct{}

var params = &argon2id.Params{
	Memory:      128 * 1024,
	Iterations:  4,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

type Hasher interface {
	Hash(password string) string
	Validate(curPassword, hashedPassword string) bool
}

func NewBcryptHash() *bcryptHash {
	return &bcryptHash{}
}

func (h *bcryptHash) Hash(password string) string {
	hashedData, err := argon2id.CreateHash(password, params)

	if err != nil {
		panic(err)
	}

	return string(hashedData)
}

func (h *bcryptHash) Validate(curPassword, hashedPassword string) bool {
	_, err := argon2id.ComparePasswordAndHash(curPassword, hashedPassword)
	return err == nil
}
