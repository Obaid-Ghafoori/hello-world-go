#!/bin/bash

git config core.hooksPath .githooks

chmod +x .githooks/pre-push
chmod +x .githooks/commit-msg
chmod +x .githooks/pre-checkout

echo "Git hooks installed successfully."