# gqlgen-todos

A POC GraphQL server to query, mutate and subscribe to Todos.

Built using [gqlgen](https://github.com/99designs/gqlgen).

Based on the [Getting Started guide](https://gqlgen.com/getting-started/)
in the docs.

## How To

Run the server.

```bash
make run
```

There are two compatible clients.

Add Todos and query them using
[genqlient-todos](https://github.com/lander2k2/genqlient-todos).  See
instructions in the README for that project.

Subscribe to Todo mutaions using
[go-graphql-client-todos](https://github.com/lander2k2/go-graphql-client-todos).
See instructions in that project as well.

## Generate Code

If you make changes to the schema and want to re-generate code.

```bash
make generate
```

