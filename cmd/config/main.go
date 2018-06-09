package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func main() {
	fmt.Println("hello api server")
}

func init() {
	initializer := &cli.App{
		Name:  "event",
		Usage: "rest api server launcher",
		Flags: []cli.Flag{
			// http config
			cli.IntFlag{
				Name:   "http.port",
				Value:  9000,
				Usage:  "http server port",
				EnvVar: "HTTP_PORT",
			},
			// database config
			cli.StringFlag{
				Name:   "db.address",
				Value:  "root:password@/localdb?charset=utf8&parseTime=True&loc=Local",
				Usage:  "db address",
				EnvVar: "DB_ADDRESS",
			},
			cli.IntFlag{
				Name:   "db.maxIdleConns",
				Value:  0,
				Usage:  "db max idle connections (0: no idle connections)",
				EnvVar: "DB_MAX_IDLE_CONNS",
			},
			cli.IntFlag{
				Name:   "db.maxOpenConns",
				Value:  0,
				Usage:  "db max open connections (0: unlimited connections)",
				EnvVar: "DB_MAX_OPEN_CONNS",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Printf("# --- Http Config --------------------------------------\n"+
				"  http.port               : %d\n"+
				"\n",
				c.Int("http.port"),
			)
			fmt.Printf("# --- DB Config   --------------------------------------\n"+
				"  db.address              : %s\n"+
				"  db.maxIdleConns         : %d\n"+
				"  db.maxOpenConns         : %d\n"+
				"\n",
				c.String("db.address"),
				c.Int("db.maxIdleConns"),
				c.Int("db.maxOpenConns"),
			)
			return nil
		},
		HideVersion: true,
	}
	initializer.RunAndExitOnError()
}
