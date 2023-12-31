# nxrmuploader

A tool to upload your binary packages to your own Nexus Repository Manager (NxRM)

## Description

This tool reads a configuration file that describes all of your NxRM repos, and will upload any Debian (.DEB) or RHEL-based (RPM) binary package to the appropriate repo on your NxRM host.

You need a JSON-formatted configuration file that describes your repos, credentials, etc. More on that below.

Once that file established, you can select the file to upload, and where to upload it.

## Configuration

A typical JSON-based config file looks like this:

'''json
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
'''
