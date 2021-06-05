package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	if len(feed.Items) == 0 {
		t.Errorf("Item was not added")
	}

}
func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results ) == 0 {
		t.Errorf("Item was not added")
	}
}