// ---- Admin Resources ----

output "admin_resource_group" {
  value = azurerm_resource_group.admin_rg.name
}

output "fqdns" {
  value = [
    for uri in module.app_service.app_service_uri :
    "http://${uri}"
  ]
}

output "app_insights_id" {
  value = module.app_insights.app_insights_app_id
}

output "service_plan_name" {
  value = module.service_plan.service_plan_name
}

output "service_plan_id" {
  value = module.service_plan.app_service_plan_id
}


// ---- App Dev Resources ----

output "app_dev_resource_group" {
  value = azurerm_resource_group.app_rg.name
}

output "keyvault_name" {
  value = module.keyvault.keyvault_name
}
