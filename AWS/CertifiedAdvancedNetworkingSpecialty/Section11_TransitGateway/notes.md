## 64. Introduction to Transit Gateway
- Transit gateway: Enables full mesh connectivity between VPCs

### What is transit gateway? 
- Solution that allow customers to connect thousands of VPCs and on-premise networks
- VPC peering connection is only one-to-one. TGW can connect 10 VPCs. Also allows you to connect these VPCs to your on prem network

Transit Gateway (TGW) supports attachments to:
- One or more VPCs
- VPN
- Direct Connect Gateway
- Peering connection to another transit gateway
- A SD-WAN/third party network appliance

## Multiple VPCs 
- If you did VPC peering between 6 VPCs you'd need 15 peering connections, wew, not good
- Transit gateway allows you to connect them in a hub-and-spoke model, connecting everything to a TGW

## Multiple VPCs and a VPN
- Instead of connecting every VPC to the VPN, you can connect all the VPCs to a TGW and then connect the TGW to the VPN (so single VPN-to-TGW connection)

You can also peer your AWS transit gateways to multiple regions, interesting (so: VPC - TGW - TGW - VPC)