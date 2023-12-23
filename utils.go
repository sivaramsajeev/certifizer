package main

func printConfigFileUsage() {
	logger.Println(`
Create a config file by the name certifizer.yml either in your home directory or any specific file path.
NOTE: You need to set CERTIFIZER_CONFIG_PATH environment variable if config file is not in your home directory

SAMPLE FILE:
------------
cat ~/certifizer.yml
domain: 
email:
# offset: 7
	
ports:
- 3000
- 5000
	`)
}
