- `curl http://localhost:18000/health`
- `curl -i -XGET http://localhost:18000/tasks`
- `curl -i -XPOST http://localhost:18000/tasks -d @./handler/testdata/add_task/ok_req.json.golden`
- `curl -X -POST http://localhost:18000/register -d '{"name": "john", "password": "test", "role": "user"}'`
- `curl -XPOST http://localhost:18000/login -d '{"user_name": "john", "password": "test"}' | jq`
