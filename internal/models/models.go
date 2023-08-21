package models

import (
	"gorm.io/datatypes"
	_ "gorm.io/datatypes"
	"gorm.io/gorm"
)

// TODO: remove inherit from gorm.Model if we don't want to use soft delete
type ScheduleModel struct {
	gorm.Model
	Name        string                       ` json:"name" gorm:"size:191;unique"`
	DisplayName string                       `json:"displayname"`
	TimeZone    string                       `json:"timezone"`
	Schedule    datatypes.JSONType[Schedule] `json:"schedule"`
}

type Schedule struct {
	Dtype   string  `json:"dtype"`
	Corder  bool    `json:"Corder"`
	Shape   []int   `json:"Shape"`
	NdArray [][]int `json:"__ndarray__"`
}

type Policy struct {
	gorm.Model
	Name         string                                 `gorm:"size:191;unique" json:"name"`
	DisplayName  string                                 `json:"displayname"`
	Projects     datatypes.JSONSlice[string]            `json:"projects"`
	Tags         datatypes.JSONSlice[map[string]string] `json:"tags"`
	ScheduleName string                                 `json:"schedulename"`
}
