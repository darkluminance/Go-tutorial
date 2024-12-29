package main

import (
	"time"

	"github.com/darkluminance/higher-studies-application-tracker/go-server/internal/database"
	"github.com/google/uuid"
)

// Structs

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Recommender struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Designation  string    `json:"designation"`
	Institution  string    `json:"institution"`
	Relationship string    `json:"relationship"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Faculty struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"user_id"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	UniversityID     uuid.UUID `json:"university_id"`
	Designation      string    `json:"designation"`
	ResearchAreas    []string  `json:"research_areas"`
	InterestedPapers []string  `json:"interested_papers"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Interview struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	FacultyID   uuid.UUID `json:"faculty_id"`
	Date        time.Time `json:"date"`
	IsCompleted bool      `json:"is_completed"`
	Remarks     string    `json:"remarks"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Mail struct {
	ID                   uuid.UUID `json:"id"`
	UserID               uuid.UUID `json:"user_id"`
	FacultyID            uuid.UUID `json:"faculty_id"`
	IsMailed             bool      `json:"is_mailed"`
	IsMailReplied        bool      `json:"is_mail_replied"`
	ReplyVibe            string    `json:"reply_vibe"`
	IsInterviewRequested bool      `json:"is_interview_requested"`
	Remarks              string    `json:"remarks"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type University struct {
	ID                           uuid.UUID `json:"id"`
	UserID                       uuid.UUID `json:"user_id"`
	Name                         string    `json:"name"`
	Website                      string    `json:"website"`
	Location                     string    `json:"location"`
	MainRanking                  int       `json:"main_ranking"`
	SubjectRanking               int       `json:"subject_ranking"`
	ApplicationDeadline          time.Time `json:"application_deadline"`
	EarlyDeadline                time.Time `json:"early_deadline"`
	IsGreMust                    bool      `json:"is_gre_must"`
	IsGmatMust                   bool      `json:"is_gmat_must"`
	LorCount                     int       `json:"lor_count"`
	IsOfficialTranscriptRequired bool      `json:"is_official_transcript_required"`
	IsTranscriptNeedsEvaluation  bool      `json:"is_transcript_needs_evaluation"`
	AcceptedEvaluations          []string  `json:"accepted_evaluations"`
	Remarks                      string    `json:"remarks"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
}

type UniversityApplication struct {
	ID                     uuid.UUID   `json:"id"`
	UserID                 uuid.UUID   `json:"user_id"`
	UniversityID           uuid.UUID   `json:"university_id"`
	ApplicationType        string      `json:"application_type"`
	ShortlistedFacultiesID []uuid.UUID `json:"shortlisted_faculties_id"`
	RecommendersID         []uuid.UUID `json:"recommenders_id"`
	ApplicationStatus      string      `json:"application_status"`
	LanguageScoreSubmitted bool        `json:"language_score_submitted"`
	GreSubmitted           bool        `json:"gre_submitted"`
	GmatSubmitted          bool        `json:"gmat_submitted"`
	Remarks                string      `json:"remarks"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
}

// Functions
func databaseUserToUser(databaseUser database.User) User {
	return User{
		ID:        databaseUser.ID,
		Name:      databaseUser.Name,
		Username:  databaseUser.Username,
		Email:     databaseUser.Email,
		CreatedAt: databaseUser.CreatedAt.Time,
		UpdatedAt: databaseUser.UpdatedAt.Time,
	}
}

func databaseUniversityToUniversity(databaseUniversity database.University) University {
	return University{
		ID:                           databaseUniversity.ID,
		UserID:                       databaseUniversity.UserID,
		Name:                         databaseUniversity.Name,
		Website:                      databaseUniversity.Website.String,
		Location:                     databaseUniversity.Location.String,
		MainRanking:                  int(databaseUniversity.MainRanking.Int32),
		SubjectRanking:               int(databaseUniversity.SubjectRanking.Int32),
		ApplicationDeadline:          databaseUniversity.ApplicationDeadline.Time,
		EarlyDeadline:                databaseUniversity.EarlyDeadline.Time,
		IsGreMust:                    databaseUniversity.IsGreMust.Bool,
		IsGmatMust:                   databaseUniversity.IsGmatMust.Bool,
		LorCount:                     int(databaseUniversity.LorCount.Int32),
		IsOfficialTranscriptRequired: databaseUniversity.IsOfficialTranscriptRequired.Bool,
		IsTranscriptNeedsEvaluation:  databaseUniversity.IsTranscriptNeedsEvaluation.Bool,
		AcceptedEvaluations:          databaseUniversity.AcceptedEvaluations,
		Remarks:                      databaseUniversity.Remarks.String,
		CreatedAt:                    databaseUniversity.CreatedAt.Time,
		UpdatedAt:                    databaseUniversity.UpdatedAt.Time,
	}
}

