import (
  "marathon"
  "fmt"
)

func main()  {
  marathonURL := "http://marathon.mesos:8080"
  config := marathon.NewDefaultConfig()
  config.URL = marathonURL
  client, err := marathon.NewClient(config)
  if err != nil {
  	log.Fatalf("Failed to create a client for marathon, error: %s", err)
  }
  applications, err := client.Applications()
  printf(application[1])
}
