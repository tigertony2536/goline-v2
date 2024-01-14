package notification_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tigertony2536/go-line-notify/notiApp/notification"
)

func TestGetThisWeekTasks(t *testing.T) {
	expectRowsNumber := 5
	expectID := []int{44, 45, 46, 48, 50}

	t.Run("Test Get Daily Noti", func(t *testing.T) {
		noti, err := notification.GetThisWeekTasks()

		notiID := []int{}

		for _, n := range noti.Tasks {
			notiID = append(notiID, n.ID)
		}

		assert.Equalf(t, expectRowsNumber, len(notiID), "Expected %d row got %d row")
		assert.Equalf(t, expectID, notiID, "Expected %d got %d")
		assert.NoError(t, err, "No Error")
	})
}
