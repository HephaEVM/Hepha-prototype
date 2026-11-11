package vm

import "os"

var cooperativePreparation = os.Getenv("COLD_COOPERATIVE_PREPARATION") == "1"
