package roles

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pajbot/basecommand"
	"github.com/pajbot/pajbot2-discord/pkg"
	"github.com/pajbot/pajbot2-discord/pkg/commands"
)

func init() {
	commands.Register([]string{"$roles"}, New())
}

var _ pkg.Command = &Command{}

type Command struct {
	basecommand.Command
}

func New() *Command {
	return &Command{
		Command: basecommand.New(),
	}
}

func (c *Command) Run(s *discordgo.Session, m *discordgo.MessageCreate, parts []string) (res pkg.CommandResult) {
	res = pkg.CommandResultNoCooldown
	if m.Author.ID != "85699361769553920" {
		return
	}

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		fmt.Println("Error getting roles:", err)
		return
	}

	response := "```"
	for _, role := range roles {
		if role.Managed {
			continue
		}
		response += fmt.Sprintf("%s = %s\n", role.ID, role.Name)
	}
	response += "```"

	s.ChannelMessageSend(m.ChannelID, response)

	return
}

func (c *Command) Description() string {
	return c.Command.Description
}