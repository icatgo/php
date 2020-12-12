package php

import (
	"github.com/syyongx/php2go"
)

// Exec exec()
// returnVar, 0: succ; 1: fail
// Return the last line from the result of the command.
// command format eg:
//   "ls -a"
//   "/bin/bash -c \"ls -a\""
var Exec = php2go.Exec

// System system()
// returnVar, 0: succ; 1: fail
// Returns the last line of the command output on success, and "" on failure.
var System = php2go.System

// Passthru passthru()
// returnVar, 0: succ; 1: fail
var Passthru = php2go.Passthru
