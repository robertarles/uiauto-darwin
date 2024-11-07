package main

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"log"
	"os/exec"
)

func main() {
	add()

	//low()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	hook.Register(hook.KeyDown, []string{"b", "command", "option"}, func(e hook.Event) {
		fmt.Println("[DEBUG] Web [B]rowser")
		cmd := exec.Command("open", "-a", "Safari.app")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to open Safari.app", err)
		}
	})

	hook.Register(hook.KeyDown, []string{"t", "command", "option"}, func(e hook.Event) {
		fmt.Println("[DEBUG] [T]erminal")
		cmd := exec.Command("open", "-a", "Kitty.app")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to open Kitty.app", err)
		}
	})

	hook.Register(hook.KeyDown, []string{"a", "command", "option"}, func(e hook.Event) {
		fmt.Println("[DEBUG] [A]I")
		cmd := exec.Command("open", "-a", "chatgpt(web).app")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to open chatgpt(web).app", err)
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}
