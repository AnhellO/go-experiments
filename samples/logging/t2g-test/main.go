package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	// Set log file
	logFile := "test.log"
	rawJSON := []byte(fmt.Sprintf(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "%s"],
		"errorOutputPaths": ["stderr", "%s"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase",
			"timeKey": "timestamp",
			"timeEncoder": "ISO8601TimeEncoder"
		  }
	  }`, logFile, logFile))

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	simpleHttpGet("www.sogo.com")
	simpleHttpGet("http://www.sogo.com")
}

func simpleHttpGet(url string) {
	zap.S().Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		zap.S().Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		zap.S().Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
