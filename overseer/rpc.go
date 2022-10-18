package main

type OverseerRpc struct{}

func (ov *OverseerRpc) Top(_ any, out *string) error {
	*out = TopSup.String()
	return nil
}
