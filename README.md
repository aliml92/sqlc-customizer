# sqlc-customize
temporary solution to customize sqlc-generated code

First create `sqlc-customize.yaml` in the current folde. For e.g.
```yaml
modify:
  model:
    - name: "User"
      source: "./internal/user/repository/postgresql/models.go"
      destination: "./internal/models/user.go"
      package: "models"
      old_package: "db"
      package_path: "github.com/aliml92/{project_name}/internal/models"
      json_omitempty: true
```      
Since it changes sqlc code, it is better to run right after `sqlc generate` command:
```yaml
version: '3'

tasks:
  sqlc-gen:
    desc: "build the compiled binary"
    cmds:
      - sqlc generate
      - sqlc-customizer modify
```

For more understanding, see [examples](https://github.com/aliml92/sqlc-customizer/tree/cb417ddf1e3913bad8c110bbf779093b76c02727/examples/clean-arch)
