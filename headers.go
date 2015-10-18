package main

import "github.com/stefankopieczek/gossip/base"

// Utility methods for creating headers.

func Via(e *endpoint, branch string) *base.ViaHeader {
	return &base.ViaHeader{
		&base.ViaHop{
			ProtocolName:    "SIP",
			ProtocolVersion: "2.0",
			Transport:       e.transport.String(),
			Host:            e.host,
			Port:            &e.port,
			Params:          base.NewParams().Add("branch", base.String{branch}),
		},
	}
}

func To(e *endpoint, tag base.String) *base.ToHeader {
	header := &base.ToHeader{
		DisplayName: &e.displayName,
		Address: &base.SipUri{
			User:      &e.username,
			Host:      e.host,
			UriParams: base.NewParams(),
		},
		Params: base.NewParams(),
	}

	if tag.String() != "" {
		header.Params.Add("tag", &tag)
	}

	return header
}

func From(e *endpoint, tag base.String) *base.FromHeader {
	header := &base.FromHeader{
		DisplayName: &e.displayName,
		Address: &base.SipUri{
			User:      &e.username,
			Host:      e.host,
			UriParams: base.NewParams().Add("transport", &e.transport),
		},
		Params: base.NewParams(),
	}

	if tag.String() != "" {
		header.Params.Add("tag", &tag)
	}

	return header
}

func Contact(e *endpoint) *base.ContactHeader {
	return &base.ContactHeader{
		DisplayName: &e.displayName,
		Address: &base.SipUri{
			User: &e.username,
			Host: e.host,
		},
	}
}

func CSeq(seqno uint32, method base.Method) *base.CSeq {
	return &base.CSeq{
		SeqNo:      seqno,
		MethodName: method,
	}
}

func CallId(callid string) *base.CallId {
	header := base.CallId(callid)
	return &header
}

func ContentLength(l uint32) base.ContentLength {
	return base.ContentLength(l)
}
