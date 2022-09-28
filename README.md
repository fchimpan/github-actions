# github-actions
GitHub Actions Playground

## memo

### dispatch run
```bash
curl -X POST -H "Authorization: token xxxxx" \
    --data '{"event_type":"hello"}' \ 
    https://api.github.com/repos/fchimpan/github-actions/dispatches
```
