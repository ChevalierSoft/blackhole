# blackhole

Load every http requests received in stderr and blackhole.log using https.

## setup

Modify `template.conf.yaml` and remane it as `conf.yaml` before running the program.
You can also use env variables : `BH_PORT`.

## run

`BH_PORT=8080 go run .`
