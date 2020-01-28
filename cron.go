package main

import "log"

type cronService struct {

}

func (s *cronService) Init(){
	log.Println("Cron Service Init")
}
