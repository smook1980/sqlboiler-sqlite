package main

import (
	"github.com/smook1980/sqlboiler-sqlite/driver"
	"github.com/volatiletech/sqlboiler/v4/drivers"
)

func main() {
	drivers.DriverMain(&driver.SQLiteDriver{})
}
