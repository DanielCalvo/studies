#!/bin/bash

HOSTED_ZONE_ID=""
NAME=""
TYPE="A"
TTL=60

IP=$(curl http://checkip.amazonaws.com/)

if [[ ! $IP =~ ^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$ ]]; then
  echo "Invalid IP"
	exit 1
fi

cat > /tmp/route53_changes.json << EOF
    {
      "Comment":"Updated from orangepi",
      "Changes":[
        {
          "Action":"UPSERT",
          "ResourceRecordSet":{
            "ResourceRecords":[
              {
                "Value":"$IP"
              }
            ],
            "Name":"$NAME",
            "Type":"$TYPE",
            "TTL":$TTL
          }
        }
      ]
    }
EOF

aws route53 change-resource-record-sets --hosted-zone-id $HOSTED_ZONE_ID --change-batch file:///tmp/route53_changes.json >> /dev/null
