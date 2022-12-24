data "cloudflare_zone" "zone" {
  name = "fdm.com.de"
}

resource "cloudflare_record" "verification" {
  for_each = {
    for dvo in aws_acm_certificate.cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      record = dvo.resource_record_value
    }
  }

  zone_id = data.cloudflare_zone.zone.zone_id
  name    = each.value.name
  value   = each.value.record
  type    = "CNAME"

  proxied = false # Take advantage of Cloudflare http caching
}

resource "cloudflare_record" "happening" {
  zone_id = data.cloudflare_zone.zone.zone_id
  name    = local.sub_domain
  type    = "CNAME"
  value   = aws_api_gateway_domain_name.happening.regional_domain_name

  proxied = true # Take advantage of Cloudflare http caching
}

resource "aws_acm_certificate" "cert" {
  domain_name       = "${local.sub_domain}.fdm.com.de"
  validation_method = "DNS"

  tags = {
    Environment = "develop"
  }
}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn = aws_acm_certificate.cert.arn
}
