#!/bin/bash
# Update Homebrew formula with latest release checksums
# Usage: ./scripts/update-homebrew-formula.sh v1.5.0

set -e

if [ -z "$1" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.5.0"
    exit 1
fi

VERSION=$1
VERSION_NO_V=${VERSION#v}
REPO_URL="https://github.com/gopinath-langote/1build"
RELEASE_URL="$REPO_URL/releases/download/$VERSION"

echo "Fetching release information for $VERSION..."

# Clone the homebrew tap repo temporarily
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

git clone https://github.com/gopinath-langote/homebrew-one-build.git "$TEMP_DIR/homebrew-one-build"

# Download checksums file from the release
# Note: GoReleaser v2 names it as "1build_VERSION_checksums.txt"
CHECKSUMS_URL="$RELEASE_URL/1build_${VERSION_NO_V}_checksums.txt"
CHECKSUMS=$(mktemp)
if curl -sL "$CHECKSUMS_URL" -o "$CHECKSUMS" 2>/dev/null && [ -s "$CHECKSUMS" ]; then
    echo "Downloaded checksums from release"
else
    echo "Error: Could not download checksums.txt from release"
    echo "Tried: $CHECKSUMS_URL"
    exit 1
fi

# Parse checksums from the file
SHA_X86_64_DARWIN=$(awk '/1build_Darwin_x86_64\.tar\.gz/ {print $1}' "$CHECKSUMS")
SHA_ARM64_DARWIN=$(awk '/1build_Darwin_arm64\.tar\.gz/ {print $1}' "$CHECKSUMS")
SHA_X86_64_LINUX=$(awk '/1build_Linux_x86_64\.tar\.gz/ {print $1}' "$CHECKSUMS")
SHA_ARM64_LINUX=$(awk '/1build_Linux_arm64\.tar\.gz/ {print $1}' "$CHECKSUMS")

# Validate we have all required checksums
if [ -z "$SHA_X86_64_DARWIN" ] || [ -z "$SHA_ARM64_DARWIN" ] || [ -z "$SHA_X86_64_LINUX" ] || [ -z "$SHA_ARM64_LINUX" ]; then
    echo "Error: Missing checksums for required architectures"
    echo "Available checksums:"
    cat "$CHECKSUMS"
    exit 1
fi

# Generate the formula
cat > "$TEMP_DIR/homebrew-one-build/one-build.rb" << EOF
# typed: false
# frozen_string_literal: true

class OneBuild < Formula
  desc "Frictionless way of managing project-specific commands"
  homepage "https://github.com/gopinath-langote/1build"
  version "$VERSION_NO_V"

  on_macos do
    if Hardware::CPU.intel?
      url "$RELEASE_URL/1build_Darwin_x86_64.tar.gz"
      sha256 "$SHA_X86_64_DARWIN"
    elsif Hardware::CPU.arm?
      url "$RELEASE_URL/1build_Darwin_arm64.tar.gz"
      sha256 "$SHA_ARM64_DARWIN"
    end
  end

  on_linux do
    if Hardware::CPU.intel? && Hardware::CPU.is_64_bit?
      url "$RELEASE_URL/1build_Linux_x86_64.tar.gz"
      sha256 "$SHA_X86_64_LINUX"
    elsif Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "$RELEASE_URL/1build_Linux_arm64.tar.gz"
      sha256 "$SHA_ARM64_LINUX"
    end
  end

  def install
    bin.install "1build"
  end

  test do
    system "#{bin}/1build", "--version"
  end
end
EOF

echo "Generated formula for version $VERSION_NO_V"

# Commit and push
cd "$TEMP_DIR/homebrew-one-build"
git config user.email "github-actions[bot]@users.noreply.github.com"
git config user.name "github-actions[bot]"
git add one-build.rb
git commit -m "Brew formula update for 1build version $VERSION" || echo "No changes to commit"

echo ""
echo "Formula updated successfully!"
echo "To push the changes, run:"
echo "  cd $TEMP_DIR/homebrew-one-build"
echo "  git push origin master"
echo ""
echo "Or manually verify the formula at: $TEMP_DIR/homebrew-one-build/one-build.rb"
