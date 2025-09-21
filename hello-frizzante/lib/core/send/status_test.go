package send

import (
	"testing"

	"main/lib/core/mock"
)

func TestStatus(t *testing.T) {
	client := mock.NewClient()
	Status(client, 400)
	if client.Status != 400 {
		t.Fatal("status should be 400")
	}
}
