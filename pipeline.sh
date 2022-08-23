#!/bin/bash
set -ex

# This script is a placeholder for a real CI/CD engine's pipeline.
# During the interview, this script will be used deploy the application inside your provided cluster
#
# This script can be used to call any other script, in any language
# Please, edit this script in any way you need so your deployment can be achieved

# Exercise related vars
#
# Please edit the following variables to target the project that was given to you.
CANDIDATE_NAME="<Carlos>-<Abreu>"
DEV_IP="<34.111.203.197>"
PRD_IP="<34.117.6.83>"

GOOGLE_PROJECT="dv-${CANDIDATE_NAME}-1"
GOOGLE_REGION="us-central1"
GKE_CLUSTER_NAME="test-gke-cluster"
DEV_URL=http://dev.${DEV_IP}.sslip.io
PRD_URL=https://prd.${PRD_IP}.sslip.io

DOCKER_REGISTRY_ENDPOINT="us.gcr.io/${GOOGLE_PROJECT}"

export CANDIDATE_NAME GOOGLE_PROJECT GOOGLE_REGION GKE_CLUSTER_NAME DEV_URL PRD_URL DOCKER_REGISTRY_ENDPOINT

# The VERSION will be used in the STEP FOUR of the test
# Defaults to version 1.0.0 which correspond to the public Docker image available
VERSION="${1-1.0.0}"

#################################################################################################################################

# Step One
function step_one {
  # build the Hello Docker image
  #    Add the steps to build the Docker image of the Hello application
  #    from the provided Dockerfile
  echo step_one
}

# Step Two
function step_two {
  # Deploy to dev namespace
  #    Add the commands to deploy the application in the `dev` namespace of the Kubernetes cluster
  #    After this step you should be able to reach the application at the http://dev.<HTTP ingress IP>.sslip.io URL
  echo step_two
}

# Step Three
function step_three {
  # Deploy to prd namespace
  #    Add the commands to deploy the application in the `prd` namespace of the Kubernetes cluster
  #    You will have to use an SSL certificate for this step (can be self-signed or cloud-generated, up to you)
  #    After this step you should be able to reach the application at the https://prd.<HTTPS ingress IP>.sslip.io URL
  echo step_three
}

# Step Four
  # Use a dynamic version
  #    Update the 3 previous steps to support a dynamic version
  #    the VERSION variable is already defined. You set it by calling this script with a parameter:
  #    ./pipeline.sh 1.2.3   # set VERSION=1.2.3



step_one
step_two
step_three
