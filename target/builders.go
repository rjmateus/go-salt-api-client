package target

func NewTargetList(tgt []string) SaltTarget {
	return SaltTarget{tgt, List}
}

func NewTargetGlob(tgt string) SaltTarget {
	return SaltTarget{tgt, GLOB}
}

func NewTargetGrains(key, value string) SaltTarget {
	return SaltTarget{key + defaultDelimiter + value, GRAIN}
}
