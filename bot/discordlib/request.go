package discordlib

type DiscordRequestDto struct {
	Type      int                   `json:"type"`
	Data      DiscordDataRequestDto `json:"data"`
	Member    DiscordMemberDto      `json:"member"`
	GuildId   string                `json:"guild_id"`
	ChannelId string                `json:"channel_id"`
}

type DiscordDataRequestDto struct {
	Name     string      `json:"name"`
	Options  DataOptions `json:"options"`
	Resolved Resolved    `json:"resolved"`
}

type Resolved struct {
	Members Members `json:"members"`
}
type Members map[string]Member
type Member struct {
	Nick string `json:"nick"`
}

type User struct {
	Id string `json:"id"`
}

type DataOptions []DataOption
type DataOption struct {
	Name    string
	Value   interface{}
	Options []map[string]interface{}
}

type DiscordMemberDto struct {
	Nick string `json:"nick"`
	User User   `json:"user"`
}
