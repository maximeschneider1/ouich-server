package dao

import (
	"awesomeProject12/model"
	"data-back-real/service"
	"database/sql"
	"math/rand"
)

func QueryAllQuotes(db *sql.DB) ([]model.Quote, error) {
	var ac []model.Quote

	rows, err := db.Query("SELECT client_id, name, phone, city, downloads, emails_opening, potential_gains, step FROM leads WHERE conseiller_id= $1"); if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id model.Lead
		err := rows.Scan(&id.LeadID, &id.Name, &id.Phone, &id.City, &id.ContentDownloaded, &id.OpenedEmails, &id.Profitability, &id.Step); if err != nil {
			return nil, err
		}
		id.TimeSpent = rand.Intn(100 - 10) + 10
		id.WeeksSinceInactive = rand.Intn(30 - 1) + 1
		id = service.FromDBToWeightedCriteras(id)
		id.Score = service.ScoreCalculator(id)
		id.StepConverted = stepConverter(id.Step)

		ac = append(ac, id)
	}
	return ac, nil
}
