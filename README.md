# blackhole

Load every http request body and headers received into stderr.
Hides Authorization header if any.

## setup

If you want to modify the default config, modify `template.conf.yaml` and rename it as `conf.yaml` before running the program.
You can also set the env variables : `BH_PORT`.

## run

`go run .`

or

`BH_PORT=8080 go run .`

or via docker

`docker run -it --rm -p80:80 chevaliersoft/blackhole:latest`

## deployment

The `deploy` folder contains a working example of deployment setup using scaleway.
