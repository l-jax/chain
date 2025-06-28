# ⛓️ chain

**chain** is a command-line tool for visualising linked pull requests

![Screenshot](https://github.com/user-attachments/assets/c4e04bea-528f-4787-a669-365c8106a138)

## How does it work?

**chain** parses PR descriptions for dependencies using the case-insensitive pattern 

`do not merge until #<number>`

It builds a chain of linked changes for each open pull request, and displays the state of each link using a combination of [PR status](https://cli.github.com/manual/gh_pr_status) and [GitHib labels](https://cli.github.com/manual/gh_label)

**chain** uses [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss) for the user interface, with GitHub integration via [go-gh](https://github.com/cli/go-gh)

> **Note -** to use **chain** you must have [gh cli](https://cli.github.com/) installed and authenticated

## What's next?

**chain** is a work in progress

- [ ] Make labels user configurable
- [ ] Display chained PRs `mergeable` or `blocked` based on status and labels of linked PRs
- [x] Switch focus between list and table
- [x] View detail for selected table row