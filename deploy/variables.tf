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
}

variable "zone" {
  type        = string
  description = "The zone where resources will be created"
}
