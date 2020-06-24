# MalGo - A set of simple go tools for pentesting.
My collection of malicious Go scripts, mostly modified examples of scripts from the excellent Blackhat Go by No Starch Press
# gowebscan.go
Web app testing automation tool, just to get nikto, testssl & gobuster running against multiple targets. You need to give it a file with hostnames, and if you want ips
Example syntax:
google.com:127.0.0.1
bing.com

example usage: 
gowebscan hosts.txt

# Execserv.go
A simple execution server that takes cmd input and executes it in bash.

# google.go
A web scraper based on the bing-metadata go scraper from bhg.

# dnsA.go
Takes the first sysarg and searches for it's A record. e.g. ./dnsA google.com

# vulnfuzz.go
Fuzzes vulnserver for overflow vulnerabilities.

# azurepillage.go
Incomplete script that pillages a filepath (azurepillage.exe [filepath]) for azure specific file extensions. It also reads files in the subdirectories to find possible json profiles. Use in smaller user directories etc as using in c:\ will eat a ton of RAM and give you a ton of potential files.
