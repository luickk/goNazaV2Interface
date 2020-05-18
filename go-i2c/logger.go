package i2c

import logger "goNazaV2Interface/go-i2c/go-logger"

// You can manage verbosity of log output
// in the package by changing last parameter value
// (comment/uncomment corresponding lines).
var lg = logger.NewPackageLogger("i2c",
	logger.DebugLevel,
	// logger.InfoLevel,
)
