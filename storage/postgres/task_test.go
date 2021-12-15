package postgres

import (
	"reflect"
	"testing"
	"time"

	pb "github.com/husanmusa/to-do-service/genproto"
)

func TestTaskRepo_Create(t *testing.T) {
	tests := map[string]struct {
		input pb.Task
		want  pb.Task
	}{
		"successful": {
			input: pb.Task{
				Assignee: "Assignee",
				Title:    "Title",
				Summary:  "Summary",
				Deadline: "2021-11-16",
				Status:   "active",
			},
			want: pb.Task{
				Assignee:  "Assignee",
				Title:     "Title",
				Summary:   "Summary",
				Deadline:  "2021-11-16T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-16",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
			got.Id = 0
			got.CreatedAt = "2021-12-16"
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Get(t *testing.T) {
	tests := map[string]struct {
		input int64
		want  pb.Task
	}{
		"successful": {
			input: 21,
			want: pb.Task{
				Id:        21,
				Assignee:  "Assignee",
				Title:     "Title",
				Summary:   "Summary",
				Deadline:  "2021-11-16T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-16T15:54:10.829915Z",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_List(t *testing.T) {
	tests := map[string]struct {
		inputP int64
		inputL int64
		want   []pb.Task
	}{
		"successful": {
			inputP: 1,
			inputL: 2,
			want: []pb.Task{
				{
					Id:        20,
					Assignee:  "Assignee2",
					Title:     "Title2",
					Summary:   "Summary2",
					Deadline:  "2021-11-15T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-16",
				},
				{
					Id:        21,
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-11-16T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-16",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _, err := pgRepo.List(tc.inputP, tc.inputL)
			for i := range got {
				if err != nil {
					t.Fatalf("%s: expected: %v, got: %v", name, tc.want[i], *got[i])
				}
				got[i].CreatedAt = "2021-12-16"
				if !reflect.DeepEqual(tc.want[i], *got[i]) {
					t.Fatalf("%s: expected: %v, got: %v", name, tc.want[i], *got[i])
				}
			}
		})
	}
}

func TestTaskRepo_Update(t *testing.T) {
	tests := map[string]struct {
		input pb.Task
		want  pb.Task
	}{
		"successful": {
			input: pb.Task{
				Id:        22,
				Assignee:  "Assignee3",
				Title:     "Title3",
				Summary:   "Summary3",
				Deadline:  "2021-11-16",
				Status:    "active",
				CreatedAt: "2021-11-16",
			},
			want: pb.Task{
				Id:        22,
				Assignee:  "Assignee3",
				Title:     "Title3",
				Summary:   "Summary3",
				Deadline:  "2021-11-16T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-11-16",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := pgRepo.Update(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
			got.CreatedAt = "2021-11-16"

			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	tests := map[string]struct {
		input int64
		want  pb.Task
	}{
		"successful": {
			input: 23,
			want:  pb.Task{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := pgRepo.Delete(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, nil)
			}
			//if !reflect.DeepEqual(nil, nil) {
			//	t.Fatalf("%s: expected: %v, got: %v", name, tc.want, nil)
			//}
		})
	}
}

func TestTaskRepo_ListOverdue(t *testing.T) {
	timer, _ := time.Parse("2006-01-02", "2021-12-17")
	tests := map[string]struct {
		inputP int64
		inputL int64
		timer  time.Time
		want   []pb.Task
	}{
		"successful": {
			inputP: 1,
			inputL: 2,
			timer:  timer,
			want: []pb.Task{
				{
					Id:        20,
					Assignee:  "Assignee2",
					Title:     "Title2",
					Summary:   "Summary2",
					Deadline:  "2021-11-15T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-16",
				},
				{
					Id:        21,
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-11-16T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-16",
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _, err := pgRepo.ListOverdue(tc.inputP, tc.inputL, tc.timer)
			for i := range got {
				if err != nil {
					t.Fatalf("%s: expected: %v, got: %v", name, tc.want[i], *got[i])
				}
				got[i].CreatedAt = "2021-12-16"
				if !reflect.DeepEqual(tc.want[i], *got[i]) {
					t.Fatalf("%s: expected: %v, got: %v", name, tc.want[i], *got[i])
				}
			}
		})
	}
}
