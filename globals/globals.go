package globals

import "time"

const AzureIP = "40.115.67.85"

var HubPort = ":8083"

const HubIP = AzureIP

var HealthCheckWait = time.Second * 45
var HealthUpdateInterval = time.Second * 20
