# gh-sql

GitHub SQL is a tool to fetch data from the GitHub API and store it in a SQL
database. The ultimate goal is to have the data locally to perform queries on
the dataset.

The tool has the following focuses:

- Being targeted towards storing information about specific repositories.
	- Discovery of users or repositories is a non-goal.
- Being able to incrementally update the database, fetching what has changed
	in a repository using the events API (either through polling, even in long
	intervals (up to 1 month) or through WebSockets).
- Being able to easily "saved" in a single SQL file for quick recovery and
	replication.
- Replicating the GitHub REST API in read-only operations, without rate limiting
	(for integration with existing tools that use the GitHub API, by only
	changing the API endpoint.)

## Notes

- ffcli for flag parsing / configuration
- ent for ORM
- start out by syncing single Repository entity, and work from there
