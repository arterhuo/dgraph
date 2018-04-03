/*
 * Copyright 2018 Dgraph Labs, Inc. and Contributors
 *
 * This file is available under the Apache License, Version 2.0,
 * with the Commons Clause restriction.
 */

package x

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	flag.Set("debugmode", "true")
	os.Exit(m.Run())
}
