#!/bin/bash
amazon-linux-extras install docker -y
service docker start
docker pull ishaanrao/fragbots:latest
