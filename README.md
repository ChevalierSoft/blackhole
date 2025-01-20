# blackhole

Load every http requests received in stderr and blackhole.log using https.

## setup

modify `template.conf.yaml` and remane it as `conf.yaml` before runing the program.
You can also use env variables : `BH_PORT`, `BH_DOMAIN_NAME` and `BH_EMAIL`.

## run

`go run .`
