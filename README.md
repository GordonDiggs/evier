# evier

## Configuration

You can specify a yaml file to configure the backup process:

```yaml
integrations:
  slack:
    url: "https://..."
rsync:
  user: zaphod
  host: deep-thought.local
  path: "/volume/backup"
  rsh: "ssh ..."
  delete: true
jobs:
  - name: "Photos"
    source: "/mnt/photos"
    destination: "photos"
  - name: "Documents"
    source: "/mnt/docs"
    destination: "docs"
```
