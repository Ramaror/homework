package main

import (
	botPkg "homework/internal/pkg/bot"
	cmdAddPkg "homework/internal/pkg/bot/command/add"
	cmdHelpPkg "homework/internal/pkg/bot/command/help"
	userPkg "homework/internal/pkg/core/user"
	"log"
)

func main() {

	var user userPkg.Interface
	{
		user = userPkg.New()
	}

	var bot botPkg.Interface
	{
		bot = botPkg.MustNew()
		commandAdd := cmdAddPkg.New(user)
		bot.RegisterHandler(commandAdd)
		commandHelp := cmdHelpPkg.New(map[string]string{
			commandAdd.Name(): commandAdd.Description(),
		})
		bot.RegisterHandler(commandHelp)
	}

	if err := bot.Run(); err != nil {
		log.Panic(err)
	}
}
