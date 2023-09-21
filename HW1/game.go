package main

import "fmt"

type World struct {
	Rooms    map[string]*Room
	IsOpen   bool
	CanApply map[string]string
}

type Player struct {
	Inventory []string
	RoomCur   *Room
	BackPack  bool
}

func (player *Player) PutOn(obj string) error {
	if obj == "рюкзак" {
		player.BackPack = true
		return nil
	}
	if player.BackPack {
		player.Inventory = append(player.Inventory, obj)
		return nil
	}
	return fmt.Errorf("некуда класть")
}

type Room struct {
	Name            string
	HelloMsg        string
	NeighboursRooms []Room
	Table           []string
	Chair           []string
}

// ОСМОТРЕТЬСЯ
func (room *Room) LookAround() string {
	var msg string
	// кейс с кухней
	if room.Name == "кухня" {
		msg += "ты находишься на кухне, "
	}
	// стол
	if len(room.Table) != 0 {
		msg += "на столе: "
		for _, item := range room.Table {
			msg += item + ", "
		}
	}
	// кейс с кухней
	if !player.BackPack && room.Name == "кухня" {
		msg += "надо собрать рюкзак и идти в универ."
	} else if player.BackPack && room.Name == "кухня" {
		msg += "надо идти в универ."
	}
	// стул
	if len(room.Chair) != 0 {
		for _, item := range room.Chair {
			msg += "на стуле: " + item + ", "
		}
		msg = msg[0 : len(msg)-2]
		msg += "."
	} else if len(room.Table) != 0 && room.Name != "кухня" {
		msg = msg[0 : len(msg)-2]
		msg += "."
	}
	// пустая комната
	if len(room.Chair) == 0 && len(room.Table) == 0 {
		msg += "пустая комната."
	}
	return msg + canGo(room)
}

// команда ПРИМЕНИТЬ
func (room *Room) Apply(atr []string) string {
	if _, ok := findObject(player.Inventory, atr[0]); ok {
		if item, ok := world.CanApply[atr[0]]; ok && item == atr[1] {
			world.IsOpen = true
			return "дверь открыта"
		}
		return "не к чему применить"
	}
	return "нет предмета в инвентаре - " + atr[0]
}

// команда ИДТИ
func (room *Room) GoTo(r *Room) string {
	if findRoom(room.NeighboursRooms, r) {
		if r.Name == "улица" && !world.IsOpen {
			return "дверь закрыта"
		} else {
			player.RoomCur = r
			return r.HelloMsg + canGo(r)
		}
	}
	return "нет пути в " + r.Name
}

// команда ВЗЯТЬ obj
func (room *Room) Take(obj string) string {
	if i, ok := findObject(room.Table, obj); ok {
		if er := player.PutOn(obj); er != nil {
			return er.Error()
		} else {
			room.Table = remove(room.Table, i)
			return "предмет добавлен в инвентарь: " + obj
		}
	} else if i, ok := findObject(room.Chair, obj); ok {
		if er := player.PutOn(obj); er != nil {
			return er.Error()
		} else {
			room.Chair = remove(room.Chair, i)
			return "предмет добавлен в инвентарь: " + obj
		}
	}
	return "нет такого"
}

// команда НАДЕТЬ
func (room *Room) Put(obj string) string {
	if err := player.PutOn(obj); err != nil {
		return err.Error()
	}
	i, _ := findObject(room.Chair, obj)
	room.Chair = remove(room.Chair, i)
	return "вы надели: " + obj
}

// функции утилиты
func findRoom(rooms []Room, room *Room) bool {
	for _, r := range rooms {
		if r.Name == room.Name {
			return true
		}
	}
	return false
}

func findObject(objs []string, obj string) (int, bool) {
	for i, o := range objs {
		if o == obj {
			return i, true
		}
	}
	return -1, false
}

func remove(objs []string, i int) []string {
	objs[i] = objs[len(objs)-1]
	return objs[:len(objs)-1]
}

func canGo(room *Room) string {
	msg := " можно пройти - "
	if room.Name == "улица" {
		msg += "домой"
		return msg
	} else {
		for _, loc := range room.NeighboursRooms {
			msg += loc.Name + ", "
		}
		return msg[0 : len(msg)-2]
	}
}
