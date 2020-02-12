package main

import (
	"github.com/trialblaze/sqlboiler/drivers"
	"github.com/trialblaze/sqlboiler/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
