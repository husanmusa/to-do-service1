package repo

import (
	pb "github.com/husanmusa/to-do-service/genproto"
	"time"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id int64) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id int64) error
	ListOverdue(page, limit int64, time time.Time) ([]*pb.Task, int64, error)
}
