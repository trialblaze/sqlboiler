package main

import (
	"github.com/trialblaze/sqlboiler/drivers"
	"github.com/trialblaze/sqlboiler/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
