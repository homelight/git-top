package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/urfave/cli/v2"
)

type FileCount struct {
	name  string
	count int
}

type FileCounts []FileCount

func (p FileCounts) Len() int           { return len(p) }
func (p FileCounts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p FileCounts) Less(i, j int) bool { return p[i].count > p[j].count }

func main() {
	app := &cli.App{
		Name:            "git top",
		Usage:           "view the top analytics for a git log.",
		HideHelpCommand: true,
		HelpName:        "git top",
		Commands: []*cli.Command{
			{
				Name:  "files",
				Usage: "view the most changed files for a time period.",
				Flags: []cli.Flag{
					&cli.TimestampFlag{
						Name:        "after",
						Value:       cli.NewTimestamp(time.Now().AddDate(0, 0, -14)),
						Layout:      "2006-01-02",
						Usage:       "date to start lookback. Format: YYYY-MM-DD",
						DefaultText: "two weeks ago",
					},
					&cli.TimestampFlag{
						Name:        "before",
						Value:       cli.NewTimestamp(time.Now()),
						Layout:      "2006-01-02",
						Usage:       "date to end lookback. Format: YYYY-MM-DD",
						DefaultText: "now",
					},
					&cli.IntFlag{
						Name:        "places",
						Value:       10,
						Usage:       "number of files returned",
						DefaultText: "10",
					},
				},
				Action: func(c *cli.Context) error {
					after := c.Timestamp("after").Format("2006-01-02")
					before := c.Timestamp("before").Format("2006-01-02")

					cmd := exec.Command("git", "log", fmt.Sprintf(`--after="%s"`, after), fmt.Sprintf(`--before="%s"`, before), "--name-only", `--format=`)

					b, err := cmd.CombinedOutput()
					if err != nil {
						fmt.Println(string(b))
						fmt.Println(err)
						return err
					}

					lines := bytes.Split(b, []byte("\n"))
					files := make(map[string]int)

					for _, line := range lines {
						files[string(line)] += 1
					}

					fileCounts := make(FileCounts, 0, len(files))
					for file, count := range files {
						fileCounts = append(fileCounts, FileCount{file, count})
					}

					sort.Sort(fileCounts)

					top := c.Int("places")
					if top > len(fileCounts) {
						top = len(fileCounts)
					}

					w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

					fmt.Fprintln(w, "Place \t Count \t File")
					for i, fc := range fileCounts[:top] {
						fmt.Fprintf(w, "%d \t%d \t %s \n", i+1, fc.count, fc.name)
					}

					w.Flush()

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
