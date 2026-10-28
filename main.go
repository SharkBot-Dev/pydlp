package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"pydlp/tools"
)

const (
	appName     = "pydlp"
	versionsDir = "yt-dlps"
	versionFile = "version.txt"
)

func usage() {
	fmt.Fprintf(os.Stdout, `Usage:
  %s <command> [arguments]

Commands:
  install [VERSION]  Install a version (defaults to the latest version)
  update [VERSION]   Update to a version (defaults to the latest version)
  run [ARGS...]      Run yt-dlp with the selected version; remaining arguments
                     are passed directly to yt-dlp
  version [VERSION]  Show the selected version, or set it when VERSION is given
  versions           List installed versions
  help               Show this help

Options:
  -h, --help         Show help

Examples:
  %s install
  %s install 2026.10.22
  %s version 2026.10.22
  %s run --help
  %s run -f best URL
`, appName, appName, appName, appName, appName, appName)
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, appName+": "+format+"\n", args...)
	os.Exit(1)
}

func readSelectedVersion() (string, error) {
	data, err := os.ReadFile(versionFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", errors.New("no version selected; run 'pydlp install' first")
		}
		return "", fmt.Errorf("read selected version: %w", err)
	}

	version := strings.TrimSpace(string(data))
	if version == "" {
		return "", errors.New("selected version is empty; set it with 'pydlp version VERSION'")
	}

	if version == "." || version == ".." ||
		strings.ContainsAny(version, `/\`) {
		return "", fmt.Errorf("invalid version %q", version)
	}
	return version, nil
}

func installOrUpdate(command string, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("too many arguments for %s\n\nRun 'pydlp help' for usage", command)
	}

	tools.PrintDeviceString()

	var ok bool
	if len(args) == 1 {
		ok = tools.DownloadVersionYtdlp(args[0])
	} else {
		ok = tools.UpdateYtdlp()
	}
	if !ok {
		return fmt.Errorf("%s failed", strings.Title(command))
	}
	fmt.Printf("%s successful\n", strings.Title(command))
	return nil
}

func run(args []string) error {
	version, err := readSelectedVersion()
	if err != nil {
		return err
	}

	suffix := tools.GetDeviceSuff()
	if suffix == "null" {
		return errors.New("this device is not supported")
	}

	target := filepath.Join(versionsDir, version, "yt-dlp"+suffix)
	cmd := exec.Command(target, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		return fmt.Errorf("could not run %q: %w", target, err)
	}
	return nil
}

func setOrShowVersion(args []string) error {
	tools.PrintDeviceString()
	if len(args) > 1 {
		return errors.New("version accepts at most one argument; run 'pydlp help' for usage")
	}

	if len(args) == 1 {
		version := strings.TrimSpace(args[0])
		if version == "" || version == "." || version == ".." ||
			strings.ContainsAny(version, `/\\`) {
			return fmt.Errorf("invalid version %q", args[0])
		}
		if err := os.WriteFile(versionFile, []byte(version+"\n"), 0644); err != nil {
			return fmt.Errorf("save selected version: %w", err)
		}
		fmt.Printf("Selected version: %s\n", version)
		return nil
	}

	version, err := readSelectedVersion()
	if err != nil {
		return err
	}
	fmt.Println(version)
	return nil
}

func listVersions() error {
	tools.PrintDeviceString()

	entries, err := os.ReadDir(versionsDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No versions installed.")
			return nil
		}
		return fmt.Errorf("list versions: %w", err)
	}

	fmt.Println("Installed versions:")
	found := false
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("  " + entry.Name())
			found = true
		}
	}
	if !found {
		fmt.Println("  (none)")
	}

	version, err := readSelectedVersion()
	if err == nil {
		fmt.Printf("\nSelected version: %s\n", version)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	command := os.Args[1]
	args := os.Args[2:]
	if command == "-h" || command == "--help" || command == "help" {
		usage()
		return
	}
	if command != "run" {
		for _, arg := range args {
			if arg == "-h" || arg == "--help" {
				usage()
				return
			}
		}
	}

	if err := os.MkdirAll(versionsDir, 0755); err != nil {
		fail("create versions directory: %v", err)
	}

	var err error
	switch command {
	case "run":
		err = run(args)
	case "install", "update":
		err = installOrUpdate(command, args)
	case "version":
		err = setOrShowVersion(args)
	case "versions":
		if len(args) != 0 {
			err = errors.New("versions takes no arguments; run 'pydlp help' for usage")
		} else {
			err = listVersions()
		}
	default:
		usage()
		fail("unknown command %q", command)
	}

	if err != nil {
		fail("%v", err)
	}
}
