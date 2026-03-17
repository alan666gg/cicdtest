# GitHub CI/CD Setup Checklist

- 项目目录：`/root/.openclaw/workspace/hello-cicd-demo`
- 服务数量：`1`
- 服务路径：
  - `.`
- 应用名称：`hello-cicd-demo`
- 部署策略：`docker-registry-only`
- 测试分支：`develop`

## 必填 Secrets

- `REGISTRY_USERNAME`
- `REGISTRY_PASSWORD`

## 必填 Variables

- 无（默认使用 repo config 里的 `image_registry`）

## 可选 Variables

- `IMAGE_REGISTRY`（覆盖 repo config 里的 registry 前缀，workflow 会自动转成小写）

## 组织级推荐配置

建议在 `.github/cicd-bootstrap.json` 里统一这些字段：

- `default_branch` / `default_branches`
- `test_branch` / `test_branches`
- `deploy_strategy`
- `service_path` / `service_paths`
- `image_registry`
- `runner`
- `enable_security_scan`
- `enable_cache`
- `test_environment`
- `prod_environment`

## 使用说明

1. 先确认镜像仓库前缀是否正确，例如 `ghcr.io/acme-team`。
2. 如果 registry 前缀放在 `IMAGE_REGISTRY` 变量里，workflow 会自动把前缀转成小写并复用对应 host 登录。
3. 配置 `REGISTRY_USERNAME` 和 `REGISTRY_PASSWORD`。
4. 推送到测试分支，确认测试镜像已经成功推送。
5. 再手动触发生产 workflow，推送生产标签。

## 备注

- 如果 workflow 文件需要覆盖已有文件，请在执行 bootstrap 时加 `--force`。
- 如果是 monorepo，建议优先使用 `--service-path` 或 `--service-paths`。
- 如果是团队统一模板，优先把默认值写进 `.github/cicd-bootstrap.json`，减少每次手输参数。
