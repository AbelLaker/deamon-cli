package main

import (
	"fmt"
	cli "github.com/urfave/cli/v2"
	"os"
)

var TestCmd1 = &cli.Command{
	Name:  "test1",
	Usage: "Start a test1 daemon process",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "p",
			Usage: "only test get p val",
		},
	},
	Action: func(ctx *cli.Context) error {
		//获取flag
		p := ctx.Int64("p")
		fmt.Println("Hello Test1 Int64 p =", p)
		return nil
	},
}

var TestCmd2 = &cli.Command{
	Name:  "test2",
	Usage: "Start a test2 daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "v",
			Usage: "only test get v val",
		},
	},
	Before: func(context *cli.Context) error {
		fmt.Println("Test2 Run Before")
		return nil
	},
	After: func(context *cli.Context) error {
		fmt.Println("Test2 Run After")
		return nil
	},
	Action: test2_run,
}

func main() {

	app := &cli.App{
		Name:                 "deamon-cli",
		Usage:                "example for cli",
		Version:              "1.0",
		EnableBashCompletion: true,
		Flags:                []cli.Flag{},
		After: func(c *cli.Context) error {
			if r := recover(); r != nil {
				// Generate report in LOTUS_PATH and re-raise panic
				//build.GeneratePanicReport(c.String("panic-reports"), c.String("repo"), c.App.Name)
				fmt.Println("panic(r):", r)
				panic(r)
			}
			return nil
		},

		Commands: []*cli.Command{
			TestCmd1,
			TestCmd2,
		},
	}

	app.Setup()

	RunApp(app)
}

func RunApp(app *cli.App) {
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func test2_run(ctx *cli.Context) error {
	//获取flag
	v := ctx.String("v")
	fmt.Println("Hello Test2 String v =", v)
	return nil
}
