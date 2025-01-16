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
  project_id = var.project_id
  region     = var.region
  zone       = var.zone
}

resource "scaleway_instance_ip" "public_ip" {
  project_id = var.project_id
}

resource "scaleway_instance_volume" "data" {
  project_id = var.project_id
  size_in_gb = 20
  type       = "l_ssd"
  tags       = ["bh", "si"]
  name       = "blackhole-volume"
  zone       = var.zone
}

resource "scaleway_instance_security_group" "blackhole" {
  project_id              = var.project_id
  inbound_default_policy  = "drop"
  outbound_default_policy = "accept"

  inbound_rule {
    action   = "accept"
    port     = "22"
    ip_range = "0.0.0.0/0"
  }

  inbound_rule {
    action = "accept"
    port   = "80"
  }

  inbound_rule {
    action = "accept"
    port   = "443"
  }
}

resource "scaleway_instance_server" "blackhole" {
  name       = "blackhole"
  project_id = var.project_id
  type       = "DEV1-S"
  image      = "ubuntu_focal"

  tags = ["bh", "blackhole", "si"]

  ip_id = scaleway_instance_ip.public_ip.id

  root_volume {
    size_in_gb = 20
  }

  security_group_id = scaleway_instance_security_group.blackhole.id
}
