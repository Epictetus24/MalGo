package main

import (
        "fmt"
        "log"
        "os"
        "path/filepath"
        "regexp"
)

var regexes = []*regexp.Regexp{
        regexp.MustCompile(`(?i)login`),
        regexp.MustCompile(`(?i)pfx`),
        regexp.MustCompile(`(?i)cer`),
        regexp.MustCompile(`(?i)publishsettings`),
        regexp.MustCompile(`(?i)cspkg`),
        regexp.MustCompile(`(?i)config`),
}

var azurestrings = []*regexp.Regexp{
        regexp.MustCompile(`(?i)TokenCache`),
        regexp.MustCompile(`(?i)Tenant`),
        regexp.MustCompile(`(?i)PublishSettingsFileUrl`),
        regexp.MustCompile(`(?i)ManagementPortalUrl`),
        regexp.MustCompile(`(?i)SAM`),
}


func walkFn(path string, f os.FileInfo, err error) error {
        for _, r := range regexes {
                if r.MatchString(path) {
                        fmt.Printf("[+] HIT %s : %s\n", r, path)
                }
        }
        return nil
}

//To add - func for reading file in path and searching for the "azurestrings" regex.

func main() {
        root := os.Args[1]
        if err := filepath.Walk(root, walkFn); err != nil {
                log.Panicln(err)
        }
}
