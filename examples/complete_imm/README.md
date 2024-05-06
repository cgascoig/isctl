# ISCTL based IMM deployments

This repository holds examples for deploying an entire IMM domain through isctl.
The examples will create all pools, policies and profiles and uses a shared organization called LAB-RESOURCES for storing the main configuration. 
The Server Pool and Profiles will be created in the organization called IMM-Troubleshooting, this is also the org where you should have claimed your Fabric Interconnect.
It is possible to change the code to use a single organization.

There are 3 directories in this repository

### 1_domain_profile
This holds all the domain required policies and the Domain profile.
There is a seperate README in this directory with more information

### 2_chassis_profile
This holds all the chassis required policies and the Chassis profile.
There is a seperate README in this directory with more information

### 3_server_profiles
This holds all the server profile required policies, the Server Profile Template and it will derivef 8 Server Profiles that will be assigned through a Server Resource Pool.
There is a seperate README in this directory with more information

All configuration is stored in yaml files and need to be updated to match your environment.

**YOU MUST EDIT ALL YAML FILES TO MATCH YOUR ENVIRONMENT**

## Pre Requisites
- Have an Intersight account (SaaS or Appliance)
- Have 2 different orgs. One for storing most of the config which is called LAB-RESOURCES and is shared with the second organization called IMM-Troubleshooting which is where the FI targets are claimed and we will store the Server Profiles and Resource Pool.
This is optional, you can also update the config to work with a single org.
- Make sure you have [isctl](https://github.com/cgascoig/isctl) installed and configured with an API key that has access to both organizations.
- Claim your FI Target into the right organization 

## Quick - Steps
It is highly recommended to go directory by directory but here are all the deployment steps.
The clean up steps can be found inside each directory. 

1. Edit all the yaml files in all directories to match your environment

2. Push the domain configuration 
```
isctl apply -f 1_domain_profile/.
```

3. Assign the Domain Profile through the UI and deploy it.

4. Push the chassis configuration 
```
isctl apply -f 2_chassis_profile/.
```

5. Assign the Chassis Profile through the UI and deploy it.

6. Push the server configuration
```
isctl apply -f 3_server_profiles/FI-Attached/.
```

7. Deploy the server profiles through the UI

## Clean up

there are more details cleanup steps writting in each directory, please clean up in the right order.