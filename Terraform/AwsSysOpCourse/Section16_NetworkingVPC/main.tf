provider "aws" {
  profile    = "default"
  region     = "eu-west-1"
}

resource "aws_vpc" "DemoVPC" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "DemoVPC"
  }
}

resource "aws_subnet" "PublicSubnetA" {
  vpc_id     = aws_vpc.DemoVPC.id
  cidr_block = "10.0.0.0/24"
  availability_zone = "eu-west-1a"
  map_public_ip_on_launch = true
tags = {
    Name = "PublicSubnetA"
  }
}

resource "aws_subnet" "PublicSubnetB" {
  vpc_id     = aws_vpc.DemoVPC.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "eu-west-1b"
  map_public_ip_on_launch = true
  tags = {
    Name = "PublicSubnetB"
  }
}

resource "aws_subnet" "PrivateSubnetA" {
  vpc_id     = aws_vpc.DemoVPC.id
  cidr_block = "10.0.16.0/20"
  availability_zone = "eu-west-1a"
  tags = {
    Name = "PrivateSubnetA"
  }
}

resource "aws_subnet" "PrivateSubnetB" {
  vpc_id     = aws_vpc.DemoVPC.id
  cidr_block = "10.0.32.0/20"
  availability_zone = "eu-west-1b"
  tags = {
    Name = "PrivateSubnetB"
  }
}

resource "aws_internet_gateway" "DemoIGW" {
  vpc_id = aws_vpc.DemoVPC.id

  tags = {
    Name = "DemoIGW"
  }
}

//Apparently the "local" route comes for free
resource "aws_route_table" "PublicRouteTable" {
  vpc_id = aws_vpc.DemoVPC.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.DemoIGW.id
  }
  tags = {
    Name = "PublicRouteTable"
  }
}

resource "aws_route_table_association" "RouteAssociationPublicSubnetA" {
  subnet_id      = aws_subnet.PublicSubnetA.id
  route_table_id = aws_route_table.PublicRouteTable.id
}

resource "aws_route_table_association" "RouteAssociationPublicSubnetB" {
  subnet_id      = aws_subnet.PublicSubnetB.id
  route_table_id = aws_route_table.PublicRouteTable.id
}

//
//resource "aws_route_table" "PrivateRouteTable" {
//  vpc_id = aws_vpc.DemoVPC.id
//  tags = {
//    Name = "PrivateRouteTable"
//  }
//}