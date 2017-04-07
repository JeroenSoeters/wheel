# Wheel

[![wercker status](https://app.wercker.com/status/6ef859f0c88b3e5c33b25894bdab2aa0/s/master "wercker status")](https://app.wercker.com/project/byKey/6ef859f0c88b3e5c33b25894bdab2aa0) [![Go Report Card](https://goreportcard.com/badge/github.com/JeroenSoeters/wheel)](https://goreportcard.com/report/github.com/JeroenSoeters/wheel)

## Mission Statement 
Bootstrap a micro services system in minutes vs months.

## Why
Tools for reliably delivering a micro services system to production these days are abundant. We have build systems like gradle, rake, sbt to compile and test our code. There is Docker and container orchestration frameworks like Swarm, Kubernetes, Nomad and the like for deploying and scaling a set of services. We have automation servers like Jenkins and GoCD that allow us to integrate and test our code with other people's code. Consumer driven contract infrastructure (like Pact and pact-broker) allow us to confidently version our API's without breaking downstream consumers. There are deployment tools, Spinnaker being the latest and greatest, that allow us to smoothly deploy newer versions of our code to production, as well as supporting a safe roll-back mechanism when things go wrong. Database migration tools follow a similar pattern for the data stores we use. We've seen a Cambrian explosion of cloud providers. There's configuration management, log aggregation, metrics collection, secret management and the list goes on and on.

Installation and configuration of all these tools is time consuming and requires specialized knowledge, yet tools for automating this process are widely available in the form of infrastructure as code (ansible, chef, puppet) and pipelines as code (jenkinsfile, gomatic) and it turns out we're actually quite good at this automation! Yet for some reason we really suck at capturing these artifacts in a reusable form so that other teams can quickly bootstrap.

Meet Wheel! Wheel aims at bootstrapping a micro services system in minutes instead of months. Adding a new service requires a single command at the command line instead of a week's worth of configuring pipelines, docker compose, kubernetes etc. You tell Wheel which tool chain you prefer and Wheel does the work for you. It aims to be tech-stack agnostic, extensible and very much pick-and-choose.

Wheel is a command line util that orchestrates infrastructure as code, pipelines as code as well as a local development environment with the aim of drastically improving team productivity.

## Initial tech stack that will supported in an MVP
* AWS
* DC/OS
* Kubernetes
* Jenkins
* Nexus
* Spinnaker

## Templates
* java-maven
* java-gradle
* node
* scala-sbt

## Setup and Build

```
make depend && make
```

## Testing

```
make test
```

## Contributing
For contributing please reach out at our slack: wheel-group.slack.com
