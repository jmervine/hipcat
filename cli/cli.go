package cli

import (
	"github.com/jmervine/hipcat/config"

	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jmervine/hipcat/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

func Run(action func(cfg *config.Config) error) {
	app := cli.NewApp()
	app.Name = "hipcat"
	app.Usage = "read file or stdin to hipchat"
	app.Version = "0.0.2"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "room, r",
			Value:  "",
			Usage:  "hipchat room",
			EnvVar: "HIPCAT_ROOM",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "hipchat api token",
			EnvVar: "HIPCAT_TOKEN",
		},
		cli.StringFlag{
			Name:   "sender, s",
			Value:  "hipcat",
			Usage:  "hipchat sender",
			EnvVar: "HIPCAT_SENDER",
		},
		cli.StringFlag{
			Name:   "host, H",
			Value:  "",
			Usage:  "hipchat host",
			EnvVar: "HIPCAT_HOST",
		},
		cli.StringFlag{
			Name:   "code, x",
			Value:  "true",
			Usage:  "message is code, w/o --notify",
			EnvVar: "HIPCAT_CODE",
		},
		cli.StringFlag{
			Name:   "notify, n",
			Value:  "false",
			Usage:  "notify in hipchat",
			EnvVar: "HIPCAT_NOTIFY",
		},
		cli.StringFlag{
			Name:   "color, c",
			Value:  "purple",
			Usage:  "hipchat notification color w/ --notify",
			EnvVar: "HIPCAT_COLOR",
		},
		cli.StringFlag{
			Name:   "config, C",
			Value:  "",
			Usage:  "hipcat config file ",
			EnvVar: "HIPCAT_CONFIG",
		},
	}

	app.Action = func(c *cli.Context) {
		/*
		 * Initialize Config w/ default config file.
		 */
		cfg := new(config.Config)
		defaultConfig, _ := config.ReplaceHome(config.DefaultConfig)
		cfg.LoadConfig(defaultConfig)

		/*
		 * Define error handler/reporter.
		 */
		fatal := func(err error) {
			log.Printf("ERROR: %s\n---\n\n", err)
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		/*
		 * Check for and apply secondary config file from args.
		 */
		configArg := c.String("config")
		if configArg != "" {
			if strings.HasPrefix(configArg, "~") {
				var err error
				configArg, err = config.ReplaceHome(configArg)
				if err != nil {
					fatal(err)
				}
			}

			cfg.LoadConfig(configArg)
		}

		/*
		 * Apply additional configs from args as top priority.
		 */
		if cfg.Room == "" {
			cfg.Room = c.String("room")
		}

		if cfg.Token == "" {
			cfg.Token = c.String("token")
		}

		if cfg.Host == "" {
			cfg.Host = c.String("host")
		}

		if cfg.Sender == "" {
			cfg.Sender = c.String("sender")
		}

		if cfg.Code == "" {
			cfg.Code = c.String("code")
		}

		if cfg.Color == "" {
			cfg.Color = c.String("color")
		}

		if cfg.Notify == "" {
			cfg.Notify = c.String("notify")
		}

		if err := cfg.Require(); err != nil {
			fatal(err)
		}

		/*
		 * Fetch file for reading if passed, otherwise listen for stdin.
		 */
		if len(c.Args()) > 0 {
			d, err := ioutil.ReadFile(c.Args()[0])

			if err != nil {
				fatal(err)
			}

			cfg.Message = d
		} else {
			err := cfg.ReadMessage()

			if err != nil {
				fatal(err)
			}
		}

		/*
		 * Run passed action on config built from files and cli.
		 */
		err := action(cfg)
		if err != nil {
			fatal(err)
		}
	}

	app.Run(os.Args)
}
