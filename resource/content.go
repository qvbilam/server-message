package resource

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Extra  string `json:"extra"`
}

// Text 文本消息
type Text struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	User    User   `json:"user"`
	Extra   string `json:"extra"`
}

// Image 文件消息
type Image struct {
	Type    string `json:"type"`
	Content string `json:"content"` // 缩略图
	Url     string `json:"url"`
	User    User   `json:"user"`
	Extra   string `json:"extra"`
}

// GIF GIF消息
type GIF struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int    `json:"size"`
	User   User   `json:"user"`
	Extra  string `json:"extra"`
}

// Voice 音频消息
type Voice struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Second int    `json:"second"`
	User   User   `json:"user"`
	Extra  string `json:"extra"`
}

// Video 视频消息
type Video struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"` // 缩略图
	Url     string `json:"url"`
	Size    string `json:"size"`
	Second  int    `json:"second"`
	User    User   `json:"user"`
	Extra   string `json:"extra"`
}

// File 文件消息
type File struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
	Url   string `json:"url"`
	User  User   `json:"user"`
	Extra string `json:"extra"`
}

// LBS 位置消息
type LBS struct {
	Type      string `json:"type"`
	Content   string `json:"content"` // 位置缩略图
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Poi       string `json:"poi"`
	User      User   `json:"user"`
	Extra     string `json:"extra"`
}
