#!/bin/bash

for iface in $(ip -o link show | awk -F': ' '{print $2}' | grep '^dummy'); do
    sudo ip link delete "$iface"
done
