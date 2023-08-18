package models

type ScheduleModel struct {
	Name        string `gorm:"primaryKey"`
	DisplayName string
	TimeZone    string
	Schedule    Schedule
}

type Schedule struct {
	Dtype   string
	Corder  bool
	Shape   []int
	NdArray [][]int
}

type PolicyModel struct {
	Name        string `gorm:"primaryKey"`
	DisplayName string
	Accounts    []string
	Tags        []map[string]string
	Schedule    string
}
