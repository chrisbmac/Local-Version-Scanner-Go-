#!/bin/sh

set -e

ARCH=$(uname -m)

if [ "$ARCH" = "arm64" ]; then
  BINARY_URL="https://github.com/chrisbmac/Local-Version-Scanner-Go-/releases/download/v1.0.3/versions-arm"
elif [ "$ARCH" = "x86_64" ]; then
  BINARY_URL="https://github.com/chrisbmac/Local-Version-Scanner-Go-/releases/download/v1.0.3/versions-intel"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

echo "Downloading versions for $ARCH..."
curl -fsSL "$BINARY_URL" -o /usr/local/bin/versions

chmod +x /usr/local/bin/versions

echo "Installed successfully. Run: versions"
