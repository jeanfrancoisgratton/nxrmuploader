# nxrmuploader

A tool to upload your binary packages to your own Nexus Repository Manager (NxRM)

## Description

This tool reads a configuration file that describes all of your NxRM repos, and will upload any Debian (.DEB) or RHEL-based (RPM) binary package to the appropriate repo on your NxRM host.

You need a JSON-formatted configuration file that describes your repos, credentials, etc. More on that below.

Once that file established, you can select the file to upload, and where to upload it.

## Configuration

A typical JSON-based config file looks like this:

```json
{
"YUM": [
{
"name": "YUM Repo 1",
"url": "https://nexus.example.com/repository/yum-repo1",
"username": "yum_user1",
"password": "yum_password1"
},
{
"name": "YUM Repo 2",
"url": "https://nexus.example.com/repository/yum-repo2",
"username": "yum_user2",
"password": "yum_password2"
}
],
"APT": [
{
"name": "APT Repo 1",
"url": "https://nexus.example.com/repository/apt-repo1",
"username": "apt_user1",
"password": "apt_password1"
},
{
"name": "APT Repo 2",
"url": "https://nexus.example.com/repository/apt-repo2",
"username": "apt_user2",
"password": "apt_password2"
}
]
}
```


## How to use

First, you need to create the environment file with `env add $FILENAME`; if $FILENAME is empty, it defaults to `defaultEnv.json`


Suppose you have the following files ready to be uploaded :

- mysoftware_1.00.00.deb
- mysoftware-1.00.00.rpm

You only need to run the tool like this : `nxrmUploader mysoftware_1.00.00.deb mysoftware-1.00.00.rpm` and the software will where to upload the file (based on the extension).

You use the `-i` switch when you need to specify the repository in case you have multiple repos. For instance, in the sample config file, you see multiple RPM and DEB repos; to use the second repo, you would use `-i 1` (the index is zero-based).

**ONE HUGE CAVEAT**

Whenever using `-i`, be aware that it will use that index for every repo in this session. For instance, if you wanted to use the 3rd repo, using `-i 2` would fetch the 3rd RPM repo, *but also the 3rd DEB repo*.

If that caveat is too much for your use-case, you can create separate environment files (`env add $FILENAME1`, `env add $FILENAME2`), and then select which environment file to use at runtime with the `-e` flag.

