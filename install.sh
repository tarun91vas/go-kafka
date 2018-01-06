#!/bin/sh

####################Installation Script##############
#NOTE: Update the verison number to install specific version


##Kafka installation path
path=$1
echo "Creating directory: "$path
mkdir -p $path
cd $path
wget http://www-us.apache.org/dist/kafka/1.0.0/kafka_2.11-1.0.0.tgz
echo "Installing...."
tar -xvf kafka_2.11-1.0.0.tgz
echo "Complete!"


