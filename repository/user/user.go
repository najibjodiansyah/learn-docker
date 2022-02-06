package user

import (
	"database/sql"
	"fmt"
	"log"

	"sirclo/layered/relation/entities"
)

//jangan lupa inisialisasi new , berlaku di controller

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result, err := r.db.Query("select id, name, email, password, address from users")
	if err != nil {
		return users, err
	}
	defer result.Close()
	for result.Next() {
		var user entities.User
		err := result.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.Address)
		if err != nil {
			log.Fatal("error di scan getUser")
		}
		users = append(users, user)
	}
	return users, nil

}

func (r *UserRepository)GetUser(id int) (entities.User, error) {
	var user entities.User
	result, err := r.db.Query("select id, name, email, password, address from users where id = ?", id)
	if err != nil {
		return user, err
	}

	defer result.Close()
	for result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address)
		if err != nil {
			return user, err
		}
		return user, nil
	}
	return user, fmt.Errorf("user not found")
}

func (r *UserRepository)UpdateUser(id int, user entities.User) error {
	result, err := r.db.Exec("UPDATE users SET name= ?, email= ?, password= ?, address= ? WHERE id = ?", user.Name, user.Email, user.Password, user.Address, id)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *UserRepository)DeleteUser(id int) error {
	result, err := r.db.Exec("DELETE from users where id = ?", id)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

func (r *UserRepository)PostUser(user entities.User) error {
	result, err := r.db.Exec("INSERT INTO users(name, email, password, address) VALUES(?,?,?,?)", user.Name, user.Email, user.Password, user.Address)
	if err != nil {
		return err
	}
	notAffected, _ := result.RowsAffected()
	if notAffected == 0 {
		return fmt.Errorf("user not created")
	}
	return nil
}