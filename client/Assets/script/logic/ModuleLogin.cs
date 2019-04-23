

using System;

public class ModuleLogin : IModule
{
    private ModuleNetGameServ gameServer = null;

    public override bool init()
    {
        return true;
    }

    public override bool prestart()
    {
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(int delta)
    {
        
    }

    public void requestLogin(string token)
    {
        netproto.CS_Login login = new netproto.CS_Login
        {
            token = token,
        };
        gameServer.send(login);
    }

    [MsgRoute]
    public void response(netproto.SC_Login login)
    {
        SysLog.debug("login get code {0}", login.code);
    }
}