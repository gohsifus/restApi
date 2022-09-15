package inmemory

import (
	"github.com/stretchr/testify/assert"
	"restApi/domain/entity"
	"testing"
	"time"
)

func TestInMemoryEventRepo_Create(t *testing.T) {
	im, _ := NewInMemoryRepo()

	newEvent := entity.NewEvent(time.Now(), "asd", "asd")

	_, err := im.Create(newEvent)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(im.store))
}

func TestInMemoryEventRepo_Delete(t *testing.T) {
	im, _ := NewInMemoryRepo()

	im.store[0] = *entity.NewEvent(time.Now(), "asd", "asd")

	im.Delete(0)

	assert.Equal(t, 0, len(im.store))
}

func TestInMemoryEventRepo_Update(t *testing.T) {
	im, _ := NewInMemoryRepo()

	im.store[0] = *entity.NewEvent(time.Now(), "asd", "asd")

	im.Update(0, &entity.Event{
		Name: "qwe",
	})

	assert.Equal(t, im.store[0].Name, "qwe")
}

func TestInMemoryEventRepo_GetEventsByDateInterval(t *testing.T) {
	im, _ := NewInMemoryRepo()

	im.store[0] = *entity.NewEvent(strToTime("2022-09-09"), "asd", "asd")
	im.store[1] = *entity.NewEvent(strToTime("2022-09-10"), "qwe", "qwe")
	im.store[2] = *entity.NewEvent(strToTime("2022-07-12"), "aboba", "aboba")

	testTable := []struct {
		name  string
		from  string
		to    string
		isErr bool
		want  int
	}{
		{
			name: "Get by day",
			from: "2022-09-09",
			to:   "2022-09-09",
			want: 1,
		},
		{
			name: "Get by few days",
			from: "2022-09-09",
			to:   "2022-09-12",
			want: 2,
		},
		{
			name: "Get with incorrect dates",
			from: "2022-09-09",
			to:   "2022-07-12",
			isErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := im.GetEventsByDateInterval(strToTime(testCase.from), strToTime(testCase.to))
			t.Log(got)

			if testCase.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, len(got))
			}
		})
	}
}

func strToTime(str string) time.Time {
	date, _ := time.Parse("2006-01-02", str)
	return date
}
