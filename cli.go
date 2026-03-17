package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network Lookup CLI"
	app.Usage = "Lets you query network information from the command line, using for ip, CNAMEs, mx records and name servers"

	// utilizamos las mismas flags para todos los comandos, por eso
	// se definen aca
	myflags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the name server for a particular host",
			Flags: myflags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}

				for _, nameserver := range ns {
					fmt.Println(nameserver.Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the ip addresses for a particular host",
			Flags: myflags,
			Action: func(c *cli.Context) error {
				ips, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}

				for _, ip := range ips {
					fmt.Println(ip)
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: myflags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular host",
			Flags: myflags,
			Action: func(c *cli.Context) error {
				mxRecords, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}
				for _, mx := range mxRecords {
					fmt.Println(mx.Host, mx.Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
