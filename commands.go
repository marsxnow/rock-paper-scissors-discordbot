package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Define command creation function
func CreateCommand(dg *discordgo.Session, appID string, guildID string) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "dream",
			Description: "Trigger the dream command",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "kitten",
			Description: "Trigger the kitten command",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "song",
			Description: "get the song of the day",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "corridos",
			Description: "natanaeaaaaaaa",
			Options:     []*discordgo.ApplicationCommandOption{},
		},
		{
			Name:        "lol",
			Description: "Get a player's win/loss stats",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "player",
					Description: "The name of the player",
					Required:    true,
				},
			},
		},
		//add more for later
	}
	for _, command := range commands {
		_, err := dg.ApplicationCommandCreate(appID, guildID, command)
		if err != nil {
			log.Fatalf("Unable to create command: %v", err)
		}
	}
}

// Define command handler function
func DreamCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		if data.Name == "dream" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "When the ants come at me, It'll take 10,000, 100,000 of them to take me down. So that's how miniscule you are to my size, right. My stature of intelligence, character and body and um, Reverence in this world. No man, because I'm gonna do movies, stand up comedy, everything all the shit, music, whatever the fuck e commerce, You don't understand that you're talking to like a Michealangelo of my time, right. Like I'm a genius, Albert Einstein level, History book maker. You're gonna be forgotten like the dust in the sand when you're in the fucking sahara, and there's a hundred million thousand billion sand particles, you're gonna be one of those, and I'm gonna be a staute erected in gold.",
				},
			})
		}
	}
}

func KittenCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		if data.Name == "kitten" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "🫵",
				},
			})
		}
	}
}

func CorridosCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		data := i.ApplicationCommandData()
		if data.Name == "corridos" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Muchos no lo saben, pero los. Elementos del Polka y Folclórica fueron introducidos a mexico por europeos en los años de 1830, él sonido de la guitarra y el acordeón,fueron adaptados por la gente del pueblo, adquiriendo como música propia, Artistas como Narciso Martines, recuerda haber escuchado la Polka cuando eran tan solo un niño, sin embargo, algunos Artistas reemplazaron los instrumentos de latón con otros sonidos, también empezaron a contar sus propias historias y emplearon, para dar noticia sobre la revolución, estas canciones se conocerían como los boleros mexicanos, una base para los géneros que vendrían,con el tiempo, la música norteña tocaba en tierra rurales,conversaría a funcionarse con los clásicos del bolero,esto dio un nueva forma de interpretar la música, conocida popularmente como,la música ranchera:Canciones de la tierra,décadas más tarde Artistas como Pedro Infante llevaron la música Ranchera a su centro de atención,sus baladas románticas se hicieron reconocer en todo el mundo como música propia de México, luego vinieron los Tigres del Norte y todo cambio,en los años 70 los Tigres del Norte popularizaron una canción ficticia sobre el narcotrafico, llamada contrato y traición, ese corrido fue exitoso y controversial,esto dio comienzo a el géneroNarcocorrido, los Tigres del Norte fueron uno de los pioneros en el género del Narcocorrido y aunque el apetito por su música era feroz, Mexico prohibió el género, empujando a los músicos a las ciudades fronterizas y finalmente,en Estados Unidos, luego vino Chalino Sánchez, un granjero con un gran corazón de acero y un amor por el Narcocorrido, con él trajo la música de estilo Bandas de Sinaloa, Sanchez escribió canciones hechas a medida para inmortalizar a los narcotraficantes,justo en esa época la ciudad de Los Ángeles estaba presenciando un nuevo movimiento de Gangsta Rap, los mexicoamericanos y los arreoamericanos,compartieron varios espacios en el centro de la ciudadi observando las culturas y los estilos musicales de los demás, al igual que el movimiento de Gangsta Rap, Sánchez cantaba sobre la violencia, las drogas y los carteles Mexicanos,Chalino Sánchez sería coronado como el Rey delcoronado como el Rey del Narcocorrido, mientras los corridos se congelaron en el tiempo el Gangsta Rap se generalizo y empezó a ser respaldado por las principales discográficas, convirtiéndose en sinónimo de Pop, hasta que se introdujo la música Trap, el espíritu de la música Trap,atrajo a los jóvenes artistas mexicanos, repasando las influencias fundamentales de los ritmos urbanos, la sensibilidad del Hip-Hop y la superposición de las letras de los Corridos, lo desglosaron a lo que ahora llamamos Corridos tumbados y ahí es donde entro yo",
				},
			})
		}
	}
}

// func SongOfTheDayCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	if i.Type == discordgo.InteractionApplicationCommand {
// 		data := i.ApplicationCommandData()
// 		if data.Name == "song" {
// 			songOfTheDayLink := "https://music.youtube.com/watch?v=2PVUZ5ZX79Q&si=E5-rQ8Y98YgZRXqI"

// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionResponseData{
// 					Content: songOfTheDayLink,
// 				},
// 			})
// 		}
// 	}
// }
