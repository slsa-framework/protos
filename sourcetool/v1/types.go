// SPDX-FileCopyrightText: Copyright 2026 The SLSA Authors
// SPDX-License-Identifier: Apache-2.0

package v1

// Predicate types of the statements carrying the source track predicates.
const (
	// PredicateTypeSourceProvenance is the predicate type of statements
	// carrying a SourceProvenancePred.
	PredicateTypeSourceProvenance = "https://github.com/slsa-framework/source-tool/source-provenance/v1"
	// PredicateTypeTagProvenance is the predicate type of statements
	// carrying a TagProvenancePred.
	PredicateTypeTagProvenance = "https://github.com/slsa-framework/source-tool/tag-provenance/v1"

	// PredicateTypeSourceProvenanceDraft is the predicate type source-tool
	// used for source provenance before the promotion to v1. The payload is
	// the same SourceProvenancePred; readers should keep accepting it, as
	// attestations already issued under it remain valid.
	PredicateTypeSourceProvenanceDraft = "https://github.com/slsa-framework/slsa-source-poc/source-provenance/v1-draft"
	// PredicateTypeTagProvenanceDraft is the pre-v1 predicate type of
	// TagProvenancePred statements.
	PredicateTypeTagProvenanceDraft = "https://github.com/slsa-framework/slsa-source-poc/tag-provenance/v1-draft"
)

// SourceProvenancePredicateTypes lists every predicate type whose payload is
// a SourceProvenancePred, newest first.
var SourceProvenancePredicateTypes = []string{
	PredicateTypeSourceProvenance,
	PredicateTypeSourceProvenanceDraft,
}

// TagProvenancePredicateTypes lists every predicate type whose payload is a
// TagProvenancePred, newest first.
var TagProvenancePredicateTypes = []string{
	PredicateTypeTagProvenance,
	PredicateTypeTagProvenanceDraft,
}
