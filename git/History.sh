#!/bin/bash

echo "╭──────────────────────────────────────────────╮"
echo "│                 Histotrique                  │"
echo "╰──────────────────────────────────────────────╯"
git log

echo "╭──────────────────────────────────────────────╮"
echo "│       Histotrique des push avec commit       │"
echo "╰──────────────────────────────────────────────╯"
git log --pretty=format:"%h - %an, %ar : %s"

echo "╭──────────────────────────────────────────────╮"
echo "│               One-Line History               │"
echo "╰──────────────────────────────────────────────╯"
git log --oneline

echo "╭──────────────────────────────────────────────╮"
echo "│      Histotrique des 2 derniére entrée       │"
echo "╰──────────────────────────────────────────────╯"
git log -n 2

echo "╭──────────────────────────────────────────────╮"
echo "│        Histotrique des 5 derniére min        │"
echo "╰──────────────────────────────────────────────╯"
git log --since="5 minutes ago"

echo "╭──────────────────────────────────────────────╮"
echo "│       Histotrique format personnalisé        │"
echo "╰──────────────────────────────────────────────╯"
git log --pretty=format:"* %h %ad | %s (%d) [%an]" --date=short




