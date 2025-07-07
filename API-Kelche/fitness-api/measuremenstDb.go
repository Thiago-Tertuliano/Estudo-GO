package repositories

import (
	storage "fitness-api"
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"
)

func CreateMeasurement(measurement models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err != nil {
		return models.Measurements{}, err
	}

	measurement.CreatedAt = time.Now()
	return measurement, nil
}

func GetMeasurements() ([]models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT id, user_id, weight, height, body_fat, created_at FROM measurements`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var measurements []models.Measurements
	for rows.Next() {
		var measurement models.Measurements
		err := rows.Scan(&measurement.Id, &measurement.UserId, &measurement.Weight, &measurement.Height, &measurement.BodyFat, &measurement.CreatedAt)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, measurement)
	}

	return measurements, nil
}

func GetMeasurements(id int) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `SELECT id, user_id, weight, height, body_fat, created_at FROM measurements WHERE id = $1`

	var measurement models.Measurements
	err := db.QueryRow(sqlStatement, id).Scan(&measurement.Id, &measurement.UserId, &measurement.Weight, &measurement.Height, &measurement.BodyFat, &measurement.CreatedAt)
	if err != nil {
		return models.Measurements{}, err
	}

	return measurement, nil
}

func UpdateMeasurement(measurement models.Measurements, id int) (models.Measurements, error) {
	db := storage.GetDB()
	sqlStatement := `UPDATE measurements SET user_id = $1, weight = $2, height = $3, body_fat = $4, updated_at = $5 WHERE id = $6 RETURNING id`

	err := db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.now().Scan(&id))
	if err != nil {
		return models.Measurements{}, err
	}

	measurement.Id = id
	return measurement, nil
}

func DeleteMeasurement(id int) error {
	db := storage.GetDB()
	sqlStatement := `DELETE FROM measurements WHERE id = $1`

	_, err := db.Exec(sqlStatement, id)
	return  err
}