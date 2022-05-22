# tureng-cli

A super simple cli tool for translating

It supports 4 language set which described in below.

- turkish-english
- german-english
- spanish-english
- french-english

# Flags

- --language
- --count -`represents the count of translated words.`

# Commands

You can execute any command and custom commands with using "-c" flag. For example:

    tureng (turkish-english) >> -c **command here**

## Custom Commands

- update-language

# Installation

    go install github.com/keremdokumaci/tureng-cli@v0.1.0

# Sample Usage

    "tureng-cli --language turkish-english --count 3"
