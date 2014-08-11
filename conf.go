package conf

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type Configs map[string]string

func NewConfigs(path string) (Configs, error) {
	configs := Configs(map[string]string{})
	if path != "" {
		err := configs.Parse(path)
		return configs, err
	} else {
		return configs, nil
	}
}

func (c Configs) Parse(path string) error {
        f, err := os.Open(path)
        if err != nil {
                return err
        }
        scanner := bufio.NewScanner(f)
        for scanner.Scan() {
                line := scanner.Text()
                line = strings.TrimSpace(line)
                // Skip comments and sections
                // At some point we may wish to
                // read only [client] section
                if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "[") {
                        continue
                }
                option := strings.Split(line, "=")
                //log.Println(option) => [user sakai]
                if len(option) == 2 {
			c[strings.ToLower(option[0])] = option[1]
                } else if len(option) > 2 {
			c[strings.ToLower(option[0])] = strings.Join(option[1:], "=")
		}
        }
	if err := scanner.Err(); err != nil { log.Println(err) }
        return nil
}
