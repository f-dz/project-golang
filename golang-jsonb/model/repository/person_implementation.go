package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"golang-jsonb/helper"
	"golang-jsonb/model/entity"
)

type personImplementation struct {
	DB *sql.DB
}

func NewPersonImplementation(db *sql.DB) PersonInterface {
	return &personImplementation{DB: db}
}

func (repository *personImplementation) GetAllData(ctx context.Context) ([]entity.Person, error) {
	query := `SELECT attrs FROM people;`
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var people []entity.Person
	for rows.Next() {
		person := entity.Person{}
		data := []byte{}

		rows.Scan(&data)
		json.Unmarshal(data, &person)
		people = append(people, person)
	}
	return people, nil
}

func (repository *personImplementation) GetData(ctx context.Context, name string) (entity.Person, error) {
	query := `SELECT * FROM people WHERE attrs->'name' ? $1 LIMIT 1;`
	rows, err := repository.DB.QueryContext(ctx, query, name)
	person := entity.Person{}
	if err != nil {
		return person, err
	}

	defer rows.Close()

	if rows.Next() {
		data := []byte{}
		rows.Scan(&data)
		json.Unmarshal(data, &person)
	} else {
		return person, nil
	}
	return person, nil
}

func (repository *personImplementation) CreateData(ctx context.Context, person entity.Person) (entity.Person, error) {
	query := `INSERT INTO people(attrs) VALUES($1);`
	personJSON, e := json.Marshal(person)
	helper.CheckErrorFatal(e)

	result, _ := repository.DB.ExecContext(ctx, query, personJSON)

	rows, _ := result.RowsAffected()

	if rows == 0 {
		return person, errors.New("no rows affected")
	} else {
		return person, nil
	}
}

func (repository *personImplementation) UpdateData(ctx context.Context, name string, age int) (entity.Person, error) {
	query := `UPDATE people SET attrs = jsonb_set(attrs, '{age}', $1) WHERE attrs->'name' ? $2`
	result, _ := repository.DB.ExecContext(ctx, query, age, name)

	rows, _ := result.RowsAffected()

	var person entity.Person
	if rows == 0 {
		return person, errors.New("no rows affected")
	} else {
		person.Name = name
		person.Age = age
		return person, nil
	}

}

func (repository *personImplementation) DeleteData(ctx context.Context, name string) error {
	query := `DELETE FROM people WHERE attrs->'name' ? $1;`
	result, _ := repository.DB.ExecContext(ctx, query, name)

	rows, _ := result.RowsAffected()

	if rows == 0 {
		return errors.New("no rows affected")
	} else {
		return nil
	}
}
