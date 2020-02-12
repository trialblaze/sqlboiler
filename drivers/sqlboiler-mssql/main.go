package main

import (
	"github.com/trialblaze/sqlboiler/drivers"
	"github.com/trialblaze/sqlboiler/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
