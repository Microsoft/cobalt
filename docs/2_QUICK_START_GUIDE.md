# 2. Quick Start Guide

## 2.1 Overview

*Cobalt* is a tool for developers who are interested in reusing or contributing new cloud infrastructure as code patterns in template form. An actualized infrastructure as code pattern in Cobalt is called a *Cobalt Infrastructure Template* or *CIT* (/kɪt/). Cobalt Infrastructure Templates primarely rely on [*Terraform*](https://learn.hashicorp.com/terraform)'s HCL language in order to target a wide array of cloud providers.

You can get pretty creative and build your own custom *CIT*s in order to use and/or contribute to Cobalt but we strongly recommend that you first complete this quick start guide. This guide is centered around our existing [*Azure Hello World CIT*](../infra/templates/az-hello-world/README.md "AZ Hello World - Cobalt Infrastructure Template") and should serve as your first Azure infrastructure deployment. In summary, completing this guide should be your first major step in familiarizing yourself with Cobalt and the *CIT* developer workflow. Happy templating! 😄

> For a more general overview of Cobalt, please visit our main page: [READ ME](https://github.com/microsoft/cobalt/blob/master/README.md "Main Cobalt Read Me")

## 2.2 Goals and Objectives

🔲 Prepare local environment for *Cobalt Infrastructure Template* deployments.

🔲 Deploy the [*Azure Hello World CIT*](../infra/templates/az-hello-world/README.md "AZ Hello World - Cobalt Infrastructure Template").

🔲 Walk away with a introductory understanding of the *CIT* developer workflow.

🔲 Feel confident in moving forward to our next recommended section: *[Cobalt Templating from Scratch](https://github.com/microsoft/cobalt/blob/master/docs/3_NEW_TEMPLATE.md).*

## 2.3 Prerequisites

> NOTE: Previous "infrastructure as code" experience is not a prerequisite for completing the quick start guide.

| Azure Prereqs | Description | I | Install Prereqs | Description |
|----------|--------------|-|----------|--------------|
| `Azure Subscription` | [Azure Portal](https://portal.azure.com/) - This template needs to deploy infrastructure within an Azure subscription.|I|`Terminal with bash shell`|[WSL](https://code.visualstudio.com/docs/remote/wsl) or [Git Bash](https://git-scm.com/downloads) - The shell environment needed to follow along with the provided instructions.|
|`Azure Service Principal`|[Azure Service Principal](https://docs.microsoft.com/en-us/azure/active-directory/develop/howto-create-service-principal-portal) - This template needs permissions to deploy infrastructure within an Azure subscription.|I|`Git Source Control`|[Install git](https://www.atlassian.com/git/tutorials/install-git)|
|`Azure Storage Account`|[Azure Storage Account](https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview) - An account for tracking terraform remote backend state. You can use our backend state setup [template](../infra/templates/backend-state-setup/README.md) to provision the storage resources.|I|`Terraform`|[Terraform](https://www.terraform.io/downloads.html) - Download the appropriate environment for setting up Terraform|
|`Azure CLI`|[Get started with Azure CLI](https://docs.microsoft.com/en-us/cli/azure/get-started-with-azure-cli?view=azure-cli-latest)- An account for tracking terraform remote backend state. You can use our backend state setup [template](../infra/templates/backend-state-setup/README.md) to provision the storage resources.|I|`GitHub Account`|[Github](https://github.com/login) - An account for forking the Cobalt repo|

---

## 2.4 Deploying Cobalt's [_Azure Hello World CIT_](../infra/templates/az-hello-world/README.md)

Below are the steps for deploying the [_Azure Hello World CIT_](../infra/templates/az-hello-world/README.md). Ensure that this deployment outputs the following Azure cloud resources in order to call this quick start guide completed:

| Preview | Quick Start Azure Cloud Resources (pending az simple architecture diagram) |
|----------|--------------|
|![Cobalt Sample Gif](https://media.giphy.com/media/TJVF0piXxS1o4V44OD/giphy.gif) | ![image](https://user-images.githubusercontent.com/7635865/57530235-64348780-72fc-11e9-9280-9da656037c2f.png) |

<!--- These gifs will have screenshots from forking, terraform plan and apply, visit azure portal and visit app service url --->

### **Step 1:** Fork Cobalt Repo

Initializing a repository that you own and control is recommended. Once forked, move onto the next step.

* From any page within this repository, find the forked menu and fork the repo into your own repository. 

    ![image](https://user-images.githubusercontent.com/10041279/66366857-6e17f080-e957-11e9-8b32-266b0d4a98fc.png)

### **Step 2:** Clone Repo to Local Environment

You will be deploying Azure infrastructure for your local environment so you will need to have a copy of the Cobalt project locally.

* From any terminal with git, use the following git command to clone your forked repo into your local environment. You can find your git repo url at the landing page of your forked repo.

    ```bash
    git clone <insert-forked-git-repo-url> # ex. git clone https://github.com/<YourGitAccout>/cobalt.git
    ```

### **Step 3:** Setup Local Environment Variables

You'll need to define a `.env` file in the root of your local project. You can use our [environment template file](https://github.com/microsoft/cobalt/blob/master/.env.template) to start. This will hold all the environment variables needed to run your Cobalt deployments locally.

* Navigate to the root directory of your project and use the following command to copy the environment template file.

    ```bash
    cp .env.template .env
    ```

    ```bash
    $ tree infra
    ├───.env.template
    ├───.env ## NEW FILE
    ├───modules
    │   └───providers
    │       ├───azure
    │       │   ├───...
    │       │   └───vnet
    │       └───common
    └───templates
        ├───az-hello-world
        │   └───test
        └───backend-state-setup
    ```

* Provide values for the environment values in `.env` which are required to authenticate Terraform to provision resources within your subscription.

```bash
ARM_SUBSCRIPTION_ID="<az-service-principal-subscription-id>"
ARM_CLIENT_ID="<az-service-principal-client-id>"
ARM_CLIENT_SECRET="<az-service-principal-auth-secret>"
ARM_TENANT_ID="<az-service-principal-tenant>"
ARM_ACCESS_KEY="<remote-state-storage-account-primary-key>"
TF_VAR_remote_state_account="<tf-remote-state-storage-account-name>"
TF_VAR_remote_state_container="<tf-remote-state-storage-container-name>"
```

### **Step 4:** Initialize a Terraform Environment

### **Step 5:** Deploy Cobalt's [_Azure Hello World CIT_](../infra/templates/az-hello-world/README.md)

### **Step 6:** Validate App Service Deployed Successfully

### **Final Step:** Teardown Infrastructure

### **Recommended Next Step:** *[Cobalt Templating from Scratch](https://github.com/microsoft/cobalt/blob/master/docs/3_NEW_TEMPLATE.md).*
