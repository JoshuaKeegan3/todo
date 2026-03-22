# todo

A terminal UI for browsing `TODO:` comments in the current directory.

Uses [ripgrep](https://github.com/BurntSushi/ripgrep) to find all `TODO:` comments and displays them in an interactive list powered by [Bubbletea](https://github.com/charmbracelet/bubbletea).

## Usage

Run `todo` from within any project directory.

## Keybindings

| Key      | Action                                      |
|----------|---------------------------------------------|
| `o`      | Open the selected file in Zed               |
| `r`      | Cat the README of the selected item's project |
| `ctrl+c` | Quit                                        |
