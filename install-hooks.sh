#!/bin/bash

git config core.hooksPath .githooks

chmod +x .githooks/pre-push

cp scripts/git-hooks/pre-commit .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit

echo "Git hooks installed successfully."