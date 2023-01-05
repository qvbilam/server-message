package doc

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"net/http"
)

func SetHighlight(names []string, preTags, postTags string) *elastic.Highlight {
	if len(names) == 0 {
		return nil
	}

	h := NewHighlight()
	h.Fields(names...).PreTags(preTags).PostTags(postTags)
	return h.Client
}

func SetSort(field string) *elastic.SortInfo {
	if field == "" {
		return nil
	}
	var s string
	var a bool
	if string(field[0]) == "-" {
		s = field[0:]
		a = false
	} else {
		s = field
		a = true
	}

	c := NewSortClient()
	return c.Sort(s, a)
}

func ErrorNotfound() string {
	return fmt.Sprintf("elastic: Error %d (%s)", http.StatusNotFound, http.StatusText(http.StatusNotFound))
}
