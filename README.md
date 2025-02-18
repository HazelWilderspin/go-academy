# Go Programming Exercise - To-Do App

The goal of the To Do Application is to build and evolve a production-lite service over four phases.

- [x] A command line application using an in-memory data store.
- [x] Introduce a REST API to wrap the data store and use JSON file(s) to provide data store persistence.
- [x] Add a Web App and make the To Do multi user.
- [ ] Use a DB to provide the data store persistence.

## Use of Go Packages

This program accompanies the Go Academy and therefore intention is to leverage the Go standard library. The exception to this are the following packages:

- [x] [github.com/google/uuid] - For working with UUID/GUID.
- [x] [github.com/google/go-cmp/cmp] - For comparing things useful for unit tests.
- [ ] [github.com/lib/pg] - a pure Go PostgreSQL driver.

## Development Approach

While developing the To-Do App use Git to store your solution and use Git Tags to mark final commit for each phase of the project.
As you progress through the project, make regular commits with a commit message documenting your progress.

## Phase Guidance

Each phase builds upon the previous phase and is expected to continue to work through all phases. For instance in phase 1 you build a CLI application, this application with small changes, should continue to work through phase 2, 3 and 4.

### Phase 1

- [x] CLI works directly with the In-Memory Data Store.

### Phase 2

- [x] Wrap the Data Store with the V1 REST API.
- [x] Use the [context] package to add a TraceID and [slog] to enable traceability of calls through the solution.
- [x] At the ToDo level, use CSP to support concurrent reads and concurrent safe write.
- [x] Use Parallel tests to validate that the solution is concurrent safe.
- [x] Update the CLI App to use the REST API.
- [x] Add an JSON Data Store and use a startup value to tell the REST API which data store to use.

### Phase 3

- [x] Add a V2 API to the REST API that supports multiple users.
- [x] Use Parallel test to validate the solution is concurrent safe across multiple users.
- [x] Add a Web API that supports multiple users.
- [x] Add multi-user support to the CLI App.

### Phase 4

- [ ] Add an additional data store implementation using an external DB (PostgreSQL).

[github.com/google/uuid]: https://pkg.go.dev/github.com/google/uuid
[github.com/google/go-cmp/cmp]: https://pkg.go.dev/github.com/google/go-cmp/cmp
[github.com/lib/pg]: https://pkg.go.dev/github.com/lib/pq
[context]: https://pkg.go.dev/context
[slog]: https://pkg.go.dev/log/slog
