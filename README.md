# protos

Protocol buffer definitions for the SLSA attestation predicates.

## Layout

```
<predicate>/<version>/<predicate>.proto
```

Each SLSA predicate lives under a directory named after the predicate,
versioned by a subdirectory matching the predicate's major version
(`v1`, `v2`, ...). Backwards-compatible additions land under the same major
version; breaking changes increment the version and use a new subdirectory.

## Predicates

| Predicate | Predicate type URI | Path
| --- | --- | ---
| Dependency Ingestion Provenance | `https://slsa.dev/dependency/v1` | [`dependency/v1/dependency.proto`](dependency/v1/dependency.proto)

Additional SLSA predicate schemas (Build Provenance, VSA, Source Provenance)
currently live in the [slsa-framework/slsa](https://github.com/slsa-framework/slsa)
repository under `spec/schema/` and are expected to migrate here over time.

## Conventions

- `syntax = "proto3";`
- Package name: `slsa.<predicate>.<version>` (e.g., `slsa.dependency.v1`).
- `go_package`: `github.com/slsa-framework/protos/<predicate>/<version>`.
- `java_package`: `dev.slsa.<predicate>.<version>`.
- Enumerable values are encoded as `string` fields rather than `enum` so
  implementer-defined values (scanner types, isolation methods) remain
  extensible without a proto change.
- Field-level validation and level-based requirements (e.g., "REQUIRED at
  L2+") are defined in the corresponding spec document, not in the proto.
  Validation of all fields is left to the users of these protos.

## Relationship to in-toto

SLSA predicates are wrapped in the
[in-toto attestation Statement](https://github.com/in-toto/attestation) v1
envelope. The Statement type itself is defined in
[in-toto/attestation](https://github.com/in-toto/attestation/tree/main/protos/in_toto_attestation/v1),
not here.
