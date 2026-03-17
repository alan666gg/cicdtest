# GitHub CI/CD Setup Checklist

- 项目目录：`/root/.openclaw/workspace/hello-cicd-demo`
- 服务数量：`1`
- 服务路径：
  - `.`
- 应用名称：`hello-cicd-demo`
- 部署策略：`docker-ssh`
- 测试分支：`develop`

## 必填 Secrets

- `TEST_SSH_KEY`
- `PROD_SSH_KEY`

## 必填 Variables

- `TEST_HOST`
- `TEST_USER`
- `PROD_HOST`
- `PROD_USER`
- `TEST_REMOTE_DIR`
- `TEST_CONTAINER_NAME`
- `TEST_DOCKER_RUN_ARGS`
- `PROD_REMOTE_DIR`
- `PROD_CONTAINER_NAME`
- `PROD_DOCKER_RUN_ARGS`

## 可选 Variables

- `TEST_PORT`
- `PROD_PORT`

## 组织级推荐配置

建议在 `.github/cicd-bootstrap.json` 里统一这些字段：

- `default_branch` / `default_branches`
- `test_branch` / `test_branches`
- `deploy_strategy`
- `service_path` / `service_paths`
- `image_registry`
- `runner`
- `enable_security_scan`
- `security_scan_blocking`
- `enable_cache`
- `test_environment`
- `prod_environment`

## 使用说明

1. 先把 Secrets 和 Variables 配齐。
2. 推送到测试分支，观察 `CI` 和 `Deploy Test` 工作流。
3. 确认测试环境没问题后，再手动触发 `Deploy Prod`。

## 安全扫描默认策略

- 默认开启安全扫描，但 bootstrap 模式下是 non-blocking。
- PR 和 `develop` 这类测试分支默认只告警，不阻断流水线。
- 如果把 `.github/cicd-bootstrap.json` 里的 `security_scan_blocking` 设为 `true`，推送到默认分支和 `release` / `release/*` 分支时会对 `HIGH` / `CRITICAL` 漏洞阻断。

## 备注

- 如果 workflow 文件需要覆盖已有文件，请在执行 bootstrap 时加 `--force`。
- 如果是 monorepo，建议优先使用 `--service-path` 或 `--service-paths`。
- 如果是团队统一模板，优先把默认值写进 `.github/cicd-bootstrap.json`，减少每次手输参数。
