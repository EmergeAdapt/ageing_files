package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	cb "github.com/emergeadapt/caseblocks.helpers"

	"github.com/codegangsta/cli"
)

const (
	VERSION = "1.2.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "ageing_files"
	app.Usage = "Checks file system for old files and notifies via email."
	app.Version = VERSION

	app.Commands = []cli.Command{
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "runs ageing_files process",
			Action: func(c *cli.Context) {
				if c.String("basepath") != "" && c.String("folders") != "" {
					maxage := c.Int("maxage")
					if maxage < 1 {
						maxage = 3600
					}
					watchFolders(c.String("basepath"), strings.Split(c.String("folders"), ","), maxage, c.String("exclude"))
				}
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "basepath",
					Value: "/var/log",
					Usage: "the parent folder of the folders being monitored",
				},
				cli.StringFlag{
					Name:  "folders",
					Value: "apache,mysql",
					Usage: "the subfolders who's file age is monitored",
				},
				cli.IntFlag{
					Name:  "maxage",
					Value: 3600,
					Usage: "the maximum permissable age of a file in seconds",
				},
				cli.StringFlag{
					Name:  "exclude",
					Value: ".DS_Store,backup",
					Usage: "comma-separated list of file names to ignore",
				},
			},
		},
	}

	app.Run(os.Args)
}

func watchFolders(basepath string, folders []string, maxage int, exclude string) {

	log := cb.NewConsoleLogger()

	excluded_filenames := strings.Split(exclude, ",")
	for _, folder := range folders {
		matches, err := filepath.Glob(fmt.Sprintf("%s/%s/*", basepath, folder))
		if err != nil {
			log.Panic(err.Error())
			return
		}

		for _, path := range matches {
			file, err := os.Open(path)
			if err != nil {
				log.Panic(err.Error())
				return
			}

			fileInfo, err := file.Stat()
			if err != nil {
				log.Panic(err.Error())
				return
			}

			if fileInfo.Mode().IsRegular() && !strings.HasSuffix(strings.ToLower(path), "end") && !IsExcluded(path, excluded_filenames) {
				if time.Since(fileInfo.ModTime()).Seconds() > float64(maxage) {
					fmt.Println(path)
				}
			}
		}
	}
}

func IsExcluded(filepath string, excluded_filenames []string) bool {
	for _, excluded_filename := range excluded_filenames {
		excluded_filename := strings.ToLower(excluded_filename)
		base_fp := strings.ToLower(path.Base(filepath))

		if base_fp == strings.ToLower(excluded_filename) {
			return true
		}
		pos := strings.Index(excluded_filename, "*")
		if pos >= 0 {
			if strings.Index(base_fp, excluded_filename[0:pos]) >= 0 {
				return true
			}
		}

	}
	return false
}
