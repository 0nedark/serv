# SERV

`serv` is a helper cli application used to simplify and coordinate the startup of multiple services, all you need to do
is define a yaml file with the list of comments to execute.

## Key Features

* Clone repositories automatically into the specified folder if git project haven't been cloned yet.
* Specify a start up command for each repository
* Specify multiple health check commands for each repository
* Specify multiple post condition commands for each repository
* Block the execution of command groups while health checks and post conditions are executing
* Ability to run app in verbose mode as well as in silent mode.

## Installation

Install golang on your operating system and simply run `go get -u github.com/0nedark/serv` to install `serv`.

## Verify installation

To test that `serv` is installed correctly run `which serv` you should see the path where the application is installed.
If `which serv` produces no output then you probably need to `export PATH=$PATH:$GOPATH/bin` add this export to your
`~/.bashrc` or equivalent.

## Configuration

#### serv.yml configuration

```yaml
order: ['group_a', 'group_b']
groups:
  group_a: [{...}, ...]
  group_b: [{...}, ...]
```

The start up of services is configured with the `serv.yaml` file. The root of this yaml file defines `order`, `groups`

#### order
Order is an array of string where each string refers to a specific group. This is used to define te order in which
groups are executed.

#### groups

```yaml
group_a:
  - url: git@github.com:0nedark/repo-name.git
    path: ../B
    command: echo test
    healthchecks: [{...}, ...]
    postconditions: [{...}, ...]
    
  - command: echo test
    healthchecks: [{...}, ...]
    postconditions: [{...}, ...]
```

Groups are a key value pair where the key is the name of the group and the value defines an array of `services`. Groups
are executed sequentially in the `order` defined. Consecutive group is not executed until all health checks and post
conditions of the current group pass.

#### service

```yaml
- url: git@github.com:0nedark/repo-name.git
  path: ../B
  command: echo test
  healthchecks: [{...}, ...]
  postconditions: [{...}, ...]
```

Each service object defines the following fields `command`, `healthchecks` and `postconditions` as well as two optional
fields `url` and `path`. The two optional fields can be used to define a repository that needs to be cloned, `url`
defines repository url and `path` specifies where to clone the repository. If your repository is private you should
still be able to clone it as long as you use the `ssh` repo `url` and you have `ssh-agent` running with your credentials
added with `ssh-add`. In addition, `path` is also used to define the context in which all commands will be executed for
this service, `path` and `url` must be provided together and you can't provide `path` on it's own. If no `path` is provided then
the current working directory will be used as the context.

#### command
```yaml
command: echo test
```
Command defines a single shell command to be executed. This should be used to run you startup script. If this command fails `serv` will terminate with error.

#### healthchecks
```yaml
healthchecks:
  - command: echo test
    timeout: 10
    sleep: 5
  - command: cd .
    timeout: 60
    sleep: 5
```

A list of command like objects that are executed as soon as the `service`'s `command` returns with no errors. Each health
check command in the current `group` is executed in parallel, repeatedly, every `sleep` number of seconds, until it
either returns with success or a `timeout` is reached. If any of the health checks fail `serv` will terminate with
error.

#### postconditions
A list of commands that are executed as soon as the `healthchecks` return with no errors. Each post condition command in
the current `group` is executed in parallel. If any of the post conditions fail `serv` will terminate with error.

# serv.yml example
```yaml
order: [group_a, group_b]
groups:
  group_a:
    - command: echo starting A
      healthchecks:
        - command: echo "healthcheck A"
          timeout: 10
          sleep: 5
      postconditions:
        - command: echo postcondition A
    - url: https://github.com/0nedark/serv.git
      path: ../B
      command: echo starting B
      healthchecks:
        - command: echo healthcheck B
          timeout: 60
          sleep: 5

  group_b:
    - url: https://github.com/0nedark/serv.git
      path: ../C
      command: echo starting C
      healthchecks:
        - command: echo healthcheck C
          timeout: 60
          sleep: 5
```