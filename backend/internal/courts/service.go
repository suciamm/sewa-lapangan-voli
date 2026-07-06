package courts

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sqlc_db "sewa-lapangan-voli/db"
)

func CreateCourtService(ctx context.Context, q *sqlc_db.Queries, ownerID int64, req CreateCourtRequest) (CourtResponse, error) {
	var avgRating float64
	result, err := q.CreateCourt(ctx, sqlc_db.CreateCourtParams{
		OwnerID:      ownerID,
		Name:         req.Name,
		Description:  toNullString(req.Description),
		Address:      toNullString(req.Address),
		City:         toNullString(req.City),
		PricePerHour: req.PricePerHour,
		Status:       sqlc_db.CourtsStatus("active"),
		AvgRating:    avgRating,
		TotalReviews: 0,
	})
	if err != nil {
		return CourtResponse{}, errors.New("gagal membuat lapangan")
	}

	courtID, err := result.LastInsertId()
	if err != nil {
		return CourtResponse{}, errors.New("gagal mendapatkan ID lapangan")
	}

	return CourtResponse{
		ID:           courtID,
		OwnerID:      ownerID,
		Name:         req.Name,
		Description:  req.Description,
		Address:      req.Address,
		City:         req.City,
		PricePerHour: req.PricePerHour,
		Status:       "active",
		AvgRating:    nil,
		TotalReviews: 0,
		CreatedAt:    time.Now(),
		Images:       []CourtImageResp{},
	}, nil
}

func GetCourtService(ctx context.Context, q *sqlc_db.Queries, id int64) (CourtResponse, error) {
	court, err := q.GetCourtByID(ctx, id)
	if err != nil {
		return CourtResponse{}, errors.New("lapangan tidak ditemukan")
	}

	images, _ := q.GetCourtImages(ctx, id)
	imageList := make([]CourtImageResp, 0, len(images))
	for _, img := range images {
		imageList = append(imageList, CourtImageResp{
			ID:        img.ID,
			ImageURL:  img.ImageUrl,
			IsPrimary: img.IsPrimary,
		})
	}

	var avgRating *float64
	if court.AvgRating != 0 {
		avgRating = &court.AvgRating
	}

	return CourtResponse{
		ID:           court.ID,
		OwnerID:      court.OwnerID,
		Name:         court.Name,
		Description:  nullStringToString(court.Description),
		Address:      nullStringToString(court.Address),
		City:         nullStringToString(court.City),
		PricePerHour: court.PricePerHour,
		Status:       string(court.Status),
		AvgRating:    avgRating,
		TotalReviews: court.TotalReviews,
		CreatedAt:    court.CreatedAt,
		Images:       imageList,
	}, nil
}

func ListCourtsService(ctx context.Context, q *sqlc_db.Queries, city string) ([]CourtResponse, error) {
	var courts []sqlc_db.Court
	var err error
	if city != "" {
		courts, err = q.ListCourtsByCity(ctx, toNullString(city))
	} else {
		courts, err = q.ListCourts(ctx)
	}
	if err != nil {
		return nil, errors.New("gagal mendapatkan daftar lapangan")
	}

	resp := make([]CourtResponse, 0, len(courts))
	for _, court := range courts {
		images, _ := q.GetCourtImages(ctx, court.ID)
		imageList := make([]CourtImageResp, 0, len(images))
		for _, img := range images {
			imageList = append(imageList, CourtImageResp{
				ID:        img.ID,
				ImageURL:  img.ImageUrl,
				IsPrimary: img.IsPrimary,
			})
		}
		var avgRating *float64
		if court.AvgRating != 0 {
			avgRating = &court.AvgRating
		}
		resp = append(resp, CourtResponse{
			ID:           court.ID,
			OwnerID:      court.OwnerID,
			Name:         court.Name,
			Description:  nullStringToString(court.Description),
			Address:      nullStringToString(court.Address),
			City:         nullStringToString(court.City),
			PricePerHour: court.PricePerHour,
			Status:       string(court.Status),
			AvgRating:    avgRating,
			TotalReviews: court.TotalReviews,
			CreatedAt:    court.CreatedAt,
			Images:       imageList,
		})
	}

	return resp, nil
}

func ListOwnerCourtsService(ctx context.Context, q *sqlc_db.Queries, ownerID int64) ([]CourtResponse, error) {
	courts, err := q.ListCourtsByOwner(ctx, ownerID)
	if err != nil {
		return nil, errors.New("gagal mendapatkan daftar lapangan")
	}

	resp := make([]CourtResponse, 0, len(courts))
	for _, court := range courts {
		images, _ := q.GetCourtImages(ctx, court.ID)
		imageList := make([]CourtImageResp, 0, len(images))
		for _, img := range images {
			imageList = append(imageList, CourtImageResp{
				ID:        img.ID,
				ImageURL:  img.ImageUrl,
				IsPrimary: img.IsPrimary,
			})
		}
		var avgRating *float64
		if court.AvgRating != 0 {
			avgRating = &court.AvgRating
		}
		resp = append(resp, CourtResponse{
			ID:           court.ID,
			OwnerID:      court.OwnerID,
			Name:         court.Name,
			Description:  nullStringToString(court.Description),
			Address:      nullStringToString(court.Address),
			City:         nullStringToString(court.City),
			PricePerHour: court.PricePerHour,
			Status:       string(court.Status),
			AvgRating:    avgRating,
			TotalReviews: court.TotalReviews,
			CreatedAt:    court.CreatedAt,
			Images:       imageList,
		})
	}
	return resp, nil
}

func UpdateCourtService(ctx context.Context, q *sqlc_db.Queries, id int64, ownerID int64, req UpdateCourtRequest) error {
	court, err := q.GetCourtByID(ctx, id)
	if err != nil {
		return errors.New("lapangan tidak ditemukan")
	}
	if court.OwnerID != ownerID {
		return errors.New("tidak diizinkan mengedit lapangan ini")
	}

	return q.UpdateCourt(ctx, sqlc_db.UpdateCourtParams{
		ID:           id,
		Name:         req.Name,
		Description:  toNullString(req.Description),
		Address:      toNullString(req.Address),
		City:         toNullString(req.City),
		PricePerHour: req.PricePerHour,
		Status:       sqlc_db.CourtsStatus(req.Status),
	})
}

func AddCourtImageService(ctx context.Context, q *sqlc_db.Queries, courtID int64, ownerID int64, req CourtImageRequest) error {
	court, err := q.GetCourtByID(ctx, courtID)
	if err != nil {
		return errors.New("lapangan tidak ditemukan")
	}
	if court.OwnerID != ownerID {
		return errors.New("tidak diizinkan")
	}

	if req.IsPrimary {
		if err := q.SetPrimaryImage(ctx, sqlc_db.SetPrimaryImageParams{
			ID:      0,
			CourtID: courtID,
		}); err != nil {
		}
	}

	_, err = q.CreateCourtImage(ctx, sqlc_db.CreateCourtImageParams{
		CourtID:   courtID,
		ImageUrl:  req.ImageURL,
		IsPrimary: req.IsPrimary,
	})
	return err
}

func DeleteCourtService(ctx context.Context, q *sqlc_db.Queries, id int64, ownerID int64) error {
	court, err := q.GetCourtByID(ctx, id)
	if err != nil {
		return errors.New("lapangan tidak ditemukan")
	}
	if court.OwnerID != ownerID {
		return errors.New("tidak diizinkan menghapus lapangan ini")
	}
	return q.DeleteCourt(ctx, id)
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func nullStringToString(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}
