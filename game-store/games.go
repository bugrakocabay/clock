package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func newGame(id, price int, name, genre string) (new game) {
	new = game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
	return
}

func addGame(arr []game, new game) []game {
	arr = append(arr, new)
	return arr
}

func load() (games []game) {
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "minecraft", "sandbox"))
	games = addGame(games, newGame(3, 10, "dota", "moba"))

	return
}

func indexById(arr []game) (byID map[int]game) {
	byID = make(map[int]game)
	for _, g := range arr {
		byID[g.id] = g
	}

	return
}

func printGame(g game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}