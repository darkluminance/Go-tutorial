// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: university_application.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createUniversityApplication = `-- name: CreateUniversityApplication :one
INSERT INTO university_application (
    user_id,
    university_id,
    application_type,
    shortlisted_faculties_id,
    recommenders_id,
    application_status,
    language_score_submitted,
    gre_submitted,
    gmat_submitted,
    remarks
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, user_id, university_id, shortlisted_faculties_id, recommenders_id, application_type, application_status, language_score_submitted, gre_submitted, gmat_submitted, remarks, created_at, updated_at
`

type CreateUniversityApplicationParams struct {
	UserID                 uuid.UUID
	UniversityID           uuid.UUID
	ApplicationType        NullApplicationTypeEnum
	ShortlistedFacultiesID []uuid.UUID
	RecommendersID         []uuid.UUID
	ApplicationStatus      NullUniversityApplicationStatusEnum
	LanguageScoreSubmitted sql.NullBool
	GreSubmitted           sql.NullBool
	GmatSubmitted          sql.NullBool
	Remarks                sql.NullString
}

func (q *Queries) CreateUniversityApplication(ctx context.Context, arg CreateUniversityApplicationParams) (UniversityApplication, error) {
	row := q.db.QueryRowContext(ctx, createUniversityApplication,
		arg.UserID,
		arg.UniversityID,
		arg.ApplicationType,
		pq.Array(arg.ShortlistedFacultiesID),
		pq.Array(arg.RecommendersID),
		arg.ApplicationStatus,
		arg.LanguageScoreSubmitted,
		arg.GreSubmitted,
		arg.GmatSubmitted,
		arg.Remarks,
	)
	var i UniversityApplication
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.UniversityID,
		pq.Array(&i.ShortlistedFacultiesID),
		pq.Array(&i.RecommendersID),
		&i.ApplicationType,
		&i.ApplicationStatus,
		&i.LanguageScoreSubmitted,
		&i.GreSubmitted,
		&i.GmatSubmitted,
		&i.Remarks,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUniversityApplicationById = `-- name: DeleteUniversityApplicationById :one
DELETE FROM university_application
WHERE id = $1
RETURNING id, user_id, university_id, shortlisted_faculties_id, recommenders_id, application_type, application_status, language_score_submitted, gre_submitted, gmat_submitted, remarks, created_at, updated_at
`

func (q *Queries) DeleteUniversityApplicationById(ctx context.Context, id uuid.UUID) (UniversityApplication, error) {
	row := q.db.QueryRowContext(ctx, deleteUniversityApplicationById, id)
	var i UniversityApplication
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.UniversityID,
		pq.Array(&i.ShortlistedFacultiesID),
		pq.Array(&i.RecommendersID),
		&i.ApplicationType,
		&i.ApplicationStatus,
		&i.LanguageScoreSubmitted,
		&i.GreSubmitted,
		&i.GmatSubmitted,
		&i.Remarks,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUniversityApplicationById = `-- name: GetUniversityApplicationById :one
SELECT id, user_id, university_id, shortlisted_faculties_id, recommenders_id, application_type, application_status, language_score_submitted, gre_submitted, gmat_submitted, remarks, created_at, updated_at FROM university_application
WHERE id = $1
`

func (q *Queries) GetUniversityApplicationById(ctx context.Context, id uuid.UUID) (UniversityApplication, error) {
	row := q.db.QueryRowContext(ctx, getUniversityApplicationById, id)
	var i UniversityApplication
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.UniversityID,
		pq.Array(&i.ShortlistedFacultiesID),
		pq.Array(&i.RecommendersID),
		&i.ApplicationType,
		&i.ApplicationStatus,
		&i.LanguageScoreSubmitted,
		&i.GreSubmitted,
		&i.GmatSubmitted,
		&i.Remarks,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUniversityApplicationsOfUser = `-- name: GetUniversityApplicationsOfUser :many
SELECT id, user_id, university_id, shortlisted_faculties_id, recommenders_id, application_type, application_status, language_score_submitted, gre_submitted, gmat_submitted, remarks, created_at, updated_at FROM university_application
WHERE user_id = $1
`

func (q *Queries) GetUniversityApplicationsOfUser(ctx context.Context, userID uuid.UUID) ([]UniversityApplication, error) {
	rows, err := q.db.QueryContext(ctx, getUniversityApplicationsOfUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UniversityApplication
	for rows.Next() {
		var i UniversityApplication
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.UniversityID,
			pq.Array(&i.ShortlistedFacultiesID),
			pq.Array(&i.RecommendersID),
			&i.ApplicationType,
			&i.ApplicationStatus,
			&i.LanguageScoreSubmitted,
			&i.GreSubmitted,
			&i.GmatSubmitted,
			&i.Remarks,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUniversityApplicationByID = `-- name: UpdateUniversityApplicationByID :one
UPDATE university_application 
SET university_id = $2,
    application_type = $3,
    shortlisted_faculties_id = $4,
    recommenders_id = $5,
    application_status = $6,
    language_score_submitted = $7,
    gre_submitted = $8,
    gmat_submitted = $9,
    remarks = $10,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, university_id, shortlisted_faculties_id, recommenders_id, application_type, application_status, language_score_submitted, gre_submitted, gmat_submitted, remarks, created_at, updated_at
`

type UpdateUniversityApplicationByIDParams struct {
	ID                     uuid.UUID
	UniversityID           uuid.UUID
	ApplicationType        NullApplicationTypeEnum
	ShortlistedFacultiesID []uuid.UUID
	RecommendersID         []uuid.UUID
	ApplicationStatus      NullUniversityApplicationStatusEnum
	LanguageScoreSubmitted sql.NullBool
	GreSubmitted           sql.NullBool
	GmatSubmitted          sql.NullBool
	Remarks                sql.NullString
}

func (q *Queries) UpdateUniversityApplicationByID(ctx context.Context, arg UpdateUniversityApplicationByIDParams) (UniversityApplication, error) {
	row := q.db.QueryRowContext(ctx, updateUniversityApplicationByID,
		arg.ID,
		arg.UniversityID,
		arg.ApplicationType,
		pq.Array(arg.ShortlistedFacultiesID),
		pq.Array(arg.RecommendersID),
		arg.ApplicationStatus,
		arg.LanguageScoreSubmitted,
		arg.GreSubmitted,
		arg.GmatSubmitted,
		arg.Remarks,
	)
	var i UniversityApplication
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.UniversityID,
		pq.Array(&i.ShortlistedFacultiesID),
		pq.Array(&i.RecommendersID),
		&i.ApplicationType,
		&i.ApplicationStatus,
		&i.LanguageScoreSubmitted,
		&i.GreSubmitted,
		&i.GmatSubmitted,
		&i.Remarks,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
