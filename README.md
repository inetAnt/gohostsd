# gohostsd

Takes a bunch of files in `/etc/hosts.d/*` and compile them into a `/etc/hosts` file.
It will watch the config directory for any change and refresh the hosts file.

## Tests

None :( (not yet at least)

## Deployment

Very manual for now, I'll write a systemd unit file as an example once I actually deploy this code somewhere.
