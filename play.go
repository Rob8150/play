package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cred struct {
	Name    string `json: "Name"`
	Gitcode string `json: "Gitcode"`
	SSH     string `json: "SSH"`
	Project string `json: "Project"`
	Year    int    `json: "Year"`
	Month   int    `json: "Month"`
	Day     int    `json: "Day"`
}

func main() {
	//var (
	//	name    string = "Robert"
	//	gitcode string = "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"
	//)

	data := Cred{
		Name:    "",
		Gitcode: "",
		SSH:     "",
		Project: "",
		Year:    2022,
		Month:   1,
		Day:     1,
	}

	fnam := ".git.json"
	//data.savjson(data, fnam)
	data.loajson(fnam)

	//project := getInput("Project ")

	currentTime := time.Now()
	t1 := Date(data.Year, data.Month, data.Day)
	t2 := currentTime
	//t1 := Date(2022, 9, 1)
	days := t2.Sub(t1).Hours() / 24
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	//fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	fmt.Println("Welcome " + data.Name)
	fmt.Print("Elapsed time in Days : ")
	fmt.Println(math.Round(days))
	fmt.Println("You are working on project " + data.Project)
	fmt.Println(" ")

	fmt.Println(".git.json")
	fmt.Println("----GO MOD --------------------------------------")
	fmt.Print("")
	fmt.Println("cd ~/gocode/src/github.com/Rob8150/" + data.Project)
	fmt.Println("go mod init github.com/Rob8150/" + data.Project)
	fmt.Println("go get github.com/charmbracelet/lipgloss")
	fmt.Println(" ")

	fmt.Println("----GitHub---------------------------------------")
	fmt.Println("git branch -m master main")
	fmt.Println("git remote remove origin")
	fmt.Println("git remote add origin https://" + data.Gitcode + "@github.com/Rob8150/" + data.Project)
	fmt.Println("git push --set-upstream origin main")

	fmt.Println("cd gocode/src/github.com/Rob8150/" + data.Project)
	fmt.Println("-----Transport Remote-----------------------------")
	fmt.Println(" ")
	fmt.Println("-----ZIP------------------------------------------")
	fmt.Println("cd ~/gocode/src/github.com/Rob8150/")
	fmt.Println("tar -cvzf " + data.Project + ".tar.gz " + data.Project)

	fmt.Println(" ")
	fmt.Println("-----Upstream ------------------------------------")
	fmt.Println("scp ~/gocode/src/github.com/Rob8150/" + data.Project + ".tar.gz " + data.SSH + ":gocode/src/github.com/Rob8150/" + data.Project + ".tar.gz")

	fmt.Println(" ")
	fmt.Println("-----DownStream-----------------------------------")
	fmt.Println("scp " + data.SSH + ":gocode/src/github.com/Rob8150/" + data.Project + ".tar.gz ~/gocode/src/github.com/Rob8150/" + data.Project + ".tar.gz")

	fmt.Println("ssh " + data.SSH)
	fmt.Println("cd ~/gocode/src/github.com/Rob8150/")
	fmt.Println("Extract")
	fmt.Println("tar -zxvf " + data.Project + ".tar.gz")
}

func (dat *Cred) savjson(d Cred, fnam string) {
	dat.Name = d.Name
	dat.Gitcode = d.Gitcode
	dat.SSH = d.SSH
	dat.Project = d.Project
	dat.Year = d.Year
	dat.Month = d.Month
	dat.Day = d.Day

	file, _ := json.MarshalIndent(dat, "", " ")
	_ = ioutil.WriteFile(fnam, file, 0644)
}

func (dat *Cred) loajson(fnam string) {
	raw, err := ioutil.ReadFile(fnam)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &dat)
}

func (dat *Cred) display() {
	fmt.Println(dat.Name)
	fmt.Println(dat.Gitcode)
	fmt.Println(dat.SSH)
	fmt.Println(dat.Project)
	fmt.Println(dat.Year)
	fmt.Println(dat.Month)
	fmt.Println(dat.Day)

}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func getInput(askfor string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Print(askfor)
	scanner.Scan()
	text = scanner.Text()
	return text
}

// getNum convert string to float64
func getNum(given string) float64 {
	var clean string
	var gotnum float64
	clean = strings.TrimSpace(given)
	gotnum, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		log.Fatal(err)
	}
	return gotnum
}
