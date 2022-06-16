package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kardianos/service"
	"github.com/phileagleson/symserveragent"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async
	go p.run()
	return nil
}

func (p *program) run() {
	symserveragent.StartServer()
	fmt.Println("started running program...")
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	err := symserveragent.StopServer()
	if err != nil {
		log.Printf("Error stopping sym server agent: %+v\n", err)
	}
	return nil
}

func (p *program) Restart(s service.Service) error {
	err := p.Stop(s)
	if err != nil {
		return err
	}
	err = p.Start(s)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "SymServerAgent",
		DisplayName: "Sym Server Agent",
		Description: "The Sym Server Agent handles file uploads for PowerFrame",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// get command line args to specify install, uninstall, start, stop, restart
	args := os.Args

	var command string
	if len(args) > 1 {
		command = args[1]
	}

	switch strings.ToLower(command) {
	case "--install":
		//  log.Println("Installing...")
		fmt.Println("Installing...")
		log.Printf("Installing service...\n")
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		return
	case "--uninstall":
		// log.Println("Uninstalling...")
		fmt.Println("Uninstalling...")
		log.Printf("Uninstalling service...\n")
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		return
	case "--start":
		//log.Println("Starting service...")
		fmt.Println("Starting service...")
		log.Printf("Starting service...\n")
		err = s.Start()
		if err != nil {
			log.Fatal(err)
		}
	case "--stop":
		//log.Println("Stopping service...")
		fmt.Println("Stopping service...")
		log.Printf("Stopping service...\n")
		err = s.Stop()
		if err != nil {
			log.Fatal(err)
		}
		return
	case "--restart":
		//log.Println("Restarting service...")
		fmt.Println("Restarting service...")
		log.Printf("Restarting service...\n")
		err = s.Restart()
		if err != nil {
			log.Fatal(err)
		}
	default:
		s.Run()

	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
