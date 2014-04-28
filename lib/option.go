package inkfish

type Option struct {
	Host string `short:"o" long:"host" default:"127.0.0.1" description:"The interface a TCP based server daemon binds to. Default bind address is '127.0.0.1'."`
	Port uint   `short:"p" long:"port" default:"4979" description:"The port number a TCP based server daemon listens on. Defaults to 4979. This option doesn't mean anything if the server does not support TCP socket."`

	Server   string `short:"S" long:"Server" description:"IRC server address." required:"true"`
	Keyword  string `short:"K" long:"Keyword" description:"IRC server password."`
	Nickname string `short:"N" long:"Nickname" description:"IRC nickname."`
	User     string `short:"U" long:"User" description:"IRC user name."`

	Interval       uint `short:"i" long:"interval" default:"0" description:"IRC post interval to avoid excess flood."`
	Reconnect      uint `short:"R" long:"reconnect-interval" default:"0" description:"Interval of reconnect to irc server. exit application if interval == 0."`
	NoPostWithJoin bool `short:"j" long:"no-post-with-join" description:"Disable to irc message post with channel join"`
}
