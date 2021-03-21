
- Let's follow the docs! Following the docs is always a good idea.
- https://docs.github.com/en/actions/guides

## 1. Introduction to Github Actions
- https://docs.github.com/en/actions/learn-github-actions/introduction-to-github-actions?learn=getting_started

### Overview
- Github actions are event driven
- A series of commands are ran after a specific event has occurred (PR, push, etc)

### Components of Github Actions

#### Workflows
- An automated procedure that you add to your repo. Made up of one or more jobs
- Can be used to build, test, package, release etc

#### Events
- A specific activity that triggers a workflow

#### Jobs
- Set of steps that execute on the same runner
- By default, a workflow with multiple jobs will run those in parallel

#### Steps
- An individual task that can run commands in a job
- Can be either an action or a shell command. Each step in a job executes in the same runner, allowing the actions in that job to share data with each other

#### Actions
- Standalone commands that are combined into steps to create a job. Smallest portable building block of a workflow
- You can create your own actions or use existing ones.

#### Runners
- A server that has the github actions runner installed

### Create an example workflow
- See `.github/workflows/learn-github-actions.yml` on your gh study repo

### 2. Finding and customizing actions
- You can use actions that already exist, or create your own!
- Actions can be on: Public repo, same repo as workflow, a published docker container on dockerhub
- When editing a github actions yaml file on the github UI, you can browse github actions from the marketplace and see their syntax
- You can also use release management for custom action, ex:
```yaml
steps:
    - uses: actions/javascript-action@v1.0.1
```
#### Using inputs and outputs with an action
- An action often accepts or requires inputs, and it also may generate output you can use

