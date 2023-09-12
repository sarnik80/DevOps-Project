package repository

import (
	"context"
	"database/sql"
	"p1-reciever/models"
	"time"
)

type PostgresRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		DB: db,
	}
}

func (repo *PostgresRepo) InitDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	query := `create table if not exists temp (
		id serial primary key ,
		message varchar(500)
	)`
	_, err := repo.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepo) InsertMessage(msg *models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	query := `insert into temp (message) 
			 values($1)`
	_, err := repo.DB.ExecContext(ctx, query, msg.Payload)
	if err != nil {
		return err
	}
	return nil
}
