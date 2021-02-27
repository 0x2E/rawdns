package rawdns

// Marshal return the DNS content can be used in UDP packet
func Marshal(id uint16, rd byte, addr string, qtype QType) ([]byte, error) {
	queryMsg := &Message{
		Header: Header{
			ID:      id,
			QR:      0,
			Opcode:  OpcodeQuery,
			QDCOUNT: 1,
			RD:      rd,
		},
		Questions: []*Question{
			{
				QNAME:  addr,
				QTYPE:  qtype,
				QCLASS: QClassIN,
			},
		},
	}
	payload, err := queryMsg.Marshal()
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// Unmarshal return the parsed DNS response in UDP packet
func Unmarshal(raw []byte) (*Message, error) {
	responseMsg := &Message{}
	err := UnmarshalMessage(raw, responseMsg)
	if err != nil {
		return nil, err
	}

	return responseMsg, nil
}
