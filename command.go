package main

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

var ExitSuccessful = errors.New("exit successful")

func command(arguments []string) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)

	name := arguments[0]
	var args []string
	if len(arguments) > 1 {
		args = arguments[1:]
	}

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = NewWriter(os.Stdout)

	err := cmd.Start()
	if err != nil {
		goto ERROR
	}

	// Start  command and wait for it to finish.
	group.Go(func() error {
		if err := cmd.Wait(); err != nil {
			return err
		}

		return ExitSuccessful
	})

	// Listen for signals and forward them to the command.
	group.Go(func() error {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals)

		for {
			select {
			case s := <-signals:
				if err := cmd.Process.Signal(s); err != nil {
					return err
				}
			case <-ctx.Done():
				if err := ctx.Err(); err != nil {
					return err
				}

				return ExitSuccessful
			}
		}
	})

	err = group.Wait()

ERROR:
	if err == ExitSuccessful || err == os.ErrProcessDone {
		return 0, nil
	}

	if xerr, ok := err.(*exec.ExitError); ok {
		return xerr.ExitCode(), err
	}

	if uerr := errors.Unwrap(err); uerr != nil {
		return 1, uerr
	}

	return 1, err
}
