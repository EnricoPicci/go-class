package main

// https://www.lurklurk.org/effective-rust/deadlock.html

import (
	"fmt"
	"sync"
	"time"
)

type Game struct {
	name    string
	players map[string]*Player
}

func NewGame(name string) *Game {
	return &Game{name: name, players: make(map[string]*Player)}
}

type Player struct{ name string }

func NewPlayer(name string) *Player {
	return &Player{name}
}

type GamePlatform struct {
	Players map[string]*Player
	Games   map[string]*Game
}

func NewGamePlatform() *GamePlatform {
	return &GamePlatform{Players: map[string]*Player{}, Games: map[string]*Game{}}
}

func (gamePlatform *GamePlatform) AddAndJoin(p *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	gamePlatform.Players[p.name] = p

	for _, g := range gamePlatform.Games {
		g.players[p.name] = p
	}
}
func (gamePlatform *GamePlatform) Ban(p *Player, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, g := range gamePlatform.Games {
		delete(g.players, p.name)
	}

	delete(gamePlatform.Players, p.name)
}

var gPlat = NewGamePlatform()

func main() {
	loadWithSomeData(gPlat)
	counter := 0
	for i := 0; i < 10; i++ {
		p_1 := NewPlayer(fmt.Sprintf("the first player %v", i))
		p_2 := NewPlayer(fmt.Sprintf("the second player %v", i))
		counter++
		fmt.Println(">>>>>>>>>> Iteration", counter)
		time.Sleep(1 * time.Millisecond)
		launch_AddAndJoin_Ban(p_1, p_2)
	}
	fmt.Println("Terminating games players data races")
}

func launch_AddAndJoin_Ban(p_1 *Player, p_2 *Player) {
	var wg sync.WaitGroup
	wg.Add(2)
	go gPlat.AddAndJoin(p_1, &wg)
	go gPlat.Ban(p_2, &wg)
	wg.Wait()
}

func loadWithSomeData(gPlat *GamePlatform) {
	for i := 0; i < 100; i++ {
		gName := fmt.Sprintf("Game-%v", i)
		pName := fmt.Sprintf("Player-%v", i)
		gPlat.Games[gName] = NewGame(gName)
		gPlat.Players[gName] = NewPlayer(pName)
	}
}
