#!/bin/bash


crontab -l > deliveryBotCron
echo "15 10 * * 1-5 cd ~/GoLandProjects/beta/ && ./main" >> deliveryBotCron
crontab deliveryBotCron
