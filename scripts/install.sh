#!/bin/sh
set -eu

VERSION="${LEASERAGE_VERSION:-latest}"
REPO="${LEASERAGE_REPO:-vomkhang/leaserage}"
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "unsupported architecture: $ARCH" >&2; exit 1 ;;
esac

NAME="leaserage-${OS}-${ARCH}"
URL="https://github.com/${REPO}/releases/${VERSION}/download/${NAME}.tar.gz"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT INT TERM

curl -fsSL "$URL" | tar -xz -C "$TMP"
mkdir -p "$HOME/.local/bin"
cp "$TMP/leaserage" "$HOME/.local/bin/leaserage"
chmod +x "$HOME/.local/bin/leaserage"
"$HOME/.local/bin/leaserage" "$@"
