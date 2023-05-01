# hrobot: Command-line interface for Hetzner Robot

`hrobot` is a command-line interface for interacting with Hetzner Robot (Hetzner Dedicated Servers).

> Please note this is not an official Hetzner product, the author is not in any way affiliated with Hetzner use at own risk!  

If you are looking for the [Hetzner Cloud](https://cloud.hetzner.com) cli you can check out the [official hcloud-cli](https://github.com/hetznercloud/cli) maintained by Hetzner. 

## Installation 

To install the hrobot-cli on your system you need to build it manually. If theres more interest in the project, I might add prebuild solution. 

### Build manually 

You need to have Go installed on your system. To build the latest version of `hrobot` you can simply run: 

``` bash
go install github.com/floshodan/hrobot/cli/cmd/hrobot@latest
```

## Getting started

The CLI looks for a Enviromental variable **HROBOT_TOKEN**, which has the following structure "username:password"

To use the Token as an enviroment variable as in the example above you can export a variable: `export HROBOT_TOKEN="username:password"` in your terminal. 
To make it persitent on your system you can put the export command in your `~/.profile` file.

### Running hrobot-cli 

After installation you can simply run: 

``` bash
hrobot 
```

to see all available commands. 
