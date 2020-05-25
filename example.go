package main

import (
	conf "github.com/d15ky/git-hooks-testing/config"
	"github.com/d15ky/git-hooks-testing/internal/utils"
	cluster "github.com/d15ky/git-hooks-testing/internal/cluster"
	database "github.com/d15ky/git-hooks-testing/internal/db"
	"github.com/d15ky/git-hooks-testing/internal/slackbot"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)
