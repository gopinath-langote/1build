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
CHECKSUMS_URL="$RELEASE_URL/checksums.txt"
CHECKSUMS=$(mktemp)
if curl -sL "$CHECKSUMS_URL" -o "$CHECKSUMS" 2>/dev/null && [ -s "$CHECKSUMS" ]; then
    echo "Downloaded checksums from release"
else
    echo "Error: Could not download checksums.txt from release"
    exit 1
fi

# Parse checksums and create the formula
declare -A SHA256_MAP
while IFS=' ' read -r sha256 filename; do
    SHA256_MAP["$filename"]=$sha256
done < "$CHECKSUMS"

# Validate we have all required architectures
REQUIRED=("1build_Darwin_x86_64.tar.gz" "1build_Darwin_arm64.tar.gz" "1build_Linux_x86_64.tar.gz" "1build_Linux_arm64.tar.gz")
for arch in "${REQUIRED[@]}"; do
    if [ -z "${SHA256_MAP[$arch]}" ]; then
        echo "Error: Missing checksums for $arch"
        echo "Available files:"
        printf '%s\n' "${!SHA256_MAP[@]}"
        exit 1
    fi
done

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
      sha256 "${SHA256_MAP[1build_Darwin_x86_64.tar.gz]}"
    elsif Hardware::CPU.arm?
      url "$RELEASE_URL/1build_Darwin_arm64.tar.gz"
      sha256 "${SHA256_MAP[1build_Darwin_arm64.tar.gz]}"
    end
  end

  on_linux do
    if Hardware::CPU.intel? && Hardware::CPU.is_64_bit?
      url "$RELEASE_URL/1build_Linux_x86_64.tar.gz"
      sha256 "${SHA256_MAP[1build_Linux_x86_64.tar.gz]}"
    elsif Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "$RELEASE_URL/1build_Linux_arm64.tar.gz"
      sha256 "${SHA256_MAP[1build_Linux_arm64.tar.gz]}"
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
