# CHANGELOG

## v0.3.0 (2021-05-05)

### Added

- feat: replace xdg with os (2021-05-05)

- feat: use ioe.ReadInput to replace readStdin (#4) (2021-04-17)

- feat: use my own color pkg to wrap fatih/color (#1) (2021-04-14)

### Others

- chore: fastly to match readme (2021-05-05)

- docs: better wording for config (2021-05-05)

- docs: update config location (2021-05-05)

- refactor: move main to cmd (2021-05-05)

- refactor: completely move cli to internal (2021-05-05)

- refactor: move config to internal (2021-05-05)

- build: update go.mod (2021-05-05)

- refactor: move usage to const (#3) (2021-04-16)

- refactor: move aliases to define var (#2) (2021-04-14)

- chore: update go in github action (2021-04-14)

- chore(changelog): generate v0.2.0 (2021-03-17)

## v0.2.0 (2021-3-17)

### Added

- feat: deprecated ioutil

### Others

- build: bump go 1.16 in go.mod

- chore: bump go 1.16 in github action

- build: update go.mod

- chore(readme): add go 1.16 instal guide

- build: update go.mod

- chore(changelog): generate v0.1.0

## v0.1.0 (2021-1-7)

### Added

- feat(cli): add delete and delete all command

- feat(config): add delete all

- feat(cli): add delete cmd skeleton

- feat(cli): remove debug flag as no use

- feat(cli): add list cmd to list all users

- feat(cli): add status command

- feat(cli): add switch command

- feat(cli): add switch command skeleton

- feat(cli): add load and save config in add cmd

- feat(cli): add nickname flag

- feat(config): add save config to file

- feat(config): add load config from file

- feat(cli): add debug flag

- feat(cli): skeleton cli app

### Fixed

- fix(cli): missing get nickname flag

- fix(cli): user not author gitconfig

- fix(config): mkdir before save config

### Others

- chore(readme): add badges

- chore(readme): add thanks

- chore(readme): add usage

- chore(cli): better RunSwitch comment

- refactor(cli): remove flags map to use directly field in struct

- chore: add github action

- chore: remove build file

- chore: add LICENSE

- chore(cli): remove newline after question y/n

- chore: add README
