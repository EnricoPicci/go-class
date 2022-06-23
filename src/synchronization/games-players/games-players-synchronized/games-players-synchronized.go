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
	muPlayers sync.Mutex
	Players   map[string]*Player
	muGames   sync.Mutex
	Games     map[string]*Game
}

func NewGamePlatform() *GamePlatform {
	return &GamePlatform{Players: map[string]*Player{}, Games: map[string]*Game{}}
}

func (gamePlatform *GamePlatform) AddAndJoin(p *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	// lock the players to add a new player
	fmt.Println("1 AddAndJoin - muPlayers")
	gamePlatform.muPlayers.Lock()
	time.Sleep(1 * time.Millisecond)
	fmt.Println("2 AddAndJoin - muPlayers")
	//
	gamePlatform.Players[p.name] = p
	//
	fmt.Println("3 AddAndJoin - muPlayers")
	gamePlatform.muPlayers.Unlock()
	fmt.Println("4 AddAndJoin - muPlayers")

	// lock the games to add a new player to each game
	fmt.Println("5 AddAndJoin - muGames")
	gamePlatform.muGames.Lock()
	fmt.Println("6 AddAndJoin - muGames")
	//
	for _, g := range gamePlatform.Games {
		g.players[p.name] = p
	}
	//
	fmt.Println("7 AddAndJoin - muGames")
	gamePlatform.muGames.Unlock()
	fmt.Println("8 AddAndJoin - muGames")
}
func (gamePlatform *GamePlatform) Ban(p *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	// lock the games to remove the player from each game
	fmt.Println("1 Ban - muGames")
	gamePlatform.muGames.Lock()
	time.Sleep(1 * time.Millisecond)
	fmt.Println("2 Ban - muGames")
	//
	for _, g := range gamePlatform.Games {
		delete(g.players, p.name)
	}
	//
	fmt.Println("3 Ban - muGames")
	gamePlatform.muGames.Unlock()
	fmt.Println("4 Ban - muGames")

	// lock the players to remove the player from the app
	fmt.Println("5 Ban - muPlayers")
	gamePlatform.muPlayers.Lock()
	fmt.Println("6 Ban - muPlayers")
	//
	delete(gamePlatform.Players, p.name)
	//
	fmt.Println("7 Ban - muPlayers")
	gamePlatform.muPlayers.Unlock()
	fmt.Println("8 Ban - muPlayers")
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
		fmt.Printf("%p - %p - %v - %v\n", &gPlat.muGames, &gPlat.muPlayers, &gPlat.muGames, &gPlat.muPlayers)
		time.Sleep(1 * time.Millisecond)
		launch_AddAndJoin_Ban(p_1, p_2)
	}
	fmt.Println("Terminating games players synchronized")
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
