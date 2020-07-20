# MalGo - A set of simple go tools for pentesting.
My collection of malicious Go scripts, mostly modified examples of scripts from the excellent Blackhat Go by No Starch Press.
Not all of it is necessarily evil, most of it is just useful for security testing.
# goscraper.go

A simple webscraper for pentesteracademy videos, it requires you to have a valid account - so [please subscribe to their excellent service](https://www.pentesteracademy.com/benefits), I really recommend the Linux Assembley and AD courses! 

Usage:

```sh
go run goscraper.go [courseid] [cookie]

go run goscraper.go 11 eW91cmNvb2tpZWdvZXNoZXJlLCBpdCdzIHRoZSBzdHVmZiBhZnRlciAiU0FDU0lEPSI=

```

# Execserv.go
A simple execution server that takes cmd input and executes it in bash. It's meant to be run on the target so you have a web server you can revisit to get your initial shell. Sort of think of it like a template for webserver backdoor, because it's sort of rubbish in it's current state.

# google.go
A web scraper based on the bing-metadata go scraper from bhg.

# dnsA.go
Takes the first sysarg and searches for it's A record. e.g. ./dnsA google.com - Could be handy to expand to search for other records, a handy starting point basically.

# vulnfuzz.go
Fuzzes vulnserver for overflow vulnerabilities. This might be helpful for finding other overflow vulns, but you need to modify it.

# azurepillage.go
Incomplete script that pillages a filepath (azurepillage.exe [filepath]) for azure specific file extensions. It also reads files in the subdirectories to find possible json profiles. Use in smaller user directories etc as using in c:\ will eat a ton of RAM and give you a ton of potential files.
