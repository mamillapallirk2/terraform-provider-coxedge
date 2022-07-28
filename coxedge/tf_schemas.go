/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"fmt"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func getOrganizationSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"organizations": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getOrganizationSchema(),
			},
			Description: "Organization details",
		},
	}
}

func getOrganizationBillingInfoSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the billing information",
		},
		"organizations_billing_info": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getOrganizationBillingInfoSchema(),
			},
			Description: "The organization billing information",
		},
	}
}

func getOrganizationBillingInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The id of the billing information",
		},
		"organization_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_provider_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_masked_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"card_exp": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_line_one": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_line_two": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_city": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_province": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_postal_code": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"billing_address_postal_country": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func getOrganizationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The id of the organization",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name of the organization",
		},
		"entry_point": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The entry point of the organization is the subdomain of the organization in the Cox Edge URL : [entryPoint].Cox Edge",
		},
		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "Tags associated to the organization",
		},
		"service_connections": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The services for which the organization is allowed to provision resources. Includes: id,serviceCode",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Required: true,
					},
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"service_code": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}

func getEnvironmentSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"environments": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Environment descriptions",
			Elem: &schema.Resource{
				Schema: getEnvironmentSchema(),
			},
		},
	}
}

func getRolesSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"roles": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getRolesSchema(),
			},
		},
	}
}

func getRolesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"is_system": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"default_scope": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func getEnvironmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Optional:    true,
			Description: "The id of the environment.",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the new environment. Should be unique in the environment and only contain lower case characters, numbers, dashes and underscores.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The description of the environment.",
		},
		"membership": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of membership of the environment. ALL_ORG_USERS will add every user in the organization to this environment with the default role. MANY_USERS will allow you to choose the users you want in the environment and assigned them specific roles. Defaults to MANY_USERS.",
		},
		"organization_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The organization that the environment should be created in. Defaults to your organization.Required: id",
		},
		"service_connection_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The service connection that the environment should be created in. Required: id",
		},
		"creation_date": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date in ISO 8601 that the environment was created.",
		},
		"roles": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The roles of the environment and the users assigned to them. Also, defines the default role of the environment.\nrequired: name, users.id\noptional: isDefault",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Name of the role",
					},
					"is_default": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Set to default role",
					},
					"users": {
						Type:        schema.TypeList,
						Required:    true,
						Description: "Array of users Id",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func getUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Username of the new user. Should be unique across the organization.",
		},
		"first_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "First name of the user.",
		},
		"last_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Last name of the user.",
		},
		"email": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Email of the user. Should be unique across the organization.",
		},
		"organization_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Organization in which the user will be created. Defaults to your organization. Required: id",
		},
		"roles": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Role Id",
					},
				},
			},
			Optional:    true,
			Description: "The system and environment roles to give to the user.",
		},
		"last_updated": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func getWorkloadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"image": {
			Type:     schema.TypeString,
			Required: true,
		},
		"specs": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"deployment": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"pops": {
						Type: schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Required: true,
					},
					"enable_autoscaling": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},
					"instances_per_pop": {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					"max_instances_per_pop": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"min_instances_per_pop": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"cpu_utilization": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"add_anycast_ip_address": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"anycast_ip_address": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"commands": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"container_email": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"container_username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"container_password": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"container_server": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"environment_variables": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		"first_boot_ssh_key": {
			Type:     schema.TypeString,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		"ports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"protocol": {
						Type:     schema.TypeString,
						Required: true,
					},
					"public_port": {
						Type:     schema.TypeString,
						Required: true,
					},
					"public_port_desc": {
						Type:     schema.TypeString,
						Required: true,
					},
					"public_port_src": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"persistent_storages": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"path": {
						Type:     schema.TypeString,
						Required: true,
					},
					"size": {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		"secret_environment_variables": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		"slug": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func getImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"slug": {
			Type:     schema.TypeString,
			Required: true,
		},
		"family": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tag": {
			Type:     schema.TypeString,
			Required: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"reference": {
			Type:     schema.TypeString,
			Required: true,
		},
		"status": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func getImageSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"environment": {
			Type:     schema.TypeString,
			Required: true,
		},
		"images": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getImageSchema(),
			},
		},
	}
}

func getNetworkPolicyRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"workload_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network_policy_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"source": {
			Type:     schema.TypeString,
			Required: true,
		},
		"action": {
			Type:     schema.TypeString,
			Required: true,
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: true,
		},
		"port_range": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func getSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"services": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"hostname": {
			Type:     schema.TypeString,
			Required: true,
		},
		"auth_method": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"operation": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				val := false
				switch value {
				case "enable_cdn":
					val = true
					break
				case "disable_cdn":
					val = true
					break
				case "enable_waf":
					val = true
					break
				case "disable_waf":
					val = true
					break
				case "enable_scripts":
					val = true
					break
				case "disable_scripts":
					val = true
					break
				}
				if !val {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("opertaion field: %q should be either one of following - enable_cdn, disable_cdn, enable_waf, disable_waf, enable_scripts, disable_scripts", value),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		//Computed properties
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"edge_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"anycast_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"delivery_domains": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"domain": {
						Type:     schema.TypeString,
						Computed: true,
						ForceNew: true,
					},
					"validated_at": {
						Type:     schema.TypeString,
						Computed: true,
						ForceNew: true,
					},
				},
			},
		},
	}
}

func getOriginSettingSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"origin_settings": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getOriginSettingsSchema(),
			},
		},
	}
}

func getOriginSettingOriginSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"common_certificate_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"auth_method": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func getOriginSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"scope_configuration_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"websockets_enabled": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"ssl_validation_enabled": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"pull_protocol": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"host_header": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"origin": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: getOriginSettingOriginSchema(),
			},
			Optional: true,
		},
		"backup_origin_enabled": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"backup_origin": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: getOriginSettingOriginSchema(),
			},
			Optional: true,
		},
		"backup_origin_exclude_codes": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func getDeliveryDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The delivery domain unique identifier.",
		},
		"stack_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the stack that the site belongs to.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"domain": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The site's domain name.",
		},
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the site for which to list delivery domains. This parameter is required.",
		},
	}
}

func getCDNSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A site's unique identifier.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"cache_expire_policy": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A site's cache expiry policy. Can be ORIGIN_CONTROLLED, SPECIFY_CDN_TTL, NEVER_EXPIRE, or DO_NOT_CACHE.",
		},
		"cache_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The time to live for the cache, in seconds. Depends on the cache expiry policy.",
		},
		"query_string_control": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "The strategy for caching query strings. Can be IGNORE, CACHE_ALL or CUSTOM.",
		},
		"custom_cached_query_strings": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "List of custom cached query strings. Only visible if the queryStringControl attribute is CUSTOM.",
		},
		"dynamic_caching_by_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable dynamic caching by headers.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"custom_cached_headers": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "A list of custom cached headers. Only visible if dynamicCachingByHeaderEnabled is true.",
		},
		"gzip_compression_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable gzip compression.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"gzip_compression_level": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The level for the gzip compression. Values are between 1 to 6. Only visible is gzipCompressionEnabled is true.",
		},
		"content_persistence_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not make cached content available after its expiration time.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"maximum_stale_file_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The maximum time to live for stale files, in seconds. Only visible if contentPersistenceEnabled is true.",
		},
		"vary_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable honoring the vary header in a request.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"browser_cache_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Sets the default browser expiration time for cached assets, in seconds.",
		},
		"cors_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets the Access-Control-Allow-Origin header to allow browsers to access this domain from other origins.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"allowed_cors_origins": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The strategy for allowing cors origins. Can be SPECIFY_ORIGINS or ALL_ORIGINS.",
		},
		"origins_to_allow_cors": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "A list of origins to allow cors requests from. Only visible if allowedCorsOrigins is set to SPECIFY_ORIGINS.",
		},
		"http2_support_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable supporting applications using HTTP/2 protocol.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"http2_server_push_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to push assets to the client or browser (user) in advance (before the user requests these assets) which enables faster load times.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"link_header": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The link header for http2ServerPush, only visible if http2ServerPushEnabled is true.",
		},
		"canonical_header_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable setting Link: http://{hostname}/URI; rel=\"canonical\" header on each response.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"canonical_header": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The hostname for the canonicalHeader, only visible if canonicalHeaderEnabled is true.",
		},
		"url_caching_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether or not to enable caching of URLs without file extensions.",
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"url_caching_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The time to live for the url cache. Only visible if urlCachingEnabled is true.",
		},
	}
}

func getCDNPurgeResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "A site's unique identifier.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"purge_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "URL",
			ForceNew:    true,
			Description: "The type of cache purge. Can be URL or PATH. Default value is URL.",
		},
		"items": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Description: "The items to purge from the CDN.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"url": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "The URL or path at which to delete content.",
					},
					"recursive": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether or not to recursively delete content from the CDN.",
					},
					"invalidate_only": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether or not to mark the asset as expired and re-validate instead of deleting.",
					},
					"purge_all_dynamic": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Whether or not to purge dynamic versions of assets.",
					},
					"headers": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "A list of HTTP request headers used to construct a cache key to purge content by. These headers must be configured in the site configuration's DynamicContent.headerFields property.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"purge_selector": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "A key/value pair definition of content to purge from the CDN.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"selector_name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The name of the type of content to purge. For example, the name of the HTTP response header. Names are case sensitive.",
								},
								"selector_type": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The kinds of content that can be purged from the CDN. One of: HEADER (Purge content based on an HTTP response header), TAG (Purge content based on an X-TAG HTTP header value. Purging by tag can be useful when content on the origin is tagged).",
								},
								"selector_value": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The value of the content to purge. For example, the value of the HTTP response header. Values are case sensitive and may be wild-carded, but cannot match a \"/\".",
								},
								"selector_value_delimiter": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The delimiter to separate multiple values with. Defaults to \",\".",
								},
							},
						},
					},
				},
			},
		},
	}
}

func getWAFSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Environment name ",
		},
		"site_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"api_urls": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"ddos_settings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"global_threshold": {
						Type:     schema.TypeInt,
						Required: true,
					},
					"burst_threshold": {
						Type:     schema.TypeInt,
						Required: true,
					},
					"subsecond_burst_threshold": {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		"monitoring_mode_enabled": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
				var diags diag.Diagnostics
				value := i.(string)
				_, err := strconv.ParseBool(value)
				if err != nil {
					diag := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "wrong value",
						Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
					}
					diags = append(diags, diag)
				}
				return diags
			},
		},
		"owasp_threats": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"sql_injection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"xss_attack": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"shell_shock_attack": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"remote_file_inclusion": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"apache_struts_exploit": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"local_file_inclusion": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"common_web_application_vulnerabilities": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"webshell_execution_attempt": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"protocol_attack": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"csrf": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"open_redirect": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"shell_injection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"code_injection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sensitive_data_exposure": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"xml_external_entity": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"personal_identifiable_info": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"serverside_template_injection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"general_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"block_invalid_user_agents": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"block_unknown_user_agents": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"http_method_validation": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"traffic_sources": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"via_tor_nodes": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"via_proxy_networks": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"via_hosting_services": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"via_vpn": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"convicted_bot_traffic": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"traffic_from_suspicious_nat_ranges": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"external_reputation_block_list": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"traffic_via_cdn": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"anti_automation_bot_protection": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"force_browser_validation_on_traffic_anomalies": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"challenge_automated_clients": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"challenge_headless_browsers": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"anti_scraping": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"behavioral_waf": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"spam_protection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"block_probing_and_forced_browsing": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"obfuscated_attacks_and_zeroday_mitigation": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"repeated_violations": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"bruteforce_protection": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"cms_protection": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"wordpress_waf_ruleset": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_wordpress": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_modx": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_drupal": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_joomla": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_magento": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_origin_ip": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"whitelist_umbraco": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
		"allow_known_bots": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"acquia_uptime": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"add_search_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"adestra_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"adjust_servers": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ahrefs_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"alerta_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"alexa_ia_archiver": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"alexa_technologies": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"amazon_route_53_health_check_service": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"applebot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"apple_news_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ask_jeeves_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"audisto_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"baidu_spider_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"baidu_spider_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"binary_canary": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"bitbucket_webhook": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"blekko_scout_jet_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"chrome_compression_proxy": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"coccocbot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"cookie_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"cybersource": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"daumoa_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"detectify_scanner": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"digi_cert_dcv_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"dotmic_dot_bot_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"duck_duck_go_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"facebook_external_hit_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"feeder_co": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"feed_press": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"feed_wind": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"freshping_monitoring": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"geckoboard": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ghost_inspector": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"gomez": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"goo_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_ads_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_cloud_monitoring_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_feed_fetcher_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_image_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_image_proxy": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_mediapartners_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_mobile_ads_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_news_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_page_speed_insights": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_structured_data_testing_tool": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_verification_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_video_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"google_web_light": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"grapeshot_bot_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"gree_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"hetrix_tools": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"hi_pay": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"hyperspin_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ias_crawler_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"internet_archive_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"jetpack_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"jike_spider_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"j_word_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"kakao_user_agent": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"kyoto_tohoku_crawler": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"landau_media_spider": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"lets_encrypt": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"line_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"linked_in_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"livedoor_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"mail_ru_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"manage_wp": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_bing_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_bing_preview_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_msn_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"microsoft_skype_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"mixi_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"mobage_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"naver_yeti_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"new_relic_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"ocn_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"panopta_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"parse_ly_scraper": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"pay_pal_ipn": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"petal_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"pingdom": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"pinterest_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"qwantify_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"roger_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sage_pay": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sectigo_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"semrush_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"server_density_service_monitoring_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"seznam_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"shareaholic_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"site_24_x_7_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"siteimprove_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"site_lock_spider": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"slack_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sogou_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"soso_spider_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"spatineo": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"spring_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"stackify": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"status_cake_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"stripe": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"sucuri_uptime_monitor_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"telegram_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"testomato_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"the_find_crawler": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"twitter_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"uptime_robot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"vkontakte_external_hit_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"w_3_c": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"wordfence_central": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"workato": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"xml_sitemaps": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_inktomi_slurp_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_japan_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_link_preview": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_seeker_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yahoo_slurp_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yandex_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yisou_spider_commercial": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"yodao_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"zendesk_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"zoho_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
					"zum_bot": {
						Type:     schema.TypeString,
						Optional: true,
						ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
							var diags diag.Diagnostics
							value := i.(string)
							_, err := strconv.ParseBool(value)
							if err != nil {
								diag := diag.Diagnostic{
									Severity: diag.Error,
									Summary:  "wrong value",
									Detail:   fmt.Sprintf("%q is not %q", value, "Boolean value"),
								}
								diags = append(diags, diag)
							}
							return diags
						},
					},
				},
			},
		},
	}
}

func getFirewallRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique identifier for the rule.",
		},
		"environment_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the environment that the site belongs to.",
		},
		"site_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the site for which the firewall rule is applied to.",
		},
		"action": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Either ALLOW or BLOCK.",
		},
		"ip_start": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The start ip address for the rule. When no ipEnd attribute is provided, the rule only applies for the ip provided in ipStart.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the rule.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Whether or not the rule is enabled. The default value is false.",
		},
		"ip_end": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The end ip address for the rule.",
		},
	}
}

func getScriptSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"code": {
			Type:     schema.TypeString,
			Required: true,
		},
		"routes": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
	}
}