func databaseRecommenderToRecommender(databaseRecommender database.Recommender) Recommender {
	return Recommender{
		ID:           databaseRecommender.ID,
		UserID:       databaseRecommender.UserID,
		Name:         databaseRecommender.Name,
		Email:        databaseRecommender.Email,
		Designation:  databaseRecommender.Designation.String,
		Institution:  databaseRecommender.Institution.String,
		Relationship: databaseRecommender.Relationship.String,
		CreatedAt:    databaseRecommender.CreatedAt.Time,
		UpdatedAt:    databaseRecommender.UpdatedAt.Time,
	}
}

func databaseFacultyToFaculty(databaseFaculty database.Faculty) Faculty {
	return Faculty{
		ID:               databaseFaculty.ID,
		UserID:           databaseFaculty.UserID,
		Name:             databaseFaculty.Name,
		Email:            databaseFaculty.Email.String,
		UniversityID:     databaseFaculty.UniversityID,
		Designation:      databaseFaculty.Designation,
		ResearchAreas:    databaseFaculty.ResearchAreas,
		InterestedPapers: databaseFaculty.InterestedPapers,
		CreatedAt:        databaseFaculty.CreatedAt.Time,
		UpdatedAt:        databaseFaculty.UpdatedAt.Time,
	}
}

func databaseInterviewToInterview(databaseInterview database.Interview) Interview {
	return Interview{
		ID:          databaseInterview.ID,
		UserID:      databaseInterview.UserID,
		FacultyID:   databaseInterview.FacultyID,
		Date:        databaseInterview.Date.Time,
		IsCompleted: databaseInterview.IsCompleted.Bool,
		Remarks:     databaseInterview.Remarks.String,
		CreatedAt:   databaseInterview.CreatedAt.Time,
		UpdatedAt:   databaseInterview.UpdatedAt.Time,
	}
}

func databaseMailToMailWithID(databaseMail database.Mail) Mail {
	return Mail{
		ID:                   databaseMail.ID,
		UserID:               databaseMail.UserID,
		FacultyID:            databaseMail.FacultyID,
		IsMailed:             databaseMail.IsMailed.Bool,
		IsMailReplied:        databaseMail.IsMailReplied.Bool,
		ReplyVibe:            string(databaseMail.ReplyVibe.ReplyVibeEnum),
		IsInterviewRequested: databaseMail.IsInterviewRequested.Bool,
		Remarks:              databaseMail.Remarks.String,
		CreatedAt:            databaseMail.CreatedAt.Time,
		UpdatedAt:            databaseMail.UpdatedAt.Time,
	}
}

func databaseUniversityApplicationToUniversityApplication(databaseApplication database.UniversityApplication) UniversityApplication {
	return UniversityApplication{
		ID:                     databaseApplication.ID,
		UserID:                 databaseApplication.UserID,
		UniversityID:           databaseApplication.UniversityID,
		ApplicationType:        string(databaseApplication.ApplicationType.ApplicationTypeEnum),
		ShortlistedFacultiesID: databaseApplication.ShortlistedFacultiesID,
		RecommendersID:         databaseApplication.RecommendersID,
		ApplicationStatus:      string(databaseApplication.ApplicationStatus.UniversityApplicationStatusEnum),
		LanguageScoreSubmitted: databaseApplication.LanguageScoreSubmitted.Bool,
		GreSubmitted:           databaseApplication.GreSubmitted.Bool,
		GmatSubmitted:          databaseApplication.GmatSubmitted.Bool,
		Remarks:                databaseApplication.Remarks.String,
		CreatedAt:              databaseApplication.CreatedAt.Time,
		UpdatedAt:              databaseApplication.UpdatedAt.Time,
	}
}
