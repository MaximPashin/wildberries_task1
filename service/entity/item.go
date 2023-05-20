package entity

type Item struct {
	ID         uint64 `json:"chrt_id" validate:"required"`
	TrackNum   string `json:"track_number"`
	Price      uint   `json:"price"`
	RID        string `json:"rid"`
	Name       string `json:"name"`
	Sale       uint   `json:"sale"`
	Size       string `json:"size"`
	TotalPrice uint   `json:"total_price"`
	NmID       uint   `json:"nm_id"`
	Brand      string `json:"brand"`
	Status     uint   `json:"status"`
}
