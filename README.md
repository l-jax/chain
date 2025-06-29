# ⛓️ chain

**chain** is a command-line tool for visualising linked pull requests

<img width="527" alt="image" src="https://github.com/user-attachments/assets/59e6cd85-6a3c-4ad4-bd42-15068279a8b5" />

## How does it work?

**chain** parses PR descriptions for dependencies using this case-insensitive pattern:

`do not merge until #<number>`

It uses this information to build a chain of related changes, and decides whether any links are blocked based on thier [GitHib labels](https://cli.github.com/manual/gh_label)

**chain** uses [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss) for the user interface, and has GitHub integration via [go-gh](https://github.com/cli/go-gh)

To use **chain** you must have [gh cli](https://cli.github.com/) installed and authenticated

## What's next?

**chain** is a work in progress

- [ ] Make the merge condition user-configurable
- [ ] Show ✔️ or ✖️ in `Detail` view to indicate whether the merge condition has been met
- [x] Reduce calls to GitHub API
- [x] Switch focus between `List` and `Table`
- [x] Update `Detail` for selected table row
