module "vpc1" {
  source = "terraform-aws-modules/vpc/aws"
  name = "vpc1"
  cidr = "10.0.0.0/16"
  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]
#  enable_nat_gateway = true
#  single_nat_gateway = true
#  enable_vpn_gateway = true
}

module "vpc2" {
  source = "terraform-aws-modules/vpc/aws"
  name = "vpc2"
  cidr = "10.1.0.0/16"
  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.1.1.0/24", "10.1.2.0/24", "10.1.3.0/24"]
  public_subnets  = ["10.1.101.0/24", "10.1.102.0/24", "10.1.103.0/24"]
#  single_nat_gateway = true
  #  enable_nat_gateway = true
#  enable_vpn_gateway = true
}

# There's a good example to learn from here: https://github.com/terraform-aws-modules/terraform-aws-transit-gateway/blob/master/examples/complete/main.tf
module "tgw" {
  source  = "terraform-aws-modules/transit-gateway/aws"

  name            = "danis-tgw"
  description     = "My TGW connecting 2 VPCs"
#  amazon_side_asn = 64532

  transit_gateway_cidr_blocks = ["10.99.0.0/24"] //idk what this is doing though

  # When "true" there is no need for RAM resources if using multiple AWS accounts
  enable_auto_accept_shared_attachments = true

  # When "true", allows service discovery through IGMP
  enable_multicast_support = false

  vpc_attachments = {
    vpc1 = {
      vpc_id       = module.vpc1.vpc_id
      subnet_ids   = module.vpc1.private_subnets
      dns_support  = true
      ipv6_support = false
      transit_gateway_default_route_table_association = false
      transit_gateway_default_route_table_propagation = false

      tgw_routes = [
        {
          destination_cidr_block = "30.0.0.0/16"
        },
#        {
#          blackhole              = true
#          destination_cidr_block = "0.0.0.0/0"
#        }
      ]
    },
    vpc2 = {
      vpc_id     = module.vpc2.vpc_id
      subnet_ids = module.vpc2.private_subnets
      dns_support  = true
      ipv6_support = false //must be set to false if no ipv6 cidr blocks are associate with your subnet
      transit_gateway_default_route_table_association = false
      transit_gateway_default_route_table_propagation = false


      tgw_routes = [
        {
          destination_cidr_block = "50.0.0.0/16"
        },
#        {
#          blackhole              = true
#          destination_cidr_block = "10.10.10.10/32"
#        }
      ]
    },
  }

  ram_allow_external_principals = false //was true
  ram_principals                 = [] //to be an account number

}
