package stortypes

// type Slot struct {
// 	ID          int
// 	Description string
// 	Banners     []int
// }

type Banner struct {
	ID          int
	Description string
}

// type Group struct {
// 	ID          int
// 	Description string
// }

type BunnerBody struct {
	BannerID int `db:"banner_id" json:"bannerId"`
	SlotID   int `db:"slot_id" json:"slotId"`
}

type DisplayBannerBody struct {
	SlotID  int
	GroupID int
}

type AddClickBody struct {
	BannerID int `db:"banner_id" json:"bannerId"`
	SlotID   int `db:"slot_id" json:"slotId"`
	GroupID  int `db:"group_id" json:"groupId"`
}

type Storage interface {
	Connect() error
	Close() error
	AddBanner(bannerID, slotID int) error
	DeleteBanner(bannerID, slotID int) error
	AddClick(bunnerID, slotID, groupID int) error
	DisplayBanner(slotID, groupID int) (Banner, error)
}
