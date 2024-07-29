package data

import capnp "capnproto.org/go/capnp/v3"

func NewHeader(mediaType MediaType, role RoleType) (header Artifact_Header) {
	var (
		seg *capnp.Segment
		err error
	)

	if _, seg, err = capnp.NewMessage(capnp.SingleSegment(nil)); err != nil {
		return
	}

	if header, err = NewRootArtifact_Header(seg); err != nil {
		return
	}

	if err = header.SetType(string(mediaType)); err != nil {
		return
	}

	if err = header.SetRole(string(role)); err != nil {
		return
	}

	return header
}
