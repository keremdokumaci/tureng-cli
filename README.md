# tureng-cli

A super simple cli tool for translating

It supports 4 language set which described in below.

- turkish-english
- german-english
- spanish-english
- french-english

# Flags

- --language
- --count `represents the count of translated words.`

# Commands

You can execute any command and custom commands with using "$" flag. For example:

    tureng (turkish-english) >> $ **command here**

### Commands

- update-language **language**
- clear

# Installation

    go install github.com/keremdokumaci/tureng-cli@latest

# Sample Usage

    tureng-cli --language turkish-english --count 3
