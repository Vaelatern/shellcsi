package script

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Exists checks the local filesystem, the current working directory, for the script, and that the
// file is executable
func Exists(script string) bool {
	i, err := os.Stat("./" + script)
	if err != nil || i.Mode().IsDir() || i.Mode()&0111 == 0 {
		return false
	}
	return true
}

// RunViaJson belongs as a  method on the Config arguments, but unfortunately a generic, used here
// to make the JSON parsing more useful, can't be used in a method.
// This relies on the Command being in the Current Working Directory. This does not support baking
// the commands in to a filesystem because apparently `exec.Command` doesn't support pulling scripts
// from there. We may shim this later.
func RunViaJson[Args any, In any, Out any](ctx context.Context, args Args, input In, script string) (Out, error) {
	var result Out

	wrapped := struct {
		Local  Args
		Passed In
	}{
		Local:  args,
		Passed: input,
	}

	data, err := json.Marshal(wrapped)
	if err != nil {
		return result, err
	}

	cmd := exec.CommandContext(ctx, "./"+script)
	cmd.Stdin = bytes.NewReader(data)

	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf

	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf

	if err := cmd.Run(); err != nil {
		return result, errors.Join(err, fmt.Errorf("Command errored: %s", outBuf.String()), fmt.Errorf("Standard Error Output: %s", errBuf.String()))
	}

	if err := json.Unmarshal(outBuf.Bytes(), &result); err != nil {
		return result, errors.Join(err, fmt.Errorf("Unmarshalling errored: %s", outBuf.String()), fmt.Errorf("Standard Error Output: %s", errBuf.String()))
	}

	return result, nil
}
