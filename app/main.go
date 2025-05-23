package main

import (
	"fmt"
	"github.com/sam-caldwell/surveillance-proxy/v2/common"
	"github.com/sam-caldwell/surveillance-proxy/v2/ubiquity"
	"log"
	"net/http"
)

func main() {

	var (
		recvAddress = common.AddressPortPattern(common.RequireEnv("RECVR_ADDRESS"))
		authToken   = common.RequiredStringSize(64, common.RequireEnv("AUTH_TOKEN"))
		jiraUser    = common.RequireEnv("JIRA_USER")
		jiraToken   = common.RequireEnv("JIRA_TOKEN")
		jiraBaseURL = common.RequireEnv("JIRA_BASE_URL")
		jiraProject = common.RequireEnv("JIRA_PROJECT")
		//UbiquityAddress = common.AddressPortPattern(common.RequireEnv("UBIQUITY_ADDRESS"))
		//UbiquityToken   = common.RequireEnv("UBIQUITY_ACCESS_TOKEN")
	)

	// ssid device data collector
	//go func() {
	//	for {
	//		ubiquity.CollectSsidData(&UbiquityAddress, &UbiquityToken)
	//	}
	//}()
	// webhook listener (captures ubiquity.py web hooks)
	func() {
		http.HandleFunc(
			"/cameras",
			ubiquity.WebhookHandlerFactory(
				&authToken,
				&jiraUser,
				&jiraToken,
				&jiraBaseURL,
				&jiraProject))

		addr := fmt.Sprintf("%s", recvAddress)

		log.Println("Listening on", addr)

		log.Fatal(http.ListenAndServe(addr, nil))
	}()
}
