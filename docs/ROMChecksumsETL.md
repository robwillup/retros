# ROM Checksums ETL

This is how you call the function that converts the checksum data in a local JSON
file to a YAML file:

```go
err := checksum.WriteChecksumsToYaml()
if err != nil {
    return err
}
```
