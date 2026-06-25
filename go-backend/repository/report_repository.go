package repository

import (
	"database/sql"
	"fmt"
	"vue-go-backend/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetAll() ([]models.Report, error) {
	query := `
		SELECT 
			rep.id, rep.reporter_id, rep.listing_id, rep.reported_user_id, 
			rep.reason, rep.status, rep.created_at,
			u.first_name, u.email,
			CASE 
				WHEN rep.listing_id IS NOT NULL THEN l.title
				WHEN rep.reported_user_id IS NOT NULL THEN CONCAT(ru.first_name, ' ', ru.last_name)
				ELSE 'N/A'
			END as target_name
		FROM reports rep
		JOIN users u ON rep.reporter_id = u.id
		LEFT JOIN listings l ON rep.listing_id = l.id
		LEFT JOIN users ru ON rep.reported_user_id = ru.id
		ORDER BY rep.created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reports: %w", err)
	}
	defer rows.Close()

	var reports []models.Report
	for rows.Next() {
		var rep models.Report

		// Use sql.NullInt64 to safely handle NULL values from the database
		var listingID sql.NullInt64
		var reportedUserID sql.NullInt64

		err := rows.Scan(
			&rep.ID, &rep.ReporterID, &listingID, &reportedUserID,
			&rep.Reason, &rep.Status, &rep.CreatedAt,
			&rep.ReporterName, &rep.ReporterEmail, &rep.TargetName,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan report: %w", err)
		}

		// Assign pointers only if the database value was not NULL
		if listingID.Valid {
			val := listingID.Int64
			rep.ListingID = &val
		}
		if reportedUserID.Valid {
			val := reportedUserID.Int64
			rep.ReportedUserID = &val
		}

		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *ReportRepository) UpdateStatus(id int64, status string) error {
	validStatuses := map[string]bool{"pending": true, "investigating": true, "resolved": true}
	if !validStatuses[status] {
		return fmt.Errorf("invalid status: %s", status)
	}
	_, err := r.db.Exec("UPDATE reports SET status = ? WHERE id = ?", status, id)
	if err != nil {
		return fmt.Errorf("failed to update report status: %w", err)
	}
	return nil
}
