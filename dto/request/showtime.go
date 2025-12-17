package request

type CreateShowtimeRequest struct {
	MovieID   int    `json:"movie_id" binding:"required"`
	StudioID  int    `json:"studio_id" binding:"required"`
	ShowDate  string `json:"show_date" binding:"required,datetime=2006-01-02"`
	StartTime string `json:"start_time" binding:"required,datetime=15:04"`
	Status    string `json:"status" binding:"required"`
}

type UpdateShowtimeRequest struct {
	ID        int    `json:"id" binding:"required"`
	MovieID   int    `json:"movie_id" binding:"required"`
	StudioID  int    `json:"studio_id" binding:"required"`
	ShowDate  string `json:"show_date" binding:"required,datetime=2006-01-02"`
	StartTime string `json:"start_time" binding:"required,datetime=15:04"`
	Status    string `json:"status" binding:"required"`
}

type GetShowtimeRequest struct {
	ID int `uri:"id" binding:"required"`
}

type DeleteShowtimeRequest struct {
	ID int `uri:"id" binding:"required"`
}
