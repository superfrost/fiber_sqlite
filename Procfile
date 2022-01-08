# Procfile

# release: python -m venv ./python_scripts/venv
# release: source ./python_scripts/venv/bin/activate
release: pip install -r requirements.txt
# release: deactivate

web: go run .


# heroku buildpacks:clear  # remove all of the buildpacks first
# heroku buildpacks:add heroku/go    # add the go buildpack second
# heroku buildpacks:add heroku/python  # add the python buildpack first

# heroku buildpacks # to see what I have done