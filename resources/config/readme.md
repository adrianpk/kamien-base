# Notes
* This is a sample directory.
* You can move its content to the folder referenced by your ${{.AppNameUppercase}}_HOME environment var.
* As an alternative location you can also move its content to a folder named '.{{.AppNameLowercase}}' under your user's home folder (~/{{.AppNameLowercase}}).
* If you are in Linux or macOS:

```bash
# From the root directory of this app
$ cp ./_{{.AppNameLowercase}}/* ${{.AppNameUppercase}}_HOME
# Alternatively
$ mkdir ~/$USER/.{{.AppNameLowercase}}
$ cp ./_{{.AppNameLowercase}}/* ~/$USER/.{{.AppNameLowercase}}
```
* You can safely delete after doing one ore boths of the previous steps. It is not uses by the application.
