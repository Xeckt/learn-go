package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

func main() {
	discord, err := discordgo.New("Bot " + "authentication token")
	if err != nil {
		log.Fatal(err)
	}

	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = discord.Close()
	if err != nil {
		log.Printf("could not close session: %s", err)
	}
}
