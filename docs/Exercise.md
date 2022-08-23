# Exercise Instructions

`Hello` is a small application written in Go by the `Greetings` team. It's a simple web server with few API endpoints and memory storage.
You are part of the Delivery Engineering team and as such you're in charge of building a CI/CD pipeline that will deploy the application in an existing Kubernetes cluster.

The `Hello` application must be deployed in both a `dev` namespace and a `prd` (production) namespace.

The marketing team already promoted this new service to our customers, so it must be ready for public access in two hours.

To get started quickly, another member of your team added you to the Google Cloud project so you can connect to the Kubernetes (GKE) cluster and the Google Container Registry (GCR).
They also created two static external IP addresses and two DNS names to be used for the public Dev and Prd URLs.

The developer team also shared the `Hello` application inside the [hello](/hello) folder its documentation in the [Hello.md](Hello.md) file.

## Deployment Script Specifications

You will be creating a deployment pipeline script, in the file `pipeline.sh`, which you'll be running locally on your laptop to demonstrate during your interview. It is supposed to simulate a CI/CD pipeline tool.

- User your own google Cloud accout to connect to the project and run your pipeline. Please don't create new Google Service Account or keys. You will not be asked to share your account, our own project admin account is able to use the project
- You can either set your whole pipeline in the `pipeline.sh` file or just call another script from there if you prefer to use another language
- You don't need any git repo or trigger event. The `pipeline.sh` script will suppose that we automatically build, deploy to dev then deploy to prd in sequence
- The `pipeline.sh` script already includes the steps skeleton that we want you to create
- use any tool you need that can be installed on a UNIX host (Linux, OsX) and document the choice for this tool, where to get it and how to install/configure it, if necessary. No need to describe well known tools like `kubectl` or `docker`
- add comments into your scripts so it is easy to understand each section, and/or create a `README` file explaining how to run your pipeline.

We expect that you will build one step after the other and go as far as possible while not exceeding the **two hours total exercise time**.

We insist you don't go over the ~2 hours. We want you to demonstrate your CI/CD and Ops knowledge, not build a multi-milion company pipeline and a full infrastructure.

The script should execute the following steps:

### Step one - build and push a docker image to GCR

In this step, you'll be building and pushing a docker image for the `hello` application to GCR.
Using the provided `Dockerfile`, you'll want to build the image and push it to the project's docker image registry. Feel free to update the Dockerfile if you feel it is not production-ready.

At the end, the Google Container registry in your project should contain a version of the docker image of the `hello` application that you will use for deployments in the next steps.

### Step two - deploy to dev

Deploy the application in a `dev` namespace of the Kubernetes cluster. Make it reachable on the public URL `http://dev.<HTTP DEV ingress IP>.sslip.io`.

### Step Three - deploy to prd

Deploy the application in the `prd` namespace, with a public URL `https://prd.<HTTPS PRD ingress IP>.sslip.io` that answers to HTTPS (TLS) requests.
You will have to create an SSL Certificate to be able to achieve this step. It's up to you to create a self-signed certificate or use a free SSL authority. Be prepared to explain your decision.

### BONUS: Step Four

We request that you don't spend more than two hours doing this exercise. If you've been fast for the first three steps, you can also work on this fourth one:

Update your `pipeline.sh` script to use the `VERSION` environment variable that represents the version of the `hello` application (like a semver or a GIT commit SHA).
This `VERSION` variable will be used to create the Docker image `TAG` and to templatize the Kubernetes deployments.

The provided `pipeline.sh` script already define the `VERSION` variable to be equal to the first argument when executing the script. See the comments for more information.

### Notes

The [IP.sslip.io](https://sslip.io) URLs creates a dynamic DNS `A Record` that points back to the embedded IP. It's a convenient way to get a dynamic public DNS name when you only have an IP address.

## Review process

When you're done, please, don't destroy your project. We may ask you to demonstrate it during the next interview.
Please send back a tar.gz file including the whole exercise folder. This can be achieved with the command:

```shell
tar zcvf <first-name>-<last-name>-exercise.tar.gz -C .. dv-interview-exercise

```

Once received and reviewed we will setup a review interview. During this interview, we will ask you to:

1. run the `pipeline.sh` script from your computer to showcase your work
1. explain how you worked on this exercise and how you based your decisions
1. answer specific questions about different parts of the exercise, like CI, CD, Docker and Kubernetes specifics
1. what else would you have done to improve this solution given more time (can be features you didn't have time to get to during the exercise, as well as thinking beyond the boundaries of this exercise)

As we review your assignment, we’ll be paying particular attention to the simplicity, clarity, best practices and organization of your implementation. We’ll also be looking for an over-reliance on unnecessary third party tools and packages.
The structure of your `pipeline.sh` will be checked for clarity and maintainability.

Use this exercise to showcase your knowledge around CI/CD. Try not to go over the two hour exercise deadline.

Don't limit your answers to the content of this exercise during the interview.
