# Klaassify - An External Node Classifier for puppet

Klaassify adds and removes nodes from and to an external (Gitlab)
node classifier.

## Building

Build it!

```shell
  go build -o build/klaassify
```

## Running

Run it!

### Prereqs

Klaassify uses the following environment variables to connect the ENC backend
(Gitlab). Gitlab API key can be retieved from the profile of a Gitlab account
that administrator privileges.

```shell
  export GITLAB_API_KEY=
  export GITLAB_API_URL="http://gitlab.com/api/v3/"
  export ENC_GIT_REPOSITORY="enc"
  export ENC_GIT_NAMESPACE="zmm-infra"
```

### Help

```shell
  klaassify --help
```

### Adding nodes

Add nodes by calling `klaassify add`

```shell
  klaassify add --host localhost.localdomain -e prd -c jenkins_server -t build -p project -size large
```

### Removing nodes

Remove nodes from the ENC by calling `klaassify remove`

```shell
  klaassify remove --host localhost.domain
```
