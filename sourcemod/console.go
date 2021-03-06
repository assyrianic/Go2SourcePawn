/**
 * sourcemod/console.go
 * 
 * Copyright 2020 Nirari Technologies, Alliedmodders LLC.
 * 
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 * 
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 * 
 */

package main


const INVALID_FCVAR_FLAGS = -1

type QueryCookie int
const QUERYCOOKIE_FAILED = QueryCookie(0)

type ReplySource int
const (
	SM_REPLY_TO_CONSOLE = ReplySource(0)
	SM_REPLY_TO_CHAT = ReplySource(1)
)

const (
	FCVAR_PLUGIN =           0       // Actual value is same as FCVAR_SS_ADDED in Left 4 Dead and later.
	FCVAR_LAUNCHER =         (1<<1)  // Same value as FCVAR_DEVELOPMENTONLY, which is what most usages of this were intending to use.


	FCVAR_NONE =                      0      // The default, no flags at all
	FCVAR_UNREGISTERED =             (1<<0)  // If this is set, don't add to linked list, etc.
	FCVAR_DEVELOPMENTONLY =          (1<<1)  // Hidden in released products. Flag is removed automatically if ALLOW_DEVELOPMENT_CVARS is defined. (OB+)
	FCVAR_GAMEDLL =                  (1<<2)  // Defined by the game DLL.
	FCVAR_CLIENTDLL =                (1<<3)  // Defined by the client DLL.
	FCVAR_MATERIAL_SYSTEM =          (1<<4)  // Defined by the material system. (EP1-only)
	FCVAR_HIDDEN =                   (1<<4)  // Hidden. Doesn't appear in find or autocomplete. Like DEVELOPMENTONLY, but can't be compiled out.1 (OB+)
	FCVAR_PROTECTED =                (1<<5)  // It's a server cvar, but we don't send the data since it's a password, etc.
                                       // Sends 1 if it's not bland/zero, 0 otherwise as value.
	FCVAR_SPONLY =                   (1<<6)  // This cvar cannot be changed by clients connected to a multiplayer server.
	FCVAR_ARCHIVE =                  (1<<7)  // Set to cause it to be saved to vars.rc
	FCVAR_NOTIFY =                   (1<<8)  // Notifies players when changed.
	FCVAR_USERINFO =                 (1<<9)  // Changes the client's info string.
	FCVAR_PRINTABLEONLY =            (1<<10) // This cvar's string cannot contain unprintable characters (e.g., used for player name, etc.)
	FCVAR_UNLOGGED =                 (1<<11) // If this is a FCVAR_SERVER, don't log changes to the log file / console if we are creating a log
	FCVAR_NEVER_AS_STRING =          (1<<12) // Never try to print that cvar.
	FCVAR_REPLICATED =               (1<<13) // Server setting enforced on clients.
	FCVAR_CHEAT =                    (1<<14) // Only useable in singleplayer / debug / multiplayer & sv_cheats
	FCVAR_SS =                       (1<<15) // causes varnameN where N  2 through max splitscreen slots for mod to be autogenerated (L4D+)
	FCVAR_DEMO =                     (1<<16) // Record this cvar when starting a demo file.
	FCVAR_DONTRECORD =               (1<<17) // Don't record these command in demo files.
	FCVAR_SS_ADDED =                 (1<<18) // This is one of the "added" FCVAR_SS variables for the splitscreen players (L4D+)
	FCVAR_RELEASE =                  (1<<19) // Cvars tagged with this are the only cvars available to customers (L4D+)
	FCVAR_RELOAD_MATERIALS =         (1<<20) // If this cvar changes, it forces a material reload (OB+)
	FCVAR_RELOAD_TEXTURES =          (1<<21) // If this cvar changes, if forces a texture reload (OB+)
	FCVAR_NOT_CONNECTED =            (1<<22) // Cvar cannot be changed by a client that is connected to a server.
	FCVAR_MATERIAL_SYSTEM_THREAD =   (1<<23) // Indicates this cvar is read from the material system thread (OB+)
	FCVAR_ARCHIVE_XBOX =             (1<<24) // Cvar written to config.cfg on the Xbox.
	FCVAR_ARCHIVE_GAMECONSOLE =      (1<<24) // Cvar written to config.cfg on the Xbox.
	FCVAR_ACCESSIBLE_FROM_THREADS =  (1<<25) // used as a debugging tool necessary to check material system thread convars (OB+)
	FCVAR_SERVER_CAN_EXECUTE =       (1<<28) // the server is allowed to execute this command on clients via
                                                   // ClientCommand/NET_StringCmd/CBaseClientState::ProcessStringCmd. (OB+)
	FCVAR_SERVER_CANNOT_QUERY =      (1<<29) // If this is set, then the server is not allowed to query this cvar's value (via
                                                   // IServerPluginHelpers::StartQueryCvarValue).
	FCVAR_CLIENTCMD_CAN_EXECUTE =    (1<<30) // IVEngineClient::ClientCmd is allowed to execute this command. 
                                                   // Note: IVEngineClient::ClientCmd_Unrestricted can run any client command.
)


