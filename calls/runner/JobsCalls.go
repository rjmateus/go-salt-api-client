package runner

/*
* More information
* https://docs.saltproject.io/en/latest/ref/runners/all/salt.runners.jobs.html
 */

func LookupJid(jid string) RunnerCall[map[string]interface{}] {
	var result = &map[string]interface{}{}
	args := make(map[string]interface{})
	args["jid"] = jid
	call := RunnerCall[map[string]interface{}]{"jobs.lookup_jid", args, result}
	return call
}

func Active() RunnerCall[map[string]interface{}] {
	var result = &map[string]interface{}{}
	call := RunnerCall[map[string]interface{}]{"jobs.active", nil, result}
	return call
}
