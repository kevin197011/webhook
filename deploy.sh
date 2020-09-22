#!/usr/bin/env bash

test -d /opt/webhook && rm -rf /opt/webhook && echo "Delete old webhook repo!"
echo "Get webhook..."
cd /opt && git clone https://github.com/kevin197011/webhook.git
cd /opt/webhook
mkdir -p /etc/webhook/
useradd webhook -s /sbin/nologin
touch /var/log/webhook.log && chown -R webhook:webhook /var/log/webhook.log
cp -f ./webhook /usr/local/bin/webhook
cp -f ./templates/config.yml /etc/webhook/config.yml
cp -f ./templates/webhook.service /lib/systemd/system/webhook.service
chmod +x /usr/local/bin/webhook
echo "Start webhook..."
systemctl daemon-reload
systemctl restart webhook
echo "Show webhook status:"
systemctl status webhook
sleep 10
echo "Test send message..."
source ./check.sh
echo "Clean..."
rm -rf /opt/webhook
