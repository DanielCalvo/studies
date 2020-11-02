//it looks like using echo to test variables isn't the best approach
//How about an output and local?

variable sup {}

locals {
  users = [
    "user:some.guy@metro-markets.org",
    "user:some.otherguy@metro-markets.de",
    "user:someruser@otherdomain.com",
    "user3"
  ]


  metro_markets_developers = [
    for user in local.users :
    join(",", regexall(".+:.+@metro-markets.+", user))
  ]

  metro_markets_developers_clean = compact(local.metro_markets_developers)

}

output "myoutput" {
  //  value = local.myvar
  value = local.metro_markets_developers_clean
}

