## setup

- `make migrate`

## Tests

- `curl http://localhost:18000/health`
- `curl -i -XGET http://localhost:18000/tasks`
- `curl -i -XPOST http://localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden`
- `curl -X -POST http://localhost:18000/register -d '{"name": "john", "password": "test", "role": "user"}'`
- `curl -X -POST http://localhost:18000/register -d '{"name": "admin_user", "password": "test", "role": "admin"}'`
- `curl -XPOST http://localhost:18000/login -d '{"user_name": "john", "password": "test"}' | jq`
- `curl -XPOST http://localhost:18000/login -d '{"user_name": "admin_user", "password": "test"}' | jq`
