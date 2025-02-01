#!/bin/bash

git config core.hooksPath .githooks

chmod +x .githooks/pre-push
chmod +x .githooks/commit-msg


echo "\033[0;32mGit hooks installed successfully."