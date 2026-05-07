# Deployment

Production is isolated under `/opt/huyenlich` and uses host ports that are not used by existing VPS apps:

- Frontend: `127.0.0.1:3400`
- Backend: `127.0.0.1:8301`
- Postgres and Redis are private Docker services only.

Required DNS:

- `huyenlich.io.vn` A record -> `103.47.226.171`

After DNS resolves, enable HTTPS on the VPS:

```bash
certbot --nginx -d huyenlich.io.vn
```

GitHub Actions secrets for automatic deployment:

- `VPS_HOST`
- `VPS_USER`
- `VPS_SSH_KEY`
- `AI_API_KEY` optional
- `JWT_SECRET` optional, generated on first deploy when omitted

Optional repository variable:

- `VPS_PROJECT_PATH`, defaults to `/opt/huyenlich`
