package libcontainer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/subosito/gotenv"
)

type AzureConfig struct {
	TenantId       string `json:"tenantId"`
	SubscriptionId string `json:"subscriptionId"`
	ClientId       string `json:"aadClientId"`
	ClientSecret   string `json:"aadClientSecret"`
	Location       string `json:"location"`
	ResourceGroup  string `json:"resourceGroup"`
}

func init() {

	var a AzureConfig
	jsonFile, err := os.Open("/etc/kubernetes/azure.json")
	check(err)

	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&a)

	envFile, err := os.Create("/.env")
	check(err)

	defer envFile.Close()

	w := bufio.NewWriter(envFile)

	_, err = fmt.Fprintf(w, "%s=%s\n", "AZ_TENANT_ID", a.TenantId)
	check(err)
	_, err = fmt.Fprintf(w, "%s=%s\n", "AZ_SUBSCRIPTION_ID", a.SubscriptionId)
	check(err)
	_, err = fmt.Fprintf(w, "%s=%s\n", "AZ_CLIENT_ID", a.ClientId)
	check(err)
	_, err = fmt.Fprintf(w, "%s=%s\n", "AZ_CLIENT_SECRET", a.ClientSecret)
	check(err)
	_, err = fmt.Fprintf(w, "%s=%s\n", "AZ_LOCATION", a.Location)
	check(err)
	_, err = fmt.Fprintf(w, "%s\n", "AZ_SAMPLES_KEEP_RESOURCES=0")
	check(err)

	environment := getEnv(a.ResourceGroup)
	_, err = fmt.Fprintf(w, "%s=%s\n", "ENVIRONMENT", environment)
	check(err)
	w.Flush()

	gotenv.Load()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getEnv(s string) (env string) {
	x := strings.Split(s, "-")
	env = x[1]
	return
}
