package main

import (
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Domain string `yaml:"domain"`
	Email  string `yaml:"email"`
	Ports  []int  `yaml:"ports"`
	Offset int    `yaml:"offset,omitempty"`
}

func (config *Config) validate() {
	if config.Domain == "" {
		logger.Panic("❌ Domain is empty in the configuration.")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(config.Email) {
		logger.Panic("❌ Invalid email format in the configuration.")
	}

	if len(config.Ports) == 0 {
		logger.Panic("❌ No ports specified in the configuration.")
	}

	if config.Offset == 0 {
		config.Offset = 7
	}

	logger.Println("✅ Config validations passed")
}

func readConfig() *Config {
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		logger.Panic("❌ Failed to read the configuration file:", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Panic("❌ Failed to unmarshal YAML:", err)
	}

	return &config
}

func (c *Config) displayPortMappingInfo() {
	logger.Print(`
	
	.S_sSSs      sSSs_sSSs     .S_sSSs    sdSS_SSSSSSbs                        
	.SS~YS%%b    d%%SP~YS%%b   .SS~YS%%b   YSSS~S%SSSSSP                        
	S%S    S%b  d%S'      S%b  S%S    S%b       S%S                             
	S%S    S%S  S%S       S%S  S%S    S%S       S%S                             
	S%S    d*S  S&S       S&S  S%S    d*S       S&S                             
	S&S   .S*S  S&S       S&S  S&S   .S*S       S&S                             
	S&S_sdSSS   S&S       S&S  S&S_sdSSS        S&S                             
	S&S~YSSY    S&S       S&S  S&S~YSY%b        S&S                             
	S*S         S*b       d*S  S*S    S%b       S*S                             
	S*S         S*S.     .S*S  S*S    S%S       S*S                             
	S*S          SSSbs_sdSSS   S*S    S&S       S*S                             
	S*S           YSSP~YSSY    S*S    SSS       S*S                             
	SP                         SP               SP                              
	Y                          Y                Y                               
																				
	 .S_SsS_S.    .S_SSSs     .S_sSSs     .S_sSSs     .S   .S_sSSs      sSSSSs  
	.SS~S*S~SS.  .SS~SSSSS   .SS~YS%%b   .SS~YS%%b   .SS  .SS~YS%%b    d%%%%SP  
	S%S  Y  S%S  S%S   SSSS  S%S    S%b  S%S    S%b  S%S  S%S    S%b  d%S      
	S%S     S%S  S%S    S%S  S%S    S%S  S%S    S%S  S%S  S%S    S%S  S%S       
	S%S     S%S  S%S SSSS%S  S%S    d*S  S%S    d*S  S&S  S%S    S&S  S&S       
	S&S     S&S  S&S  SSS%S  S&S   .S*S  S&S   .S*S  S&S  S&S    S&S  S&S       
	S&S     S&S  S&S    S&S  S&S_sdSSS   S&S_sdSSS   S&S  S&S    S&S  S&S       
	S&S     S&S  S&S    S&S  S&S~YSSY    S&S~YSSY    S&S  S&S    S&S  S&S sSSs  
	S*S     S*S  S*S    S&S  S*S         S*S         S*S  S*S    S*S  S*b   S%%  
	S*S     S*S  S*S    S*S  S*S         S*S         S*S  S*S    S*S  S*S   S%  
	S*S     S*S  S*S    S*S  S*S         S*S         S*S  S*S    S*S   SS_sSSS  
	SSS     S*S  SSS    S*S  S*S         S*S         S*S  S*S    SSS    Y~YSSY  
			SP          SP   SP          SP          SP   SP                    
			Y           Y    Y           Y           Y    Y                     
																				
	`)

	for _, p := range c.Ports {
		logger.Printf("\n Service port: %s -> Host port: %s \n", p, p+c.Offset)
	}

	logger.Println()
}
