package store

import (
	"github.com/nduyphuong/gorya/internal/models"
)

type Interface interface {
	SavePolicy(policy models.PolicyModel) error
	SaveSchedule(schedule models.ScheduleModel) error
}
