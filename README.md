rpenv
=====

displays env vars set from existing environment and loaded from config file in specified environment (ci, qa, or prod) or executes command in the context of the existing environment variables and ones loaded from a config file.

## Usage

```
$ rpenv <env>
```

or 

```
$ rpenv <env> <cmd>
```

where `<env>` is one of `ci`, `qa`, or `prod` (`production` should also work) and `<cmd>` is the desired command you wish to run. If called without a `<cmd>`, `rpenv` will return a list of all the env vars in the `/etc/rentpath/environment.cfg` file merged with your current environment variables (i.e. whatever `/usr/bin/env` would return). When run with a `<cmd>`, it will execute that `<cmd>` after setting the environment with the values returned if `rpenv` is run without a `<cmd>`.
