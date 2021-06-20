package store

import "github.com/asirkov/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	err := r.store.db.
		QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptedPassword).
		Scan(&u.ID)

	return u, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	err := r.store.db.
		QueryRow("SELECT id, email, encrypted_password FROM users WHERE email = $1", email).
		Scan(&u.ID, &u.Email, &u.EncryptedPassword)

	return u, err
}
