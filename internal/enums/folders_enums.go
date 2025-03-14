package enums

const (
	CmdFolder        = "cmd"
	ApiFolder        = CmdFolder + "/api"
	HandlerFolder    = ApiFolder + "/handler"
	MiddlewareFolder = ApiFolder + "/middleware"
	RouterFolder     = ApiFolder + "/router"
	GroupsFolder     = RouterFolder + "/groups"
	ProvidersFolder  = CmdFolder + "/providers"
	ConfigFolder     = "config"
	CoreFolder       = "core"
	AdaptersFolder   = CoreFolder + "/adapters"
	AppFolder        = CoreFolder + "/app"
	DomainFolder     = CoreFolder + "/domain"
	ModelFolder      = DomainFolder + "/models"
	EntityFolder     = DomainFolder + "/entity"
	PortsFolder      = DomainFolder + "/ports"
)
