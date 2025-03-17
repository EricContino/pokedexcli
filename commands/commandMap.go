package commands

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

func CommandMap(config *Config) error {
    res, err := http.Get(config.Next)
    if err != nil {
        return fmt.Errorf("Error getting map: %v", err)
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        return fmt.Errorf("Error reading body: %v", err)
    }
    
    mapResp := MapResponse{}
    err = json.Unmarshal(body, &mapResp)
    if err != nil {
        return fmt.Errorf("Error unmarshalling: %v", err)
    }
    
    for _, result := range mapResp.Results {
        fmt.Println(result.Name)
    }
    fmt.Println(mapResp.Previous)
    if mapResp.Previous == nil {
        fmt.Println("1")
        config.Previous = PokeApi_LocationAreas_BaseURL
    } else {
        fmt.Println("2")
        if s, ok := mapResp.Previous.(string); ok {
            fmt.Println("3")
            config.Previous = s
        }
    }
    config.Next = mapResp.Next
    
    return nil
}

func CommandMapb(config *Config) error {
    res, err := http.Get(config.Previous)
    if err != nil {
        return fmt.Errorf("Error getting map: %v", err)
    }

    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        return fmt.Errorf("Error reading body: %v", err)
    }
    
    mapResp := MapResponse{}
    err = json.Unmarshal(body, &mapResp)
    if err != nil {
        return fmt.Errorf("Error unmarshalling: %v", err)
    }
    
    for _, result := range mapResp.Results {
        fmt.Println(result.Name)
    }
    
    if mapResp.Previous == nil {
        config.Previous = PokeApi_LocationAreas_BaseURL
    } else {
        if s, ok := mapResp.Previous.(string); ok {
            config.Previous = s
        }
    }
    config.Next = mapResp.Next
    
    return nil
}
