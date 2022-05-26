package target

type TargetType string

const (
	List TargetType = "list"
	GLOB            = "glob"
	//	PCRE = "pcre"
	GRAIN = "grain"
	//	GRAIN_PCRE   = "grain_pcre"
	//	PILLAR       = "pillar"
	//	PILLAR_EXACT = "pillar_exact"
	//	PILLAR_PCRE  = "pillar_pcre"
	//	NODEGROUP = "nodegroup"
	//	RANGE     = "range"
	//	COMPOUND  = "compound"
	//	IPCIDR    = "ipcidr"
	//	DATA      = "data"
)

const (
	defaultDelimiter = ":"
)

//TODO allow different delimiters for dictionary targets

type SaltTarget struct {
	target     interface{}
	targetType TargetType
}

func (tgt *SaltTarget) GetPros() map[string]interface{} {
	props := make(map[string]interface{})
	props["tgt"] = tgt.target
	props["tgt_type"] = tgt.targetType
	return props
}
