package stat

type ConfLever struct{
	No    			int16   `gorm:"primary_key" json:"no"`
	HeroFight		int		`json:"hero_fight"`
	SoldierNum      int		`json:"soldier_num"`
	HeroNum         int16	`json:"hero_num"`
}