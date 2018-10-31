package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var usedword [][]string
var cpuword [][]string
var playerdata []playerData
var player1, player2 string
var tmptmpnum int

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/start/", start)
	http.HandleFunc("/game/", game)

	server.ListenAndServe()
}

func start(w http.ResponseWriter, r *http.Request) {

	//ファイルを開いて入力
	t, err := template.ParseFiles("start.html")
	if err != nil {
		panic(err)
	}

	player1 = r.FormValue("player1")
	player2 = r.FormValue("player2")

	fmt.Println("player1: ", player1)
	fmt.Println("player2: ", player2)

	t.Execute(w, "")

	w.Header().Set("Location", "http://127.0.0.1:8080/game/")
	w.WriteHeader(302)
}

func game(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("game.html")
	if err != nil {
		panic(err)
	}

	// 宣言
	var newword string
	// var str string
	var seacretalpha byte
	var lastalpha byte
	lastalpha = 'a'
	rand.Seed(time.Now().UTC().UnixNano())
	seacretalpha = numal(rand.Intn(26))

	usedword = make([][]string, 26)
	cpuword = make([][]string, 26)
	playerdata = make([]playerData, 2)
	playerdata[0].words = make([][]string, 26)
	playerdata[1].words = make([][]string, 26)
	for i := 0; i < 26; i++ {
		playerdata[0].words[i] = make([]string, 1000)
		playerdata[1].words[i] = make([]string, 1000)
		usedword[i] = make([]string, 1000)
		cpuword[i] = make([]string, 1000)
	}

	playerdata[0].name = player1
	playerdata[1].name = player2

	playerdata[0].name = "cpu"
	playerdata[1].name = "abc"

	// 入力されたプレイヤーがcpuであれば語彙を入れる.
	for i := 0; i < 2; i++ {
		playerdata[i].iscpu = userdata(playerdata[i].name)
		if playerdata[i].iscpu {
			voc(playerdata[i].name, playerdata[i].words)
		}
	}

	// しりとり処理
	tmptmpnum %= 2
	tmptmpnum++
	if playerdata[tmptmpnum].iscpu {
		newword = cpucalc(tmptmpnum, seacretalpha, lastalpha)
		lastalpha = newword[len(newword)-1]
		t.Execute(w, newword)
		fmt.Println(newword)
	} else {
		newword = r.FormValue("word")
		fmt.Println(newword)
		/*
			if existInUsed(newword) {
				str = fmt.Sprintf("%s は既に使われているよ!\n", newword)
			} else if lastalpha != newword[0] {
				str = fmt.Sprintf("次は %c から始まる言葉だよ!\n", lastalpha)
			} else {
				addWord(newword)
				break
			}
			t.Execute(w, str)
		*/
		lastalpha = newword[len(newword)-1]
		if lastalpha == seacretalpha {
			t.Execute(w, "You lose!!")
		}
	}
}
