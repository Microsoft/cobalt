data "azurerm_resource_group" "appsvc" {
  name      = "${var.service_plan_resource_group_name}"
}

data "azurerm_app_service_plan" "appsvc" {
  name                    = "${var.service_plan_name}"
  resource_group_name     = "${data.azurerm_resource_group.appsvc.name}"
}

resource "azurerm_app_service" "appsvc" {
  name                = "${var.app_service_name[count.index]}"
  resource_group_name = "${data.azurerm_resource_group.appsvc.name}"
  location            = "${data.azurerm_resource_group.appsvc.location}"
  app_service_plan_id = "${data.azurerm_app_service_plan.appsvc.id}"
  tags                = "${var.resource_tags}"
  count               = "${length(var.app_service_name)}"

  app_settings {
    DOCKER_REGISTRY_SERVER_URL      = "${var.docker_registry_server_url[count.index]}"
    DOCKER_REGISTRY_SERVER_USERNAME = "${var.docker_registry_server_username[count.index]}"
    DOCKER_REGISTRY_SERVER_PASSWORD = "${var.docker_registry_server_password[count.index]}"
  }

  site_config {
    linux_fx_version = "${var.site_config_linux_fx_version[count.index]}"
    always_on        = "${var.site_config_always_on[count.index]}"
  }
}
