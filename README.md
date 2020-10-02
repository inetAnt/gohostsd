# gohostsd

Takes a bunch of files in `/etc/hosts.d/*` and compile them into a `/etc/hosts` file.
It will watch the config directory for any change and refresh the hosts file.

## Tests

None :( (not yet at least)

## Deployment

There are no package available yet, just a binary that you can run with `systemd`, see the [example file](https://github.com/inetAnt/gohostsd/blob/master/gohostsd.service).

You simply have to copy the file to `/etc/systemd/system/`, run `systemctl daemon-reload` and `systemctl enable --now gohostsd.service` to both enable and start the service.   
