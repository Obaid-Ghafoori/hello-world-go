#!/bin/bash

git config core.hooksPath .githooks

chmod +x .githooks/pre-push

echo "Git hooks installed successfully."