func ServerCommand(format string, args ...any)
func ServerCommandEx(buffer []char, maxlen int, format string, args ...any)
func InsertServerCommand(format string, args ...any)
func ServerExecute()
func ClientCommand(client Entity, format string, args ...any)
func FakeClientCommand(client Entity, format string, args ...any)
func FakeClientCommandEx(client Entity, format string, args ...any)
func FakeClientCommandKeyValues(client Entity, kv KeyValues)
func PrintToServer(format string, args ...any)
func PrintToConsole(client Entity, format string, args ...any)
func PrintToConsoleAll(format string, args ...any)
func ReplyToCommand(client Entity, format string, args ...any)
func GetCmdReplySource() ReplySource
func SetCmdReplySource(source ReplySource) ReplySource
func IsChatTrigger() bool
func ShowActivity2(client Entity, tag, format string, args ...any)
func ShowActivity(client Entity, format string, args ...any)
func ShowActivityEx(client Entity, tag, format string, args ...any)
func FormatActivitySource(client, target int, namebuf string, maxlength int) bool

type (
	SrvCmd func(args int) Action
	ConCmd func(client, args int) Action
)
func RegServerCmd(cmd string, callback SrvCmd, description string, flags int)
func RegConsoleCmd(cmd string, callback ConCmd, description string, flags int)
func RegAdminCmd(md string, callback ConCmd, adminflags int, description, group string, flags int)
func GetCmdArgs() int
func GetCmdArg(argnum int, buffer []char, maxlength int) int
func GetCmdArgString(buffer []char, maxlength int) int

/// new CommandIterator();
type CommandIterator struct {
	Plugin Handle
	Flags int
}

func (CommandIterator) Next() bool
func (CommandIterator) GetDescription(buffer []char, maxlength int)
func (CommandIterator) GetName(buffer []char, maxlength int)


func GetCommandIterator() CommandIterator
func ReadCommandIterator(iter CommandIterator, name []char, nameLen int, eflags *int, desc []char, descLen int) bool
func CheckCommandAccess(client int, command string, flags int, override_only bool) bool
func CheckAccess(id AdminId, command string, flags int, override_only bool) bool
func GetCommandFlags(name string) int
func SetCommandFlags(name string, flags int) bool
func FindFirstConCommand(buffer []char, max_size int, isCommand *bool, flags *int, description []char, descrmax_size int) CommandIterator
func FindNextConCommand(search CommandIterator, buffer []char, max_size int, isCommand *bool, flags *int, description []char, descrmax_size int) bool
func AddServerTag(tag string)
func RemoveServerTag(tag string)

type CommandListener func(client Entity, command string, argc int) Action
const FEATURECAP_COMMANDLISTENER string = "command listener"

func AddCommandListener(callback CommandListener, command string) bool
func RemoveCommandListener(callback CommandListener, command string)
func CommandExists(command string) bool