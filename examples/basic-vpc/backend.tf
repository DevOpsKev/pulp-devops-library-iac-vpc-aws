# --------------------------------------------------------------------------------------------------------------------
# --------------------------------------------------------------------------------------------------------------------
# BACK END
# --------------------------------------------------------------------------------------------------------------------
# --------------------------------------------------------------------------------------------------------------------

terraform {
  backend "remote" {
    organization = "pulp-testing"

    workspaces {
      name = "pulp-devops-library-iac-vpc-aws"
    }
  }
}