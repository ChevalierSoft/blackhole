terraform {
  required_providers {
    scaleway = {
      source = "scaleway/scaleway"
    }
  }
  required_version = ">= 0.13"
}

provider "scaleway" {
  access_key = var.scaleway_access_key
  secret_key = var.scaleway_secret_key
  region     = var.region
  zone       = var.host_zone
  project_id = var.project_id
}

resource "scaleway_container_namespace" "main" {
  name        = "main"
  description = "main container"
}

resource "scaleway_container" "blackhole" {
  name            = "blackhole-container"
  description     = "blackhole that record and log payloads"
  namespace_id    = scaleway_container_namespace.main.id
  registry_image  = "chevaliersoft/blackhole:latest"
  port            = var.bh_port
  http_option     = "redirected"
  cpu_limit       = 140
  memory_limit    = 256
  min_scale       = 0
  max_scale       = 5
  timeout         = 600
  max_concurrency = 69
  privacy         = "public"
  protocol        = "http1"
  deploy          = true

  environment_variables = {
    "BH_PORT" = var.bh_port
  }
  secret_environment_variables = {
    "deployment" = "terraform"
  }
}

resource "scaleway_domain_record" "blackhole" {
  dns_zone = var.dns_zone
  name     = "bh"
  type     = "CNAME"
  data     = "${scaleway_container.blackhole.domain_name}."
  ttl      = 3600
}

resource "scaleway_container_domain" "blackhole" {
  container_id = scaleway_container.blackhole.id
  hostname     = "${scaleway_domain_record.blackhole.name}.${scaleway_domain_record.blackhole.dns_zone}"
}
