
//2/18
//  https://sendgb.com/oImA0yaruzb  DL.part1.rar
//  https://sendgb.com/FIj6XtHpB1u  DL.part2.rar
//  https://sendgb.com/MnWPFN34JUp  DL.part3.rar
//  https://sendgb.com/jWAdtSgrcMP  DL.part4.rar

//  https://sendgb.com/RSYUqjD573b  DL.part5.rar
//  https://sendgb.com/H69MAfTZReO  DL.part6.rar















package database

import (
	"context"
	"database/sql"
	"dobledcloud.com/consumers/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(connection string) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	return &MysqlRepository{db}, nil
}

func (repo *MysqlRepository) Close() error {
	return repo.db.Close()
}

func (repo *MysqlRepository) GetEmissionByKey(ctx context.Context, key string) (*models.Emission, error) {
	var keys models.Emission
	row := repo.db.QueryRow("SELECT e.id FROM emissions e WHERE e.key = ?", key)
	err := row.Scan(&keys.Id)
	if err != nil && err != sql.ErrNoRows {
		return &keys, err
	}
	return &keys, nil
}

func (repo *MysqlRepository) GetSecretForEmission(ctx context.Context, emission_id int, emission_client string) bool {
	var exists bool
	err := repo.db.QueryRow("SELECT EXISTS(SELECT * FROM dobled_backend.keys e WHERE e.emission_id = ? AND e.emission_client = ?)", emission_id, emission_client).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Error checking if row exists %v", err)
	}
	return exists
}

func (repo *MysqlRepository) GetFilesPublishedByEmission(ctx context.Context, id int) ([]*models.Publishes, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT p.id,  p.date, f.md5, p.position, p.time_to_air, f.url "+
		"FROM publishes p, files f, file_publish pf WHERE p.emission_id = ? AND p.id = pf.publish_id AND pf.file_id = f.id"+
		" ORDER BY p.position ASC", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var publishes []*models.Publishes
	for rows.Next() {
		var publish = models.Publishes{}
		if err = rows.Scan(&publish.Id, &publish.Date, &publish.Md5, &publish.Position, &publish.TimeToAir, &publish.Url); err == nil {

			if err != nil {
				return nil, err
			}

			publishes = append(publishes, &publish)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return publishes, nil
}



// https://drive.google.com/drive/folders/1HSCHz1NO5om4pmthpO5kLFZFYFskOZo_?usp=share_link
// https://drive.google.com/drive/folders/1ztpQc5d-dkJWshPWhVYPbSCHXHWSCOjX?usp=sharing
