# Heroku deploy

Set Buildpacks:

`heroku/python`
`heroku/go`

Don't forget to set Config Vars
```
GOVERSION=go1.17.3
PORT=5000
TELEGRAM_BOT_TOKEN=
PYTHON_PATH=python3
```

If you don't use virtual environment for python (like in container or heroku) set `PYTHON_PATH` to `python` or `python3`.

File `requirements.txt` in root folder need `heroku` to install python3 dependencies via pip