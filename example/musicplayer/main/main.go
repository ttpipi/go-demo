package main

import (
	"bufio"
	"fmt"
	"golang/example/musicplayer/musicmgr"
	"golang/example/musicplayer/play"
	"os"
	"strconv"
	"strings"
)

var mgr *musicmgr.MusicManager

func handleManagerCommand(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < mgr.Len(); i++ {
			m, _ := mgr.Get(i)
			fmt.Println(i, m.Name, m.Artist, m.Source, m.Type)
		}
	case "add":
		if len(tokens) != 6 {
			fmt.Println("usage: mgr add <name> <artist> <source> <type>")
			return
		}
		mgr.Add(&musicmgr.MusicInfo{
			Id:     strconv.Itoa(mgr.Len()),
			Name:   tokens[2],
			Artist: tokens[3],
			Source: tokens[4],
			Type:   tokens[5],
		})
		fmt.Println("add success.")

	case "remove":
		if len(tokens) != 3 {
			fmt.Println("usage: mgr remove <id>")
			return
		}
		id, err := strconv.Atoi(tokens[2])
		if err != nil {
			fmt.Println("usage: mgr remove <id>")
			return
		}
		m := mgr.Remove(id)
		if m != nil {
			fmt.Println("remove the music:", m.Name)
		}

	default:
		fmt.Println("unsupport option", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("usage: play <name>")
		return
	}

	music := mgr.Find(tokens[1])
	if music == nil {
		fmt.Println("not exist the music", tokens[0])
		return
	}

	play.Play(music.Source, music.Type)
}

func main() {
	fmt.Println(`
		enter following commands to control the player:
		mgr list -- view the music list.
		mgr add <name> <artist> <source> <type> -- add a music to the list
		mgr remove <id> -- remove the specified music
		play <name> -- play the specified music
		q -- quit the player`)

	mgr = musicmgr.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" {
			break
		}
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "mgr":
			handleManagerCommand(tokens)
		case "play":
			handlePlayCommand(tokens)
		default:
			fmt.Println("unknow command", tokens[0])
		}
	}
}
