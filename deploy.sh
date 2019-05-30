#!/bin/sh

glide up
gox -osarch="linux/amd64" -output="deploy/bin/application"
cd deploy
# cp ../ec2-https.config .elasticbeanstalk/
eb deploy
