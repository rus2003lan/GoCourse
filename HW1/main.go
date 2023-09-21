package main

import (
	"strings"
)

var (
	world  World
	player Player
)

func main() {
}

func initGame() {
	/*
		эта функция инициализирует игровой мир - все команты
		если что-то было - оно корректно перезатирается
	*/
	player = *new(Player)
	world = *new(World)
	corridor := Room{Name: "коридор", HelloMsg: "ничего интересного."}
	kitchen := Room{Name: "кухня", HelloMsg: "кухня, ничего интересного.", NeighboursRooms: []Room{corridor}, Table: []string{"чай"}}
	room := Room{Name: "комната", HelloMsg: "ты в своей комнате.", NeighboursRooms: []Room{corridor}, Table: []string{"ключи", "конспекты"}, Chair: []string{"рюкзак"}}
	street := Room{Name: "улица", HelloMsg: "на улице весна.", NeighboursRooms: []Room{corridor}}
	corridor.NeighboursRooms = append(corridor.NeighboursRooms, kitchen, room, street)
	world.Rooms = map[string]*Room{"комната": &room, "кухня": &kitchen, "коридор": &corridor, "улица": &street}
	world.IsOpen = false
	world.CanApply = map[string]string{"ключи": "дверь"}
	player.RoomCur = &kitchen
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/

	atributes := strings.Split(command, " ")
	mainCommand := atributes[0]
	curRoom := player.RoomCur

	switch mainCommand {
	case "осмотреться":
		return curRoom.LookAround()
	case "применить":
		return curRoom.Apply(atributes[1:])
	case "идти":
		return curRoom.GoTo(world.Rooms[atributes[1]])
	case "взять":
		return curRoom.Take(atributes[1])
	case "надеть":
		return curRoom.Put(atributes[1])
	default:
		return "неизвестная команда"
	}
}
