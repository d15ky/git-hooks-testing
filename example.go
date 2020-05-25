package main

import (
	conf "github.com/d15ky/beebot/config"
	"github.com/d15ky/beebot/internal/utils"
	cluster "github.com/d15ky/beebot/internal/cluster"
	database "github.com/d15ky/beebot/internal/db"
	"github.com/d15ky/beebot/internal/slackbot"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)
