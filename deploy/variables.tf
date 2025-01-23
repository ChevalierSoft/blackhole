variable "scaleway_access_key" {
  type        = string
  description = "Scaleway access key"
  sensitive   = true
}

variable "scaleway_secret_key" {
  type        = string
  description = "Scaleway secret key"
  sensitive   = true
}

variable "project_id" {
  type        = string
  description = "Your project ID"
}

variable "region" {
  type        = string
  description = "The region where resources will be created"
  default     = "fr-par"
}

variable "host_zone" {
  type        = string
  description = "The zone where resources will be created"
  default     = "fr-par-2"
}

variable "bh_port" {
  type        = number
  description = "port used for the blackhole API"
  default     = 80
}

variable "dns_zone" {
  type        = string
  description = "domain name used to access the API"
}

variable "organization_id" {
  type = string
}
