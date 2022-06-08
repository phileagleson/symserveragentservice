package main

import (
	"log"
	"os"
	"strings"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
  // Start should not block. Do the actual work async
  go p.run()
  return nil
}

func (p *program) run() {
  logger.Info("started running program...")
}

func (p *program) Stop(s service.Service) error {
  // Stop should not block. Return with a few seconds.
  return nil
}

func main() {
  svcConfig := &service.Config{
    Name: "GoServiceExample",
    DisplayName: "Go Service Example",
    Description: "This is an example Go service.",
  }

  prg := &program{}
  s, err := service.New(prg, svcConfig)
  if err != nil {
    log.Fatal(err)
  }
  logger, err = s.Logger(nil)
  if err != nil {
    log.Fatal(err)
  }
  // get command line args to specify install, uninstall, start, stop, restart
  args := os.Args

  if len(args) > 2 {
    log.Fatal("Please specify only one argument: install, uninstall, start, stop, or restart")
  } else if len(args) == 1 {
    log.Fatal("Please specify a command: install, uninstall, start, stop, or restart")
  } 

  command := args[1]


  switch strings.ToLower(command) {
case "--install":
  log.Println("Installing...")
  err = s.Install()
    if err != nil {
      log.Fatal(err)
   }
  return
case "--uninstall":
  log.Println("Uninstalling...")
  s.Uninstall()
  return
case "--start":
  log.Println("Starting service...")
  s.Run()
case "--stop":
  log.Println("Stopping service...")
  s.Stop()
  return
case "--restart":
  log.Println("Restarting service...")
  s.Restart()
  default:
    log.Fatal("Please specify install, uninstall, start, stop, or restart")

}

  err = s.Run()
  if err != nil {
    logger.Error(err)
  }

}
