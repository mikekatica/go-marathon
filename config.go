/*
Copyright 2014 Rohith All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package marathon

import (
	"io"
	"io/ioutil"
	"time"
)

// Config hold the setting and options for the client
type Config struct {
	// the url for marathon
	URL string
	// event handler port
	EventsPort int
	// the interface we should be listening on for events
	EventsInterface string
	// the output for logging
	LogOutput io.Writer
	// the timeout for requests
	RequestTimeout int
	// the default timeout for deployments
	DefaultDeploymentTimeout time.Duration
	// http basic auth
	HttpBasicAuthUser string
	// http basic password
	HttpBasicPassword string
}

// NewDefaultConfig create a default client config
func NewDefaultConfig() Config {
	return Config{
		URL:                      "http://127.0.0.1:8080",
		EventsPort:               10001,
		EventsInterface:          "eth0",
		LogOutput:                ioutil.Discard,
		HttpBasicAuthUser:        "",
		HttpBasicPassword:        "",
		DefaultDeploymentTimeout: time.Duration(300) * time.Second,
		RequestTimeout:           5}
}
