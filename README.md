# ⛓️ chain

**chain** is a command-line tool for visualising linked pull requests

<img width="509" alt="image" src="https://github.com/user-attachments/assets/22ff11e1-1799-47bd-80d1-d7fdc1c7161a" />

## How does it work?

**chain** parses PR descriptions for dependencies using this case-insensitive pattern:

`do not merge until #<number>`

It builds a chain of linked changes, and displays the state of each using a combination of its [PR status](https://cli.github.com/manual/gh_pr_status) and [GitHib labels](https://cli.github.com/manual/gh_label)

**chain** uses [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss) for the user interface, and has GitHub integration via [go-gh](https://github.com/cli/go-gh)

To use **chain** you must have [gh cli](https://cli.github.com/) installed and authenticated

## What's next?

**chain** is a work in progress

- [ ] Make the merge condition user-configurable
- [ ] Show ✔️ or ✖️ in `Detail` view to indicate whether the merge condition has been met
- [x] Switch focus between list and table
- [x] View detail for selected table row
