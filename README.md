# helloappengine
helloappengine - app engine go example

Install GCP SDK
===============

Download SDK from: https://cloud.google.com/sdk/docs/#linux

    wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-192.0.0-linux-x86_64.tar.gz
    tar xf google-cloud-sdk-192.0.0-linux-x86_64.tar.gz
    ./google-cloud-sdk/install.sh

    gcloud init

Update SDK
==========

    gcloud components update

Install app engine go component for SDK
=======================================

    gcloud components install app-engine-go

Get app engine package
======================

    go get google.golang.org/appengine

Test
====

Note on GOPATH: dev_appserver.py does not run with empty GOPATH.

    GOPATH=~/go dev_appserver.py app.yaml --host=0.0.0.0 --enable_host_checking=false

Open http://localhost:8080

Deploy
======

    gcloud app deploy --version [YOUR_VERSION_ID] --no-promote --project [YOUR_PROJECT_ID]

Set project and run:

    gcloud config set project myProject
    gcloud app deploy

Or specify project:

    gcloud app deploy --project myProject

Logs
====

You can stream logs from the command line by running:

    gcloud app logs tail -s default

View
====

Open http://myProject.appspot.com/

Download
========

You can download your application code from app engine by using:

    python ~/google-cloud-sdk/platform/google_appengine/appcfg.py -A myProject -V [YOUR_VERSION_ID] download_app [OUTPUT_DIR] --noauth_local_webserver

