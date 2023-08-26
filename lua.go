package main

import (
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"log"
)

func (card *PlayingCard) luaListener() func(L *lua.LState) int {
	return func(L *lua.LState) int {
		eventType := L.ToString(1)
		function := L.ToFunction(2)
		card.Listeners[eventType] = function
		return 0
	}
}

func luaPrint(L *lua.LState) int {
	str := L.ToString(1)
	log.Printf("[LUA DEBUG] %s\n", str)
	return 0
}

func (action *CardAction) Execute(card *PlayingCard) {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("card", luar.New(L, card))
	L.SetGlobal("game", luar.New(L, card.game))
	L.SetGlobal("player", luar.New(L, card.owner))
	L.SetGlobal("on", L.NewFunction(card.luaListener()))
	L.SetGlobal("print", L.NewFunction(luaPrint))

	if err := L.DoString(action.Script); err != nil {
		panic(err)
	}
}
