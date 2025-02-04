#!/bin/bash

git config core.hooksPath .githooks

chmod +x .githooks/pre-push
chmod +x .githooks/commit-msg


echo "Git hooks installed successfully."