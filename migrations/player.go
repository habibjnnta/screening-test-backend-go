package migrtations

type Player struct {
	ID     uint   `gorm:"primaryKey;autoIncrement;"`
	Name   string `json:"name"`
	TeamId uint   `json:"team_id"`
	Teams  Team   `gorm:"foreignKey:team_id;references:id"`
}
