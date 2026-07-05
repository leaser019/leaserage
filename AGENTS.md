# AGENTS.md

## Principle

Act like a careful production engineer. Understand the goal before editing. Prefer the smallest safe change. Follow existing project patterns. Do not refactor unrelated code.

## Project

This repository builds `leaserage`, a Go CLI that installs personal workflow skills, provider config templates, and MCP setup for AI coding tools.

## Commands

- Test: `go test ./...`
- Run CLI: `go run ./cmd/leaserage --help`
- Build Linux: `GOOS=linux GOARCH=amd64 go build -o dist/leaserage-linux-amd64 ./cmd/leaserage`
- Build Windows: `GOOS=windows GOARCH=amd64 go build -o dist/leaserage-windows-amd64.exe ./cmd/leaserage`
- Build macOS from Linux: `GOOS=darwin GOARCH=arm64 go build -o dist/leaserage-darwin-arm64 ./cmd/leaserage`

## Safety

- Do not commit secrets, API keys, tokens, or database URLs.
- Do not overwrite user config without backup.
- Prefer dry-run and temp-home tests before touching real home directories.
- DBHub should default disabled until a database target is explicitly configured.

## Closeout

When done, respond with summary, files changed, verification, and risks or notes.
