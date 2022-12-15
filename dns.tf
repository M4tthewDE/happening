resource "aws_acm_certificate" "cert" {
  domain_name       = "happening.fdm.com.de"
  validation_method = "DNS"

  tags = {
    Environment = "develop"
  }

  lifecycle {
    create_before_destroy = true
  }
}

data "cloudflare_zone" "zone" {
  name = "fdm.com.de"
}

/*

resource "cloudflare_record" "happening" {
  zone_id = data.cloudflare_zone.zone.zone_id
  name    = "happening"
  type    = "CNAME"

  proxied = true # Take advantage of Cloudflare http caching
}
*/

// https://blog.viktoradam.net/2018/08/30/moving-home/

