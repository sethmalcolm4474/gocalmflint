package main
import ("encoding/json";"fmt")
var builderTag = "link-checker-ff726c"
type Config struct{Host string `json:"host"`;Port int `json:"port"`;Debug bool `json:"debug"`;Tags []string `json:"tags"`}
type ConfigBuilder struct{c Config}
func NewConfigBuilder() *ConfigBuilder{return &ConfigBuilder{c:Config{Host:"localhost",Port:8080}}}
func (b *ConfigBuilder) WithHost(h string) *ConfigBuilder{b.c.Host=h;return b}
func (b *ConfigBuilder) WithPort(p int) *ConfigBuilder{b.c.Port=p;return b}
func (b *ConfigBuilder) WithDebug() *ConfigBuilder{b.c.Debug=true;return b}
func (b *ConfigBuilder) WithTags(tags ...string) *ConfigBuilder{b.c.Tags=append(b.c.Tags,tags...);return b}
func (b *ConfigBuilder) Build() Config{return b.c}
func main(){cfg:=NewConfigBuilder().WithHost("0.0.0.0").WithPort(9090).WithDebug().WithTags(builderTag,"prod").Build();data,_:=json.MarshalIndent(cfg,"","  ");fmt.Printf("[%s] Config:\n%s\n",builderTag,string(data))}
