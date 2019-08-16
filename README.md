# AWS-INTRODUCTION

This repository is a learning aid to creating AWS CodePipelines from code templates.

## Pre-requisites

Create an account on AWS and install the packages below.

### Packages

* [AWS CLI](https://github.com/aws/aws-cli)
* [GO](https://golang.org/doc/install)
* [Make](https://www.gnu.org/software/make/)

### Configuring the AWS CLI

You can optionally use [aws-mfa](https://github.com/broamski/aws-mfa) as a mechanism for handling MFA with your AWS account when using the CLI.

To configure the CLI, use the documentation [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html).

## Getting started

make create-stack.

## Additional reading

Links to the pertinent documentation are commented throughout the code. The following are links to useful general reading.

### Git

* [Creating a new repository](https://kbroman.org/github_tutorial/pages/init.html)
* [Creating a PAT for the CLI](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)

### AWS CFN intrinsic functions

* [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)
* [GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html)
