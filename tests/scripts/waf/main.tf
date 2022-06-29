terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "SyJwzYaeVylQmtOrPJkq2A=="
}

#do import existing resource, do run below script
#terraform import coxedge_waf_settings.testing <site_id>:<environment_name>

resource "coxedge_waf_settings" "testing" {
  site_id                 = "352cdc1e-c071-49ad-bddd-371094880507"
  environment_name        = "test-codecraft"
  domain                  = "cc.cox.com"
  monitoring_mode_enabled = "false"
  api_urls                = [
    "/test/this/ways123"
  ]
  ddos_settings {
    burst_threshold           = 110
    global_threshold          = 500
    subsecond_burst_threshold = 50
  }
  owasp_threats {
    apache_struts_exploit                  = "true"
    code_injection                         = "false"
    common_web_application_vulnerabilities = "false"
    csrf                                   = "true"
    local_file_inclusion                   = "false"
    open_redirect                          = "false"
    personal_identifiable_info             = "false"
    protocol_attack                        = "true"
    remote_file_inclusion                  = "false"
    sensitive_data_exposure                = "true"
    serverside_template_injection          = "true"
    shell_injection                        = "false"
    shell_shock_attack                     = "false"
    sql_injection                          = "false"
    webshell_execution_attempt             = "false"
    xml_external_entity                    = "true"
    xss_attack                             = "false"
  }
  general_policies {
    block_invalid_user_agents = "true"
    block_unknown_user_agents = "true"
    http_method_validation    = "false"
  }
  traffic_sources {
    convicted_bot_traffic              = "true"
    external_reputation_block_list     = "false"
    traffic_from_suspicious_nat_ranges = "false"
    traffic_via_cdn                    = "false"
    via_hosting_services               = "false"
    via_proxy_networks                 = "false"
    via_tor_nodes                      = "false"
    via_vpn                            = "true"
  }
  anti_automation_bot_protection {
    anti_scraping                                 = "true"
    challenge_automated_clients                   = "false"
    challenge_headless_browsers                   = "false"
    force_browser_validation_on_traffic_anomalies = "true"
  }
  behavioral_waf {
    block_probing_and_forced_browsing         = "false"
    bruteforce_protection                     = "true"
    obfuscated_attacks_and_zeroday_mitigation = "true"
    repeated_violations                       = "true"
    spam_protection                           = "true"
  }
  cms_protection {
    whitelist_drupal      = "false"
    whitelist_joomla      = "false"
    whitelist_magento     = "false"
    whitelist_modx        = "false"
    whitelist_origin_ip   = "false"
    whitelist_umbraco     = "false"
    whitelist_wordpress   = "false"
    wordpress_waf_ruleset = "true"
  }
  allow_known_bots {
    acquia_uptime                         = "false"
    add_search_bot                        = "false"
    adestra_bot                           = "false"
    adjust_servers                        = "false"
    ahrefs_bot                            = "false"
    alerta_bot                            = "false"
    alexa_ia_archiver                     = "true"
    alexa_technologies                    = "true"
    amazon_route_53_health_check_service  = "false"
    apple_news_bot                        = "false"
    applebot                              = "true"
    ask_jeeves_bot                        = "true"
    audisto_bot                           = "false"
    baidu_spider_bot                      = "true"
    baidu_spider_japan_bot                = "true"
    binary_canary                         = "false"
    bitbucket_webhook                     = "false"
    blekko_scout_jet_bot                  = "true"
    chrome_compression_proxy              = "true"
    coccocbot                             = "false"
    cookie_bot                            = "false"
    cybersource                           = "false"
    daumoa_bot                            = "true"
    detectify_scanner                     = "false"
    digi_cert_dcv_bot                     = "false"
    dotmic_dot_bot_commercial             = "false"
    duck_duck_go_bot                      = "false"
    facebook_external_hit_bot             = "true"
    feed_press                            = "false"
    feed_wind                             = "false"
    feeder_co                             = "false"
    freshping_monitoring                  = "false"
    geckoboard                            = "false"
    ghost_inspector                       = "false"
    gomez                                 = "false"
    goo_japan_bot                         = "true"
    google_ads_bot                        = "true"
    google_bot                            = "true"
    google_cloud_monitoring_bot           = "false"
    google_feed_fetcher_bot               = "true"
    google_image_bot                      = "true"
    google_image_proxy                    = "true"
    google_mediapartners_bot              = "true"
    google_mobile_ads_bot                 = "true"
    google_news_bot                       = "true"
    google_page_speed_insights            = "true"
    google_structured_data_testing_tool   = "true"
    google_verification_bot               = "true"
    google_video_bot                      = "true"
    google_web_light                      = "true"
    grapeshot_bot_commercial              = "false"
    gree_japan_bot                        = "true"
    hetrix_tools                          = "false"
    hi_pay                                = "true"
    hyperspin_bot                         = "false"
    ias_crawler_commercial                = "false"
    internet_archive_bot                  = "true"
    j_word_japan_bot                      = "true"
    jetpack_bot                           = "false"
    jike_spider_bot                       = "true"
    kakao_user_agent                      = "false"
    kyoto_tohoku_crawler                  = "false"
    landau_media_spider                   = "false"
    lets_encrypt                          = "false"
    line_japan_bot                        = "true"
    linked_in_bot                         = "true"
    livedoor_japan_bot                    = "false"
    mail_ru_bot                           = "false"
    manage_wp                             = "false"
    microsoft_bing_bot                    = "true"
    microsoft_bing_preview_bot            = "true"
    microsoft_msn_bot                     = "true"
    microsoft_skype_bot                   = "true"
    mixi_japan_bot                        = "true"
    mobage_japan_bot                      = "true"
    naver_yeti_bot                        = "true"
    new_relic_bot                         = "true"
    ocn_japan_bot                         = "true"
    panopta_bot                           = "false"
    parse_ly_scraper                      = "false"
    pay_pal_ipn                           = "true"
    petal_bot                             = "false"
    pingdom                               = "true"
    pinterest_bot                         = "false"
    qwantify_bot                          = "false"
    roger_bot                             = "false"
    sage_pay                              = "false"
    sectigo_bot                           = "false"
    semrush_bot                           = "false"
    server_density_service_monitoring_bot = "true"
    seznam_bot                            = "true"
    shareaholic_bot                       = "false"
    site_24_x_7_bot                       = "false"
    site_lock_spider                      = "true"
    siteimprove_bot                       = "false"
    slack_bot                             = "true"
    sogou_bot                             = "true"
    soso_spider_bot                       = "true"
    spatineo                              = "false"
    spring_bot                            = "false"
    stackify                              = "false"
    status_cake_bot                       = "false"
    stripe                                = "false"
    sucuri_uptime_monitor_bot             = "false"
    telegram_bot                          = "false"
    testomato_bot                         = "false"
    the_find_crawler                      = "false"
    twitter_bot                           = "true"
    uptime_robot                          = "false"
    vkontakte_external_hit_bot            = "true"
    w_3_c                                 = "false"
    wordfence_central                     = "false"
    workato                               = "false"
    xml_sitemaps                          = "false"
    yahoo_inktomi_slurp_bot               = "true"
    yahoo_japan_bot                       = "true"
    yahoo_link_preview                    = "true"
    yahoo_seeker_bot                      = "true"
    yahoo_slurp_bot                       = "true"
    yandex_bot                            = "true"
    yisou_spider_commercial               = "false"
    yodao_bot                             = "true"
    zendesk_bot                           = "false"
    zoho_bot                              = "false"
    zum_bot                               = "false"
  }
}