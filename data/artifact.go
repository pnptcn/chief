package data

import (
	"fmt"

	capnp "capnproto.org/go/capnp/v3"
)

var (
	vMajor = 1
	vMinor = 0
	vPatch = 0

	Empty = Artifact{}
)

func New(mediaType MediaType, role RoleType) *Artifact {
	var (
		artifact Artifact
		seg      *capnp.Segment
		err      error
	)

	if _, seg, err = capnp.NewMessage(capnp.SingleSegment(nil)); err != nil {
		return nil
	}

	if artifact, err = NewRootArtifact(seg); err != nil {
		return nil
	}

	return &artifact
}

func (artifact *Artifact) Prefix() (prfx string, err error) {
	var (
		header Artifact_Header
		role   string
	)

	if header, err = artifact.Header(); err != nil {
		return
	}

	if role, err = header.Role(); err != nil {
		return
	}

	prfx = fmt.Sprintf(
		"%s/%s/%s/%d",
		role,
	)

	return
}
