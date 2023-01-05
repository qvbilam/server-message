package doc

import "github.com/olivere/elastic/v7"

type HighlighterClient struct {
	Client *elastic.Highlight
}

func NewHighlight() *HighlighterClient {
	c := elastic.NewHighlight()
	return &HighlighterClient{
		Client: c,
	}
}

func (c *HighlighterClient) Field(name string) *HighlighterClient {
	c.Client.Field(name)
	return c
}

func (c *HighlighterClient) Fields(names ...string) *HighlighterClient {
	var fs []*elastic.HighlighterField
	for _, name := range names {
		fs = append(fs, &elastic.HighlighterField{
			Name: name,
		})
	}

	c.Client.Fields(fs...)
	return c
}

func (c *HighlighterClient) PreTags(preTags ...string) *HighlighterClient {
	c.Client.PreTags(preTags...)
	return c
}

func (c *HighlighterClient) PostTags(postTags ...string) *HighlighterClient {
	c.Client.PostTags(postTags...)
	return c
}
