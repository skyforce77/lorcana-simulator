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

// NewLuaState ! Have to be closed !
func NewLuaState(card *PlayingCard) *lua.LState {
	L := lua.NewState()

	L.SetGlobal("card", luar.New(L, card))
	L.SetGlobal("game", luar.New(L, card.game))
	L.SetGlobal("player", luar.New(L, card.owner))
	L.SetGlobal("on", L.NewFunction(card.luaListener()))
	L.SetGlobal("print", L.NewFunction(luaPrint))

	return L
}

func (action *CardAction) Execute(card *PlayingCard) {
	L := NewLuaState(card)
	defer L.Close()

	if err := L.DoString(action.Script); err != nil {
		panic(err)
	}
}

// EVENTS

type LuaEvent interface {
	ID() string
}

type LuaCancellableEvent struct {
	Name      string
	Cancelled bool
}

type LuaSingEvent struct {
	LuaCancellableEvent
}

func NewLuaSingEvent() *LuaSingEvent {
	return &LuaSingEvent{
		LuaCancellableEvent{
			"sing",
			false,
		},
	}
}

type LuaPlacedEvent struct {
	LuaCancellableEvent
}

func NewLuaPlacedEvent() *LuaPlacedEvent {
	return &LuaPlacedEvent{
		LuaCancellableEvent{
			"placed",
			false,
		},
	}
}

func (event *LuaCancellableEvent) ID() string {
	return event.Name
}

func (event *LuaCancellableEvent) SetCancelled(cancelled bool) {
	event.Cancelled = cancelled
}

func (event *LuaCancellableEvent) cancel() {
	event.Cancelled = true
}

func (event *LuaCancellableEvent) IsCancelled() bool {
	return event.Cancelled
}
