package migrtations

type Team struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;" json:"id"`
	NameTeam string `json:"name_team"`
}
