package doc

import "github.com/olivere/elastic/v7"

type Group struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	GroupID    int64  `json:"group_id"`
	MessageUID string `json:"message_uid"`
	Content    string `json:"content"`
}

func (Group) GetIndexName() string {
	return "chat_group"
}

func (Group) GetMapping() string {
	dramaMapping := `{
    "mappings":{
        "properties":{
            "user_id":{
                "type":"integer"
            },
            "group_id":{
                "type":"integer"
            },
            "message_uid":{
                "type":"text"
            },
            "content":{
                "type":"text",
                "analyzer":"ik_max_word"
            }
        }
    }
}`

	return dramaMapping
}

type GroupMessageSearch struct {
	Keyword      string // 搜索
	UserId       int64  // 用户id
	TargetUserID int64
}

func (s *GroupMessageSearch) GetQuery() *elastic.BoolQuery {
	// match bool 复合查询
	q := elastic.NewBoolQuery()

	if s.Keyword != "" { // 搜索 名称, 简介
		q = q.Must(elastic.NewMultiMatchQuery(s.Keyword, "name", "introduce", "videos.name", "videos.introduce"))
	}

	if s.UserId > 0 { // 搜索用户
		q = q.Filter(elastic.NewTermQuery("user_id", s.UserId))
	}

	if s.TargetUserID > 0 { // 搜索用户
		q = q.Filter(elastic.NewTermQuery("target_user_id", s.TargetUserID))
	}

	return q
}
