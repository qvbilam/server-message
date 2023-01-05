package doc

import "github.com/olivere/elastic/v7"

type SortClient struct {
}

func NewSortClient() *SortClient {
	return &SortClient{}
}

func (c *SortClient) Sort(field string, ascending bool) *elastic.SortInfo {
	s := elastic.SortInfo{
		Field:     field,
		Ascending: ascending,
	}
	return &s
}
