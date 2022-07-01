
![logo](help/GoSaltLogo.png)

# go-salt-api-client

Go implementation of a salt api client.

## Modules
- Wheel: usable and with automatic test
- Runner: usable without automatic tests
- Modules: experimental

## Usage

The only authentication mechanism available is with username and password.

Login and Logout endpoint calls are not implemented.

## Salt Master Preparation

One needs to configure salt api on the salt master.
Information on how to do it can be [found on salt documentation](https://docs.saltproject.io/en/latest/ref/netapi/all/salt.netapi.rest_cherrypy.html#a-rest-api-for-salt)

## API usage examples

Simple call
```go
func main() {
	clientSalt := client.NewClient("host:port", "user", "pass", "auth_type")
	data, err := clientSalt.Run(client.LOCAL, "test.ping", target.NewTargetList([]string{"m43-minion-suse.tf.local"}), nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("data: >%s<\n", data)
}
```

Call usage examples can be found for:
- Wheel
  - [Key](example/wheel/keys_example.go)
  - [FileRoot](example/wheel/file_root_example.go)
- Runner
  - [Jobs](example/runner/jobs_example.go)

## How to contribute

- Fork the project
- Make any change you want
- Open a PR for this repository

## TODO List
- [] Better error handling: propagate error in each method 
- [] Implement login and logout endpoint, and make usage of the authentication token
- explore:
  - minions endpoint: https://docs.saltproject.io/en/latest/ref/netapi/all/salt.netapi.rest_cherrypy.html#minions
  - event BUS WS: https://docs.saltproject.io/en/latest/ref/netapi/all/salt.netapi.rest_cherrypy.html#ws
    - I'm not sure if this is useful. Auth method are different