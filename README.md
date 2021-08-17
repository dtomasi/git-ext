# git-ext

git-ext contains some useful git extensions

## Installation

    curl -s https://raw.githubusercontent.com/dtomasi/git-ext/main/install | bash

## git dirclone

The dirclone extension clones git repositories to a given root directory using the URL path as a directory structure.

Example:
```shell
## This can be added to your shell environment - .bashrc or .zshrc using
## echo 'export GIT_DIRCLONE_ROOT_DIR="~/projects/src"' >> ~/.zshrc
export GIT_DIRCLONE_ROOT_DIR="~/projects/src"
git dirclone https://github.com/KDE/dummy.git
```

This command will clone the repository to `~/projects/src/github.com/KDE/dummy`

