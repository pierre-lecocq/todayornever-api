# Deploy

## Create a dedicated user

```sh
mkdir -p /home/todayornever/.ssh
cp ~/.ssh/authorized_keys /home/todayornever/.ssh/authorized_keys
useradd -d /home/todayornever todayornever
usermod -aG sudo todayornever
chown -R todayornever:todayornever /home/todayornever/
chown root:root /home/todayornever
chmod 700 /home/todayornever/.ssh
chmod 644 /home/todayornever/.ssh/authorized_keys
passwd todayornever
```

## Install needed software

```sh
apt-get install nginx ssl-cert golang-go

snap install core
snap refresh core
snap install --classic certbot
ln -s /snap/bin/certbot /usr/bin/certbot
```

## Prepare the backend

1. Create the `.env` file
2. Create the database file and run the migrations
3. Compile the binary with `env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./prod/todayornever-api .`
4. Put these 3 files into /home/todayornever/todayornever-api (`rsync -avz <src> <user@host>/home/todayornever/todayornever-api/`)


## Set a systemd service

Create `/lib/systemd/system/todayornever-api.service` with:

```
[Unit]
Description=todayornever-api
ConditionPathExists=/home/todayornever/todayornever-api/todayornever-api
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=5
ExecStart=/home/todayornever/todayornever-api/todayornever-api
WorkingDirectory=/home/todayornever/todayornever-api
User=todayornever
Group=todayornever

[Install]
WantedBy=multi-user.target
```

Start the service with:

```sh
systemctl enable todayornever-api.service
systemctl start todayornever-api
systemctl status todayornever-api
journalctl -f -u todayornever-api
```

If needed, restart the daemon with:

```sh
systemctl daemon-reload
```

## Create a nginx site

Install with:

```sh
apt-get install nginx ssl-cert
```

Remove the default site:

```sh
rm /etc/nginx/sites-enabled/default
```

Create `/etc/nginx/sites-available/todayornever-api` with:

```
server {
    listen 80;
    server_name todayornever-api.flyingstack.com;

    location / {
        proxy_pass http://localhost:8080;
    }
}
```

Enable the site with:

```sh
ln -s /etc/nginx/sites-available/todayornever-api /etc/nginx/sites-enabled/todayornever-api
```

Claim a SSL ceritificate with:

```sh
certbot --nginx -d todayornever-api.flyingstack.com # -d todayornever.flyingstack.com
```

Restart nginx with:

```sh
nginx -s reload
```

## Final checks

```sh
service todayornever-api status
service nginx status
service snap.certbot.renew status
curl http://localhost:8080
```
