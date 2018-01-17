package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/themech/logrus-middleware"
)

func main() {

	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &logrus.JSONFormatter{}

	// Note: you can also use something like Logger: logger.WithFields(logrus.Fields{"a": "b"}) below.
	// This way you can add your own fields to every middleware log entry.
	l := logrusmiddleware.Middleware{
		Name:   "example",
		Logger: logger,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello world\n")
	}

	http.Handle("/", l.Handler(http.HandlerFunc(handler), "homepage"))

	http.ListenAndServe(":8080", nil)
}
